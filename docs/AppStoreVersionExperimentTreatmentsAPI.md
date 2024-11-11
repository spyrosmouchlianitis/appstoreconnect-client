# \AppStoreVersionExperimentTreatmentsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated**](AppStoreVersionExperimentTreatmentsAPI.md#AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated) | **Get** /v1/appStoreVersionExperimentTreatments/{id}/appStoreVersionExperimentTreatmentLocalizations | 
[**AppStoreVersionExperimentTreatmentsCreateInstance**](AppStoreVersionExperimentTreatmentsAPI.md#AppStoreVersionExperimentTreatmentsCreateInstance) | **Post** /v1/appStoreVersionExperimentTreatments | 
[**AppStoreVersionExperimentTreatmentsDeleteInstance**](AppStoreVersionExperimentTreatmentsAPI.md#AppStoreVersionExperimentTreatmentsDeleteInstance) | **Delete** /v1/appStoreVersionExperimentTreatments/{id} | 
[**AppStoreVersionExperimentTreatmentsGetInstance**](AppStoreVersionExperimentTreatmentsAPI.md#AppStoreVersionExperimentTreatmentsGetInstance) | **Get** /v1/appStoreVersionExperimentTreatments/{id} | 
[**AppStoreVersionExperimentTreatmentsUpdateInstance**](AppStoreVersionExperimentTreatmentsAPI.md#AppStoreVersionExperimentTreatmentsUpdateInstance) | **Patch** /v1/appStoreVersionExperimentTreatments/{id} | 



## AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated

> AppStoreVersionExperimentTreatmentLocalizationsResponse AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated(ctx, id).FilterLocale(filterLocale).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).LimitAppScreenshotSets(limitAppScreenshotSets).LimitAppPreviewSets(limitAppPreviewSets).Execute()



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
	filterLocale := []string{"Inner_example"} // []string | filter by attribute 'locale' (optional)
	fieldsAppStoreVersionExperimentTreatmentLocalizations := []string{"FieldsAppStoreVersionExperimentTreatmentLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations (optional)
	fieldsAppStoreVersionExperimentTreatments := []string{"FieldsAppStoreVersionExperimentTreatments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatments (optional)
	fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
	fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppScreenshotSets := int32(56) // int32 | maximum number of related appScreenshotSets returned (when they are included) (optional)
	limitAppPreviewSets := int32(56) // int32 | maximum number of related appPreviewSets returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated(context.Background(), id).FilterLocale(filterLocale).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).LimitAppScreenshotSets(limitAppScreenshotSets).LimitAppPreviewSets(limitAppPreviewSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated`: AppStoreVersionExperimentTreatmentLocalizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentsAppStoreVersionExperimentTreatmentLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLocale** | **[]string** | filter by attribute &#39;locale&#39; | 
 **fieldsAppStoreVersionExperimentTreatmentLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations | 
 **fieldsAppStoreVersionExperimentTreatments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatments | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppScreenshotSets** | **int32** | maximum number of related appScreenshotSets returned (when they are included) | 
 **limitAppPreviewSets** | **int32** | maximum number of related appPreviewSets returned (when they are included) | 

### Return type

[**AppStoreVersionExperimentTreatmentLocalizationsResponse**](AppStoreVersionExperimentTreatmentLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionExperimentTreatmentsCreateInstance

> AppStoreVersionExperimentTreatmentResponse AppStoreVersionExperimentTreatmentsCreateInstance(ctx).AppStoreVersionExperimentTreatmentCreateRequest(appStoreVersionExperimentTreatmentCreateRequest).Execute()



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
	appStoreVersionExperimentTreatmentCreateRequest := *openapiclient.NewAppStoreVersionExperimentTreatmentCreateRequest(*openapiclient.NewAppStoreVersionExperimentTreatmentCreateRequestData("Type_example", *openapiclient.NewAppStoreVersionExperimentTreatmentCreateRequestDataAttributes("Name_example"))) // AppStoreVersionExperimentTreatmentCreateRequest | AppStoreVersionExperimentTreatment representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsCreateInstance(context.Background()).AppStoreVersionExperimentTreatmentCreateRequest(appStoreVersionExperimentTreatmentCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionExperimentTreatmentsCreateInstance`: AppStoreVersionExperimentTreatmentResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreVersionExperimentTreatmentCreateRequest** | [**AppStoreVersionExperimentTreatmentCreateRequest**](AppStoreVersionExperimentTreatmentCreateRequest.md) | AppStoreVersionExperimentTreatment representation | 

### Return type

[**AppStoreVersionExperimentTreatmentResponse**](AppStoreVersionExperimentTreatmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionExperimentTreatmentsDeleteInstance

> AppStoreVersionExperimentTreatmentsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentsDeleteInstanceRequest struct via the builder pattern


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


## AppStoreVersionExperimentTreatmentsGetInstance

> AppStoreVersionExperimentTreatmentResponse AppStoreVersionExperimentTreatmentsGetInstance(ctx, id).FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).Include(include).LimitAppStoreVersionExperimentTreatmentLocalizations(limitAppStoreVersionExperimentTreatmentLocalizations).Execute()



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
	fieldsAppStoreVersionExperimentTreatments := []string{"FieldsAppStoreVersionExperimentTreatments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatments (optional)
	fieldsAppStoreVersionExperimentTreatmentLocalizations := []string{"FieldsAppStoreVersionExperimentTreatmentLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppStoreVersionExperimentTreatmentLocalizations := int32(56) // int32 | maximum number of related appStoreVersionExperimentTreatmentLocalizations returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsGetInstance(context.Background(), id).FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).Include(include).LimitAppStoreVersionExperimentTreatmentLocalizations(limitAppStoreVersionExperimentTreatmentLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionExperimentTreatmentsGetInstance`: AppStoreVersionExperimentTreatmentResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionExperimentTreatments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatments | 
 **fieldsAppStoreVersionExperimentTreatmentLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppStoreVersionExperimentTreatmentLocalizations** | **int32** | maximum number of related appStoreVersionExperimentTreatmentLocalizations returned (when they are included) | 

### Return type

[**AppStoreVersionExperimentTreatmentResponse**](AppStoreVersionExperimentTreatmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionExperimentTreatmentsUpdateInstance

> AppStoreVersionExperimentTreatmentResponse AppStoreVersionExperimentTreatmentsUpdateInstance(ctx, id).AppStoreVersionExperimentTreatmentUpdateRequest(appStoreVersionExperimentTreatmentUpdateRequest).Execute()



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
	appStoreVersionExperimentTreatmentUpdateRequest := *openapiclient.NewAppStoreVersionExperimentTreatmentUpdateRequest(*openapiclient.NewAppStoreVersionExperimentTreatmentUpdateRequestData("Type_example", "Id_example")) // AppStoreVersionExperimentTreatmentUpdateRequest | AppStoreVersionExperimentTreatment representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsUpdateInstance(context.Background(), id).AppStoreVersionExperimentTreatmentUpdateRequest(appStoreVersionExperimentTreatmentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionExperimentTreatmentsUpdateInstance`: AppStoreVersionExperimentTreatmentResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionExperimentTreatmentsAPI.AppStoreVersionExperimentTreatmentsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreVersionExperimentTreatmentUpdateRequest** | [**AppStoreVersionExperimentTreatmentUpdateRequest**](AppStoreVersionExperimentTreatmentUpdateRequest.md) | AppStoreVersionExperimentTreatment representation | 

### Return type

[**AppStoreVersionExperimentTreatmentResponse**](AppStoreVersionExperimentTreatmentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

