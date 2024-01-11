# ProjectSettingsPatchRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** | Current status of the project settings. | [optional] 
**SeverityThreshold** | Pointer to [**SeverityThreshold**](SeverityThreshold.md) |  | [optional] 
**TargetChannelId** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectSettingsPatchRequestDataAttributes

`func NewProjectSettingsPatchRequestDataAttributes() *ProjectSettingsPatchRequestDataAttributes`

NewProjectSettingsPatchRequestDataAttributes instantiates a new ProjectSettingsPatchRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSettingsPatchRequestDataAttributesWithDefaults

`func NewProjectSettingsPatchRequestDataAttributesWithDefaults() *ProjectSettingsPatchRequestDataAttributes`

NewProjectSettingsPatchRequestDataAttributesWithDefaults instantiates a new ProjectSettingsPatchRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *ProjectSettingsPatchRequestDataAttributes) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProjectSettingsPatchRequestDataAttributes) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProjectSettingsPatchRequestDataAttributes) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProjectSettingsPatchRequestDataAttributes) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSeverityThreshold

`func (o *ProjectSettingsPatchRequestDataAttributes) GetSeverityThreshold() SeverityThreshold`

GetSeverityThreshold returns the SeverityThreshold field if non-nil, zero value otherwise.

### GetSeverityThresholdOk

`func (o *ProjectSettingsPatchRequestDataAttributes) GetSeverityThresholdOk() (*SeverityThreshold, bool)`

GetSeverityThresholdOk returns a tuple with the SeverityThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityThreshold

`func (o *ProjectSettingsPatchRequestDataAttributes) SetSeverityThreshold(v SeverityThreshold)`

SetSeverityThreshold sets SeverityThreshold field to given value.

### HasSeverityThreshold

`func (o *ProjectSettingsPatchRequestDataAttributes) HasSeverityThreshold() bool`

HasSeverityThreshold returns a boolean if a field has been set.

### GetTargetChannelId

`func (o *ProjectSettingsPatchRequestDataAttributes) GetTargetChannelId() string`

GetTargetChannelId returns the TargetChannelId field if non-nil, zero value otherwise.

### GetTargetChannelIdOk

`func (o *ProjectSettingsPatchRequestDataAttributes) GetTargetChannelIdOk() (*string, bool)`

GetTargetChannelIdOk returns a tuple with the TargetChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChannelId

`func (o *ProjectSettingsPatchRequestDataAttributes) SetTargetChannelId(v string)`

SetTargetChannelId sets TargetChannelId field to given value.

### HasTargetChannelId

`func (o *ProjectSettingsPatchRequestDataAttributes) HasTargetChannelId() bool`

HasTargetChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


