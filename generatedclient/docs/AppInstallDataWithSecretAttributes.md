# AppInstallDataWithSecretAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The oauth2 client id for the app. | 
**ClientSecret** | **string** | The oauth2 client secret for the app. This is the only time this secret will be returned, store it securely and don’t lose it. | 

## Methods

### NewAppInstallDataWithSecretAttributes

`func NewAppInstallDataWithSecretAttributes(clientId string, clientSecret string, ) *AppInstallDataWithSecretAttributes`

NewAppInstallDataWithSecretAttributes instantiates a new AppInstallDataWithSecretAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstallDataWithSecretAttributesWithDefaults

`func NewAppInstallDataWithSecretAttributesWithDefaults() *AppInstallDataWithSecretAttributes`

NewAppInstallDataWithSecretAttributesWithDefaults instantiates a new AppInstallDataWithSecretAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AppInstallDataWithSecretAttributes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AppInstallDataWithSecretAttributes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AppInstallDataWithSecretAttributes) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AppInstallDataWithSecretAttributes) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AppInstallDataWithSecretAttributes) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AppInstallDataWithSecretAttributes) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


