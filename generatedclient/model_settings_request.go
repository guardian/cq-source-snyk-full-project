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

// checks if the SettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsRequest{}

// SettingsRequest struct for SettingsRequest
type SettingsRequest struct {
	Data SettingsRequestData `json:"data"`
}

type _SettingsRequest SettingsRequest

// NewSettingsRequest instantiates a new SettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsRequest(data SettingsRequestData) *SettingsRequest {
	this := SettingsRequest{}
	this.Data = data
	return &this
}

// NewSettingsRequestWithDefaults instantiates a new SettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsRequestWithDefaults() *SettingsRequest {
	this := SettingsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SettingsRequest) GetData() SettingsRequestData {
	if o == nil {
		var ret SettingsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetDataOk() (*SettingsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SettingsRequest) SetData(v SettingsRequestData) {
	o.Data = v
}

func (o SettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SettingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varSettingsRequest := _SettingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSettingsRequest)

	if err != nil {
		return err
	}

	*o = SettingsRequest(varSettingsRequest)

	return err
}

type NullableSettingsRequest struct {
	value *SettingsRequest
	isSet bool
}

func (v NullableSettingsRequest) Get() *SettingsRequest {
	return v.value
}

func (v *NullableSettingsRequest) Set(val *SettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsRequest(val *SettingsRequest) *NullableSettingsRequest {
	return &NullableSettingsRequest{value: val, isSet: true}
}

func (v NullableSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


