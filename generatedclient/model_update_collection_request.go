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

// checks if the UpdateCollectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCollectionRequest{}

// UpdateCollectionRequest struct for UpdateCollectionRequest
type UpdateCollectionRequest struct {
	Data UpdateCollectionRequestData `json:"data"`
}

type _UpdateCollectionRequest UpdateCollectionRequest

// NewUpdateCollectionRequest instantiates a new UpdateCollectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCollectionRequest(data UpdateCollectionRequestData) *UpdateCollectionRequest {
	this := UpdateCollectionRequest{}
	this.Data = data
	return &this
}

// NewUpdateCollectionRequestWithDefaults instantiates a new UpdateCollectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCollectionRequestWithDefaults() *UpdateCollectionRequest {
	this := UpdateCollectionRequest{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateCollectionRequest) GetData() UpdateCollectionRequestData {
	if o == nil {
		var ret UpdateCollectionRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateCollectionRequest) GetDataOk() (*UpdateCollectionRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateCollectionRequest) SetData(v UpdateCollectionRequestData) {
	o.Data = v
}

func (o UpdateCollectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCollectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UpdateCollectionRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateCollectionRequest := _UpdateCollectionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCollectionRequest)

	if err != nil {
		return err
	}

	*o = UpdateCollectionRequest(varUpdateCollectionRequest)

	return err
}

type NullableUpdateCollectionRequest struct {
	value *UpdateCollectionRequest
	isSet bool
}

func (v NullableUpdateCollectionRequest) Get() *UpdateCollectionRequest {
	return v.value
}

func (v *NullableUpdateCollectionRequest) Set(val *UpdateCollectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCollectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCollectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCollectionRequest(val *UpdateCollectionRequest) *NullableUpdateCollectionRequest {
	return &NullableUpdateCollectionRequest{value: val, isSet: true}
}

func (v NullableUpdateCollectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCollectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

