# OrgIacSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**OrgIacSettingsResponseAttributes**](OrgIacSettingsResponseAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID | [optional] 
**Type** | Pointer to **string** | Content type | [optional] 

## Methods

### NewOrgIacSettingsResponse

`func NewOrgIacSettingsResponse() *OrgIacSettingsResponse`

NewOrgIacSettingsResponse instantiates a new OrgIacSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgIacSettingsResponseWithDefaults

`func NewOrgIacSettingsResponseWithDefaults() *OrgIacSettingsResponse`

NewOrgIacSettingsResponseWithDefaults instantiates a new OrgIacSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OrgIacSettingsResponse) GetAttributes() OrgIacSettingsResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrgIacSettingsResponse) GetAttributesOk() (*OrgIacSettingsResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrgIacSettingsResponse) SetAttributes(v OrgIacSettingsResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *OrgIacSettingsResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *OrgIacSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgIacSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgIacSettingsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgIacSettingsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *OrgIacSettingsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgIacSettingsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgIacSettingsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrgIacSettingsResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


