# ProjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildArgs** | Pointer to [**ProjectAttributesBuildArgs**](ProjectAttributesBuildArgs.md) |  | [optional] 
**BusinessCriticality** | Pointer to **[]string** |  | [optional] 
**Created** | **time.Time** | The date that the project was created on | 
**Environment** | Pointer to **[]string** |  | [optional] 
**Lifecycle** | Pointer to **[]string** |  | [optional] 
**Name** | **string** | Project name. | 
**Origin** | **string** | The origin the project was added from. | 
**ReadOnly** | **bool** | Whether the project is read-only | 
**Settings** | [**ProjectSettings**](ProjectSettings.md) |  | 
**Status** | **string** | Describes if a project is currently monitored or it is de-activated. | 
**Tags** | Pointer to [**[]PatchProjectRequestDataAttributesTagsInner**](PatchProjectRequestDataAttributesTagsInner.md) |  | [optional] 
**TargetFile** | **string** | Path within the target to identify a specific file/directory/image etc. when scanning just part  of the target, and not the entity. | 
**TargetReference** | **string** | The additional information required to resolve which revision of the resource should be scanned. | 
**TargetRuntime** | Pointer to **string** | Dotnet Target, for relevant projects | [optional] 
**Type** | **string** | The package manager of the project. | 

## Methods

### NewProjectAttributes

`func NewProjectAttributes(created time.Time, name string, origin string, readOnly bool, settings ProjectSettings, status string, targetFile string, targetReference string, type_ string, ) *ProjectAttributes`

NewProjectAttributes instantiates a new ProjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAttributesWithDefaults

`func NewProjectAttributesWithDefaults() *ProjectAttributes`

NewProjectAttributesWithDefaults instantiates a new ProjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildArgs

`func (o *ProjectAttributes) GetBuildArgs() ProjectAttributesBuildArgs`

GetBuildArgs returns the BuildArgs field if non-nil, zero value otherwise.

### GetBuildArgsOk

`func (o *ProjectAttributes) GetBuildArgsOk() (*ProjectAttributesBuildArgs, bool)`

GetBuildArgsOk returns a tuple with the BuildArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildArgs

`func (o *ProjectAttributes) SetBuildArgs(v ProjectAttributesBuildArgs)`

SetBuildArgs sets BuildArgs field to given value.

### HasBuildArgs

`func (o *ProjectAttributes) HasBuildArgs() bool`

HasBuildArgs returns a boolean if a field has been set.

### GetBusinessCriticality

`func (o *ProjectAttributes) GetBusinessCriticality() []string`

GetBusinessCriticality returns the BusinessCriticality field if non-nil, zero value otherwise.

### GetBusinessCriticalityOk

`func (o *ProjectAttributes) GetBusinessCriticalityOk() (*[]string, bool)`

GetBusinessCriticalityOk returns a tuple with the BusinessCriticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCriticality

`func (o *ProjectAttributes) SetBusinessCriticality(v []string)`

SetBusinessCriticality sets BusinessCriticality field to given value.

### HasBusinessCriticality

`func (o *ProjectAttributes) HasBusinessCriticality() bool`

HasBusinessCriticality returns a boolean if a field has been set.

### GetCreated

`func (o *ProjectAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProjectAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProjectAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEnvironment

`func (o *ProjectAttributes) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProjectAttributes) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProjectAttributes) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProjectAttributes) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLifecycle

`func (o *ProjectAttributes) GetLifecycle() []string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ProjectAttributes) GetLifecycleOk() (*[]string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ProjectAttributes) SetLifecycle(v []string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ProjectAttributes) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetName

`func (o *ProjectAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetOrigin

`func (o *ProjectAttributes) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ProjectAttributes) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ProjectAttributes) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetReadOnly

`func (o *ProjectAttributes) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ProjectAttributes) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ProjectAttributes) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetSettings

`func (o *ProjectAttributes) GetSettings() ProjectSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ProjectAttributes) GetSettingsOk() (*ProjectSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ProjectAttributes) SetSettings(v ProjectSettings)`

SetSettings sets Settings field to given value.


### GetStatus

`func (o *ProjectAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *ProjectAttributes) GetTags() []PatchProjectRequestDataAttributesTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectAttributes) GetTagsOk() (*[]PatchProjectRequestDataAttributesTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectAttributes) SetTags(v []PatchProjectRequestDataAttributesTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetFile

`func (o *ProjectAttributes) GetTargetFile() string`

GetTargetFile returns the TargetFile field if non-nil, zero value otherwise.

### GetTargetFileOk

`func (o *ProjectAttributes) GetTargetFileOk() (*string, bool)`

GetTargetFileOk returns a tuple with the TargetFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFile

`func (o *ProjectAttributes) SetTargetFile(v string)`

SetTargetFile sets TargetFile field to given value.


### GetTargetReference

`func (o *ProjectAttributes) GetTargetReference() string`

GetTargetReference returns the TargetReference field if non-nil, zero value otherwise.

### GetTargetReferenceOk

`func (o *ProjectAttributes) GetTargetReferenceOk() (*string, bool)`

GetTargetReferenceOk returns a tuple with the TargetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetReference

`func (o *ProjectAttributes) SetTargetReference(v string)`

SetTargetReference sets TargetReference field to given value.


### GetTargetRuntime

`func (o *ProjectAttributes) GetTargetRuntime() string`

GetTargetRuntime returns the TargetRuntime field if non-nil, zero value otherwise.

### GetTargetRuntimeOk

`func (o *ProjectAttributes) GetTargetRuntimeOk() (*string, bool)`

GetTargetRuntimeOk returns a tuple with the TargetRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRuntime

`func (o *ProjectAttributes) SetTargetRuntime(v string)`

SetTargetRuntime sets TargetRuntime field to given value.

### HasTargetRuntime

`func (o *ProjectAttributes) HasTargetRuntime() bool`

HasTargetRuntime returns a boolean if a field has been set.

### GetType

`func (o *ProjectAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectAttributes) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


