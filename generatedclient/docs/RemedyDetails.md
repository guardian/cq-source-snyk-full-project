# RemedyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradePackage** | Pointer to **string** | A minimum version to upgrade to in order to remedy the issue. | [optional] 

## Methods

### NewRemedyDetails

`func NewRemedyDetails() *RemedyDetails`

NewRemedyDetails instantiates a new RemedyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemedyDetailsWithDefaults

`func NewRemedyDetailsWithDefaults() *RemedyDetails`

NewRemedyDetailsWithDefaults instantiates a new RemedyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradePackage

`func (o *RemedyDetails) GetUpgradePackage() string`

GetUpgradePackage returns the UpgradePackage field if non-nil, zero value otherwise.

### GetUpgradePackageOk

`func (o *RemedyDetails) GetUpgradePackageOk() (*string, bool)`

GetUpgradePackageOk returns a tuple with the UpgradePackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePackage

`func (o *RemedyDetails) SetUpgradePackage(v string)`

SetUpgradePackage sets UpgradePackage field to given value.

### HasUpgradePackage

`func (o *RemedyDetails) HasUpgradePackage() bool`

HasUpgradePackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


