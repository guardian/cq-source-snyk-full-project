# Remedy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A markdown-formatted optional description of this remedy. | [optional] 
**Details** | Pointer to [**RemedyDetails**](RemedyDetails.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the remedy. Always ‘indeterminate’. | [optional] 

## Methods

### NewRemedy

`func NewRemedy() *Remedy`

NewRemedy instantiates a new Remedy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemedyWithDefaults

`func NewRemedyWithDefaults() *Remedy`

NewRemedyWithDefaults instantiates a new Remedy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Remedy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Remedy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Remedy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Remedy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *Remedy) GetDetails() RemedyDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Remedy) GetDetailsOk() (*RemedyDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Remedy) SetDetails(v RemedyDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Remedy) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetType

`func (o *Remedy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Remedy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Remedy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Remedy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


