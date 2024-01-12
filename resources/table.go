package resources

import (
	"context"
	"math"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/guardian/cq-source-snyk-full-project/client"
	openapi "github.com/guardian/cq-source-snyk-full-project/generatedclient"
)

type Tag = openapi.PatchProjectRequestDataAttributesTagsInner

type Project struct {
	ID                    string         `json:"id,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Origin                string         `json:"origin,omitempty"`
	IssueCountsBySeverity map[string]int `json:"issueCountsBySeverity,omitempty"`
	Tags                  []Tag          `json:"tags,omitempty"`
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
		snykProjects, _, err := snykClient.ProjectsAPI.ListOrgProjects(ctx, orgID).Execute()
		if err != nil {
			return err
		}

		for _, snykProject := range snykProjects.Data {
			project := Project{
				ID:                    snykProject.Id,
				Name:                  snykProject.Attributes.Name,
				Origin:                snykProject.Attributes.Origin,
				IssueCountsBySeverity: map[string]int{"critical": int(math.Round(float64(*snykProject.Meta.LatestIssueCounts.Critical))), "bar": 2},
				Tags:                  snykProject.Attributes.Tags,
				OrgID:                 orgID,
			}

			res <- project
		}
	}

	return nil
}
