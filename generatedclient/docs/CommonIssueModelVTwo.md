# CommonIssueModelVTwo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**CommonIssueModelVTwoAttributes**](CommonIssueModelVTwoAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The Snyk ID of the vulnerability. | [optional] 
**Type** | Pointer to **string** | The type of the REST resource. Always ‘issue’. | [optional] 

## Methods

### NewCommonIssueModelVTwo

`func NewCommonIssueModelVTwo() *CommonIssueModelVTwo`

NewCommonIssueModelVTwo instantiates a new CommonIssueModelVTwo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonIssueModelVTwoWithDefaults

`func NewCommonIssueModelVTwoWithDefaults() *CommonIssueModelVTwo`

NewCommonIssueModelVTwoWithDefaults instantiates a new CommonIssueModelVTwo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CommonIssueModelVTwo) GetAttributes() CommonIssueModelVTwoAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CommonIssueModelVTwo) GetAttributesOk() (*CommonIssueModelVTwoAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CommonIssueModelVTwo) SetAttributes(v CommonIssueModelVTwoAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CommonIssueModelVTwo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *CommonIssueModelVTwo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonIssueModelVTwo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonIssueModelVTwo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommonIssueModelVTwo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CommonIssueModelVTwo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonIssueModelVTwo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonIssueModelVTwo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommonIssueModelVTwo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


