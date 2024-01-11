# CommonIssueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**CommonIssueModelAttributes**](CommonIssueModelAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The Snyk ID of the vulnerability. | [optional] 
**Type** | Pointer to **string** | The type of the REST resource. Always ‘issue’. | [optional] 

## Methods

### NewCommonIssueModel

`func NewCommonIssueModel() *CommonIssueModel`

NewCommonIssueModel instantiates a new CommonIssueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonIssueModelWithDefaults

`func NewCommonIssueModelWithDefaults() *CommonIssueModel`

NewCommonIssueModelWithDefaults instantiates a new CommonIssueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CommonIssueModel) GetAttributes() CommonIssueModelAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CommonIssueModel) GetAttributesOk() (*CommonIssueModelAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CommonIssueModel) SetAttributes(v CommonIssueModelAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CommonIssueModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *CommonIssueModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonIssueModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonIssueModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommonIssueModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CommonIssueModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonIssueModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonIssueModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommonIssueModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


