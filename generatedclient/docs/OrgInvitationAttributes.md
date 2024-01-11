# OrgInvitationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the invitee. | 
**IsActive** | **bool** | The active status of the invitation. is_active of true indicates that the invitation is still waiting to be accepted. Invitations are considered inactive when accepted or revoked.  | 
**Role** | **string** | The role public ID that will be granted to to invitee on acceptance. | 

## Methods

### NewOrgInvitationAttributes

`func NewOrgInvitationAttributes(email string, isActive bool, role string, ) *OrgInvitationAttributes`

NewOrgInvitationAttributes instantiates a new OrgInvitationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgInvitationAttributesWithDefaults

`func NewOrgInvitationAttributesWithDefaults() *OrgInvitationAttributes`

NewOrgInvitationAttributesWithDefaults instantiates a new OrgInvitationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrgInvitationAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgInvitationAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgInvitationAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsActive

`func (o *OrgInvitationAttributes) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *OrgInvitationAttributes) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *OrgInvitationAttributes) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetRole

`func (o *OrgInvitationAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgInvitationAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgInvitationAttributes) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


