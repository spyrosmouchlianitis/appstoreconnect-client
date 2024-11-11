# \AppCustomProductPageVersionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated**](AppCustomProductPageVersionsAPI.md#AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated) | **Get** /v1/appCustomProductPageVersions/{id}/appCustomProductPageLocalizations | 
[**AppCustomProductPageVersionsCreateInstance**](AppCustomProductPageVersionsAPI.md#AppCustomProductPageVersionsCreateInstance) | **Post** /v1/appCustomProductPageVersions | 
[**AppCustomProductPageVersionsGetInstance**](AppCustomProductPageVersionsAPI.md#AppCustomProductPageVersionsGetInstance) | **Get** /v1/appCustomProductPageVersions/{id} | 
[**AppCustomProductPageVersionsUpdateInstance**](AppCustomProductPageVersionsAPI.md#AppCustomProductPageVersionsUpdateInstance) | **Patch** /v1/appCustomProductPageVersions/{id} | 



## AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated

> AppCustomProductPageLocalizationsResponse AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated(ctx, id).FilterLocale(filterLocale).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).LimitAppScreenshotSets(limitAppScreenshotSets).LimitAppPreviewSets(limitAppPreviewSets).Execute()



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
	fieldsAppCustomProductPageLocalizations := []string{"FieldsAppCustomProductPageLocalizations_example"} // []string | the fields to include for returned resources of type appCustomProductPageLocalizations (optional)
	fieldsAppCustomProductPageVersions := []string{"FieldsAppCustomProductPageVersions_example"} // []string | the fields to include for returned resources of type appCustomProductPageVersions (optional)
	fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
	fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppScreenshotSets := int32(56) // int32 | maximum number of related appScreenshotSets returned (when they are included) (optional)
	limitAppPreviewSets := int32(56) // int32 | maximum number of related appPreviewSets returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated(context.Background(), id).FilterLocale(filterLocale).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).LimitAppScreenshotSets(limitAppScreenshotSets).LimitAppPreviewSets(limitAppPreviewSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated`: AppCustomProductPageLocalizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageVersionsAppCustomProductPageLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLocale** | **[]string** | filter by attribute &#39;locale&#39; | 
 **fieldsAppCustomProductPageLocalizations** | **[]string** | the fields to include for returned resources of type appCustomProductPageLocalizations | 
 **fieldsAppCustomProductPageVersions** | **[]string** | the fields to include for returned resources of type appCustomProductPageVersions | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppScreenshotSets** | **int32** | maximum number of related appScreenshotSets returned (when they are included) | 
 **limitAppPreviewSets** | **int32** | maximum number of related appPreviewSets returned (when they are included) | 

### Return type

[**AppCustomProductPageLocalizationsResponse**](AppCustomProductPageLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomProductPageVersionsCreateInstance

> AppCustomProductPageVersionResponse AppCustomProductPageVersionsCreateInstance(ctx).AppCustomProductPageVersionCreateRequest(appCustomProductPageVersionCreateRequest).Execute()



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
	appCustomProductPageVersionCreateRequest := *openapiclient.NewAppCustomProductPageVersionCreateRequest(*openapiclient.NewAppCustomProductPageVersionCreateRequestData("Type_example", *openapiclient.NewAppCustomProductPageVersionCreateRequestDataRelationships(*openapiclient.NewAppCustomProductPageVersionCreateRequestDataRelationshipsAppCustomProductPage(*openapiclient.NewAppCustomProductPageVersionRelationshipsAppCustomProductPageData("Type_example", "Id_example"))))) // AppCustomProductPageVersionCreateRequest | AppCustomProductPageVersion representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsCreateInstance(context.Background()).AppCustomProductPageVersionCreateRequest(appCustomProductPageVersionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageVersionsCreateInstance`: AppCustomProductPageVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageVersionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appCustomProductPageVersionCreateRequest** | [**AppCustomProductPageVersionCreateRequest**](AppCustomProductPageVersionCreateRequest.md) | AppCustomProductPageVersion representation | 

### Return type

[**AppCustomProductPageVersionResponse**](AppCustomProductPageVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomProductPageVersionsGetInstance

> AppCustomProductPageVersionResponse AppCustomProductPageVersionsGetInstance(ctx, id).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).Include(include).LimitAppCustomProductPageLocalizations(limitAppCustomProductPageLocalizations).Execute()



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
	fieldsAppCustomProductPageVersions := []string{"FieldsAppCustomProductPageVersions_example"} // []string | the fields to include for returned resources of type appCustomProductPageVersions (optional)
	fieldsAppCustomProductPageLocalizations := []string{"FieldsAppCustomProductPageLocalizations_example"} // []string | the fields to include for returned resources of type appCustomProductPageLocalizations (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppCustomProductPageLocalizations := int32(56) // int32 | maximum number of related appCustomProductPageLocalizations returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsGetInstance(context.Background(), id).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).Include(include).LimitAppCustomProductPageLocalizations(limitAppCustomProductPageLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageVersionsGetInstance`: AppCustomProductPageVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageVersionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCustomProductPageVersions** | **[]string** | the fields to include for returned resources of type appCustomProductPageVersions | 
 **fieldsAppCustomProductPageLocalizations** | **[]string** | the fields to include for returned resources of type appCustomProductPageLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppCustomProductPageLocalizations** | **int32** | maximum number of related appCustomProductPageLocalizations returned (when they are included) | 

### Return type

[**AppCustomProductPageVersionResponse**](AppCustomProductPageVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomProductPageVersionsUpdateInstance

> AppCustomProductPageVersionResponse AppCustomProductPageVersionsUpdateInstance(ctx, id).AppCustomProductPageVersionUpdateRequest(appCustomProductPageVersionUpdateRequest).Execute()



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
	appCustomProductPageVersionUpdateRequest := *openapiclient.NewAppCustomProductPageVersionUpdateRequest(*openapiclient.NewAppCustomProductPageVersionUpdateRequestData("Type_example", "Id_example")) // AppCustomProductPageVersionUpdateRequest | AppCustomProductPageVersion representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsUpdateInstance(context.Background(), id).AppCustomProductPageVersionUpdateRequest(appCustomProductPageVersionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPageVersionsUpdateInstance`: AppCustomProductPageVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPageVersionsAPI.AppCustomProductPageVersionsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPageVersionsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appCustomProductPageVersionUpdateRequest** | [**AppCustomProductPageVersionUpdateRequest**](AppCustomProductPageVersionUpdateRequest.md) | AppCustomProductPageVersion representation | 

### Return type

[**AppCustomProductPageVersionResponse**](AppCustomProductPageVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

