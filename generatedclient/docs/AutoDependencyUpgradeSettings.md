# AutoDependencyUpgradeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoredDependencies** | Pointer to **[]string** | Dependencies which should NOT be included in an automatic upgrade operation. | [optional] 
**IsEnabled** | Pointer to **bool** | Automatically raise pull requests to update out-of-date dependencies. | [optional] 
**IsInherited** | Pointer to **bool** | Apply the auto dependency integration settings of the Organization to this project. | [optional] 
**IsMajorUpgradeEnabled** | Pointer to **bool** | Include major version in dependency upgrade recommendation. | [optional] 
**Limit** | Pointer to **float32** | Limit of dependency upgrade PRs which can be opened simultaneously. When the limit is reached, no new upgrade PRs are created. If specified, must be between 1 and 10. | [optional] 
**MinimumAge** | Pointer to **float32** | Minimum dependency maturity period in days. If specified, must be between 1 and 365. | [optional] 

## Methods

### NewAutoDependencyUpgradeSettings

`func NewAutoDependencyUpgradeSettings() *AutoDependencyUpgradeSettings`

NewAutoDependencyUpgradeSettings instantiates a new AutoDependencyUpgradeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoDependencyUpgradeSettingsWithDefaults

`func NewAutoDependencyUpgradeSettingsWithDefaults() *AutoDependencyUpgradeSettings`

NewAutoDependencyUpgradeSettingsWithDefaults instantiates a new AutoDependencyUpgradeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoredDependencies

`func (o *AutoDependencyUpgradeSettings) GetIgnoredDependencies() []string`

GetIgnoredDependencies returns the IgnoredDependencies field if non-nil, zero value otherwise.

### GetIgnoredDependenciesOk

`func (o *AutoDependencyUpgradeSettings) GetIgnoredDependenciesOk() (*[]string, bool)`

GetIgnoredDependenciesOk returns a tuple with the IgnoredDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredDependencies

`func (o *AutoDependencyUpgradeSettings) SetIgnoredDependencies(v []string)`

SetIgnoredDependencies sets IgnoredDependencies field to given value.

### HasIgnoredDependencies

`func (o *AutoDependencyUpgradeSettings) HasIgnoredDependencies() bool`

HasIgnoredDependencies returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AutoDependencyUpgradeSettings) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AutoDependencyUpgradeSettings) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AutoDependencyUpgradeSettings) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AutoDependencyUpgradeSettings) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsInherited

`func (o *AutoDependencyUpgradeSettings) GetIsInherited() bool`

GetIsInherited returns the IsInherited field if non-nil, zero value otherwise.

### GetIsInheritedOk

`func (o *AutoDependencyUpgradeSettings) GetIsInheritedOk() (*bool, bool)`

GetIsInheritedOk returns a tuple with the IsInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInherited

`func (o *AutoDependencyUpgradeSettings) SetIsInherited(v bool)`

SetIsInherited sets IsInherited field to given value.

### HasIsInherited

`func (o *AutoDependencyUpgradeSettings) HasIsInherited() bool`

HasIsInherited returns a boolean if a field has been set.

### GetIsMajorUpgradeEnabled

`func (o *AutoDependencyUpgradeSettings) GetIsMajorUpgradeEnabled() bool`

GetIsMajorUpgradeEnabled returns the IsMajorUpgradeEnabled field if non-nil, zero value otherwise.

### GetIsMajorUpgradeEnabledOk

`func (o *AutoDependencyUpgradeSettings) GetIsMajorUpgradeEnabledOk() (*bool, bool)`

GetIsMajorUpgradeEnabledOk returns a tuple with the IsMajorUpgradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorUpgradeEnabled

`func (o *AutoDependencyUpgradeSettings) SetIsMajorUpgradeEnabled(v bool)`

SetIsMajorUpgradeEnabled sets IsMajorUpgradeEnabled field to given value.

### HasIsMajorUpgradeEnabled

`func (o *AutoDependencyUpgradeSettings) HasIsMajorUpgradeEnabled() bool`

HasIsMajorUpgradeEnabled returns a boolean if a field has been set.

### GetLimit

`func (o *AutoDependencyUpgradeSettings) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AutoDependencyUpgradeSettings) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AutoDependencyUpgradeSettings) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AutoDependencyUpgradeSettings) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMinimumAge

`func (o *AutoDependencyUpgradeSettings) GetMinimumAge() float32`

GetMinimumAge returns the MinimumAge field if non-nil, zero value otherwise.

### GetMinimumAgeOk

`func (o *AutoDependencyUpgradeSettings) GetMinimumAgeOk() (*float32, bool)`

GetMinimumAgeOk returns a tuple with the MinimumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAge

`func (o *AutoDependencyUpgradeSettings) SetMinimumAge(v float32)`

SetMinimumAge sets MinimumAge field to given value.

### HasMinimumAge

`func (o *AutoDependencyUpgradeSettings) HasMinimumAge() bool`

HasMinimumAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


