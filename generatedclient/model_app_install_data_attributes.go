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

// checks if the AppInstallDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInstallDataAttributes{}

// AppInstallDataAttributes struct for AppInstallDataAttributes
type AppInstallDataAttributes struct {
	// The oauth2 client id for the app.
	ClientId *string `json:"client_id,omitempty"`
}

// NewAppInstallDataAttributes instantiates a new AppInstallDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInstallDataAttributes() *AppInstallDataAttributes {
	this := AppInstallDataAttributes{}
	return &this
}

// NewAppInstallDataAttributesWithDefaults instantiates a new AppInstallDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInstallDataAttributesWithDefaults() *AppInstallDataAttributes {
	this := AppInstallDataAttributes{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AppInstallDataAttributes) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInstallDataAttributes) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AppInstallDataAttributes) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AppInstallDataAttributes) SetClientId(v string) {
	o.ClientId = &v
}

func (o AppInstallDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInstallDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	return toSerialize, nil
}

type NullableAppInstallDataAttributes struct {
	value *AppInstallDataAttributes
	isSet bool
}

func (v NullableAppInstallDataAttributes) Get() *AppInstallDataAttributes {
	return v.value
}

func (v *NullableAppInstallDataAttributes) Set(val *AppInstallDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInstallDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInstallDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInstallDataAttributes(val *AppInstallDataAttributes) *NullableAppInstallDataAttributes {
	return &NullableAppInstallDataAttributes{value: val, isSet: true}
}

func (v NullableAppInstallDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInstallDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


