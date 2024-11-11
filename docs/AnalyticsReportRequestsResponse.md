# AnalyticsReportRequestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AnalyticsReportRequest**](AnalyticsReportRequest.md) |  | 
**Included** | Pointer to [**[]AnalyticsReport**](AnalyticsReport.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAnalyticsReportRequestsResponse

`func NewAnalyticsReportRequestsResponse(data []AnalyticsReportRequest, links PagedDocumentLinks, ) *AnalyticsReportRequestsResponse`

NewAnalyticsReportRequestsResponse instantiates a new AnalyticsReportRequestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsReportRequestsResponseWithDefaults

`func NewAnalyticsReportRequestsResponseWithDefaults() *AnalyticsReportRequestsResponse`

NewAnalyticsReportRequestsResponseWithDefaults instantiates a new AnalyticsReportRequestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AnalyticsReportRequestsResponse) GetData() []AnalyticsReportRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalyticsReportRequestsResponse) GetDataOk() (*[]AnalyticsReportRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalyticsReportRequestsResponse) SetData(v []AnalyticsReportRequest)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AnalyticsReportRequestsResponse) GetIncluded() []AnalyticsReport`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AnalyticsReportRequestsResponse) GetIncludedOk() (*[]AnalyticsReport, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AnalyticsReportRequestsResponse) SetIncluded(v []AnalyticsReport)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AnalyticsReportRequestsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AnalyticsReportRequestsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnalyticsReportRequestsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnalyticsReportRequestsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AnalyticsReportRequestsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AnalyticsReportRequestsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AnalyticsReportRequestsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AnalyticsReportRequestsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


