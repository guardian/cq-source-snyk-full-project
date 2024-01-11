# ProjectOfCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Meta** | [**ProjectMeta**](ProjectMeta.md) |  | 
**Relationships** | [**ProjectOfCollectionRelationships**](ProjectOfCollectionRelationships.md) |  | 
**Type** | **string** |  | 

## Methods

### NewProjectOfCollection

`func NewProjectOfCollection(id string, meta ProjectMeta, relationships ProjectOfCollectionRelationships, type_ string, ) *ProjectOfCollection`

NewProjectOfCollection instantiates a new ProjectOfCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectOfCollectionWithDefaults

`func NewProjectOfCollectionWithDefaults() *ProjectOfCollection`

NewProjectOfCollectionWithDefaults instantiates a new ProjectOfCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectOfCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectOfCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectOfCollection) SetId(v string)`

SetId sets Id field to given value.


### GetMeta

`func (o *ProjectOfCollection) GetMeta() ProjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProjectOfCollection) GetMetaOk() (*ProjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProjectOfCollection) SetMeta(v ProjectMeta)`

SetMeta sets Meta field to given value.


### GetRelationships

`func (o *ProjectOfCollection) GetRelationships() ProjectOfCollectionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ProjectOfCollection) GetRelationshipsOk() (*ProjectOfCollectionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ProjectOfCollection) SetRelationships(v ProjectOfCollectionRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *ProjectOfCollection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectOfCollection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectOfCollection) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


