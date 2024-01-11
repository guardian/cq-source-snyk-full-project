# GetUserInstalledApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PublicApp**](PublicApp.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**PaginatedLinks**](PaginatedLinks.md) |  | 

## Methods

### NewGetUserInstalledApps200Response

`func NewGetUserInstalledApps200Response(data []PublicApp, jsonapi JsonApi, links PaginatedLinks, ) *GetUserInstalledApps200Response`

NewGetUserInstalledApps200Response instantiates a new GetUserInstalledApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInstalledApps200ResponseWithDefaults

`func NewGetUserInstalledApps200ResponseWithDefaults() *GetUserInstalledApps200Response`

NewGetUserInstalledApps200ResponseWithDefaults instantiates a new GetUserInstalledApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUserInstalledApps200Response) GetData() []PublicApp`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserInstalledApps200Response) GetDataOk() (*[]PublicApp, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserInstalledApps200Response) SetData(v []PublicApp)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *GetUserInstalledApps200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetUserInstalledApps200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetUserInstalledApps200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *GetUserInstalledApps200Response) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUserInstalledApps200Response) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUserInstalledApps200Response) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

