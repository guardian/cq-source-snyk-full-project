# \SlackAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChannelNameById**](SlackAPI.md#GetChannelNameById) | **Get** /orgs/{org_id}/slack_app/{tenant_id}/channels/{channel_id} | Get Slack Channel name by Slack Channel ID.
[**ListChannels**](SlackAPI.md#ListChannels) | **Get** /orgs/{org_id}/slack_app/{tenant_id}/channels | Get a list of Slack channels



## GetChannelNameById

> GetChannelNameById200Response GetChannelNameById(ctx, orgId, channelId, tenantId).Version(version).Execute()

Get Slack Channel name by Slack Channel ID.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	channelId := "channelId_example" // string | Slack Channel ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackAPI.GetChannelNameById(context.Background(), orgId, channelId, tenantId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.GetChannelNameById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelNameById`: GetChannelNameById200Response
	fmt.Fprintf(os.Stdout, "Response from `SlackAPI.GetChannelNameById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**channelId** | **string** | Slack Channel ID | 
**tenantId** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelNameByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 




### Return type

[**GetChannelNameById200Response**](GetChannelNameById200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannels

> ListChannels200Response ListChannels(ctx, orgId, tenantId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of Slack channels



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Tenant ID
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(100) // int32 | Number of results to return per page (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackAPI.ListChannels(context.Background(), orgId, tenantId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackAPI.ListChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChannels`: ListChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `SlackAPI.ListChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**tenantId** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 1000]

### Return type

[**ListChannels200Response**](ListChannels200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

