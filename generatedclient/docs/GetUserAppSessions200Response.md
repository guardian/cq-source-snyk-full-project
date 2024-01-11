# GetUserAppSessions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SessionData**](SessionData.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 

## Methods

### NewGetUserAppSessions200Response

`func NewGetUserAppSessions200Response(data []SessionData, jsonapi JsonApi, ) *GetUserAppSessions200Response`

NewGetUserAppSessions200Response instantiates a new GetUserAppSessions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserAppSessions200ResponseWithDefaults

`func NewGetUserAppSessions200ResponseWithDefaults() *GetUserAppSessions200Response`

NewGetUserAppSessions200ResponseWithDefaults instantiates a new GetUserAppSessions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUserAppSessions200Response) GetData() []SessionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserAppSessions200Response) GetDataOk() (*[]SessionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserAppSessions200Response) SetData(v []SessionData)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *GetUserAppSessions200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetUserAppSessions200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetUserAppSessions200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *GetUserAppSessions200Response) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUserAppSessions200Response) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUserAppSessions200Response) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUserAppSessions200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


