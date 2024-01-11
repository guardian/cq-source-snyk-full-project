# \InvitesAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgInvitation**](InvitesAPI.md#CreateOrgInvitation) | **Post** /orgs/{org_id}/invites | Invite a user to an organization
[**DeleteOrgInvitation**](InvitesAPI.md#DeleteOrgInvitation) | **Delete** /orgs/{org_id}/invites/{invite_id} | Cancel a pending user invitations to an organization.
[**ListOrgInvitation**](InvitesAPI.md#ListOrgInvitation) | **Get** /orgs/{org_id}/invites | List pending user invitations to an organization.



## CreateOrgInvitation

> CreateOrgInvitation201Response CreateOrgInvitation(ctx, orgId).Version(version).CreateOrgInvitationRequest(createOrgInvitationRequest).Execute()

Invite a user to an organization



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the org the user is being invited to
	createOrgInvitationRequest := *openapiclient.NewCreateOrgInvitationRequest(*openapiclient.NewOrgInvitationPostData(*openapiclient.NewOrgInvitationPostAttributes("example@email.com", interface{}(f1968726-1dca-42d4-a4dc-80cab99e2b6c)), "Type_example")) // CreateOrgInvitationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.CreateOrgInvitation(context.Background(), orgId).Version(version).CreateOrgInvitationRequest(createOrgInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.CreateOrgInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgInvitation`: CreateOrgInvitation201Response
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.CreateOrgInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org the user is being invited to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **createOrgInvitationRequest** | [**CreateOrgInvitationRequest**](CreateOrgInvitationRequest.md) |  | 

### Return type

[**CreateOrgInvitation201Response**](CreateOrgInvitation201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgInvitation

> DeleteOrgInvitation(ctx, orgId, inviteId).Version(version).Execute()

Cancel a pending user invitations to an organization.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the org the user is being invited to
	inviteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the pending invite to cancel
	version := "2021-06-04" // string | The requested version of the endpoint to process the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvitesAPI.DeleteOrgInvitation(context.Background(), orgId, inviteId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.DeleteOrgInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org the user is being invited to | 
**inviteId** | **string** | The id of the pending invite to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgInvitationRequest struct via the builder pattern


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


## ListOrgInvitation

> ListOrgInvitation200Response ListOrgInvitation(ctx, orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()

List pending user invitations to an organization.



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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the org the user is being invited to
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.ListOrgInvitation(context.Background(), orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.ListOrgInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgInvitation`: ListOrgInvitation200Response
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.ListOrgInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of the org the user is being invited to | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]

### Return type

[**ListOrgInvitation200Response**](ListOrgInvitation200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

