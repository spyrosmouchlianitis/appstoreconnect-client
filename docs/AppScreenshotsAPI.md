# \AppScreenshotsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppScreenshotsCreateInstance**](AppScreenshotsAPI.md#AppScreenshotsCreateInstance) | **Post** /v1/appScreenshots | 
[**AppScreenshotsDeleteInstance**](AppScreenshotsAPI.md#AppScreenshotsDeleteInstance) | **Delete** /v1/appScreenshots/{id} | 
[**AppScreenshotsGetInstance**](AppScreenshotsAPI.md#AppScreenshotsGetInstance) | **Get** /v1/appScreenshots/{id} | 
[**AppScreenshotsUpdateInstance**](AppScreenshotsAPI.md#AppScreenshotsUpdateInstance) | **Patch** /v1/appScreenshots/{id} | 



## AppScreenshotsCreateInstance

> AppScreenshotResponse AppScreenshotsCreateInstance(ctx).AppScreenshotCreateRequest(appScreenshotCreateRequest).Execute()



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
	appScreenshotCreateRequest := *openapiclient.NewAppScreenshotCreateRequest(*openapiclient.NewAppScreenshotCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewAppScreenshotCreateRequestDataRelationships(*openapiclient.NewAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet(*openapiclient.NewAppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner("Type_example", "Id_example"))))) // AppScreenshotCreateRequest | AppScreenshot representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppScreenshotsAPI.AppScreenshotsCreateInstance(context.Background()).AppScreenshotCreateRequest(appScreenshotCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotsAPI.AppScreenshotsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppScreenshotsCreateInstance`: AppScreenshotResponse
	fmt.Fprintf(os.Stdout, "Response from `AppScreenshotsAPI.AppScreenshotsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appScreenshotCreateRequest** | [**AppScreenshotCreateRequest**](AppScreenshotCreateRequest.md) | AppScreenshot representation | 

### Return type

[**AppScreenshotResponse**](AppScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotsDeleteInstance

> AppScreenshotsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppScreenshotsAPI.AppScreenshotsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotsAPI.AppScreenshotsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppScreenshotsDeleteInstanceRequest struct via the builder pattern


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


## AppScreenshotsGetInstance

> AppScreenshotResponse AppScreenshotsGetInstance(ctx, id).FieldsAppScreenshots(fieldsAppScreenshots).Include(include).Execute()



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
	fieldsAppScreenshots := []string{"FieldsAppScreenshots_example"} // []string | the fields to include for returned resources of type appScreenshots (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppScreenshotsAPI.AppScreenshotsGetInstance(context.Background(), id).FieldsAppScreenshots(fieldsAppScreenshots).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotsAPI.AppScreenshotsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppScreenshotsGetInstance`: AppScreenshotResponse
	fmt.Fprintf(os.Stdout, "Response from `AppScreenshotsAPI.AppScreenshotsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppScreenshots** | **[]string** | the fields to include for returned resources of type appScreenshots | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppScreenshotResponse**](AppScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotsUpdateInstance

> AppScreenshotResponse AppScreenshotsUpdateInstance(ctx, id).AppScreenshotUpdateRequest(appScreenshotUpdateRequest).Execute()



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
	appScreenshotUpdateRequest := *openapiclient.NewAppScreenshotUpdateRequest(*openapiclient.NewAppScreenshotUpdateRequestData("Type_example", "Id_example")) // AppScreenshotUpdateRequest | AppScreenshot representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppScreenshotsAPI.AppScreenshotsUpdateInstance(context.Background(), id).AppScreenshotUpdateRequest(appScreenshotUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotsAPI.AppScreenshotsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppScreenshotsUpdateInstance`: AppScreenshotResponse
	fmt.Fprintf(os.Stdout, "Response from `AppScreenshotsAPI.AppScreenshotsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appScreenshotUpdateRequest** | [**AppScreenshotUpdateRequest**](AppScreenshotUpdateRequest.md) | AppScreenshot representation | 

### Return type

[**AppScreenshotResponse**](AppScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

