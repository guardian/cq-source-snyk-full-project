# VersioningSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Expression** | **string** | The regular expression used to describe the format of tags. Keep in mind that backslashes in the expression need to be escaped due to being encompassed in a JSON string.  | 
**Label** | Pointer to **string** | A customizable string that can be set for a custom versioning schema to describe its meaning. This label has no function.  | [optional] 
**IsSelected** | **bool** | Whether this image should be the recommendation. Only one image can be selected at a given time. Setting this as true will remove previous selection.  | 

## Methods

### NewVersioningSchema

`func NewVersioningSchema(type_ string, expression string, isSelected bool, ) *VersioningSchema`

NewVersioningSchema instantiates a new VersioningSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersioningSchemaWithDefaults

`func NewVersioningSchemaWithDefaults() *VersioningSchema`

NewVersioningSchemaWithDefaults instantiates a new VersioningSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VersioningSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersioningSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersioningSchema) SetType(v string)`

SetType sets Type field to given value.


### GetExpression

`func (o *VersioningSchema) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *VersioningSchema) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *VersioningSchema) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetLabel

`func (o *VersioningSchema) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VersioningSchema) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VersioningSchema) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VersioningSchema) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIsSelected

`func (o *VersioningSchema) GetIsSelected() bool`

GetIsSelected returns the IsSelected field if non-nil, zero value otherwise.

### GetIsSelectedOk

`func (o *VersioningSchema) GetIsSelectedOk() (*bool, bool)`

GetIsSelectedOk returns a tuple with the IsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelected

`func (o *VersioningSchema) SetIsSelected(v bool)`

SetIsSelected sets IsSelected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


