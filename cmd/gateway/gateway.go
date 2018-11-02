package main

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/youngderekm/grpc-cookies-example/servicedef"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"encoding/gob"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/romana/rlog"
	"google.golang.org/grpc/metadata"
)

// Session is the structure holding session data that gorilla's session
// store will convert into a cookie.
// In this example we only set the UserID, but other session data could
// be handled the same way.
type Session struct {
	// Our user ID
	UserID int
}

// our type for context keys
type contextKey int

const (
	// name of session we're setting
	defaultSessionID = "grpc-cookies-example-session"

	// secret key used to encrypt session
	cookieStorageKey = "asdflkjasflkjasldfkjs"

	// How long our profile session can last, in seconds, unless renewed.
	sessionLength = 60 * 60 * 4

	// our unique key used for storing the request in the context
	requestContextKey contextKey = 0
)

// Global session store
var sessionStore sessions.Store

// Middleware that will store the http.Request into the Context
type gatewayMiddleware struct {
}

// the gatewayResponseModifier function needs the request object to
// be able to use gorilla's session stores correctly.  the only way to pass it on
// is through the context, so store it here
func (middleware *gatewayMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		ctx = context.WithValue(ctx, requestContextKey, r)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// pull the request from context (set in middleware above)
func getRequestFromContext(ctx context.Context) *http.Request {
	return ctx.Value(requestContextKey).(*http.Request)
}

// get the first metadata value with the given name
func firstMetadataWithName(md runtime.ServerMetadata, name string) string {
	values := md.HeaderMD.Get(name)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

// extract the user ID from the server meta data, if available.  user ID will be 0 if
// it wasn't set in the metadata
func getUserIDFromServerMetadata(md runtime.ServerMetadata) (int, error) {
	userIDString := firstMetadataWithName(md, "gateway-session-userId")
	if userIDString == "" {
		return 0, nil
	}
	if userIDString != "" {
		userID, err := strconv.Atoi(userIDString)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}
	return 0, nil
}

// return a boolean from the server metadata, or a default value if the named
// field does not exist in the metadata
func getBoolFromServerMetadata(md runtime.ServerMetadata, name string, defaultValue bool) (bool, error) {
	boolString := firstMetadataWithName(md, name)
	if boolString != "" {
		value, err := strconv.ParseBool(boolString)
		if err != nil {
			return defaultValue, err
		}
		return value, nil
	}
	return defaultValue, nil
}

// look to see if the gRPC method we called tells us to do something special with the session
// (create, update, delete).  If so, take that action here.
// see
//   https://github.com/grpc-ecosystem/grpc-gateway/blob/master/docs/_docs/customizingyourgateway.md
func gatewayResponseModifier(ctx context.Context, response http.ResponseWriter, _ proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return fmt.Errorf("Failed to extract ServerMetadata from context")
	}
	// did the gRPC method set a user ID in the metadata?
	userID, err := getUserIDFromServerMetadata(md)
	if err != nil {
		return err
	}

	if userID != 0 {
		rlog.Debugf("gRPC call set userId to %d", userID)

		// pull the request from context (set in middleware above)
		request := getRequestFromContext(ctx)

		// create or get the session
		session, err := sessionStore.New(request, defaultSessionID)
		if err != nil {
			rlog.Error(err, "couldn't create a session")
			return err
		}
		session.Options.MaxAge = sessionLength
		session.Options.Path = "/"

		// create a session for the user.  This session is converted by gorilla
		// into a session cookie
		userIDSession := &Session{
			UserID: userID,
		}

		// put the userId into session
		session.Values["userId"] = userIDSession
		// save the session, creating a cookie from it
		if err := sessionStore.Save(request, response, session); err != nil {
			rlog.Error(err, "couldn't save the session as a cookie")
			return err
		}
	}

	// did the gRPC method called set a flag telling us to delete the session?
	deleteSession, err := getBoolFromServerMetadata(md, "gateway-session-delete", false)
	if err != nil {
		return err
	}
	if deleteSession {
		// pull the request from context (set in middleware above)
		r := getRequestFromContext(ctx)

		// as documented, to delete session, set max age to -1
		session, err := sessionStore.New(r, defaultSessionID)
		if err != nil {
			rlog.Error(err, "couldn't create empty session")
			return err
		}
		session.Options.MaxAge = -1
		session.Options.Path = "/"
		// "save" the session with maxage = -1, clearing it
		if err := sessionStore.Save(r, response, session); err != nil {
			rlog.Error(err, "couldn't delete session")
			return err
		}
	}

	return nil
}

// look up session and pass userId in to context if it exists
func gatewayMetadataAnnotator(_ context.Context, r *http.Request) metadata.MD {
	session, err := sessionStore.Get(r, defaultSessionID)
	if err != nil {
		// no session, or invalid session, so pass along no extra metadata
		return metadata.Pairs()
	}
	if userIDSessionValue, ok := session.Values["userId"]; ok {
		// convert back to a Session
		userIDSession := userIDSessionValue.(*Session)
		userID := userIDSession.UserID
		// set user ID from session in the gRPC metadata
		return metadata.Pairs("userId", strconv.Itoa(userID))
	}
	// otherwise pass no extra metadata along
	return metadata.Pairs()
}

func serve() error {
	const wait = time.Second * 15

	// session handling
	gob.Register(&Session{})

	cookieStore := sessions.NewCookieStore([]byte(cookieStorageKey))
	cookieStore.Options = &sessions.Options{HttpOnly: true}
	sessionStore = cookieStore

	r := mux.NewRouter()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// gRPC endpoint the gateway will be calling
	grpcEndpoint := "localhost:50051"

	// use our hook to modify the response after the gRPC call comes back
	gwmux := runtime.NewServeMux(runtime.WithForwardResponseOption(gatewayResponseModifier), runtime.WithMetadata(gatewayMetadataAnnotator))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := servicedef.RegisterAuthApiHandlerFromEndpoint(ctx, gwmux, grpcEndpoint, opts)
	if err != nil {
		return err
	}

	loggingHandler := handlers.CombinedLoggingHandler(os.Stderr, gwmux)

	middleware := gatewayMiddleware{}
	r.Use(middleware.Middleware)

	r.PathPrefix("/").Handler(loggingHandler)

	gatewayAddr := "localhost:8081"
	srv := &http.Server{
		Addr: gatewayAddr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		rlog.Infof("Listening for JSON requests at %s", gatewayAddr)
		if err := srv.ListenAndServe(); err != nil {
			rlog.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(shutdownCtx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	rlog.Info("shutting down gRPC gateway")
	os.Exit(0)
	return nil
}

func main() {
	// log at debug for this demo
	os.Setenv("RLOG_LOG_LEVEL", "DEBUG")
	rlog.UpdateEnv()

	if err := serve(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
