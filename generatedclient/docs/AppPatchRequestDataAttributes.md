# AppPatchRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtlSeconds** | Pointer to **float32** | The access token time to live for your app, in seconds. It only affects the newly generated access tokens, existing access token will  continue to have their previous time to live as expiration. | [optional] 
**Name** | Pointer to **string** | New name of the app to display to users during authorization flow. | [optional] 
**RedirectUris** | Pointer to **[]string** | List of allowed redirect URIs to call back after authentication. | [optional] 

## Methods

### NewAppPatchRequestDataAttributes

`func NewAppPatchRequestDataAttributes() *AppPatchRequestDataAttributes`

NewAppPatchRequestDataAttributes instantiates a new AppPatchRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPatchRequestDataAttributesWithDefaults

`func NewAppPatchRequestDataAttributesWithDefaults() *AppPatchRequestDataAttributes`

NewAppPatchRequestDataAttributesWithDefaults instantiates a new AppPatchRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTtlSeconds

`func (o *AppPatchRequestDataAttributes) GetAccessTokenTtlSeconds() float32`

GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field if non-nil, zero value otherwise.

### GetAccessTokenTtlSecondsOk

`func (o *AppPatchRequestDataAttributes) GetAccessTokenTtlSecondsOk() (*float32, bool)`

GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtlSeconds

`func (o *AppPatchRequestDataAttributes) SetAccessTokenTtlSeconds(v float32)`

SetAccessTokenTtlSeconds sets AccessTokenTtlSeconds field to given value.

### HasAccessTokenTtlSeconds

`func (o *AppPatchRequestDataAttributes) HasAccessTokenTtlSeconds() bool`

HasAccessTokenTtlSeconds returns a boolean if a field has been set.

### GetName

`func (o *AppPatchRequestDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPatchRequestDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPatchRequestDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppPatchRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRedirectUris

`func (o *AppPatchRequestDataAttributes) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *AppPatchRequestDataAttributes) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *AppPatchRequestDataAttributes) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *AppPatchRequestDataAttributes) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


