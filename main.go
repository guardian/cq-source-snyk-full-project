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
			"Authorization": {Key: "token <SNYK_API_KEY>"},
		},
	)

	client := openapi.NewAPIClient(openapi.NewConfiguration())

	resp, http_resp, err := client.ProjectsAPI.ListOrgProjects(auth, "foobar").Version("2024-01-04").Execute()

	fmt.Println(resp)
	fmt.Println(http_resp)
	fmt.Println(err)

	// serve.Source(plugin.Plugin())
}
