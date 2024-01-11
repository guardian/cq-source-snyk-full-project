# CreateCollection201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**CollectionAttributes**](CollectionAttributes.md) |  | 
**Id** | **string** |  | 
**Relationships** | [**CollectionRelationships**](CollectionRelationships.md) |  | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateCollection201ResponseData

`func NewCreateCollection201ResponseData(attributes CollectionAttributes, id string, relationships CollectionRelationships, ) *CreateCollection201ResponseData`

NewCreateCollection201ResponseData instantiates a new CreateCollection201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollection201ResponseDataWithDefaults

`func NewCreateCollection201ResponseDataWithDefaults() *CreateCollection201ResponseData`

NewCreateCollection201ResponseDataWithDefaults instantiates a new CreateCollection201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CreateCollection201ResponseData) GetAttributes() CollectionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateCollection201ResponseData) GetAttributesOk() (*CollectionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateCollection201ResponseData) SetAttributes(v CollectionAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *CreateCollection201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCollection201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCollection201ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *CreateCollection201ResponseData) GetRelationships() CollectionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CreateCollection201ResponseData) GetRelationshipsOk() (*CollectionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CreateCollection201ResponseData) SetRelationships(v CollectionRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *CreateCollection201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCollection201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCollection201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateCollection201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


