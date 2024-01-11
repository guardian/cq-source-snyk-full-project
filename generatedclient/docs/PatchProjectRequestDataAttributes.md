# PatchProjectRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessCriticality** | Pointer to **[]string** |  | [optional] 
**Environment** | Pointer to **[]string** |  | [optional] 
**Lifecycle** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]PatchProjectRequestDataAttributesTagsInner**](PatchProjectRequestDataAttributesTagsInner.md) |  | [optional] 
**TestFrequency** | Pointer to **string** | Test frequency of a project. Also controls when automated PRs may be created. | [optional] 

## Methods

### NewPatchProjectRequestDataAttributes

`func NewPatchProjectRequestDataAttributes() *PatchProjectRequestDataAttributes`

NewPatchProjectRequestDataAttributes instantiates a new PatchProjectRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchProjectRequestDataAttributesWithDefaults

`func NewPatchProjectRequestDataAttributesWithDefaults() *PatchProjectRequestDataAttributes`

NewPatchProjectRequestDataAttributesWithDefaults instantiates a new PatchProjectRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessCriticality

`func (o *PatchProjectRequestDataAttributes) GetBusinessCriticality() []string`

GetBusinessCriticality returns the BusinessCriticality field if non-nil, zero value otherwise.

### GetBusinessCriticalityOk

`func (o *PatchProjectRequestDataAttributes) GetBusinessCriticalityOk() (*[]string, bool)`

GetBusinessCriticalityOk returns a tuple with the BusinessCriticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCriticality

`func (o *PatchProjectRequestDataAttributes) SetBusinessCriticality(v []string)`

SetBusinessCriticality sets BusinessCriticality field to given value.

### HasBusinessCriticality

`func (o *PatchProjectRequestDataAttributes) HasBusinessCriticality() bool`

HasBusinessCriticality returns a boolean if a field has been set.

### GetEnvironment

`func (o *PatchProjectRequestDataAttributes) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PatchProjectRequestDataAttributes) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PatchProjectRequestDataAttributes) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PatchProjectRequestDataAttributes) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLifecycle

`func (o *PatchProjectRequestDataAttributes) GetLifecycle() []string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *PatchProjectRequestDataAttributes) GetLifecycleOk() (*[]string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *PatchProjectRequestDataAttributes) SetLifecycle(v []string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *PatchProjectRequestDataAttributes) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetTags

`func (o *PatchProjectRequestDataAttributes) GetTags() []PatchProjectRequestDataAttributesTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchProjectRequestDataAttributes) GetTagsOk() (*[]PatchProjectRequestDataAttributesTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchProjectRequestDataAttributes) SetTags(v []PatchProjectRequestDataAttributesTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchProjectRequestDataAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTestFrequency

`func (o *PatchProjectRequestDataAttributes) GetTestFrequency() string`

GetTestFrequency returns the TestFrequency field if non-nil, zero value otherwise.

### GetTestFrequencyOk

`func (o *PatchProjectRequestDataAttributes) GetTestFrequencyOk() (*string, bool)`

GetTestFrequencyOk returns a tuple with the TestFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFrequency

`func (o *PatchProjectRequestDataAttributes) SetTestFrequency(v string)`

SetTestFrequency sets TestFrequency field to given value.

### HasTestFrequency

`func (o *PatchProjectRequestDataAttributes) HasTestFrequency() bool`

HasTestFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


