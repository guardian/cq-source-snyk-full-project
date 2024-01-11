# PublicAppDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**Context** | Pointer to **string** | Allow installing the app to a org/group or to a user, default tenant. | [optional] 
**Name** | **string** | New name of the app to display to users during authorization flow. | 
**Scopes** | Pointer to **[]string** | The scopes this app is allowed to request during authorization. | [optional] 

## Methods

### NewPublicAppDataAttributes

`func NewPublicAppDataAttributes(clientId string, name string, ) *PublicAppDataAttributes`

NewPublicAppDataAttributes instantiates a new PublicAppDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAppDataAttributesWithDefaults

`func NewPublicAppDataAttributesWithDefaults() *PublicAppDataAttributes`

NewPublicAppDataAttributesWithDefaults instantiates a new PublicAppDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PublicAppDataAttributes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PublicAppDataAttributes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PublicAppDataAttributes) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetContext

`func (o *PublicAppDataAttributes) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PublicAppDataAttributes) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PublicAppDataAttributes) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PublicAppDataAttributes) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetName

`func (o *PublicAppDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicAppDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicAppDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetScopes

`func (o *PublicAppDataAttributes) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PublicAppDataAttributes) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PublicAppDataAttributes) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PublicAppDataAttributes) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


