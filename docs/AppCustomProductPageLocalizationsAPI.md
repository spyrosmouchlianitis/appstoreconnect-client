# \AppCustomProductPageLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelated**](AppCustomProductPageLocalizationsAPI.md#AppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelated) | **Get** /v1/appCustomProductPageLocalizations/{id}/appPreviewSets | 
[**AppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelated**](AppCustomProductPageLocalizationsAPI.md#AppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelated) | **Get** /v1/appCustomProductPageLocalizations/{id}/appScreenshotSets | 
[**AppCustomProductPageLocalizationsCreateInstance**](AppCustomProductPageLocalizationsAPI.md#AppCustomProductPageLocalizationsCreateInstance) | **Post** /v1/appCustomProductPageLocalizations | 
[**AppCustomProductPageLocalizationsDeleteInstance**](AppCustomProductPageLocalizationsAPI.md#AppCustomProductPageLocalizationsDeleteInstance) | **Delete** /v1/appCustomProductPageLocalizations/{id} | 
[**AppCustomProductPageLocalizationsGetInstance**](AppCustomProductPageLocalizationsAPI.md#AppCustomProductPageLocalizationsGetInstance) | **Get** /v1/appCustomProductPageLocalizations/{id} | 
[**AppCustomProductPageLocalizationsUpdateInstance**](AppCustomProductPageLocalizationsAPI.md#AppCustomProductPageLocalizationsUpdateInstance) | **Patch** /v1/appCustomProductPageLocalizations/{id} | 



## AppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelated

> AppPreviewSetsResponse AppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelated(ctx, id).FilterPreviewType(filterPreviewType).FilterAppStoreVersionLocalization(filterAppStoreVersionLocalization).FilterAppStoreVersionExperimentTreatmentLocalization(filterAppStoreVersionExperimentTreatmentLocalization).FieldsAppPreviewSets(fieldsAppPreviewSets).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppPreviews(fieldsAppPreviews).Limit(limit).Include(include).LimitAppPreviews(limitAppPreviews).Execute()



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
	filterPreviewType := []string{"FilterPreviewType_example"} // []string | filter by attribute 'previewType' (optional)
	filterAppStoreVersionLocalization := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersionLocalization' (optional)
	filterAppStoreVersionExperimentTreatmentLocalization := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersionExperimentTreatmentLocalization' (optional)
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
	resp, r, err := apiClient.AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelated(context.Background(), id).FilterPreviewType(filterPreviewType).FilterAppStoreVersionLocalization(filterAppStoreVersionLocalization).FilterAppStoreVersionExperimentTreatmentLocalization(filterAppStoreVersionExperimentTreatmentLocalization).FieldsAppPreviewSets(fieldsAppPreviewSets).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppPreviews(fieldsAppPreviews).Limit(limit).Include(include).LimitAppPreviews(limitAppPreviews).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelated`: AppPreviewSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageLocalizationsAppPreviewSetsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPreviewType** | **[]string** | filter by attribute &#39;previewType&#39; | 
 **filterAppStoreVersionLocalization** | **[]string** | filter by id(s) of related &#39;appStoreVersionLocalization&#39; | 
 **filterAppStoreVersionExperimentTreatmentLocalization** | **[]string** | filter by id(s) of related &#39;appStoreVersionExperimentTreatmentLocalization&#39; | 
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


## AppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelated

> AppScreenshotSetsResponse AppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelated(ctx, id).FilterScreenshotDisplayType(filterScreenshotDisplayType).FilterAppStoreVersionLocalization(filterAppStoreVersionLocalization).FilterAppStoreVersionExperimentTreatmentLocalization(filterAppStoreVersionExperimentTreatmentLocalization).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppScreenshots(fieldsAppScreenshots).Limit(limit).Include(include).LimitAppScreenshots(limitAppScreenshots).Execute()



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
	filterScreenshotDisplayType := []string{"FilterScreenshotDisplayType_example"} // []string | filter by attribute 'screenshotDisplayType' (optional)
	filterAppStoreVersionLocalization := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersionLocalization' (optional)
	filterAppStoreVersionExperimentTreatmentLocalization := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersionExperimentTreatmentLocalization' (optional)
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
	resp, r, err := apiClient.AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelated(context.Background(), id).FilterScreenshotDisplayType(filterScreenshotDisplayType).FilterAppStoreVersionLocalization(filterAppStoreVersionLocalization).FilterAppStoreVersionExperimentTreatmentLocalization(filterAppStoreVersionExperimentTreatmentLocalization).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppStoreVersionExperimentTreatmentLocalizations(fieldsAppStoreVersionExperimentTreatmentLocalizations).FieldsAppScreenshots(fieldsAppScreenshots).Limit(limit).Include(include).LimitAppScreenshots(limitAppScreenshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelated`: AppScreenshotSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageLocalizationsAppScreenshotSetsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterScreenshotDisplayType** | **[]string** | filter by attribute &#39;screenshotDisplayType&#39; | 
 **filterAppStoreVersionLocalization** | **[]string** | filter by id(s) of related &#39;appStoreVersionLocalization&#39; | 
 **filterAppStoreVersionExperimentTreatmentLocalization** | **[]string** | filter by id(s) of related &#39;appStoreVersionExperimentTreatmentLocalization&#39; | 
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


## AppCustomProductPageLocalizationsCreateInstance

> AppCustomProductPageLocalizationResponse AppCustomProductPageLocalizationsCreateInstance(ctx).AppCustomProductPageLocalizationCreateRequest(appCustomProductPageLocalizationCreateRequest).Execute()



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
	appCustomProductPageLocalizationCreateRequest := *openapiclient.NewAppCustomProductPageLocalizationCreateRequest(*openapiclient.NewAppCustomProductPageLocalizationCreateRequestData("Type_example", *openapiclient.NewAppCustomProductPageLocalizationInlineCreateAttributes("Locale_example"), *openapiclient.NewAppCustomProductPageLocalizationCreateRequestDataRelationships(*openapiclient.NewAppCustomProductPageLocalizationCreateRequestDataRelationshipsAppCustomProductPageVersion(*openapiclient.NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData("Type_example", "Id_example"))))) // AppCustomProductPageLocalizationCreateRequest | AppCustomProductPageLocalization representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsCreateInstance(context.Background()).AppCustomProductPageLocalizationCreateRequest(appCustomProductPageLocalizationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageLocalizationsCreateInstance`: AppCustomProductPageLocalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appCustomProductPageLocalizationCreateRequest** | [**AppCustomProductPageLocalizationCreateRequest**](AppCustomProductPageLocalizationCreateRequest.md) | AppCustomProductPageLocalization representation | 

### Return type

[**AppCustomProductPageLocalizationResponse**](AppCustomProductPageLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomProductPageLocalizationsDeleteInstance

> AppCustomProductPageLocalizationsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppCustomProductPageLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## AppCustomProductPageLocalizationsGetInstance

> AppCustomProductPageLocalizationResponse AppCustomProductPageLocalizationsGetInstance(ctx, id).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Include(include).LimitAppPreviewSets(limitAppPreviewSets).LimitAppScreenshotSets(limitAppScreenshotSets).Execute()



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
	fieldsAppCustomProductPageLocalizations := []string{"FieldsAppCustomProductPageLocalizations_example"} // []string | the fields to include for returned resources of type appCustomProductPageLocalizations (optional)
	fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
	fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppPreviewSets := int32(56) // int32 | maximum number of related appPreviewSets returned (when they are included) (optional)
	limitAppScreenshotSets := int32(56) // int32 | maximum number of related appScreenshotSets returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsGetInstance(context.Background(), id).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Include(include).LimitAppPreviewSets(limitAppPreviewSets).LimitAppScreenshotSets(limitAppScreenshotSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageLocalizationsGetInstance`: AppCustomProductPageLocalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCustomProductPageLocalizations** | **[]string** | the fields to include for returned resources of type appCustomProductPageLocalizations | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppPreviewSets** | **int32** | maximum number of related appPreviewSets returned (when they are included) | 
 **limitAppScreenshotSets** | **int32** | maximum number of related appScreenshotSets returned (when they are included) | 

### Return type

[**AppCustomProductPageLocalizationResponse**](AppCustomProductPageLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomProductPageLocalizationsUpdateInstance

> AppCustomProductPageLocalizationResponse AppCustomProductPageLocalizationsUpdateInstance(ctx, id).AppCustomProductPageLocalizationUpdateRequest(appCustomProductPageLocalizationUpdateRequest).Execute()



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
	appCustomProductPageLocalizationUpdateRequest := *openapiclient.NewAppCustomProductPageLocalizationUpdateRequest(*openapiclient.NewAppCustomProductPageLocalizationUpdateRequestData("Type_example", "Id_example")) // AppCustomProductPageLocalizationUpdateRequest | AppCustomProductPageLocalization representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsUpdateInstance(context.Background(), id).AppCustomProductPageLocalizationUpdateRequest(appCustomProductPageLocalizationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageLocalizationsUpdateInstance`: AppCustomProductPageLocalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageLocalizationsAPI.AppCustomProductPageLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appCustomProductPageLocalizationUpdateRequest** | [**AppCustomProductPageLocalizationUpdateRequest**](AppCustomProductPageLocalizationUpdateRequest.md) | AppCustomProductPageLocalization representation | 

### Return type

[**AppCustomProductPageLocalizationResponse**](AppCustomProductPageLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

