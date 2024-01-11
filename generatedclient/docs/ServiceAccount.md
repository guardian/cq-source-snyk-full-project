# ServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**GetManyGroupServiceAccount200ResponseDataInnerAttributes**](GetManyGroupServiceAccount200ResponseDataInnerAttributes.md) |  | 
**Id** | **string** |  | 
**Links** | Pointer to [**GetManyGroupServiceAccount200ResponseDataInnerLinks**](GetManyGroupServiceAccount200ResponseDataInnerLinks.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewServiceAccount

`func NewServiceAccount(attributes GetManyGroupServiceAccount200ResponseDataInnerAttributes, id string, type_ string, ) *ServiceAccount`

NewServiceAccount instantiates a new ServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountWithDefaults

`func NewServiceAccountWithDefaults() *ServiceAccount`

NewServiceAccountWithDefaults instantiates a new ServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ServiceAccount) GetAttributes() GetManyGroupServiceAccount200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceAccount) GetAttributesOk() (*GetManyGroupServiceAccount200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceAccount) SetAttributes(v GetManyGroupServiceAccount200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *ServiceAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccount) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *ServiceAccount) GetLinks() GetManyGroupServiceAccount200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceAccount) GetLinksOk() (*GetManyGroupServiceAccount200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceAccount) SetLinks(v GetManyGroupServiceAccount200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceAccount) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetType

`func (o *ServiceAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceAccount) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


