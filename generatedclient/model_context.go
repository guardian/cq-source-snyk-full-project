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

// Context Allow installing the app to a org/group or to a user, default tenant.
type Context string

// List of Context
const (
	TENANT Context = "tenant"
	USER Context = "user"
)

// All allowed values of Context enum
var AllowedContextEnumValues = []Context{
	"tenant",
	"user",
}

func (v *Context) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Context(value)
	for _, existing := range AllowedContextEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Context", value)
}

// NewContextFromValue returns a pointer to a valid Context
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContextFromValue(v string) (*Context, error) {
	ev := Context(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Context: valid values are %v", v, AllowedContextEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Context) IsValid() bool {
	for _, existing := range AllowedContextEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Context value
func (v Context) Ptr() *Context {
	return &v
}

type NullableContext struct {
	value *Context
	isSet bool
}

func (v NullableContext) Get() *Context {
	return v.value
}

func (v *NullableContext) Set(val *Context) {
	v.value = val
	v.isSet = true
}

func (v NullableContext) IsSet() bool {
	return v.isSet
}

func (v *NullableContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContext(val *Context) *NullableContext {
	return &NullableContext{value: val, isSet: true}
}

func (v NullableContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
