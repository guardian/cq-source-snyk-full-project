# UpdateServiceAccountSecretRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Operation to perform:   * &#x60;replace&#x60; - Replace existing secrets with a new generated secret.   * &#x60;create&#x60; - Add a new secret, preserving existing secrets. A maximum of to two secrets can exist at a time.   * &#x60;delete&#x60; - Remove an existing secret by value. At least one secret must remain per service account.  | 
**Secret** | Pointer to **string** | Secret to delete when using &#x60;delete&#x60; mode | [optional] 

## Methods

### NewUpdateServiceAccountSecretRequestDataAttributes

`func NewUpdateServiceAccountSecretRequestDataAttributes(mode string, ) *UpdateServiceAccountSecretRequestDataAttributes`

NewUpdateServiceAccountSecretRequestDataAttributes instantiates a new UpdateServiceAccountSecretRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceAccountSecretRequestDataAttributesWithDefaults

`func NewUpdateServiceAccountSecretRequestDataAttributesWithDefaults() *UpdateServiceAccountSecretRequestDataAttributes`

NewUpdateServiceAccountSecretRequestDataAttributesWithDefaults instantiates a new UpdateServiceAccountSecretRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *UpdateServiceAccountSecretRequestDataAttributes) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateServiceAccountSecretRequestDataAttributes) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateServiceAccountSecretRequestDataAttributes) SetMode(v string)`

SetMode sets Mode field to given value.


### GetSecret

`func (o *UpdateServiceAccountSecretRequestDataAttributes) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateServiceAccountSecretRequestDataAttributes) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateServiceAccountSecretRequestDataAttributes) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateServiceAccountSecretRequestDataAttributes) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


