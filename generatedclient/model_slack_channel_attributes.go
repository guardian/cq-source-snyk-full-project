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

// checks if the SlackChannelAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackChannelAttributes{}

// SlackChannelAttributes struct for SlackChannelAttributes
type SlackChannelAttributes struct {
	// Name of the Slack Channel
	Name *string `json:"name,omitempty"`
	// Channel type
	Type *string `json:"type,omitempty"`
}

// NewSlackChannelAttributes instantiates a new SlackChannelAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackChannelAttributes() *SlackChannelAttributes {
	this := SlackChannelAttributes{}
	return &this
}

// NewSlackChannelAttributesWithDefaults instantiates a new SlackChannelAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackChannelAttributesWithDefaults() *SlackChannelAttributes {
	this := SlackChannelAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SlackChannelAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackChannelAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SlackChannelAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SlackChannelAttributes) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SlackChannelAttributes) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackChannelAttributes) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SlackChannelAttributes) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SlackChannelAttributes) SetType(v string) {
	o.Type = &v
}

func (o SlackChannelAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackChannelAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSlackChannelAttributes struct {
	value *SlackChannelAttributes
	isSet bool
}

func (v NullableSlackChannelAttributes) Get() *SlackChannelAttributes {
	return v.value
}

func (v *NullableSlackChannelAttributes) Set(val *SlackChannelAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackChannelAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackChannelAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackChannelAttributes(val *SlackChannelAttributes) *NullableSlackChannelAttributes {
	return &NullableSlackChannelAttributes{value: val, isSet: true}
}

func (v NullableSlackChannelAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackChannelAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

