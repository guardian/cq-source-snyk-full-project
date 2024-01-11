# GetProjectSettingsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ProjectSettingsData**](ProjectSettingsData.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**PaginatedLinks**](PaginatedLinks.md) |  | 

## Methods

### NewGetProjectSettingsCollection

`func NewGetProjectSettingsCollection(data []ProjectSettingsData, jsonapi JsonApi, links PaginatedLinks, ) *GetProjectSettingsCollection`

NewGetProjectSettingsCollection instantiates a new GetProjectSettingsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectSettingsCollectionWithDefaults

`func NewGetProjectSettingsCollectionWithDefaults() *GetProjectSettingsCollection`

NewGetProjectSettingsCollectionWithDefaults instantiates a new GetProjectSettingsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetProjectSettingsCollection) GetData() []ProjectSettingsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetProjectSettingsCollection) GetDataOk() (*[]ProjectSettingsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetProjectSettingsCollection) SetData(v []ProjectSettingsData)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *GetProjectSettingsCollection) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetProjectSettingsCollection) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetProjectSettingsCollection) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *GetProjectSettingsCollection) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetProjectSettingsCollection) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetProjectSettingsCollection) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


