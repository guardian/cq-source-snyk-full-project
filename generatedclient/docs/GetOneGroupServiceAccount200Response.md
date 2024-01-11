# GetOneGroupServiceAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetManyGroupServiceAccount200ResponseDataInner**](GetManyGroupServiceAccount200ResponseDataInner.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewGetOneGroupServiceAccount200Response

`func NewGetOneGroupServiceAccount200Response(data GetManyGroupServiceAccount200ResponseDataInner, jsonapi JsonApi, links Links, ) *GetOneGroupServiceAccount200Response`

NewGetOneGroupServiceAccount200Response instantiates a new GetOneGroupServiceAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOneGroupServiceAccount200ResponseWithDefaults

`func NewGetOneGroupServiceAccount200ResponseWithDefaults() *GetOneGroupServiceAccount200Response`

NewGetOneGroupServiceAccount200ResponseWithDefaults instantiates a new GetOneGroupServiceAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetOneGroupServiceAccount200Response) GetData() GetManyGroupServiceAccount200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOneGroupServiceAccount200Response) GetDataOk() (*GetManyGroupServiceAccount200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOneGroupServiceAccount200Response) SetData(v GetManyGroupServiceAccount200ResponseDataInner)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *GetOneGroupServiceAccount200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetOneGroupServiceAccount200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetOneGroupServiceAccount200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *GetOneGroupServiceAccount200Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetOneGroupServiceAccount200Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetOneGroupServiceAccount200Response) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


