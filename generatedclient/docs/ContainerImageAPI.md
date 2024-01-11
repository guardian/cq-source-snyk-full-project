# \ContainerImageAPI

All URIs are relative to *https://api.snyk.io/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainerImage**](ContainerImageAPI.md#GetContainerImage) | **Get** /orgs/{org_id}/container_images/{image_id} | Get instance of container image
[**ListContainerImage**](ContainerImageAPI.md#ListContainerImage) | **Get** /orgs/{org_id}/container_images | List instances of container image
[**ListImageTargetRefs**](ContainerImageAPI.md#ListImageTargetRefs) | **Get** /orgs/{org_id}/container_images/{image_id}/relationships/image_target_refs | List instances of image target references for a container image



## GetContainerImage

> GetContainerImage200Response GetContainerImage(ctx, orgId, imageId).Version(version).Execute()

Get instance of container image



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
	orgId := "f59045b3-f093-40c3-871d-a334ae30c568" // string | Org ID
	imageId := "sha256:2bd864580926b790a22c8b96fd74496fe87b3c59c0774fe144bab2788e78e676" // string | Image ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerImageAPI.GetContainerImage(context.Background(), orgId, imageId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImageAPI.GetContainerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerImage`: GetContainerImage200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainerImageAPI.GetContainerImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | The requested version of the endpoint to process the request | 



### Return type

[**GetContainerImage200Response**](GetContainerImage200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerImage

> ListContainerImage200Response ListContainerImage(ctx, orgId).Version(version).ImageIds(imageIds).Platform(platform).Names(names).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List instances of container image



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
	orgId := "f59045b3-f093-40c3-871d-a334ae30c568" // string | Org ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	imageIds := []string{"Inner_example"} // []string | A comma-separated list of Image IDs (optional)
	platform := openapiclient.Platform("aix/ppc64") // Platform | The image Operating System and processor architecture (optional)
	names := []string{"gcr.io/snyk/redis:5"} // []string | The container registry names (optional)
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerImageAPI.ListContainerImage(context.Background(), orgId).Version(version).ImageIds(imageIds).Platform(platform).Names(names).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImageAPI.ListContainerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerImage`: ListContainerImage200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainerImageAPI.ListContainerImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | The requested version of the endpoint to process the request | 
 **imageIds** | **[]string** | A comma-separated list of Image IDs | 
 **platform** | [**Platform**](Platform.md) | The image Operating System and processor architecture | 
 **names** | **[]string** | The container registry names | 
 **limit** | **int32** | Number of results to return per page | [default to 10]
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 

### Return type

[**ListContainerImage200Response**](ListContainerImage200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageTargetRefs

> ListImageTargetRefs200Response ListImageTargetRefs(ctx, orgId, imageId).Version(version).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List instances of image target references for a container image



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
	orgId := "f59045b3-f093-40c3-871d-a334ae30c568" // string | Org ID
	imageId := "sha256:2bd864580926b790a22c8b96fd74496fe87b3c59c0774fe144bab2788e78e676" // string | Image ID
	version := "2021-06-04" // string | The requested version of the endpoint to process the request
	limit := int32(10) // int32 | Number of results to return per page (optional) (default to 10)
	startingAfter := "v1.eyJpZCI6IjEwMDAifQo=" // string | Return the page of results immediately after this cursor (optional)
	endingBefore := "v1.eyJpZCI6IjExMDAifQo=" // string | Return the page of results immediately before this cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerImageAPI.ListImageTargetRefs(context.Background(), orgId, imageId).Version(version).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImageAPI.ListImageTargetRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImageTargetRefs`: ListImageTargetRefs200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainerImageAPI.ListImageTargetRefs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageTargetRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | The requested version of the endpoint to process the request | 
 **limit** | **int32** | Number of results to return per page | [default to 10]
 **startingAfter** | **string** | Return the page of results immediately after this cursor | 
 **endingBefore** | **string** | Return the page of results immediately before this cursor | 

### Return type

[**ListImageTargetRefs200Response**](ListImageTargetRefs200Response.md)

### Authorization

[APIToken](../README.md#APIToken), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

