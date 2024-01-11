# ProjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imported** | **time.Time** | The time the project was imported | 
**IssuesCriticalCount** | **float32** | The sum of critical severity issues of the project | 
**IssuesHighCount** | **float32** | The sum of high severity issues of the project | 
**IssuesLowCount** | **float32** | The sum of low severity issues of the project | 
**IssuesMediumCount** | **float32** | The sum of medium severity issues of the project | 
**LastTestedAt** | **time.Time** | The time the project was last tested | 

## Methods

### NewProjectMeta

`func NewProjectMeta(imported time.Time, issuesCriticalCount float32, issuesHighCount float32, issuesLowCount float32, issuesMediumCount float32, lastTestedAt time.Time, ) *ProjectMeta`

NewProjectMeta instantiates a new ProjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectMetaWithDefaults

`func NewProjectMetaWithDefaults() *ProjectMeta`

NewProjectMetaWithDefaults instantiates a new ProjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImported

`func (o *ProjectMeta) GetImported() time.Time`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *ProjectMeta) GetImportedOk() (*time.Time, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *ProjectMeta) SetImported(v time.Time)`

SetImported sets Imported field to given value.


### GetIssuesCriticalCount

`func (o *ProjectMeta) GetIssuesCriticalCount() float32`

GetIssuesCriticalCount returns the IssuesCriticalCount field if non-nil, zero value otherwise.

### GetIssuesCriticalCountOk

`func (o *ProjectMeta) GetIssuesCriticalCountOk() (*float32, bool)`

GetIssuesCriticalCountOk returns a tuple with the IssuesCriticalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesCriticalCount

`func (o *ProjectMeta) SetIssuesCriticalCount(v float32)`

SetIssuesCriticalCount sets IssuesCriticalCount field to given value.


### GetIssuesHighCount

`func (o *ProjectMeta) GetIssuesHighCount() float32`

GetIssuesHighCount returns the IssuesHighCount field if non-nil, zero value otherwise.

### GetIssuesHighCountOk

`func (o *ProjectMeta) GetIssuesHighCountOk() (*float32, bool)`

GetIssuesHighCountOk returns a tuple with the IssuesHighCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesHighCount

`func (o *ProjectMeta) SetIssuesHighCount(v float32)`

SetIssuesHighCount sets IssuesHighCount field to given value.


### GetIssuesLowCount

`func (o *ProjectMeta) GetIssuesLowCount() float32`

GetIssuesLowCount returns the IssuesLowCount field if non-nil, zero value otherwise.

### GetIssuesLowCountOk

`func (o *ProjectMeta) GetIssuesLowCountOk() (*float32, bool)`

GetIssuesLowCountOk returns a tuple with the IssuesLowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesLowCount

`func (o *ProjectMeta) SetIssuesLowCount(v float32)`

SetIssuesLowCount sets IssuesLowCount field to given value.


### GetIssuesMediumCount

`func (o *ProjectMeta) GetIssuesMediumCount() float32`

GetIssuesMediumCount returns the IssuesMediumCount field if non-nil, zero value otherwise.

### GetIssuesMediumCountOk

`func (o *ProjectMeta) GetIssuesMediumCountOk() (*float32, bool)`

GetIssuesMediumCountOk returns a tuple with the IssuesMediumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesMediumCount

`func (o *ProjectMeta) SetIssuesMediumCount(v float32)`

SetIssuesMediumCount sets IssuesMediumCount field to given value.


### GetLastTestedAt

`func (o *ProjectMeta) GetLastTestedAt() time.Time`

GetLastTestedAt returns the LastTestedAt field if non-nil, zero value otherwise.

### GetLastTestedAtOk

`func (o *ProjectMeta) GetLastTestedAtOk() (*time.Time, bool)`

GetLastTestedAtOk returns a tuple with the LastTestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestedAt

`func (o *ProjectMeta) SetLastTestedAt(v time.Time)`

SetLastTestedAt sets LastTestedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


