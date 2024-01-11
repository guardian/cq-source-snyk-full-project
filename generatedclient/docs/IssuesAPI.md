# \IssuesAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchIssuesPerPurl**](IssuesAPI.md#FetchIssuesPerPurl) | **Get** /orgs/{org_id}/packages/{purl}/issues | List issues for a package
[**ListIssuesForManyPurls**](IssuesAPI.md#ListIssuesForManyPurls) | **Post** /orgs/{org_id}/packages/issues | List issues for a given set of packages  (Currently not available to all customers)



## FetchIssuesPerPurl

> IssuesResponse FetchIssuesPerPurl(ctx, purl, orgId).Version(version).Offset(offset).Limit(limit).Execute()

List issues for a package



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
	purl := "pkg%3Amaven%2Fcom.fasterxml.woodstox%2Fwoodstox-core%405.0.0" // string | A URI-encoded Package URL (purl). Supported purl types are apk, cargo, cocoapods, composer, deb, gem, generic, golang, hex, maven, npm, nuget, pub, pypi, rpm, and swift. A version for the package is also required.
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for an organization
	offset := float32(8.14) // float32 | Specify the number of results to skip before returning results. Must be greater than or equal to 0. Default is 0. (optional)
	limit := float32(8.14) // float32 | Specify the number of results to return. Must be greater than 0 and less than 1000. Default is 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.FetchIssuesPerPurl(context.Background(), purl, orgId).Version(version).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.FetchIssuesPerPurl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchIssuesPerPurl`: IssuesResponse
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.FetchIssuesPerPurl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purl** | **string** | A URI-encoded Package URL (purl). Supported purl types are apk, cargo, cocoapods, composer, deb, gem, generic, golang, hex, maven, npm, nuget, pub, pypi, rpm, and swift. A version for the package is also required. | 
**orgId** | **string** | Unique identifier for an organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIssuesPerPurlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **offset** | **float32** | Specify the number of results to skip before returning results. Must be greater than or equal to 0. Default is 0. | 
 **limit** | **float32** | Specify the number of results to return. Must be greater than 0 and less than 1000. Default is 1000. | 

### Return type

[**IssuesResponse**](IssuesResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIssuesForManyPurls

> IssuesWithPurlsResponse ListIssuesForManyPurls(ctx, orgId).Version(version).BulkPackageUrlsRequestBody(bulkPackageUrlsRequestBody).Execute()

List issues for a given set of packages  (Currently not available to all customers)



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
	bulkPackageUrlsRequestBody := *openapiclient.NewBulkPackageUrlsRequestBody(*openapiclient.NewBulkPackageUrlsRequestBodyData(*openapiclient.NewBulkPackageUrlsRequestBodyDataAttributes([]string{"Purls_example"}))) // BulkPackageUrlsRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.ListIssuesForManyPurls(context.Background(), orgId).Version(version).BulkPackageUrlsRequestBody(bulkPackageUrlsRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.ListIssuesForManyPurls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIssuesForManyPurls`: IssuesWithPurlsResponse
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.ListIssuesForManyPurls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique identifier for an organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIssuesForManyPurlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **bulkPackageUrlsRequestBody** | [**BulkPackageUrlsRequestBody**](BulkPackageUrlsRequestBody.md) |  | 

### Return type

[**IssuesWithPurlsResponse**](IssuesWithPurlsResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

