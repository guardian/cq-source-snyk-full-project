/*
Snyk API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: REST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VersioningSchemaSemverType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersioningSchemaSemverType{}

// VersioningSchemaSemverType struct for VersioningSchemaSemverType
type VersioningSchemaSemverType struct {
	Type string `json:"type"`
}

type _VersioningSchemaSemverType VersioningSchemaSemverType

// NewVersioningSchemaSemverType instantiates a new VersioningSchemaSemverType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersioningSchemaSemverType(type_ string) *VersioningSchemaSemverType {
	this := VersioningSchemaSemverType{}
	this.Type = type_
	return &this
}

// NewVersioningSchemaSemverTypeWithDefaults instantiates a new VersioningSchemaSemverType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersioningSchemaSemverTypeWithDefaults() *VersioningSchemaSemverType {
	this := VersioningSchemaSemverType{}
	return &this
}

// GetType returns the Type field value
func (o *VersioningSchemaSemverType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VersioningSchemaSemverType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VersioningSchemaSemverType) SetType(v string) {
	o.Type = v
}

func (o VersioningSchemaSemverType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersioningSchemaSemverType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *VersioningSchemaSemverType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVersioningSchemaSemverType := _VersioningSchemaSemverType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVersioningSchemaSemverType)

	if err != nil {
		return err
	}

	*o = VersioningSchemaSemverType(varVersioningSchemaSemverType)

	return err
}

type NullableVersioningSchemaSemverType struct {
	value *VersioningSchemaSemverType
	isSet bool
}

func (v NullableVersioningSchemaSemverType) Get() *VersioningSchemaSemverType {
	return v.value
}

func (v *NullableVersioningSchemaSemverType) Set(val *VersioningSchemaSemverType) {
	v.value = val
	v.isSet = true
}

func (v NullableVersioningSchemaSemverType) IsSet() bool {
	return v.isSet
}

func (v *NullableVersioningSchemaSemverType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersioningSchemaSemverType(val *VersioningSchemaSemverType) *NullableVersioningSchemaSemverType {
	return &NullableVersioningSchemaSemverType{value: val, isSet: true}
}

func (v NullableVersioningSchemaSemverType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersioningSchemaSemverType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

