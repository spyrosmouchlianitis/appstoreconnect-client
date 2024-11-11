# AlternativeDistributionKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AlternativeDistributionKey**](AlternativeDistributionKey.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAlternativeDistributionKeysResponse

`func NewAlternativeDistributionKeysResponse(data []AlternativeDistributionKey, links PagedDocumentLinks, ) *AlternativeDistributionKeysResponse`

NewAlternativeDistributionKeysResponse instantiates a new AlternativeDistributionKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionKeysResponseWithDefaults

`func NewAlternativeDistributionKeysResponseWithDefaults() *AlternativeDistributionKeysResponse`

NewAlternativeDistributionKeysResponseWithDefaults instantiates a new AlternativeDistributionKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlternativeDistributionKeysResponse) GetData() []AlternativeDistributionKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlternativeDistributionKeysResponse) GetDataOk() (*[]AlternativeDistributionKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlternativeDistributionKeysResponse) SetData(v []AlternativeDistributionKey)`

SetData sets Data field to given value.


### GetLinks

`func (o *AlternativeDistributionKeysResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionKeysResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionKeysResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AlternativeDistributionKeysResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AlternativeDistributionKeysResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AlternativeDistributionKeysResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AlternativeDistributionKeysResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


