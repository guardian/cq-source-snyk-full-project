# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | An application-specific error code, expressed as a string value. | [optional] 
**Detail** | **string** | A human-readable explanation specific to this occurrence of the problem. | 
**Id** | Pointer to **string** | A unique identifier for this particular occurrence of the problem. | [optional] 
**Links** | Pointer to [**ErrorLink**](ErrorLink.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to [**GetAppInstallsForGroup400ResponseErrorsInnerSource**](GetAppInstallsForGroup400ResponseErrorsInnerSource.md) |  | [optional] 
**Status** | **string** | The HTTP status code applicable to this problem, expressed as a string value. | 
**Title** | Pointer to **string** | A short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization. | [optional] 

## Methods

### NewError

`func NewError(detail string, status string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *Error) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Error) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Error) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetId

`func (o *Error) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Error) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Error) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Error) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *Error) GetLinks() ErrorLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Error) GetLinksOk() (*ErrorLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Error) SetLinks(v ErrorLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Error) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *Error) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Error) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Error) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Error) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSource

`func (o *Error) GetSource() GetAppInstallsForGroup400ResponseErrorsInnerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Error) GetSourceOk() (*GetAppInstallsForGroup400ResponseErrorsInnerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Error) SetSource(v GetAppInstallsForGroup400ResponseErrorsInnerSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Error) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *Error) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *Error) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Error) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Error) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Error) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


