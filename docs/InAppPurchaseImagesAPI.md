# \InAppPurchaseImagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InAppPurchaseImagesCreateInstance**](InAppPurchaseImagesAPI.md#InAppPurchaseImagesCreateInstance) | **Post** /v1/inAppPurchaseImages | 
[**InAppPurchaseImagesDeleteInstance**](InAppPurchaseImagesAPI.md#InAppPurchaseImagesDeleteInstance) | **Delete** /v1/inAppPurchaseImages/{id} | 
[**InAppPurchaseImagesGetInstance**](InAppPurchaseImagesAPI.md#InAppPurchaseImagesGetInstance) | **Get** /v1/inAppPurchaseImages/{id} | 
[**InAppPurchaseImagesUpdateInstance**](InAppPurchaseImagesAPI.md#InAppPurchaseImagesUpdateInstance) | **Patch** /v1/inAppPurchaseImages/{id} | 



## InAppPurchaseImagesCreateInstance

> InAppPurchaseImageResponse InAppPurchaseImagesCreateInstance(ctx).InAppPurchaseImageCreateRequest(inAppPurchaseImageCreateRequest).Execute()



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
	inAppPurchaseImageCreateRequest := *openapiclient.NewInAppPurchaseImageCreateRequest(*openapiclient.NewInAppPurchaseImageCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewInAppPurchaseImageCreateRequestDataRelationships(*openapiclient.NewInAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2(*openapiclient.NewAppRelationshipsInAppPurchasesDataInner("Type_example", "Id_example"))))) // InAppPurchaseImageCreateRequest | InAppPurchaseImage representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchaseImagesAPI.InAppPurchaseImagesCreateInstance(context.Background()).InAppPurchaseImageCreateRequest(inAppPurchaseImageCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchaseImagesAPI.InAppPurchaseImagesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchaseImagesCreateInstance`: InAppPurchaseImageResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchaseImagesAPI.InAppPurchaseImagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchaseImagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inAppPurchaseImageCreateRequest** | [**InAppPurchaseImageCreateRequest**](InAppPurchaseImageCreateRequest.md) | InAppPurchaseImage representation | 

### Return type

[**InAppPurchaseImageResponse**](InAppPurchaseImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchaseImagesDeleteInstance

> InAppPurchaseImagesDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.InAppPurchaseImagesAPI.InAppPurchaseImagesDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchaseImagesAPI.InAppPurchaseImagesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInAppPurchaseImagesDeleteInstanceRequest struct via the builder pattern


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


## InAppPurchaseImagesGetInstance

> InAppPurchaseImageResponse InAppPurchaseImagesGetInstance(ctx, id).FieldsInAppPurchaseImages(fieldsInAppPurchaseImages).Include(include).Execute()



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
	fieldsInAppPurchaseImages := []string{"FieldsInAppPurchaseImages_example"} // []string | the fields to include for returned resources of type inAppPurchaseImages (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchaseImagesAPI.InAppPurchaseImagesGetInstance(context.Background(), id).FieldsInAppPurchaseImages(fieldsInAppPurchaseImages).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchaseImagesAPI.InAppPurchaseImagesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchaseImagesGetInstance`: InAppPurchaseImageResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchaseImagesAPI.InAppPurchaseImagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchaseImagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchaseImages** | **[]string** | the fields to include for returned resources of type inAppPurchaseImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchaseImageResponse**](InAppPurchaseImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchaseImagesUpdateInstance

> InAppPurchaseImageResponse InAppPurchaseImagesUpdateInstance(ctx, id).InAppPurchaseImageUpdateRequest(inAppPurchaseImageUpdateRequest).Execute()



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
	inAppPurchaseImageUpdateRequest := *openapiclient.NewInAppPurchaseImageUpdateRequest(*openapiclient.NewInAppPurchaseImageUpdateRequestData("Type_example", "Id_example")) // InAppPurchaseImageUpdateRequest | InAppPurchaseImage representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchaseImagesAPI.InAppPurchaseImagesUpdateInstance(context.Background(), id).InAppPurchaseImageUpdateRequest(inAppPurchaseImageUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchaseImagesAPI.InAppPurchaseImagesUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchaseImagesUpdateInstance`: InAppPurchaseImageResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchaseImagesAPI.InAppPurchaseImagesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchaseImagesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inAppPurchaseImageUpdateRequest** | [**InAppPurchaseImageUpdateRequest**](InAppPurchaseImageUpdateRequest.md) | InAppPurchaseImage representation | 

### Return type

[**InAppPurchaseImageResponse**](InAppPurchaseImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

