# GetManyGroupServiceAccount200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtlSeconds** | Pointer to **float32** | The time, in seconds, that a generated access token will be valid for. Defaults to 1hr if unset. Only provided when auth_type is oauth_private_key_jwt. | [optional] 
**ApiKey** | Pointer to **string** | The Snyk API Key for this service account. Only returned on creation, and only when auth_type is api_key. | [optional] 
**AuthType** | **string** | The authentication strategy for the service account:   * api_key - Regular Snyk API Key.   * oauth_client_secret - OAuth2 client_credentials grant, which returns a client secret that can be used to retrieve an access token.   * oauth_private_key_jwt - OAuth2 client_credentials grant, using private_key_jwt client_assertion as laid out OIDC Connect Core 1.0, section 9. | 
**ClientId** | Pointer to **string** | The service account&#39;s attached client-id. Used to request an access-token. Only provided when auth_type is oauth_private_key_jwt. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret used for obtaining access tokens. Only sent on creation of new service accounts and cannot be retrieved thereafter. Only provided when auth_type is oauth_client_secret. | [optional] 
**JwksUrl** | Pointer to **string** | A JWKs URL used to verify signed JWT requests against. Must be https. Only provided when auth_type is oauth_private_key_jwt. | [optional] 
**Level** | Pointer to **string** | The level of access for the service account:   * Group - the service account was created at the Group level.   * Org - the service account was created at the Org level. | [optional] 
**Name** | **string** | A human-friendly name of the service account. | 
**RoleId** | **string** | The ID of the role which the Service Account is associated with. | 

## Methods

### NewGetManyGroupServiceAccount200ResponseDataInnerAttributes

`func NewGetManyGroupServiceAccount200ResponseDataInnerAttributes(authType string, name string, roleId string, ) *GetManyGroupServiceAccount200ResponseDataInnerAttributes`

NewGetManyGroupServiceAccount200ResponseDataInnerAttributes instantiates a new GetManyGroupServiceAccount200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManyGroupServiceAccount200ResponseDataInnerAttributesWithDefaults

`func NewGetManyGroupServiceAccount200ResponseDataInnerAttributesWithDefaults() *GetManyGroupServiceAccount200ResponseDataInnerAttributes`

NewGetManyGroupServiceAccount200ResponseDataInnerAttributesWithDefaults instantiates a new GetManyGroupServiceAccount200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTtlSeconds

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetAccessTokenTtlSeconds() float32`

GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field if non-nil, zero value otherwise.

### GetAccessTokenTtlSecondsOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetAccessTokenTtlSecondsOk() (*float32, bool)`

GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtlSeconds

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetAccessTokenTtlSeconds(v float32)`

SetAccessTokenTtlSeconds sets AccessTokenTtlSeconds field to given value.

### HasAccessTokenTtlSeconds

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasAccessTokenTtlSeconds() bool`

HasAccessTokenTtlSeconds returns a boolean if a field has been set.

### GetApiKey

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAuthType

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetClientId

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetJwksUrl

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetLevel

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetRoleId

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *GetManyGroupServiceAccount200ResponseDataInnerAttributes) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


