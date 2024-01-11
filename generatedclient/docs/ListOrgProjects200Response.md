# ListOrgProjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ListOrgProjects200ResponseDataInner**](ListOrgProjects200ResponseDataInner.md) |  | [optional] 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | Pointer to [**ListOrgProjects200ResponseMeta**](ListOrgProjects200ResponseMeta.md) |  | [optional] 

## Methods

### NewListOrgProjects200Response

`func NewListOrgProjects200Response(jsonapi JsonApi, links Links, ) *ListOrgProjects200Response`

NewListOrgProjects200Response instantiates a new ListOrgProjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrgProjects200ResponseWithDefaults

`func NewListOrgProjects200ResponseWithDefaults() *ListOrgProjects200Response`

NewListOrgProjects200ResponseWithDefaults instantiates a new ListOrgProjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListOrgProjects200Response) GetData() []ListOrgProjects200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListOrgProjects200Response) GetDataOk() (*[]ListOrgProjects200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListOrgProjects200Response) SetData(v []ListOrgProjects200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ListOrgProjects200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetJsonapi

`func (o *ListOrgProjects200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *ListOrgProjects200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *ListOrgProjects200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *ListOrgProjects200Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListOrgProjects200Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListOrgProjects200Response) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *ListOrgProjects200Response) GetMeta() ListOrgProjects200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListOrgProjects200Response) GetMetaOk() (*ListOrgProjects200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListOrgProjects200Response) SetMeta(v ListOrgProjects200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListOrgProjects200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


