# CreateOrgServiceAccountRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTtlSeconds** | Pointer to **float32** | The time, in seconds, that a generated access token will be valid for. Defaults to 1 hour if unset. Only used when auth_type is one of the oauth_* variants. | [optional] 
**AuthType** | **string** | Authentication strategy for the service account:   * api_key - Regular Snyk API Key.   * oauth_client_secret - OAuth2 client_credentials grant, which returns a client secret that can be used to retrieve an access token.   * oauth_private_key_jwt - OAuth2 client_credentials grant, using private_key_jwt client_assertion as laid out in OIDC Connect Core 1.0, section 9. | 
**JwksUrl** | Pointer to **string** | A JWKs URL hosting your public keys, used to verify signed JWT requests. Must be https. Required only when auth_type is oauth_private_key_jwt. | [optional] 
**Name** | **string** | A human-friendly name for the service account. | 
**RoleId** | **string** | The ID of the role which the created service account should use. Obtained in the Snyk UI, via \&quot;Group Page\&quot; -&gt; \&quot;Settings\&quot; -&gt; \&quot;Member Roles\&quot; -&gt; \&quot;Create new Role\&quot;. Can be shared among multiple accounts. | 

## Methods

### NewCreateOrgServiceAccountRequestDataAttributes

`func NewCreateOrgServiceAccountRequestDataAttributes(authType string, name string, roleId string, ) *CreateOrgServiceAccountRequestDataAttributes`

NewCreateOrgServiceAccountRequestDataAttributes instantiates a new CreateOrgServiceAccountRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrgServiceAccountRequestDataAttributesWithDefaults

`func NewCreateOrgServiceAccountRequestDataAttributesWithDefaults() *CreateOrgServiceAccountRequestDataAttributes`

NewCreateOrgServiceAccountRequestDataAttributesWithDefaults instantiates a new CreateOrgServiceAccountRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTtlSeconds

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetAccessTokenTtlSeconds() float32`

GetAccessTokenTtlSeconds returns the AccessTokenTtlSeconds field if non-nil, zero value otherwise.

### GetAccessTokenTtlSecondsOk

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetAccessTokenTtlSecondsOk() (*float32, bool)`

GetAccessTokenTtlSecondsOk returns a tuple with the AccessTokenTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTtlSeconds

`func (o *CreateOrgServiceAccountRequestDataAttributes) SetAccessTokenTtlSeconds(v float32)`

SetAccessTokenTtlSeconds sets AccessTokenTtlSeconds field to given value.

### HasAccessTokenTtlSeconds

`func (o *CreateOrgServiceAccountRequestDataAttributes) HasAccessTokenTtlSeconds() bool`

HasAccessTokenTtlSeconds returns a boolean if a field has been set.

### GetAuthType

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateOrgServiceAccountRequestDataAttributes) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetJwksUrl

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *CreateOrgServiceAccountRequestDataAttributes) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *CreateOrgServiceAccountRequestDataAttributes) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetName

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrgServiceAccountRequestDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetRoleId

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateOrgServiceAccountRequestDataAttributes) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateOrgServiceAccountRequestDataAttributes) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


