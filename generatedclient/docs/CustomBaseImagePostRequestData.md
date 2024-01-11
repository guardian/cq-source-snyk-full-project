# CustomBaseImagePostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**CustomBaseImageAttributes**](CustomBaseImageAttributes.md) |  | 
**Type** | **string** | This should always be \&quot;custom_base_image\&quot; | 

## Methods

### NewCustomBaseImagePostRequestData

`func NewCustomBaseImagePostRequestData(attributes CustomBaseImageAttributes, type_ string, ) *CustomBaseImagePostRequestData`

NewCustomBaseImagePostRequestData instantiates a new CustomBaseImagePostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomBaseImagePostRequestDataWithDefaults

`func NewCustomBaseImagePostRequestDataWithDefaults() *CustomBaseImagePostRequestData`

NewCustomBaseImagePostRequestDataWithDefaults instantiates a new CustomBaseImagePostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CustomBaseImagePostRequestData) GetAttributes() CustomBaseImageAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomBaseImagePostRequestData) GetAttributesOk() (*CustomBaseImageAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomBaseImagePostRequestData) SetAttributes(v CustomBaseImageAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *CustomBaseImagePostRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomBaseImagePostRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomBaseImagePostRequestData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


