# UpdateGroupServiceAccountRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**UpdateGroupServiceAccountRequestDataAttributes**](UpdateGroupServiceAccountRequestDataAttributes.md) |  | 
**Id** | **string** | The ID of the service account. Must match the id in the url path. | 
**Type** | **string** | The Resource type. | 

## Methods

### NewUpdateGroupServiceAccountRequestData

`func NewUpdateGroupServiceAccountRequestData(attributes UpdateGroupServiceAccountRequestDataAttributes, id string, type_ string, ) *UpdateGroupServiceAccountRequestData`

NewUpdateGroupServiceAccountRequestData instantiates a new UpdateGroupServiceAccountRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupServiceAccountRequestDataWithDefaults

`func NewUpdateGroupServiceAccountRequestDataWithDefaults() *UpdateGroupServiceAccountRequestData`

NewUpdateGroupServiceAccountRequestDataWithDefaults instantiates a new UpdateGroupServiceAccountRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateGroupServiceAccountRequestData) GetAttributes() UpdateGroupServiceAccountRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateGroupServiceAccountRequestData) GetAttributesOk() (*UpdateGroupServiceAccountRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateGroupServiceAccountRequestData) SetAttributes(v UpdateGroupServiceAccountRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *UpdateGroupServiceAccountRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateGroupServiceAccountRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateGroupServiceAccountRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateGroupServiceAccountRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateGroupServiceAccountRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateGroupServiceAccountRequestData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


