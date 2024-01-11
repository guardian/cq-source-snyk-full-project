# Coordinate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remedies** | Pointer to [**[]Remedy**](Remedy.md) |  | [optional] 
**Representation** | Pointer to **[]string** | The affected versions of this vulnerability. | [optional] 

## Methods

### NewCoordinate

`func NewCoordinate() *Coordinate`

NewCoordinate instantiates a new Coordinate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoordinateWithDefaults

`func NewCoordinateWithDefaults() *Coordinate`

NewCoordinateWithDefaults instantiates a new Coordinate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemedies

`func (o *Coordinate) GetRemedies() []Remedy`

GetRemedies returns the Remedies field if non-nil, zero value otherwise.

### GetRemediesOk

`func (o *Coordinate) GetRemediesOk() (*[]Remedy, bool)`

GetRemediesOk returns a tuple with the Remedies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedies

`func (o *Coordinate) SetRemedies(v []Remedy)`

SetRemedies sets Remedies field to given value.

### HasRemedies

`func (o *Coordinate) HasRemedies() bool`

HasRemedies returns a boolean if a field has been set.

### GetRepresentation

`func (o *Coordinate) GetRepresentation() []string`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *Coordinate) GetRepresentationOk() (*[]string, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *Coordinate) SetRepresentation(v []string)`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *Coordinate) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


