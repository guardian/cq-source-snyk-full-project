# LinkProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | A string containing the link’s URL. | 
**Meta** | Pointer to **map[string]interface{}** | Free-form object that may contain non-standard information. | [optional] 

## Methods

### NewLinkProperty

`func NewLinkProperty(href string, ) *LinkProperty`

NewLinkProperty instantiates a new LinkProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPropertyWithDefaults

`func NewLinkPropertyWithDefaults() *LinkProperty`

NewLinkPropertyWithDefaults instantiates a new LinkProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LinkProperty) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LinkProperty) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LinkProperty) SetHref(v string)`

SetHref sets Href field to given value.


### GetMeta

`func (o *LinkProperty) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LinkProperty) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LinkProperty) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LinkProperty) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


