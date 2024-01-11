# \SBOMAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSbom**](SBOMAPI.md#GetSbom) | **Get** /orgs/{org_id}/projects/{project_id}/sbom | Get a project’s SBOM document



## GetSbom

> map[string]interface{} GetSbom(ctx, orgId, projectId).Version(version).Format(format).Execute()

Get a project’s SBOM document



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for an organization
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for a project
	format := "cyclonedx1.4+json" // string | The desired SBOM format of the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMAPI.GetSbom(context.Background(), orgId, projectId).Version(version).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.GetSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbom`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SBOMAPI.GetSbom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique identifier for an organization | 
**projectId** | **string** | Unique identifier for a project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **format** | **string** | The desired SBOM format of the response. | 

### Return type

**map[string]interface{}**

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.api+json, application/vnd.cyclonedx+json, application/vnd.cyclonedx+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

