# SastEnablement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**SastEnablementAttributes**](SastEnablementAttributes.md) |  | 
**Id** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewSastEnablement

`func NewSastEnablement(attributes SastEnablementAttributes, id string, type_ string, ) *SastEnablement`

NewSastEnablement instantiates a new SastEnablement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSastEnablementWithDefaults

`func NewSastEnablementWithDefaults() *SastEnablement`

NewSastEnablementWithDefaults instantiates a new SastEnablement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SastEnablement) GetAttributes() SastEnablementAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SastEnablement) GetAttributesOk() (*SastEnablementAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SastEnablement) SetAttributes(v SastEnablementAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *SastEnablement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SastEnablement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SastEnablement) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SastEnablement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SastEnablement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SastEnablement) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


