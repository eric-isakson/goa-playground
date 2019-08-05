// Code generated by goa v3.0.3, DO NOT EDIT.
//
// playground gRPC client CLI support package
//
// Command:
// $ goa gen github.com/eric-isakson/goa-playground/design

package cli

import (
	"flag"
	"fmt"
	"os"

	securedservicec "github.com/eric-isakson/goa-playground/gen/grpc/secured_service/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `secured-service (signin|secure)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` secured-service signin --username "user" --password "password"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		securedServiceFlags = flag.NewFlagSet("secured-service", flag.ContinueOnError)

		securedServiceSigninFlags        = flag.NewFlagSet("signin", flag.ExitOnError)
		securedServiceSigninUsernameFlag = securedServiceSigninFlags.String("username", "REQUIRED", "")
		securedServiceSigninPasswordFlag = securedServiceSigninFlags.String("password", "REQUIRED", "")

		securedServiceSecureFlags       = flag.NewFlagSet("secure", flag.ExitOnError)
		securedServiceSecureMessageFlag = securedServiceSecureFlags.String("message", "", "")
		securedServiceSecureTokenFlag   = securedServiceSecureFlags.String("token", "REQUIRED", "")
	)
	securedServiceFlags.Usage = securedServiceUsage
	securedServiceSigninFlags.Usage = securedServiceSigninUsage
	securedServiceSecureFlags.Usage = securedServiceSecureUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "secured-service":
			svcf = securedServiceFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "secured-service":
			switch epn {
			case "signin":
				epf = securedServiceSigninFlags

			case "secure":
				epf = securedServiceSecureFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "secured-service":
			c := securedservicec.NewClient(cc, opts...)
			switch epn {
			case "signin":
				endpoint = c.Signin()
				data, err = securedservicec.BuildSigninPayload(*securedServiceSigninUsernameFlag, *securedServiceSigninPasswordFlag)
			case "secure":
				endpoint = c.Secure()
				data, err = securedservicec.BuildSecurePayload(*securedServiceSecureMessageFlag, *securedServiceSecureTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// secured-serviceUsage displays the usage of the secured-service command and
// its subcommands.
func securedServiceUsage() {
	fmt.Fprintf(os.Stderr, `The secured service exposes endpoints that require valid authorization credentials.
Usage:
    %s [globalflags] secured-service COMMAND [flags]

COMMAND:
    signin: Creates a valid JWT
    secure: This action is secured with the jwt scheme

Additional help:
    %s secured-service COMMAND --help
`, os.Args[0], os.Args[0])
}
func securedServiceSigninUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] secured-service signin -username STRING -password STRING

Creates a valid JWT
    -username STRING: 
    -password STRING: 

Example:
    `+os.Args[0]+` secured-service signin --username "user" --password "password"
`, os.Args[0])
}

func securedServiceSecureUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] secured-service secure -message JSON -token STRING

This action is secured with the jwt scheme
    -message JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` secured-service secure --message '{
      "test": "values"
   }' --token "Nemo id occaecati molestias."
`, os.Args[0])
}
