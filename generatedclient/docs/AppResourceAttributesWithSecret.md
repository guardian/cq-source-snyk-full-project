# AppResourceAttributesWithSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtlSeconds** | **float32** | The access token time to live for your app, in seconds. It only affects the newly generated access tokens, existing access token will  continue to have their previous time to live as expiration. | 
**ClientId** | **string** | The oauth2 client id for the app. | 
**ClientSecret** | **string** | The oauth2 client secret for the app. This is the only time this secret will be returned, store it securely and don’t lose it. | 
**Context** | [**Context**](Context.md) |  | 
**GrantType** | [**GrantType**](GrantType.md) |  | 
**IsConfidential** | **bool** | A boolean to indicate if an app is confidential or not as per the OAuth2 RFC. Confidential apps can securely store secrets. Examples of non-confidential apps are full web-based or CLIs. | 
**IsPublic** | **bool** | A boolean to indicate if an app is publicly available or not. | 
**Name** | **string** | New name of the app to display to users during authorization flow. | 
**OrgPublicId** | Pointer to **string** |  | [optional] 
**RedirectUris** | **[]string** | List of allowed redirect URIs to call back after authentication. | 
**Scopes** | **[]string** | The scopes this app is allowed to request during authorization. | 

## Methods

### NewAppResourceAttributesWithSecret

`func NewAppResourceAttributesWithSecret(accessTokenTtlSeconds float32, clientId string, clientSecret string, context Context, grantType GrantType, isConfidential bool, isPublic bool, name string, redirectUris []string, scopes []string, ) *AppResourceAttributesWithSecret`

NewAppResourceAttributesWithSecret instantiates a new AppResourceAttributesWithSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppResourceAttributesWithSecretWithDefaults

`func NewAppResourceAttributesWithSecretWithDefaults() *AppResourceAttributesWithSecret`

NewAppResourceAttributesWithSecretWithDefaults instantiates a new AppResourceAttributesWithSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTtlSeconds

`func (o *AppResourceAttributesWithSecret) GetAccessTokenTtlSeconds() float32`

GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field if non-nil, zero value otherwise.

### GetAccessTokenTtlSecondsOk

`func (o *AppResourceAttributesWithSecret) GetAccessTokenTtlSecondsOk() (*float32, bool)`

GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtlSeconds

`func (o *AppResourceAttributesWithSecret) SetAccessTokenTtlSeconds(v float32)`

SetAccessTokenTtlSeconds sets AccessTokenTtlSeconds field to given value.


### GetClientId

`func (o *AppResourceAttributesWithSecret) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AppResourceAttributesWithSecret) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AppResourceAttributesWithSecret) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AppResourceAttributesWithSecret) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AppResourceAttributesWithSecret) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AppResourceAttributesWithSecret) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetContext

`func (o *AppResourceAttributesWithSecret) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AppResourceAttributesWithSecret) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AppResourceAttributesWithSecret) SetContext(v Context)`

SetContext sets Context field to given value.


### GetGrantType

`func (o *AppResourceAttributesWithSecret) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *AppResourceAttributesWithSecret) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *AppResourceAttributesWithSecret) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.


### GetIsConfidential

`func (o *AppResourceAttributesWithSecret) GetIsConfidential() bool`

GetIsConfidential returns the IsConfidential field if non-nil, zero value otherwise.

### GetIsConfidentialOk

`func (o *AppResourceAttributesWithSecret) GetIsConfidentialOk() (*bool, bool)`

GetIsConfidentialOk returns a tuple with the IsConfidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfidential

`func (o *AppResourceAttributesWithSecret) SetIsConfidential(v bool)`

SetIsConfidential sets IsConfidential field to given value.


### GetIsPublic

`func (o *AppResourceAttributesWithSecret) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AppResourceAttributesWithSecret) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AppResourceAttributesWithSecret) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetName

`func (o *AppResourceAttributesWithSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppResourceAttributesWithSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppResourceAttributesWithSecret) SetName(v string)`

SetName sets Name field to given value.


### GetOrgPublicId

`func (o *AppResourceAttributesWithSecret) GetOrgPublicId() string`

GetOrgPublicId returns the OrgPublicId field if non-nil, zero value otherwise.

### GetOrgPublicIdOk

`func (o *AppResourceAttributesWithSecret) GetOrgPublicIdOk() (*string, bool)`

GetOrgPublicIdOk returns a tuple with the OrgPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPublicId

`func (o *AppResourceAttributesWithSecret) SetOrgPublicId(v string)`

SetOrgPublicId sets OrgPublicId field to given value.

### HasOrgPublicId

`func (o *AppResourceAttributesWithSecret) HasOrgPublicId() bool`

HasOrgPublicId returns a boolean if a field has been set.

### GetRedirectUris

`func (o *AppResourceAttributesWithSecret) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *AppResourceAttributesWithSecret) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *AppResourceAttributesWithSecret) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetScopes

`func (o *AppResourceAttributesWithSecret) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AppResourceAttributesWithSecret) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AppResourceAttributesWithSecret) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


