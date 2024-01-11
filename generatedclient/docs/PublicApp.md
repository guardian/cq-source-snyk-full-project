# PublicApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**PublicAppAttributes**](PublicAppAttributes.md) |  | [optional] 
**Id** | **string** |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewPublicApp

`func NewPublicApp(id string, type_ string, ) *PublicApp`

NewPublicApp instantiates a new PublicApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAppWithDefaults

`func NewPublicAppWithDefaults() *PublicApp`

NewPublicAppWithDefaults instantiates a new PublicApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PublicApp) GetAttributes() PublicAppAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PublicApp) GetAttributesOk() (*PublicAppAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PublicApp) SetAttributes(v PublicAppAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PublicApp) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *PublicApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicApp) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *PublicApp) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PublicApp) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PublicApp) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PublicApp) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetType

`func (o *PublicApp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicApp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicApp) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


