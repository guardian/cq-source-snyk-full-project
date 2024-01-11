# OrgIacSettingsResponseAttributesCustomRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InheritFromParent** | Pointer to [**InheritFromParent**](InheritFromParent.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the custom rules feature is enabled or not. | [optional] 
**OciRegistryTag** | Pointer to **string** | The tag for an OCI artifact inside an OCI registry. | [optional] 
**OciRegistryUrl** | Pointer to **string** | The URL to an OCI registry. | [optional] 
**Parents** | Pointer to [**OrgIacSettingsResponseAttributesCustomRulesParents**](OrgIacSettingsResponseAttributesCustomRulesParents.md) |  | [optional] 
**Updated** | Pointer to **time.Time** | The last time the settings were updated. | [optional] 

## Methods

### NewOrgIacSettingsResponseAttributesCustomRules

`func NewOrgIacSettingsResponseAttributesCustomRules() *OrgIacSettingsResponseAttributesCustomRules`

NewOrgIacSettingsResponseAttributesCustomRules instantiates a new OrgIacSettingsResponseAttributesCustomRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgIacSettingsResponseAttributesCustomRulesWithDefaults

`func NewOrgIacSettingsResponseAttributesCustomRulesWithDefaults() *OrgIacSettingsResponseAttributesCustomRules`

NewOrgIacSettingsResponseAttributesCustomRulesWithDefaults instantiates a new OrgIacSettingsResponseAttributesCustomRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInheritFromParent

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetInheritFromParent() InheritFromParent`

GetInheritFromParent returns the InheritFromParent field if non-nil, zero value otherwise.

### GetInheritFromParentOk

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetInheritFromParentOk() (*InheritFromParent, bool)`

GetInheritFromParentOk returns a tuple with the InheritFromParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritFromParent

`func (o *OrgIacSettingsResponseAttributesCustomRules) SetInheritFromParent(v InheritFromParent)`

SetInheritFromParent sets InheritFromParent field to given value.

### HasInheritFromParent

`func (o *OrgIacSettingsResponseAttributesCustomRules) HasInheritFromParent() bool`

HasInheritFromParent returns a boolean if a field has been set.

### GetIsEnabled

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *OrgIacSettingsResponseAttributesCustomRules) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *OrgIacSettingsResponseAttributesCustomRules) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetOciRegistryTag

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetOciRegistryTag() string`

GetOciRegistryTag returns the OciRegistryTag field if non-nil, zero value otherwise.

### GetOciRegistryTagOk

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetOciRegistryTagOk() (*string, bool)`

GetOciRegistryTagOk returns a tuple with the OciRegistryTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegistryTag

`func (o *OrgIacSettingsResponseAttributesCustomRules) SetOciRegistryTag(v string)`

SetOciRegistryTag sets OciRegistryTag field to given value.

### HasOciRegistryTag

`func (o *OrgIacSettingsResponseAttributesCustomRules) HasOciRegistryTag() bool`

HasOciRegistryTag returns a boolean if a field has been set.

### GetOciRegistryUrl

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetOciRegistryUrl() string`

GetOciRegistryUrl returns the OciRegistryUrl field if non-nil, zero value otherwise.

### GetOciRegistryUrlOk

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetOciRegistryUrlOk() (*string, bool)`

GetOciRegistryUrlOk returns a tuple with the OciRegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegistryUrl

`func (o *OrgIacSettingsResponseAttributesCustomRules) SetOciRegistryUrl(v string)`

SetOciRegistryUrl sets OciRegistryUrl field to given value.

### HasOciRegistryUrl

`func (o *OrgIacSettingsResponseAttributesCustomRules) HasOciRegistryUrl() bool`

HasOciRegistryUrl returns a boolean if a field has been set.

### GetParents

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetParents() OrgIacSettingsResponseAttributesCustomRulesParents`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetParentsOk() (*OrgIacSettingsResponseAttributesCustomRulesParents, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *OrgIacSettingsResponseAttributesCustomRules) SetParents(v OrgIacSettingsResponseAttributesCustomRulesParents)`

SetParents sets Parents field to given value.

### HasParents

`func (o *OrgIacSettingsResponseAttributesCustomRules) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetUpdated

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *OrgIacSettingsResponseAttributesCustomRules) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *OrgIacSettingsResponseAttributesCustomRules) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *OrgIacSettingsResponseAttributesCustomRules) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


