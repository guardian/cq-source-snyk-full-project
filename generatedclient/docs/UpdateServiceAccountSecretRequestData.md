# UpdateServiceAccountSecretRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**UpdateServiceAccountSecretRequestDataAttributes**](UpdateServiceAccountSecretRequestDataAttributes.md) |  | 
**Type** | **string** | The Resource type. | 

## Methods

### NewUpdateServiceAccountSecretRequestData

`func NewUpdateServiceAccountSecretRequestData(attributes UpdateServiceAccountSecretRequestDataAttributes, type_ string, ) *UpdateServiceAccountSecretRequestData`

NewUpdateServiceAccountSecretRequestData instantiates a new UpdateServiceAccountSecretRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceAccountSecretRequestDataWithDefaults

`func NewUpdateServiceAccountSecretRequestDataWithDefaults() *UpdateServiceAccountSecretRequestData`

NewUpdateServiceAccountSecretRequestDataWithDefaults instantiates a new UpdateServiceAccountSecretRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateServiceAccountSecretRequestData) GetAttributes() UpdateServiceAccountSecretRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateServiceAccountSecretRequestData) GetAttributesOk() (*UpdateServiceAccountSecretRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateServiceAccountSecretRequestData) SetAttributes(v UpdateServiceAccountSecretRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *UpdateServiceAccountSecretRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateServiceAccountSecretRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateServiceAccountSecretRequestData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


