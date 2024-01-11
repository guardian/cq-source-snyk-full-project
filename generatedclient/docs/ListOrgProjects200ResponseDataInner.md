# ListOrgProjects200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**ProjectAttributes**](ProjectAttributes.md) |  | 
**Id** | **string** | Resource ID. | 
**Meta** | Pointer to [**ListOrgProjects200ResponseDataInnerMeta**](ListOrgProjects200ResponseDataInnerMeta.md) |  | [optional] 
**Relationships** | Pointer to [**ProjectRelationships**](ProjectRelationships.md) |  | [optional] 
**Type** | **string** | The Resource type. | 

## Methods

### NewListOrgProjects200ResponseDataInner

`func NewListOrgProjects200ResponseDataInner(attributes ProjectAttributes, id string, type_ string, ) *ListOrgProjects200ResponseDataInner`

NewListOrgProjects200ResponseDataInner instantiates a new ListOrgProjects200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrgProjects200ResponseDataInnerWithDefaults

`func NewListOrgProjects200ResponseDataInnerWithDefaults() *ListOrgProjects200ResponseDataInner`

NewListOrgProjects200ResponseDataInnerWithDefaults instantiates a new ListOrgProjects200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ListOrgProjects200ResponseDataInner) GetAttributes() ProjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ListOrgProjects200ResponseDataInner) GetAttributesOk() (*ProjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ListOrgProjects200ResponseDataInner) SetAttributes(v ProjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *ListOrgProjects200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOrgProjects200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOrgProjects200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMeta

`func (o *ListOrgProjects200ResponseDataInner) GetMeta() ListOrgProjects200ResponseDataInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListOrgProjects200ResponseDataInner) GetMetaOk() (*ListOrgProjects200ResponseDataInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListOrgProjects200ResponseDataInner) SetMeta(v ListOrgProjects200ResponseDataInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListOrgProjects200ResponseDataInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRelationships

`func (o *ListOrgProjects200ResponseDataInner) GetRelationships() ProjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ListOrgProjects200ResponseDataInner) GetRelationshipsOk() (*ProjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ListOrgProjects200ResponseDataInner) SetRelationships(v ProjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ListOrgProjects200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *ListOrgProjects200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListOrgProjects200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListOrgProjects200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


