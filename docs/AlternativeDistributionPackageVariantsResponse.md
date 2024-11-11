# AlternativeDistributionPackageVariantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AlternativeDistributionPackageVariant**](AlternativeDistributionPackageVariant.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAlternativeDistributionPackageVariantsResponse

`func NewAlternativeDistributionPackageVariantsResponse(data []AlternativeDistributionPackageVariant, links PagedDocumentLinks, ) *AlternativeDistributionPackageVariantsResponse`

NewAlternativeDistributionPackageVariantsResponse instantiates a new AlternativeDistributionPackageVariantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionPackageVariantsResponseWithDefaults

`func NewAlternativeDistributionPackageVariantsResponseWithDefaults() *AlternativeDistributionPackageVariantsResponse`

NewAlternativeDistributionPackageVariantsResponseWithDefaults instantiates a new AlternativeDistributionPackageVariantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlternativeDistributionPackageVariantsResponse) GetData() []AlternativeDistributionPackageVariant`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlternativeDistributionPackageVariantsResponse) GetDataOk() (*[]AlternativeDistributionPackageVariant, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlternativeDistributionPackageVariantsResponse) SetData(v []AlternativeDistributionPackageVariant)`

SetData sets Data field to given value.


### GetLinks

`func (o *AlternativeDistributionPackageVariantsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionPackageVariantsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionPackageVariantsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AlternativeDistributionPackageVariantsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AlternativeDistributionPackageVariantsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AlternativeDistributionPackageVariantsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AlternativeDistributionPackageVariantsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


