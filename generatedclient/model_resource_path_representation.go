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

// checks if the ResourcePathRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcePathRepresentation{}

// ResourcePathRepresentation An object that contains an opaque identifying string.
type ResourcePathRepresentation struct {
	ResourcePath string `json:"resource_path"`
}

type _ResourcePathRepresentation ResourcePathRepresentation

// NewResourcePathRepresentation instantiates a new ResourcePathRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePathRepresentation(resourcePath string) *ResourcePathRepresentation {
	this := ResourcePathRepresentation{}
	this.ResourcePath = resourcePath
	return &this
}

// NewResourcePathRepresentationWithDefaults instantiates a new ResourcePathRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePathRepresentationWithDefaults() *ResourcePathRepresentation {
	this := ResourcePathRepresentation{}
	return &this
}

// GetResourcePath returns the ResourcePath field value
func (o *ResourcePathRepresentation) GetResourcePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourcePath
}

// GetResourcePathOk returns a tuple with the ResourcePath field value
// and a boolean to check if the value has been set.
func (o *ResourcePathRepresentation) GetResourcePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePath, true
}

// SetResourcePath sets field value
func (o *ResourcePathRepresentation) SetResourcePath(v string) {
	o.ResourcePath = v
}

func (o ResourcePathRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcePathRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_path"] = o.ResourcePath
	return toSerialize, nil
}

func (o *ResourcePathRepresentation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource_path",
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

	varResourcePathRepresentation := _ResourcePathRepresentation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourcePathRepresentation)

	if err != nil {
		return err
	}

	*o = ResourcePathRepresentation(varResourcePathRepresentation)

	return err
}

type NullableResourcePathRepresentation struct {
	value *ResourcePathRepresentation
	isSet bool
}

func (v NullableResourcePathRepresentation) Get() *ResourcePathRepresentation {
	return v.value
}

func (v *NullableResourcePathRepresentation) Set(val *ResourcePathRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePathRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePathRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePathRepresentation(val *ResourcePathRepresentation) *NullableResourcePathRepresentation {
	return &NullableResourcePathRepresentation{value: val, isSet: true}
}

func (v NullableResourcePathRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePathRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

