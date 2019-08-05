package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("playground", func() {
	Title("Playground API")
	Description("This API is used for experimenting with Goa features.")
	Docs(func() { // Documentation links
		Description("Eric Isakson's Goa Playground README")
		URL("https://github.com/eric-isakson/goa-playground/tree/master/README.md")
	})
	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("playground", func() {
		Description("playground hosts the Playground Service.")

		// List the services hosted by this server.
		Services("secured_service")

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://localhost:8000/playground")
			URI("grpc://localhost:8080")
		})
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
	Scope("api:read", "Read-only access")
})

var _ = Service("secured_service", func() {
	Description("The secured service exposes endpoints that require valid authorization credentials.")

	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	GRPC(func() {
		Response("unauthorized", CodeUnauthenticated)
	})

	Method("signin", func() {
		Description("Creates a valid JWT")

		// The signin endpoint is secured via basic auth
		Security(BasicAuth)

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			UsernameField(1, "username", String, "Username used to perform signin", func() {
				Example("user")
			})
			PasswordField(2, "password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Required("username", "password")
		})

		Result(Creds)

		HTTP(func() {
			POST("/signin")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("secure", func() {
		Description("This action is secured with the jwt scheme")

		Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			Scope("api:read") // Enforce presence of "api:read" scope in JWT claims.
		})

		Payload(func() {
			Field(1, "test", String, func() {
				Description("Are these exposed for unauthenticated requests?")
				Enum("exposed", "enum", "values")
			})
			TokenField(2, "token", String, func() {
				Description("JWT used for authentication")
			})
			Required("token")
		})

		Result(String)

		Error("invalid-scopes", String, "Token scopes are invalid")

		HTTP(func() {
			GET("/secure")
			Param("test")
			Response(StatusOK)
			Response("invalid-scopes", StatusForbidden)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("invalid-scopes", CodeUnauthenticated)
		})
	})
})

// Creds defines the credentials to use for authenticating to service methods.
var Creds = Type("Creds", func() {
	Field(1, "jwt", String, "JWT token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Required("jwt")
})
