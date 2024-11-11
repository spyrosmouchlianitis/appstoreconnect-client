# AnalyticsReportRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AnalyticsReportRequest**](AnalyticsReportRequest.md) |  | 
**Included** | Pointer to [**[]AnalyticsReport**](AnalyticsReport.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAnalyticsReportRequestResponse

`func NewAnalyticsReportRequestResponse(data AnalyticsReportRequest, links DocumentLinks, ) *AnalyticsReportRequestResponse`

NewAnalyticsReportRequestResponse instantiates a new AnalyticsReportRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsReportRequestResponseWithDefaults

`func NewAnalyticsReportRequestResponseWithDefaults() *AnalyticsReportRequestResponse`

NewAnalyticsReportRequestResponseWithDefaults instantiates a new AnalyticsReportRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AnalyticsReportRequestResponse) GetData() AnalyticsReportRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalyticsReportRequestResponse) GetDataOk() (*AnalyticsReportRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalyticsReportRequestResponse) SetData(v AnalyticsReportRequest)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AnalyticsReportRequestResponse) GetIncluded() []AnalyticsReport`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AnalyticsReportRequestResponse) GetIncludedOk() (*[]AnalyticsReport, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AnalyticsReportRequestResponse) SetIncluded(v []AnalyticsReport)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AnalyticsReportRequestResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AnalyticsReportRequestResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnalyticsReportRequestResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnalyticsReportRequestResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


