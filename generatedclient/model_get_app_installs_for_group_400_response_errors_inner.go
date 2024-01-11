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

// checks if the GetAppInstallsForGroup400ResponseErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAppInstallsForGroup400ResponseErrorsInner{}

// GetAppInstallsForGroup400ResponseErrorsInner struct for GetAppInstallsForGroup400ResponseErrorsInner
type GetAppInstallsForGroup400ResponseErrorsInner struct {
	// An application-specific error code, expressed as a string value.
	Code *string `json:"code,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`
	// A unique identifier for this particular occurrence of the problem.
	Id *string `json:"id,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Source *GetAppInstallsForGroup400ResponseErrorsInnerSource `json:"source,omitempty"`
	// The HTTP status code applicable to this problem, expressed as a string value.
	Status string `json:"status"`
	// A short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
	Title *string `json:"title,omitempty"`
}

type _GetAppInstallsForGroup400ResponseErrorsInner GetAppInstallsForGroup400ResponseErrorsInner

// NewGetAppInstallsForGroup400ResponseErrorsInner instantiates a new GetAppInstallsForGroup400ResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAppInstallsForGroup400ResponseErrorsInner(detail string, status string) *GetAppInstallsForGroup400ResponseErrorsInner {
	this := GetAppInstallsForGroup400ResponseErrorsInner{}
	this.Detail = detail
	this.Status = status
	return &this
}

// NewGetAppInstallsForGroup400ResponseErrorsInnerWithDefaults instantiates a new GetAppInstallsForGroup400ResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAppInstallsForGroup400ResponseErrorsInnerWithDefaults() *GetAppInstallsForGroup400ResponseErrorsInner {
	this := GetAppInstallsForGroup400ResponseErrorsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetCode(v string) {
	o.Code = &v
}

// GetDetail returns the Detail field value
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetDetail(v string) {
	o.Detail = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetMeta() map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetSource() GetAppInstallsForGroup400ResponseErrorsInnerSource {
	if o == nil || IsNil(o.Source) {
		var ret GetAppInstallsForGroup400ResponseErrorsInnerSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetSourceOk() (*GetAppInstallsForGroup400ResponseErrorsInnerSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given GetAppInstallsForGroup400ResponseErrorsInnerSource and assigns it to the Source field.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetSource(v GetAppInstallsForGroup400ResponseErrorsInnerSource) {
	o.Source = &v
}

// GetStatus returns the Status field value
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetStatus(v string) {
	o.Status = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetTitle(v string) {
	o.Title = &v
}

func (o GetAppInstallsForGroup400ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAppInstallsForGroup400ResponseErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

func (o *GetAppInstallsForGroup400ResponseErrorsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"detail",
		"status",
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

	varGetAppInstallsForGroup400ResponseErrorsInner := _GetAppInstallsForGroup400ResponseErrorsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAppInstallsForGroup400ResponseErrorsInner)

	if err != nil {
		return err
	}

	*o = GetAppInstallsForGroup400ResponseErrorsInner(varGetAppInstallsForGroup400ResponseErrorsInner)

	return err
}

type NullableGetAppInstallsForGroup400ResponseErrorsInner struct {
	value *GetAppInstallsForGroup400ResponseErrorsInner
	isSet bool
}

func (v NullableGetAppInstallsForGroup400ResponseErrorsInner) Get() *GetAppInstallsForGroup400ResponseErrorsInner {
	return v.value
}

func (v *NullableGetAppInstallsForGroup400ResponseErrorsInner) Set(val *GetAppInstallsForGroup400ResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAppInstallsForGroup400ResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAppInstallsForGroup400ResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAppInstallsForGroup400ResponseErrorsInner(val *GetAppInstallsForGroup400ResponseErrorsInner) *NullableGetAppInstallsForGroup400ResponseErrorsInner {
	return &NullableGetAppInstallsForGroup400ResponseErrorsInner{value: val, isSet: true}
}

func (v NullableGetAppInstallsForGroup400ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAppInstallsForGroup400ResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


