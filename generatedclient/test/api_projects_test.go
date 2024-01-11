/*
Snyk API

Testing ProjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ProjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsAPIService DeleteOrgProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var projectId string

		httpRes, err := apiClient.ProjectsAPI.DeleteOrgProject(context.Background(), orgId, projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetOrgProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.GetOrgProject(context.Background(), orgId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService ListOrgProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ProjectsAPI.ListOrgProjects(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService UpdateOrgProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var projectId string

		resp, httpRes, err := apiClient.ProjectsAPI.UpdateOrgProject(context.Background(), orgId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}