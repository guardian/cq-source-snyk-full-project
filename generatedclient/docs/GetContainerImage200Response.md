# GetContainerImage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Image**](Image.md) |  | [optional] 
**Jsonapi** | Pointer to [**JsonApi**](JsonApi.md) |  | [optional] 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 

## Methods

### NewGetContainerImage200Response

`func NewGetContainerImage200Response() *GetContainerImage200Response`

NewGetContainerImage200Response instantiates a new GetContainerImage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerImage200ResponseWithDefaults

`func NewGetContainerImage200ResponseWithDefaults() *GetContainerImage200Response`

NewGetContainerImage200ResponseWithDefaults instantiates a new GetContainerImage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetContainerImage200Response) GetData() Image`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContainerImage200Response) GetDataOk() (*Image, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContainerImage200Response) SetData(v Image)`

SetData sets Data field to given value.

### HasData

`func (o *GetContainerImage200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetJsonapi

`func (o *GetContainerImage200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetContainerImage200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetContainerImage200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.

### HasJsonapi

`func (o *GetContainerImage200Response) HasJsonapi() bool`

HasJsonapi returns a boolean if a field has been set.

### GetLinks

`func (o *GetContainerImage200Response) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetContainerImage200Response) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetContainerImage200Response) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetContainerImage200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


