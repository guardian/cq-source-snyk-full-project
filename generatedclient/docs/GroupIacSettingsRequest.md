# GroupIacSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**GroupIacSettingsRequestAttributes**](GroupIacSettingsRequestAttributes.md) |  | 
**Type** | **string** | Content type | 

## Methods

### NewGroupIacSettingsRequest

`func NewGroupIacSettingsRequest(attributes GroupIacSettingsRequestAttributes, type_ string, ) *GroupIacSettingsRequest`

NewGroupIacSettingsRequest instantiates a new GroupIacSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupIacSettingsRequestWithDefaults

`func NewGroupIacSettingsRequestWithDefaults() *GroupIacSettingsRequest`

NewGroupIacSettingsRequestWithDefaults instantiates a new GroupIacSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *GroupIacSettingsRequest) GetAttributes() GroupIacSettingsRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GroupIacSettingsRequest) GetAttributesOk() (*GroupIacSettingsRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GroupIacSettingsRequest) SetAttributes(v GroupIacSettingsRequestAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *GroupIacSettingsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupIacSettingsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupIacSettingsRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


