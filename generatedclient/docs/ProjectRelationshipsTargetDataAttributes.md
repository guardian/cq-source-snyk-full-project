# ProjectRelationshipsTargetDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The human readable name that represents this target. These are generated based on the provided properties, and the source. In the future we may support updating this value.  | [optional] 
**Url** | Pointer to **NullableString** | The URL for the resource. We do not use this as part of our representation of the identity of the target, as it can      be changed externally to Snyk We are reliant on individual integrations providing us with this value. Currently it is only provided by the CLI  | [optional] 

## Methods

### NewProjectRelationshipsTargetDataAttributes

`func NewProjectRelationshipsTargetDataAttributes() *ProjectRelationshipsTargetDataAttributes`

NewProjectRelationshipsTargetDataAttributes instantiates a new ProjectRelationshipsTargetDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRelationshipsTargetDataAttributesWithDefaults

`func NewProjectRelationshipsTargetDataAttributesWithDefaults() *ProjectRelationshipsTargetDataAttributes`

NewProjectRelationshipsTargetDataAttributesWithDefaults instantiates a new ProjectRelationshipsTargetDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ProjectRelationshipsTargetDataAttributes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProjectRelationshipsTargetDataAttributes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProjectRelationshipsTargetDataAttributes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ProjectRelationshipsTargetDataAttributes) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectRelationshipsTargetDataAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectRelationshipsTargetDataAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectRelationshipsTargetDataAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectRelationshipsTargetDataAttributes) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ProjectRelationshipsTargetDataAttributes) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ProjectRelationshipsTargetDataAttributes) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


