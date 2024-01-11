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

// checks if the OrgIacSettingsResponseAttributesCustomRulesParents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgIacSettingsResponseAttributesCustomRulesParents{}

// OrgIacSettingsResponseAttributesCustomRulesParents Contains all parents the org can inherit settings from.
type OrgIacSettingsResponseAttributesCustomRulesParents struct {
	Group *OrgIacSettingsResponseAttributesCustomRulesParentsGroup `json:"group,omitempty"`
}

// NewOrgIacSettingsResponseAttributesCustomRulesParents instantiates a new OrgIacSettingsResponseAttributesCustomRulesParents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgIacSettingsResponseAttributesCustomRulesParents() *OrgIacSettingsResponseAttributesCustomRulesParents {
	this := OrgIacSettingsResponseAttributesCustomRulesParents{}
	return &this
}

// NewOrgIacSettingsResponseAttributesCustomRulesParentsWithDefaults instantiates a new OrgIacSettingsResponseAttributesCustomRulesParents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgIacSettingsResponseAttributesCustomRulesParentsWithDefaults() *OrgIacSettingsResponseAttributesCustomRulesParents {
	this := OrgIacSettingsResponseAttributesCustomRulesParents{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *OrgIacSettingsResponseAttributesCustomRulesParents) GetGroup() OrgIacSettingsResponseAttributesCustomRulesParentsGroup {
	if o == nil || IsNil(o.Group) {
		var ret OrgIacSettingsResponseAttributesCustomRulesParentsGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgIacSettingsResponseAttributesCustomRulesParents) GetGroupOk() (*OrgIacSettingsResponseAttributesCustomRulesParentsGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *OrgIacSettingsResponseAttributesCustomRulesParents) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given OrgIacSettingsResponseAttributesCustomRulesParentsGroup and assigns it to the Group field.
func (o *OrgIacSettingsResponseAttributesCustomRulesParents) SetGroup(v OrgIacSettingsResponseAttributesCustomRulesParentsGroup) {
	o.Group = &v
}

func (o OrgIacSettingsResponseAttributesCustomRulesParents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgIacSettingsResponseAttributesCustomRulesParents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableOrgIacSettingsResponseAttributesCustomRulesParents struct {
	value *OrgIacSettingsResponseAttributesCustomRulesParents
	isSet bool
}

func (v NullableOrgIacSettingsResponseAttributesCustomRulesParents) Get() *OrgIacSettingsResponseAttributesCustomRulesParents {
	return v.value
}

func (v *NullableOrgIacSettingsResponseAttributesCustomRulesParents) Set(val *OrgIacSettingsResponseAttributesCustomRulesParents) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgIacSettingsResponseAttributesCustomRulesParents) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgIacSettingsResponseAttributesCustomRulesParents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgIacSettingsResponseAttributesCustomRulesParents(val *OrgIacSettingsResponseAttributesCustomRulesParents) *NullableOrgIacSettingsResponseAttributesCustomRulesParents {
	return &NullableOrgIacSettingsResponseAttributesCustomRulesParents{value: val, isSet: true}
}

func (v NullableOrgIacSettingsResponseAttributesCustomRulesParents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgIacSettingsResponseAttributesCustomRulesParents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

