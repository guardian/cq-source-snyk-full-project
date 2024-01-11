# OrgRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberRole** | Pointer to [**NullableMemberRoleRelationship**](MemberRoleRelationship.md) |  | [optional] 

## Methods

### NewOrgRelationships

`func NewOrgRelationships() *OrgRelationships`

NewOrgRelationships instantiates a new OrgRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgRelationshipsWithDefaults

`func NewOrgRelationshipsWithDefaults() *OrgRelationships`

NewOrgRelationshipsWithDefaults instantiates a new OrgRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberRole

`func (o *OrgRelationships) GetMemberRole() MemberRoleRelationship`

GetMemberRole returns the MemberRole field if non-nil, zero value otherwise.

### GetMemberRoleOk

`func (o *OrgRelationships) GetMemberRoleOk() (*MemberRoleRelationship, bool)`

GetMemberRoleOk returns a tuple with the MemberRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRole

`func (o *OrgRelationships) SetMemberRole(v MemberRoleRelationship)`

SetMemberRole sets MemberRole field to given value.

### HasMemberRole

`func (o *OrgRelationships) HasMemberRole() bool`

HasMemberRole returns a boolean if a field has been set.

### SetMemberRoleNil

`func (o *OrgRelationships) SetMemberRoleNil(b bool)`

 SetMemberRoleNil sets the value for MemberRole to be an explicit nil

### UnsetMemberRole
`func (o *OrgRelationships) UnsetMemberRole()`

UnsetMemberRole ensures that no value is present for MemberRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


