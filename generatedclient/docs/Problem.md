# Problem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisclosedAt** | Pointer to **time.Time** | When this problem was disclosed to the public. | [optional] 
**DiscoveredAt** | Pointer to **time.Time** | When this problem was first discovered. | [optional] 
**Id** | **string** |  | 
**Source** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** | When this problem was last updated. | [optional] 
**Url** | Pointer to **string** | An optional URL for this problem. | [optional] 

## Methods

### NewProblem

`func NewProblem(id string, source string, ) *Problem`

NewProblem instantiates a new Problem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemWithDefaults

`func NewProblemWithDefaults() *Problem`

NewProblemWithDefaults instantiates a new Problem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclosedAt

`func (o *Problem) GetDisclosedAt() time.Time`

GetDisclosedAt returns the DisclosedAt field if non-nil, zero value otherwise.

### GetDisclosedAtOk

`func (o *Problem) GetDisclosedAtOk() (*time.Time, bool)`

GetDisclosedAtOk returns a tuple with the DisclosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosedAt

`func (o *Problem) SetDisclosedAt(v time.Time)`

SetDisclosedAt sets DisclosedAt field to given value.

### HasDisclosedAt

`func (o *Problem) HasDisclosedAt() bool`

HasDisclosedAt returns a boolean if a field has been set.

### GetDiscoveredAt

`func (o *Problem) GetDiscoveredAt() time.Time`

GetDiscoveredAt returns the DiscoveredAt field if non-nil, zero value otherwise.

### GetDiscoveredAtOk

`func (o *Problem) GetDiscoveredAtOk() (*time.Time, bool)`

GetDiscoveredAtOk returns a tuple with the DiscoveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredAt

`func (o *Problem) SetDiscoveredAt(v time.Time)`

SetDiscoveredAt sets DiscoveredAt field to given value.

### HasDiscoveredAt

`func (o *Problem) HasDiscoveredAt() bool`

HasDiscoveredAt returns a boolean if a field has been set.

### GetId

`func (o *Problem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Problem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Problem) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *Problem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Problem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Problem) SetSource(v string)`

SetSource sets Source field to given value.


### GetUpdatedAt

`func (o *Problem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Problem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Problem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Problem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Problem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Problem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Problem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Problem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


