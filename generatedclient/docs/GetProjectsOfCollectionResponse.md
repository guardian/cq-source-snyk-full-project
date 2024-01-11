# GetProjectsOfCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ProjectOfCollection**](ProjectOfCollection.md) |  | [optional] 
**Jsonapi** | Pointer to [**JsonApi**](JsonApi.md) |  | [optional] 
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 

## Methods

### NewGetProjectsOfCollectionResponse

`func NewGetProjectsOfCollectionResponse() *GetProjectsOfCollectionResponse`

NewGetProjectsOfCollectionResponse instantiates a new GetProjectsOfCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectsOfCollectionResponseWithDefaults

`func NewGetProjectsOfCollectionResponseWithDefaults() *GetProjectsOfCollectionResponse`

NewGetProjectsOfCollectionResponseWithDefaults instantiates a new GetProjectsOfCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetProjectsOfCollectionResponse) GetData() []ProjectOfCollection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetProjectsOfCollectionResponse) GetDataOk() (*[]ProjectOfCollection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetProjectsOfCollectionResponse) SetData(v []ProjectOfCollection)`

SetData sets Data field to given value.

### HasData

`func (o *GetProjectsOfCollectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetJsonapi

`func (o *GetProjectsOfCollectionResponse) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetProjectsOfCollectionResponse) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetProjectsOfCollectionResponse) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.

### HasJsonapi

`func (o *GetProjectsOfCollectionResponse) HasJsonapi() bool`

HasJsonapi returns a boolean if a field has been set.

### GetLinks

`func (o *GetProjectsOfCollectionResponse) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetProjectsOfCollectionResponse) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetProjectsOfCollectionResponse) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetProjectsOfCollectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


