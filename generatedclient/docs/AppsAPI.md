# \AppsAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppsAPI.md#CreateApp) | **Post** /orgs/{org_id}/apps | Create a new app for an organization.
[**CreateGroupAppInstall**](AppsAPI.md#CreateGroupAppInstall) | **Post** /groups/{group_id}/apps/installs | Install a Snyk Apps to this group.
[**CreateOrgApp**](AppsAPI.md#CreateOrgApp) | **Post** /orgs/{org_id}/apps/creations | Create a new Snyk App for an organization.
[**CreateOrgAppInstall**](AppsAPI.md#CreateOrgAppInstall) | **Post** /orgs/{org_id}/apps/installs | Install a Snyk Apps to this organization.
[**DeleteApp**](AppsAPI.md#DeleteApp) | **Delete** /orgs/{org_id}/apps/{client_id} | Delete an app
[**DeleteAppBot**](AppsAPI.md#DeleteAppBot) | **Delete** /orgs/{org_id}/app_bots/{bot_id} | Revoke app bot authorization
[**DeleteAppByID**](AppsAPI.md#DeleteAppByID) | **Delete** /orgs/{org_id}/apps/creations/{app_id} | Delete an app by its App ID.
[**DeleteAppOrgInstallByID**](AppsAPI.md#DeleteAppOrgInstallByID) | **Delete** /orgs/{org_id}/apps/installs/{install_id} | Revoke app authorization for an Snyk Organization with install ID.
[**DeleteGroupAppInstallByID**](AppsAPI.md#DeleteGroupAppInstallByID) | **Delete** /groups/{group_id}/apps/installs/{install_id} | Revoke app authorization for an Snyk Group with install ID.
[**DeleteUserAppInstallByID**](AppsAPI.md#DeleteUserAppInstallByID) | **Delete** /self/apps/installs/{install_id} | Revoke access for an app by install ID.
[**GetApp**](AppsAPI.md#GetApp) | **Get** /orgs/{org_id}/apps/{client_id} | Get an app by client id
[**GetAppBots**](AppsAPI.md#GetAppBots) | **Get** /orgs/{org_id}/app_bots | Get a list of app bots authorized to an organization.
[**GetAppByID**](AppsAPI.md#GetAppByID) | **Get** /orgs/{org_id}/apps/creations/{app_id} | Get a Snyk App by its App ID.
[**GetAppInstallsForGroup**](AppsAPI.md#GetAppInstallsForGroup) | **Get** /groups/{group_id}/apps/installs | Get a list of apps installed for a group.
[**GetAppInstallsForOrg**](AppsAPI.md#GetAppInstallsForOrg) | **Get** /orgs/{org_id}/apps/installs | Get a list of apps installed for an organization.
[**GetAppInstallsForUser**](AppsAPI.md#GetAppInstallsForUser) | **Get** /self/apps/installs | Get a list of apps installed for an user.
[**GetApps**](AppsAPI.md#GetApps) | **Get** /orgs/{org_id}/apps | Get a list of apps created by an organization.
[**GetOrgApps**](AppsAPI.md#GetOrgApps) | **Get** /orgs/{org_id}/apps/creations | Get a list of apps created by an organization.
[**GetUserAppSessions**](AppsAPI.md#GetUserAppSessions) | **Get** /self/apps/{app_id}/sessions | Get a list of active OAuth sessions for the app.
[**GetUserInstalledApps**](AppsAPI.md#GetUserInstalledApps) | **Get** /self/apps | Get a list of apps that can act on your behalf.
[**ManageAppCreationSecret**](AppsAPI.md#ManageAppCreationSecret) | **Post** /orgs/{org_id}/apps/creations/{app_id}/secrets | Manage client secret for the Snyk App.
[**ManageSecrets**](AppsAPI.md#ManageSecrets) | **Post** /orgs/{org_id}/apps/{client_id}/secrets | Manage client secrets for an app.
[**RevokeUserAppSession**](AppsAPI.md#RevokeUserAppSession) | **Delete** /self/apps/{app_id}/sessions/{session_id} | Revoke an active user app session.
[**RevokeUserInstalledApp**](AppsAPI.md#RevokeUserInstalledApp) | **Delete** /self/apps/{app_id} | Revoke an app
[**UpdateApp**](AppsAPI.md#UpdateApp) | **Patch** /orgs/{org_id}/apps/{client_id} | Update app attributes that are name, redirect URIs, and access token time to live
[**UpdateAppCreationByID**](AppsAPI.md#UpdateAppCreationByID) | **Patch** /orgs/{org_id}/apps/creations/{app_id} | Update app creation attributes such as name, redirect URIs, and access token time to live using the App ID.
[**UpdateGroupAppInstallSecret**](AppsAPI.md#UpdateGroupAppInstallSecret) | **Post** /groups/{group_id}/apps/installs/{install_id}/secrets | Manage client secret for non-interactive Snyk App installations.
[**UpdateOrgAppInstallSecret**](AppsAPI.md#UpdateOrgAppInstallSecret) | **Post** /orgs/{org_id}/apps/installs/{install_id}/secrets | Manage client secret for non-interactive Snyk App installations.



## CreateApp

> AppPostResponse CreateApp(ctx, orgId).Version(version).AppPostRequest(appPostRequest).Execute()

Create a new app for an organization.



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
	appPostRequest := *openapiclient.NewAppPostRequest(*openapiclient.NewAppPostRequestData(*openapiclient.NewAppPostRequestDataAttributes(openapiclient.Context("tenant"), "My App", []string{"RedirectUris_example"}, []string{"Scopes_example"}), "Type_example")) // AppPostRequest | app to be created (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.CreateApp(context.Background(), orgId).Version(version).AppPostRequest(appPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CreateApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApp`: AppPostResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CreateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **appPostRequest** | [**AppPostRequest**](AppPostRequest.md) | app to be created | 

### Return type

[**AppPostResponse**](AppPostResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupAppInstall

> CreateGroupAppInstall201Response CreateGroupAppInstall(ctx, groupId).Version(version).CreateGroupAppInstallRequest(createGroupAppInstallRequest).Execute()

Install a Snyk Apps to this group.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Group ID
	createGroupAppInstallRequest := *openapiclient.NewCreateGroupAppInstallRequest(*openapiclient.NewCreateGroupAppInstallRequestData(), *openapiclient.NewCreateGroupAppInstallRequestRelationships(*openapiclient.NewCreateGroupAppInstallRequestRelationshipsApp(*openapiclient.NewCreateGroupAppInstallRequestRelationshipsAppData("Id_example", "app")))) // CreateGroupAppInstallRequest | App Install to be created (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.CreateGroupAppInstall(context.Background(), groupId).Version(version).CreateGroupAppInstallRequest(createGroupAppInstallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CreateGroupAppInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupAppInstall`: CreateGroupAppInstall201Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CreateGroupAppInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupAppInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **createGroupAppInstallRequest** | [**CreateGroupAppInstallRequest**](CreateGroupAppInstallRequest.md) | App Install to be created | 

### Return type

[**CreateGroupAppInstall201Response**](CreateGroupAppInstall201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgApp

> AppPostResponse CreateOrgApp(ctx, orgId).Version(version).AppPostRequest(appPostRequest).Execute()

Create a new Snyk App for an organization.



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
	appPostRequest := *openapiclient.NewAppPostRequest(*openapiclient.NewAppPostRequestData(*openapiclient.NewAppPostRequestDataAttributes(openapiclient.Context("tenant"), "My App", []string{"RedirectUris_example"}, []string{"Scopes_example"}), "Type_example")) // AppPostRequest | Snyk App details for app to be created. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.CreateOrgApp(context.Background(), orgId).Version(version).AppPostRequest(appPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CreateOrgApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgApp`: AppPostResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CreateOrgApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **appPostRequest** | [**AppPostRequest**](AppPostRequest.md) | Snyk App details for app to be created. | 

### Return type

[**AppPostResponse**](AppPostResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgAppInstall

> CreateGroupAppInstall201Response CreateOrgAppInstall(ctx, orgId).Version(version).CreateGroupAppInstallRequest(createGroupAppInstallRequest).Execute()

Install a Snyk Apps to this organization.



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
	createGroupAppInstallRequest := *openapiclient.NewCreateGroupAppInstallRequest(*openapiclient.NewCreateGroupAppInstallRequestData(), *openapiclient.NewCreateGroupAppInstallRequestRelationships(*openapiclient.NewCreateGroupAppInstallRequestRelationshipsApp(*openapiclient.NewCreateGroupAppInstallRequestRelationshipsAppData("Id_example", "app")))) // CreateGroupAppInstallRequest | App Install to be created (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.CreateOrgAppInstall(context.Background(), orgId).Version(version).CreateGroupAppInstallRequest(createGroupAppInstallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CreateOrgAppInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgAppInstall`: CreateGroupAppInstall201Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CreateOrgAppInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgAppInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **createGroupAppInstallRequest** | [**CreateGroupAppInstallRequest**](CreateGroupAppInstallRequest.md) | App Install to be created | 

### Return type

[**CreateGroupAppInstall201Response**](CreateGroupAppInstall201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> DeleteApp(ctx, orgId, clientId).Version(version).Execute()

Delete an app



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
	clientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.DeleteApp(context.Background(), orgId, clientId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


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


## DeleteAppBot

> DeleteAppBot(ctx, botId, orgId).Version(version).Execute()

Revoke app bot authorization



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
	botId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the app bot
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.DeleteAppBot(context.Background(), botId, orgId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteAppBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | The ID of the app bot | 
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppBotRequest struct via the builder pattern


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


## DeleteAppByID

> DeleteAppByID(ctx, orgId, appId).Version(version).Execute()

Delete an app by its App ID.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.DeleteAppByID(context.Background(), orgId, appId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteAppByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**appId** | **string** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppByIDRequest struct via the builder pattern


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


## DeleteAppOrgInstallByID

> DeleteAppOrgInstallByID(ctx, orgId, installId).Version(version).Execute()

Revoke app authorization for an Snyk Organization with install ID.



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
	installId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Install ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.DeleteAppOrgInstallByID(context.Background(), orgId, installId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteAppOrgInstallByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**installId** | **string** | Install ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppOrgInstallByIDRequest struct via the builder pattern


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


## DeleteGroupAppInstallByID

> DeleteGroupAppInstallByID(ctx, groupId, installId).Version(version).Execute()

Revoke app authorization for an Snyk Group with install ID.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Group ID
	installId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Install ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.DeleteGroupAppInstallByID(context.Background(), groupId, installId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteGroupAppInstallByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 
**installId** | **string** | Install ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupAppInstallByIDRequest struct via the builder pattern


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


## DeleteUserAppInstallByID

> DeleteUserAppInstallByID(ctx, installId).Version(version).Execute()

Revoke access for an app by install ID.



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
	installId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Install ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.DeleteUserAppInstallByID(context.Background(), installId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteUserAppInstallByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installId** | **string** | Install ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserAppInstallByIDRequest struct via the builder pattern


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


## GetApp

> GetAppByID200Response GetApp(ctx, orgId, clientId).Version(version).Execute()

Get an app by client id



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	clientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Client ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetApp(context.Background(), orgId, clientId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApp`: GetAppByID200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 

### Return type

[**GetAppByID200Response**](GetAppByID200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppBots

> GetAppBots200Response GetAppBots(ctx, orgId).Version(version).Expand(expand).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of app bots authorized to an organization.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	expand := []string{"Expand_example"} // []string | Expand relationships. (optional)
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppBots(context.Background(), orgId).Version(version).Expand(expand).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppBots`: GetAppBots200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppBots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppBotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **expand** | **[]string** | Expand relationships. | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetAppBots200Response**](GetAppBots200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppByID

> GetAppByID200Response GetAppByID(ctx, orgId, appId).Version(version).Execute()

Get a Snyk App by its App ID.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppByID(context.Background(), orgId, appId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppByID`: GetAppByID200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**appId** | **string** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 

### Return type

[**GetAppByID200Response**](GetAppByID200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstallsForGroup

> GetAppInstallsForGroup200Response GetAppInstallsForGroup(ctx, groupId).Version(version).Expand(expand).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of apps installed for a group.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Group ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	expand := []string{"Expand_example"} // []string | Expand relationships. (optional)
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppInstallsForGroup(context.Background(), groupId).Version(version).Expand(expand).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppInstallsForGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstallsForGroup`: GetAppInstallsForGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppInstallsForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstallsForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **expand** | **[]string** | Expand relationships. | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetAppInstallsForGroup200Response**](GetAppInstallsForGroup200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstallsForOrg

> GetAppInstallsForGroup200Response GetAppInstallsForOrg(ctx, orgId).Version(version).Expand(expand).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of apps installed for an organization.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	expand := []string{"Expand_example"} // []string | Expand relationships. (optional)
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppInstallsForOrg(context.Background(), orgId).Version(version).Expand(expand).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppInstallsForOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstallsForOrg`: GetAppInstallsForGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppInstallsForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstallsForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **expand** | **[]string** | Expand relationships. | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetAppInstallsForGroup200Response**](GetAppInstallsForGroup200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppInstallsForUser

> GetAppInstallsForGroup200Response GetAppInstallsForUser(ctx).Version(version).Expand(expand).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of apps installed for an user.



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
	expand := []string{"Expand_example"} // []string | Expand relationships. (optional)
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppInstallsForUser(context.Background()).Version(version).Expand(expand).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppInstallsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppInstallsForUser`: GetAppInstallsForGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppInstallsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppInstallsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 
 **expand** | **[]string** | Expand relationships. | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetAppInstallsForGroup200Response**](GetAppInstallsForGroup200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApps

> GetApps200Response GetApps(ctx, orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of apps created by an organization.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetApps(context.Background(), orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApps`: GetApps200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetApps200Response**](GetApps200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgApps

> GetApps200Response GetOrgApps(ctx, orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of apps created by an organization.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Org ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetOrgApps(context.Background(), orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetOrgApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgApps`: GetApps200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetOrgApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetApps200Response**](GetApps200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAppSessions

> GetUserAppSessions200Response GetUserAppSessions(ctx, appId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of active OAuth sessions for the app.



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
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App ID
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetUserAppSessions(context.Background(), appId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetUserAppSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAppSessions`: GetUserAppSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetUserAppSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAppSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetUserAppSessions200Response**](GetUserAppSessions200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInstalledApps

> GetUserInstalledApps200Response GetUserInstalledApps(ctx).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of apps that can act on your behalf.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetUserInstalledApps(context.Background()).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetUserInstalledApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInstalledApps`: GetUserInstalledApps200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetUserInstalledApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetUserInstalledApps200Response**](GetUserInstalledApps200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageAppCreationSecret

> ManageAppCreationSecret200Response ManageAppCreationSecret(ctx, orgId, appId).Version(version).UpdateGroupAppInstallSecretRequest(updateGroupAppInstallSecretRequest).Execute()

Manage client secret for the Snyk App.



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
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App ID
	updateGroupAppInstallSecretRequest := *openapiclient.NewUpdateGroupAppInstallSecretRequest(*openapiclient.NewUpdateGroupAppInstallSecretRequestData(*openapiclient.NewUpdateGroupAppInstallSecretRequestDataAttributes("Mode_example"), "Type_example")) // UpdateGroupAppInstallSecretRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ManageAppCreationSecret(context.Background(), orgId, appId).Version(version).UpdateGroupAppInstallSecretRequest(updateGroupAppInstallSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ManageAppCreationSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManageAppCreationSecret`: ManageAppCreationSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ManageAppCreationSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**appId** | **string** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageAppCreationSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **updateGroupAppInstallSecretRequest** | [**UpdateGroupAppInstallSecretRequest**](UpdateGroupAppInstallSecretRequest.md) |  | 

### Return type

[**ManageAppCreationSecret200Response**](ManageAppCreationSecret200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageSecrets

> ManageAppCreationSecret200Response ManageSecrets(ctx, orgId, clientId).Version(version).UpdateGroupAppInstallSecretRequestDataAttributes(updateGroupAppInstallSecretRequestDataAttributes).Execute()

Manage client secrets for an app.



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
	clientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Client ID
	updateGroupAppInstallSecretRequestDataAttributes := *openapiclient.NewUpdateGroupAppInstallSecretRequestDataAttributes("Mode_example") // UpdateGroupAppInstallSecretRequestDataAttributes |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ManageSecrets(context.Background(), orgId, clientId).Version(version).UpdateGroupAppInstallSecretRequestDataAttributes(updateGroupAppInstallSecretRequestDataAttributes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ManageSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManageSecrets`: ManageAppCreationSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ManageSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **updateGroupAppInstallSecretRequestDataAttributes** | [**UpdateGroupAppInstallSecretRequestDataAttributes**](UpdateGroupAppInstallSecretRequestDataAttributes.md) |  | 

### Return type

[**ManageAppCreationSecret200Response**](ManageAppCreationSecret200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeUserAppSession

> RevokeUserAppSession(ctx, appId, sessionId).Version(version).Execute()

Revoke an active user app session.



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
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App ID
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Session ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.RevokeUserAppSession(context.Background(), appId, sessionId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.RevokeUserAppSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | App ID | 
**sessionId** | **string** | Session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserAppSessionRequest struct via the builder pattern


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


## RevokeUserInstalledApp

> RevokeUserInstalledApp(ctx, appId).Version(version).Execute()

Revoke an app



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
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.RevokeUserInstalledApp(context.Background(), appId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.RevokeUserInstalledApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserInstalledAppRequest struct via the builder pattern


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


## UpdateApp

> GetAppByID200Response UpdateApp(ctx, orgId, clientId).Version(version).AppPatchRequest(appPatchRequest).Execute()

Update app attributes that are name, redirect URIs, and access token time to live



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
	clientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Client ID
	appPatchRequest := *openapiclient.NewAppPatchRequest(*openapiclient.NewAppPatchRequestData()) // AppPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateApp(context.Background(), orgId, clientId).Version(version).AppPatchRequest(appPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApp`: GetAppByID200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **appPatchRequest** | [**AppPatchRequest**](AppPatchRequest.md) |  | 

### Return type

[**GetAppByID200Response**](GetAppByID200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppCreationByID

> GetAppByID200Response UpdateAppCreationByID(ctx, orgId, appId).Version(version).AppPatchRequest(appPatchRequest).Execute()

Update app creation attributes such as name, redirect URIs, and access token time to live using the App ID.



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
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | App ID
	appPatchRequest := *openapiclient.NewAppPatchRequest(*openapiclient.NewAppPatchRequestData()) // AppPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateAppCreationByID(context.Background(), orgId, appId).Version(version).AppPatchRequest(appPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateAppCreationByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppCreationByID`: GetAppByID200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateAppCreationByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**appId** | **string** | App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppCreationByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **appPatchRequest** | [**AppPatchRequest**](AppPatchRequest.md) |  | 

### Return type

[**GetAppByID200Response**](GetAppByID200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupAppInstallSecret

> UpdateGroupAppInstallSecret200Response UpdateGroupAppInstallSecret(ctx, groupId, installId).Version(version).UpdateGroupAppInstallSecretRequest(updateGroupAppInstallSecretRequest).Execute()

Manage client secret for non-interactive Snyk App installations.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Group ID
	installId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Install ID
	updateGroupAppInstallSecretRequest := *openapiclient.NewUpdateGroupAppInstallSecretRequest(*openapiclient.NewUpdateGroupAppInstallSecretRequestData(*openapiclient.NewUpdateGroupAppInstallSecretRequestDataAttributes("Mode_example"), "Type_example")) // UpdateGroupAppInstallSecretRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateGroupAppInstallSecret(context.Background(), groupId, installId).Version(version).UpdateGroupAppInstallSecretRequest(updateGroupAppInstallSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateGroupAppInstallSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupAppInstallSecret`: UpdateGroupAppInstallSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateGroupAppInstallSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID | 
**installId** | **string** | Install ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupAppInstallSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **updateGroupAppInstallSecretRequest** | [**UpdateGroupAppInstallSecretRequest**](UpdateGroupAppInstallSecretRequest.md) |  | 

### Return type

[**UpdateGroupAppInstallSecret200Response**](UpdateGroupAppInstallSecret200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgAppInstallSecret

> UpdateGroupAppInstallSecret200Response UpdateOrgAppInstallSecret(ctx, orgId, installId).Version(version).UpdateGroupAppInstallSecretRequest(updateGroupAppInstallSecretRequest).Execute()

Manage client secret for non-interactive Snyk App installations.



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
	installId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Install ID
	updateGroupAppInstallSecretRequest := *openapiclient.NewUpdateGroupAppInstallSecretRequest(*openapiclient.NewUpdateGroupAppInstallSecretRequestData(*openapiclient.NewUpdateGroupAppInstallSecretRequestDataAttributes("Mode_example"), "Type_example")) // UpdateGroupAppInstallSecretRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateOrgAppInstallSecret(context.Background(), orgId, installId).Version(version).UpdateGroupAppInstallSecretRequest(updateGroupAppInstallSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateOrgAppInstallSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgAppInstallSecret`: UpdateGroupAppInstallSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateOrgAppInstallSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**installId** | **string** | Install ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgAppInstallSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **updateGroupAppInstallSecretRequest** | [**UpdateGroupAppInstallSecretRequest**](UpdateGroupAppInstallSecretRequest.md) |  | 

### Return type

[**UpdateGroupAppInstallSecret200Response**](UpdateGroupAppInstallSecret200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

