# VersioningSchemaSingleSelectionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSelected** | **bool** | Whether this image should be the recommendation. Only one image can be selected at a given time. Setting this as true will remove previous selection.  | 
**Type** | **string** |  | 

## Methods

### NewVersioningSchemaSingleSelectionType

`func NewVersioningSchemaSingleSelectionType(isSelected bool, type_ string, ) *VersioningSchemaSingleSelectionType`

NewVersioningSchemaSingleSelectionType instantiates a new VersioningSchemaSingleSelectionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersioningSchemaSingleSelectionTypeWithDefaults

`func NewVersioningSchemaSingleSelectionTypeWithDefaults() *VersioningSchemaSingleSelectionType`

NewVersioningSchemaSingleSelectionTypeWithDefaults instantiates a new VersioningSchemaSingleSelectionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSelected

`func (o *VersioningSchemaSingleSelectionType) GetIsSelected() bool`

GetIsSelected returns the IsSelected field if non-nil, zero value otherwise.

### GetIsSelectedOk

`func (o *VersioningSchemaSingleSelectionType) GetIsSelectedOk() (*bool, bool)`

GetIsSelectedOk returns a tuple with the IsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelected

`func (o *VersioningSchemaSingleSelectionType) SetIsSelected(v bool)`

SetIsSelected sets IsSelected field to given value.


### GetType

`func (o *VersioningSchemaSingleSelectionType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersioningSchemaSingleSelectionType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersioningSchemaSingleSelectionType) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


