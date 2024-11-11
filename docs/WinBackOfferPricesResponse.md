# WinBackOfferPricesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WinBackOfferPrice**](WinBackOfferPrice.md) |  | 
**Included** | Pointer to [**[]SubscriptionOfferCodePricesResponseIncludedInner**](SubscriptionOfferCodePricesResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewWinBackOfferPricesResponse

`func NewWinBackOfferPricesResponse(data []WinBackOfferPrice, links PagedDocumentLinks, ) *WinBackOfferPricesResponse`

NewWinBackOfferPricesResponse instantiates a new WinBackOfferPricesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWinBackOfferPricesResponseWithDefaults

`func NewWinBackOfferPricesResponseWithDefaults() *WinBackOfferPricesResponse`

NewWinBackOfferPricesResponseWithDefaults instantiates a new WinBackOfferPricesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WinBackOfferPricesResponse) GetData() []WinBackOfferPrice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WinBackOfferPricesResponse) GetDataOk() (*[]WinBackOfferPrice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WinBackOfferPricesResponse) SetData(v []WinBackOfferPrice)`

SetData sets Data field to given value.


### GetIncluded

`func (o *WinBackOfferPricesResponse) GetIncluded() []SubscriptionOfferCodePricesResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WinBackOfferPricesResponse) GetIncludedOk() (*[]SubscriptionOfferCodePricesResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WinBackOfferPricesResponse) SetIncluded(v []SubscriptionOfferCodePricesResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WinBackOfferPricesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *WinBackOfferPricesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WinBackOfferPricesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WinBackOfferPricesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *WinBackOfferPricesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WinBackOfferPricesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WinBackOfferPricesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WinBackOfferPricesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


