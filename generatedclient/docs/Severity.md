# Severity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **NullableFloat32** | The CVSSv3 value of the vulnerability. | [optional] 
**Source** | Pointer to **string** | The source of this severity. The value must be the id of a referenced problem or class, in which case that problem or class is the source of this issue. If source is omitted, this severity is sourced internally in the Snyk application. | [optional] 
**Vector** | Pointer to **NullableString** | The CVSSv3 value of the vulnerability. | [optional] 

## Methods

### NewSeverity

`func NewSeverity() *Severity`

NewSeverity instantiates a new Severity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeverityWithDefaults

`func NewSeverityWithDefaults() *Severity`

NewSeverityWithDefaults instantiates a new Severity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *Severity) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Severity) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Severity) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Severity) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetScore

`func (o *Severity) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Severity) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Severity) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Severity) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *Severity) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *Severity) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetSource

`func (o *Severity) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Severity) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Severity) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Severity) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVector

`func (o *Severity) GetVector() string`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *Severity) GetVectorOk() (*string, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *Severity) SetVector(v string)`

SetVector sets Vector field to given value.

### HasVector

`func (o *Severity) HasVector() bool`

HasVector returns a boolean if a field has been set.

### SetVectorNil

`func (o *Severity) SetVectorNil(b bool)`

 SetVectorNil sets the value for Vector to be an explicit nil

### UnsetVector
`func (o *Severity) UnsetVector()`

UnsetVector ensures that no value is present for Vector, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


