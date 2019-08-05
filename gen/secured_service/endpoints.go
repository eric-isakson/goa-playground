// Code generated by goa v3.0.3, DO NOT EDIT.
//
// secured_service endpoints
//
// Command:
// $ goa gen github.com/eric-isakson/goa-playground/design

package securedservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "secured_service" service endpoints.
type Endpoints struct {
	Signin goa.Endpoint
	Secure goa.Endpoint
}

// NewEndpoints wraps the methods of the "secured_service" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Signin: NewSigninEndpoint(s, a.BasicAuth),
		Secure: NewSecureEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "secured_service" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Signin = m(e.Signin)
	e.Secure = m(e.Secure)
}

// NewSigninEndpoint returns an endpoint function that calls the method
// "signin" of service "secured_service".
func NewSigninEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SigninPayload)
		var err error
		sc := security.BasicScheme{
			Name:           "basic",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
		}
		ctx, err = authBasicFn(ctx, p.Username, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return s.Signin(ctx, p)
	}
}

// NewSecureEndpoint returns an endpoint function that calls the method
// "secure" of service "secured_service".
func NewSecureEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SecurePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:read"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Secure(ctx, p)
	}
}