# ListChannels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SlackChannel**](SlackChannel.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**PaginatedLinks**](PaginatedLinks.md) |  | 

## Methods

### NewListChannels200Response

`func NewListChannels200Response(data []SlackChannel, jsonapi JsonApi, links PaginatedLinks, ) *ListChannels200Response`

NewListChannels200Response instantiates a new ListChannels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChannels200ResponseWithDefaults

`func NewListChannels200ResponseWithDefaults() *ListChannels200Response`

NewListChannels200ResponseWithDefaults instantiates a new ListChannels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListChannels200Response) GetData() []SlackChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListChannels200Response) GetDataOk() (*[]SlackChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListChannels200Response) SetData(v []SlackChannel)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *ListChannels200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *ListChannels200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *ListChannels200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *ListChannels200Response) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListChannels200Response) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListChannels200Response) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


