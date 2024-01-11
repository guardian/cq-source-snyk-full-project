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

// checks if the AppBotRelationshipsApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppBotRelationshipsApp{}

// AppBotRelationshipsApp struct for AppBotRelationshipsApp
type AppBotRelationshipsApp struct {
	Data *PublicApp `json:"data,omitempty"`
}

// NewAppBotRelationshipsApp instantiates a new AppBotRelationshipsApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppBotRelationshipsApp() *AppBotRelationshipsApp {
	this := AppBotRelationshipsApp{}
	return &this
}

// NewAppBotRelationshipsAppWithDefaults instantiates a new AppBotRelationshipsApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppBotRelationshipsAppWithDefaults() *AppBotRelationshipsApp {
	this := AppBotRelationshipsApp{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppBotRelationshipsApp) GetData() PublicApp {
	if o == nil || IsNil(o.Data) {
		var ret PublicApp
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBotRelationshipsApp) GetDataOk() (*PublicApp, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppBotRelationshipsApp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PublicApp and assigns it to the Data field.
func (o *AppBotRelationshipsApp) SetData(v PublicApp) {
	o.Data = &v
}

func (o AppBotRelationshipsApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppBotRelationshipsApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppBotRelationshipsApp struct {
	value *AppBotRelationshipsApp
	isSet bool
}

func (v NullableAppBotRelationshipsApp) Get() *AppBotRelationshipsApp {
	return v.value
}

func (v *NullableAppBotRelationshipsApp) Set(val *AppBotRelationshipsApp) {
	v.value = val
	v.isSet = true
}

func (v NullableAppBotRelationshipsApp) IsSet() bool {
	return v.isSet
}

func (v *NullableAppBotRelationshipsApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppBotRelationshipsApp(val *AppBotRelationshipsApp) *NullableAppBotRelationshipsApp {
	return &NullableAppBotRelationshipsApp{value: val, isSet: true}
}

func (v NullableAppBotRelationshipsApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppBotRelationshipsApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


