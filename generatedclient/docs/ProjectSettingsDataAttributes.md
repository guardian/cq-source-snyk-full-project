# ProjectSettingsDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | **bool** | Current status of the project settings. | 
**SeverityThreshold** | [**SeverityThreshold**](SeverityThreshold.md) |  | 
**TargetChannelId** | **string** |  | 
**TargetChannelName** | **string** |  | 
**TargetProjectName** | **string** | The target file name for the project. | 

## Methods

### NewProjectSettingsDataAttributes

`func NewProjectSettingsDataAttributes(isActive bool, severityThreshold SeverityThreshold, targetChannelId string, targetChannelName string, targetProjectName string, ) *ProjectSettingsDataAttributes`

NewProjectSettingsDataAttributes instantiates a new ProjectSettingsDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSettingsDataAttributesWithDefaults

`func NewProjectSettingsDataAttributesWithDefaults() *ProjectSettingsDataAttributes`

NewProjectSettingsDataAttributesWithDefaults instantiates a new ProjectSettingsDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *ProjectSettingsDataAttributes) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProjectSettingsDataAttributes) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProjectSettingsDataAttributes) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetSeverityThreshold

`func (o *ProjectSettingsDataAttributes) GetSeverityThreshold() SeverityThreshold`

GetSeverityThreshold returns the SeverityThreshold field if non-nil, zero value otherwise.

### GetSeverityThresholdOk

`func (o *ProjectSettingsDataAttributes) GetSeverityThresholdOk() (*SeverityThreshold, bool)`

GetSeverityThresholdOk returns a tuple with the SeverityThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityThreshold

`func (o *ProjectSettingsDataAttributes) SetSeverityThreshold(v SeverityThreshold)`

SetSeverityThreshold sets SeverityThreshold field to given value.


### GetTargetChannelId

`func (o *ProjectSettingsDataAttributes) GetTargetChannelId() string`

GetTargetChannelId returns the TargetChannelId field if non-nil, zero value otherwise.

### GetTargetChannelIdOk

`func (o *ProjectSettingsDataAttributes) GetTargetChannelIdOk() (*string, bool)`

GetTargetChannelIdOk returns a tuple with the TargetChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChannelId

`func (o *ProjectSettingsDataAttributes) SetTargetChannelId(v string)`

SetTargetChannelId sets TargetChannelId field to given value.


### GetTargetChannelName

`func (o *ProjectSettingsDataAttributes) GetTargetChannelName() string`

GetTargetChannelName returns the TargetChannelName field if non-nil, zero value otherwise.

### GetTargetChannelNameOk

`func (o *ProjectSettingsDataAttributes) GetTargetChannelNameOk() (*string, bool)`

GetTargetChannelNameOk returns a tuple with the TargetChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChannelName

`func (o *ProjectSettingsDataAttributes) SetTargetChannelName(v string)`

SetTargetChannelName sets TargetChannelName field to given value.


### GetTargetProjectName

`func (o *ProjectSettingsDataAttributes) GetTargetProjectName() string`

GetTargetProjectName returns the TargetProjectName field if non-nil, zero value otherwise.

### GetTargetProjectNameOk

`func (o *ProjectSettingsDataAttributes) GetTargetProjectNameOk() (*string, bool)`

GetTargetProjectNameOk returns a tuple with the TargetProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProjectName

`func (o *ProjectSettingsDataAttributes) SetTargetProjectName(v string)`

SetTargetProjectName sets TargetProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


