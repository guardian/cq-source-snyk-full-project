# RecurringTestsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **string** | Test frequency of a project. Also controls when automated PRs may be created. | [optional] 

## Methods

### NewRecurringTestsSettings

`func NewRecurringTestsSettings() *RecurringTestsSettings`

NewRecurringTestsSettings instantiates a new RecurringTestsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringTestsSettingsWithDefaults

`func NewRecurringTestsSettingsWithDefaults() *RecurringTestsSettings`

NewRecurringTestsSettingsWithDefaults instantiates a new RecurringTestsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *RecurringTestsSettings) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RecurringTestsSettings) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RecurringTestsSettings) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *RecurringTestsSettings) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


