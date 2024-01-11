# CollectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsGenerated** | Pointer to **bool** |  | [optional] 
**Name** | **string** | User-defined name of the collection | 

## Methods

### NewCollectionAttributes

`func NewCollectionAttributes(name string, ) *CollectionAttributes`

NewCollectionAttributes instantiates a new CollectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionAttributesWithDefaults

`func NewCollectionAttributesWithDefaults() *CollectionAttributes`

NewCollectionAttributesWithDefaults instantiates a new CollectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsGenerated

`func (o *CollectionAttributes) GetIsGenerated() bool`

GetIsGenerated returns the IsGenerated field if non-nil, zero value otherwise.

### GetIsGeneratedOk

`func (o *CollectionAttributes) GetIsGeneratedOk() (*bool, bool)`

GetIsGeneratedOk returns a tuple with the IsGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenerated

`func (o *CollectionAttributes) SetIsGenerated(v bool)`

SetIsGenerated sets IsGenerated field to given value.

### HasIsGenerated

`func (o *CollectionAttributes) HasIsGenerated() bool`

HasIsGenerated returns a boolean if a field has been set.

### GetName

`func (o *CollectionAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionAttributes) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


