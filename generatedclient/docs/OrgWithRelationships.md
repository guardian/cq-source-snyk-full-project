# OrgWithRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**OrgAttributes**](OrgAttributes.md) |  | 
**Id** | **string** | The Snyk ID of the organization. | 
**Relationships** | Pointer to [**OrgRelationships**](OrgRelationships.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewOrgWithRelationships

`func NewOrgWithRelationships(attributes OrgAttributes, id string, type_ string, ) *OrgWithRelationships`

NewOrgWithRelationships instantiates a new OrgWithRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWithRelationshipsWithDefaults

`func NewOrgWithRelationshipsWithDefaults() *OrgWithRelationships`

NewOrgWithRelationshipsWithDefaults instantiates a new OrgWithRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OrgWithRelationships) GetAttributes() OrgAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrgWithRelationships) GetAttributesOk() (*OrgAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrgWithRelationships) SetAttributes(v OrgAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *OrgWithRelationships) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgWithRelationships) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgWithRelationships) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *OrgWithRelationships) GetRelationships() OrgRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrgWithRelationships) GetRelationshipsOk() (*OrgRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrgWithRelationships) SetRelationships(v OrgRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrgWithRelationships) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *OrgWithRelationships) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgWithRelationships) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgWithRelationships) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


