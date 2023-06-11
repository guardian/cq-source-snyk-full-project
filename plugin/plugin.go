package plugin

import (
	"github.com/guardian/cq-source-snyk-full-project/client"
	"github.com/guardian/cq-source-snyk-full-project/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"guardian-snyk-full-project",
		Version,
		schema.Tables{
			resources.SampleTable(),
		},
		client.New,
	)
}
