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

// checks if the CollectionRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionRelationships{}

// CollectionRelationships struct for CollectionRelationships
type CollectionRelationships struct {
	CreatedByUser CollectionRelationshipsCreatedByUser `json:"created_by_user"`
	Org CollectionRelationshipsOrg `json:"org"`
}

type _CollectionRelationships CollectionRelationships

// NewCollectionRelationships instantiates a new CollectionRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionRelationships(createdByUser CollectionRelationshipsCreatedByUser, org CollectionRelationshipsOrg) *CollectionRelationships {
	this := CollectionRelationships{}
	this.CreatedByUser = createdByUser
	this.Org = org
	return &this
}

// NewCollectionRelationshipsWithDefaults instantiates a new CollectionRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionRelationshipsWithDefaults() *CollectionRelationships {
	this := CollectionRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value
func (o *CollectionRelationships) GetCreatedByUser() CollectionRelationshipsCreatedByUser {
	if o == nil {
		var ret CollectionRelationshipsCreatedByUser
		return ret
	}

	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value
// and a boolean to check if the value has been set.
func (o *CollectionRelationships) GetCreatedByUserOk() (*CollectionRelationshipsCreatedByUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// SetCreatedByUser sets field value
func (o *CollectionRelationships) SetCreatedByUser(v CollectionRelationshipsCreatedByUser) {
	o.CreatedByUser = v
}

// GetOrg returns the Org field value
func (o *CollectionRelationships) GetOrg() CollectionRelationshipsOrg {
	if o == nil {
		var ret CollectionRelationshipsOrg
		return ret
	}

	return o.Org
}

// GetOrgOk returns a tuple with the Org field value
// and a boolean to check if the value has been set.
func (o *CollectionRelationships) GetOrgOk() (*CollectionRelationshipsOrg, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Org, true
}

// SetOrg sets field value
func (o *CollectionRelationships) SetOrg(v CollectionRelationshipsOrg) {
	o.Org = v
}

func (o CollectionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_by_user"] = o.CreatedByUser
	toSerialize["org"] = o.Org
	return toSerialize, nil
}

func (o *CollectionRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_by_user",
		"org",
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

	varCollectionRelationships := _CollectionRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionRelationships)

	if err != nil {
		return err
	}

	*o = CollectionRelationships(varCollectionRelationships)

	return err
}

type NullableCollectionRelationships struct {
	value *CollectionRelationships
	isSet bool
}

func (v NullableCollectionRelationships) Get() *CollectionRelationships {
	return v.value
}

func (v *NullableCollectionRelationships) Set(val *CollectionRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionRelationships(val *CollectionRelationships) *NullableCollectionRelationships {
	return &NullableCollectionRelationships{value: val, isSet: true}
}

func (v NullableCollectionRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


