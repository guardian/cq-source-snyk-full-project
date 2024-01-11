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

// checks if the OrgWithRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgWithRelationships{}

// OrgWithRelationships struct for OrgWithRelationships
type OrgWithRelationships struct {
	Attributes OrgAttributes `json:"attributes"`
	// The Snyk ID of the organization.
	Id string `json:"id"`
	Relationships *OrgRelationships `json:"relationships,omitempty"`
	Type string `json:"type"`
}

type _OrgWithRelationships OrgWithRelationships

// NewOrgWithRelationships instantiates a new OrgWithRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgWithRelationships(attributes OrgAttributes, id string, type_ string) *OrgWithRelationships {
	this := OrgWithRelationships{}
	this.Attributes = attributes
	this.Id = id
	this.Type = type_
	return &this
}

// NewOrgWithRelationshipsWithDefaults instantiates a new OrgWithRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgWithRelationshipsWithDefaults() *OrgWithRelationships {
	this := OrgWithRelationships{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *OrgWithRelationships) GetAttributes() OrgAttributes {
	if o == nil {
		var ret OrgAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrgWithRelationships) GetAttributesOk() (*OrgAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrgWithRelationships) SetAttributes(v OrgAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value
func (o *OrgWithRelationships) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrgWithRelationships) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrgWithRelationships) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrgWithRelationships) GetRelationships() OrgRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret OrgRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgWithRelationships) GetRelationshipsOk() (*OrgRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrgWithRelationships) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrgRelationships and assigns it to the Relationships field.
func (o *OrgWithRelationships) SetRelationships(v OrgRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value
func (o *OrgWithRelationships) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgWithRelationships) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrgWithRelationships) SetType(v string) {
	o.Type = v
}

func (o OrgWithRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgWithRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *OrgWithRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributes",
		"id",
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

	varOrgWithRelationships := _OrgWithRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrgWithRelationships)

	if err != nil {
		return err
	}

	*o = OrgWithRelationships(varOrgWithRelationships)

	return err
}

type NullableOrgWithRelationships struct {
	value *OrgWithRelationships
	isSet bool
}

func (v NullableOrgWithRelationships) Get() *OrgWithRelationships {
	return v.value
}

func (v *NullableOrgWithRelationships) Set(val *OrgWithRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgWithRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgWithRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgWithRelationships(val *OrgWithRelationships) *NullableOrgWithRelationships {
	return &NullableOrgWithRelationships{value: val, isSet: true}
}

func (v NullableOrgWithRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgWithRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


