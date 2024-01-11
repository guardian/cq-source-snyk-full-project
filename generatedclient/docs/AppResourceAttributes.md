# AppResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtlSeconds** | **float32** | The access token time to live for your app, in seconds. It only affects the newly generated access tokens, existing access token will  continue to have their previous time to live as expiration. | 
**ClientId** | Pointer to **string** | The oauth2 client id for the app. | [optional] 
**Context** | [**Context**](Context.md) |  | 
**GrantType** | [**GrantType**](GrantType.md) |  | 
**IsConfidential** | **bool** | A boolean to indicate if an app is confidential or not as per the OAuth2 RFC. Confidential apps can securely store secrets. Examples of non-confidential apps are full web-based or CLIs. | 
**IsPublic** | **bool** | A boolean to indicate if an app is publicly available or not. | 
**Name** | **string** | New name of the app to display to users during authorization flow. | 
**OrgPublicId** | Pointer to **string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** | List of allowed redirect URIs to call back after authentication. | [optional] 
**Scopes** | **[]string** | The scopes this app is allowed to request during authorization. | 

## Methods

### NewAppResourceAttributes

`func NewAppResourceAttributes(accessTokenTtlSeconds float32, context Context, grantType GrantType, isConfidential bool, isPublic bool, name string, scopes []string, ) *AppResourceAttributes`

NewAppResourceAttributes instantiates a new AppResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppResourceAttributesWithDefaults

`func NewAppResourceAttributesWithDefaults() *AppResourceAttributes`

NewAppResourceAttributesWithDefaults instantiates a new AppResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTtlSeconds

`func (o *AppResourceAttributes) GetAccessTokenTtlSeconds() float32`

GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field if non-nil, zero value otherwise.

### GetAccessTokenTtlSecondsOk

`func (o *AppResourceAttributes) GetAccessTokenTtlSecondsOk() (*float32, bool)`

GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtlSeconds

`func (o *AppResourceAttributes) SetAccessTokenTtlSeconds(v float32)`

SetAccessTokenTtlSeconds sets AccessTokenTtlSeconds field to given value.


### GetClientId

`func (o *AppResourceAttributes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AppResourceAttributes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AppResourceAttributes) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AppResourceAttributes) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetContext

`func (o *AppResourceAttributes) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AppResourceAttributes) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AppResourceAttributes) SetContext(v Context)`

SetContext sets Context field to given value.


### GetGrantType

`func (o *AppResourceAttributes) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *AppResourceAttributes) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *AppResourceAttributes) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.


### GetIsConfidential

`func (o *AppResourceAttributes) GetIsConfidential() bool`

GetIsConfidential returns the IsConfidential field if non-nil, zero value otherwise.

### GetIsConfidentialOk

`func (o *AppResourceAttributes) GetIsConfidentialOk() (*bool, bool)`

GetIsConfidentialOk returns a tuple with the IsConfidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfidential

`func (o *AppResourceAttributes) SetIsConfidential(v bool)`

SetIsConfidential sets IsConfidential field to given value.


### GetIsPublic

`func (o *AppResourceAttributes) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AppResourceAttributes) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AppResourceAttributes) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetName

`func (o *AppResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppResourceAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetOrgPublicId

`func (o *AppResourceAttributes) GetOrgPublicId() string`

GetOrgPublicId returns the OrgPublicId field if non-nil, zero value otherwise.

### GetOrgPublicIdOk

`func (o *AppResourceAttributes) GetOrgPublicIdOk() (*string, bool)`

GetOrgPublicIdOk returns a tuple with the OrgPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPublicId

`func (o *AppResourceAttributes) SetOrgPublicId(v string)`

SetOrgPublicId sets OrgPublicId field to given value.

### HasOrgPublicId

`func (o *AppResourceAttributes) HasOrgPublicId() bool`

HasOrgPublicId returns a boolean if a field has been set.

### GetRedirectUris

`func (o *AppResourceAttributes) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *AppResourceAttributes) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *AppResourceAttributes) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *AppResourceAttributes) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *AppResourceAttributes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AppResourceAttributes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AppResourceAttributes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


