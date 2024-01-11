# \AuditLogsAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGroupAuditLogs**](AuditLogsAPI.md#ListGroupAuditLogs) | **Get** /groups/{group_id}/audit_logs/search | Search Group audit logs.
[**ListOrgAuditLogs**](AuditLogsAPI.md#ListOrgAuditLogs) | **Get** /orgs/{org_id}/audit_logs/search | Search Organization audit logs.



## ListGroupAuditLogs

> ListGroupAuditLogs200Response ListGroupAuditLogs(ctx, groupId).Version(version).Cursor(cursor).From(from).To(to).Size(size).SortOrder(sortOrder).UserId(userId).ProjectId(projectId).Event(event).ExcludeEvent(excludeEvent).Execute()

Search Group audit logs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	groupId := "0d3728ec-eebf-484d-9907-ba238019f10b" // string | The ID of the Group.
	cursor := "cursor_example" // string | The ID for the next page of results. (optional)
	from := time.Now() // string | The start date (inclusive) of the audit logs search. Example: 2023-07-27  (optional)
	to := time.Now() // string | The end date (inclusive) of the audit logs search. Example: 2023-07-27  (optional)
	size := int32(10) // int32 | Number of results to return per page. (optional)
	sortOrder := "ASC" // string | Order in which results are returned. (optional) (default to "DESC")
	userId := "0d3728ec-eebf-484d-9907-ba238019f10b" // string | Filter logs by user ID. (optional)
	projectId := "0d3728ec-eebf-484d-9907-ba238019f10b" // string | Filter logs by project ID. (optional)
	event := "api.access" // string | Filter logs by event type, cannot be used in conjunction with exclude_event parameter. (optional)
	excludeEvent := "api.access" // string | Exclude event type from results, cannot be used in conjunctions with event parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.ListGroupAuditLogs(context.Background(), groupId).Version(version).Cursor(cursor).From(from).To(to).Size(size).SortOrder(sortOrder).UserId(userId).ProjectId(projectId).Event(event).ExcludeEvent(excludeEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.ListGroupAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupAuditLogs`: ListGroupAuditLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.ListGroupAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **cursor** | **string** | The ID for the next page of results. | 
 **from** | **string** | The start date (inclusive) of the audit logs search. Example: 2023-07-27  | 
 **to** | **string** | The end date (inclusive) of the audit logs search. Example: 2023-07-27  | 
 **size** | **int32** | Number of results to return per page. | 
 **sortOrder** | **string** | Order in which results are returned. | [default to &quot;DESC&quot;]
 **userId** | **string** | Filter logs by user ID. | 
 **projectId** | **string** | Filter logs by project ID. | 
 **event** | **string** | Filter logs by event type, cannot be used in conjunction with exclude_event parameter. | 
 **excludeEvent** | **string** | Exclude event type from results, cannot be used in conjunctions with event parameter. | 

### Return type

[**ListGroupAuditLogs200Response**](ListGroupAuditLogs200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgAuditLogs

> ListGroupAuditLogs200Response ListOrgAuditLogs(ctx, orgId).Version(version).Cursor(cursor).From(from).To(to).Size(size).SortOrder(sortOrder).UserId(userId).ProjectId(projectId).Event(event).ExcludeEvent(excludeEvent).Execute()

Search Organization audit logs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	orgId := "0d3728ec-eebf-484d-9907-ba238019f10b" // string | The ID of the organization.
	cursor := "cursor_example" // string | The ID for the next page of results. (optional)
	from := time.Now() // string | The start date (inclusive) of the audit logs search. Example: 2023-07-27  (optional)
	to := time.Now() // string | The end date (inclusive) of the audit logs search. Example: 2023-07-27  (optional)
	size := int32(10) // int32 | Number of results to return per page. (optional)
	sortOrder := "ASC" // string | Order in which results are returned. (optional) (default to "DESC")
	userId := "0d3728ec-eebf-484d-9907-ba238019f10b" // string | Filter logs by user ID. (optional)
	projectId := "0d3728ec-eebf-484d-9907-ba238019f10b" // string | Filter logs by project ID. (optional)
	event := "api.access" // string | Filter logs by event type, cannot be used in conjunction with exclude_event parameter. (optional)
	excludeEvent := "api.access" // string | Exclude event type from results, cannot be used in conjunctions with event parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.ListOrgAuditLogs(context.Background(), orgId).Version(version).Cursor(cursor).From(from).To(to).Size(size).SortOrder(sortOrder).UserId(userId).ProjectId(projectId).Event(event).ExcludeEvent(excludeEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.ListOrgAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgAuditLogs`: ListGroupAuditLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.ListOrgAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **cursor** | **string** | The ID for the next page of results. | 
 **from** | **string** | The start date (inclusive) of the audit logs search. Example: 2023-07-27  | 
 **to** | **string** | The end date (inclusive) of the audit logs search. Example: 2023-07-27  | 
 **size** | **int32** | Number of results to return per page. | 
 **sortOrder** | **string** | Order in which results are returned. | [default to &quot;DESC&quot;]
 **userId** | **string** | Filter logs by user ID. | 
 **projectId** | **string** | Filter logs by project ID. | 
 **event** | **string** | Filter logs by event type, cannot be used in conjunction with exclude_event parameter. | 
 **excludeEvent** | **string** | Exclude event type from results, cannot be used in conjunctions with event parameter. | 

### Return type

[**ListGroupAuditLogs200Response**](ListGroupAuditLogs200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

