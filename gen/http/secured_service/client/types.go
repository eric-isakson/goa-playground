// Code generated by goa v3.0.3, DO NOT EDIT.
//
// secured_service HTTP client types
//
// Command:
// $ goa gen github.com/eric-isakson/goa-playground/design

package client

import (
	securedservice "github.com/eric-isakson/goa-playground/gen/secured_service"
	goa "goa.design/goa/v3/pkg"
)

// SigninResponseBody is the type of the "secured_service" service "signin"
// endpoint HTTP response body.
type SigninResponseBody struct {
	// JWT token
	JWT *string `form:"jwt,omitempty" json:"jwt,omitempty" xml:"jwt,omitempty"`
}

// SigninUnauthorizedResponseBody is the type of the "secured_service" service
// "signin" endpoint HTTP response body for the "unauthorized" error.
type SigninUnauthorizedResponseBody string

// SecureInvalidScopesResponseBody is the type of the "secured_service" service
// "secure" endpoint HTTP response body for the "invalid-scopes" error.
type SecureInvalidScopesResponseBody string

// SecureUnauthorizedResponseBody is the type of the "secured_service" service
// "secure" endpoint HTTP response body for the "unauthorized" error.
type SecureUnauthorizedResponseBody string

// NewSigninCredsOK builds a "secured_service" service "signin" endpoint result
// from a HTTP "OK" response.
func NewSigninCredsOK(body *SigninResponseBody) *securedservice.Creds {
	v := &securedservice.Creds{
		JWT: *body.JWT,
	}
	return v
}

// NewSigninUnauthorized builds a secured_service service signin endpoint
// unauthorized error.
func NewSigninUnauthorized(body SigninUnauthorizedResponseBody) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)
	return v
}

// NewSecureInvalidScopes builds a secured_service service secure endpoint
// invalid-scopes error.
func NewSecureInvalidScopes(body SecureInvalidScopesResponseBody) securedservice.InvalidScopes {
	v := securedservice.InvalidScopes(body)
	return v
}

// NewSecureUnauthorized builds a secured_service service secure endpoint
// unauthorized error.
func NewSecureUnauthorized(body SecureUnauthorizedResponseBody) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)
	return v
}

// ValidateSigninResponseBody runs the validations defined on SigninResponseBody
func ValidateSigninResponseBody(body *SigninResponseBody) (err error) {
	if body.JWT == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("jwt", "body"))
	}
	return
}
