package main

import (
	"os"

	"github.com/danielgtaylor/restish/cli"
	"github.com/danielgtaylor/restish/oauth"
	"github.com/danielgtaylor/restish/openapi"
)

var version string = "dev"
var commit string
var date string

func main() {
	if version == "dev" {
		// Try to add the executable modification time to the dev version.
		if info, err := os.Stat(os.Args[0]); err == nil {
			version += "-" + info.ModTime().Format("2006-01-02-15:04")
		}
	}

	cli.Init("restish", version)

	// Register default encodings, content type handlers, and link parsers.
	cli.Defaults()

	// Register format loaders to auto-discover API descriptions
	cli.AddLoader(openapi.New())

	// Register auth schemes
	cli.AddAuth("oauth-client-credentials", &oauth.ClientCredentialsHandler{})
	cli.AddAuth("oauth-authorization-code", &oauth.AuthorizationCodeHandler{})

	// Run the CLI, parsing arguments, making requests, and printing responses.
	cli.Run()
}
