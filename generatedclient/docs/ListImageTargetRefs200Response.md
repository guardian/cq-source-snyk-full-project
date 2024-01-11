# ListImageTargetRefs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ImageTargetRef**](ImageTargetRef.md) |  | [optional] 
**Jsonapi** | Pointer to [**JsonApi**](JsonApi.md) |  | [optional] 
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 

## Methods

### NewListImageTargetRefs200Response

`func NewListImageTargetRefs200Response() *ListImageTargetRefs200Response`

NewListImageTargetRefs200Response instantiates a new ListImageTargetRefs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageTargetRefs200ResponseWithDefaults

`func NewListImageTargetRefs200ResponseWithDefaults() *ListImageTargetRefs200Response`

NewListImageTargetRefs200ResponseWithDefaults instantiates a new ListImageTargetRefs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListImageTargetRefs200Response) GetData() []ImageTargetRef`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListImageTargetRefs200Response) GetDataOk() (*[]ImageTargetRef, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListImageTargetRefs200Response) SetData(v []ImageTargetRef)`

SetData sets Data field to given value.

### HasData

`func (o *ListImageTargetRefs200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetJsonapi

`func (o *ListImageTargetRefs200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *ListImageTargetRefs200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *ListImageTargetRefs200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.

### HasJsonapi

`func (o *ListImageTargetRefs200Response) HasJsonapi() bool`

HasJsonapi returns a boolean if a field has been set.

### GetLinks

`func (o *ListImageTargetRefs200Response) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListImageTargetRefs200Response) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListImageTargetRefs200Response) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListImageTargetRefs200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

