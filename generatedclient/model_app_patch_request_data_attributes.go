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

// checks if the AppPatchRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPatchRequestDataAttributes{}

// AppPatchRequestDataAttributes struct for AppPatchRequestDataAttributes
type AppPatchRequestDataAttributes struct {
	// The access token time to live for your app, in seconds. It only affects the newly generated access tokens, existing access token will  continue to have their previous time to live as expiration.
	AccessTokenTtlSeconds *float32 `json:"access_token_ttl_seconds,omitempty"`
	// New name of the app to display to users during authorization flow.
	Name *string `json:"name,omitempty"`
	// List of allowed redirect URIs to call back after authentication.
	RedirectUris []string `json:"redirect_uris,omitempty"`
}

// NewAppPatchRequestDataAttributes instantiates a new AppPatchRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPatchRequestDataAttributes() *AppPatchRequestDataAttributes {
	this := AppPatchRequestDataAttributes{}
	return &this
}

// NewAppPatchRequestDataAttributesWithDefaults instantiates a new AppPatchRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPatchRequestDataAttributesWithDefaults() *AppPatchRequestDataAttributes {
	this := AppPatchRequestDataAttributes{}
	return &this
}

// GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field value if set, zero value otherwise.
func (o *AppPatchRequestDataAttributes) GetAccessTokenTtlSeconds() float32 {
	if o == nil || IsNil(o.AccessTokenTtlSeconds) {
		var ret float32
		return ret
	}
	return *o.AccessTokenTtlSeconds
}

// GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPatchRequestDataAttributes) GetAccessTokenTtlSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.AccessTokenTtlSeconds) {
		return nil, false
	}
	return o.AccessTokenTtlSeconds, true
}

// HasAccessTokenTtlSeconds returns a boolean if a field has been set.
func (o *AppPatchRequestDataAttributes) HasAccessTokenTtlSeconds() bool {
	if o != nil && !IsNil(o.AccessTokenTtlSeconds) {
		return true
	}

	return false
}

// SetAccessTokenTtlSeconds gets a reference to the given float32 and assigns it to the AccessTokenTtlSeconds field.
func (o *AppPatchRequestDataAttributes) SetAccessTokenTtlSeconds(v float32) {
	o.AccessTokenTtlSeconds = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppPatchRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPatchRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppPatchRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppPatchRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *AppPatchRequestDataAttributes) GetRedirectUris() []string {
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPatchRequestDataAttributes) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *AppPatchRequestDataAttributes) HasRedirectUris() bool {
	if o != nil && !IsNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *AppPatchRequestDataAttributes) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

func (o AppPatchRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPatchRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessTokenTtlSeconds) {
		toSerialize["access_token_ttl_seconds"] = o.AccessTokenTtlSeconds
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	return toSerialize, nil
}

type NullableAppPatchRequestDataAttributes struct {
	value *AppPatchRequestDataAttributes
	isSet bool
}

func (v NullableAppPatchRequestDataAttributes) Get() *AppPatchRequestDataAttributes {
	return v.value
}

func (v *NullableAppPatchRequestDataAttributes) Set(val *AppPatchRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPatchRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPatchRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPatchRequestDataAttributes(val *AppPatchRequestDataAttributes) *NullableAppPatchRequestDataAttributes {
	return &NullableAppPatchRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppPatchRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPatchRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

