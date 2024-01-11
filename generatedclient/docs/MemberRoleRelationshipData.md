# MemberRoleRelationshipData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**OrgRoleAttributes**](OrgRoleAttributes.md) |  | [optional] 
**Id** | **string** | The Snyk ID of the organization role. | 
**Type** | **string** |  | 

## Methods

### NewMemberRoleRelationshipData

`func NewMemberRoleRelationshipData(id string, type_ string, ) *MemberRoleRelationshipData`

NewMemberRoleRelationshipData instantiates a new MemberRoleRelationshipData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberRoleRelationshipDataWithDefaults

`func NewMemberRoleRelationshipDataWithDefaults() *MemberRoleRelationshipData`

NewMemberRoleRelationshipDataWithDefaults instantiates a new MemberRoleRelationshipData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MemberRoleRelationshipData) GetAttributes() OrgRoleAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MemberRoleRelationshipData) GetAttributesOk() (*OrgRoleAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MemberRoleRelationshipData) SetAttributes(v OrgRoleAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MemberRoleRelationshipData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MemberRoleRelationshipData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberRoleRelationshipData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberRoleRelationshipData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MemberRoleRelationshipData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemberRoleRelationshipData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemberRoleRelationshipData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


