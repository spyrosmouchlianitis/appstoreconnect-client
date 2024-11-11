# AnalyticsReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AnalyticsReport**](AnalyticsReport.md) |  | 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAnalyticsReportResponse

`func NewAnalyticsReportResponse(data AnalyticsReport, links DocumentLinks, ) *AnalyticsReportResponse`

NewAnalyticsReportResponse instantiates a new AnalyticsReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsReportResponseWithDefaults

`func NewAnalyticsReportResponseWithDefaults() *AnalyticsReportResponse`

NewAnalyticsReportResponseWithDefaults instantiates a new AnalyticsReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AnalyticsReportResponse) GetData() AnalyticsReport`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalyticsReportResponse) GetDataOk() (*AnalyticsReport, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalyticsReportResponse) SetData(v AnalyticsReport)`

SetData sets Data field to given value.


### GetLinks

`func (o *AnalyticsReportResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnalyticsReportResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnalyticsReportResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


