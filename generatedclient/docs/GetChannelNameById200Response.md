# GetChannelNameById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SlackChannel**](SlackChannel.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**SelfLink**](SelfLink.md) |  | 

## Methods

### NewGetChannelNameById200Response

`func NewGetChannelNameById200Response(data SlackChannel, jsonapi JsonApi, links SelfLink, ) *GetChannelNameById200Response`

NewGetChannelNameById200Response instantiates a new GetChannelNameById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChannelNameById200ResponseWithDefaults

`func NewGetChannelNameById200ResponseWithDefaults() *GetChannelNameById200Response`

NewGetChannelNameById200ResponseWithDefaults instantiates a new GetChannelNameById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetChannelNameById200Response) GetData() SlackChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetChannelNameById200Response) GetDataOk() (*SlackChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetChannelNameById200Response) SetData(v SlackChannel)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *GetChannelNameById200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetChannelNameById200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetChannelNameById200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *GetChannelNameById200Response) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetChannelNameById200Response) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetChannelNameById200Response) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


