# OrgInvitationPostData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**OrgInvitationPostAttributes**](OrgInvitationPostAttributes.md) |  | 
**Relationships** | Pointer to [**OrgInvitationRelationships**](OrgInvitationRelationships.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewOrgInvitationPostData

`func NewOrgInvitationPostData(attributes OrgInvitationPostAttributes, type_ string, ) *OrgInvitationPostData`

NewOrgInvitationPostData instantiates a new OrgInvitationPostData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgInvitationPostDataWithDefaults

`func NewOrgInvitationPostDataWithDefaults() *OrgInvitationPostData`

NewOrgInvitationPostDataWithDefaults instantiates a new OrgInvitationPostData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OrgInvitationPostData) GetAttributes() OrgInvitationPostAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrgInvitationPostData) GetAttributesOk() (*OrgInvitationPostAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrgInvitationPostData) SetAttributes(v OrgInvitationPostAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrgInvitationPostData) GetRelationships() OrgInvitationRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrgInvitationPostData) GetRelationshipsOk() (*OrgInvitationRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrgInvitationPostData) SetRelationships(v OrgInvitationRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrgInvitationPostData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *OrgInvitationPostData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgInvitationPostData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgInvitationPostData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


