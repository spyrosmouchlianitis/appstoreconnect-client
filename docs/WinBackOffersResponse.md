# WinBackOffersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WinBackOffer**](WinBackOffer.md) |  | 
**Included** | Pointer to [**[]WinBackOfferPrice**](WinBackOfferPrice.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewWinBackOffersResponse

`func NewWinBackOffersResponse(data []WinBackOffer, links PagedDocumentLinks, ) *WinBackOffersResponse`

NewWinBackOffersResponse instantiates a new WinBackOffersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWinBackOffersResponseWithDefaults

`func NewWinBackOffersResponseWithDefaults() *WinBackOffersResponse`

NewWinBackOffersResponseWithDefaults instantiates a new WinBackOffersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WinBackOffersResponse) GetData() []WinBackOffer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WinBackOffersResponse) GetDataOk() (*[]WinBackOffer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WinBackOffersResponse) SetData(v []WinBackOffer)`

SetData sets Data field to given value.


### GetIncluded

`func (o *WinBackOffersResponse) GetIncluded() []WinBackOfferPrice`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WinBackOffersResponse) GetIncludedOk() (*[]WinBackOfferPrice, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WinBackOffersResponse) SetIncluded(v []WinBackOfferPrice)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WinBackOffersResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *WinBackOffersResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WinBackOffersResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WinBackOffersResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *WinBackOffersResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WinBackOffersResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WinBackOffersResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WinBackOffersResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


