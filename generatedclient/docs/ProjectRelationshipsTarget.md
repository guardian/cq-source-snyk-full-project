# ProjectRelationshipsTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ProjectRelationshipsTargetData**](ProjectRelationshipsTargetData.md) |  | 
**Links** | [**RelatedLink**](RelatedLink.md) |  | 
**Meta** | Pointer to **map[string]interface{}** | Free-form object that may contain non-standard information. | [optional] 

## Methods

### NewProjectRelationshipsTarget

`func NewProjectRelationshipsTarget(data ProjectRelationshipsTargetData, links RelatedLink, ) *ProjectRelationshipsTarget`

NewProjectRelationshipsTarget instantiates a new ProjectRelationshipsTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRelationshipsTargetWithDefaults

`func NewProjectRelationshipsTargetWithDefaults() *ProjectRelationshipsTarget`

NewProjectRelationshipsTargetWithDefaults instantiates a new ProjectRelationshipsTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProjectRelationshipsTarget) GetData() ProjectRelationshipsTargetData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProjectRelationshipsTarget) GetDataOk() (*ProjectRelationshipsTargetData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProjectRelationshipsTarget) SetData(v ProjectRelationshipsTargetData)`

SetData sets Data field to given value.


### GetLinks

`func (o *ProjectRelationshipsTarget) GetLinks() RelatedLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectRelationshipsTarget) GetLinksOk() (*RelatedLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectRelationshipsTarget) SetLinks(v RelatedLink)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *ProjectRelationshipsTarget) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProjectRelationshipsTarget) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProjectRelationshipsTarget) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProjectRelationshipsTarget) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


