# \AppStoreVersionSubmissionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionSubmissionsCreateInstance**](AppStoreVersionSubmissionsAPI.md#AppStoreVersionSubmissionsCreateInstance) | **Post** /v1/appStoreVersionSubmissions | 
[**AppStoreVersionSubmissionsDeleteInstance**](AppStoreVersionSubmissionsAPI.md#AppStoreVersionSubmissionsDeleteInstance) | **Delete** /v1/appStoreVersionSubmissions/{id} | 



## AppStoreVersionSubmissionsCreateInstance

> AppStoreVersionSubmissionResponse AppStoreVersionSubmissionsCreateInstance(ctx).AppStoreVersionSubmissionCreateRequest(appStoreVersionSubmissionCreateRequest).Execute()



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
	appStoreVersionSubmissionCreateRequest := *openapiclient.NewAppStoreVersionSubmissionCreateRequest(*openapiclient.NewAppStoreVersionSubmissionCreateRequestData("Type_example", *openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationships(*openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData("Type_example", "Id_example"))))) // AppStoreVersionSubmissionCreateRequest | AppStoreVersionSubmission representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionSubmissionsAPI.AppStoreVersionSubmissionsCreateInstance(context.Background()).AppStoreVersionSubmissionCreateRequest(appStoreVersionSubmissionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionSubmissionsAPI.AppStoreVersionSubmissionsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionSubmissionsCreateInstance`: AppStoreVersionSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionSubmissionsAPI.AppStoreVersionSubmissionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionSubmissionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreVersionSubmissionCreateRequest** | [**AppStoreVersionSubmissionCreateRequest**](AppStoreVersionSubmissionCreateRequest.md) | AppStoreVersionSubmission representation | 

### Return type

[**AppStoreVersionSubmissionResponse**](AppStoreVersionSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionSubmissionsDeleteInstance

> AppStoreVersionSubmissionsDeleteInstance(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppStoreVersionSubmissionsAPI.AppStoreVersionSubmissionsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionSubmissionsAPI.AppStoreVersionSubmissionsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionSubmissionsDeleteInstanceRequest struct via the builder pattern


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

