# SettingsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**SettingsAttributes**](SettingsAttributes.md) |  | 
**Type** | **string** |  | 

## Methods

### NewSettingsRequestData

`func NewSettingsRequestData(attributes SettingsAttributes, type_ string, ) *SettingsRequestData`

NewSettingsRequestData instantiates a new SettingsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsRequestDataWithDefaults

`func NewSettingsRequestDataWithDefaults() *SettingsRequestData`

NewSettingsRequestDataWithDefaults instantiates a new SettingsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SettingsRequestData) GetAttributes() SettingsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SettingsRequestData) GetAttributesOk() (*SettingsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SettingsRequestData) SetAttributes(v SettingsAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *SettingsRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SettingsRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SettingsRequestData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


