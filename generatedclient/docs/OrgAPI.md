# \OrgAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrg**](OrgAPI.md#GetOrg) | **Get** /orgs/{org_id} | Get organization
[**ListOrgs**](OrgAPI.md#ListOrgs) | **Get** /orgs | List accessible organizations
[**UpdateOrg**](OrgAPI.md#UpdateOrg) | **Patch** /orgs/{org_id} | Update organization



## GetOrg

> GetOrg200Response GetOrg(ctx, orgId).Version(version).Execute()

Get organization



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
	orgId := "b667f176-df52-4b0a-9954-117af6b05ab7" // string | Unique identifier for org

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgAPI.GetOrg(context.Background(), orgId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.GetOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrg`: GetOrg200Response
	fmt.Fprintf(os.Stdout, "Response from `OrgAPI.GetOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique identifier for org | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


### Return type

[**GetOrg200Response**](GetOrg200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgs

> ListOrgs200Response ListOrgs(ctx).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).GroupId(groupId).IsPersonal(isPersonal).Slug(slug).Name(name).Expand(expand).Execute()

List accessible organizations



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
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | If set, only return organizations within the specified group (optional)
	isPersonal := true // bool | If true, only return organizations that are not part of a group. (optional)
	slug := "slug_example" // string | Only return orgs whose slug exactly matches this value. (optional)
	name := "name_example" // string | Only return orgs whose name contains this value. (optional)
	expand := []string{"Expand_example"} // []string | Expand the specified related resources in the response to include their attributes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgAPI.ListOrgs(context.Background()).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).GroupId(groupId).IsPersonal(isPersonal).Slug(slug).Name(name).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.ListOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgs`: ListOrgs200Response
	fmt.Fprintf(os.Stdout, "Response from `OrgAPI.ListOrgs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]
 **groupId** | **string** | If set, only return organizations within the specified group | 
 **isPersonal** | **bool** | If true, only return organizations that are not part of a group. | 
 **slug** | **string** | Only return orgs whose slug exactly matches this value. | 
 **name** | **string** | Only return orgs whose name contains this value. | 
 **expand** | **[]string** | Expand the specified related resources in the response to include their attributes. | 

### Return type

[**ListOrgs200Response**](ListOrgs200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrg

> UpdateOrg200Response UpdateOrg(ctx, orgId).Version(version).UpdateOrgRequest(updateOrgRequest).Execute()

Update organization



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
	orgId := "b667f176-df52-4b0a-9954-117af6b05ab7" // string | Unique identifier for org
	updateOrgRequest := *openapiclient.NewUpdateOrgRequest(*openapiclient.NewUpdateOrgRequestData()) // UpdateOrgRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgAPI.UpdateOrg(context.Background(), orgId).Version(version).UpdateOrgRequest(updateOrgRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.UpdateOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrg`: UpdateOrg200Response
	fmt.Fprintf(os.Stdout, "Response from `OrgAPI.UpdateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique identifier for org | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **updateOrgRequest** | [**UpdateOrgRequest**](UpdateOrgRequest.md) |  | 

### Return type

[**UpdateOrg200Response**](UpdateOrg200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

