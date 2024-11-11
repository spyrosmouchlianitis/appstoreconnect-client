# BetaAppLocalizationsWithoutIncludesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaAppLocalization**](BetaAppLocalization.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaAppLocalizationsWithoutIncludesResponse

`func NewBetaAppLocalizationsWithoutIncludesResponse(data []BetaAppLocalization, links PagedDocumentLinks, ) *BetaAppLocalizationsWithoutIncludesResponse`

NewBetaAppLocalizationsWithoutIncludesResponse instantiates a new BetaAppLocalizationsWithoutIncludesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAppLocalizationsWithoutIncludesResponseWithDefaults

`func NewBetaAppLocalizationsWithoutIncludesResponseWithDefaults() *BetaAppLocalizationsWithoutIncludesResponse`

NewBetaAppLocalizationsWithoutIncludesResponseWithDefaults instantiates a new BetaAppLocalizationsWithoutIncludesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaAppLocalizationsWithoutIncludesResponse) GetData() []BetaAppLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaAppLocalizationsWithoutIncludesResponse) GetDataOk() (*[]BetaAppLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaAppLocalizationsWithoutIncludesResponse) SetData(v []BetaAppLocalization)`

SetData sets Data field to given value.


### GetLinks

`func (o *BetaAppLocalizationsWithoutIncludesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaAppLocalizationsWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaAppLocalizationsWithoutIncludesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaAppLocalizationsWithoutIncludesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaAppLocalizationsWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaAppLocalizationsWithoutIncludesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaAppLocalizationsWithoutIncludesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


