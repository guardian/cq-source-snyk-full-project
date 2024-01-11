# GetOrgProject200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**ProjectAttributes**](ProjectAttributes.md) |  | 
**Id** | **string** | The Resource ID. | 
**Meta** | Pointer to [**ListOrgProjects200ResponseDataInnerMeta**](ListOrgProjects200ResponseDataInnerMeta.md) |  | [optional] 
**Relationships** | Pointer to [**ProjectRelationships**](ProjectRelationships.md) |  | [optional] 
**Type** | **string** | The Resource type. | 

## Methods

### NewGetOrgProject200ResponseData

`func NewGetOrgProject200ResponseData(attributes ProjectAttributes, id string, type_ string, ) *GetOrgProject200ResponseData`

NewGetOrgProject200ResponseData instantiates a new GetOrgProject200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrgProject200ResponseDataWithDefaults

`func NewGetOrgProject200ResponseDataWithDefaults() *GetOrgProject200ResponseData`

NewGetOrgProject200ResponseDataWithDefaults instantiates a new GetOrgProject200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *GetOrgProject200ResponseData) GetAttributes() ProjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetOrgProject200ResponseData) GetAttributesOk() (*ProjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetOrgProject200ResponseData) SetAttributes(v ProjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *GetOrgProject200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrgProject200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrgProject200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMeta

`func (o *GetOrgProject200ResponseData) GetMeta() ListOrgProjects200ResponseDataInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetOrgProject200ResponseData) GetMetaOk() (*ListOrgProjects200ResponseDataInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetOrgProject200ResponseData) SetMeta(v ListOrgProjects200ResponseDataInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetOrgProject200ResponseData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRelationships

`func (o *GetOrgProject200ResponseData) GetRelationships() ProjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetOrgProject200ResponseData) GetRelationshipsOk() (*ProjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetOrgProject200ResponseData) SetRelationships(v ProjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetOrgProject200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *GetOrgProject200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrgProject200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrgProject200ResponseData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


