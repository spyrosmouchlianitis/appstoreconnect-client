# AlternativeDistributionPackageVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AlternativeDistributionPackageVersion**](AlternativeDistributionPackageVersion.md) |  | 
**Included** | Pointer to [**[]AlternativeDistributionPackageVersionsResponseIncludedInner**](AlternativeDistributionPackageVersionsResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAlternativeDistributionPackageVersionsResponse

`func NewAlternativeDistributionPackageVersionsResponse(data []AlternativeDistributionPackageVersion, links PagedDocumentLinks, ) *AlternativeDistributionPackageVersionsResponse`

NewAlternativeDistributionPackageVersionsResponse instantiates a new AlternativeDistributionPackageVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionPackageVersionsResponseWithDefaults

`func NewAlternativeDistributionPackageVersionsResponseWithDefaults() *AlternativeDistributionPackageVersionsResponse`

NewAlternativeDistributionPackageVersionsResponseWithDefaults instantiates a new AlternativeDistributionPackageVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlternativeDistributionPackageVersionsResponse) GetData() []AlternativeDistributionPackageVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlternativeDistributionPackageVersionsResponse) GetDataOk() (*[]AlternativeDistributionPackageVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlternativeDistributionPackageVersionsResponse) SetData(v []AlternativeDistributionPackageVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AlternativeDistributionPackageVersionsResponse) GetIncluded() []AlternativeDistributionPackageVersionsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AlternativeDistributionPackageVersionsResponse) GetIncludedOk() (*[]AlternativeDistributionPackageVersionsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AlternativeDistributionPackageVersionsResponse) SetIncluded(v []AlternativeDistributionPackageVersionsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AlternativeDistributionPackageVersionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AlternativeDistributionPackageVersionsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionPackageVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionPackageVersionsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AlternativeDistributionPackageVersionsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AlternativeDistributionPackageVersionsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AlternativeDistributionPackageVersionsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AlternativeDistributionPackageVersionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


