# WinBackOfferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**WinBackOffer**](WinBackOffer.md) |  | 
**Included** | Pointer to [**[]WinBackOfferPrice**](WinBackOfferPrice.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewWinBackOfferResponse

`func NewWinBackOfferResponse(data WinBackOffer, links DocumentLinks, ) *WinBackOfferResponse`

NewWinBackOfferResponse instantiates a new WinBackOfferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWinBackOfferResponseWithDefaults

`func NewWinBackOfferResponseWithDefaults() *WinBackOfferResponse`

NewWinBackOfferResponseWithDefaults instantiates a new WinBackOfferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WinBackOfferResponse) GetData() WinBackOffer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WinBackOfferResponse) GetDataOk() (*WinBackOffer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WinBackOfferResponse) SetData(v WinBackOffer)`

SetData sets Data field to given value.


### GetIncluded

`func (o *WinBackOfferResponse) GetIncluded() []WinBackOfferPrice`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WinBackOfferResponse) GetIncludedOk() (*[]WinBackOfferPrice, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WinBackOfferResponse) SetIncluded(v []WinBackOfferPrice)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WinBackOfferResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *WinBackOfferResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WinBackOfferResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WinBackOfferResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


