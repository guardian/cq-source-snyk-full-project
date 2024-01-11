# OrgIacSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**OrgIacSettingsRequestAttributes**](OrgIacSettingsRequestAttributes.md) |  | 
**Type** | **string** | Content type | 

## Methods

### NewOrgIacSettingsRequest

`func NewOrgIacSettingsRequest(attributes OrgIacSettingsRequestAttributes, type_ string, ) *OrgIacSettingsRequest`

NewOrgIacSettingsRequest instantiates a new OrgIacSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgIacSettingsRequestWithDefaults

`func NewOrgIacSettingsRequestWithDefaults() *OrgIacSettingsRequest`

NewOrgIacSettingsRequestWithDefaults instantiates a new OrgIacSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OrgIacSettingsRequest) GetAttributes() OrgIacSettingsRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrgIacSettingsRequest) GetAttributesOk() (*OrgIacSettingsRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrgIacSettingsRequest) SetAttributes(v OrgIacSettingsRequestAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *OrgIacSettingsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgIacSettingsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgIacSettingsRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


