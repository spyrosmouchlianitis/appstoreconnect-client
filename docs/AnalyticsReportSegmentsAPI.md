# \AnalyticsReportSegmentsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsReportSegmentsGetInstance**](AnalyticsReportSegmentsAPI.md#AnalyticsReportSegmentsGetInstance) | **Get** /v1/analyticsReportSegments/{id} | 



## AnalyticsReportSegmentsGetInstance

> AnalyticsReportSegmentResponse AnalyticsReportSegmentsGetInstance(ctx, id).FieldsAnalyticsReportSegments(fieldsAnalyticsReportSegments).Execute()



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
	fieldsAnalyticsReportSegments := []string{"FieldsAnalyticsReportSegments_example"} // []string | the fields to include for returned resources of type analyticsReportSegments (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportSegmentsAPI.AnalyticsReportSegmentsGetInstance(context.Background(), id).FieldsAnalyticsReportSegments(fieldsAnalyticsReportSegments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportSegmentsAPI.AnalyticsReportSegmentsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReportSegmentsGetInstance`: AnalyticsReportSegmentResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportSegmentsAPI.AnalyticsReportSegmentsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportSegmentsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAnalyticsReportSegments** | **[]string** | the fields to include for returned resources of type analyticsReportSegments | 

### Return type

[**AnalyticsReportSegmentResponse**](AnalyticsReportSegmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

