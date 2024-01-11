# VersioningSchemaCustomType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | The regular expression used to describe the format of tags. Keep in mind that backslashes in the expression need to be escaped due to being encompassed in a JSON string.  | 
**Label** | Pointer to **string** | A customizable string that can be set for a custom versioning schema to describe its meaning. This label has no function.  | [optional] 
**Type** | **string** |  | 

## Methods

### NewVersioningSchemaCustomType

`func NewVersioningSchemaCustomType(expression string, type_ string, ) *VersioningSchemaCustomType`

NewVersioningSchemaCustomType instantiates a new VersioningSchemaCustomType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersioningSchemaCustomTypeWithDefaults

`func NewVersioningSchemaCustomTypeWithDefaults() *VersioningSchemaCustomType`

NewVersioningSchemaCustomTypeWithDefaults instantiates a new VersioningSchemaCustomType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *VersioningSchemaCustomType) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *VersioningSchemaCustomType) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *VersioningSchemaCustomType) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetLabel

`func (o *VersioningSchemaCustomType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VersioningSchemaCustomType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VersioningSchemaCustomType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VersioningSchemaCustomType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *VersioningSchemaCustomType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersioningSchemaCustomType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersioningSchemaCustomType) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


