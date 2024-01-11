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

// checks if the ProjectRelationshipsTargetDataMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectRelationshipsTargetDataMeta{}

// ProjectRelationshipsTargetDataMeta struct for ProjectRelationshipsTargetDataMeta
type ProjectRelationshipsTargetDataMeta struct {
	// A collection of properties regarding integration data
	IntegrationData map[string]interface{} `json:"integration_data,omitempty"`
}

// NewProjectRelationshipsTargetDataMeta instantiates a new ProjectRelationshipsTargetDataMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRelationshipsTargetDataMeta() *ProjectRelationshipsTargetDataMeta {
	this := ProjectRelationshipsTargetDataMeta{}
	return &this
}

// NewProjectRelationshipsTargetDataMetaWithDefaults instantiates a new ProjectRelationshipsTargetDataMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRelationshipsTargetDataMetaWithDefaults() *ProjectRelationshipsTargetDataMeta {
	this := ProjectRelationshipsTargetDataMeta{}
	return &this
}

// GetIntegrationData returns the IntegrationData field value if set, zero value otherwise.
func (o *ProjectRelationshipsTargetDataMeta) GetIntegrationData() map[string]interface{} {
	if o == nil || IsNil(o.IntegrationData) {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationData
}

// GetIntegrationDataOk returns a tuple with the IntegrationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRelationshipsTargetDataMeta) GetIntegrationDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IntegrationData) {
		return map[string]interface{}{}, false
	}
	return o.IntegrationData, true
}

// HasIntegrationData returns a boolean if a field has been set.
func (o *ProjectRelationshipsTargetDataMeta) HasIntegrationData() bool {
	if o != nil && !IsNil(o.IntegrationData) {
		return true
	}

	return false
}

// SetIntegrationData gets a reference to the given map[string]interface{} and assigns it to the IntegrationData field.
func (o *ProjectRelationshipsTargetDataMeta) SetIntegrationData(v map[string]interface{}) {
	o.IntegrationData = v
}

func (o ProjectRelationshipsTargetDataMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectRelationshipsTargetDataMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntegrationData) {
		toSerialize["integration_data"] = o.IntegrationData
	}
	return toSerialize, nil
}

type NullableProjectRelationshipsTargetDataMeta struct {
	value *ProjectRelationshipsTargetDataMeta
	isSet bool
}

func (v NullableProjectRelationshipsTargetDataMeta) Get() *ProjectRelationshipsTargetDataMeta {
	return v.value
}

func (v *NullableProjectRelationshipsTargetDataMeta) Set(val *ProjectRelationshipsTargetDataMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRelationshipsTargetDataMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRelationshipsTargetDataMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRelationshipsTargetDataMeta(val *ProjectRelationshipsTargetDataMeta) *NullableProjectRelationshipsTargetDataMeta {
	return &NullableProjectRelationshipsTargetDataMeta{value: val, isSet: true}
}

func (v NullableProjectRelationshipsTargetDataMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRelationshipsTargetDataMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


