// Code generated by goa v3.0.3, DO NOT EDIT.
//
// secured_service gRPC client CLI support package
//
// Command:
// $ goa gen github.com/eric-isakson/goa-playground/design

package client

import (
	"encoding/json"
	"fmt"

	secured_servicepb "github.com/eric-isakson/goa-playground/gen/grpc/secured_service/pb"
	securedservice "github.com/eric-isakson/goa-playground/gen/secured_service"
)

// BuildSigninPayload builds the payload for the secured_service signin
// endpoint from CLI flags.
func BuildSigninPayload(securedServiceSigninUsername string, securedServiceSigninPassword string) (*securedservice.SigninPayload, error) {
	var username string
	{
		username = securedServiceSigninUsername
	}
	var password string
	{
		password = securedServiceSigninPassword
	}
	v := &securedservice.SigninPayload{}
	v.Username = username
	v.Password = password
	return v, nil
}

// BuildSecurePayload builds the payload for the secured_service secure
// endpoint from CLI flags.
func BuildSecurePayload(securedServiceSecureMessage string, securedServiceSecureToken string) (*securedservice.SecurePayload, error) {
	var err error
	var message secured_servicepb.SecureRequest
	{
		if securedServiceSecureMessage != "" {
			err = json.Unmarshal([]byte(securedServiceSecureMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"test\": \"values\"\n   }'")
			}
		}
	}
	var token string
	{
		token = securedServiceSecureToken
	}
	v := &securedservice.SecurePayload{}
	if message.Test != "" {
		v.Test = &message.Test
	}
	v.Token = token
	return v, nil
}
