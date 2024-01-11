# IssuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CommonIssueModel**](CommonIssueModel.md) |  | [optional] 
**Jsonapi** | Pointer to [**JsonApi**](JsonApi.md) |  | [optional] 
**Links** | Pointer to [**PaginatedLinks**](PaginatedLinks.md) |  | [optional] 
**Meta** | Pointer to [**IssuesMeta**](IssuesMeta.md) |  | [optional] 

## Methods

### NewIssuesResponse

`func NewIssuesResponse() *IssuesResponse`

NewIssuesResponse instantiates a new IssuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesResponseWithDefaults

`func NewIssuesResponseWithDefaults() *IssuesResponse`

NewIssuesResponseWithDefaults instantiates a new IssuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IssuesResponse) GetData() []CommonIssueModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IssuesResponse) GetDataOk() (*[]CommonIssueModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IssuesResponse) SetData(v []CommonIssueModel)`

SetData sets Data field to given value.

### HasData

`func (o *IssuesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetJsonapi

`func (o *IssuesResponse) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *IssuesResponse) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *IssuesResponse) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.

### HasJsonapi

`func (o *IssuesResponse) HasJsonapi() bool`

HasJsonapi returns a boolean if a field has been set.

### GetLinks

`func (o *IssuesResponse) GetLinks() PaginatedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IssuesResponse) GetLinksOk() (*PaginatedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IssuesResponse) SetLinks(v PaginatedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IssuesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *IssuesResponse) GetMeta() IssuesMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IssuesResponse) GetMetaOk() (*IssuesMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IssuesResponse) SetMeta(v IssuesMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IssuesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


