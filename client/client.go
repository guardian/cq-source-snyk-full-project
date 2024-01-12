package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	openapi "github.com/guardian/cq-source-snyk-full-project/generatedclient"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger        zerolog.Logger
	Organisations []string
	SnykClient    *openapi.APIClient
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "ID"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	snykSpec := Spec{}
	err := s.UnmarshalSpec(&snykSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	config := openapi.NewConfiguration()
	config.AddDefaultHeader("Authorization", "token "+snykSpec.APIKey)
	client := openapi.NewAPIClient(config)

	if err != nil {
		return nil, fmt.Errorf("failed to create Snyk client: %w", err)
	}

	// init orgs
	orgs, _, err := client.OrgAPI.ListOrgs(ctx).Version("2024-01-04").Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to list Snyk orgs: %w", err)
	}

	var orgIDs []string
	for _, org := range orgs.Data {
		orgIDs = append(orgIDs, org.Id)
	}

	return &Client{
		Logger:        logger,
		SnykClient:    client,
		Organisations: orgIDs,
	}, nil
}
