# GetAppInstallsForGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppInstallData**](AppInstallData.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**PaginatedLinks**](PaginatedLinks.md) |  | 

## Methods

### NewGetAppInstallsForGroup200Response

`func NewGetAppInstallsForGroup200Response(data []AppInstallData, jsonapi JsonApi, links PaginatedLinks, ) *GetAppInstallsForGroup200Response`

NewGetAppInstallsForGroup200Response instantiates a new GetAppInstallsForGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppInstallsForGroup200ResponseWithDefaults

`func NewGetAppInstallsForGroup200ResponseWithDefaults() *GetAppInstallsForGroup200Response`

NewGetAppInstallsForGroup200ResponseWithDefaults instantiates a new GetAppInstallsForGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAppInstallsForGroup200Response) GetData() []AppInstallData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAppInstallsForGroup200Response) GetDataOk() (*[]AppInstallData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAppInstallsForGroup200Response) SetData(v []AppInstallData)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *GetAppInstallsForGroup200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetAppInstallsForGroup200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetAppInstallsForGroup200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *GetAppInstallsForGroup200Response) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAppInstallsForGroup200Response) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAppInstallsForGroup200Response) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


