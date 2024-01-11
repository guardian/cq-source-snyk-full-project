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

// checks if the CustomBaseImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomBaseImageResponse{}

// CustomBaseImageResponse struct for CustomBaseImageResponse
type CustomBaseImageResponse struct {
	Data CustomBaseImageResource `json:"data"`
	Jsonapi JsonApi `json:"jsonapi"`
	Links *SelfLink `json:"links,omitempty"`
}

type _CustomBaseImageResponse CustomBaseImageResponse

// NewCustomBaseImageResponse instantiates a new CustomBaseImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomBaseImageResponse(data CustomBaseImageResource, jsonapi JsonApi) *CustomBaseImageResponse {
	this := CustomBaseImageResponse{}
	this.Data = data
	this.Jsonapi = jsonapi
	return &this
}

// NewCustomBaseImageResponseWithDefaults instantiates a new CustomBaseImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomBaseImageResponseWithDefaults() *CustomBaseImageResponse {
	this := CustomBaseImageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CustomBaseImageResponse) GetData() CustomBaseImageResource {
	if o == nil {
		var ret CustomBaseImageResource
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomBaseImageResponse) GetDataOk() (*CustomBaseImageResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomBaseImageResponse) SetData(v CustomBaseImageResource) {
	o.Data = v
}

// GetJsonapi returns the Jsonapi field value
func (o *CustomBaseImageResponse) GetJsonapi() JsonApi {
	if o == nil {
		var ret JsonApi
		return ret
	}

	return o.Jsonapi
}

// GetJsonapiOk returns a tuple with the Jsonapi field value
// and a boolean to check if the value has been set.
func (o *CustomBaseImageResponse) GetJsonapiOk() (*JsonApi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jsonapi, true
}

// SetJsonapi sets field value
func (o *CustomBaseImageResponse) SetJsonapi(v JsonApi) {
	o.Jsonapi = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomBaseImageResponse) GetLinks() SelfLink {
	if o == nil || IsNil(o.Links) {
		var ret SelfLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBaseImageResponse) GetLinksOk() (*SelfLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomBaseImageResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLink and assigns it to the Links field.
func (o *CustomBaseImageResponse) SetLinks(v SelfLink) {
	o.Links = &v
}

func (o CustomBaseImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomBaseImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["jsonapi"] = o.Jsonapi
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *CustomBaseImageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"jsonapi",
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

	varCustomBaseImageResponse := _CustomBaseImageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomBaseImageResponse)

	if err != nil {
		return err
	}

	*o = CustomBaseImageResponse(varCustomBaseImageResponse)

	return err
}

type NullableCustomBaseImageResponse struct {
	value *CustomBaseImageResponse
	isSet bool
}

func (v NullableCustomBaseImageResponse) Get() *CustomBaseImageResponse {
	return v.value
}

func (v *NullableCustomBaseImageResponse) Set(val *CustomBaseImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomBaseImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomBaseImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomBaseImageResponse(val *CustomBaseImageResponse) *NullableCustomBaseImageResponse {
	return &NullableCustomBaseImageResponse{value: val, isSet: true}
}

func (v NullableCustomBaseImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomBaseImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


