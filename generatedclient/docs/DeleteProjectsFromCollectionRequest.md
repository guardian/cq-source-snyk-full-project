# DeleteProjectsFromCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DeleteProjectsFromCollectionRequestDataInner**](DeleteProjectsFromCollectionRequestDataInner.md) | IDs of items to remove from a collection | 

## Methods

### NewDeleteProjectsFromCollectionRequest

`func NewDeleteProjectsFromCollectionRequest(data []DeleteProjectsFromCollectionRequestDataInner, ) *DeleteProjectsFromCollectionRequest`

NewDeleteProjectsFromCollectionRequest instantiates a new DeleteProjectsFromCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteProjectsFromCollectionRequestWithDefaults

`func NewDeleteProjectsFromCollectionRequestWithDefaults() *DeleteProjectsFromCollectionRequest`

NewDeleteProjectsFromCollectionRequestWithDefaults instantiates a new DeleteProjectsFromCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeleteProjectsFromCollectionRequest) GetData() []DeleteProjectsFromCollectionRequestDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteProjectsFromCollectionRequest) GetDataOk() (*[]DeleteProjectsFromCollectionRequestDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteProjectsFromCollectionRequest) SetData(v []DeleteProjectsFromCollectionRequestDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


