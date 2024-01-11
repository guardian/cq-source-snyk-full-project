# PackageMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The package’s name | [optional] 
**Namespace** | Pointer to **string** | A name prefix, such as a maven group id or docker image owner | [optional] 
**Type** | Pointer to **string** | The package type or protocol | [optional] 
**Url** | Pointer to **string** | The purl of the package | [optional] 
**Version** | Pointer to **string** | The version of the package | [optional] 

## Methods

### NewPackageMeta

`func NewPackageMeta() *PackageMeta`

NewPackageMeta instantiates a new PackageMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageMetaWithDefaults

`func NewPackageMetaWithDefaults() *PackageMeta`

NewPackageMetaWithDefaults instantiates a new PackageMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PackageMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *PackageMeta) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PackageMeta) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PackageMeta) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PackageMeta) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetType

`func (o *PackageMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PackageMeta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *PackageMeta) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PackageMeta) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PackageMeta) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PackageMeta) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersion

`func (o *PackageMeta) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageMeta) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageMeta) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackageMeta) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


