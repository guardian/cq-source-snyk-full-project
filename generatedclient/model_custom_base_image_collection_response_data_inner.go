/*
Snyk API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: REST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CustomBaseImageCollectionResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomBaseImageCollectionResponseDataInner{}

// CustomBaseImageCollectionResponseDataInner struct for CustomBaseImageCollectionResponseDataInner
type CustomBaseImageCollectionResponseDataInner struct {
	Attributes *CustomBaseImageSnapshot `json:"attributes,omitempty"`
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewCustomBaseImageCollectionResponseDataInner instantiates a new CustomBaseImageCollectionResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomBaseImageCollectionResponseDataInner() *CustomBaseImageCollectionResponseDataInner {
	this := CustomBaseImageCollectionResponseDataInner{}
	return &this
}

// NewCustomBaseImageCollectionResponseDataInnerWithDefaults instantiates a new CustomBaseImageCollectionResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomBaseImageCollectionResponseDataInnerWithDefaults() *CustomBaseImageCollectionResponseDataInner {
	this := CustomBaseImageCollectionResponseDataInner{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomBaseImageCollectionResponseDataInner) GetAttributes() CustomBaseImageSnapshot {
	if o == nil || IsNil(o.Attributes) {
		var ret CustomBaseImageSnapshot
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageCollectionResponseDataInner) GetAttributesOk() (*CustomBaseImageSnapshot, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomBaseImageCollectionResponseDataInner) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CustomBaseImageSnapshot and assigns it to the Attributes field.
func (o *CustomBaseImageCollectionResponseDataInner) SetAttributes(v CustomBaseImageSnapshot) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomBaseImageCollectionResponseDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageCollectionResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomBaseImageCollectionResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomBaseImageCollectionResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomBaseImageCollectionResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageCollectionResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomBaseImageCollectionResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomBaseImageCollectionResponseDataInner) SetType(v string) {
	o.Type = &v
}

func (o CustomBaseImageCollectionResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomBaseImageCollectionResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCustomBaseImageCollectionResponseDataInner struct {
	value *CustomBaseImageCollectionResponseDataInner
	isSet bool
}

func (v NullableCustomBaseImageCollectionResponseDataInner) Get() *CustomBaseImageCollectionResponseDataInner {
	return v.value
}

func (v *NullableCustomBaseImageCollectionResponseDataInner) Set(val *CustomBaseImageCollectionResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomBaseImageCollectionResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomBaseImageCollectionResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomBaseImageCollectionResponseDataInner(val *CustomBaseImageCollectionResponseDataInner) *NullableCustomBaseImageCollectionResponseDataInner {
	return &NullableCustomBaseImageCollectionResponseDataInner{value: val, isSet: true}
}

func (v NullableCustomBaseImageCollectionResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomBaseImageCollectionResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

