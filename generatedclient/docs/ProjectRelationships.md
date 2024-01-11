# ProjectRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Importer** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**Organization** | [**Relationship**](Relationship.md) |  | 
**Owner** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**Target** | [**ProjectRelationshipsTarget**](ProjectRelationshipsTarget.md) |  | 

## Methods

### NewProjectRelationships

`func NewProjectRelationships(organization Relationship, target ProjectRelationshipsTarget, ) *ProjectRelationships`

NewProjectRelationships instantiates a new ProjectRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRelationshipsWithDefaults

`func NewProjectRelationshipsWithDefaults() *ProjectRelationships`

NewProjectRelationshipsWithDefaults instantiates a new ProjectRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImporter

`func (o *ProjectRelationships) GetImporter() Relationship`

GetImporter returns the Importer field if non-nil, zero value otherwise.

### GetImporterOk

`func (o *ProjectRelationships) GetImporterOk() (*Relationship, bool)`

GetImporterOk returns a tuple with the Importer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporter

`func (o *ProjectRelationships) SetImporter(v Relationship)`

SetImporter sets Importer field to given value.

### HasImporter

`func (o *ProjectRelationships) HasImporter() bool`

HasImporter returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectRelationships) GetOrganization() Relationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectRelationships) GetOrganizationOk() (*Relationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectRelationships) SetOrganization(v Relationship)`

SetOrganization sets Organization field to given value.


### GetOwner

`func (o *ProjectRelationships) GetOwner() Relationship`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ProjectRelationships) GetOwnerOk() (*Relationship, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ProjectRelationships) SetOwner(v Relationship)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ProjectRelationships) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTarget

`func (o *ProjectRelationships) GetTarget() ProjectRelationshipsTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ProjectRelationships) GetTargetOk() (*ProjectRelationshipsTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ProjectRelationships) SetTarget(v ProjectRelationshipsTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


