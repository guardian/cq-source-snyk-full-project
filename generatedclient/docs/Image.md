# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**ImageAttributes**](ImageAttributes.md) |  | 
**Id** | **string** |  | 
**Relationships** | Pointer to [**ImageRelationships**](ImageRelationships.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewImage

`func NewImage(attributes ImageAttributes, id string, type_ string, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Image) GetAttributes() ImageAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Image) GetAttributesOk() (*ImageAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Image) SetAttributes(v ImageAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *Image) GetRelationships() ImageRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Image) GetRelationshipsOk() (*ImageRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Image) SetRelationships(v ImageRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Image) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *Image) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Image) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Image) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


