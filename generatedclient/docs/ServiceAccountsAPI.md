# \ServiceAccountsAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupServiceAccount**](ServiceAccountsAPI.md#CreateGroupServiceAccount) | **Post** /groups/{group_id}/service_accounts | Create a service account for a group.
[**CreateOrgServiceAccount**](ServiceAccountsAPI.md#CreateOrgServiceAccount) | **Post** /orgs/{org_id}/service_accounts | Create a service account for an organization.
[**DeleteOneGroupServiceAccount**](ServiceAccountsAPI.md#DeleteOneGroupServiceAccount) | **Delete** /groups/{group_id}/service_accounts/{serviceaccount_id} | Delete a group service account.
[**DeleteServiceAccount**](ServiceAccountsAPI.md#DeleteServiceAccount) | **Delete** /orgs/{org_id}/service_accounts/{serviceaccount_id} | Delete a service account in an organization.
[**GetManyGroupServiceAccount**](ServiceAccountsAPI.md#GetManyGroupServiceAccount) | **Get** /groups/{group_id}/service_accounts | Get a list of group service accounts.
[**GetManyOrgServiceAccounts**](ServiceAccountsAPI.md#GetManyOrgServiceAccounts) | **Get** /orgs/{org_id}/service_accounts | Get a list of organization service accounts.
[**GetOneGroupServiceAccount**](ServiceAccountsAPI.md#GetOneGroupServiceAccount) | **Get** /groups/{group_id}/service_accounts/{serviceaccount_id} | Get a group service account.
[**GetOneOrgServiceAccount**](ServiceAccountsAPI.md#GetOneOrgServiceAccount) | **Get** /orgs/{org_id}/service_accounts/{serviceaccount_id} | Get an organization service account.
[**UpdateGroupServiceAccount**](ServiceAccountsAPI.md#UpdateGroupServiceAccount) | **Patch** /groups/{group_id}/service_accounts/{serviceaccount_id} | Update a group service account.
[**UpdateOrgServiceAccount**](ServiceAccountsAPI.md#UpdateOrgServiceAccount) | **Patch** /orgs/{org_id}/service_accounts/{serviceaccount_id} | Update an organization service account.
[**UpdateOrgServiceAccountSecret**](ServiceAccountsAPI.md#UpdateOrgServiceAccountSecret) | **Post** /orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets | Manage an organization service account&#39;s client secret.
[**UpdateServiceAccountSecret**](ServiceAccountsAPI.md#UpdateServiceAccountSecret) | **Post** /groups/{group_id}/service_accounts/{serviceaccount_id}/secrets | Manage a group service account&#39;s client secret.



## CreateGroupServiceAccount

> CreateGroupServiceAccount201Response CreateGroupServiceAccount(ctx, groupId).Version(version).CreateGroupServiceAccountRequest(createGroupServiceAccountRequest).Execute()

Create a service account for a group.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Group that is creating and owns the service account
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	createGroupServiceAccountRequest := *openapiclient.NewCreateGroupServiceAccountRequest(*openapiclient.NewCreateGroupServiceAccountRequestData(*openapiclient.NewCreateGroupServiceAccountRequestDataAttributes("AuthType_example", "Name_example", "RoleId_example"), "Type_example")) // CreateGroupServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.CreateGroupServiceAccount(context.Background(), groupId).Version(version).CreateGroupServiceAccountRequest(createGroupServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.CreateGroupServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupServiceAccount`: CreateGroupServiceAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.CreateGroupServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the Snyk Group that is creating and owns the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **createGroupServiceAccountRequest** | [**CreateGroupServiceAccountRequest**](CreateGroupServiceAccountRequest.md) |  | 

### Return type

[**CreateGroupServiceAccount201Response**](CreateGroupServiceAccount201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgServiceAccount

> GetOneGroupServiceAccount200Response CreateOrgServiceAccount(ctx, orgId).Version(version).CreateOrgServiceAccountRequest(createOrgServiceAccountRequest).Execute()

Create a service account for an organization.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Organization that is creating and will own the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	createOrgServiceAccountRequest := *openapiclient.NewCreateOrgServiceAccountRequest(*openapiclient.NewCreateOrgServiceAccountRequestData(*openapiclient.NewCreateOrgServiceAccountRequestDataAttributes("AuthType_example", "Name_example", "RoleId_example"))) // CreateOrgServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.CreateOrgServiceAccount(context.Background(), orgId).Version(version).CreateOrgServiceAccountRequest(createOrgServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.CreateOrgServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgServiceAccount`: GetOneGroupServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.CreateOrgServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the Snyk Organization that is creating and will own the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **createOrgServiceAccountRequest** | [**CreateOrgServiceAccountRequest**](CreateOrgServiceAccountRequest.md) |  | 

### Return type

[**GetOneGroupServiceAccount200Response**](GetOneGroupServiceAccount200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOneGroupServiceAccount

> DeleteOneGroupServiceAccount(ctx, groupId, serviceaccountId).Version(version).Execute()

Delete a group service account.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Group that owns the service account.
	serviceaccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceAccountsAPI.DeleteOneGroupServiceAccount(context.Background(), groupId, serviceaccountId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.DeleteOneGroupServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the Snyk Group that owns the service account. | 
**serviceaccountId** | **string** | The ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOneGroupServiceAccountRequest struct via the builder pattern


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


## DeleteServiceAccount

> DeleteServiceAccount(ctx, orgId, serviceaccountId).Version(version).Execute()

Delete a service account in an organization.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of org to which the service account belongs.
	serviceaccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceAccountsAPI.DeleteServiceAccount(context.Background(), orgId, serviceaccountId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.DeleteServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of org to which the service account belongs. | 
**serviceaccountId** | **string** | The ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


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


## GetManyGroupServiceAccount

> GetManyGroupServiceAccount200Response GetManyGroupServiceAccount(ctx, groupId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of group service accounts.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Group that owns the service accounts.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.GetManyGroupServiceAccount(context.Background(), groupId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.GetManyGroupServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManyGroupServiceAccount`: GetManyGroupServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.GetManyGroupServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the Snyk Group that owns the service accounts. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManyGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetManyGroupServiceAccount200Response**](GetManyGroupServiceAccount200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManyOrgServiceAccounts

> GetManyGroupServiceAccount200Response GetManyOrgServiceAccounts(ctx, orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

Get a list of organization service accounts.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Organization that owns the service accounts.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.GetManyOrgServiceAccounts(context.Background(), orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.GetManyOrgServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManyOrgServiceAccounts`: GetManyGroupServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.GetManyOrgServiceAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the Snyk Organization that owns the service accounts. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManyOrgServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**GetManyGroupServiceAccount200Response**](GetManyGroupServiceAccount200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneGroupServiceAccount

> GetOneGroupServiceAccount200Response GetOneGroupServiceAccount(ctx, groupId, serviceaccountId).Version(version).Execute()

Get a group service account.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Group that owns the service account.
	serviceaccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.GetOneGroupServiceAccount(context.Background(), groupId, serviceaccountId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.GetOneGroupServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOneGroupServiceAccount`: GetOneGroupServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.GetOneGroupServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the Snyk Group that owns the service account. | 
**serviceaccountId** | **string** | The ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOneGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 

### Return type

[**GetOneGroupServiceAccount200Response**](GetOneGroupServiceAccount200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneOrgServiceAccount

> GetOneGroupServiceAccount200Response GetOneOrgServiceAccount(ctx, orgId, serviceaccountId).Version(version).Execute()

Get an organization service account.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Organization that owns the service account.
	serviceaccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.GetOneOrgServiceAccount(context.Background(), orgId, serviceaccountId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.GetOneOrgServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOneOrgServiceAccount`: GetOneGroupServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.GetOneOrgServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the Snyk Organization that owns the service account. | 
**serviceaccountId** | **string** | The ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOneOrgServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 

### Return type

[**GetOneGroupServiceAccount200Response**](GetOneGroupServiceAccount200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupServiceAccount

> GetOneGroupServiceAccount200Response UpdateGroupServiceAccount(ctx, groupId, serviceaccountId).Version(version).UpdateGroupServiceAccountRequest(updateGroupServiceAccountRequest).Execute()

Update a group service account.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Group that owns the service account.
	serviceaccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	updateGroupServiceAccountRequest := *openapiclient.NewUpdateGroupServiceAccountRequest(*openapiclient.NewUpdateGroupServiceAccountRequestData(*openapiclient.NewUpdateGroupServiceAccountRequestDataAttributes("Name_example"), "Id_example", "Type_example")) // UpdateGroupServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.UpdateGroupServiceAccount(context.Background(), groupId, serviceaccountId).Version(version).UpdateGroupServiceAccountRequest(updateGroupServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.UpdateGroupServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupServiceAccount`: GetOneGroupServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.UpdateGroupServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the Snyk Group that owns the service account. | 
**serviceaccountId** | **string** | The ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 
 **updateGroupServiceAccountRequest** | [**UpdateGroupServiceAccountRequest**](UpdateGroupServiceAccountRequest.md) |  | 

### Return type

[**GetOneGroupServiceAccount200Response**](GetOneGroupServiceAccount200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgServiceAccount

> GetOneGroupServiceAccount200Response UpdateOrgServiceAccount(ctx, orgId, serviceaccountId).Version(version).UpdateOrgServiceAccountRequest(updateOrgServiceAccountRequest).Execute()

Update an organization service account.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Organization that owns the service account.
	serviceaccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	updateOrgServiceAccountRequest := *openapiclient.NewUpdateOrgServiceAccountRequest(*openapiclient.NewUpdateOrgServiceAccountRequestData(*openapiclient.NewUpdateOrgServiceAccountRequestDataAttributes("Name_example"), "Id_example", "Type_example")) // UpdateOrgServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.UpdateOrgServiceAccount(context.Background(), orgId, serviceaccountId).Version(version).UpdateOrgServiceAccountRequest(updateOrgServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.UpdateOrgServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgServiceAccount`: GetOneGroupServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.UpdateOrgServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the Snyk Organization that owns the service account. | 
**serviceaccountId** | **string** | The ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 
 **updateOrgServiceAccountRequest** | [**UpdateOrgServiceAccountRequest**](UpdateOrgServiceAccountRequest.md) |  | 

### Return type

[**GetOneGroupServiceAccount200Response**](GetOneGroupServiceAccount200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgServiceAccountSecret

> CreateGroupServiceAccount201Response UpdateOrgServiceAccountSecret(ctx, orgId, serviceaccountId).Version(version).UpdateServiceAccountSecretRequest(updateServiceAccountSecretRequest).Execute()

Manage an organization service account's client secret.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Organization that owns the service account.
	serviceaccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	updateServiceAccountSecretRequest := *openapiclient.NewUpdateServiceAccountSecretRequest(*openapiclient.NewUpdateServiceAccountSecretRequestData(*openapiclient.NewUpdateServiceAccountSecretRequestDataAttributes("Mode_example"), "Type_example")) // UpdateServiceAccountSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.UpdateOrgServiceAccountSecret(context.Background(), orgId, serviceaccountId).Version(version).UpdateServiceAccountSecretRequest(updateServiceAccountSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.UpdateOrgServiceAccountSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgServiceAccountSecret`: CreateGroupServiceAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.UpdateOrgServiceAccountSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the Snyk Organization that owns the service account. | 
**serviceaccountId** | **string** | The ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgServiceAccountSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 
 **updateServiceAccountSecretRequest** | [**UpdateServiceAccountSecretRequest**](UpdateServiceAccountSecretRequest.md) |  | 

### Return type

[**CreateGroupServiceAccount201Response**](CreateGroupServiceAccount201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccountSecret

> CreateGroupServiceAccount201Response UpdateServiceAccountSecret(ctx, groupId, serviceaccountId).Version(version).UpdateServiceAccountSecretRequest(updateServiceAccountSecretRequest).Execute()

Manage a group service account's client secret.



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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Snyk Group that owns the service account.
	serviceaccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service account.
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	updateServiceAccountSecretRequest := *openapiclient.NewUpdateServiceAccountSecretRequest(*openapiclient.NewUpdateServiceAccountSecretRequestData(*openapiclient.NewUpdateServiceAccountSecretRequestDataAttributes("Mode_example"), "Type_example")) // UpdateServiceAccountSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountsAPI.UpdateServiceAccountSecret(context.Background(), groupId, serviceaccountId).Version(version).UpdateServiceAccountSecretRequest(updateServiceAccountSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsAPI.UpdateServiceAccountSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceAccountSecret`: CreateGroupServiceAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsAPI.UpdateServiceAccountSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the Snyk Group that owns the service account. | 
**serviceaccountId** | **string** | The ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 
 **updateServiceAccountSecretRequest** | [**UpdateServiceAccountSecretRequest**](UpdateServiceAccountSecretRequest.md) |  | 

### Return type

[**CreateGroupServiceAccount201Response**](CreateGroupServiceAccount201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

