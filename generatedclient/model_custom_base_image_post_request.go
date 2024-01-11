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

// checks if the CustomBaseImagePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomBaseImagePostRequest{}

// CustomBaseImagePostRequest struct for CustomBaseImagePostRequest
type CustomBaseImagePostRequest struct {
	Data CustomBaseImagePostRequestData `json:"data"`
}

type _CustomBaseImagePostRequest CustomBaseImagePostRequest

// NewCustomBaseImagePostRequest instantiates a new CustomBaseImagePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomBaseImagePostRequest(data CustomBaseImagePostRequestData) *CustomBaseImagePostRequest {
	this := CustomBaseImagePostRequest{}
	this.Data = data
	return &this
}

// NewCustomBaseImagePostRequestWithDefaults instantiates a new CustomBaseImagePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomBaseImagePostRequestWithDefaults() *CustomBaseImagePostRequest {
	this := CustomBaseImagePostRequest{}
	return &this
}

// GetData returns the Data field value
func (o *CustomBaseImagePostRequest) GetData() CustomBaseImagePostRequestData {
	if o == nil {
		var ret CustomBaseImagePostRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomBaseImagePostRequest) GetDataOk() (*CustomBaseImagePostRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomBaseImagePostRequest) SetData(v CustomBaseImagePostRequestData) {
	o.Data = v
}

func (o CustomBaseImagePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomBaseImagePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CustomBaseImagePostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCustomBaseImagePostRequest := _CustomBaseImagePostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomBaseImagePostRequest)

	if err != nil {
		return err
	}

	*o = CustomBaseImagePostRequest(varCustomBaseImagePostRequest)

	return err
}

type NullableCustomBaseImagePostRequest struct {
	value *CustomBaseImagePostRequest
	isSet bool
}

func (v NullableCustomBaseImagePostRequest) Get() *CustomBaseImagePostRequest {
	return v.value
}

func (v *NullableCustomBaseImagePostRequest) Set(val *CustomBaseImagePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomBaseImagePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomBaseImagePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomBaseImagePostRequest(val *CustomBaseImagePostRequest) *NullableCustomBaseImagePostRequest {
	return &NullableCustomBaseImagePostRequest{value: val, isSet: true}
}

func (v NullableCustomBaseImagePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomBaseImagePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


