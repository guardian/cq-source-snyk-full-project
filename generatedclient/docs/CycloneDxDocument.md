# CycloneDxDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BomFormat** | **string** |  | 
**Components** | Pointer to [**[]CycloneDxComponent**](CycloneDxComponent.md) | A list of included software components | [optional] 
**Dependencies** | [**[]CycloneDxDependency**](CycloneDxDependency.md) |  | 
**Metadata** | [**CycloneDxMetadata**](CycloneDxMetadata.md) |  | 
**SpecVersion** | **string** |  | 
**Version** | **int32** |  | 

## Methods

### NewCycloneDxDocument

`func NewCycloneDxDocument(bomFormat string, dependencies []CycloneDxDependency, metadata CycloneDxMetadata, specVersion string, version int32, ) *CycloneDxDocument`

NewCycloneDxDocument instantiates a new CycloneDxDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCycloneDxDocumentWithDefaults

`func NewCycloneDxDocumentWithDefaults() *CycloneDxDocument`

NewCycloneDxDocumentWithDefaults instantiates a new CycloneDxDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBomFormat

`func (o *CycloneDxDocument) GetBomFormat() string`

GetBomFormat returns the BomFormat field if non-nil, zero value otherwise.

### GetBomFormatOk

`func (o *CycloneDxDocument) GetBomFormatOk() (*string, bool)`

GetBomFormatOk returns a tuple with the BomFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomFormat

`func (o *CycloneDxDocument) SetBomFormat(v string)`

SetBomFormat sets BomFormat field to given value.


### GetComponents

`func (o *CycloneDxDocument) GetComponents() []CycloneDxComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *CycloneDxDocument) GetComponentsOk() (*[]CycloneDxComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *CycloneDxDocument) SetComponents(v []CycloneDxComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *CycloneDxDocument) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDependencies

`func (o *CycloneDxDocument) GetDependencies() []CycloneDxDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *CycloneDxDocument) GetDependenciesOk() (*[]CycloneDxDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *CycloneDxDocument) SetDependencies(v []CycloneDxDependency)`

SetDependencies sets Dependencies field to given value.


### GetMetadata

`func (o *CycloneDxDocument) GetMetadata() CycloneDxMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CycloneDxDocument) GetMetadataOk() (*CycloneDxMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CycloneDxDocument) SetMetadata(v CycloneDxMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpecVersion

`func (o *CycloneDxDocument) GetSpecVersion() string`

GetSpecVersion returns the SpecVersion field if non-nil, zero value otherwise.

### GetSpecVersionOk

`func (o *CycloneDxDocument) GetSpecVersionOk() (*string, bool)`

GetSpecVersionOk returns a tuple with the SpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecVersion

`func (o *CycloneDxDocument) SetSpecVersion(v string)`

SetSpecVersion sets SpecVersion field to given value.


### GetVersion

`func (o *CycloneDxDocument) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CycloneDxDocument) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CycloneDxDocument) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


