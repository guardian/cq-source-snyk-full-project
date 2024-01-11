# CollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**CollectionAttributes**](CollectionAttributes.md) |  | 
**Id** | **string** |  | 
**Meta** | [**CollectionMeta**](CollectionMeta.md) |  | 
**Relationships** | [**CollectionRelationships**](CollectionRelationships.md) |  | 
**Type** | **string** |  | 

## Methods

### NewCollectionResponse

`func NewCollectionResponse(attributes CollectionAttributes, id string, meta CollectionMeta, relationships CollectionRelationships, type_ string, ) *CollectionResponse`

NewCollectionResponse instantiates a new CollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithDefaults

`func NewCollectionResponseWithDefaults() *CollectionResponse`

NewCollectionResponseWithDefaults instantiates a new CollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CollectionResponse) GetAttributes() CollectionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CollectionResponse) GetAttributesOk() (*CollectionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CollectionResponse) SetAttributes(v CollectionAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *CollectionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMeta

`func (o *CollectionResponse) GetMeta() CollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CollectionResponse) GetMetaOk() (*CollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CollectionResponse) SetMeta(v CollectionMeta)`

SetMeta sets Meta field to given value.


### GetRelationships

`func (o *CollectionResponse) GetRelationships() CollectionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CollectionResponse) GetRelationshipsOk() (*CollectionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CollectionResponse) SetRelationships(v CollectionRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *CollectionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectionResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


