# AuditLogSearchItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **time.Time** |  | 
**Event** | **string** |  | 
**GroupId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewAuditLogSearchItemsInner

`func NewAuditLogSearchItemsInner(created time.Time, event string, ) *AuditLogSearchItemsInner`

NewAuditLogSearchItemsInner instantiates a new AuditLogSearchItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogSearchItemsInnerWithDefaults

`func NewAuditLogSearchItemsInnerWithDefaults() *AuditLogSearchItemsInner`

NewAuditLogSearchItemsInnerWithDefaults instantiates a new AuditLogSearchItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *AuditLogSearchItemsInner) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AuditLogSearchItemsInner) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AuditLogSearchItemsInner) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *AuditLogSearchItemsInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreated

`func (o *AuditLogSearchItemsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuditLogSearchItemsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuditLogSearchItemsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEvent

`func (o *AuditLogSearchItemsInner) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AuditLogSearchItemsInner) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AuditLogSearchItemsInner) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetGroupId

`func (o *AuditLogSearchItemsInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AuditLogSearchItemsInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AuditLogSearchItemsInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AuditLogSearchItemsInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetOrgId

`func (o *AuditLogSearchItemsInner) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuditLogSearchItemsInner) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuditLogSearchItemsInner) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuditLogSearchItemsInner) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *AuditLogSearchItemsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AuditLogSearchItemsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AuditLogSearchItemsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AuditLogSearchItemsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


