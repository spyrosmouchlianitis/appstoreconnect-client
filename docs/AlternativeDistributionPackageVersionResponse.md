# AlternativeDistributionPackageVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AlternativeDistributionPackageVersion**](AlternativeDistributionPackageVersion.md) |  | 
**Included** | Pointer to [**[]AlternativeDistributionPackageVersionsResponseIncludedInner**](AlternativeDistributionPackageVersionsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAlternativeDistributionPackageVersionResponse

`func NewAlternativeDistributionPackageVersionResponse(data AlternativeDistributionPackageVersion, links DocumentLinks, ) *AlternativeDistributionPackageVersionResponse`

NewAlternativeDistributionPackageVersionResponse instantiates a new AlternativeDistributionPackageVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionPackageVersionResponseWithDefaults

`func NewAlternativeDistributionPackageVersionResponseWithDefaults() *AlternativeDistributionPackageVersionResponse`

NewAlternativeDistributionPackageVersionResponseWithDefaults instantiates a new AlternativeDistributionPackageVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlternativeDistributionPackageVersionResponse) GetData() AlternativeDistributionPackageVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlternativeDistributionPackageVersionResponse) GetDataOk() (*AlternativeDistributionPackageVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlternativeDistributionPackageVersionResponse) SetData(v AlternativeDistributionPackageVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AlternativeDistributionPackageVersionResponse) GetIncluded() []AlternativeDistributionPackageVersionsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AlternativeDistributionPackageVersionResponse) GetIncludedOk() (*[]AlternativeDistributionPackageVersionsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AlternativeDistributionPackageVersionResponse) SetIncluded(v []AlternativeDistributionPackageVersionsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AlternativeDistributionPackageVersionResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AlternativeDistributionPackageVersionResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionPackageVersionResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionPackageVersionResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


