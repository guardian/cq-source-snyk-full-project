# CreateGroupServiceAccountRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtlSeconds** | Pointer to **float32** | The time, in seconds, that a generated access token will be valid for. Defaults to 1 hour if unset. Only used when auth_type is one of the oauth_* variants. | [optional] 
**AuthType** | **string** | Authentication strategy for the service account:   * api_key - Regular Snyk API Key.   * oauth_client_secret - OAuth2 client_credentials grant, which returns a client secret that can be used to retrieve an access token.   * oauth_private_key_jwt - OAuth2 client_credentials grant, using private_key_jwt client_assertion as laid out in OIDC Connect Core 1.0, section 9. | 
**JwksUrl** | Pointer to **string** | A JWKs URL hosting your public keys, used to verify signed JWT requests. Must be https. Required only when auth_type is oauth_private_key_jwt. | [optional] 
**Name** | **string** | A human-friendly name for the service account. | 
**RoleId** | **string** | The ID of the role which the created service account should use. | 

## Methods

### NewCreateGroupServiceAccountRequestDataAttributes

`func NewCreateGroupServiceAccountRequestDataAttributes(authType string, name string, roleId string, ) *CreateGroupServiceAccountRequestDataAttributes`

NewCreateGroupServiceAccountRequestDataAttributes instantiates a new CreateGroupServiceAccountRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupServiceAccountRequestDataAttributesWithDefaults

`func NewCreateGroupServiceAccountRequestDataAttributesWithDefaults() *CreateGroupServiceAccountRequestDataAttributes`

NewCreateGroupServiceAccountRequestDataAttributesWithDefaults instantiates a new CreateGroupServiceAccountRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTtlSeconds

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetAccessTokenTtlSeconds() float32`

GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field if non-nil, zero value otherwise.

### GetAccessTokenTtlSecondsOk

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetAccessTokenTtlSecondsOk() (*float32, bool)`

GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtlSeconds

`func (o *CreateGroupServiceAccountRequestDataAttributes) SetAccessTokenTtlSeconds(v float32)`

SetAccessTokenTtlSeconds sets AccessTokenTtlSeconds field to given value.

### HasAccessTokenTtlSeconds

`func (o *CreateGroupServiceAccountRequestDataAttributes) HasAccessTokenTtlSeconds() bool`

HasAccessTokenTtlSeconds returns a boolean if a field has been set.

### GetAuthType

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateGroupServiceAccountRequestDataAttributes) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetJwksUrl

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *CreateGroupServiceAccountRequestDataAttributes) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *CreateGroupServiceAccountRequestDataAttributes) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetName

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupServiceAccountRequestDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetRoleId

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateGroupServiceAccountRequestDataAttributes) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateGroupServiceAccountRequestDataAttributes) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


