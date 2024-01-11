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

// checks if the ListContainerImage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListContainerImage200Response{}

// ListContainerImage200Response struct for ListContainerImage200Response
type ListContainerImage200Response struct {
	Data []Image `json:"data,omitempty"`
	Jsonapi *JsonApi `json:"jsonapi,omitempty"`
	Links *PaginatedLinks `json:"links,omitempty"`
}

// NewListContainerImage200Response instantiates a new ListContainerImage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListContainerImage200Response() *ListContainerImage200Response {
	this := ListContainerImage200Response{}
	return &this
}

// NewListContainerImage200ResponseWithDefaults instantiates a new ListContainerImage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListContainerImage200ResponseWithDefaults() *ListContainerImage200Response {
	this := ListContainerImage200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListContainerImage200Response) GetData() []Image {
	if o == nil || IsNil(o.Data) {
		var ret []Image
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainerImage200Response) GetDataOk() ([]Image, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListContainerImage200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Image and assigns it to the Data field.
func (o *ListContainerImage200Response) SetData(v []Image) {
	o.Data = v
}

// GetJsonapi returns the Jsonapi field value if set, zero value otherwise.
func (o *ListContainerImage200Response) GetJsonapi() JsonApi {
	if o == nil || IsNil(o.Jsonapi) {
		var ret JsonApi
		return ret
	}
	return *o.Jsonapi
}

// GetJsonapiOk returns a tuple with the Jsonapi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainerImage200Response) GetJsonapiOk() (*JsonApi, bool) {
	if o == nil || IsNil(o.Jsonapi) {
		return nil, false
	}
	return o.Jsonapi, true
}

// HasJsonapi returns a boolean if a field has been set.
func (o *ListContainerImage200Response) HasJsonapi() bool {
	if o != nil && !IsNil(o.Jsonapi) {
		return true
	}

	return false
}

// SetJsonapi gets a reference to the given JsonApi and assigns it to the Jsonapi field.
func (o *ListContainerImage200Response) SetJsonapi(v JsonApi) {
	o.Jsonapi = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListContainerImage200Response) GetLinks() PaginatedLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListContainerImage200Response) GetLinksOk() (*PaginatedLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListContainerImage200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedLinks and assigns it to the Links field.
func (o *ListContainerImage200Response) SetLinks(v PaginatedLinks) {
	o.Links = &v
}

func (o ListContainerImage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListContainerImage200Response) ToMap() (map[string]interface{}, error) {
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

type NullableListContainerImage200Response struct {
	value *ListContainerImage200Response
	isSet bool
}

func (v NullableListContainerImage200Response) Get() *ListContainerImage200Response {
	return v.value
}

func (v *NullableListContainerImage200Response) Set(val *ListContainerImage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListContainerImage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListContainerImage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListContainerImage200Response(val *ListContainerImage200Response) *NullableListContainerImage200Response {
	return &NullableListContainerImage200Response{value: val, isSet: true}
}

func (v NullableListContainerImage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListContainerImage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


