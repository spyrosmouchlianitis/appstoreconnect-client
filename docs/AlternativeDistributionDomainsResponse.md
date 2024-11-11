# AlternativeDistributionDomainsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AlternativeDistributionDomain**](AlternativeDistributionDomain.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAlternativeDistributionDomainsResponse

`func NewAlternativeDistributionDomainsResponse(data []AlternativeDistributionDomain, links PagedDocumentLinks, ) *AlternativeDistributionDomainsResponse`

NewAlternativeDistributionDomainsResponse instantiates a new AlternativeDistributionDomainsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionDomainsResponseWithDefaults

`func NewAlternativeDistributionDomainsResponseWithDefaults() *AlternativeDistributionDomainsResponse`

NewAlternativeDistributionDomainsResponseWithDefaults instantiates a new AlternativeDistributionDomainsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlternativeDistributionDomainsResponse) GetData() []AlternativeDistributionDomain`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlternativeDistributionDomainsResponse) GetDataOk() (*[]AlternativeDistributionDomain, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlternativeDistributionDomainsResponse) SetData(v []AlternativeDistributionDomain)`

SetData sets Data field to given value.


### GetLinks

`func (o *AlternativeDistributionDomainsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionDomainsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionDomainsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AlternativeDistributionDomainsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AlternativeDistributionDomainsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AlternativeDistributionDomainsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AlternativeDistributionDomainsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


