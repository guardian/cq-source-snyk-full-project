# CreateGroupAppInstall201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppInstallWithClient**](AppInstallWithClient.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | [**PaginatedLinks**](PaginatedLinks.md) |  | 

## Methods

### NewCreateGroupAppInstall201Response

`func NewCreateGroupAppInstall201Response(data AppInstallWithClient, jsonapi JsonApi, links PaginatedLinks, ) *CreateGroupAppInstall201Response`

NewCreateGroupAppInstall201Response instantiates a new CreateGroupAppInstall201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupAppInstall201ResponseWithDefaults

`func NewCreateGroupAppInstall201ResponseWithDefaults() *CreateGroupAppInstall201Response`

NewCreateGroupAppInstall201ResponseWithDefaults instantiates a new CreateGroupAppInstall201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateGroupAppInstall201Response) GetData() AppInstallWithClient`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateGroupAppInstall201Response) GetDataOk() (*AppInstallWithClient, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateGroupAppInstall201Response) SetData(v AppInstallWithClient)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *CreateGroupAppInstall201Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *CreateGroupAppInstall201Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *CreateGroupAppInstall201Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *CreateGroupAppInstall201Response) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateGroupAppInstall201Response) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateGroupAppInstall201Response) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


