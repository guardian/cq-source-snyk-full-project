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

// checks if the PaginatedLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedLinks{}

// PaginatedLinks struct for PaginatedLinks
type PaginatedLinks struct {
	First *LinkProperty `json:"first,omitempty"`
	Last *LinkProperty `json:"last,omitempty"`
	Next *LinkProperty `json:"next,omitempty"`
	Prev *LinkProperty `json:"prev,omitempty"`
	Self *LinkProperty `json:"self,omitempty"`
}

// NewPaginatedLinks instantiates a new PaginatedLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLinks() *PaginatedLinks {
	this := PaginatedLinks{}
	return &this
}

// NewPaginatedLinksWithDefaults instantiates a new PaginatedLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLinksWithDefaults() *PaginatedLinks {
	this := PaginatedLinks{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *PaginatedLinks) GetFirst() LinkProperty {
	if o == nil || IsNil(o.First) {
		var ret LinkProperty
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLinks) GetFirstOk() (*LinkProperty, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *PaginatedLinks) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given LinkProperty and assigns it to the First field.
func (o *PaginatedLinks) SetFirst(v LinkProperty) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *PaginatedLinks) GetLast() LinkProperty {
	if o == nil || IsNil(o.Last) {
		var ret LinkProperty
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLinks) GetLastOk() (*LinkProperty, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *PaginatedLinks) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given LinkProperty and assigns it to the Last field.
func (o *PaginatedLinks) SetLast(v LinkProperty) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedLinks) GetNext() LinkProperty {
	if o == nil || IsNil(o.Next) {
		var ret LinkProperty
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLinks) GetNextOk() (*LinkProperty, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given LinkProperty and assigns it to the Next field.
func (o *PaginatedLinks) SetNext(v LinkProperty) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *PaginatedLinks) GetPrev() LinkProperty {
	if o == nil || IsNil(o.Prev) {
		var ret LinkProperty
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLinks) GetPrevOk() (*LinkProperty, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *PaginatedLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given LinkProperty and assigns it to the Prev field.
func (o *PaginatedLinks) SetPrev(v LinkProperty) {
	o.Prev = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PaginatedLinks) GetSelf() LinkProperty {
	if o == nil || IsNil(o.Self) {
		var ret LinkProperty
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLinks) GetSelfOk() (*LinkProperty, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PaginatedLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given LinkProperty and assigns it to the Self field.
func (o *PaginatedLinks) SetSelf(v LinkProperty) {
	o.Self = &v
}

func (o PaginatedLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullablePaginatedLinks struct {
	value *PaginatedLinks
	isSet bool
}

func (v NullablePaginatedLinks) Get() *PaginatedLinks {
	return v.value
}

func (v *NullablePaginatedLinks) Set(val *PaginatedLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLinks(val *PaginatedLinks) *NullablePaginatedLinks {
	return &NullablePaginatedLinks{value: val, isSet: true}
}

func (v NullablePaginatedLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


