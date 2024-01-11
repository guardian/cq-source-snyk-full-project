# BulkPackageUrlsRequestBodyDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purls** | **[]string** | An array of Package URLs (purl). Supported purl types are apk, cargo, cocoapods, composer, deb, gem, generic, golang, hex, maven, npm, nuget, pub, pypi, rpm, and swift. A version for the package is also required. | 

## Methods

### NewBulkPackageUrlsRequestBodyDataAttributes

`func NewBulkPackageUrlsRequestBodyDataAttributes(purls []string, ) *BulkPackageUrlsRequestBodyDataAttributes`

NewBulkPackageUrlsRequestBodyDataAttributes instantiates a new BulkPackageUrlsRequestBodyDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPackageUrlsRequestBodyDataAttributesWithDefaults

`func NewBulkPackageUrlsRequestBodyDataAttributesWithDefaults() *BulkPackageUrlsRequestBodyDataAttributes`

NewBulkPackageUrlsRequestBodyDataAttributesWithDefaults instantiates a new BulkPackageUrlsRequestBodyDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurls

`func (o *BulkPackageUrlsRequestBodyDataAttributes) GetPurls() []string`

GetPurls returns the Purls field if non-nil, zero value otherwise.

### GetPurlsOk

`func (o *BulkPackageUrlsRequestBodyDataAttributes) GetPurlsOk() (*[]string, bool)`

GetPurlsOk returns a tuple with the Purls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurls

`func (o *BulkPackageUrlsRequestBodyDataAttributes) SetPurls(v []string)`

SetPurls sets Purls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


