/*
Snyk API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: REST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AutoDependencyUpgradeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoDependencyUpgradeSettings{}

// AutoDependencyUpgradeSettings Automatically create pull requests on recurring tests for dependencies as upgrades become available. If not specified, settings will be inherited from the Project's integration.
type AutoDependencyUpgradeSettings struct {
	// Dependencies which should NOT be included in an automatic upgrade operation.
	IgnoredDependencies []string `json:"ignored_dependencies,omitempty"`
	// Automatically raise pull requests to update out-of-date dependencies.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Apply the auto dependency integration settings of the Organization to this project.
	IsInherited *bool `json:"is_inherited,omitempty"`
	// Include major version in dependency upgrade recommendation.
	IsMajorUpgradeEnabled *bool `json:"is_major_upgrade_enabled,omitempty"`
	// Limit of dependency upgrade PRs which can be opened simultaneously. When the limit is reached, no new upgrade PRs are created. If specified, must be between 1 and 10.
	Limit *float32 `json:"limit,omitempty"`
	// Minimum dependency maturity period in days. If specified, must be between 1 and 365.
	MinimumAge *float32 `json:"minimum_age,omitempty"`
}

// NewAutoDependencyUpgradeSettings instantiates a new AutoDependencyUpgradeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoDependencyUpgradeSettings() *AutoDependencyUpgradeSettings {
	this := AutoDependencyUpgradeSettings{}
	return &this
}

// NewAutoDependencyUpgradeSettingsWithDefaults instantiates a new AutoDependencyUpgradeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoDependencyUpgradeSettingsWithDefaults() *AutoDependencyUpgradeSettings {
	this := AutoDependencyUpgradeSettings{}
	return &this
}

// GetIgnoredDependencies returns the IgnoredDependencies field value if set, zero value otherwise.
func (o *AutoDependencyUpgradeSettings) GetIgnoredDependencies() []string {
	if o == nil || IsNil(o.IgnoredDependencies) {
		var ret []string
		return ret
	}
	return o.IgnoredDependencies
}

// GetIgnoredDependenciesOk returns a tuple with the IgnoredDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoDependencyUpgradeSettings) GetIgnoredDependenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.IgnoredDependencies) {
		return nil, false
	}
	return o.IgnoredDependencies, true
}

// HasIgnoredDependencies returns a boolean if a field has been set.
func (o *AutoDependencyUpgradeSettings) HasIgnoredDependencies() bool {
	if o != nil && !IsNil(o.IgnoredDependencies) {
		return true
	}

	return false
}

// SetIgnoredDependencies gets a reference to the given []string and assigns it to the IgnoredDependencies field.
func (o *AutoDependencyUpgradeSettings) SetIgnoredDependencies(v []string) {
	o.IgnoredDependencies = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *AutoDependencyUpgradeSettings) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoDependencyUpgradeSettings) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *AutoDependencyUpgradeSettings) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *AutoDependencyUpgradeSettings) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsInherited returns the IsInherited field value if set, zero value otherwise.
func (o *AutoDependencyUpgradeSettings) GetIsInherited() bool {
	if o == nil || IsNil(o.IsInherited) {
		var ret bool
		return ret
	}
	return *o.IsInherited
}

// GetIsInheritedOk returns a tuple with the IsInherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoDependencyUpgradeSettings) GetIsInheritedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInherited) {
		return nil, false
	}
	return o.IsInherited, true
}

// HasIsInherited returns a boolean if a field has been set.
func (o *AutoDependencyUpgradeSettings) HasIsInherited() bool {
	if o != nil && !IsNil(o.IsInherited) {
		return true
	}

	return false
}

// SetIsInherited gets a reference to the given bool and assigns it to the IsInherited field.
func (o *AutoDependencyUpgradeSettings) SetIsInherited(v bool) {
	o.IsInherited = &v
}

// GetIsMajorUpgradeEnabled returns the IsMajorUpgradeEnabled field value if set, zero value otherwise.
func (o *AutoDependencyUpgradeSettings) GetIsMajorUpgradeEnabled() bool {
	if o == nil || IsNil(o.IsMajorUpgradeEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMajorUpgradeEnabled
}

// GetIsMajorUpgradeEnabledOk returns a tuple with the IsMajorUpgradeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoDependencyUpgradeSettings) GetIsMajorUpgradeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMajorUpgradeEnabled) {
		return nil, false
	}
	return o.IsMajorUpgradeEnabled, true
}

// HasIsMajorUpgradeEnabled returns a boolean if a field has been set.
func (o *AutoDependencyUpgradeSettings) HasIsMajorUpgradeEnabled() bool {
	if o != nil && !IsNil(o.IsMajorUpgradeEnabled) {
		return true
	}

	return false
}

// SetIsMajorUpgradeEnabled gets a reference to the given bool and assigns it to the IsMajorUpgradeEnabled field.
func (o *AutoDependencyUpgradeSettings) SetIsMajorUpgradeEnabled(v bool) {
	o.IsMajorUpgradeEnabled = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AutoDependencyUpgradeSettings) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoDependencyUpgradeSettings) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AutoDependencyUpgradeSettings) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *AutoDependencyUpgradeSettings) SetLimit(v float32) {
	o.Limit = &v
}

// GetMinimumAge returns the MinimumAge field value if set, zero value otherwise.
func (o *AutoDependencyUpgradeSettings) GetMinimumAge() float32 {
	if o == nil || IsNil(o.MinimumAge) {
		var ret float32
		return ret
	}
	return *o.MinimumAge
}

// GetMinimumAgeOk returns a tuple with the MinimumAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoDependencyUpgradeSettings) GetMinimumAgeOk() (*float32, bool) {
	if o == nil || IsNil(o.MinimumAge) {
		return nil, false
	}
	return o.MinimumAge, true
}

// HasMinimumAge returns a boolean if a field has been set.
func (o *AutoDependencyUpgradeSettings) HasMinimumAge() bool {
	if o != nil && !IsNil(o.MinimumAge) {
		return true
	}

	return false
}

// SetMinimumAge gets a reference to the given float32 and assigns it to the MinimumAge field.
func (o *AutoDependencyUpgradeSettings) SetMinimumAge(v float32) {
	o.MinimumAge = &v
}

func (o AutoDependencyUpgradeSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoDependencyUpgradeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IgnoredDependencies) {
		toSerialize["ignored_dependencies"] = o.IgnoredDependencies
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.IsInherited) {
		toSerialize["is_inherited"] = o.IsInherited
	}
	if !IsNil(o.IsMajorUpgradeEnabled) {
		toSerialize["is_major_upgrade_enabled"] = o.IsMajorUpgradeEnabled
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.MinimumAge) {
		toSerialize["minimum_age"] = o.MinimumAge
	}
	return toSerialize, nil
}

type NullableAutoDependencyUpgradeSettings struct {
	value *AutoDependencyUpgradeSettings
	isSet bool
}

func (v NullableAutoDependencyUpgradeSettings) Get() *AutoDependencyUpgradeSettings {
	return v.value
}

func (v *NullableAutoDependencyUpgradeSettings) Set(val *AutoDependencyUpgradeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoDependencyUpgradeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoDependencyUpgradeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoDependencyUpgradeSettings(val *AutoDependencyUpgradeSettings) *NullableAutoDependencyUpgradeSettings {
	return &NullableAutoDependencyUpgradeSettings{value: val, isSet: true}
}

func (v NullableAutoDependencyUpgradeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoDependencyUpgradeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

