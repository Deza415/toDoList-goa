// Code generated by goa v3.20.1, DO NOT EDIT.
//
// todo HTTP client CLI support package
//
// Command:
// $ goa gen github.com/Deza415/toDoList-goa/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	todoc "github.com/Deza415/toDoList-goa/gen/http/todo/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `todo (list|create)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` todo list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		todoFlags = flag.NewFlagSet("todo", flag.ContinueOnError)

		todoListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		todoCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		todoCreateBodyFlag = todoCreateFlags.String("body", "REQUIRED", "")
	)
	todoFlags.Usage = todoUsage
	todoListFlags.Usage = todoListUsage
	todoCreateFlags.Usage = todoCreateUsage

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
		case "todo":
			svcf = todoFlags
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
		case "todo":
			switch epn {
			case "list":
				epf = todoListFlags

			case "create":
				epf = todoCreateFlags

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
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "todo":
			c := todoc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
			case "create":
				endpoint = c.Create()
				data, err = todoc.BuildCreatePayload(*todoCreateBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// todoUsage displays the usage of the todo command and its subcommands.
func todoUsage() {
	fmt.Fprintf(os.Stderr, `The todo service manages todo items
Usage:
    %[1]s [globalflags] todo COMMAND [flags]

COMMAND:
    list: List all todo items
    create: Create a new todo item

Additional help:
    %[1]s todo COMMAND --help
`, os.Args[0])
}
func todoListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo list

List all todo items

Example:
    %[1]s todo list
`, os.Args[0])
}

func todoCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todo create -body JSON

Create a new todo item
    -body JSON: 

Example:
    %[1]s todo create --body '{
      "title": "0"
   }'
`, os.Args[0])
}
