# CoordinateVTwo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remedies** | Pointer to [**[]Remedy**](Remedy.md) |  | [optional] 
**Representations** | [**[]CoordinateVTwoRepresentationsInner**](CoordinateVTwoRepresentationsInner.md) | The affected versions of this vulnerability. | 

## Methods

### NewCoordinateVTwo

`func NewCoordinateVTwo(representations []CoordinateVTwoRepresentationsInner, ) *CoordinateVTwo`

NewCoordinateVTwo instantiates a new CoordinateVTwo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoordinateVTwoWithDefaults

`func NewCoordinateVTwoWithDefaults() *CoordinateVTwo`

NewCoordinateVTwoWithDefaults instantiates a new CoordinateVTwo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemedies

`func (o *CoordinateVTwo) GetRemedies() []Remedy`

GetRemedies returns the Remedies field if non-nil, zero value otherwise.

### GetRemediesOk

`func (o *CoordinateVTwo) GetRemediesOk() (*[]Remedy, bool)`

GetRemediesOk returns a tuple with the Remedies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedies

`func (o *CoordinateVTwo) SetRemedies(v []Remedy)`

SetRemedies sets Remedies field to given value.

### HasRemedies

`func (o *CoordinateVTwo) HasRemedies() bool`

HasRemedies returns a boolean if a field has been set.

### GetRepresentations

`func (o *CoordinateVTwo) GetRepresentations() []CoordinateVTwoRepresentationsInner`

GetRepresentations returns the Representations field if non-nil, zero value otherwise.

### GetRepresentationsOk

`func (o *CoordinateVTwo) GetRepresentationsOk() (*[]CoordinateVTwoRepresentationsInner, bool)`

GetRepresentationsOk returns a tuple with the Representations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentations

`func (o *CoordinateVTwo) SetRepresentations(v []CoordinateVTwoRepresentationsInner)`

SetRepresentations sets Representations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


