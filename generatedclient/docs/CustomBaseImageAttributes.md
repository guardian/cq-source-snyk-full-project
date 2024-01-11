# CustomBaseImageAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeInRecommendations** | **bool** | Whether this image should be recommended as a base image upgrade.  If set to true, this image could be shown as a base image upgrade to other projects. If set to false this image will never be recommended as an upgrade.  | 
**ProjectId** | **string** | The ID of the container project that the custom base image is based off of. The attributes of this custom base image are taken from the latest snapshot at the time of creation. This means that no changes should be made to the original project after the creation of the custom base image, as new snapshots created from any changes will NOT be picked up.  | 
**VersioningSchema** | Pointer to [**VersioningSchema**](VersioningSchema.md) |  | [optional] 

## Methods

### NewCustomBaseImageAttributes

`func NewCustomBaseImageAttributes(includeInRecommendations bool, projectId string, ) *CustomBaseImageAttributes`

NewCustomBaseImageAttributes instantiates a new CustomBaseImageAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomBaseImageAttributesWithDefaults

`func NewCustomBaseImageAttributesWithDefaults() *CustomBaseImageAttributes`

NewCustomBaseImageAttributesWithDefaults instantiates a new CustomBaseImageAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeInRecommendations

`func (o *CustomBaseImageAttributes) GetIncludeInRecommendations() bool`

GetIncludeInRecommendations returns the IncludeInRecommendations field if non-nil, zero value otherwise.

### GetIncludeInRecommendationsOk

`func (o *CustomBaseImageAttributes) GetIncludeInRecommendationsOk() (*bool, bool)`

GetIncludeInRecommendationsOk returns a tuple with the IncludeInRecommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInRecommendations

`func (o *CustomBaseImageAttributes) SetIncludeInRecommendations(v bool)`

SetIncludeInRecommendations sets IncludeInRecommendations field to given value.


### GetProjectId

`func (o *CustomBaseImageAttributes) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CustomBaseImageAttributes) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CustomBaseImageAttributes) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersioningSchema

`func (o *CustomBaseImageAttributes) GetVersioningSchema() VersioningSchema`

GetVersioningSchema returns the VersioningSchema field if non-nil, zero value otherwise.

### GetVersioningSchemaOk

`func (o *CustomBaseImageAttributes) GetVersioningSchemaOk() (*VersioningSchema, bool)`

GetVersioningSchemaOk returns a tuple with the VersioningSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioningSchema

`func (o *CustomBaseImageAttributes) SetVersioningSchema(v VersioningSchema)`

SetVersioningSchema sets VersioningSchema field to given value.

### HasVersioningSchema

`func (o *CustomBaseImageAttributes) HasVersioningSchema() bool`

HasVersioningSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


