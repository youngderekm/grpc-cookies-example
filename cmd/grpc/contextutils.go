package main

import (
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SetUserIDInContext sets the user ID into the context.  This has the effect of logging the user
// in as that userId.  The grpc json gateway will set the UID in the user's session in this case
func SetUserIDInContext(ctx context.Context, userID int) {
	// create a header that the gateway will watch for
	header := metadata.Pairs("gateway-session-userId", strconv.Itoa(userID))
	// send the header back to the gateway
	grpc.SendHeader(ctx, header)
}

// SetDeleteSessionFlagInContext sets a flag telling the gateway to delete the session, if any
func SetDeleteSessionFlagInContext(ctx context.Context) {
	// create a header that the gateway will watch for
	header := metadata.Pairs("gateway-session-delete", "true")
	// send the header back to the gateway
	grpc.SendHeader(ctx, header)
}

// get the first metadata value with the given name
func firstMetadataWithName(md metadata.MD, name string) string {
	values := md.Get(name)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

// GetUIDFromContext returns the userId that has been stored in Context, if available.
// This will return 0 if the user has not logged in.  If there is an error attempting to return
// the userId it will be returned.  It's valid for this function to return 0 with no
// error, which indicates the user has not logged in.
func GetUIDFromContext(ctx context.Context) (int, error) {
	// retrieve incoming metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// get the first (and presumably only) user ID from the request metadata
		userIDString := firstMetadataWithName(md, "userId")
		if userIDString != "" {
			userID, err := strconv.Atoi(userIDString)
			if err != nil {
				return 0, errors.Wrap(err, "unable to parse userId")
			}
			return userID, nil
		}
	}
	return 0, nil
}
