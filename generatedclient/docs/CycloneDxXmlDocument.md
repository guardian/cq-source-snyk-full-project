# CycloneDxXmlDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | [**[]CycloneDxComponent**](CycloneDxComponent.md) | A list of included software components | 
**Dependencies** | [**[]CycloneDxDependency**](CycloneDxDependency.md) |  | 
**Metadata** | [**CycloneDxMetadata**](CycloneDxMetadata.md) |  | 

## Methods

### NewCycloneDxXmlDocument

`func NewCycloneDxXmlDocument(components []CycloneDxComponent, dependencies []CycloneDxDependency, metadata CycloneDxMetadata, ) *CycloneDxXmlDocument`

NewCycloneDxXmlDocument instantiates a new CycloneDxXmlDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCycloneDxXmlDocumentWithDefaults

`func NewCycloneDxXmlDocumentWithDefaults() *CycloneDxXmlDocument`

NewCycloneDxXmlDocumentWithDefaults instantiates a new CycloneDxXmlDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *CycloneDxXmlDocument) GetComponents() []CycloneDxComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *CycloneDxXmlDocument) GetComponentsOk() (*[]CycloneDxComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *CycloneDxXmlDocument) SetComponents(v []CycloneDxComponent)`

SetComponents sets Components field to given value.


### GetDependencies

`func (o *CycloneDxXmlDocument) GetDependencies() []CycloneDxDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *CycloneDxXmlDocument) GetDependenciesOk() (*[]CycloneDxDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *CycloneDxXmlDocument) SetDependencies(v []CycloneDxDependency)`

SetDependencies sets Dependencies field to given value.


### GetMetadata

`func (o *CycloneDxXmlDocument) GetMetadata() CycloneDxMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CycloneDxXmlDocument) GetMetadataOk() (*CycloneDxMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CycloneDxXmlDocument) SetMetadata(v CycloneDxMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


