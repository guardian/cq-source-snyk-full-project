# CreateOrgInvitation201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**OrgInvitation**](OrgInvitation.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewCreateOrgInvitation201Response

`func NewCreateOrgInvitation201Response(data OrgInvitation, jsonapi JsonApi, ) *CreateOrgInvitation201Response`

NewCreateOrgInvitation201Response instantiates a new CreateOrgInvitation201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrgInvitation201ResponseWithDefaults

`func NewCreateOrgInvitation201ResponseWithDefaults() *CreateOrgInvitation201Response`

NewCreateOrgInvitation201ResponseWithDefaults instantiates a new CreateOrgInvitation201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateOrgInvitation201Response) GetData() OrgInvitation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateOrgInvitation201Response) GetDataOk() (*OrgInvitation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateOrgInvitation201Response) SetData(v OrgInvitation)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *CreateOrgInvitation201Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *CreateOrgInvitation201Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *CreateOrgInvitation201Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *CreateOrgInvitation201Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateOrgInvitation201Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateOrgInvitation201Response) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateOrgInvitation201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


