# \SastSettingsAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSastSettings**](SastSettingsAPI.md#GetSastSettings) | **Get** /orgs/{org_id}/settings/sast | Retrieves the SAST settings for an org
[**UpdateOrgSastSettings**](SastSettingsAPI.md#UpdateOrgSastSettings) | **Patch** /orgs/{org_id}/settings/sast | Enable/Disable the Snyk Code settings for an org



## GetSastSettings

> GetSastSettings200Response GetSastSettings(ctx, orgId).Version(version).Execute()

Retrieves the SAST settings for an org



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the org for which we want to retrieve the SAST settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SastSettingsAPI.GetSastSettings(context.Background(), orgId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SastSettingsAPI.GetSastSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSastSettings`: GetSastSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SastSettingsAPI.GetSastSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org for which we want to retrieve the SAST settings | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSastSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


### Return type

[**GetSastSettings200Response**](GetSastSettings200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSastSettings

> GetSastSettings200Response UpdateOrgSastSettings(ctx, orgId).Version(version).UpdateOrgSastSettingsRequest(updateOrgSastSettingsRequest).Execute()

Enable/Disable the Snyk Code settings for an org



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the org for which we want to update the Snyk Code setting
	updateOrgSastSettingsRequest := *openapiclient.NewUpdateOrgSastSettingsRequest(*openapiclient.NewUpdateOrgSastSettingsRequestData(*openapiclient.NewUpdateOrgSastSettingsRequestDataAttributes(false), "Id_example", "Type_example")) // UpdateOrgSastSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SastSettingsAPI.UpdateOrgSastSettings(context.Background(), orgId).Version(version).UpdateOrgSastSettingsRequest(updateOrgSastSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SastSettingsAPI.UpdateOrgSastSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSastSettings`: GetSastSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SastSettingsAPI.UpdateOrgSastSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org for which we want to update the Snyk Code setting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSastSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **updateOrgSastSettingsRequest** | [**UpdateOrgSastSettingsRequest**](UpdateOrgSastSettingsRequest.md) |  | 

### Return type

[**GetSastSettings200Response**](GetSastSettings200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

