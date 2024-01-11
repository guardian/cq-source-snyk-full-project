# UpdateOrgProject200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**ProjectAttributes**](ProjectAttributes.md) |  | 
**Id** | **string** | The Resource ID. | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Meta** | Pointer to [**UpdateOrgProject200ResponseDataMeta**](UpdateOrgProject200ResponseDataMeta.md) |  | [optional] 
**Relationships** | Pointer to [**ProjectRelationships**](ProjectRelationships.md) |  | [optional] 
**Type** | **string** | The Resource type. | 

## Methods

### NewUpdateOrgProject200ResponseData

`func NewUpdateOrgProject200ResponseData(attributes ProjectAttributes, id string, type_ string, ) *UpdateOrgProject200ResponseData`

NewUpdateOrgProject200ResponseData instantiates a new UpdateOrgProject200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgProject200ResponseDataWithDefaults

`func NewUpdateOrgProject200ResponseDataWithDefaults() *UpdateOrgProject200ResponseData`

NewUpdateOrgProject200ResponseDataWithDefaults instantiates a new UpdateOrgProject200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateOrgProject200ResponseData) GetAttributes() ProjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateOrgProject200ResponseData) GetAttributesOk() (*ProjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateOrgProject200ResponseData) SetAttributes(v ProjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *UpdateOrgProject200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrgProject200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrgProject200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *UpdateOrgProject200ResponseData) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateOrgProject200ResponseData) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateOrgProject200ResponseData) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateOrgProject200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *UpdateOrgProject200ResponseData) GetMeta() UpdateOrgProject200ResponseDataMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UpdateOrgProject200ResponseData) GetMetaOk() (*UpdateOrgProject200ResponseDataMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UpdateOrgProject200ResponseData) SetMeta(v UpdateOrgProject200ResponseDataMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UpdateOrgProject200ResponseData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRelationships

`func (o *UpdateOrgProject200ResponseData) GetRelationships() ProjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UpdateOrgProject200ResponseData) GetRelationshipsOk() (*ProjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UpdateOrgProject200ResponseData) SetRelationships(v ProjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UpdateOrgProject200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *UpdateOrgProject200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOrgProject200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOrgProject200ResponseData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


