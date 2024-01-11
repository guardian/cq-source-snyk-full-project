# OrgAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The time the organization was created. | [optional] 
**GroupId** | Pointer to **string** | The Snyk ID of the group to which the organization belongs. | [optional] 
**IsPersonal** | **bool** | Whether the organization is independent (that is, not part of a group). | 
**Name** | **string** | The display name of the organization. | 
**Slug** | **string** | The canonical (unique and URL-friendly) name of the organization. | 
**UpdatedAt** | Pointer to **time.Time** | The time the organization was last modified. | [optional] 

## Methods

### NewOrgAttributes

`func NewOrgAttributes(isPersonal bool, name string, slug string, ) *OrgAttributes`

NewOrgAttributes instantiates a new OrgAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgAttributesWithDefaults

`func NewOrgAttributesWithDefaults() *OrgAttributes`

NewOrgAttributesWithDefaults instantiates a new OrgAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrgAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrgAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGroupId

`func (o *OrgAttributes) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OrgAttributes) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OrgAttributes) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *OrgAttributes) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIsPersonal

`func (o *OrgAttributes) GetIsPersonal() bool`

GetIsPersonal returns the IsPersonal field if non-nil, zero value otherwise.

### GetIsPersonalOk

`func (o *OrgAttributes) GetIsPersonalOk() (*bool, bool)`

GetIsPersonalOk returns a tuple with the IsPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonal

`func (o *OrgAttributes) SetIsPersonal(v bool)`

SetIsPersonal sets IsPersonal field to given value.


### GetName

`func (o *OrgAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OrgAttributes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrgAttributes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrgAttributes) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetUpdatedAt

`func (o *OrgAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrgAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


