# OrgInvitationPostAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the invitee. | 
**Role** | **interface{}** | The role public ID that will be granted to to invitee on acceptance. | 

## Methods

### NewOrgInvitationPostAttributes

`func NewOrgInvitationPostAttributes(email string, role interface{}, ) *OrgInvitationPostAttributes`

NewOrgInvitationPostAttributes instantiates a new OrgInvitationPostAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgInvitationPostAttributesWithDefaults

`func NewOrgInvitationPostAttributesWithDefaults() *OrgInvitationPostAttributes`

NewOrgInvitationPostAttributesWithDefaults instantiates a new OrgInvitationPostAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrgInvitationPostAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgInvitationPostAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgInvitationPostAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *OrgInvitationPostAttributes) GetRole() interface{}`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgInvitationPostAttributes) GetRoleOk() (*interface{}, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgInvitationPostAttributes) SetRole(v interface{})`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *OrgInvitationPostAttributes) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *OrgInvitationPostAttributes) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


