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

// checks if the CollectionMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionMeta{}

// CollectionMeta struct for CollectionMeta
type CollectionMeta struct {
	// The sum of critical severity issues of the collection's projects
	IssuesCriticalCount float32 `json:"issues_critical_count"`
	// The sum of high severity issues of the collection's projects
	IssuesHighCount float32 `json:"issues_high_count"`
	// The sum of low severity issues of the collection's projects
	IssuesLowCount float32 `json:"issues_low_count"`
	// The sum of medium severity issues of the collection's projects
	IssuesMediumCount float32 `json:"issues_medium_count"`
	// The amount of projects belonging to this collection
	ProjectsCount float32 `json:"projects_count"`
}

type _CollectionMeta CollectionMeta

// NewCollectionMeta instantiates a new CollectionMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionMeta(issuesCriticalCount float32, issuesHighCount float32, issuesLowCount float32, issuesMediumCount float32, projectsCount float32) *CollectionMeta {
	this := CollectionMeta{}
	this.IssuesCriticalCount = issuesCriticalCount
	this.IssuesHighCount = issuesHighCount
	this.IssuesLowCount = issuesLowCount
	this.IssuesMediumCount = issuesMediumCount
	this.ProjectsCount = projectsCount
	return &this
}

// NewCollectionMetaWithDefaults instantiates a new CollectionMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionMetaWithDefaults() *CollectionMeta {
	this := CollectionMeta{}
	return &this
}

// GetIssuesCriticalCount returns the IssuesCriticalCount field value
func (o *CollectionMeta) GetIssuesCriticalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IssuesCriticalCount
}

// GetIssuesCriticalCountOk returns a tuple with the IssuesCriticalCount field value
// and a boolean to check if the value has been set.
func (o *CollectionMeta) GetIssuesCriticalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuesCriticalCount, true
}

// SetIssuesCriticalCount sets field value
func (o *CollectionMeta) SetIssuesCriticalCount(v float32) {
	o.IssuesCriticalCount = v
}

// GetIssuesHighCount returns the IssuesHighCount field value
func (o *CollectionMeta) GetIssuesHighCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IssuesHighCount
}

// GetIssuesHighCountOk returns a tuple with the IssuesHighCount field value
// and a boolean to check if the value has been set.
func (o *CollectionMeta) GetIssuesHighCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuesHighCount, true
}

// SetIssuesHighCount sets field value
func (o *CollectionMeta) SetIssuesHighCount(v float32) {
	o.IssuesHighCount = v
}

// GetIssuesLowCount returns the IssuesLowCount field value
func (o *CollectionMeta) GetIssuesLowCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IssuesLowCount
}

// GetIssuesLowCountOk returns a tuple with the IssuesLowCount field value
// and a boolean to check if the value has been set.
func (o *CollectionMeta) GetIssuesLowCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuesLowCount, true
}

// SetIssuesLowCount sets field value
func (o *CollectionMeta) SetIssuesLowCount(v float32) {
	o.IssuesLowCount = v
}

// GetIssuesMediumCount returns the IssuesMediumCount field value
func (o *CollectionMeta) GetIssuesMediumCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IssuesMediumCount
}

// GetIssuesMediumCountOk returns a tuple with the IssuesMediumCount field value
// and a boolean to check if the value has been set.
func (o *CollectionMeta) GetIssuesMediumCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuesMediumCount, true
}

// SetIssuesMediumCount sets field value
func (o *CollectionMeta) SetIssuesMediumCount(v float32) {
	o.IssuesMediumCount = v
}

// GetProjectsCount returns the ProjectsCount field value
func (o *CollectionMeta) GetProjectsCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProjectsCount
}

// GetProjectsCountOk returns a tuple with the ProjectsCount field value
// and a boolean to check if the value has been set.
func (o *CollectionMeta) GetProjectsCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectsCount, true
}

// SetProjectsCount sets field value
func (o *CollectionMeta) SetProjectsCount(v float32) {
	o.ProjectsCount = v
}

func (o CollectionMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issues_critical_count"] = o.IssuesCriticalCount
	toSerialize["issues_high_count"] = o.IssuesHighCount
	toSerialize["issues_low_count"] = o.IssuesLowCount
	toSerialize["issues_medium_count"] = o.IssuesMediumCount
	toSerialize["projects_count"] = o.ProjectsCount
	return toSerialize, nil
}

func (o *CollectionMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"issues_critical_count",
		"issues_high_count",
		"issues_low_count",
		"issues_medium_count",
		"projects_count",
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

	varCollectionMeta := _CollectionMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionMeta)

	if err != nil {
		return err
	}

	*o = CollectionMeta(varCollectionMeta)

	return err
}

type NullableCollectionMeta struct {
	value *CollectionMeta
	isSet bool
}

func (v NullableCollectionMeta) Get() *CollectionMeta {
	return v.value
}

func (v *NullableCollectionMeta) Set(val *CollectionMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionMeta(val *CollectionMeta) *NullableCollectionMeta {
	return &NullableCollectionMeta{value: val, isSet: true}
}

func (v NullableCollectionMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

