# SlackChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**SlackChannelAttributes**](SlackChannelAttributes.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSlackChannel

`func NewSlackChannel() *SlackChannel`

NewSlackChannel instantiates a new SlackChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackChannelWithDefaults

`func NewSlackChannelWithDefaults() *SlackChannel`

NewSlackChannelWithDefaults instantiates a new SlackChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SlackChannel) GetAttributes() SlackChannelAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SlackChannel) GetAttributesOk() (*SlackChannelAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SlackChannel) SetAttributes(v SlackChannelAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SlackChannel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *SlackChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlackChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlackChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SlackChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SlackChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlackChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlackChannel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SlackChannel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


