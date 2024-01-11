# \SlackSettingsAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSlackDefaultNotificationSettings**](SlackSettingsAPI.md#CreateSlackDefaultNotificationSettings) | **Post** /orgs/{org_id}/slack_app/{bot_id} | Create new Slack notification default settings.
[**CreateSlackProjectNotificationSettings**](SlackSettingsAPI.md#CreateSlackProjectNotificationSettings) | **Post** /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} | Create a new Slack settings override for a given project.
[**DeleteSlackDefaultNotificationSettings**](SlackSettingsAPI.md#DeleteSlackDefaultNotificationSettings) | **Delete** /orgs/{org_id}/slack_app/{bot_id} | Remove the given Slack App integration
[**DeleteSlackProjectNotificationSettings**](SlackSettingsAPI.md#DeleteSlackProjectNotificationSettings) | **Delete** /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} | Remove Slack settings override for a project.
[**GetSlackDefaultNotificationSettings**](SlackSettingsAPI.md#GetSlackDefaultNotificationSettings) | **Get** /orgs/{org_id}/slack_app/{bot_id} | Get Slack integration default notification settings.
[**GetSlackProjectNotificationSettingsCollection**](SlackSettingsAPI.md#GetSlackProjectNotificationSettingsCollection) | **Get** /orgs/{org_id}/slack_app/{bot_id}/projects | Slack notification settings overrides for projects
[**UpdateSlackProjectNotificationSettings**](SlackSettingsAPI.md#UpdateSlackProjectNotificationSettings) | **Patch** /orgs/{org_id}/slack_app/{bot_id}/projects/{project_id} | Update Slack notification settings for a project.



## CreateSlackDefaultNotificationSettings

> GetSlackDefaultNotificationSettings200Response CreateSlackDefaultNotificationSettings(ctx, orgId, botId).Version(version).SettingsRequest(settingsRequest).Execute()

Create new Slack notification default settings.



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
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Bot ID
	settingsRequest := *openapiclient.NewSettingsRequest(*openapiclient.NewSettingsRequestData(*openapiclient.NewSettingsAttributes(openapiclient.SeverityThreshold("low"), "slack://channel?team=team-id&id=channel-id"), "Type_example")) // SettingsRequest | Create new Slack notification default settings for a tenant. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackSettingsAPI.CreateSlackDefaultNotificationSettings(context.Background(), orgId, botId).Version(version).SettingsRequest(settingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackSettingsAPI.CreateSlackDefaultNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSlackDefaultNotificationSettings`: GetSlackDefaultNotificationSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SlackSettingsAPI.CreateSlackDefaultNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**botId** | **string** | Bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlackDefaultNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **settingsRequest** | [**SettingsRequest**](SettingsRequest.md) | Create new Slack notification default settings for a tenant. | 

### Return type

[**GetSlackDefaultNotificationSettings200Response**](GetSlackDefaultNotificationSettings200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSlackProjectNotificationSettings

> CreateSlackProjectNotificationSettings201Response CreateSlackProjectNotificationSettings(ctx, orgId, projectId, botId).Version(version).SettingsRequest(settingsRequest).Execute()

Create a new Slack settings override for a given project.



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Bot ID
	settingsRequest := *openapiclient.NewSettingsRequest(*openapiclient.NewSettingsRequestData(*openapiclient.NewSettingsAttributes(openapiclient.SeverityThreshold("low"), "slack://channel?team=team-id&id=channel-id"), "Type_example")) // SettingsRequest | Create new Slack notification default settings for a tenant. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackSettingsAPI.CreateSlackProjectNotificationSettings(context.Background(), orgId, projectId, botId).Version(version).SettingsRequest(settingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackSettingsAPI.CreateSlackProjectNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSlackProjectNotificationSettings`: CreateSlackProjectNotificationSettings201Response
	fmt.Fprintf(os.Stdout, "Response from `SlackSettingsAPI.CreateSlackProjectNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**projectId** | **string** | Project ID | 
**botId** | **string** | Bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlackProjectNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 



 **settingsRequest** | [**SettingsRequest**](SettingsRequest.md) | Create new Slack notification default settings for a tenant. | 

### Return type

[**CreateSlackProjectNotificationSettings201Response**](CreateSlackProjectNotificationSettings201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSlackDefaultNotificationSettings

> DeleteSlackDefaultNotificationSettings(ctx, orgId, botId).Version(version).Execute()

Remove the given Slack App integration



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
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Bot ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackSettingsAPI.DeleteSlackDefaultNotificationSettings(context.Background(), orgId, botId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackSettingsAPI.DeleteSlackDefaultNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**botId** | **string** | Bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlackDefaultNotificationSettingsRequest struct via the builder pattern


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


## DeleteSlackProjectNotificationSettings

> DeleteSlackProjectNotificationSettings(ctx, orgId, projectId, botId).Version(version).Execute()

Remove Slack settings override for a project.



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Bot ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlackSettingsAPI.DeleteSlackProjectNotificationSettings(context.Background(), orgId, projectId, botId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackSettingsAPI.DeleteSlackProjectNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**projectId** | **string** | Project ID | 
**botId** | **string** | Bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlackProjectNotificationSettingsRequest struct via the builder pattern


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


## GetSlackDefaultNotificationSettings

> GetSlackDefaultNotificationSettings200Response GetSlackDefaultNotificationSettings(ctx, orgId, botId).Version(version).Execute()

Get Slack integration default notification settings.



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
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Bot ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackSettingsAPI.GetSlackDefaultNotificationSettings(context.Background(), orgId, botId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackSettingsAPI.GetSlackDefaultNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlackDefaultNotificationSettings`: GetSlackDefaultNotificationSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SlackSettingsAPI.GetSlackDefaultNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**botId** | **string** | Bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackDefaultNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 



### Return type

[**GetSlackDefaultNotificationSettings200Response**](GetSlackDefaultNotificationSettings200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackProjectNotificationSettingsCollection

> GetProjectSettingsCollection GetSlackProjectNotificationSettingsCollection(ctx, orgId, botId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Slack notification settings overrides for projects



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
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Bot ID
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackSettingsAPI.GetSlackProjectNotificationSettingsCollection(context.Background(), orgId, botId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackSettingsAPI.GetSlackProjectNotificationSettingsCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlackProjectNotificationSettingsCollection`: GetProjectSettingsCollection
	fmt.Fprintf(os.Stdout, "Response from `SlackSettingsAPI.GetSlackProjectNotificationSettingsCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**botId** | **string** | Bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackProjectNotificationSettingsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetProjectSettingsCollection**](GetProjectSettingsCollection.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlackProjectNotificationSettings

> CreateSlackProjectNotificationSettings201Response UpdateSlackProjectNotificationSettings(ctx, orgId, botId, projectId).Version(version).ProjectSettingsPatchRequest(projectSettingsPatchRequest).Execute()

Update Slack notification settings for a project.



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
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Bot ID
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
	projectSettingsPatchRequest := *openapiclient.NewProjectSettingsPatchRequest() // ProjectSettingsPatchRequest | Update existing project specific settings for a project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlackSettingsAPI.UpdateSlackProjectNotificationSettings(context.Background(), orgId, botId, projectId).Version(version).ProjectSettingsPatchRequest(projectSettingsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlackSettingsAPI.UpdateSlackProjectNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSlackProjectNotificationSettings`: CreateSlackProjectNotificationSettings201Response
	fmt.Fprintf(os.Stdout, "Response from `SlackSettingsAPI.UpdateSlackProjectNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**botId** | **string** | Bot ID | 
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSlackProjectNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 



 **projectSettingsPatchRequest** | [**ProjectSettingsPatchRequest**](ProjectSettingsPatchRequest.md) | Update existing project specific settings for a project. | 

### Return type

[**CreateSlackProjectNotificationSettings201Response**](CreateSlackProjectNotificationSettings201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

