# AppPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AppDataWithSecret**](AppDataWithSecret.md) |  | [optional] 
**Jsonapi** | Pointer to [**JsonApi**](JsonApi.md) |  | [optional] 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 

## Methods

### NewAppPostResponse

`func NewAppPostResponse() *AppPostResponse`

NewAppPostResponse instantiates a new AppPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPostResponseWithDefaults

`func NewAppPostResponseWithDefaults() *AppPostResponse`

NewAppPostResponseWithDefaults instantiates a new AppPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppPostResponse) GetData() AppDataWithSecret`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppPostResponse) GetDataOk() (*AppDataWithSecret, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppPostResponse) SetData(v AppDataWithSecret)`

SetData sets Data field to given value.

### HasData

`func (o *AppPostResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetJsonapi

`func (o *AppPostResponse) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *AppPostResponse) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *AppPostResponse) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.

### HasJsonapi

`func (o *AppPostResponse) HasJsonapi() bool`

HasJsonapi returns a boolean if a field has been set.

### GetLinks

`func (o *AppPostResponse) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPostResponse) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPostResponse) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppPostResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


