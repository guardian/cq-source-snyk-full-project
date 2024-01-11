# PublicAppData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**PublicAppDataAttributes**](PublicAppDataAttributes.md) |  | [optional] 
**Id** | **string** |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewPublicAppData

`func NewPublicAppData(id string, type_ string, ) *PublicAppData`

NewPublicAppData instantiates a new PublicAppData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAppDataWithDefaults

`func NewPublicAppDataWithDefaults() *PublicAppData`

NewPublicAppDataWithDefaults instantiates a new PublicAppData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PublicAppData) GetAttributes() PublicAppDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PublicAppData) GetAttributesOk() (*PublicAppDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PublicAppData) SetAttributes(v PublicAppDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PublicAppData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *PublicAppData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicAppData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicAppData) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *PublicAppData) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PublicAppData) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PublicAppData) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PublicAppData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetType

`func (o *PublicAppData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicAppData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicAppData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


