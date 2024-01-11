# GroupIacSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**GroupIacSettingsResponseAttributes**](GroupIacSettingsResponseAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID | [optional] 
**Type** | Pointer to **string** | Content type | [optional] 

## Methods

### NewGroupIacSettingsResponse

`func NewGroupIacSettingsResponse() *GroupIacSettingsResponse`

NewGroupIacSettingsResponse instantiates a new GroupIacSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupIacSettingsResponseWithDefaults

`func NewGroupIacSettingsResponseWithDefaults() *GroupIacSettingsResponse`

NewGroupIacSettingsResponseWithDefaults instantiates a new GroupIacSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *GroupIacSettingsResponse) GetAttributes() GroupIacSettingsResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GroupIacSettingsResponse) GetAttributesOk() (*GroupIacSettingsResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GroupIacSettingsResponse) SetAttributes(v GroupIacSettingsResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GroupIacSettingsResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *GroupIacSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupIacSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupIacSettingsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupIacSettingsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GroupIacSettingsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupIacSettingsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupIacSettingsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupIacSettingsResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


