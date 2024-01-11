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

// checks if the UpdateOrgServiceAccountRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrgServiceAccountRequestDataAttributes{}

// UpdateOrgServiceAccountRequestDataAttributes struct for UpdateOrgServiceAccountRequestDataAttributes
type UpdateOrgServiceAccountRequestDataAttributes struct {
	// A human-friendly name for the service account. Must be unique within the organization.
	Name string `json:"name"`
}

type _UpdateOrgServiceAccountRequestDataAttributes UpdateOrgServiceAccountRequestDataAttributes

// NewUpdateOrgServiceAccountRequestDataAttributes instantiates a new UpdateOrgServiceAccountRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrgServiceAccountRequestDataAttributes(name string) *UpdateOrgServiceAccountRequestDataAttributes {
	this := UpdateOrgServiceAccountRequestDataAttributes{}
	this.Name = name
	return &this
}

// NewUpdateOrgServiceAccountRequestDataAttributesWithDefaults instantiates a new UpdateOrgServiceAccountRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrgServiceAccountRequestDataAttributesWithDefaults() *UpdateOrgServiceAccountRequestDataAttributes {
	this := UpdateOrgServiceAccountRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateOrgServiceAccountRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateOrgServiceAccountRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateOrgServiceAccountRequestDataAttributes) SetName(v string) {
	o.Name = v
}

func (o UpdateOrgServiceAccountRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrgServiceAccountRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UpdateOrgServiceAccountRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varUpdateOrgServiceAccountRequestDataAttributes := _UpdateOrgServiceAccountRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOrgServiceAccountRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = UpdateOrgServiceAccountRequestDataAttributes(varUpdateOrgServiceAccountRequestDataAttributes)

	return err
}

type NullableUpdateOrgServiceAccountRequestDataAttributes struct {
	value *UpdateOrgServiceAccountRequestDataAttributes
	isSet bool
}

func (v NullableUpdateOrgServiceAccountRequestDataAttributes) Get() *UpdateOrgServiceAccountRequestDataAttributes {
	return v.value
}

func (v *NullableUpdateOrgServiceAccountRequestDataAttributes) Set(val *UpdateOrgServiceAccountRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrgServiceAccountRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrgServiceAccountRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrgServiceAccountRequestDataAttributes(val *UpdateOrgServiceAccountRequestDataAttributes) *NullableUpdateOrgServiceAccountRequestDataAttributes {
	return &NullableUpdateOrgServiceAccountRequestDataAttributes{value: val, isSet: true}
}

func (v NullableUpdateOrgServiceAccountRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrgServiceAccountRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


