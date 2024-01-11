# SbomResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SbomResource**](SbomResource.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 

## Methods

### NewSbomResponse

`func NewSbomResponse(data SbomResource, jsonapi JsonApi, ) *SbomResponse`

NewSbomResponse instantiates a new SbomResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbomResponseWithDefaults

`func NewSbomResponseWithDefaults() *SbomResponse`

NewSbomResponseWithDefaults instantiates a new SbomResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SbomResponse) GetData() SbomResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SbomResponse) GetDataOk() (*SbomResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SbomResponse) SetData(v SbomResource)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *SbomResponse) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *SbomResponse) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *SbomResponse) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


