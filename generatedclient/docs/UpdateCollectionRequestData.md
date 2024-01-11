# UpdateCollectionRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**CreateCollectionRequestDataAttributes**](CreateCollectionRequestDataAttributes.md) |  | 
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewUpdateCollectionRequestData

`func NewUpdateCollectionRequestData(attributes CreateCollectionRequestDataAttributes, type_ string, ) *UpdateCollectionRequestData`

NewUpdateCollectionRequestData instantiates a new UpdateCollectionRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCollectionRequestDataWithDefaults

`func NewUpdateCollectionRequestDataWithDefaults() *UpdateCollectionRequestData`

NewUpdateCollectionRequestDataWithDefaults instantiates a new UpdateCollectionRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateCollectionRequestData) GetAttributes() CreateCollectionRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateCollectionRequestData) GetAttributesOk() (*CreateCollectionRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateCollectionRequestData) SetAttributes(v CreateCollectionRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *UpdateCollectionRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCollectionRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCollectionRequestData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCollectionRequestData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateCollectionRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCollectionRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCollectionRequestData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


