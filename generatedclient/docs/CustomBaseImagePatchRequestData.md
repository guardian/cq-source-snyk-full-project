# CustomBaseImagePatchRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**CustomBaseImagePatchRequestDataAttributes**](CustomBaseImagePatchRequestDataAttributes.md) |  | 
**Id** | Pointer to **string** | The ID of the custom base image that should be updated. (Same one used in the URI) | [optional] 
**Type** | **string** | This should always be \&quot;custom_base_image\&quot; | 

## Methods

### NewCustomBaseImagePatchRequestData

`func NewCustomBaseImagePatchRequestData(attributes CustomBaseImagePatchRequestDataAttributes, type_ string, ) *CustomBaseImagePatchRequestData`

NewCustomBaseImagePatchRequestData instantiates a new CustomBaseImagePatchRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomBaseImagePatchRequestDataWithDefaults

`func NewCustomBaseImagePatchRequestDataWithDefaults() *CustomBaseImagePatchRequestData`

NewCustomBaseImagePatchRequestDataWithDefaults instantiates a new CustomBaseImagePatchRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CustomBaseImagePatchRequestData) GetAttributes() CustomBaseImagePatchRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomBaseImagePatchRequestData) GetAttributesOk() (*CustomBaseImagePatchRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomBaseImagePatchRequestData) SetAttributes(v CustomBaseImagePatchRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *CustomBaseImagePatchRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomBaseImagePatchRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomBaseImagePatchRequestData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomBaseImagePatchRequestData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomBaseImagePatchRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomBaseImagePatchRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomBaseImagePatchRequestData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


