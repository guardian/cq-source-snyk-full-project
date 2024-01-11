# CollectionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuesCriticalCount** | **float32** | The sum of critical severity issues of the collection&#39;s projects | 
**IssuesHighCount** | **float32** | The sum of high severity issues of the collection&#39;s projects | 
**IssuesLowCount** | **float32** | The sum of low severity issues of the collection&#39;s projects | 
**IssuesMediumCount** | **float32** | The sum of medium severity issues of the collection&#39;s projects | 
**ProjectsCount** | **float32** | The amount of projects belonging to this collection | 

## Methods

### NewCollectionMeta

`func NewCollectionMeta(issuesCriticalCount float32, issuesHighCount float32, issuesLowCount float32, issuesMediumCount float32, projectsCount float32, ) *CollectionMeta`

NewCollectionMeta instantiates a new CollectionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionMetaWithDefaults

`func NewCollectionMetaWithDefaults() *CollectionMeta`

NewCollectionMetaWithDefaults instantiates a new CollectionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuesCriticalCount

`func (o *CollectionMeta) GetIssuesCriticalCount() float32`

GetIssuesCriticalCount returns the IssuesCriticalCount field if non-nil, zero value otherwise.

### GetIssuesCriticalCountOk

`func (o *CollectionMeta) GetIssuesCriticalCountOk() (*float32, bool)`

GetIssuesCriticalCountOk returns a tuple with the IssuesCriticalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesCriticalCount

`func (o *CollectionMeta) SetIssuesCriticalCount(v float32)`

SetIssuesCriticalCount sets IssuesCriticalCount field to given value.


### GetIssuesHighCount

`func (o *CollectionMeta) GetIssuesHighCount() float32`

GetIssuesHighCount returns the IssuesHighCount field if non-nil, zero value otherwise.

### GetIssuesHighCountOk

`func (o *CollectionMeta) GetIssuesHighCountOk() (*float32, bool)`

GetIssuesHighCountOk returns a tuple with the IssuesHighCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesHighCount

`func (o *CollectionMeta) SetIssuesHighCount(v float32)`

SetIssuesHighCount sets IssuesHighCount field to given value.


### GetIssuesLowCount

`func (o *CollectionMeta) GetIssuesLowCount() float32`

GetIssuesLowCount returns the IssuesLowCount field if non-nil, zero value otherwise.

### GetIssuesLowCountOk

`func (o *CollectionMeta) GetIssuesLowCountOk() (*float32, bool)`

GetIssuesLowCountOk returns a tuple with the IssuesLowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesLowCount

`func (o *CollectionMeta) SetIssuesLowCount(v float32)`

SetIssuesLowCount sets IssuesLowCount field to given value.


### GetIssuesMediumCount

`func (o *CollectionMeta) GetIssuesMediumCount() float32`

GetIssuesMediumCount returns the IssuesMediumCount field if non-nil, zero value otherwise.

### GetIssuesMediumCountOk

`func (o *CollectionMeta) GetIssuesMediumCountOk() (*float32, bool)`

GetIssuesMediumCountOk returns a tuple with the IssuesMediumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesMediumCount

`func (o *CollectionMeta) SetIssuesMediumCount(v float32)`

SetIssuesMediumCount sets IssuesMediumCount field to given value.


### GetProjectsCount

`func (o *CollectionMeta) GetProjectsCount() float32`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *CollectionMeta) GetProjectsCountOk() (*float32, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *CollectionMeta) SetProjectsCount(v float32)`

SetProjectsCount sets ProjectsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


