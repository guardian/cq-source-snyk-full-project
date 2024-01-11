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

// checks if the IssuesMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuesMeta{}

// IssuesMeta struct for IssuesMeta
type IssuesMeta struct {
	Package *PackageMeta `json:"package,omitempty"`
}

// NewIssuesMeta instantiates a new IssuesMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuesMeta() *IssuesMeta {
	this := IssuesMeta{}
	return &this
}

// NewIssuesMetaWithDefaults instantiates a new IssuesMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuesMetaWithDefaults() *IssuesMeta {
	this := IssuesMeta{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *IssuesMeta) GetPackage() PackageMeta {
	if o == nil || IsNil(o.Package) {
		var ret PackageMeta
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesMeta) GetPackageOk() (*PackageMeta, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *IssuesMeta) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given PackageMeta and assigns it to the Package field.
func (o *IssuesMeta) SetPackage(v PackageMeta) {
	o.Package = &v
}

func (o IssuesMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuesMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	return toSerialize, nil
}

type NullableIssuesMeta struct {
	value *IssuesMeta
	isSet bool
}

func (v NullableIssuesMeta) Get() *IssuesMeta {
	return v.value
}

func (v *NullableIssuesMeta) Set(val *IssuesMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuesMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuesMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuesMeta(val *IssuesMeta) *NullableIssuesMeta {
	return &NullableIssuesMeta{value: val, isSet: true}
}

func (v NullableIssuesMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuesMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


