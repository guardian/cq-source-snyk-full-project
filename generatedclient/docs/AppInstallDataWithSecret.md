# AppInstallDataWithSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**AppInstallDataWithSecretAttributes**](AppInstallDataWithSecretAttributes.md) |  | 
**Id** | **string** |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Relationships** | Pointer to [**AppInstallDataRelationships**](AppInstallDataRelationships.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAppInstallDataWithSecret

`func NewAppInstallDataWithSecret(attributes AppInstallDataWithSecretAttributes, id string, type_ string, ) *AppInstallDataWithSecret`

NewAppInstallDataWithSecret instantiates a new AppInstallDataWithSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstallDataWithSecretWithDefaults

`func NewAppInstallDataWithSecretWithDefaults() *AppInstallDataWithSecret`

NewAppInstallDataWithSecretWithDefaults instantiates a new AppInstallDataWithSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *AppInstallDataWithSecret) GetAttributes() AppInstallDataWithSecretAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppInstallDataWithSecret) GetAttributesOk() (*AppInstallDataWithSecretAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppInstallDataWithSecret) SetAttributes(v AppInstallDataWithSecretAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *AppInstallDataWithSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppInstallDataWithSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppInstallDataWithSecret) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *AppInstallDataWithSecret) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppInstallDataWithSecret) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppInstallDataWithSecret) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppInstallDataWithSecret) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *AppInstallDataWithSecret) GetRelationships() AppInstallDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppInstallDataWithSecret) GetRelationshipsOk() (*AppInstallDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppInstallDataWithSecret) SetRelationships(v AppInstallDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppInstallDataWithSecret) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *AppInstallDataWithSecret) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppInstallDataWithSecret) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppInstallDataWithSecret) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


