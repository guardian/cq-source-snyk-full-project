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

// checks if the CreateGroupAppInstall201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupAppInstall201Response{}

// CreateGroupAppInstall201Response struct for CreateGroupAppInstall201Response
type CreateGroupAppInstall201Response struct {
	Data AppInstallWithClient `json:"data"`
	Jsonapi JsonApi `json:"jsonapi"`
	Links PaginatedLinks `json:"links"`
}

type _CreateGroupAppInstall201Response CreateGroupAppInstall201Response

// NewCreateGroupAppInstall201Response instantiates a new CreateGroupAppInstall201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupAppInstall201Response(data AppInstallWithClient, jsonapi JsonApi, links PaginatedLinks) *CreateGroupAppInstall201Response {
	this := CreateGroupAppInstall201Response{}
	this.Data = data
	this.Jsonapi = jsonapi
	this.Links = links
	return &this
}

// NewCreateGroupAppInstall201ResponseWithDefaults instantiates a new CreateGroupAppInstall201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupAppInstall201ResponseWithDefaults() *CreateGroupAppInstall201Response {
	this := CreateGroupAppInstall201Response{}
	return &this
}

// GetData returns the Data field value
func (o *CreateGroupAppInstall201Response) GetData() AppInstallWithClient {
	if o == nil {
		var ret AppInstallWithClient
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateGroupAppInstall201Response) GetDataOk() (*AppInstallWithClient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateGroupAppInstall201Response) SetData(v AppInstallWithClient) {
	o.Data = v
}

// GetJsonapi returns the Jsonapi field value
func (o *CreateGroupAppInstall201Response) GetJsonapi() JsonApi {
	if o == nil {
		var ret JsonApi
		return ret
	}

	return o.Jsonapi
}

// GetJsonapiOk returns a tuple with the Jsonapi field value
// and a boolean to check if the value has been set.
func (o *CreateGroupAppInstall201Response) GetJsonapiOk() (*JsonApi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jsonapi, true
}

// SetJsonapi sets field value
func (o *CreateGroupAppInstall201Response) SetJsonapi(v JsonApi) {
	o.Jsonapi = v
}

// GetLinks returns the Links field value
func (o *CreateGroupAppInstall201Response) GetLinks() PaginatedLinks {
	if o == nil {
		var ret PaginatedLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CreateGroupAppInstall201Response) GetLinksOk() (*PaginatedLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CreateGroupAppInstall201Response) SetLinks(v PaginatedLinks) {
	o.Links = v
}

func (o CreateGroupAppInstall201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupAppInstall201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["jsonapi"] = o.Jsonapi
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CreateGroupAppInstall201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"jsonapi",
		"links",
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

	varCreateGroupAppInstall201Response := _CreateGroupAppInstall201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGroupAppInstall201Response)

	if err != nil {
		return err
	}

	*o = CreateGroupAppInstall201Response(varCreateGroupAppInstall201Response)

	return err
}

type NullableCreateGroupAppInstall201Response struct {
	value *CreateGroupAppInstall201Response
	isSet bool
}

func (v NullableCreateGroupAppInstall201Response) Get() *CreateGroupAppInstall201Response {
	return v.value
}

func (v *NullableCreateGroupAppInstall201Response) Set(val *CreateGroupAppInstall201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupAppInstall201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupAppInstall201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupAppInstall201Response(val *CreateGroupAppInstall201Response) *NullableCreateGroupAppInstall201Response {
	return &NullableCreateGroupAppInstall201Response{value: val, isSet: true}
}

func (v NullableCreateGroupAppInstall201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupAppInstall201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

