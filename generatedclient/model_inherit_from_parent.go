/*
Snyk API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: REST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// InheritFromParent Which parent to inherit settings from.
type InheritFromParent string

// List of InheritFromParent
const (
	GROUP InheritFromParent = "group"
)

// All allowed values of InheritFromParent enum
var AllowedInheritFromParentEnumValues = []InheritFromParent{
	"group",
}

func (v *InheritFromParent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InheritFromParent(value)
	for _, existing := range AllowedInheritFromParentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InheritFromParent", value)
}

// NewInheritFromParentFromValue returns a pointer to a valid InheritFromParent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInheritFromParentFromValue(v string) (*InheritFromParent, error) {
	ev := InheritFromParent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InheritFromParent: valid values are %v", v, AllowedInheritFromParentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InheritFromParent) IsValid() bool {
	for _, existing := range AllowedInheritFromParentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InheritFromParent value
func (v InheritFromParent) Ptr() *InheritFromParent {
	return &v
}

type NullableInheritFromParent struct {
	value *InheritFromParent
	isSet bool
}

func (v NullableInheritFromParent) Get() *InheritFromParent {
	return v.value
}

func (v *NullableInheritFromParent) Set(val *InheritFromParent) {
	v.value = val
	v.isSet = true
}

func (v NullableInheritFromParent) IsSet() bool {
	return v.isSet
}

func (v *NullableInheritFromParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInheritFromParent(val *InheritFromParent) *NullableInheritFromParent {
	return &NullableInheritFromParent{value: val, isSet: true}
}

func (v NullableInheritFromParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInheritFromParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

