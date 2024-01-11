# GetOrg200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Org**](Org.md) |  | [optional] 
**Jsonapi** | Pointer to [**JsonApi**](JsonApi.md) |  | [optional] 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 

## Methods

### NewGetOrg200Response

`func NewGetOrg200Response() *GetOrg200Response`

NewGetOrg200Response instantiates a new GetOrg200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrg200ResponseWithDefaults

`func NewGetOrg200ResponseWithDefaults() *GetOrg200Response`

NewGetOrg200ResponseWithDefaults instantiates a new GetOrg200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetOrg200Response) GetData() Org`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrg200Response) GetDataOk() (*Org, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrg200Response) SetData(v Org)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrg200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetJsonapi

`func (o *GetOrg200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetOrg200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetOrg200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.

### HasJsonapi

`func (o *GetOrg200Response) HasJsonapi() bool`

HasJsonapi returns a boolean if a field has been set.

### GetLinks

`func (o *GetOrg200Response) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetOrg200Response) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetOrg200Response) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetOrg200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


