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

// checks if the UpdateServiceAccountSecretRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceAccountSecretRequestData{}

// UpdateServiceAccountSecretRequestData struct for UpdateServiceAccountSecretRequestData
type UpdateServiceAccountSecretRequestData struct {
	Attributes UpdateServiceAccountSecretRequestDataAttributes `json:"attributes"`
	// The Resource type.
	Type string `json:"type"`
}

type _UpdateServiceAccountSecretRequestData UpdateServiceAccountSecretRequestData

// NewUpdateServiceAccountSecretRequestData instantiates a new UpdateServiceAccountSecretRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceAccountSecretRequestData(attributes UpdateServiceAccountSecretRequestDataAttributes, type_ string) *UpdateServiceAccountSecretRequestData {
	this := UpdateServiceAccountSecretRequestData{}
	this.Attributes = attributes
	this.Type = type_
	return &this
}

// NewUpdateServiceAccountSecretRequestDataWithDefaults instantiates a new UpdateServiceAccountSecretRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceAccountSecretRequestDataWithDefaults() *UpdateServiceAccountSecretRequestData {
	this := UpdateServiceAccountSecretRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *UpdateServiceAccountSecretRequestData) GetAttributes() UpdateServiceAccountSecretRequestDataAttributes {
	if o == nil {
		var ret UpdateServiceAccountSecretRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceAccountSecretRequestData) GetAttributesOk() (*UpdateServiceAccountSecretRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UpdateServiceAccountSecretRequestData) SetAttributes(v UpdateServiceAccountSecretRequestDataAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value
func (o *UpdateServiceAccountSecretRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceAccountSecretRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateServiceAccountSecretRequestData) SetType(v string) {
	o.Type = v
}

func (o UpdateServiceAccountSecretRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceAccountSecretRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *UpdateServiceAccountSecretRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributes",
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

	varUpdateServiceAccountSecretRequestData := _UpdateServiceAccountSecretRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateServiceAccountSecretRequestData)

	if err != nil {
		return err
	}

	*o = UpdateServiceAccountSecretRequestData(varUpdateServiceAccountSecretRequestData)

	return err
}

type NullableUpdateServiceAccountSecretRequestData struct {
	value *UpdateServiceAccountSecretRequestData
	isSet bool
}

func (v NullableUpdateServiceAccountSecretRequestData) Get() *UpdateServiceAccountSecretRequestData {
	return v.value
}

func (v *NullableUpdateServiceAccountSecretRequestData) Set(val *UpdateServiceAccountSecretRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceAccountSecretRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceAccountSecretRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceAccountSecretRequestData(val *UpdateServiceAccountSecretRequestData) *NullableUpdateServiceAccountSecretRequestData {
	return &NullableUpdateServiceAccountSecretRequestData{value: val, isSet: true}
}

func (v NullableUpdateServiceAccountSecretRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceAccountSecretRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

