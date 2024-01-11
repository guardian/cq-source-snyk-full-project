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

// checks if the VersioningSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersioningSchema{}

// VersioningSchema The versioning scheme used by images in the repository.  A versioning schema is a system for identifying and organizing different versions of a project.  It is used to track changes and updates to the project over time, and to help users identify which version they are using.  A versioning schema typically consists of a series of numbers or labels that are incremented to reflect the progression of versions.  For example, a versioning schema might use a series of numbers, such as \"1.0\", \"1.1\", \"2.0\", and so on, to indicate major and minor releases of a product.  A consistent and well-defined versioning schema helps users and tools understand and track the development of a project. 
type VersioningSchema struct {
	Type string `json:"type"`
	// The regular expression used to describe the format of tags. Keep in mind that backslashes in the expression need to be escaped due to being encompassed in a JSON string. 
	Expression string `json:"expression"`
	// A customizable string that can be set for a custom versioning schema to describe its meaning. This label has no function. 
	Label *string `json:"label,omitempty"`
	// Whether this image should be the recommendation. Only one image can be selected at a given time. Setting this as true will remove previous selection. 
	IsSelected bool `json:"is_selected"`
}

type _VersioningSchema VersioningSchema

// NewVersioningSchema instantiates a new VersioningSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersioningSchema(type_ string, expression string, isSelected bool) *VersioningSchema {
	this := VersioningSchema{}
	this.Type = type_
	this.Expression = expression
	this.IsSelected = isSelected
	return &this
}

// NewVersioningSchemaWithDefaults instantiates a new VersioningSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersioningSchemaWithDefaults() *VersioningSchema {
	this := VersioningSchema{}
	return &this
}

// GetType returns the Type field value
func (o *VersioningSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VersioningSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VersioningSchema) SetType(v string) {
	o.Type = v
}

// GetExpression returns the Expression field value
func (o *VersioningSchema) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *VersioningSchema) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *VersioningSchema) SetExpression(v string) {
	o.Expression = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *VersioningSchema) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersioningSchema) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *VersioningSchema) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *VersioningSchema) SetLabel(v string) {
	o.Label = &v
}

// GetIsSelected returns the IsSelected field value
func (o *VersioningSchema) GetIsSelected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSelected
}

// GetIsSelectedOk returns a tuple with the IsSelected field value
// and a boolean to check if the value has been set.
func (o *VersioningSchema) GetIsSelectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSelected, true
}

// SetIsSelected sets field value
func (o *VersioningSchema) SetIsSelected(v bool) {
	o.IsSelected = v
}

func (o VersioningSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersioningSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["expression"] = o.Expression
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["is_selected"] = o.IsSelected
	return toSerialize, nil
}

func (o *VersioningSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"expression",
		"is_selected",
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

	varVersioningSchema := _VersioningSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVersioningSchema)

	if err != nil {
		return err
	}

	*o = VersioningSchema(varVersioningSchema)

	return err
}

type NullableVersioningSchema struct {
	value *VersioningSchema
	isSet bool
}

func (v NullableVersioningSchema) Get() *VersioningSchema {
	return v.value
}

func (v *NullableVersioningSchema) Set(val *VersioningSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableVersioningSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableVersioningSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersioningSchema(val *VersioningSchema) *NullableVersioningSchema {
	return &NullableVersioningSchema{value: val, isSet: true}
}

func (v NullableVersioningSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersioningSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

