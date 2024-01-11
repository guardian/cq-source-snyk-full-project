# \ProjectsAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgProject**](ProjectsAPI.md#DeleteOrgProject) | **Delete** /orgs/{org_id}/projects/{project_id} | Delete project by project ID.
[**GetOrgProject**](ProjectsAPI.md#GetOrgProject) | **Get** /orgs/{org_id}/projects/{project_id} | Get project by project ID.
[**ListOrgProjects**](ProjectsAPI.md#ListOrgProjects) | **Get** /orgs/{org_id}/projects | List all Projects for an Org with the given Org ID.
[**UpdateOrgProject**](ProjectsAPI.md#UpdateOrgProject) | **Patch** /orgs/{org_id}/projects/{project_id} | Updates project by project ID.



## DeleteOrgProject

> DeleteOrgProject(ctx, orgId, projectId).Version(version).Execute()

Delete project by project ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the org to which the project belongs to.
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the project.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteOrgProject(context.Background(), orgId, projectId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteOrgProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the org to which the project belongs to. | 
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgProject

> GetOrgProject200Response GetOrgProject(ctx, orgId, projectId).Version(version).Expand(expand).MetaLatestIssueCounts(metaLatestIssueCounts).MetaLatestDependencyTotal(metaLatestDependencyTotal).Execute()

Get project by project ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the org to which the project belongs to.
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the project.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	expand := []string{"Expand_example"} // []string | Expand relationships. (optional)
	metaLatestIssueCounts := true // bool | Include a summary count for the issues found in the most recent scan of this project (optional)
	metaLatestDependencyTotal := true // bool | Include the total number of dependencies found in the most recent scan of this project (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetOrgProject(context.Background(), orgId, projectId).Version(version).Expand(expand).MetaLatestIssueCounts(metaLatestIssueCounts).MetaLatestDependencyTotal(metaLatestDependencyTotal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetOrgProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgProject`: GetOrgProject200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetOrgProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the org to which the project belongs to. | 
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 
 **expand** | **[]string** | Expand relationships. | 
 **metaLatestIssueCounts** | **bool** | Include a summary count for the issues found in the most recent scan of this project | 
 **metaLatestDependencyTotal** | **bool** | Include the total number of dependencies found in the most recent scan of this project | 

### Return type

[**GetOrgProject200Response**](GetOrgProject200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgProjects

> ListOrgProjects200Response ListOrgProjects(ctx, orgId).Version(version).TargetId(targetId).TargetReference(targetReference).TargetFile(targetFile).TargetRuntime(targetRuntime).MetaCount(metaCount).Ids(ids).Names(names).Origins(origins).Types(types).Expand(expand).MetaLatestIssueCounts(metaLatestIssueCounts).MetaLatestDependencyTotal(metaLatestDependencyTotal).CliMonitoredBefore(cliMonitoredBefore).CliMonitoredAfter(cliMonitoredAfter).ImportingUserPublicId(importingUserPublicId).Tags(tags).BusinessCriticality(businessCriticality).Environment(environment).Lifecycle(lifecycle).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

List all Projects for an Org with the given Org ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the org that the projects belong to.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	targetId := []string{"Inner_example"} // []string | Return projects that belong to the provided targets (optional)
	targetReference := "targetReference_example" // string | Return projects that match the provided target reference (optional)
	targetFile := "targetFile_example" // string | Return projects that match the provided target file (optional)
	targetRuntime := "targetRuntime_example" // string | Return projects that match the provided target runtime (optional)
	metaCount := "metaCount_example" // string | The collection count. (optional)
	ids := []string{"Inner_example"} // []string | Return projects that match the provided IDs. (optional)
	names := []string{"Inner_example"} // []string | Return projects that match the provided names. (optional)
	origins := []string{"Inner_example"} // []string | Return projects that match the provided origins. (optional)
	types := []string{"Inner_example"} // []string | Return projects that match the provided types. (optional)
	expand := []string{"Expand_example"} // []string | Expand relationships. (optional)
	metaLatestIssueCounts := true // bool | Include a summary count for the issues found in the most recent scan of this project (optional)
	metaLatestDependencyTotal := true // bool | Include the total number of dependencies found in the most recent scan of this project (optional)
	cliMonitoredBefore := time.Now() // time.Time | Filter projects uploaded and monitored before this date (encoded value) (optional)
	cliMonitoredAfter := time.Now() // time.Time | Filter projects uploaded and monitored after this date (encoded value) (optional)
	importingUserPublicId := []string{"Inner_example"} // []string | Return projects that match the provided importing user public ids. (optional)
	tags := []string{"Inner_example"} // []string | Return projects that match all the provided tags (optional)
	businessCriticality := []string{"BusinessCriticality_example"} // []string | Return projects that match all the provided business_criticality value (optional)
	environment := []string{"Environment_example"} // []string | Return projects that match all the provided environment values (optional)
	lifecycle := []string{"Lifecycle_example"} // []string | Return projects that match all the provided lifecycle values (optional)
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ListOrgProjects(context.Background(), orgId).Version(version).TargetId(targetId).TargetReference(targetReference).TargetFile(targetFile).TargetRuntime(targetRuntime).MetaCount(metaCount).Ids(ids).Names(names).Origins(origins).Types(types).Expand(expand).MetaLatestIssueCounts(metaLatestIssueCounts).MetaLatestDependencyTotal(metaLatestDependencyTotal).CliMonitoredBefore(cliMonitoredBefore).CliMonitoredAfter(cliMonitoredAfter).ImportingUserPublicId(importingUserPublicId).Tags(tags).BusinessCriticality(businessCriticality).Environment(environment).Lifecycle(lifecycle).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ListOrgProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgProjects`: ListOrgProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ListOrgProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the org that the projects belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **targetId** | **[]string** | Return projects that belong to the provided targets | 
 **targetReference** | **string** | Return projects that match the provided target reference | 
 **targetFile** | **string** | Return projects that match the provided target file | 
 **targetRuntime** | **string** | Return projects that match the provided target runtime | 
 **metaCount** | **string** | The collection count. | 
 **ids** | **[]string** | Return projects that match the provided IDs. | 
 **names** | **[]string** | Return projects that match the provided names. | 
 **origins** | **[]string** | Return projects that match the provided origins. | 
 **types** | **[]string** | Return projects that match the provided types. | 
 **expand** | **[]string** | Expand relationships. | 
 **metaLatestIssueCounts** | **bool** | Include a summary count for the issues found in the most recent scan of this project | 
 **metaLatestDependencyTotal** | **bool** | Include the total number of dependencies found in the most recent scan of this project | 
 **cliMonitoredBefore** | **time.Time** | Filter projects uploaded and monitored before this date (encoded value) | 
 **cliMonitoredAfter** | **time.Time** | Filter projects uploaded and monitored after this date (encoded value) | 
 **importingUserPublicId** | **[]string** | Return projects that match the provided importing user public ids. | 
 **tags** | **[]string** | Return projects that match all the provided tags | 
 **businessCriticality** | **[]string** | Return projects that match all the provided business_criticality value | 
 **environment** | **[]string** | Return projects that match all the provided environment values | 
 **lifecycle** | **[]string** | Return projects that match all the provided lifecycle values | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**ListOrgProjects200Response**](ListOrgProjects200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgProject

> UpdateOrgProject200Response UpdateOrgProject(ctx, orgId, projectId).Version(version).Expand(expand).PatchProjectRequest(patchProjectRequest).Execute()

Updates project by project ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Org the project belongs to.
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the project to patch.
	expand := []string{"Expand_example"} // []string | Expand relationships. (optional)
	patchProjectRequest := *openapiclient.NewPatchProjectRequest(*openapiclient.NewPatchProjectRequestData(*openapiclient.NewPatchProjectRequestDataAttributes(), "331ede0a-de94-456f-b788-166caeca58bf", *openapiclient.NewPatchProjectRequestDataRelationships(), "Type_example")) // PatchProjectRequest | The project attributes to be updated. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.UpdateOrgProject(context.Background(), orgId, projectId).Version(version).Expand(expand).PatchProjectRequest(patchProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateOrgProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgProject`: UpdateOrgProject200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateOrgProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the Org the project belongs to. | 
**projectId** | **string** | The ID of the project to patch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **expand** | **[]string** | Expand relationships. | 
 **patchProjectRequest** | [**PatchProjectRequest**](PatchProjectRequest.md) | The project attributes to be updated. | 

### Return type

[**UpdateOrgProject200Response**](UpdateOrgProject200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

