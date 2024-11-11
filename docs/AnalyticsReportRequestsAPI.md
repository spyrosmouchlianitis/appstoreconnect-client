# \AnalyticsReportRequestsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsReportRequestsCreateInstance**](AnalyticsReportRequestsAPI.md#AnalyticsReportRequestsCreateInstance) | **Post** /v1/analyticsReportRequests | 
[**AnalyticsReportRequestsDeleteInstance**](AnalyticsReportRequestsAPI.md#AnalyticsReportRequestsDeleteInstance) | **Delete** /v1/analyticsReportRequests/{id} | 
[**AnalyticsReportRequestsGetInstance**](AnalyticsReportRequestsAPI.md#AnalyticsReportRequestsGetInstance) | **Get** /v1/analyticsReportRequests/{id} | 
[**AnalyticsReportRequestsReportsGetToManyRelated**](AnalyticsReportRequestsAPI.md#AnalyticsReportRequestsReportsGetToManyRelated) | **Get** /v1/analyticsReportRequests/{id}/reports | 



## AnalyticsReportRequestsCreateInstance

> AnalyticsReportRequestResponse AnalyticsReportRequestsCreateInstance(ctx).AnalyticsReportRequestCreateRequest(analyticsReportRequestCreateRequest).Execute()



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
	analyticsReportRequestCreateRequest := *openapiclient.NewAnalyticsReportRequestCreateRequest(*openapiclient.NewAnalyticsReportRequestCreateRequestData("Type_example", *openapiclient.NewAnalyticsReportRequestCreateRequestDataAttributes("AccessType_example"), *openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // AnalyticsReportRequestCreateRequest | AnalyticsReportRequest representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportRequestsAPI.AnalyticsReportRequestsCreateInstance(context.Background()).AnalyticsReportRequestCreateRequest(analyticsReportRequestCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportRequestsAPI.AnalyticsReportRequestsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReportRequestsCreateInstance`: AnalyticsReportRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportRequestsAPI.AnalyticsReportRequestsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportRequestsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsReportRequestCreateRequest** | [**AnalyticsReportRequestCreateRequest**](AnalyticsReportRequestCreateRequest.md) | AnalyticsReportRequest representation | 

### Return type

[**AnalyticsReportRequestResponse**](AnalyticsReportRequestResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsReportRequestsDeleteInstance

> AnalyticsReportRequestsDeleteInstance(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalyticsReportRequestsAPI.AnalyticsReportRequestsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportRequestsAPI.AnalyticsReportRequestsDeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportRequestsDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsReportRequestsGetInstance

> AnalyticsReportRequestResponse AnalyticsReportRequestsGetInstance(ctx, id).FieldsAnalyticsReportRequests(fieldsAnalyticsReportRequests).FieldsAnalyticsReports(fieldsAnalyticsReports).Include(include).LimitReports(limitReports).Execute()



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
	fieldsAnalyticsReportRequests := []string{"FieldsAnalyticsReportRequests_example"} // []string | the fields to include for returned resources of type analyticsReportRequests (optional)
	fieldsAnalyticsReports := []string{"FieldsAnalyticsReports_example"} // []string | the fields to include for returned resources of type analyticsReports (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitReports := int32(56) // int32 | maximum number of related reports returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportRequestsAPI.AnalyticsReportRequestsGetInstance(context.Background(), id).FieldsAnalyticsReportRequests(fieldsAnalyticsReportRequests).FieldsAnalyticsReports(fieldsAnalyticsReports).Include(include).LimitReports(limitReports).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportRequestsAPI.AnalyticsReportRequestsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReportRequestsGetInstance`: AnalyticsReportRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportRequestsAPI.AnalyticsReportRequestsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportRequestsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAnalyticsReportRequests** | **[]string** | the fields to include for returned resources of type analyticsReportRequests | 
 **fieldsAnalyticsReports** | **[]string** | the fields to include for returned resources of type analyticsReports | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitReports** | **int32** | maximum number of related reports returned (when they are included) | 

### Return type

[**AnalyticsReportRequestResponse**](AnalyticsReportRequestResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsReportRequestsReportsGetToManyRelated

> AnalyticsReportsResponse AnalyticsReportRequestsReportsGetToManyRelated(ctx, id).FilterName(filterName).FilterCategory(filterCategory).FieldsAnalyticsReports(fieldsAnalyticsReports).Limit(limit).Execute()



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
	filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
	filterCategory := []string{"FilterCategory_example"} // []string | filter by attribute 'category' (optional)
	fieldsAnalyticsReports := []string{"FieldsAnalyticsReports_example"} // []string | the fields to include for returned resources of type analyticsReports (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportRequestsAPI.AnalyticsReportRequestsReportsGetToManyRelated(context.Background(), id).FilterName(filterName).FilterCategory(filterCategory).FieldsAnalyticsReports(fieldsAnalyticsReports).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportRequestsAPI.AnalyticsReportRequestsReportsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReportRequestsReportsGetToManyRelated`: AnalyticsReportsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportRequestsAPI.AnalyticsReportRequestsReportsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportRequestsReportsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterCategory** | **[]string** | filter by attribute &#39;category&#39; | 
 **fieldsAnalyticsReports** | **[]string** | the fields to include for returned resources of type analyticsReports | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AnalyticsReportsResponse**](AnalyticsReportsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

