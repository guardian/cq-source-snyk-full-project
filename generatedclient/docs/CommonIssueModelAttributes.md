# CommonIssueModelAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | Pointer to [**[]Coordinate**](Coordinate.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | A description of the issue in Markdown format | [optional] 
**EffectiveSeverityLevel** | Pointer to **string** | The type from enumeration of the issue’s severity level. This is usually set from the issue’s producer, but can be overridden by policies. | [optional] 
**Key** | Pointer to **string** | The Snyk vulnerability ID. | [optional] 
**Problems** | Pointer to [**[]Problem**](Problem.md) |  | [optional] 
**Severities** | Pointer to [**[]Severity**](Severity.md) | The severity level of the vulnerability: ‘low’, ‘medium’, ‘high’ or ‘critical’. | [optional] 
**Slots** | Pointer to [**Slots**](Slots.md) |  | [optional] 
**Title** | Pointer to **string** | A human-readable title for this issue. | [optional] 
**Type** | Pointer to **string** | The issue type | [optional] 
**UpdatedAt** | Pointer to **time.Time** | When the vulnerability information was last modified. | [optional] 

## Methods

### NewCommonIssueModelAttributes

`func NewCommonIssueModelAttributes() *CommonIssueModelAttributes`

NewCommonIssueModelAttributes instantiates a new CommonIssueModelAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonIssueModelAttributesWithDefaults

`func NewCommonIssueModelAttributesWithDefaults() *CommonIssueModelAttributes`

NewCommonIssueModelAttributesWithDefaults instantiates a new CommonIssueModelAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *CommonIssueModelAttributes) GetCoordinates() []Coordinate`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *CommonIssueModelAttributes) GetCoordinatesOk() (*[]Coordinate, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *CommonIssueModelAttributes) SetCoordinates(v []Coordinate)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *CommonIssueModelAttributes) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommonIssueModelAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommonIssueModelAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommonIssueModelAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommonIssueModelAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CommonIssueModelAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonIssueModelAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonIssueModelAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonIssueModelAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveSeverityLevel

`func (o *CommonIssueModelAttributes) GetEffectiveSeverityLevel() string`

GetEffectiveSeverityLevel returns the EffectiveSeverityLevel field if non-nil, zero value otherwise.

### GetEffectiveSeverityLevelOk

`func (o *CommonIssueModelAttributes) GetEffectiveSeverityLevelOk() (*string, bool)`

GetEffectiveSeverityLevelOk returns a tuple with the EffectiveSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSeverityLevel

`func (o *CommonIssueModelAttributes) SetEffectiveSeverityLevel(v string)`

SetEffectiveSeverityLevel sets EffectiveSeverityLevel field to given value.

### HasEffectiveSeverityLevel

`func (o *CommonIssueModelAttributes) HasEffectiveSeverityLevel() bool`

HasEffectiveSeverityLevel returns a boolean if a field has been set.

### GetKey

`func (o *CommonIssueModelAttributes) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CommonIssueModelAttributes) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CommonIssueModelAttributes) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CommonIssueModelAttributes) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetProblems

`func (o *CommonIssueModelAttributes) GetProblems() []Problem`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *CommonIssueModelAttributes) GetProblemsOk() (*[]Problem, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *CommonIssueModelAttributes) SetProblems(v []Problem)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *CommonIssueModelAttributes) HasProblems() bool`

HasProblems returns a boolean if a field has been set.

### GetSeverities

`func (o *CommonIssueModelAttributes) GetSeverities() []Severity`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *CommonIssueModelAttributes) GetSeveritiesOk() (*[]Severity, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *CommonIssueModelAttributes) SetSeverities(v []Severity)`

SetSeverities sets Severities field to given value.

### HasSeverities

`func (o *CommonIssueModelAttributes) HasSeverities() bool`

HasSeverities returns a boolean if a field has been set.

### GetSlots

`func (o *CommonIssueModelAttributes) GetSlots() Slots`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *CommonIssueModelAttributes) GetSlotsOk() (*Slots, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *CommonIssueModelAttributes) SetSlots(v Slots)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *CommonIssueModelAttributes) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetTitle

`func (o *CommonIssueModelAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CommonIssueModelAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CommonIssueModelAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CommonIssueModelAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *CommonIssueModelAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonIssueModelAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonIssueModelAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommonIssueModelAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommonIssueModelAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommonIssueModelAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommonIssueModelAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommonIssueModelAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


