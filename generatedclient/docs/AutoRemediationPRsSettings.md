# AutoRemediationPRsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBacklogPrsEnabled** | Pointer to **bool** | Automatically create pull requests on scheduled tests for known (backlog) vulnerabilities. | [optional] 
**IsFreshPrsEnabled** | Pointer to **bool** | Automatically create pull requests on scheduled tests for new vulnerabilities. | [optional] 
**IsPatchRemediationEnabled** | Pointer to **bool** | Include vulnerability patches in automatic pull requests. | [optional] 

## Methods

### NewAutoRemediationPRsSettings

`func NewAutoRemediationPRsSettings() *AutoRemediationPRsSettings`

NewAutoRemediationPRsSettings instantiates a new AutoRemediationPRsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRemediationPRsSettingsWithDefaults

`func NewAutoRemediationPRsSettingsWithDefaults() *AutoRemediationPRsSettings`

NewAutoRemediationPRsSettingsWithDefaults instantiates a new AutoRemediationPRsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBacklogPrsEnabled

`func (o *AutoRemediationPRsSettings) GetIsBacklogPrsEnabled() bool`

GetIsBacklogPrsEnabled returns the IsBacklogPrsEnabled field if non-nil, zero value otherwise.

### GetIsBacklogPrsEnabledOk

`func (o *AutoRemediationPRsSettings) GetIsBacklogPrsEnabledOk() (*bool, bool)`

GetIsBacklogPrsEnabledOk returns a tuple with the IsBacklogPrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBacklogPrsEnabled

`func (o *AutoRemediationPRsSettings) SetIsBacklogPrsEnabled(v bool)`

SetIsBacklogPrsEnabled sets IsBacklogPrsEnabled field to given value.

### HasIsBacklogPrsEnabled

`func (o *AutoRemediationPRsSettings) HasIsBacklogPrsEnabled() bool`

HasIsBacklogPrsEnabled returns a boolean if a field has been set.

### GetIsFreshPrsEnabled

`func (o *AutoRemediationPRsSettings) GetIsFreshPrsEnabled() bool`

GetIsFreshPrsEnabled returns the IsFreshPrsEnabled field if non-nil, zero value otherwise.

### GetIsFreshPrsEnabledOk

`func (o *AutoRemediationPRsSettings) GetIsFreshPrsEnabledOk() (*bool, bool)`

GetIsFreshPrsEnabledOk returns a tuple with the IsFreshPrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreshPrsEnabled

`func (o *AutoRemediationPRsSettings) SetIsFreshPrsEnabled(v bool)`

SetIsFreshPrsEnabled sets IsFreshPrsEnabled field to given value.

### HasIsFreshPrsEnabled

`func (o *AutoRemediationPRsSettings) HasIsFreshPrsEnabled() bool`

HasIsFreshPrsEnabled returns a boolean if a field has been set.

### GetIsPatchRemediationEnabled

`func (o *AutoRemediationPRsSettings) GetIsPatchRemediationEnabled() bool`

GetIsPatchRemediationEnabled returns the IsPatchRemediationEnabled field if non-nil, zero value otherwise.

### GetIsPatchRemediationEnabledOk

`func (o *AutoRemediationPRsSettings) GetIsPatchRemediationEnabledOk() (*bool, bool)`

GetIsPatchRemediationEnabledOk returns a tuple with the IsPatchRemediationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPatchRemediationEnabled

`func (o *AutoRemediationPRsSettings) SetIsPatchRemediationEnabled(v bool)`

SetIsPatchRemediationEnabled sets IsPatchRemediationEnabled field to given value.

### HasIsPatchRemediationEnabled

`func (o *AutoRemediationPRsSettings) HasIsPatchRemediationEnabled() bool`

HasIsPatchRemediationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


