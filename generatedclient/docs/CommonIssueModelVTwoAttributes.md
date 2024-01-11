# CommonIssueModelVTwoAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | Pointer to [**[]CoordinateVTwo**](CoordinateVTwo.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | A description of the issue in Markdown format | [optional] 
**EffectiveSeverityLevel** | Pointer to **string** | The type from enumeration of the issue’s severity level. This is usually set from the issue’s producer, but can be overridden by policies. | [optional] 
**Problems** | Pointer to [**[]Problem**](Problem.md) |  | [optional] 
**Severities** | Pointer to [**[]Severity**](Severity.md) | The severity level of the vulnerability: ‘low’, ‘medium’, ‘high’ or ‘critical’. | [optional] 
**Slots** | Pointer to [**Slots**](Slots.md) |  | [optional] 
**Title** | Pointer to **string** | A human-readable title for this issue. | [optional] 
**Type** | Pointer to **string** | The issue type | [optional] 
**UpdatedAt** | Pointer to **time.Time** | When the vulnerability information was last modified. | [optional] 

## Methods

### NewCommonIssueModelVTwoAttributes

`func NewCommonIssueModelVTwoAttributes() *CommonIssueModelVTwoAttributes`

NewCommonIssueModelVTwoAttributes instantiates a new CommonIssueModelVTwoAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonIssueModelVTwoAttributesWithDefaults

`func NewCommonIssueModelVTwoAttributesWithDefaults() *CommonIssueModelVTwoAttributes`

NewCommonIssueModelVTwoAttributesWithDefaults instantiates a new CommonIssueModelVTwoAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *CommonIssueModelVTwoAttributes) GetCoordinates() []CoordinateVTwo`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *CommonIssueModelVTwoAttributes) GetCoordinatesOk() (*[]CoordinateVTwo, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *CommonIssueModelVTwoAttributes) SetCoordinates(v []CoordinateVTwo)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *CommonIssueModelVTwoAttributes) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommonIssueModelVTwoAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommonIssueModelVTwoAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommonIssueModelVTwoAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommonIssueModelVTwoAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CommonIssueModelVTwoAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonIssueModelVTwoAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonIssueModelVTwoAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonIssueModelVTwoAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveSeverityLevel

`func (o *CommonIssueModelVTwoAttributes) GetEffectiveSeverityLevel() string`

GetEffectiveSeverityLevel returns the EffectiveSeverityLevel field if non-nil, zero value otherwise.

### GetEffectiveSeverityLevelOk

`func (o *CommonIssueModelVTwoAttributes) GetEffectiveSeverityLevelOk() (*string, bool)`

GetEffectiveSeverityLevelOk returns a tuple with the EffectiveSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSeverityLevel

`func (o *CommonIssueModelVTwoAttributes) SetEffectiveSeverityLevel(v string)`

SetEffectiveSeverityLevel sets EffectiveSeverityLevel field to given value.

### HasEffectiveSeverityLevel

`func (o *CommonIssueModelVTwoAttributes) HasEffectiveSeverityLevel() bool`

HasEffectiveSeverityLevel returns a boolean if a field has been set.

### GetProblems

`func (o *CommonIssueModelVTwoAttributes) GetProblems() []Problem`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *CommonIssueModelVTwoAttributes) GetProblemsOk() (*[]Problem, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *CommonIssueModelVTwoAttributes) SetProblems(v []Problem)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *CommonIssueModelVTwoAttributes) HasProblems() bool`

HasProblems returns a boolean if a field has been set.

### GetSeverities

`func (o *CommonIssueModelVTwoAttributes) GetSeverities() []Severity`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *CommonIssueModelVTwoAttributes) GetSeveritiesOk() (*[]Severity, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *CommonIssueModelVTwoAttributes) SetSeverities(v []Severity)`

SetSeverities sets Severities field to given value.

### HasSeverities

`func (o *CommonIssueModelVTwoAttributes) HasSeverities() bool`

HasSeverities returns a boolean if a field has been set.

### GetSlots

`func (o *CommonIssueModelVTwoAttributes) GetSlots() Slots`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *CommonIssueModelVTwoAttributes) GetSlotsOk() (*Slots, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *CommonIssueModelVTwoAttributes) SetSlots(v Slots)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *CommonIssueModelVTwoAttributes) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetTitle

`func (o *CommonIssueModelVTwoAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CommonIssueModelVTwoAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CommonIssueModelVTwoAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CommonIssueModelVTwoAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CommonIssueModelVTwoAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonIssueModelVTwoAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonIssueModelVTwoAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommonIssueModelVTwoAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommonIssueModelVTwoAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommonIssueModelVTwoAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommonIssueModelVTwoAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommonIssueModelVTwoAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


