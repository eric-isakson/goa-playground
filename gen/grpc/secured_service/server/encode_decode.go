// Code generated by goa v3.0.3, DO NOT EDIT.
//
// secured_service gRPC server encoders and decoders
//
// Command:
// $ goa gen github.com/eric-isakson/goa-playground/design

package server

import (
	"context"
	"strings"

	secured_servicepb "github.com/eric-isakson/goa-playground/gen/grpc/secured_service/pb"
	securedservice "github.com/eric-isakson/goa-playground/gen/secured_service"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/metadata"
)

// EncodeSigninResponse encodes responses from the "secured_service" service
// "signin" endpoint.
func EncodeSigninResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*securedservice.Creds)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "signin", "*securedservice.Creds", v)
	}
	resp := NewSigninResponse(result)
	return resp, nil
}

// DecodeSigninRequest decodes requests sent to "secured_service" service
// "signin" endpoint.
func DecodeSigninRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		username string
		password string
		err      error
	)
	{
		if vals := md.Get("username"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("username", "metadata"))
		} else {
			username = vals[0]
		}
		if vals := md.Get("password"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("password", "metadata"))
		} else {
			password = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *securedservice.SigninPayload
	{
		payload = NewSigninPayload(username, password)
	}
	return payload, nil
}

// EncodeSecureResponse encodes responses from the "secured_service" service
// "secure" endpoint.
func EncodeSecureResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "secure", "string", v)
	}
	resp := NewSecureResponse(result)
	return resp, nil
}

// DecodeSecureRequest decodes requests sent to "secured_service" service
// "secure" endpoint.
func DecodeSecureRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *secured_servicepb.SecureRequest
		ok      bool
	)
	{
		if message, ok = v.(*secured_servicepb.SecureRequest); !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "secure", "*secured_servicepb.SecureRequest", v)
		}
	}
	var payload *securedservice.SecurePayload
	{
		payload = NewSecurePayload(message, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}
