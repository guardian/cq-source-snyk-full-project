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

// checks if the CreateOrgInvitationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrgInvitationRequest{}

// CreateOrgInvitationRequest struct for CreateOrgInvitationRequest
type CreateOrgInvitationRequest struct {
	Data OrgInvitationPostData `json:"data"`
}

type _CreateOrgInvitationRequest CreateOrgInvitationRequest

// NewCreateOrgInvitationRequest instantiates a new CreateOrgInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrgInvitationRequest(data OrgInvitationPostData) *CreateOrgInvitationRequest {
	this := CreateOrgInvitationRequest{}
	this.Data = data
	return &this
}

// NewCreateOrgInvitationRequestWithDefaults instantiates a new CreateOrgInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrgInvitationRequestWithDefaults() *CreateOrgInvitationRequest {
	this := CreateOrgInvitationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *CreateOrgInvitationRequest) GetData() OrgInvitationPostData {
	if o == nil {
		var ret OrgInvitationPostData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateOrgInvitationRequest) GetDataOk() (*OrgInvitationPostData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateOrgInvitationRequest) SetData(v OrgInvitationPostData) {
	o.Data = v
}

func (o CreateOrgInvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrgInvitationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreateOrgInvitationRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCreateOrgInvitationRequest := _CreateOrgInvitationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrgInvitationRequest)

	if err != nil {
		return err
	}

	*o = CreateOrgInvitationRequest(varCreateOrgInvitationRequest)

	return err
}

type NullableCreateOrgInvitationRequest struct {
	value *CreateOrgInvitationRequest
	isSet bool
}

func (v NullableCreateOrgInvitationRequest) Get() *CreateOrgInvitationRequest {
	return v.value
}

func (v *NullableCreateOrgInvitationRequest) Set(val *CreateOrgInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrgInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrgInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrgInvitationRequest(val *CreateOrgInvitationRequest) *NullableCreateOrgInvitationRequest {
	return &NullableCreateOrgInvitationRequest{value: val, isSet: true}
}

func (v NullableCreateOrgInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrgInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


