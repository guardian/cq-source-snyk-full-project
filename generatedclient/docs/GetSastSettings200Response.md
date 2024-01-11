# GetSastSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SastEnablement**](SastEnablement.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewGetSastSettings200Response

`func NewGetSastSettings200Response(data SastEnablement, jsonapi JsonApi, links Links, ) *GetSastSettings200Response`

NewGetSastSettings200Response instantiates a new GetSastSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSastSettings200ResponseWithDefaults

`func NewGetSastSettings200ResponseWithDefaults() *GetSastSettings200Response`

NewGetSastSettings200ResponseWithDefaults instantiates a new GetSastSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSastSettings200Response) GetData() SastEnablement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSastSettings200Response) GetDataOk() (*SastEnablement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSastSettings200Response) SetData(v SastEnablement)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *GetSastSettings200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetSastSettings200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetSastSettings200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *GetSastSettings200Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSastSettings200Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSastSettings200Response) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


