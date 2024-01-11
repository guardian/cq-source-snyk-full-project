# GetManyGroupServiceAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetManyGroupServiceAccount200ResponseDataInner**](GetManyGroupServiceAccount200ResponseDataInner.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewGetManyGroupServiceAccount200Response

`func NewGetManyGroupServiceAccount200Response(data []GetManyGroupServiceAccount200ResponseDataInner, jsonapi JsonApi, links Links, ) *GetManyGroupServiceAccount200Response`

NewGetManyGroupServiceAccount200Response instantiates a new GetManyGroupServiceAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManyGroupServiceAccount200ResponseWithDefaults

`func NewGetManyGroupServiceAccount200ResponseWithDefaults() *GetManyGroupServiceAccount200Response`

NewGetManyGroupServiceAccount200ResponseWithDefaults instantiates a new GetManyGroupServiceAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetManyGroupServiceAccount200Response) GetData() []GetManyGroupServiceAccount200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetManyGroupServiceAccount200Response) GetDataOk() (*[]GetManyGroupServiceAccount200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetManyGroupServiceAccount200Response) SetData(v []GetManyGroupServiceAccount200ResponseDataInner)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *GetManyGroupServiceAccount200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetManyGroupServiceAccount200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetManyGroupServiceAccount200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *GetManyGroupServiceAccount200Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetManyGroupServiceAccount200Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetManyGroupServiceAccount200Response) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


