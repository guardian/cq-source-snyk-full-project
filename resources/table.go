package resources

import (
	"context"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/guardian/cq-source-snyk-full-project/client"
)

type Project struct {
	ID                    string         `json:"id,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Origin                string         `json:"origin,omitempty"`
	IssueCountsBySeverity map[string]int `json:"issueCountsBySeverity,omitempty"`
	Tags                  []snyk.Tag     `json:"tags,omitempty"`
	OrgID                 string         `json:"orgId"`
}

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "snyk_projects",
		Description: `https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects`,
		Resolver:    fetchProjects,
		Transform:   transformers.TransformWithStruct(Project{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	snykClient := c.SnykClient

	for _, orgID := range c.Organisations {
		snykProjects, _, err := snykClient.Projects.List(ctx, orgID)
		if err != nil {
			return err
		}

		for _, snykProject := range snykProjects {
			project := Project{
				ID:                    snykProject.ID,
				Name:                  snykProject.Name,
				Origin:                snykProject.Origin,
				IssueCountsBySeverity: snykProject.IssueCountsBySeverity,
				Tags:                  snykProject.Tags,
				OrgID:                 orgID,
			}

			res <- project
		}
	}

	return nil
}
