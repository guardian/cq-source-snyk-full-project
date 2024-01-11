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

// checks if the GetContainerImage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContainerImage200Response{}

// GetContainerImage200Response struct for GetContainerImage200Response
type GetContainerImage200Response struct {
	Data *Image `json:"data,omitempty"`
	Jsonapi *JsonApi `json:"jsonapi,omitempty"`
	Links *SelfLink `json:"links,omitempty"`
}

// NewGetContainerImage200Response instantiates a new GetContainerImage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContainerImage200Response() *GetContainerImage200Response {
	this := GetContainerImage200Response{}
	return &this
}

// NewGetContainerImage200ResponseWithDefaults instantiates a new GetContainerImage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContainerImage200ResponseWithDefaults() *GetContainerImage200Response {
	this := GetContainerImage200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetContainerImage200Response) GetData() Image {
	if o == nil || IsNil(o.Data) {
		var ret Image
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerImage200Response) GetDataOk() (*Image, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetContainerImage200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Image and assigns it to the Data field.
func (o *GetContainerImage200Response) SetData(v Image) {
	o.Data = &v
}

// GetJsonapi returns the Jsonapi field value if set, zero value otherwise.
func (o *GetContainerImage200Response) GetJsonapi() JsonApi {
	if o == nil || IsNil(o.Jsonapi) {
		var ret JsonApi
		return ret
	}
	return *o.Jsonapi
}

// GetJsonapiOk returns a tuple with the Jsonapi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerImage200Response) GetJsonapiOk() (*JsonApi, bool) {
	if o == nil || IsNil(o.Jsonapi) {
		return nil, false
	}
	return o.Jsonapi, true
}

// HasJsonapi returns a boolean if a field has been set.
func (o *GetContainerImage200Response) HasJsonapi() bool {
	if o != nil && !IsNil(o.Jsonapi) {
		return true
	}

	return false
}

// SetJsonapi gets a reference to the given JsonApi and assigns it to the Jsonapi field.
func (o *GetContainerImage200Response) SetJsonapi(v JsonApi) {
	o.Jsonapi = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetContainerImage200Response) GetLinks() SelfLink {
	if o == nil || IsNil(o.Links) {
		var ret SelfLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerImage200Response) GetLinksOk() (*SelfLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetContainerImage200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLink and assigns it to the Links field.
func (o *GetContainerImage200Response) SetLinks(v SelfLink) {
	o.Links = &v
}

func (o GetContainerImage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContainerImage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Jsonapi) {
		toSerialize["jsonapi"] = o.Jsonapi
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetContainerImage200Response struct {
	value *GetContainerImage200Response
	isSet bool
}

func (v NullableGetContainerImage200Response) Get() *GetContainerImage200Response {
	return v.value
}

func (v *NullableGetContainerImage200Response) Set(val *GetContainerImage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContainerImage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContainerImage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContainerImage200Response(val *GetContainerImage200Response) *NullableGetContainerImage200Response {
	return &NullableGetContainerImage200Response{value: val, isSet: true}
}

func (v NullableGetContainerImage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContainerImage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

