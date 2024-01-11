# SlackChannelAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Slack Channel | [optional] 
**Type** | Pointer to **string** | Channel type | [optional] 

## Methods

### NewSlackChannelAttributes

`func NewSlackChannelAttributes() *SlackChannelAttributes`

NewSlackChannelAttributes instantiates a new SlackChannelAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackChannelAttributesWithDefaults

`func NewSlackChannelAttributesWithDefaults() *SlackChannelAttributes`

NewSlackChannelAttributesWithDefaults instantiates a new SlackChannelAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SlackChannelAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlackChannelAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlackChannelAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SlackChannelAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SlackChannelAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlackChannelAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlackChannelAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SlackChannelAttributes) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


