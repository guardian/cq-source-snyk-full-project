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

// checks if the NugetBuildArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NugetBuildArgs{}

// NugetBuildArgs struct for NugetBuildArgs
type NugetBuildArgs struct {
	TargetFramework string `json:"target_framework"`
}

type _NugetBuildArgs NugetBuildArgs

// NewNugetBuildArgs instantiates a new NugetBuildArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNugetBuildArgs(targetFramework string) *NugetBuildArgs {
	this := NugetBuildArgs{}
	this.TargetFramework = targetFramework
	return &this
}

// NewNugetBuildArgsWithDefaults instantiates a new NugetBuildArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNugetBuildArgsWithDefaults() *NugetBuildArgs {
	this := NugetBuildArgs{}
	return &this
}

// GetTargetFramework returns the TargetFramework field value
func (o *NugetBuildArgs) GetTargetFramework() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetFramework
}

// GetTargetFrameworkOk returns a tuple with the TargetFramework field value
// and a boolean to check if the value has been set.
func (o *NugetBuildArgs) GetTargetFrameworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetFramework, true
}

// SetTargetFramework sets field value
func (o *NugetBuildArgs) SetTargetFramework(v string) {
	o.TargetFramework = v
}

func (o NugetBuildArgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NugetBuildArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target_framework"] = o.TargetFramework
	return toSerialize, nil
}

func (o *NugetBuildArgs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"target_framework",
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

	varNugetBuildArgs := _NugetBuildArgs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNugetBuildArgs)

	if err != nil {
		return err
	}

	*o = NugetBuildArgs(varNugetBuildArgs)

	return err
}

type NullableNugetBuildArgs struct {
	value *NugetBuildArgs
	isSet bool
}

func (v NullableNugetBuildArgs) Get() *NugetBuildArgs {
	return v.value
}

func (v *NullableNugetBuildArgs) Set(val *NugetBuildArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableNugetBuildArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableNugetBuildArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNugetBuildArgs(val *NugetBuildArgs) *NullableNugetBuildArgs {
	return &NullableNugetBuildArgs{value: val, isSet: true}
}

func (v NullableNugetBuildArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNugetBuildArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


