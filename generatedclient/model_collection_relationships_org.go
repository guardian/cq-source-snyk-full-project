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

// checks if the CollectionRelationshipsOrg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionRelationshipsOrg{}

// CollectionRelationshipsOrg struct for CollectionRelationshipsOrg
type CollectionRelationshipsOrg struct {
	Data CollectionRelationshipsOrgData `json:"data"`
}

type _CollectionRelationshipsOrg CollectionRelationshipsOrg

// NewCollectionRelationshipsOrg instantiates a new CollectionRelationshipsOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionRelationshipsOrg(data CollectionRelationshipsOrgData) *CollectionRelationshipsOrg {
	this := CollectionRelationshipsOrg{}
	this.Data = data
	return &this
}

// NewCollectionRelationshipsOrgWithDefaults instantiates a new CollectionRelationshipsOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionRelationshipsOrgWithDefaults() *CollectionRelationshipsOrg {
	this := CollectionRelationshipsOrg{}
	return &this
}

// GetData returns the Data field value
func (o *CollectionRelationshipsOrg) GetData() CollectionRelationshipsOrgData {
	if o == nil {
		var ret CollectionRelationshipsOrgData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CollectionRelationshipsOrg) GetDataOk() (*CollectionRelationshipsOrgData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CollectionRelationshipsOrg) SetData(v CollectionRelationshipsOrgData) {
	o.Data = v
}

func (o CollectionRelationshipsOrg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionRelationshipsOrg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CollectionRelationshipsOrg) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCollectionRelationshipsOrg := _CollectionRelationshipsOrg{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionRelationshipsOrg)

	if err != nil {
		return err
	}

	*o = CollectionRelationshipsOrg(varCollectionRelationshipsOrg)

	return err
}

type NullableCollectionRelationshipsOrg struct {
	value *CollectionRelationshipsOrg
	isSet bool
}

func (v NullableCollectionRelationshipsOrg) Get() *CollectionRelationshipsOrg {
	return v.value
}

func (v *NullableCollectionRelationshipsOrg) Set(val *CollectionRelationshipsOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionRelationshipsOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionRelationshipsOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionRelationshipsOrg(val *CollectionRelationshipsOrg) *NullableCollectionRelationshipsOrg {
	return &NullableCollectionRelationshipsOrg{value: val, isSet: true}
}

func (v NullableCollectionRelationshipsOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionRelationshipsOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


