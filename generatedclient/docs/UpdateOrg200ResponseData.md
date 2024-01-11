# UpdateOrg200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**OrgAttributes**](OrgAttributes.md) |  | [optional] 
**Id** | **string** |  | 
**Relationships** | Pointer to [**OrgRelationships**](OrgRelationships.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewUpdateOrg200ResponseData

`func NewUpdateOrg200ResponseData(id string, type_ string, ) *UpdateOrg200ResponseData`

NewUpdateOrg200ResponseData instantiates a new UpdateOrg200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrg200ResponseDataWithDefaults

`func NewUpdateOrg200ResponseDataWithDefaults() *UpdateOrg200ResponseData`

NewUpdateOrg200ResponseDataWithDefaults instantiates a new UpdateOrg200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateOrg200ResponseData) GetAttributes() OrgAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateOrg200ResponseData) GetAttributesOk() (*OrgAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateOrg200ResponseData) SetAttributes(v OrgAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateOrg200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *UpdateOrg200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrg200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrg200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *UpdateOrg200ResponseData) GetRelationships() OrgRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UpdateOrg200ResponseData) GetRelationshipsOk() (*OrgRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UpdateOrg200ResponseData) SetRelationships(v OrgRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UpdateOrg200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *UpdateOrg200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOrg200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOrg200ResponseData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


