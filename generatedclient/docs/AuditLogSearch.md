# AuditLogSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AuditLogSearchItemsInner**](AuditLogSearchItemsInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAuditLogSearch

`func NewAuditLogSearch() *AuditLogSearch`

NewAuditLogSearch instantiates a new AuditLogSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogSearchWithDefaults

`func NewAuditLogSearchWithDefaults() *AuditLogSearch`

NewAuditLogSearchWithDefaults instantiates a new AuditLogSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AuditLogSearch) GetItems() []AuditLogSearchItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AuditLogSearch) GetItemsOk() (*[]AuditLogSearchItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AuditLogSearch) SetItems(v []AuditLogSearchItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *AuditLogSearch) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetType

`func (o *AuditLogSearch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditLogSearch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditLogSearch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditLogSearch) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


