# Slots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisclosureTime** | Pointer to **time.Time** | The time at which this vulnerability was disclosed. | [optional] 
**Exploit** | Pointer to **string** | The exploit maturity. Value of ‘No Data’, ‘Not Defined’, ‘Unproven’, ‘Proof of Concept’, ‘Functional’ or ‘High’. | [optional] 
**PublicationTime** | Pointer to **string** | The time at which this vulnerability was published. | [optional] 
**References** | Pointer to [**[]SlotsReferencesInner**](SlotsReferencesInner.md) |  | [optional] 

## Methods

### NewSlots

`func NewSlots() *Slots`

NewSlots instantiates a new Slots object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlotsWithDefaults

`func NewSlotsWithDefaults() *Slots`

NewSlotsWithDefaults instantiates a new Slots object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclosureTime

`func (o *Slots) GetDisclosureTime() time.Time`

GetDisclosureTime returns the DisclosureTime field if non-nil, zero value otherwise.

### GetDisclosureTimeOk

`func (o *Slots) GetDisclosureTimeOk() (*time.Time, bool)`

GetDisclosureTimeOk returns a tuple with the DisclosureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureTime

`func (o *Slots) SetDisclosureTime(v time.Time)`

SetDisclosureTime sets DisclosureTime field to given value.

### HasDisclosureTime

`func (o *Slots) HasDisclosureTime() bool`

HasDisclosureTime returns a boolean if a field has been set.

### GetExploit

`func (o *Slots) GetExploit() string`

GetExploit returns the Exploit field if non-nil, zero value otherwise.

### GetExploitOk

`func (o *Slots) GetExploitOk() (*string, bool)`

GetExploitOk returns a tuple with the Exploit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploit

`func (o *Slots) SetExploit(v string)`

SetExploit sets Exploit field to given value.

### HasExploit

`func (o *Slots) HasExploit() bool`

HasExploit returns a boolean if a field has been set.

### GetPublicationTime

`func (o *Slots) GetPublicationTime() string`

GetPublicationTime returns the PublicationTime field if non-nil, zero value otherwise.

### GetPublicationTimeOk

`func (o *Slots) GetPublicationTimeOk() (*string, bool)`

GetPublicationTimeOk returns a tuple with the PublicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTime

`func (o *Slots) SetPublicationTime(v string)`

SetPublicationTime sets PublicationTime field to given value.

### HasPublicationTime

`func (o *Slots) HasPublicationTime() bool`

HasPublicationTime returns a boolean if a field has been set.

### GetReferences

`func (o *Slots) GetReferences() []SlotsReferencesInner`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Slots) GetReferencesOk() (*[]SlotsReferencesInner, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Slots) SetReferences(v []SlotsReferencesInner)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Slots) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


