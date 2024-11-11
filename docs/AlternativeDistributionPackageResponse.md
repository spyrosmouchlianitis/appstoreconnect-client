# AlternativeDistributionPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AlternativeDistributionPackage**](AlternativeDistributionPackage.md) |  | 
**Included** | Pointer to [**[]AlternativeDistributionPackageVersion**](AlternativeDistributionPackageVersion.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAlternativeDistributionPackageResponse

`func NewAlternativeDistributionPackageResponse(data AlternativeDistributionPackage, links DocumentLinks, ) *AlternativeDistributionPackageResponse`

NewAlternativeDistributionPackageResponse instantiates a new AlternativeDistributionPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionPackageResponseWithDefaults

`func NewAlternativeDistributionPackageResponseWithDefaults() *AlternativeDistributionPackageResponse`

NewAlternativeDistributionPackageResponseWithDefaults instantiates a new AlternativeDistributionPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlternativeDistributionPackageResponse) GetData() AlternativeDistributionPackage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlternativeDistributionPackageResponse) GetDataOk() (*AlternativeDistributionPackage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlternativeDistributionPackageResponse) SetData(v AlternativeDistributionPackage)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AlternativeDistributionPackageResponse) GetIncluded() []AlternativeDistributionPackageVersion`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AlternativeDistributionPackageResponse) GetIncludedOk() (*[]AlternativeDistributionPackageVersion, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AlternativeDistributionPackageResponse) SetIncluded(v []AlternativeDistributionPackageVersion)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AlternativeDistributionPackageResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AlternativeDistributionPackageResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionPackageResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionPackageResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


