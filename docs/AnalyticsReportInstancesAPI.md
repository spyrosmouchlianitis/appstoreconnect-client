# \AnalyticsReportInstancesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsReportInstancesGetInstance**](AnalyticsReportInstancesAPI.md#AnalyticsReportInstancesGetInstance) | **Get** /v1/analyticsReportInstances/{id} | 
[**AnalyticsReportInstancesSegmentsGetToManyRelated**](AnalyticsReportInstancesAPI.md#AnalyticsReportInstancesSegmentsGetToManyRelated) | **Get** /v1/analyticsReportInstances/{id}/segments | 



## AnalyticsReportInstancesGetInstance

> AnalyticsReportInstanceResponse AnalyticsReportInstancesGetInstance(ctx, id).FieldsAnalyticsReportInstances(fieldsAnalyticsReportInstances).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsAnalyticsReportInstances := []string{"FieldsAnalyticsReportInstances_example"} // []string | the fields to include for returned resources of type analyticsReportInstances (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportInstancesAPI.AnalyticsReportInstancesGetInstance(context.Background(), id).FieldsAnalyticsReportInstances(fieldsAnalyticsReportInstances).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportInstancesAPI.AnalyticsReportInstancesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReportInstancesGetInstance`: AnalyticsReportInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportInstancesAPI.AnalyticsReportInstancesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportInstancesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAnalyticsReportInstances** | **[]string** | the fields to include for returned resources of type analyticsReportInstances | 

### Return type

[**AnalyticsReportInstanceResponse**](AnalyticsReportInstanceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsReportInstancesSegmentsGetToManyRelated

> AnalyticsReportSegmentsResponse AnalyticsReportInstancesSegmentsGetToManyRelated(ctx, id).FieldsAnalyticsReportSegments(fieldsAnalyticsReportSegments).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsAnalyticsReportSegments := []string{"FieldsAnalyticsReportSegments_example"} // []string | the fields to include for returned resources of type analyticsReportSegments (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportInstancesAPI.AnalyticsReportInstancesSegmentsGetToManyRelated(context.Background(), id).FieldsAnalyticsReportSegments(fieldsAnalyticsReportSegments).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportInstancesAPI.AnalyticsReportInstancesSegmentsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReportInstancesSegmentsGetToManyRelated`: AnalyticsReportSegmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportInstancesAPI.AnalyticsReportInstancesSegmentsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportInstancesSegmentsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAnalyticsReportSegments** | **[]string** | the fields to include for returned resources of type analyticsReportSegments | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AnalyticsReportSegmentsResponse**](AnalyticsReportSegmentsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

