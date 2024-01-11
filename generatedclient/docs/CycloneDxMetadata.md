# CycloneDxMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**CycloneDxComponent**](CycloneDxComponent.md) |  | [optional] 
**Properties** | Pointer to [**[]CycloneDxProperty**](CycloneDxProperty.md) |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to [**[]CycloneDxTool**](CycloneDxTool.md) |  | [optional] 

## Methods

### NewCycloneDxMetadata

`func NewCycloneDxMetadata() *CycloneDxMetadata`

NewCycloneDxMetadata instantiates a new CycloneDxMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCycloneDxMetadataWithDefaults

`func NewCycloneDxMetadataWithDefaults() *CycloneDxMetadata`

NewCycloneDxMetadataWithDefaults instantiates a new CycloneDxMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *CycloneDxMetadata) GetComponent() CycloneDxComponent`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *CycloneDxMetadata) GetComponentOk() (*CycloneDxComponent, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *CycloneDxMetadata) SetComponent(v CycloneDxComponent)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *CycloneDxMetadata) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetProperties

`func (o *CycloneDxMetadata) GetProperties() []CycloneDxProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CycloneDxMetadata) GetPropertiesOk() (*[]CycloneDxProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CycloneDxMetadata) SetProperties(v []CycloneDxProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CycloneDxMetadata) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTimestamp

`func (o *CycloneDxMetadata) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CycloneDxMetadata) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CycloneDxMetadata) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CycloneDxMetadata) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTools

`func (o *CycloneDxMetadata) GetTools() []CycloneDxTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CycloneDxMetadata) GetToolsOk() (*[]CycloneDxTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CycloneDxMetadata) SetTools(v []CycloneDxTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CycloneDxMetadata) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


