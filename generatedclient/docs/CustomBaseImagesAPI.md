# \CustomBaseImagesAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomBaseImage**](CustomBaseImagesAPI.md#CreateCustomBaseImage) | **Post** /custom_base_images | Create a Custom Base Image from an existing container project
[**DeleteCustomBaseImage**](CustomBaseImagesAPI.md#DeleteCustomBaseImage) | **Delete** /custom_base_images/{custombaseimage_id} | Delete a custom base image
[**GetCustomBaseImage**](CustomBaseImagesAPI.md#GetCustomBaseImage) | **Get** /custom_base_images/{custombaseimage_id} | Get a custom base image
[**GetCustomBaseImages**](CustomBaseImagesAPI.md#GetCustomBaseImages) | **Get** /custom_base_images | Get a custom base image collection
[**UpdateCustomBaseImage**](CustomBaseImagesAPI.md#UpdateCustomBaseImage) | **Patch** /custom_base_images/{custombaseimage_id} | Update a custom base image



## CreateCustomBaseImage

> CustomBaseImageResponse CreateCustomBaseImage(ctx).Version(version).CustomBaseImagePostRequest(customBaseImagePostRequest).Execute()

Create a Custom Base Image from an existing container project



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
	customBaseImagePostRequest := *openapiclient.NewCustomBaseImagePostRequest(*openapiclient.NewCustomBaseImagePostRequestData(*openapiclient.NewCustomBaseImageAttributes(true, "2cab3939-d112-4ef0-836d-e09c87cbe69b"), "custom_base_image")) // CustomBaseImagePostRequest | The properties used in the creation of a custom base image (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomBaseImagesAPI.CreateCustomBaseImage(context.Background()).Version(version).CustomBaseImagePostRequest(customBaseImagePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomBaseImagesAPI.CreateCustomBaseImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomBaseImage`: CustomBaseImageResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomBaseImagesAPI.CreateCustomBaseImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomBaseImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 
 **customBaseImagePostRequest** | [**CustomBaseImagePostRequest**](CustomBaseImagePostRequest.md) | The properties used in the creation of a custom base image | 

### Return type

[**CustomBaseImageResponse**](CustomBaseImageResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomBaseImage

> DeleteCustomBaseImage(ctx, custombaseimageId).Version(version).Execute()

Delete a custom base image



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
	custombaseimageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for custom base image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomBaseImagesAPI.DeleteCustomBaseImage(context.Background(), custombaseimageId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomBaseImagesAPI.DeleteCustomBaseImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**custombaseimageId** | **string** | Unique identifier for custom base image | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomBaseImageRequest struct via the builder pattern


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


## GetCustomBaseImage

> CustomBaseImageResponse GetCustomBaseImage(ctx, custombaseimageId).Version(version).Execute()

Get a custom base image



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
	custombaseimageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for custom base image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomBaseImagesAPI.GetCustomBaseImage(context.Background(), custombaseimageId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomBaseImagesAPI.GetCustomBaseImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomBaseImage`: CustomBaseImageResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomBaseImagesAPI.GetCustomBaseImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**custombaseimageId** | **string** | Unique identifier for custom base image | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomBaseImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 


### Return type

[**CustomBaseImageResponse**](CustomBaseImageResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomBaseImages

> CustomBaseImageCollectionResponse GetCustomBaseImages(ctx).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).ProjectId(projectId).OrgId(orgId).GroupId(groupId).Repository(repository).Tag(tag).IncludeInRecommendations(includeInRecommendations).SortBy(sortBy).SortDirection(sortDirection).Execute()

Get a custom base image collection



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the container project that the custom base image is based off of. (optional)
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The organization ID of the custom base image (optional)
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The group ID of the custom base image (optional)
	repository := "repository_example" // string | The image repository (optional)
	tag := "tag_example" // string | The image tag (optional)
	includeInRecommendations := true // bool | Whether this image should be recommended as a base image upgrade (optional)
	sortBy := "sortBy_example" // string | Which column to sort by.  If sorting by version, the versioning schema is used.  (optional)
	sortDirection := "sortDirection_example" // string | Which direction to sort (optional) (default to "ASC")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomBaseImagesAPI.GetCustomBaseImages(context.Background()).Version(version).StartingAfter(startingAfter).EndingBefore(endingBefore).Limit(limit).ProjectId(projectId).OrgId(orgId).GroupId(groupId).Repository(repository).Tag(tag).IncludeInRecommendations(includeInRecommendations).SortBy(sortBy).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomBaseImagesAPI.GetCustomBaseImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomBaseImages`: CustomBaseImageCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomBaseImagesAPI.GetCustomBaseImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomBaseImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 
 **limit** | **int32** | Number of results to return per page | [default to 10]
 **projectId** | **string** | The ID of the container project that the custom base image is based off of. | 
 **orgId** | **string** | The organization ID of the custom base image | 
 **groupId** | **string** | The group ID of the custom base image | 
 **repository** | **string** | The image repository | 
 **tag** | **string** | The image tag | 
 **includeInRecommendations** | **bool** | Whether this image should be recommended as a base image upgrade | 
 **sortBy** | **string** | Which column to sort by.  If sorting by version, the versioning schema is used.  | 
 **sortDirection** | **string** | Which direction to sort | [default to &quot;ASC&quot;]

### Return type

[**CustomBaseImageCollectionResponse**](CustomBaseImageCollectionResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomBaseImage

> CustomBaseImageResponse UpdateCustomBaseImage(ctx, custombaseimageId).Version(version).CustomBaseImagePatchRequest(customBaseImagePatchRequest).Execute()

Update a custom base image



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
	custombaseimageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier for custom base image
	customBaseImagePatchRequest := *openapiclient.NewCustomBaseImagePatchRequest(*openapiclient.NewCustomBaseImagePatchRequestData(*openapiclient.NewCustomBaseImagePatchRequestDataAttributes(), "custom_base_image")) // CustomBaseImagePatchRequest | custom base image to be updated (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomBaseImagesAPI.UpdateCustomBaseImage(context.Background(), custombaseimageId).Version(version).CustomBaseImagePatchRequest(customBaseImagePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomBaseImagesAPI.UpdateCustomBaseImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomBaseImage`: CustomBaseImageResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomBaseImagesAPI.UpdateCustomBaseImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**custombaseimageId** | **string** | Unique identifier for custom base image | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomBaseImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 

 **customBaseImagePatchRequest** | [**CustomBaseImagePatchRequest**](CustomBaseImagePatchRequest.md) | custom base image to be updated | 

### Return type

[**CustomBaseImageResponse**](CustomBaseImageResponse.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

