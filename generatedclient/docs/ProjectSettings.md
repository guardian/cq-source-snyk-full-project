# ProjectSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDependencyUpgrade** | Pointer to [**AutoDependencyUpgradeSettings**](AutoDependencyUpgradeSettings.md) |  | [optional] 
**AutoRemediationPrs** | Pointer to [**AutoRemediationPRsSettings**](AutoRemediationPRsSettings.md) |  | [optional] 
**ManualRemediationPrs** | Pointer to [**ManualRemediationPRsSettings**](ManualRemediationPRsSettings.md) |  | [optional] 
**PullRequestAssignment** | Pointer to [**PullRequestAssignmentSettings**](PullRequestAssignmentSettings.md) |  | [optional] 
**PullRequests** | [**PullRequestsSettings**](PullRequestsSettings.md) |  | 
**RecurringTests** | [**RecurringTestsSettings**](RecurringTestsSettings.md) |  | 

## Methods

### NewProjectSettings

`func NewProjectSettings(pullRequests PullRequestsSettings, recurringTests RecurringTestsSettings, ) *ProjectSettings`

NewProjectSettings instantiates a new ProjectSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSettingsWithDefaults

`func NewProjectSettingsWithDefaults() *ProjectSettings`

NewProjectSettingsWithDefaults instantiates a new ProjectSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDependencyUpgrade

`func (o *ProjectSettings) GetAutoDependencyUpgrade() AutoDependencyUpgradeSettings`

GetAutoDependencyUpgrade returns the AutoDependencyUpgrade field if non-nil, zero value otherwise.

### GetAutoDependencyUpgradeOk

`func (o *ProjectSettings) GetAutoDependencyUpgradeOk() (*AutoDependencyUpgradeSettings, bool)`

GetAutoDependencyUpgradeOk returns a tuple with the AutoDependencyUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDependencyUpgrade

`func (o *ProjectSettings) SetAutoDependencyUpgrade(v AutoDependencyUpgradeSettings)`

SetAutoDependencyUpgrade sets AutoDependencyUpgrade field to given value.

### HasAutoDependencyUpgrade

`func (o *ProjectSettings) HasAutoDependencyUpgrade() bool`

HasAutoDependencyUpgrade returns a boolean if a field has been set.

### GetAutoRemediationPrs

`func (o *ProjectSettings) GetAutoRemediationPrs() AutoRemediationPRsSettings`

GetAutoRemediationPrs returns the AutoRemediationPrs field if non-nil, zero value otherwise.

### GetAutoRemediationPrsOk

`func (o *ProjectSettings) GetAutoRemediationPrsOk() (*AutoRemediationPRsSettings, bool)`

GetAutoRemediationPrsOk returns a tuple with the AutoRemediationPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRemediationPrs

`func (o *ProjectSettings) SetAutoRemediationPrs(v AutoRemediationPRsSettings)`

SetAutoRemediationPrs sets AutoRemediationPrs field to given value.

### HasAutoRemediationPrs

`func (o *ProjectSettings) HasAutoRemediationPrs() bool`

HasAutoRemediationPrs returns a boolean if a field has been set.

### GetManualRemediationPrs

`func (o *ProjectSettings) GetManualRemediationPrs() ManualRemediationPRsSettings`

GetManualRemediationPrs returns the ManualRemediationPrs field if non-nil, zero value otherwise.

### GetManualRemediationPrsOk

`func (o *ProjectSettings) GetManualRemediationPrsOk() (*ManualRemediationPRsSettings, bool)`

GetManualRemediationPrsOk returns a tuple with the ManualRemediationPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRemediationPrs

`func (o *ProjectSettings) SetManualRemediationPrs(v ManualRemediationPRsSettings)`

SetManualRemediationPrs sets ManualRemediationPrs field to given value.

### HasManualRemediationPrs

`func (o *ProjectSettings) HasManualRemediationPrs() bool`

HasManualRemediationPrs returns a boolean if a field has been set.

### GetPullRequestAssignment

`func (o *ProjectSettings) GetPullRequestAssignment() PullRequestAssignmentSettings`

GetPullRequestAssignment returns the PullRequestAssignment field if non-nil, zero value otherwise.

### GetPullRequestAssignmentOk

`func (o *ProjectSettings) GetPullRequestAssignmentOk() (*PullRequestAssignmentSettings, bool)`

GetPullRequestAssignmentOk returns a tuple with the PullRequestAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestAssignment

`func (o *ProjectSettings) SetPullRequestAssignment(v PullRequestAssignmentSettings)`

SetPullRequestAssignment sets PullRequestAssignment field to given value.

### HasPullRequestAssignment

`func (o *ProjectSettings) HasPullRequestAssignment() bool`

HasPullRequestAssignment returns a boolean if a field has been set.

### GetPullRequests

`func (o *ProjectSettings) GetPullRequests() PullRequestsSettings`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *ProjectSettings) GetPullRequestsOk() (*PullRequestsSettings, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *ProjectSettings) SetPullRequests(v PullRequestsSettings)`

SetPullRequests sets PullRequests field to given value.


### GetRecurringTests

`func (o *ProjectSettings) GetRecurringTests() RecurringTestsSettings`

GetRecurringTests returns the RecurringTests field if non-nil, zero value otherwise.

### GetRecurringTestsOk

`func (o *ProjectSettings) GetRecurringTestsOk() (*RecurringTestsSettings, bool)`

GetRecurringTestsOk returns a tuple with the RecurringTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTests

`func (o *ProjectSettings) SetRecurringTests(v RecurringTestsSettings)`

SetRecurringTests sets RecurringTests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


