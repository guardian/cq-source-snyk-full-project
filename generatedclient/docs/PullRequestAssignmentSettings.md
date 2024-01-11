# PullRequestAssignmentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to **[]string** | Manually specify users to assign (and all will be assigned). | [optional] 
**IsEnabled** | Pointer to **bool** | Automatically assign pull requests created by Snyk. | [optional] 
**Type** | Pointer to **string** | Automatically assign the last user to change the manifest file (\&quot;auto\&quot;), or manually specify a list of users (\&quot;manual\&quot;). | [optional] 

## Methods

### NewPullRequestAssignmentSettings

`func NewPullRequestAssignmentSettings() *PullRequestAssignmentSettings`

NewPullRequestAssignmentSettings instantiates a new PullRequestAssignmentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestAssignmentSettingsWithDefaults

`func NewPullRequestAssignmentSettingsWithDefaults() *PullRequestAssignmentSettings`

NewPullRequestAssignmentSettingsWithDefaults instantiates a new PullRequestAssignmentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *PullRequestAssignmentSettings) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *PullRequestAssignmentSettings) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *PullRequestAssignmentSettings) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *PullRequestAssignmentSettings) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PullRequestAssignmentSettings) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PullRequestAssignmentSettings) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PullRequestAssignmentSettings) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PullRequestAssignmentSettings) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetType

`func (o *PullRequestAssignmentSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PullRequestAssignmentSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PullRequestAssignmentSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PullRequestAssignmentSettings) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


