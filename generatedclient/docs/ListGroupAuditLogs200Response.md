# ListGroupAuditLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AuditLogSearch**](AuditLogSearch.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewListGroupAuditLogs200Response

`func NewListGroupAuditLogs200Response(data AuditLogSearch, jsonapi JsonApi, ) *ListGroupAuditLogs200Response`

NewListGroupAuditLogs200Response instantiates a new ListGroupAuditLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGroupAuditLogs200ResponseWithDefaults

`func NewListGroupAuditLogs200ResponseWithDefaults() *ListGroupAuditLogs200Response`

NewListGroupAuditLogs200ResponseWithDefaults instantiates a new ListGroupAuditLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListGroupAuditLogs200Response) GetData() AuditLogSearch`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListGroupAuditLogs200Response) GetDataOk() (*AuditLogSearch, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListGroupAuditLogs200Response) SetData(v AuditLogSearch)`

SetData sets Data field to given value.


### GetJsonapi

`func (o *ListGroupAuditLogs200Response) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *ListGroupAuditLogs200Response) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *ListGroupAuditLogs200Response) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.


### GetLinks

`func (o *ListGroupAuditLogs200Response) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListGroupAuditLogs200Response) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListGroupAuditLogs200Response) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListGroupAuditLogs200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


