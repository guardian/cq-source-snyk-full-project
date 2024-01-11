# GroupIacSettingsResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRules** | Pointer to [**GroupIacSettingsResponseAttributesCustomRules**](GroupIacSettingsResponseAttributesCustomRules.md) |  | [optional] 
**Updated** | Pointer to **time.Time** | The last time the settings were updated. | [optional] 

## Methods

### NewGroupIacSettingsResponseAttributes

`func NewGroupIacSettingsResponseAttributes() *GroupIacSettingsResponseAttributes`

NewGroupIacSettingsResponseAttributes instantiates a new GroupIacSettingsResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupIacSettingsResponseAttributesWithDefaults

`func NewGroupIacSettingsResponseAttributesWithDefaults() *GroupIacSettingsResponseAttributes`

NewGroupIacSettingsResponseAttributesWithDefaults instantiates a new GroupIacSettingsResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRules

`func (o *GroupIacSettingsResponseAttributes) GetCustomRules() GroupIacSettingsResponseAttributesCustomRules`

GetCustomRules returns the CustomRules field if non-nil, zero value otherwise.

### GetCustomRulesOk

`func (o *GroupIacSettingsResponseAttributes) GetCustomRulesOk() (*GroupIacSettingsResponseAttributesCustomRules, bool)`

GetCustomRulesOk returns a tuple with the CustomRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRules

`func (o *GroupIacSettingsResponseAttributes) SetCustomRules(v GroupIacSettingsResponseAttributesCustomRules)`

SetCustomRules sets CustomRules field to given value.

### HasCustomRules

`func (o *GroupIacSettingsResponseAttributes) HasCustomRules() bool`

HasCustomRules returns a boolean if a field has been set.

### GetUpdated

`func (o *GroupIacSettingsResponseAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GroupIacSettingsResponseAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GroupIacSettingsResponseAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GroupIacSettingsResponseAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


