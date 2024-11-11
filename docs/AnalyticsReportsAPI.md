# \AnalyticsReportsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsReportsGetInstance**](AnalyticsReportsAPI.md#AnalyticsReportsGetInstance) | **Get** /v1/analyticsReports/{id} | 
[**AnalyticsReportsInstancesGetToManyRelated**](AnalyticsReportsAPI.md#AnalyticsReportsInstancesGetToManyRelated) | **Get** /v1/analyticsReports/{id}/instances | 



## AnalyticsReportsGetInstance

> AnalyticsReportResponse AnalyticsReportsGetInstance(ctx, id).FieldsAnalyticsReports(fieldsAnalyticsReports).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsAnalyticsReports := []string{"FieldsAnalyticsReports_example"} // []string | the fields to include for returned resources of type analyticsReports (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsReportsGetInstance(context.Background(), id).FieldsAnalyticsReports(fieldsAnalyticsReports).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsReportsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReportsGetInstance`: AnalyticsReportResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsReportsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAnalyticsReports** | **[]string** | the fields to include for returned resources of type analyticsReports | 

### Return type

[**AnalyticsReportResponse**](AnalyticsReportResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsReportsInstancesGetToManyRelated

> AnalyticsReportInstancesResponse AnalyticsReportsInstancesGetToManyRelated(ctx, id).FilterGranularity(filterGranularity).FilterProcessingDate(filterProcessingDate).FieldsAnalyticsReportInstances(fieldsAnalyticsReportInstances).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	filterGranularity := []string{"FilterGranularity_example"} // []string | filter by attribute 'granularity' (optional)
	filterProcessingDate := []string{"Inner_example"} // []string | filter by attribute 'processingDate' (optional)
	fieldsAnalyticsReportInstances := []string{"FieldsAnalyticsReportInstances_example"} // []string | the fields to include for returned resources of type analyticsReportInstances (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsReportsInstancesGetToManyRelated(context.Background(), id).FilterGranularity(filterGranularity).FilterProcessingDate(filterProcessingDate).FieldsAnalyticsReportInstances(fieldsAnalyticsReportInstances).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsReportsInstancesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReportsInstancesGetToManyRelated`: AnalyticsReportInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsReportsInstancesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportsInstancesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterGranularity** | **[]string** | filter by attribute &#39;granularity&#39; | 
 **filterProcessingDate** | **[]string** | filter by attribute &#39;processingDate&#39; | 
 **fieldsAnalyticsReportInstances** | **[]string** | the fields to include for returned resources of type analyticsReportInstances | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AnalyticsReportInstancesResponse**](AnalyticsReportInstancesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

