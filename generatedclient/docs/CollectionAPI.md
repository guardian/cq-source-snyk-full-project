# \CollectionAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCollection**](CollectionAPI.md#CreateCollection) | **Post** /orgs/{org_id}/collections | Create a collection
[**DeleteCollection**](CollectionAPI.md#DeleteCollection) | **Delete** /orgs/{org_id}/collections/{collection_id} | Delete a collection
[**DeleteProjectsCollection**](CollectionAPI.md#DeleteProjectsCollection) | **Delete** /orgs/{org_id}/collections/{collection_id}/relationships/projects | Remove projects from a collection
[**GetCollection**](CollectionAPI.md#GetCollection) | **Get** /orgs/{org_id}/collections/{collection_id} | Get a collection
[**GetCollections**](CollectionAPI.md#GetCollections) | **Get** /orgs/{org_id}/collections | Get collections
[**GetProjectsOfCollection**](CollectionAPI.md#GetProjectsOfCollection) | **Get** /orgs/{org_id}/collections/{collection_id}/relationships/projects | Get projects from the specified collection
[**UpdateCollection**](CollectionAPI.md#UpdateCollection) | **Patch** /orgs/{org_id}/collections/{collection_id} | Edit a collection
[**UpdateCollectionWithProjects**](CollectionAPI.md#UpdateCollectionWithProjects) | **Post** /orgs/{org_id}/collections/{collection_id}/relationships/projects | Add projects to a collection



## CreateCollection

> CreateCollection201Response CreateCollection(ctx, orgId).Version(version).CreateCollectionRequest(createCollectionRequest).Execute()

Create a collection



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
	createCollectionRequest := *openapiclient.NewCreateCollectionRequest(*openapiclient.NewCreateCollectionRequestData(*openapiclient.NewCreateCollectionRequestDataAttributes("Name_example"), "resource")) // CreateCollectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAPI.CreateCollection(context.Background(), orgId).Version(version).CreateCollectionRequest(createCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.CreateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCollection`: CreateCollection201Response
	fmt.Fprintf(os.Stdout, "Response from `CollectionAPI.CreateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **createCollectionRequest** | [**CreateCollectionRequest**](CreateCollectionRequest.md) |  | 

### Return type

[**CreateCollection201Response**](CreateCollection201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> DeleteCollection(ctx, orgId, collectionId).Version(version).Execute()

Delete a collection



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for a collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAPI.DeleteCollection(context.Background(), orgId, collectionId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.DeleteCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**collectionId** | **string** | Unique identifier for a collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionRequest struct via the builder pattern


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


## DeleteProjectsCollection

> DeleteProjectsCollection(ctx, orgId, collectionId).Version(version).DeleteProjectsFromCollectionRequest(deleteProjectsFromCollectionRequest).Execute()

Remove projects from a collection



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for a collection
	deleteProjectsFromCollectionRequest := *openapiclient.NewDeleteProjectsFromCollectionRequest([]openapiclient.DeleteProjectsFromCollectionRequestDataInner{*openapiclient.NewDeleteProjectsFromCollectionRequestDataInner("Id_example", "Type_example")}) // DeleteProjectsFromCollectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAPI.DeleteProjectsCollection(context.Background(), orgId, collectionId).Version(version).DeleteProjectsFromCollectionRequest(deleteProjectsFromCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.DeleteProjectsCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**collectionId** | **string** | Unique identifier for a collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **deleteProjectsFromCollectionRequest** | [**DeleteProjectsFromCollectionRequest**](DeleteProjectsFromCollectionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollection

> CreateCollection201Response GetCollection(ctx, orgId, collectionId).Version(version).Execute()

Get a collection



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for a collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAPI.GetCollection(context.Background(), orgId, collectionId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.GetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollection`: CreateCollection201Response
	fmt.Fprintf(os.Stdout, "Response from `CollectionAPI.GetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**collectionId** | **string** | Unique identifier for a collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 



### Return type

[**CreateCollection201Response**](CreateCollection201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollections

> GetCollections200Response GetCollections(ctx, orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Sort(sort).Direction(direction).Name(name).IsGenerated(isGenerated).Execute()

Get collections



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
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)
	sort := "sort_example" // string | Return collections sorted by the specified attributes (optional)
	direction := "direction_example" // string | Return collections sorted in the specified direction (optional) (default to "DESC")
	name := "name_example" // string | Return collections which names include the provided string (optional)
	isGenerated := true // bool | Return collections where is_generated matches the provided boolean (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAPI.GetCollections(context.Background(), orgId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Sort(sort).Direction(direction).Name(name).IsGenerated(isGenerated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.GetCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollections`: GetCollections200Response
	fmt.Fprintf(os.Stdout, "Response from `CollectionAPI.GetCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]
 **sort** | **string** | Return collections sorted by the specified attributes | 
 **direction** | **string** | Return collections sorted in the specified direction | [default to &quot;DESC&quot;]
 **name** | **string** | Return collections which names include the provided string | 
 **isGenerated** | **bool** | Return collections where is_generated matches the provided boolean | 

### Return type

[**GetCollections200Response**](GetCollections200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsOfCollection

> GetProjectsOfCollectionResponse GetProjectsOfCollection(ctx, orgId, collectionId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Sort(sort).Direction(direction).TargetId(targetId).Show(show).Integration(integration).Execute()

Get projects from the specified collection



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for a collection
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)
	sort := "sort_example" // string | Return projects sorted by the specified attributes (optional)
	direction := "direction_example" // string | Return projects sorted in the specified direction (optional) (default to "DESC")
	targetId := []string{"Inner_example"} // []string | Return projects that belong to the provided targets (optional)
	show := []string{"Show_example"} // []string | Return projects that are with or without issues (optional)
	integration := []string{"Integration_example"} // []string | Return projects that match the provided integration types (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAPI.GetProjectsOfCollection(context.Background(), orgId, collectionId).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).Sort(sort).Direction(direction).TargetId(targetId).Show(show).Integration(integration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.GetProjectsOfCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsOfCollection`: GetProjectsOfCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionAPI.GetProjectsOfCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**collectionId** | **string** | Unique identifier for a collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsOfCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]
 **sort** | **string** | Return projects sorted by the specified attributes | 
 **direction** | **string** | Return projects sorted in the specified direction | [default to &quot;DESC&quot;]
 **targetId** | **[]string** | Return projects that belong to the provided targets | 
 **show** | **[]string** | Return projects that are with or without issues | 
 **integration** | **[]string** | Return projects that match the provided integration types | 

### Return type

[**GetProjectsOfCollectionResponse**](GetProjectsOfCollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> CreateCollection201Response UpdateCollection(ctx, orgId, collectionId).Version(version).UpdateCollectionRequest(updateCollectionRequest).Execute()

Edit a collection



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for a collection
	updateCollectionRequest := *openapiclient.NewUpdateCollectionRequest(*openapiclient.NewUpdateCollectionRequestData(*openapiclient.NewCreateCollectionRequestDataAttributes("Name_example"), "resource")) // UpdateCollectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAPI.UpdateCollection(context.Background(), orgId, collectionId).Version(version).UpdateCollectionRequest(updateCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.UpdateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCollection`: CreateCollection201Response
	fmt.Fprintf(os.Stdout, "Response from `CollectionAPI.UpdateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**collectionId** | **string** | Unique identifier for a collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **updateCollectionRequest** | [**UpdateCollectionRequest**](UpdateCollectionRequest.md) |  | 

### Return type

[**CreateCollection201Response**](CreateCollection201Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollectionWithProjects

> UpdateCollectionWithProjects(ctx, orgId, collectionId).Version(version).UpdateCollectionWithProjectsRequest(updateCollectionWithProjectsRequest).Execute()

Add projects to a collection



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for a collection
	updateCollectionWithProjectsRequest := *openapiclient.NewUpdateCollectionWithProjectsRequest([]openapiclient.DeleteProjectsFromCollectionRequestDataInner{*openapiclient.NewDeleteProjectsFromCollectionRequestDataInner("Id_example", "Type_example")}) // UpdateCollectionWithProjectsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAPI.UpdateCollectionWithProjects(context.Background(), orgId, collectionId).Version(version).UpdateCollectionWithProjectsRequest(updateCollectionWithProjectsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.UpdateCollectionWithProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**collectionId** | **string** | Unique identifier for a collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionWithProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


 **updateCollectionWithProjectsRequest** | [**UpdateCollectionWithProjectsRequest**](UpdateCollectionWithProjectsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

