# CreateOrgServiceAccountRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**CreateOrgServiceAccountRequestDataAttributes**](CreateOrgServiceAccountRequestDataAttributes.md) |  | 
**Type** | Pointer to **string** | The Resource type. | [optional] 

## Methods

### NewCreateOrgServiceAccountRequestData

`func NewCreateOrgServiceAccountRequestData(attributes CreateOrgServiceAccountRequestDataAttributes, ) *CreateOrgServiceAccountRequestData`

NewCreateOrgServiceAccountRequestData instantiates a new CreateOrgServiceAccountRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrgServiceAccountRequestDataWithDefaults

`func NewCreateOrgServiceAccountRequestDataWithDefaults() *CreateOrgServiceAccountRequestData`

NewCreateOrgServiceAccountRequestDataWithDefaults instantiates a new CreateOrgServiceAccountRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CreateOrgServiceAccountRequestData) GetAttributes() CreateOrgServiceAccountRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateOrgServiceAccountRequestData) GetAttributesOk() (*CreateOrgServiceAccountRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateOrgServiceAccountRequestData) SetAttributes(v CreateOrgServiceAccountRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *CreateOrgServiceAccountRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrgServiceAccountRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrgServiceAccountRequestData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOrgServiceAccountRequestData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


