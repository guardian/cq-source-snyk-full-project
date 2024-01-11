# PaginatedLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | Pointer to [**LinkProperty**](LinkProperty.md) |  | [optional] 
**Last** | Pointer to [**LinkProperty**](LinkProperty.md) |  | [optional] 
**Next** | Pointer to [**LinkProperty**](LinkProperty.md) |  | [optional] 
**Prev** | Pointer to [**LinkProperty**](LinkProperty.md) |  | [optional] 
**Self** | Pointer to [**LinkProperty**](LinkProperty.md) |  | [optional] 

## Methods

### NewPaginatedLinks

`func NewPaginatedLinks() *PaginatedLinks`

NewPaginatedLinks instantiates a new PaginatedLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedLinksWithDefaults

`func NewPaginatedLinksWithDefaults() *PaginatedLinks`

NewPaginatedLinksWithDefaults instantiates a new PaginatedLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *PaginatedLinks) GetFirst() LinkProperty`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginatedLinks) GetFirstOk() (*LinkProperty, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginatedLinks) SetFirst(v LinkProperty)`

SetFirst sets First field to given value.

### HasFirst

`func (o *PaginatedLinks) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *PaginatedLinks) GetLast() LinkProperty`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginatedLinks) GetLastOk() (*LinkProperty, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginatedLinks) SetLast(v LinkProperty)`

SetLast sets Last field to given value.

### HasLast

`func (o *PaginatedLinks) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedLinks) GetNext() LinkProperty`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedLinks) GetNextOk() (*LinkProperty, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedLinks) SetNext(v LinkProperty)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *PaginatedLinks) GetPrev() LinkProperty`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PaginatedLinks) GetPrevOk() (*LinkProperty, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PaginatedLinks) SetPrev(v LinkProperty)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PaginatedLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSelf

`func (o *PaginatedLinks) GetSelf() LinkProperty`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PaginatedLinks) GetSelfOk() (*LinkProperty, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PaginatedLinks) SetSelf(v LinkProperty)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PaginatedLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


