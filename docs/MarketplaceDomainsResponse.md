# MarketplaceDomainsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]MarketplaceDomain**](MarketplaceDomain.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewMarketplaceDomainsResponse

`func NewMarketplaceDomainsResponse(data []MarketplaceDomain, links PagedDocumentLinks, ) *MarketplaceDomainsResponse`

NewMarketplaceDomainsResponse instantiates a new MarketplaceDomainsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDomainsResponseWithDefaults

`func NewMarketplaceDomainsResponseWithDefaults() *MarketplaceDomainsResponse`

NewMarketplaceDomainsResponseWithDefaults instantiates a new MarketplaceDomainsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MarketplaceDomainsResponse) GetData() []MarketplaceDomain`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MarketplaceDomainsResponse) GetDataOk() (*[]MarketplaceDomain, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MarketplaceDomainsResponse) SetData(v []MarketplaceDomain)`

SetData sets Data field to given value.


### GetLinks

`func (o *MarketplaceDomainsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MarketplaceDomainsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MarketplaceDomainsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *MarketplaceDomainsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MarketplaceDomainsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MarketplaceDomainsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MarketplaceDomainsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


