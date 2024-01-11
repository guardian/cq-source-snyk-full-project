# PullRequestsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailOnlyForIssuesWithFix** | Pointer to **bool** | Only fail when the issues found have a fix available. | [optional] 
**Policy** | Pointer to **string** | Fail if the project has any issues (\&quot;all\&quot;), or fail if a PR is introducing a new dependency with issues (\&quot;only_new\&quot;). If this value is unset, the setting is inherited from the org default. | [optional] 
**SeverityThreshold** | Pointer to **string** | Only fail for issues greater than or equal to the specified severity. If this value is unset, the setting is inherited from the org default. | [optional] 

## Methods

### NewPullRequestsSettings

`func NewPullRequestsSettings() *PullRequestsSettings`

NewPullRequestsSettings instantiates a new PullRequestsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestsSettingsWithDefaults

`func NewPullRequestsSettingsWithDefaults() *PullRequestsSettings`

NewPullRequestsSettingsWithDefaults instantiates a new PullRequestsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailOnlyForIssuesWithFix

`func (o *PullRequestsSettings) GetFailOnlyForIssuesWithFix() bool`

GetFailOnlyForIssuesWithFix returns the FailOnlyForIssuesWithFix field if non-nil, zero value otherwise.

### GetFailOnlyForIssuesWithFixOk

`func (o *PullRequestsSettings) GetFailOnlyForIssuesWithFixOk() (*bool, bool)`

GetFailOnlyForIssuesWithFixOk returns a tuple with the FailOnlyForIssuesWithFix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnlyForIssuesWithFix

`func (o *PullRequestsSettings) SetFailOnlyForIssuesWithFix(v bool)`

SetFailOnlyForIssuesWithFix sets FailOnlyForIssuesWithFix field to given value.

### HasFailOnlyForIssuesWithFix

`func (o *PullRequestsSettings) HasFailOnlyForIssuesWithFix() bool`

HasFailOnlyForIssuesWithFix returns a boolean if a field has been set.

### GetPolicy

`func (o *PullRequestsSettings) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PullRequestsSettings) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PullRequestsSettings) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PullRequestsSettings) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSeverityThreshold

`func (o *PullRequestsSettings) GetSeverityThreshold() string`

GetSeverityThreshold returns the SeverityThreshold field if non-nil, zero value otherwise.

### GetSeverityThresholdOk

`func (o *PullRequestsSettings) GetSeverityThresholdOk() (*string, bool)`

GetSeverityThresholdOk returns a tuple with the SeverityThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityThreshold

`func (o *PullRequestsSettings) SetSeverityThreshold(v string)`

SetSeverityThreshold sets SeverityThreshold field to given value.

### HasSeverityThreshold

`func (o *PullRequestsSettings) HasSeverityThreshold() bool`

HasSeverityThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


