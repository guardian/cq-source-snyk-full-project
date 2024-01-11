# CreateGroupServiceAccountRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**CreateGroupServiceAccountRequestDataAttributes**](CreateGroupServiceAccountRequestDataAttributes.md) |  | 
**Type** | **string** | The Resource type. | 

## Methods

### NewCreateGroupServiceAccountRequestData

`func NewCreateGroupServiceAccountRequestData(attributes CreateGroupServiceAccountRequestDataAttributes, type_ string, ) *CreateGroupServiceAccountRequestData`

NewCreateGroupServiceAccountRequestData instantiates a new CreateGroupServiceAccountRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupServiceAccountRequestDataWithDefaults

`func NewCreateGroupServiceAccountRequestDataWithDefaults() *CreateGroupServiceAccountRequestData`

NewCreateGroupServiceAccountRequestDataWithDefaults instantiates a new CreateGroupServiceAccountRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CreateGroupServiceAccountRequestData) GetAttributes() CreateGroupServiceAccountRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateGroupServiceAccountRequestData) GetAttributesOk() (*CreateGroupServiceAccountRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateGroupServiceAccountRequestData) SetAttributes(v CreateGroupServiceAccountRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *CreateGroupServiceAccountRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateGroupServiceAccountRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateGroupServiceAccountRequestData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


