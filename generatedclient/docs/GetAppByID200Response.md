# GetAppByID200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AppData**](AppData.md) |  | [optional] 
**Jsonapi** | Pointer to [**JsonApi**](JsonApi.md) |  | [optional] 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 

## Methods

### NewGetAppByID200Response

`func NewGetAppByID200Response() *GetAppByID200Response`

NewGetAppByID200Response instantiates a new GetAppByID200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppByID200ResponseWithDefaults

`func NewGetAppByID200ResponseWithDefaults() *GetAppByID200Response`

NewGetAppByID200ResponseWithDefaults instantiates a new GetAppByID200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAppByID200Response) GetData() AppData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAppByID200Response) GetDataOk() (*AppData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAppByID200Response) SetData(v AppData)`

SetData sets Data field to given value.

### HasData

`func (o *GetAppByID200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetJsonapi

`func (o *GetAppByID200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *GetAppByID200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *GetAppByID200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.

### HasJsonapi

`func (o *GetAppByID200Response) HasJsonapi() bool`

HasJsonapi returns a boolean if a field has been set.

### GetLinks

`func (o *GetAppByID200Response) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAppByID200Response) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAppByID200Response) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAppByID200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


