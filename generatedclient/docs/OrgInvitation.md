# OrgInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**OrgInvitationAttributes**](OrgInvitationAttributes.md) |  | 
**Id** | **string** |  | 
**Relationships** | Pointer to [**OrgInvitationRelationships**](OrgInvitationRelationships.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewOrgInvitation

`func NewOrgInvitation(attributes OrgInvitationAttributes, id string, type_ string, ) *OrgInvitation`

NewOrgInvitation instantiates a new OrgInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgInvitationWithDefaults

`func NewOrgInvitationWithDefaults() *OrgInvitation`

NewOrgInvitationWithDefaults instantiates a new OrgInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OrgInvitation) GetAttributes() OrgInvitationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrgInvitation) GetAttributesOk() (*OrgInvitationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrgInvitation) SetAttributes(v OrgInvitationAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *OrgInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgInvitation) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *OrgInvitation) GetRelationships() OrgInvitationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrgInvitation) GetRelationshipsOk() (*OrgInvitationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrgInvitation) SetRelationships(v OrgInvitationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrgInvitation) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *OrgInvitation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgInvitation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgInvitation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


