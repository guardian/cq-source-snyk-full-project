# ErrorDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]Error**](Error.md) |  | 
**Jsonapi** | [**JsonApi**](JsonApi.md) |  | 

## Methods

### NewErrorDocument

`func NewErrorDocument(errors []Error, jsonapi JsonApi, ) *ErrorDocument`

NewErrorDocument instantiates a new ErrorDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDocumentWithDefaults

`func NewErrorDocumentWithDefaults() *ErrorDocument`

NewErrorDocumentWithDefaults instantiates a new ErrorDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ErrorDocument) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorDocument) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorDocument) SetErrors(v []Error)`

SetErrors sets Errors field to given value.


### GetJsonapi

`func (o *ErrorDocument) GetJsonapi() JsonApi`

GetJsonapi returns the Jsonapi field if non-nil, zero value otherwise.

### GetJsonapiOk

`func (o *ErrorDocument) GetJsonapiOk() (*JsonApi, bool)`

GetJsonapiOk returns a tuple with the Jsonapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonapi

`func (o *ErrorDocument) SetJsonapi(v JsonApi)`

SetJsonapi sets Jsonapi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


