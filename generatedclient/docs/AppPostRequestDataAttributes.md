# AppPostRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtlSeconds** | Pointer to **float32** | The access token time to live for your app, in seconds. It only affects the newly generated access tokens, existing access token will  continue to have their previous time to live as expiration. | [optional] 
**Context** | [**Context**](Context.md) |  | 
**Name** | **string** | New name of the app to display to users during authorization flow. | 
**RedirectUris** | **[]string** | List of allowed redirect URIs to call back after authentication. | 
**Scopes** | **[]string** | The scopes this app is allowed to request during authorization. | 

## Methods

### NewAppPostRequestDataAttributes

`func NewAppPostRequestDataAttributes(context Context, name string, redirectUris []string, scopes []string, ) *AppPostRequestDataAttributes`

NewAppPostRequestDataAttributes instantiates a new AppPostRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPostRequestDataAttributesWithDefaults

`func NewAppPostRequestDataAttributesWithDefaults() *AppPostRequestDataAttributes`

NewAppPostRequestDataAttributesWithDefaults instantiates a new AppPostRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTtlSeconds

`func (o *AppPostRequestDataAttributes) GetAccessTokenTtlSeconds() float32`

GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field if non-nil, zero value otherwise.

### GetAccessTokenTtlSecondsOk

`func (o *AppPostRequestDataAttributes) GetAccessTokenTtlSecondsOk() (*float32, bool)`

GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtlSeconds

`func (o *AppPostRequestDataAttributes) SetAccessTokenTtlSeconds(v float32)`

SetAccessTokenTtlSeconds sets AccessTokenTtlSeconds field to given value.

### HasAccessTokenTtlSeconds

`func (o *AppPostRequestDataAttributes) HasAccessTokenTtlSeconds() bool`

HasAccessTokenTtlSeconds returns a boolean if a field has been set.

### GetContext

`func (o *AppPostRequestDataAttributes) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AppPostRequestDataAttributes) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AppPostRequestDataAttributes) SetContext(v Context)`

SetContext sets Context field to given value.


### GetName

`func (o *AppPostRequestDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPostRequestDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPostRequestDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetRedirectUris

`func (o *AppPostRequestDataAttributes) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *AppPostRequestDataAttributes) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *AppPostRequestDataAttributes) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetScopes

`func (o *AppPostRequestDataAttributes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AppPostRequestDataAttributes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AppPostRequestDataAttributes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


