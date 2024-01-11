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

// checks if the GetManyGroupServiceAccount200ResponseDataInnerAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManyGroupServiceAccount200ResponseDataInnerAttributes{}

// GetManyGroupServiceAccount200ResponseDataInnerAttributes struct for GetManyGroupServiceAccount200ResponseDataInnerAttributes
type GetManyGroupServiceAccount200ResponseDataInnerAttributes struct {
	// The time, in seconds, that a generated access token will be valid for. Defaults to 1hr if unset. Only provided when auth_type is oauth_private_key_jwt.
	AccessTokenTtlSeconds *float32 `json:"access_token_ttl_seconds,omitempty"`
	// The Snyk API Key for this service account. Only returned on creation, and only when auth_type is api_key.
	ApiKey *string `json:"api_key,omitempty"`
	// The authentication strategy for the service account:   * api_key - Regular Snyk API Key.   * oauth_client_secret - OAuth2 client_credentials grant, which returns a client secret that can be used to retrieve an access token.   * oauth_private_key_jwt - OAuth2 client_credentials grant, using private_key_jwt client_assertion as laid out OIDC Connect Core 1.0, section 9.
	AuthType string `json:"auth_type"`
	// The service account's attached client-id. Used to request an access-token. Only provided when auth_type is oauth_private_key_jwt.
	ClientId *string `json:"client_id,omitempty"`
	// The client secret used for obtaining access tokens. Only sent on creation of new service accounts and cannot be retrieved thereafter. Only provided when auth_type is oauth_client_secret.
	ClientSecret *string `json:"client_secret,omitempty"`
	// A JWKs URL used to verify signed JWT requests against. Must be https. Only provided when auth_type is oauth_private_key_jwt.
	JwksUrl *string `json:"jwks_url,omitempty"`
	// The level of access for the service account:   * Group - the service account was created at the Group level.   * Org - the service account was created at the Org level.
	Level *string `json:"level,omitempty"`
	// A human-friendly name of the service account.
	Name string `json:"name"`
	// The ID of the role which the Service Account is associated with.
	RoleId string `json:"role_id"`
}

type _GetManyGroupServiceAccount200ResponseDataInnerAttributes GetManyGroupServiceAccount200ResponseDataInnerAttributes

// NewGetManyGroupServiceAccount200ResponseDataInnerAttributes instantiates a new GetManyGroupServiceAccount200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManyGroupServiceAccount200ResponseDataInnerAttributes(authType string, name string, roleId string) *GetManyGroupServiceAccount200ResponseDataInnerAttributes {
	this := GetManyGroupServiceAccount200ResponseDataInnerAttributes{}
	this.AuthType = authType
	this.Name = name
	this.RoleId = roleId
	return &this
}

// NewGetManyGroupServiceAccount200ResponseDataInnerAttributesWithDefaults instantiates a new GetManyGroupServiceAccount200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManyGroupServiceAccount200ResponseDataInnerAttributesWithDefaults() *GetManyGroupServiceAccount200ResponseDataInnerAttributes {
	this := GetManyGroupServiceAccount200ResponseDataInnerAttributes{}
	return &this
}

// GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field value if set, zero value otherwise.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetAccessTokenTtlSeconds() float32 {
	if o == nil || IsNil(o.AccessTokenTtlSeconds) {
		var ret float32
		return ret
	}
	return *o.AccessTokenTtlSeconds
}

// GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetAccessTokenTtlSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.AccessTokenTtlSeconds) {
		return nil, false
	}
	return o.AccessTokenTtlSeconds, true
}

// HasAccessTokenTtlSeconds returns a boolean if a field has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasAccessTokenTtlSeconds() bool {
	if o != nil && !IsNil(o.AccessTokenTtlSeconds) {
		return true
	}

	return false
}

// SetAccessTokenTtlSeconds gets a reference to the given float32 and assigns it to the AccessTokenTtlSeconds field.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetAccessTokenTtlSeconds(v float32) {
	o.AccessTokenTtlSeconds = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAuthType returns the AuthType field value
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetAuthType(v string) {
	o.AuthType = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetJwksUrl returns the JwksUrl field value if set, zero value otherwise.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetJwksUrl() string {
	if o == nil || IsNil(o.JwksUrl) {
		var ret string
		return ret
	}
	return *o.JwksUrl
}

// GetJwksUrlOk returns a tuple with the JwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetJwksUrlOk() (*string, bool) {
	if o == nil || IsNil(o.JwksUrl) {
		return nil, false
	}
	return o.JwksUrl, true
}

// HasJwksUrl returns a boolean if a field has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasJwksUrl() bool {
	if o != nil && !IsNil(o.JwksUrl) {
		return true
	}

	return false
}

// SetJwksUrl gets a reference to the given string and assigns it to the JwksUrl field.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetJwksUrl(v string) {
	o.JwksUrl = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetLevel(v string) {
	o.Level = &v
}

// GetName returns the Name field value
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = v
}

// GetRoleId returns the RoleId field value
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetRoleId(v string) {
	o.RoleId = v
}

func (o GetManyGroupServiceAccount200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManyGroupServiceAccount200ResponseDataInnerAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessTokenTtlSeconds) {
		toSerialize["access_token_ttl_seconds"] = o.AccessTokenTtlSeconds
	}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	toSerialize["auth_type"] = o.AuthType
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if !IsNil(o.JwksUrl) {
		toSerialize["jwks_url"] = o.JwksUrl
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	toSerialize["name"] = o.Name
	toSerialize["role_id"] = o.RoleId
	return toSerialize, nil
}

func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auth_type",
		"name",
		"role_id",
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

	varGetManyGroupServiceAccount200ResponseDataInnerAttributes := _GetManyGroupServiceAccount200ResponseDataInnerAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetManyGroupServiceAccount200ResponseDataInnerAttributes)

	if err != nil {
		return err
	}

	*o = GetManyGroupServiceAccount200ResponseDataInnerAttributes(varGetManyGroupServiceAccount200ResponseDataInnerAttributes)

	return err
}

type NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes struct {
	value *GetManyGroupServiceAccount200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes) Get() *GetManyGroupServiceAccount200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes) Set(val *GetManyGroupServiceAccount200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManyGroupServiceAccount200ResponseDataInnerAttributes(val *GetManyGroupServiceAccount200ResponseDataInnerAttributes) *NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes {
	return &NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManyGroupServiceAccount200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


