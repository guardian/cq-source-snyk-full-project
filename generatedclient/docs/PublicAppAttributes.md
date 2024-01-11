# PublicAppAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The oauth2 client id for the app. | 
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**Name** | **string** | New name of the app to display to users during authorization flow. | 
**Scopes** | Pointer to **[]string** | The scopes this app is allowed to request during authorization. | [optional] 

## Methods

### NewPublicAppAttributes

`func NewPublicAppAttributes(clientId string, name string, ) *PublicAppAttributes`

NewPublicAppAttributes instantiates a new PublicAppAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAppAttributesWithDefaults

`func NewPublicAppAttributesWithDefaults() *PublicAppAttributes`

NewPublicAppAttributesWithDefaults instantiates a new PublicAppAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PublicAppAttributes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PublicAppAttributes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PublicAppAttributes) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetContext

`func (o *PublicAppAttributes) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PublicAppAttributes) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PublicAppAttributes) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *PublicAppAttributes) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetName

`func (o *PublicAppAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicAppAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicAppAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetScopes

`func (o *PublicAppAttributes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PublicAppAttributes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PublicAppAttributes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PublicAppAttributes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


