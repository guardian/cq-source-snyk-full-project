# ListOrgProjects200ResponseDataInnerMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliMonitoredAt** | Pointer to **NullableTime** | The date that the project was last uploaded and monitored using cli. | [optional] 
**LatestDependencyTotal** | Pointer to [**LatestDependencyTotal**](LatestDependencyTotal.md) |  | [optional] 
**LatestIssueCounts** | Pointer to [**LatestIssueCounts**](LatestIssueCounts.md) |  | [optional] 

## Methods

### NewListOrgProjects200ResponseDataInnerMeta

`func NewListOrgProjects200ResponseDataInnerMeta() *ListOrgProjects200ResponseDataInnerMeta`

NewListOrgProjects200ResponseDataInnerMeta instantiates a new ListOrgProjects200ResponseDataInnerMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrgProjects200ResponseDataInnerMetaWithDefaults

`func NewListOrgProjects200ResponseDataInnerMetaWithDefaults() *ListOrgProjects200ResponseDataInnerMeta`

NewListOrgProjects200ResponseDataInnerMetaWithDefaults instantiates a new ListOrgProjects200ResponseDataInnerMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliMonitoredAt

`func (o *ListOrgProjects200ResponseDataInnerMeta) GetCliMonitoredAt() time.Time`

GetCliMonitoredAt returns the CliMonitoredAt field if non-nil, zero value otherwise.

### GetCliMonitoredAtOk

`func (o *ListOrgProjects200ResponseDataInnerMeta) GetCliMonitoredAtOk() (*time.Time, bool)`

GetCliMonitoredAtOk returns a tuple with the CliMonitoredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliMonitoredAt

`func (o *ListOrgProjects200ResponseDataInnerMeta) SetCliMonitoredAt(v time.Time)`

SetCliMonitoredAt sets CliMonitoredAt field to given value.

### HasCliMonitoredAt

`func (o *ListOrgProjects200ResponseDataInnerMeta) HasCliMonitoredAt() bool`

HasCliMonitoredAt returns a boolean if a field has been set.

### SetCliMonitoredAtNil

`func (o *ListOrgProjects200ResponseDataInnerMeta) SetCliMonitoredAtNil(b bool)`

 SetCliMonitoredAtNil sets the value for CliMonitoredAt to be an explicit nil

### UnsetCliMonitoredAt
`func (o *ListOrgProjects200ResponseDataInnerMeta) UnsetCliMonitoredAt()`

UnsetCliMonitoredAt ensures that no value is present for CliMonitoredAt, not even an explicit nil
### GetLatestDependencyTotal

`func (o *ListOrgProjects200ResponseDataInnerMeta) GetLatestDependencyTotal() LatestDependencyTotal`

GetLatestDependencyTotal returns the LatestDependencyTotal field if non-nil, zero value otherwise.

### GetLatestDependencyTotalOk

`func (o *ListOrgProjects200ResponseDataInnerMeta) GetLatestDependencyTotalOk() (*LatestDependencyTotal, bool)`

GetLatestDependencyTotalOk returns a tuple with the LatestDependencyTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDependencyTotal

`func (o *ListOrgProjects200ResponseDataInnerMeta) SetLatestDependencyTotal(v LatestDependencyTotal)`

SetLatestDependencyTotal sets LatestDependencyTotal field to given value.

### HasLatestDependencyTotal

`func (o *ListOrgProjects200ResponseDataInnerMeta) HasLatestDependencyTotal() bool`

HasLatestDependencyTotal returns a boolean if a field has been set.

### GetLatestIssueCounts

`func (o *ListOrgProjects200ResponseDataInnerMeta) GetLatestIssueCounts() LatestIssueCounts`

GetLatestIssueCounts returns the LatestIssueCounts field if non-nil, zero value otherwise.

### GetLatestIssueCountsOk

`func (o *ListOrgProjects200ResponseDataInnerMeta) GetLatestIssueCountsOk() (*LatestIssueCounts, bool)`

GetLatestIssueCountsOk returns a tuple with the LatestIssueCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestIssueCounts

`func (o *ListOrgProjects200ResponseDataInnerMeta) SetLatestIssueCounts(v LatestIssueCounts)`

SetLatestIssueCounts sets LatestIssueCounts field to given value.

### HasLatestIssueCounts

`func (o *ListOrgProjects200ResponseDataInnerMeta) HasLatestIssueCounts() bool`

HasLatestIssueCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


