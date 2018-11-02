package main

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/romana/rlog"
	"github.com/youngderekm/grpc-cookies-example/servicedef"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SignIn demonstrates setting a flag that tells the gateway to set a cookie.
// In this example we set the user ID after a simulated login event.
func (s *server) SignIn(ctx context.Context, signInRequest *servicedef.SignInRequest) (*empty.Empty, error) {
	// simulate logging in for the purposes of the demo
	if signInRequest.Username == "tester" && signInRequest.Password == "1234" {
		// normally a user ID would come from a database or other external service, but
		// for this demo we'll hardcode it
		userID := 98765
		// login is valid
		SetUserIDInContext(ctx, userID)
	} else {
		return nil, status.Error(codes.InvalidArgument, "Invalid username/password combination")
	}

	return &empty.Empty{}, nil
}

// SignOut demonstrates clearing a cookie.  This method sets a flag that
// tell the gateway the cookie should be removed.  This also demonstrates retriving
// the currently signed in user ID from the context.
func (s *server) SignOut(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {

	// get the user ID of the signed in user
	userID, err := GetUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// tell the gateway to delete the session/cookie
	SetDeleteSessionFlagInContext(ctx)

	rlog.Debugf("signout for user %d", userID)

	return &empty.Empty{}, nil
}
