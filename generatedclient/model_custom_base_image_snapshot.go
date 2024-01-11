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

// checks if the CustomBaseImageSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomBaseImageSnapshot{}

// CustomBaseImageSnapshot struct for CustomBaseImageSnapshot
type CustomBaseImageSnapshot struct {
	GroupId *string `json:"group_id,omitempty"`
	IncludeInRecommendations *bool `json:"include_in_recommendations,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
	Repository *string `json:"repository,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// NewCustomBaseImageSnapshot instantiates a new CustomBaseImageSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomBaseImageSnapshot() *CustomBaseImageSnapshot {
	this := CustomBaseImageSnapshot{}
	return &this
}

// NewCustomBaseImageSnapshotWithDefaults instantiates a new CustomBaseImageSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomBaseImageSnapshotWithDefaults() *CustomBaseImageSnapshot {
	this := CustomBaseImageSnapshot{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CustomBaseImageSnapshot) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageSnapshot) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CustomBaseImageSnapshot) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CustomBaseImageSnapshot) SetGroupId(v string) {
	o.GroupId = &v
}

// GetIncludeInRecommendations returns the IncludeInRecommendations field value if set, zero value otherwise.
func (o *CustomBaseImageSnapshot) GetIncludeInRecommendations() bool {
	if o == nil || IsNil(o.IncludeInRecommendations) {
		var ret bool
		return ret
	}
	return *o.IncludeInRecommendations
}

// GetIncludeInRecommendationsOk returns a tuple with the IncludeInRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageSnapshot) GetIncludeInRecommendationsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeInRecommendations) {
		return nil, false
	}
	return o.IncludeInRecommendations, true
}

// HasIncludeInRecommendations returns a boolean if a field has been set.
func (o *CustomBaseImageSnapshot) HasIncludeInRecommendations() bool {
	if o != nil && !IsNil(o.IncludeInRecommendations) {
		return true
	}

	return false
}

// SetIncludeInRecommendations gets a reference to the given bool and assigns it to the IncludeInRecommendations field.
func (o *CustomBaseImageSnapshot) SetIncludeInRecommendations(v bool) {
	o.IncludeInRecommendations = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *CustomBaseImageSnapshot) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageSnapshot) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *CustomBaseImageSnapshot) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *CustomBaseImageSnapshot) SetOrgId(v string) {
	o.OrgId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CustomBaseImageSnapshot) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageSnapshot) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CustomBaseImageSnapshot) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CustomBaseImageSnapshot) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *CustomBaseImageSnapshot) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageSnapshot) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *CustomBaseImageSnapshot) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *CustomBaseImageSnapshot) SetRepository(v string) {
	o.Repository = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CustomBaseImageSnapshot) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageSnapshot) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CustomBaseImageSnapshot) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CustomBaseImageSnapshot) SetTag(v string) {
	o.Tag = &v
}

func (o CustomBaseImageSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomBaseImageSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.IncludeInRecommendations) {
		toSerialize["include_in_recommendations"] = o.IncludeInRecommendations
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableCustomBaseImageSnapshot struct {
	value *CustomBaseImageSnapshot
	isSet bool
}

func (v NullableCustomBaseImageSnapshot) Get() *CustomBaseImageSnapshot {
	return v.value
}

func (v *NullableCustomBaseImageSnapshot) Set(val *CustomBaseImageSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomBaseImageSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomBaseImageSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomBaseImageSnapshot(val *CustomBaseImageSnapshot) *NullableCustomBaseImageSnapshot {
	return &NullableCustomBaseImageSnapshot{value: val, isSet: true}
}

func (v NullableCustomBaseImageSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomBaseImageSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


