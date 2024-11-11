# \SubscriptionImagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionImagesCreateInstance**](SubscriptionImagesAPI.md#SubscriptionImagesCreateInstance) | **Post** /v1/subscriptionImages | 
[**SubscriptionImagesDeleteInstance**](SubscriptionImagesAPI.md#SubscriptionImagesDeleteInstance) | **Delete** /v1/subscriptionImages/{id} | 
[**SubscriptionImagesGetInstance**](SubscriptionImagesAPI.md#SubscriptionImagesGetInstance) | **Get** /v1/subscriptionImages/{id} | 
[**SubscriptionImagesUpdateInstance**](SubscriptionImagesAPI.md#SubscriptionImagesUpdateInstance) | **Patch** /v1/subscriptionImages/{id} | 



## SubscriptionImagesCreateInstance

> SubscriptionImageResponse SubscriptionImagesCreateInstance(ctx).SubscriptionImageCreateRequest(subscriptionImageCreateRequest).Execute()



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
	subscriptionImageCreateRequest := *openapiclient.NewSubscriptionImageCreateRequest(*openapiclient.NewSubscriptionImageCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships(*openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(*openapiclient.NewPromotedPurchaseRelationshipsSubscriptionData("Type_example", "Id_example"))))) // SubscriptionImageCreateRequest | SubscriptionImage representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionImagesAPI.SubscriptionImagesCreateInstance(context.Background()).SubscriptionImageCreateRequest(subscriptionImageCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImagesAPI.SubscriptionImagesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionImagesCreateInstance`: SubscriptionImageResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionImagesAPI.SubscriptionImagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionImagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionImageCreateRequest** | [**SubscriptionImageCreateRequest**](SubscriptionImageCreateRequest.md) | SubscriptionImage representation | 

### Return type

[**SubscriptionImageResponse**](SubscriptionImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionImagesDeleteInstance

> SubscriptionImagesDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.SubscriptionImagesAPI.SubscriptionImagesDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImagesAPI.SubscriptionImagesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionImagesDeleteInstanceRequest struct via the builder pattern


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


## SubscriptionImagesGetInstance

> SubscriptionImageResponse SubscriptionImagesGetInstance(ctx, id).FieldsSubscriptionImages(fieldsSubscriptionImages).Include(include).Execute()



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
	fieldsSubscriptionImages := []string{"FieldsSubscriptionImages_example"} // []string | the fields to include for returned resources of type subscriptionImages (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionImagesAPI.SubscriptionImagesGetInstance(context.Background(), id).FieldsSubscriptionImages(fieldsSubscriptionImages).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImagesAPI.SubscriptionImagesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionImagesGetInstance`: SubscriptionImageResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionImagesAPI.SubscriptionImagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionImagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionImages** | **[]string** | the fields to include for returned resources of type subscriptionImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionImageResponse**](SubscriptionImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionImagesUpdateInstance

> SubscriptionImageResponse SubscriptionImagesUpdateInstance(ctx, id).SubscriptionImageUpdateRequest(subscriptionImageUpdateRequest).Execute()



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
	subscriptionImageUpdateRequest := *openapiclient.NewSubscriptionImageUpdateRequest(*openapiclient.NewSubscriptionImageUpdateRequestData("Type_example", "Id_example")) // SubscriptionImageUpdateRequest | SubscriptionImage representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionImagesAPI.SubscriptionImagesUpdateInstance(context.Background(), id).SubscriptionImageUpdateRequest(subscriptionImageUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionImagesAPI.SubscriptionImagesUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionImagesUpdateInstance`: SubscriptionImageResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionImagesAPI.SubscriptionImagesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionImagesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionImageUpdateRequest** | [**SubscriptionImageUpdateRequest**](SubscriptionImageUpdateRequest.md) | SubscriptionImage representation | 

### Return type

[**SubscriptionImageResponse**](SubscriptionImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

