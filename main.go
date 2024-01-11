package main

import (
	"context"
	"fmt"

	openapi "github.com/guardian/cq-source-snyk-full-project/generatedclient"
)

// import (
// 	"github.com/guardian/cq-source-snyk-full-project/plugin"

// 	"github.com/cloudquery/plugin-sdk/v3/serve"
// )

func main() {
	fmt.Println("Hello world!")

	auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)

	client := openapi.NewAPIClient(openapi.NewConfiguration())

	x := client.ProjectsAPI.ListOrgProjects(auth, "foobar")
	fmt.Println(x)

	// serve.Source(plugin.Plugin())
}
