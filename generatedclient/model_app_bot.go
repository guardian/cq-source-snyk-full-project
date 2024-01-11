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

// checks if the AppBot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppBot{}

// AppBot struct for AppBot
type AppBot struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id"`
	Links *Links `json:"links,omitempty"`
	Relationships AppBotRelationships `json:"relationships"`
	Type string `json:"type"`
}

type _AppBot AppBot

// NewAppBot instantiates a new AppBot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppBot(id string, relationships AppBotRelationships, type_ string) *AppBot {
	this := AppBot{}
	this.Id = id
	this.Relationships = relationships
	this.Type = type_
	return &this
}

// NewAppBotWithDefaults instantiates a new AppBot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppBotWithDefaults() *AppBot {
	this := AppBot{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppBot) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBot) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppBot) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *AppBot) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetId returns the Id field value
func (o *AppBot) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppBot) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppBot) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppBot) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBot) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppBot) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *AppBot) SetLinks(v Links) {
	o.Links = &v
}

// GetRelationships returns the Relationships field value
func (o *AppBot) GetRelationships() AppBotRelationships {
	if o == nil {
		var ret AppBotRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppBot) GetRelationshipsOk() (*AppBotRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppBot) SetRelationships(v AppBotRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value
func (o *AppBot) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppBot) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppBot) SetType(v string) {
	o.Type = v
}

func (o AppBot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppBot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *AppBot) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"relationships",
		"type",
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

	varAppBot := _AppBot{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppBot)

	if err != nil {
		return err
	}

	*o = AppBot(varAppBot)

	return err
}

type NullableAppBot struct {
	value *AppBot
	isSet bool
}

func (v NullableAppBot) Get() *AppBot {
	return v.value
}

func (v *NullableAppBot) Set(val *AppBot) {
	v.value = val
	v.isSet = true
}

func (v NullableAppBot) IsSet() bool {
	return v.isSet
}

func (v *NullableAppBot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppBot(val *AppBot) *NullableAppBot {
	return &NullableAppBot{value: val, isSet: true}
}

func (v NullableAppBot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppBot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


