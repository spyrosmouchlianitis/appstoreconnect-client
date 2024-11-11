# WinBackOfferCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**WinBackOfferCreateRequestData**](WinBackOfferCreateRequestData.md) |  | 
**Included** | Pointer to [**[]WinBackOfferPriceInlineCreate**](WinBackOfferPriceInlineCreate.md) |  | [optional] 

## Methods

### NewWinBackOfferCreateRequest

`func NewWinBackOfferCreateRequest(data WinBackOfferCreateRequestData, ) *WinBackOfferCreateRequest`

NewWinBackOfferCreateRequest instantiates a new WinBackOfferCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWinBackOfferCreateRequestWithDefaults

`func NewWinBackOfferCreateRequestWithDefaults() *WinBackOfferCreateRequest`

NewWinBackOfferCreateRequestWithDefaults instantiates a new WinBackOfferCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WinBackOfferCreateRequest) GetData() WinBackOfferCreateRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WinBackOfferCreateRequest) GetDataOk() (*WinBackOfferCreateRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WinBackOfferCreateRequest) SetData(v WinBackOfferCreateRequestData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *WinBackOfferCreateRequest) GetIncluded() []WinBackOfferPriceInlineCreate`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *WinBackOfferCreateRequest) GetIncludedOk() (*[]WinBackOfferPriceInlineCreate, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *WinBackOfferCreateRequest) SetIncluded(v []WinBackOfferPriceInlineCreate)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *WinBackOfferCreateRequest) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


