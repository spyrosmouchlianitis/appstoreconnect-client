# \ReviewSubmissionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReviewSubmissionsCreateInstance**](ReviewSubmissionsAPI.md#ReviewSubmissionsCreateInstance) | **Post** /v1/reviewSubmissions | 
[**ReviewSubmissionsGetCollection**](ReviewSubmissionsAPI.md#ReviewSubmissionsGetCollection) | **Get** /v1/reviewSubmissions | 
[**ReviewSubmissionsGetInstance**](ReviewSubmissionsAPI.md#ReviewSubmissionsGetInstance) | **Get** /v1/reviewSubmissions/{id} | 
[**ReviewSubmissionsItemsGetToManyRelated**](ReviewSubmissionsAPI.md#ReviewSubmissionsItemsGetToManyRelated) | **Get** /v1/reviewSubmissions/{id}/items | 
[**ReviewSubmissionsUpdateInstance**](ReviewSubmissionsAPI.md#ReviewSubmissionsUpdateInstance) | **Patch** /v1/reviewSubmissions/{id} | 



## ReviewSubmissionsCreateInstance

> ReviewSubmissionResponse ReviewSubmissionsCreateInstance(ctx).ReviewSubmissionCreateRequest(reviewSubmissionCreateRequest).Execute()



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
	reviewSubmissionCreateRequest := *openapiclient.NewReviewSubmissionCreateRequest(*openapiclient.NewReviewSubmissionCreateRequestData("Type_example", *openapiclient.NewReviewSubmissionCreateRequestDataAttributes(openapiclient.Platform("IOS")), *openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // ReviewSubmissionCreateRequest | ReviewSubmission representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewSubmissionsAPI.ReviewSubmissionsCreateInstance(context.Background()).ReviewSubmissionCreateRequest(reviewSubmissionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewSubmissionsAPI.ReviewSubmissionsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewSubmissionsCreateInstance`: ReviewSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `ReviewSubmissionsAPI.ReviewSubmissionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReviewSubmissionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewSubmissionCreateRequest** | [**ReviewSubmissionCreateRequest**](ReviewSubmissionCreateRequest.md) | ReviewSubmission representation | 

### Return type

[**ReviewSubmissionResponse**](ReviewSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewSubmissionsGetCollection

> ReviewSubmissionsResponse ReviewSubmissionsGetCollection(ctx).FilterApp(filterApp).FilterPlatform(filterPlatform).FilterState(filterState).FieldsReviewSubmissions(fieldsReviewSubmissions).FieldsReviewSubmissionItems(fieldsReviewSubmissionItems).Limit(limit).Include(include).LimitItems(limitItems).Execute()



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
	filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app'
	filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
	filterState := []string{"FilterState_example"} // []string | filter by attribute 'state' (optional)
	fieldsReviewSubmissions := []string{"FieldsReviewSubmissions_example"} // []string | the fields to include for returned resources of type reviewSubmissions (optional)
	fieldsReviewSubmissionItems := []string{"FieldsReviewSubmissionItems_example"} // []string | the fields to include for returned resources of type reviewSubmissionItems (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitItems := int32(56) // int32 | maximum number of related items returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewSubmissionsAPI.ReviewSubmissionsGetCollection(context.Background()).FilterApp(filterApp).FilterPlatform(filterPlatform).FilterState(filterState).FieldsReviewSubmissions(fieldsReviewSubmissions).FieldsReviewSubmissionItems(fieldsReviewSubmissionItems).Limit(limit).Include(include).LimitItems(limitItems).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewSubmissionsAPI.ReviewSubmissionsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewSubmissionsGetCollection`: ReviewSubmissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReviewSubmissionsAPI.ReviewSubmissionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReviewSubmissionsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterState** | **[]string** | filter by attribute &#39;state&#39; | 
 **fieldsReviewSubmissions** | **[]string** | the fields to include for returned resources of type reviewSubmissions | 
 **fieldsReviewSubmissionItems** | **[]string** | the fields to include for returned resources of type reviewSubmissionItems | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitItems** | **int32** | maximum number of related items returned (when they are included) | 

### Return type

[**ReviewSubmissionsResponse**](ReviewSubmissionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewSubmissionsGetInstance

> ReviewSubmissionResponse ReviewSubmissionsGetInstance(ctx, id).FieldsReviewSubmissions(fieldsReviewSubmissions).FieldsReviewSubmissionItems(fieldsReviewSubmissionItems).Include(include).LimitItems(limitItems).Execute()



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
	fieldsReviewSubmissions := []string{"FieldsReviewSubmissions_example"} // []string | the fields to include for returned resources of type reviewSubmissions (optional)
	fieldsReviewSubmissionItems := []string{"FieldsReviewSubmissionItems_example"} // []string | the fields to include for returned resources of type reviewSubmissionItems (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitItems := int32(56) // int32 | maximum number of related items returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewSubmissionsAPI.ReviewSubmissionsGetInstance(context.Background(), id).FieldsReviewSubmissions(fieldsReviewSubmissions).FieldsReviewSubmissionItems(fieldsReviewSubmissionItems).Include(include).LimitItems(limitItems).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewSubmissionsAPI.ReviewSubmissionsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewSubmissionsGetInstance`: ReviewSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `ReviewSubmissionsAPI.ReviewSubmissionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReviewSubmissionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsReviewSubmissions** | **[]string** | the fields to include for returned resources of type reviewSubmissions | 
 **fieldsReviewSubmissionItems** | **[]string** | the fields to include for returned resources of type reviewSubmissionItems | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitItems** | **int32** | maximum number of related items returned (when they are included) | 

### Return type

[**ReviewSubmissionResponse**](ReviewSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewSubmissionsItemsGetToManyRelated

> ReviewSubmissionItemsResponse ReviewSubmissionsItemsGetToManyRelated(ctx, id).FieldsReviewSubmissionItems(fieldsReviewSubmissionItems).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAppEvents(fieldsAppEvents).Limit(limit).Include(include).Execute()



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
	fieldsReviewSubmissionItems := []string{"FieldsReviewSubmissionItems_example"} // []string | the fields to include for returned resources of type reviewSubmissionItems (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsAppCustomProductPageVersions := []string{"FieldsAppCustomProductPageVersions_example"} // []string | the fields to include for returned resources of type appCustomProductPageVersions (optional)
	fieldsAppStoreVersionExperiments := []string{"FieldsAppStoreVersionExperiments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperiments (optional)
	fieldsAppEvents := []string{"FieldsAppEvents_example"} // []string | the fields to include for returned resources of type appEvents (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewSubmissionsAPI.ReviewSubmissionsItemsGetToManyRelated(context.Background(), id).FieldsReviewSubmissionItems(fieldsReviewSubmissionItems).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAppEvents(fieldsAppEvents).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewSubmissionsAPI.ReviewSubmissionsItemsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewSubmissionsItemsGetToManyRelated`: ReviewSubmissionItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReviewSubmissionsAPI.ReviewSubmissionsItemsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReviewSubmissionsItemsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsReviewSubmissionItems** | **[]string** | the fields to include for returned resources of type reviewSubmissionItems | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppCustomProductPageVersions** | **[]string** | the fields to include for returned resources of type appCustomProductPageVersions | 
 **fieldsAppStoreVersionExperiments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperiments | 
 **fieldsAppEvents** | **[]string** | the fields to include for returned resources of type appEvents | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**ReviewSubmissionItemsResponse**](ReviewSubmissionItemsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewSubmissionsUpdateInstance

> ReviewSubmissionResponse ReviewSubmissionsUpdateInstance(ctx, id).ReviewSubmissionUpdateRequest(reviewSubmissionUpdateRequest).Execute()



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
	reviewSubmissionUpdateRequest := *openapiclient.NewReviewSubmissionUpdateRequest(*openapiclient.NewReviewSubmissionUpdateRequestData("Type_example", "Id_example")) // ReviewSubmissionUpdateRequest | ReviewSubmission representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewSubmissionsAPI.ReviewSubmissionsUpdateInstance(context.Background(), id).ReviewSubmissionUpdateRequest(reviewSubmissionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewSubmissionsAPI.ReviewSubmissionsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewSubmissionsUpdateInstance`: ReviewSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `ReviewSubmissionsAPI.ReviewSubmissionsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReviewSubmissionsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewSubmissionUpdateRequest** | [**ReviewSubmissionUpdateRequest**](ReviewSubmissionUpdateRequest.md) | ReviewSubmission representation | 

### Return type

[**ReviewSubmissionResponse**](ReviewSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

