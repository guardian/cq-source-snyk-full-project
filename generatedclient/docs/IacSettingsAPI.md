# \IacSettingsAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIacSettingsForGroup**](IacSettingsAPI.md#GetIacSettingsForGroup) | **Get** /groups/{group_id}/settings/iac | Get the Infrastructure as Code Settings for a group
[**GetIacSettingsForOrg**](IacSettingsAPI.md#GetIacSettingsForOrg) | **Get** /orgs/{org_id}/settings/iac | Get the Infrastructure as Code Settings for an org.
[**UpdateIacSettingsForGroup**](IacSettingsAPI.md#UpdateIacSettingsForGroup) | **Patch** /groups/{group_id}/settings/iac | Update the Infrastructure as Code Settings for a group
[**UpdateIacSettingsForOrg**](IacSettingsAPI.md#UpdateIacSettingsForOrg) | **Patch** /orgs/{org_id}/settings/iac | Update the Infrastructure as Code Settings for an org



## GetIacSettingsForGroup

> GetIacSettingsForGroup200Response GetIacSettingsForGroup(ctx, groupId).Version(version).Execute()

Get the Infrastructure as Code Settings for a group



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the group whose Infrastructure as Code settings are requested

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IacSettingsAPI.GetIacSettingsForGroup(context.Background(), groupId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IacSettingsAPI.GetIacSettingsForGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIacSettingsForGroup`: GetIacSettingsForGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `IacSettingsAPI.GetIacSettingsForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The id of the group whose Infrastructure as Code settings are requested | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIacSettingsForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


### Return type

[**GetIacSettingsForGroup200Response**](GetIacSettingsForGroup200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIacSettingsForOrg

> GetIacSettingsForOrg200Response GetIacSettingsForOrg(ctx, orgId).Version(version).Execute()

Get the Infrastructure as Code Settings for an org.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the org whose Infrastructure as Code settings are requested.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IacSettingsAPI.GetIacSettingsForOrg(context.Background(), orgId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IacSettingsAPI.GetIacSettingsForOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIacSettingsForOrg`: GetIacSettingsForOrg200Response
	fmt.Fprintf(os.Stdout, "Response from `IacSettingsAPI.GetIacSettingsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org whose Infrastructure as Code settings are requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIacSettingsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


### Return type

[**GetIacSettingsForOrg200Response**](GetIacSettingsForOrg200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIacSettingsForGroup

> GetIacSettingsForGroup200Response UpdateIacSettingsForGroup(ctx, groupId).Version(version).UpdateIacSettingsForGroupRequest(updateIacSettingsForGroupRequest).Execute()

Update the Infrastructure as Code Settings for a group



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the group whose Infrastructure as Code settings are getting updated
	updateIacSettingsForGroupRequest := *openapiclient.NewUpdateIacSettingsForGroupRequest() // UpdateIacSettingsForGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IacSettingsAPI.UpdateIacSettingsForGroup(context.Background(), groupId).Version(version).UpdateIacSettingsForGroupRequest(updateIacSettingsForGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IacSettingsAPI.UpdateIacSettingsForGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIacSettingsForGroup`: GetIacSettingsForGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `IacSettingsAPI.UpdateIacSettingsForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The id of the group whose Infrastructure as Code settings are getting updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIacSettingsForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **updateIacSettingsForGroupRequest** | [**UpdateIacSettingsForGroupRequest**](UpdateIacSettingsForGroupRequest.md) |  | 

### Return type

[**GetIacSettingsForGroup200Response**](GetIacSettingsForGroup200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIacSettingsForOrg

> GetIacSettingsForOrg200Response UpdateIacSettingsForOrg(ctx, orgId).Version(version).UpdateIacSettingsForOrgRequest(updateIacSettingsForOrgRequest).Execute()

Update the Infrastructure as Code Settings for an org



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the org whose Infrastructure as Code settings are getting updated
	updateIacSettingsForOrgRequest := *openapiclient.NewUpdateIacSettingsForOrgRequest() // UpdateIacSettingsForOrgRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IacSettingsAPI.UpdateIacSettingsForOrg(context.Background(), orgId).Version(version).UpdateIacSettingsForOrgRequest(updateIacSettingsForOrgRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IacSettingsAPI.UpdateIacSettingsForOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIacSettingsForOrg`: GetIacSettingsForOrg200Response
	fmt.Fprintf(os.Stdout, "Response from `IacSettingsAPI.UpdateIacSettingsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org whose Infrastructure as Code settings are getting updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIacSettingsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **updateIacSettingsForOrgRequest** | [**UpdateIacSettingsForOrgRequest**](UpdateIacSettingsForOrgRequest.md) |  | 

### Return type

[**GetIacSettingsForOrg200Response**](GetIacSettingsForOrg200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

