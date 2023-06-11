package main

import (
	"github.com/guardian/cq-source-snyk-full-project/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
