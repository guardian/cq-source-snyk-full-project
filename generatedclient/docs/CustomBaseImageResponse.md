# CustomBaseImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CustomBaseImageResource**](CustomBaseImageResource.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 

## Methods

### NewCustomBaseImageResponse

`func NewCustomBaseImageResponse(data CustomBaseImageResource, jsonapi JsonApi, ) *CustomBaseImageResponse`

NewCustomBaseImageResponse instantiates a new CustomBaseImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomBaseImageResponseWithDefaults

`func NewCustomBaseImageResponseWithDefaults() *CustomBaseImageResponse`

NewCustomBaseImageResponseWithDefaults instantiates a new CustomBaseImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CustomBaseImageResponse) GetData() CustomBaseImageResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomBaseImageResponse) GetDataOk() (*CustomBaseImageResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomBaseImageResponse) SetData(v CustomBaseImageResource)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *CustomBaseImageResponse) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *CustomBaseImageResponse) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *CustomBaseImageResponse) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *CustomBaseImageResponse) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomBaseImageResponse) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomBaseImageResponse) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomBaseImageResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


