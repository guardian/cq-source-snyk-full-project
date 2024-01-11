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

// checks if the ProjectRelationshipsTargetData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectRelationshipsTargetData{}

// ProjectRelationshipsTargetData struct for ProjectRelationshipsTargetData
type ProjectRelationshipsTargetData struct {
	Attributes ProjectRelationshipsTargetDataAttributes `json:"attributes"`
	// The Resource ID.
	Id string `json:"id"`
	Meta *ProjectRelationshipsTargetDataMeta `json:"meta,omitempty"`
	// The Resource type.
	Type string `json:"type"`
}

type _ProjectRelationshipsTargetData ProjectRelationshipsTargetData

// NewProjectRelationshipsTargetData instantiates a new ProjectRelationshipsTargetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRelationshipsTargetData(attributes ProjectRelationshipsTargetDataAttributes, id string, type_ string) *ProjectRelationshipsTargetData {
	this := ProjectRelationshipsTargetData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = type_
	return &this
}

// NewProjectRelationshipsTargetDataWithDefaults instantiates a new ProjectRelationshipsTargetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRelationshipsTargetDataWithDefaults() *ProjectRelationshipsTargetData {
	this := ProjectRelationshipsTargetData{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *ProjectRelationshipsTargetData) GetAttributes() ProjectRelationshipsTargetDataAttributes {
	if o == nil {
		var ret ProjectRelationshipsTargetDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProjectRelationshipsTargetData) GetAttributesOk() (*ProjectRelationshipsTargetDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ProjectRelationshipsTargetData) SetAttributes(v ProjectRelationshipsTargetDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value
func (o *ProjectRelationshipsTargetData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectRelationshipsTargetData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectRelationshipsTargetData) SetId(v string) {
	o.Id = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProjectRelationshipsTargetData) GetMeta() ProjectRelationshipsTargetDataMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ProjectRelationshipsTargetDataMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRelationshipsTargetData) GetMetaOk() (*ProjectRelationshipsTargetDataMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProjectRelationshipsTargetData) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ProjectRelationshipsTargetDataMeta and assigns it to the Meta field.
func (o *ProjectRelationshipsTargetData) SetMeta(v ProjectRelationshipsTargetDataMeta) {
	o.Meta = &v
}

// GetType returns the Type field value
func (o *ProjectRelationshipsTargetData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProjectRelationshipsTargetData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProjectRelationshipsTargetData) SetType(v string) {
	o.Type = v
}

func (o ProjectRelationshipsTargetData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectRelationshipsTargetData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ProjectRelationshipsTargetData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributes",
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

	varProjectRelationshipsTargetData := _ProjectRelationshipsTargetData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectRelationshipsTargetData)

	if err != nil {
		return err
	}

	*o = ProjectRelationshipsTargetData(varProjectRelationshipsTargetData)

	return err
}

type NullableProjectRelationshipsTargetData struct {
	value *ProjectRelationshipsTargetData
	isSet bool
}

func (v NullableProjectRelationshipsTargetData) Get() *ProjectRelationshipsTargetData {
	return v.value
}

func (v *NullableProjectRelationshipsTargetData) Set(val *ProjectRelationshipsTargetData) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRelationshipsTargetData) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRelationshipsTargetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRelationshipsTargetData(val *ProjectRelationshipsTargetData) *NullableProjectRelationshipsTargetData {
	return &NullableProjectRelationshipsTargetData{value: val, isSet: true}
}

func (v NullableProjectRelationshipsTargetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRelationshipsTargetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


