# AppInstallData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**AppInstallDataAttributes**](AppInstallDataAttributes.md) |  | 
**Id** | **string** |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Relationships** | Pointer to [**AppInstallDataRelationships**](AppInstallDataRelationships.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAppInstallData

`func NewAppInstallData(attributes AppInstallDataAttributes, id string, type_ string, ) *AppInstallData`

NewAppInstallData instantiates a new AppInstallData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstallDataWithDefaults

`func NewAppInstallDataWithDefaults() *AppInstallData`

NewAppInstallDataWithDefaults instantiates a new AppInstallData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *AppInstallData) GetAttributes() AppInstallDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppInstallData) GetAttributesOk() (*AppInstallDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppInstallData) SetAttributes(v AppInstallDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *AppInstallData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppInstallData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppInstallData) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *AppInstallData) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppInstallData) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppInstallData) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppInstallData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *AppInstallData) GetRelationships() AppInstallDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppInstallData) GetRelationshipsOk() (*AppInstallDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppInstallData) SetRelationships(v AppInstallDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppInstallData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *AppInstallData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppInstallData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppInstallData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


