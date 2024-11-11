# \AppStoreVersionExperimentTreatmentLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated**](AppStoreVersionExperimentTreatmentLocalizationsAPI.md#AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated) | **Get** /v1/appStoreVersionExperimentTreatmentLocalizations/{id}/appPreviewSets | 
[**AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated**](AppStoreVersionExperimentTreatmentLocalizationsAPI.md#AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated) | **Get** /v1/appStoreVersionExperimentTreatmentLocalizations/{id}/appScreenshotSets | 
[**AppStoreVersionExperimentTreatmentLocalizationsCreateInstance**](AppStoreVersionExperimentTreatmentLocalizationsAPI.md#AppStoreVersionExperimentTreatmentLocalizationsCreateInstance) | **Post** /v1/appStoreVersionExperimentTreatmentLocalizations | 
[**AppStoreVersionExperimentTreatmentLocalizationsDeleteInstance**](AppStoreVersionExperimentTreatmentLocalizationsAPI.md#AppStoreVersionExperimentTreatmentLocalizationsDeleteInstance) | **Delete** /v1/appStoreVersionExperimentTreatmentLocalizations/{id} | 
[**AppStoreVersionExperimentTreatmentLocalizationsGetInstance**](AppStoreVersionExperimentTreatmentLocalizationsAPI.md#AppStoreVersionExperimentTreatmentLocalizationsGetInstance) | **Get** /v1/appStoreVersionExperimentTreatmentLocalizations/{id} | 



## AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated

> AppPreviewSetsResponse AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated(ctx, id).FilterPreviewType(filterPreviewType).FilterAppStoreVersionLocalization(filterAppStoreVersionLocalization).FilterAppCustomProductPageLocalization(filterAppCustomProductPageLocalization).FieldsAppPreviewSets(fieldsAppPreviewSets).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppPreviews(fieldsAppPreviews).Limit(limit).Include(include).LimitAppPreviews(limitAppPreviews).Execute()



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
	filterPreviewType := []string{"FilterPreviewType_example"} // []string | filter by attribute 'previewType' (optional)
	filterAppStoreVersionLocalization := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersionLocalization' (optional)
	filterAppCustomProductPageLocalization := []string{"Inner_example"} // []string | filter by id(s) of related 'appCustomProductPageLocalization' (optional)
	fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
	fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
	fieldsAppCustomProductPageLocalizations := []string{"FieldsAppCustomProductPageLocalizations_example"} // []string | the fields to include for returned resources of type appCustomProductPageLocalizations (optional)
	fieldsAppStoreVersionExperimentTreatmentLocalizations := []string{"FieldsAppStoreVersionExperimentTreatmentLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations (optional)
	fieldsAppPreviews := []string{"FieldsAppPreviews_example"} // []string | the fields to include for returned resources of type appPreviews (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppPreviews := int32(56) // int32 | maximum number of related appPreviews returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated(context.Background(), id).FilterPreviewType(filterPreviewType).FilterAppStoreVersionLocalization(filterAppStoreVersionLocalization).FilterAppCustomProductPageLocalization(filterAppCustomProductPageLocalization).FieldsAppPreviewSets(fieldsAppPreviewSets).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppPreviews(fieldsAppPreviews).Limit(limit).Include(include).LimitAppPreviews(limitAppPreviews).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated`: AppPreviewSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentLocalizationsAppPreviewSetsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPreviewType** | **[]string** | filter by attribute &#39;previewType&#39; | 
 **filterAppStoreVersionLocalization** | **[]string** | filter by id(s) of related &#39;appStoreVersionLocalization&#39; | 
 **filterAppCustomProductPageLocalization** | **[]string** | filter by id(s) of related &#39;appCustomProductPageLocalization&#39; | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsAppCustomProductPageLocalizations** | **[]string** | the fields to include for returned resources of type appCustomProductPageLocalizations | 
 **fieldsAppStoreVersionExperimentTreatmentLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations | 
 **fieldsAppPreviews** | **[]string** | the fields to include for returned resources of type appPreviews | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppPreviews** | **int32** | maximum number of related appPreviews returned (when they are included) | 

### Return type

[**AppPreviewSetsResponse**](AppPreviewSetsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated

> AppScreenshotSetsResponse AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated(ctx, id).FilterScreenshotDisplayType(filterScreenshotDisplayType).FilterAppStoreVersionLocalization(filterAppStoreVersionLocalization).FilterAppCustomProductPageLocalization(filterAppCustomProductPageLocalization).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppScreenshots(fieldsAppScreenshots).Limit(limit).Include(include).LimitAppScreenshots(limitAppScreenshots).Execute()



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
	filterScreenshotDisplayType := []string{"FilterScreenshotDisplayType_example"} // []string | filter by attribute 'screenshotDisplayType' (optional)
	filterAppStoreVersionLocalization := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersionLocalization' (optional)
	filterAppCustomProductPageLocalization := []string{"Inner_example"} // []string | filter by id(s) of related 'appCustomProductPageLocalization' (optional)
	fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
	fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
	fieldsAppCustomProductPageLocalizations := []string{"FieldsAppCustomProductPageLocalizations_example"} // []string | the fields to include for returned resources of type appCustomProductPageLocalizations (optional)
	fieldsAppStoreVersionExperimentTreatmentLocalizations := []string{"FieldsAppStoreVersionExperimentTreatmentLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations (optional)
	fieldsAppScreenshots := []string{"FieldsAppScreenshots_example"} // []string | the fields to include for returned resources of type appScreenshots (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppScreenshots := int32(56) // int32 | maximum number of related appScreenshots returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated(context.Background(), id).FilterScreenshotDisplayType(filterScreenshotDisplayType).FilterAppStoreVersionLocalization(filterAppStoreVersionLocalization).FilterAppCustomProductPageLocalization(filterAppCustomProductPageLocalization).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppScreenshots(fieldsAppScreenshots).Limit(limit).Include(include).LimitAppScreenshots(limitAppScreenshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated`: AppScreenshotSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentLocalizationsAppScreenshotSetsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterScreenshotDisplayType** | **[]string** | filter by attribute &#39;screenshotDisplayType&#39; | 
 **filterAppStoreVersionLocalization** | **[]string** | filter by id(s) of related &#39;appStoreVersionLocalization&#39; | 
 **filterAppCustomProductPageLocalization** | **[]string** | filter by id(s) of related &#39;appCustomProductPageLocalization&#39; | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsAppCustomProductPageLocalizations** | **[]string** | the fields to include for returned resources of type appCustomProductPageLocalizations | 
 **fieldsAppStoreVersionExperimentTreatmentLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations | 
 **fieldsAppScreenshots** | **[]string** | the fields to include for returned resources of type appScreenshots | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppScreenshots** | **int32** | maximum number of related appScreenshots returned (when they are included) | 

### Return type

[**AppScreenshotSetsResponse**](AppScreenshotSetsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionExperimentTreatmentLocalizationsCreateInstance

> AppStoreVersionExperimentTreatmentLocalizationResponse AppStoreVersionExperimentTreatmentLocalizationsCreateInstance(ctx).AppStoreVersionExperimentTreatmentLocalizationCreateRequest(appStoreVersionExperimentTreatmentLocalizationCreateRequest).Execute()



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
	appStoreVersionExperimentTreatmentLocalizationCreateRequest := *openapiclient.NewAppStoreVersionExperimentTreatmentLocalizationCreateRequest(*openapiclient.NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestData("Type_example", *openapiclient.NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes("Locale_example"), *openapiclient.NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationships(*openapiclient.NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationshipsAppStoreVersionExperimentTreatment(*openapiclient.NewAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData("Type_example", "Id_example"))))) // AppStoreVersionExperimentTreatmentLocalizationCreateRequest | AppStoreVersionExperimentTreatmentLocalization representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsCreateInstance(context.Background()).AppStoreVersionExperimentTreatmentLocalizationCreateRequest(appStoreVersionExperimentTreatmentLocalizationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionExperimentTreatmentLocalizationsCreateInstance`: AppStoreVersionExperimentTreatmentLocalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreVersionExperimentTreatmentLocalizationCreateRequest** | [**AppStoreVersionExperimentTreatmentLocalizationCreateRequest**](AppStoreVersionExperimentTreatmentLocalizationCreateRequest.md) | AppStoreVersionExperimentTreatmentLocalization representation | 

### Return type

[**AppStoreVersionExperimentTreatmentLocalizationResponse**](AppStoreVersionExperimentTreatmentLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionExperimentTreatmentLocalizationsDeleteInstance

> AppStoreVersionExperimentTreatmentLocalizationsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## AppStoreVersionExperimentTreatmentLocalizationsGetInstance

> AppStoreVersionExperimentTreatmentLocalizationResponse AppStoreVersionExperimentTreatmentLocalizationsGetInstance(ctx, id).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Include(include).LimitAppPreviewSets(limitAppPreviewSets).LimitAppScreenshotSets(limitAppScreenshotSets).Execute()



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
	fieldsAppStoreVersionExperimentTreatmentLocalizations := []string{"FieldsAppStoreVersionExperimentTreatmentLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations (optional)
	fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
	fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppPreviewSets := int32(56) // int32 | maximum number of related appPreviewSets returned (when they are included) (optional)
	limitAppScreenshotSets := int32(56) // int32 | maximum number of related appScreenshotSets returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsGetInstance(context.Background(), id).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Include(include).LimitAppPreviewSets(limitAppPreviewSets).LimitAppScreenshotSets(limitAppScreenshotSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionExperimentTreatmentLocalizationsGetInstance`: AppStoreVersionExperimentTreatmentLocalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionExperimentTreatmentLocalizationsAPI.AppStoreVersionExperimentTreatmentLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionExperimentTreatmentLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionExperimentTreatmentLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatmentLocalizations | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppPreviewSets** | **int32** | maximum number of related appPreviewSets returned (when they are included) | 
 **limitAppScreenshotSets** | **int32** | maximum number of related appScreenshotSets returned (when they are included) | 

### Return type

[**AppStoreVersionExperimentTreatmentLocalizationResponse**](AppStoreVersionExperimentTreatmentLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

