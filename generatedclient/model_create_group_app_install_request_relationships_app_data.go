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

// checks if the CreateGroupAppInstallRequestRelationshipsAppData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupAppInstallRequestRelationshipsAppData{}

// CreateGroupAppInstallRequestRelationshipsAppData struct for CreateGroupAppInstallRequestRelationshipsAppData
type CreateGroupAppInstallRequestRelationshipsAppData struct {
	Id string `json:"id"`
	Type string `json:"type"`
}

type _CreateGroupAppInstallRequestRelationshipsAppData CreateGroupAppInstallRequestRelationshipsAppData

// NewCreateGroupAppInstallRequestRelationshipsAppData instantiates a new CreateGroupAppInstallRequestRelationshipsAppData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupAppInstallRequestRelationshipsAppData(id string, type_ string) *CreateGroupAppInstallRequestRelationshipsAppData {
	this := CreateGroupAppInstallRequestRelationshipsAppData{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewCreateGroupAppInstallRequestRelationshipsAppDataWithDefaults instantiates a new CreateGroupAppInstallRequestRelationshipsAppData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupAppInstallRequestRelationshipsAppDataWithDefaults() *CreateGroupAppInstallRequestRelationshipsAppData {
	this := CreateGroupAppInstallRequestRelationshipsAppData{}
	return &this
}

// GetId returns the Id field value
func (o *CreateGroupAppInstallRequestRelationshipsAppData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateGroupAppInstallRequestRelationshipsAppData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateGroupAppInstallRequestRelationshipsAppData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *CreateGroupAppInstallRequestRelationshipsAppData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateGroupAppInstallRequestRelationshipsAppData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateGroupAppInstallRequestRelationshipsAppData) SetType(v string) {
	o.Type = v
}

func (o CreateGroupAppInstallRequestRelationshipsAppData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupAppInstallRequestRelationshipsAppData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CreateGroupAppInstallRequestRelationshipsAppData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varCreateGroupAppInstallRequestRelationshipsAppData := _CreateGroupAppInstallRequestRelationshipsAppData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGroupAppInstallRequestRelationshipsAppData)

	if err != nil {
		return err
	}

	*o = CreateGroupAppInstallRequestRelationshipsAppData(varCreateGroupAppInstallRequestRelationshipsAppData)

	return err
}

type NullableCreateGroupAppInstallRequestRelationshipsAppData struct {
	value *CreateGroupAppInstallRequestRelationshipsAppData
	isSet bool
}

func (v NullableCreateGroupAppInstallRequestRelationshipsAppData) Get() *CreateGroupAppInstallRequestRelationshipsAppData {
	return v.value
}

func (v *NullableCreateGroupAppInstallRequestRelationshipsAppData) Set(val *CreateGroupAppInstallRequestRelationshipsAppData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupAppInstallRequestRelationshipsAppData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupAppInstallRequestRelationshipsAppData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupAppInstallRequestRelationshipsAppData(val *CreateGroupAppInstallRequestRelationshipsAppData) *NullableCreateGroupAppInstallRequestRelationshipsAppData {
	return &NullableCreateGroupAppInstallRequestRelationshipsAppData{value: val, isSet: true}
}

func (v NullableCreateGroupAppInstallRequestRelationshipsAppData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupAppInstallRequestRelationshipsAppData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


