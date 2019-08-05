// Code generated by goa v3.0.3, DO NOT EDIT.
//
// secured_service gRPC client types
//
// Command:
// $ goa gen github.com/eric-isakson/goa-playground/design

package client

import (
	secured_servicepb "github.com/eric-isakson/goa-playground/gen/grpc/secured_service/pb"
	securedservice "github.com/eric-isakson/goa-playground/gen/secured_service"
)

// NewSigninRequest builds the gRPC request type from the payload of the
// "signin" endpoint of the "secured_service" service.
func NewSigninRequest() *secured_servicepb.SigninRequest {
	message := &secured_servicepb.SigninRequest{}
	return message
}

// NewSigninResult builds the result type of the "signin" endpoint of the
// "secured_service" service from the gRPC response type.
func NewSigninResult(message *secured_servicepb.SigninResponse) *securedservice.Creds {
	result := &securedservice.Creds{
		JWT: message.Jwt,
	}
	return result
}

// NewSecureRequest builds the gRPC request type from the payload of the
// "secure" endpoint of the "secured_service" service.
func NewSecureRequest(payload *securedservice.SecurePayload) *secured_servicepb.SecureRequest {
	message := &secured_servicepb.SecureRequest{}
	if payload.Fail != nil {
		message.Fail = *payload.Fail
	}
	return message
}

// NewSecureResult builds the result type of the "secure" endpoint of the
// "secured_service" service from the gRPC response type.
func NewSecureResult(message *secured_servicepb.SecureResponse) string {
	result := message.Field
	return result
}