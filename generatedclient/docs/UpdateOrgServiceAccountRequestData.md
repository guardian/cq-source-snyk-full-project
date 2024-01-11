# UpdateOrgServiceAccountRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**UpdateOrgServiceAccountRequestDataAttributes**](UpdateOrgServiceAccountRequestDataAttributes.md) |  | 
**Id** | **string** | The ID of the service account. Must match the id in the url path. | 
**Type** | **string** | The Resource type. | 

## Methods

### NewUpdateOrgServiceAccountRequestData

`func NewUpdateOrgServiceAccountRequestData(attributes UpdateOrgServiceAccountRequestDataAttributes, id string, type_ string, ) *UpdateOrgServiceAccountRequestData`

NewUpdateOrgServiceAccountRequestData instantiates a new UpdateOrgServiceAccountRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgServiceAccountRequestDataWithDefaults

`func NewUpdateOrgServiceAccountRequestDataWithDefaults() *UpdateOrgServiceAccountRequestData`

NewUpdateOrgServiceAccountRequestDataWithDefaults instantiates a new UpdateOrgServiceAccountRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateOrgServiceAccountRequestData) GetAttributes() UpdateOrgServiceAccountRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateOrgServiceAccountRequestData) GetAttributesOk() (*UpdateOrgServiceAccountRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateOrgServiceAccountRequestData) SetAttributes(v UpdateOrgServiceAccountRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *UpdateOrgServiceAccountRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrgServiceAccountRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrgServiceAccountRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateOrgServiceAccountRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOrgServiceAccountRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOrgServiceAccountRequestData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


