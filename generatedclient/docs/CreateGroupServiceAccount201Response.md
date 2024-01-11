# CreateGroupServiceAccount201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetManyGroupServiceAccount200ResponseDataInner**](GetManyGroupServiceAccount200ResponseDataInner.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewCreateGroupServiceAccount201Response

`func NewCreateGroupServiceAccount201Response(data GetManyGroupServiceAccount200ResponseDataInner, jsonapi JsonApi, ) *CreateGroupServiceAccount201Response`

NewCreateGroupServiceAccount201Response instantiates a new CreateGroupServiceAccount201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupServiceAccount201ResponseWithDefaults

`func NewCreateGroupServiceAccount201ResponseWithDefaults() *CreateGroupServiceAccount201Response`

NewCreateGroupServiceAccount201ResponseWithDefaults instantiates a new CreateGroupServiceAccount201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateGroupServiceAccount201Response) GetData() GetManyGroupServiceAccount200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateGroupServiceAccount201Response) GetDataOk() (*GetManyGroupServiceAccount200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateGroupServiceAccount201Response) SetData(v GetManyGroupServiceAccount200ResponseDataInner)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *CreateGroupServiceAccount201Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *CreateGroupServiceAccount201Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *CreateGroupServiceAccount201Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *CreateGroupServiceAccount201Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateGroupServiceAccount201Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateGroupServiceAccount201Response) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateGroupServiceAccount201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


