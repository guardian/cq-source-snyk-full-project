# GetAppInstallsForGroup400ResponseErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | An application-specific error code, expressed as a string value. | [optional] 
**Detail** | **string** | A human-readable explanation specific to this occurrence of the problem. | 
**Id** | Pointer to **string** | A unique identifier for this particular occurrence of the problem. | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to [**GetAppInstallsForGroup400ResponseErrorsInnerSource**](GetAppInstallsForGroup400ResponseErrorsInnerSource.md) |  | [optional] 
**Status** | **string** | The HTTP status code applicable to this problem, expressed as a string value. | 
**Title** | Pointer to **string** | A short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization. | [optional] 

## Methods

### NewGetAppInstallsForGroup400ResponseErrorsInner

`func NewGetAppInstallsForGroup400ResponseErrorsInner(detail string, status string, ) *GetAppInstallsForGroup400ResponseErrorsInner`

NewGetAppInstallsForGroup400ResponseErrorsInner instantiates a new GetAppInstallsForGroup400ResponseErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppInstallsForGroup400ResponseErrorsInnerWithDefaults

`func NewGetAppInstallsForGroup400ResponseErrorsInnerWithDefaults() *GetAppInstallsForGroup400ResponseErrorsInner`

NewGetAppInstallsForGroup400ResponseErrorsInnerWithDefaults instantiates a new GetAppInstallsForGroup400ResponseErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetId

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMeta

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSource

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetSource() GetAppInstallsForGroup400ResponseErrorsInnerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetSourceOk() (*GetAppInstallsForGroup400ResponseErrorsInnerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetSource(v GetAppInstallsForGroup400ResponseErrorsInnerSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetAppInstallsForGroup400ResponseErrorsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


