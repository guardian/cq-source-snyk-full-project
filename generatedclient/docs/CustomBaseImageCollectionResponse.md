# CustomBaseImageCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CustomBaseImageCollectionResponseDataInner**](CustomBaseImageCollectionResponseDataInner.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 

## Methods

### NewCustomBaseImageCollectionResponse

`func NewCustomBaseImageCollectionResponse(data []CustomBaseImageCollectionResponseDataInner, jsonapi JsonApi, ) *CustomBaseImageCollectionResponse`

NewCustomBaseImageCollectionResponse instantiates a new CustomBaseImageCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomBaseImageCollectionResponseWithDefaults

`func NewCustomBaseImageCollectionResponseWithDefaults() *CustomBaseImageCollectionResponse`

NewCustomBaseImageCollectionResponseWithDefaults instantiates a new CustomBaseImageCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CustomBaseImageCollectionResponse) GetData() []CustomBaseImageCollectionResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomBaseImageCollectionResponse) GetDataOk() (*[]CustomBaseImageCollectionResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomBaseImageCollectionResponse) SetData(v []CustomBaseImageCollectionResponseDataInner)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *CustomBaseImageCollectionResponse) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *CustomBaseImageCollectionResponse) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *CustomBaseImageCollectionResponse) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *CustomBaseImageCollectionResponse) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomBaseImageCollectionResponse) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomBaseImageCollectionResponse) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomBaseImageCollectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


