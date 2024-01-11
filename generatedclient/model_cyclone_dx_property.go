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

// checks if the CycloneDxProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CycloneDxProperty{}

// CycloneDxProperty struct for CycloneDxProperty
type CycloneDxProperty struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewCycloneDxProperty instantiates a new CycloneDxProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCycloneDxProperty() *CycloneDxProperty {
	this := CycloneDxProperty{}
	return &this
}

// NewCycloneDxPropertyWithDefaults instantiates a new CycloneDxProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCycloneDxPropertyWithDefaults() *CycloneDxProperty {
	this := CycloneDxProperty{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CycloneDxProperty) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CycloneDxProperty) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CycloneDxProperty) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CycloneDxProperty) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CycloneDxProperty) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CycloneDxProperty) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CycloneDxProperty) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CycloneDxProperty) SetValue(v string) {
	o.Value = &v
}

func (o CycloneDxProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CycloneDxProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCycloneDxProperty struct {
	value *CycloneDxProperty
	isSet bool
}

func (v NullableCycloneDxProperty) Get() *CycloneDxProperty {
	return v.value
}

func (v *NullableCycloneDxProperty) Set(val *CycloneDxProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCycloneDxProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCycloneDxProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCycloneDxProperty(val *CycloneDxProperty) *NullableCycloneDxProperty {
	return &NullableCycloneDxProperty{value: val, isSet: true}
}

func (v NullableCycloneDxProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCycloneDxProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

