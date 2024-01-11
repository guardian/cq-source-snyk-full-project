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

// checks if the UpdateOrg200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrg200ResponseData{}

// UpdateOrg200ResponseData org resource object
type UpdateOrg200ResponseData struct {
	Attributes *OrgAttributes `json:"attributes,omitempty"`
	Id string `json:"id"`
	Relationships *OrgRelationships `json:"relationships,omitempty"`
	Type string `json:"type"`
}

type _UpdateOrg200ResponseData UpdateOrg200ResponseData

// NewUpdateOrg200ResponseData instantiates a new UpdateOrg200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrg200ResponseData(id string, type_ string) *UpdateOrg200ResponseData {
	this := UpdateOrg200ResponseData{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewUpdateOrg200ResponseDataWithDefaults instantiates a new UpdateOrg200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrg200ResponseDataWithDefaults() *UpdateOrg200ResponseData {
	this := UpdateOrg200ResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateOrg200ResponseData) GetAttributes() OrgAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret OrgAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrg200ResponseData) GetAttributesOk() (*OrgAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateOrg200ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given OrgAttributes and assigns it to the Attributes field.
func (o *UpdateOrg200ResponseData) SetAttributes(v OrgAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value
func (o *UpdateOrg200ResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateOrg200ResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateOrg200ResponseData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UpdateOrg200ResponseData) GetRelationships() OrgRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret OrgRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrg200ResponseData) GetRelationshipsOk() (*OrgRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UpdateOrg200ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrgRelationships and assigns it to the Relationships field.
func (o *UpdateOrg200ResponseData) SetRelationships(v OrgRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value
func (o *UpdateOrg200ResponseData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateOrg200ResponseData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateOrg200ResponseData) SetType(v string) {
	o.Type = v
}

func (o UpdateOrg200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrg200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *UpdateOrg200ResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateOrg200ResponseData := _UpdateOrg200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOrg200ResponseData)

	if err != nil {
		return err
	}

	*o = UpdateOrg200ResponseData(varUpdateOrg200ResponseData)

	return err
}

type NullableUpdateOrg200ResponseData struct {
	value *UpdateOrg200ResponseData
	isSet bool
}

func (v NullableUpdateOrg200ResponseData) Get() *UpdateOrg200ResponseData {
	return v.value
}

func (v *NullableUpdateOrg200ResponseData) Set(val *UpdateOrg200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrg200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrg200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrg200ResponseData(val *UpdateOrg200ResponseData) *NullableUpdateOrg200ResponseData {
	return &NullableUpdateOrg200ResponseData{value: val, isSet: true}
}

func (v NullableUpdateOrg200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrg200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


