# AppBot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | **string** |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Relationships** | [**AppBotRelationships**](AppBotRelationships.md) |  | 
**Type** | **string** |  | 

## Methods

### NewAppBot

`func NewAppBot(id string, relationships AppBotRelationships, type_ string, ) *AppBot`

NewAppBot instantiates a new AppBot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppBotWithDefaults

`func NewAppBotWithDefaults() *AppBot`

NewAppBotWithDefaults instantiates a new AppBot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *AppBot) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppBot) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppBot) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppBot) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *AppBot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppBot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppBot) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *AppBot) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppBot) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppBot) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppBot) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *AppBot) GetRelationships() AppBotRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppBot) GetRelationshipsOk() (*AppBotRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppBot) SetRelationships(v AppBotRelationships)`

SetRelationships sets Relationships field to given value.


### GetType

`func (o *AppBot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppBot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppBot) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


