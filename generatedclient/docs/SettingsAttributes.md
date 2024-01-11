# SettingsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeverityThreshold** | [**SeverityThreshold**](SeverityThreshold.md) |  | 
**TargetChannelId** | **string** |  | 

## Methods

### NewSettingsAttributes

`func NewSettingsAttributes(severityThreshold SeverityThreshold, targetChannelId string, ) *SettingsAttributes`

NewSettingsAttributes instantiates a new SettingsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsAttributesWithDefaults

`func NewSettingsAttributesWithDefaults() *SettingsAttributes`

NewSettingsAttributesWithDefaults instantiates a new SettingsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverityThreshold

`func (o *SettingsAttributes) GetSeverityThreshold() SeverityThreshold`

GetSeverityThreshold returns the SeverityThreshold field if non-nil, zero value otherwise.

### GetSeverityThresholdOk

`func (o *SettingsAttributes) GetSeverityThresholdOk() (*SeverityThreshold, bool)`

GetSeverityThresholdOk returns a tuple with the SeverityThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityThreshold

`func (o *SettingsAttributes) SetSeverityThreshold(v SeverityThreshold)`

SetSeverityThreshold sets SeverityThreshold field to given value.


### GetTargetChannelId

`func (o *SettingsAttributes) GetTargetChannelId() string`

GetTargetChannelId returns the TargetChannelId field if non-nil, zero value otherwise.

### GetTargetChannelIdOk

`func (o *SettingsAttributes) GetTargetChannelIdOk() (*string, bool)`

GetTargetChannelIdOk returns a tuple with the TargetChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChannelId

`func (o *SettingsAttributes) SetTargetChannelId(v string)`

SetTargetChannelId sets TargetChannelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


