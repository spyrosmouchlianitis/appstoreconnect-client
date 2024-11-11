# \SubscriptionAppStoreReviewScreenshotsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionAppStoreReviewScreenshotsCreateInstance**](SubscriptionAppStoreReviewScreenshotsAPI.md#SubscriptionAppStoreReviewScreenshotsCreateInstance) | **Post** /v1/subscriptionAppStoreReviewScreenshots | 
[**SubscriptionAppStoreReviewScreenshotsDeleteInstance**](SubscriptionAppStoreReviewScreenshotsAPI.md#SubscriptionAppStoreReviewScreenshotsDeleteInstance) | **Delete** /v1/subscriptionAppStoreReviewScreenshots/{id} | 
[**SubscriptionAppStoreReviewScreenshotsGetInstance**](SubscriptionAppStoreReviewScreenshotsAPI.md#SubscriptionAppStoreReviewScreenshotsGetInstance) | **Get** /v1/subscriptionAppStoreReviewScreenshots/{id} | 
[**SubscriptionAppStoreReviewScreenshotsUpdateInstance**](SubscriptionAppStoreReviewScreenshotsAPI.md#SubscriptionAppStoreReviewScreenshotsUpdateInstance) | **Patch** /v1/subscriptionAppStoreReviewScreenshots/{id} | 



## SubscriptionAppStoreReviewScreenshotsCreateInstance

> SubscriptionAppStoreReviewScreenshotResponse SubscriptionAppStoreReviewScreenshotsCreateInstance(ctx).SubscriptionAppStoreReviewScreenshotCreateRequest(subscriptionAppStoreReviewScreenshotCreateRequest).Execute()



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
	subscriptionAppStoreReviewScreenshotCreateRequest := *openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequest(*openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships(*openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(*openapiclient.NewPromotedPurchaseRelationshipsSubscriptionData("Type_example", "Id_example"))))) // SubscriptionAppStoreReviewScreenshotCreateRequest | SubscriptionAppStoreReviewScreenshot representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsCreateInstance(context.Background()).SubscriptionAppStoreReviewScreenshotCreateRequest(subscriptionAppStoreReviewScreenshotCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAppStoreReviewScreenshotsCreateInstance`: SubscriptionAppStoreReviewScreenshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAppStoreReviewScreenshotsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionAppStoreReviewScreenshotCreateRequest** | [**SubscriptionAppStoreReviewScreenshotCreateRequest**](SubscriptionAppStoreReviewScreenshotCreateRequest.md) | SubscriptionAppStoreReviewScreenshot representation | 

### Return type

[**SubscriptionAppStoreReviewScreenshotResponse**](SubscriptionAppStoreReviewScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionAppStoreReviewScreenshotsDeleteInstance

> SubscriptionAppStoreReviewScreenshotsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionAppStoreReviewScreenshotsDeleteInstanceRequest struct via the builder pattern


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


## SubscriptionAppStoreReviewScreenshotsGetInstance

> SubscriptionAppStoreReviewScreenshotResponse SubscriptionAppStoreReviewScreenshotsGetInstance(ctx, id).FieldsSubscriptionAppStoreReviewScreenshots(fieldsSubscriptionAppStoreReviewScreenshots).Include(include).Execute()



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
	fieldsSubscriptionAppStoreReviewScreenshots := []string{"FieldsSubscriptionAppStoreReviewScreenshots_example"} // []string | the fields to include for returned resources of type subscriptionAppStoreReviewScreenshots (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsGetInstance(context.Background(), id).FieldsSubscriptionAppStoreReviewScreenshots(fieldsSubscriptionAppStoreReviewScreenshots).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAppStoreReviewScreenshotsGetInstance`: SubscriptionAppStoreReviewScreenshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAppStoreReviewScreenshotsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionAppStoreReviewScreenshots** | **[]string** | the fields to include for returned resources of type subscriptionAppStoreReviewScreenshots | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionAppStoreReviewScreenshotResponse**](SubscriptionAppStoreReviewScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionAppStoreReviewScreenshotsUpdateInstance

> SubscriptionAppStoreReviewScreenshotResponse SubscriptionAppStoreReviewScreenshotsUpdateInstance(ctx, id).SubscriptionAppStoreReviewScreenshotUpdateRequest(subscriptionAppStoreReviewScreenshotUpdateRequest).Execute()



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
	subscriptionAppStoreReviewScreenshotUpdateRequest := *openapiclient.NewSubscriptionAppStoreReviewScreenshotUpdateRequest(*openapiclient.NewSubscriptionAppStoreReviewScreenshotUpdateRequestData("Type_example", "Id_example")) // SubscriptionAppStoreReviewScreenshotUpdateRequest | SubscriptionAppStoreReviewScreenshot representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsUpdateInstance(context.Background(), id).SubscriptionAppStoreReviewScreenshotUpdateRequest(subscriptionAppStoreReviewScreenshotUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAppStoreReviewScreenshotsUpdateInstance`: SubscriptionAppStoreReviewScreenshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAppStoreReviewScreenshotsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionAppStoreReviewScreenshotUpdateRequest** | [**SubscriptionAppStoreReviewScreenshotUpdateRequest**](SubscriptionAppStoreReviewScreenshotUpdateRequest.md) | SubscriptionAppStoreReviewScreenshot representation | 

### Return type

[**SubscriptionAppStoreReviewScreenshotResponse**](SubscriptionAppStoreReviewScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

