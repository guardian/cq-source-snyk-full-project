# Relationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**RelationshipData**](RelationshipData.md) |  | 
**Links** | [**RelatedLink**](RelatedLink.md) |  | 
**Meta** | Pointer to **map[string]interface{}** | Free-form object that may contain non-standard information. | [optional] 

## Methods

### NewRelationship

`func NewRelationship(data RelationshipData, links RelatedLink, ) *Relationship`

NewRelationship instantiates a new Relationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipWithDefaults

`func NewRelationshipWithDefaults() *Relationship`

NewRelationshipWithDefaults instantiates a new Relationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Relationship) GetData() RelationshipData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Relationship) GetDataOk() (*RelationshipData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Relationship) SetData(v RelationshipData)`

SetData sets Data field to given value.


### GetLinks

`func (o *Relationship) GetLinks() RelatedLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Relationship) GetLinksOk() (*RelatedLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Relationship) SetLinks(v RelatedLink)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *Relationship) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Relationship) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Relationship) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Relationship) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


