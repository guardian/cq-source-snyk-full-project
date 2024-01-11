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

// checks if the ErrorLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLink{}

// ErrorLink A link that leads to further details about this particular occurrance of the problem.
type ErrorLink struct {
	About *LinkProperty `json:"about,omitempty"`
}

// NewErrorLink instantiates a new ErrorLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLink() *ErrorLink {
	this := ErrorLink{}
	return &this
}

// NewErrorLinkWithDefaults instantiates a new ErrorLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLinkWithDefaults() *ErrorLink {
	this := ErrorLink{}
	return &this
}

// GetAbout returns the About field value if set, zero value otherwise.
func (o *ErrorLink) GetAbout() LinkProperty {
	if o == nil || IsNil(o.About) {
		var ret LinkProperty
		return ret
	}
	return *o.About
}

// GetAboutOk returns a tuple with the About field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLink) GetAboutOk() (*LinkProperty, bool) {
	if o == nil || IsNil(o.About) {
		return nil, false
	}
	return o.About, true
}

// HasAbout returns a boolean if a field has been set.
func (o *ErrorLink) HasAbout() bool {
	if o != nil && !IsNil(o.About) {
		return true
	}

	return false
}

// SetAbout gets a reference to the given LinkProperty and assigns it to the About field.
func (o *ErrorLink) SetAbout(v LinkProperty) {
	o.About = &v
}

func (o ErrorLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.About) {
		toSerialize["about"] = o.About
	}
	return toSerialize, nil
}

type NullableErrorLink struct {
	value *ErrorLink
	isSet bool
}

func (v NullableErrorLink) Get() *ErrorLink {
	return v.value
}

func (v *NullableErrorLink) Set(val *ErrorLink) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLink) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLink(val *ErrorLink) *NullableErrorLink {
	return &NullableErrorLink{value: val, isSet: true}
}

func (v NullableErrorLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

