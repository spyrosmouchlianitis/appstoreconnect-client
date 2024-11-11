# \InAppPurchasesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InAppPurchasesGetInstance**](InAppPurchasesAPI.md#InAppPurchasesGetInstance) | **Get** /v1/inAppPurchases/{id} | 
[**InAppPurchasesV2AppStoreReviewScreenshotGetToOneRelated**](InAppPurchasesAPI.md#InAppPurchasesV2AppStoreReviewScreenshotGetToOneRelated) | **Get** /v2/inAppPurchases/{id}/appStoreReviewScreenshot | 
[**InAppPurchasesV2ContentGetToOneRelated**](InAppPurchasesAPI.md#InAppPurchasesV2ContentGetToOneRelated) | **Get** /v2/inAppPurchases/{id}/content | 
[**InAppPurchasesV2CreateInstance**](InAppPurchasesAPI.md#InAppPurchasesV2CreateInstance) | **Post** /v2/inAppPurchases | 
[**InAppPurchasesV2DeleteInstance**](InAppPurchasesAPI.md#InAppPurchasesV2DeleteInstance) | **Delete** /v2/inAppPurchases/{id} | 
[**InAppPurchasesV2GetInstance**](InAppPurchasesAPI.md#InAppPurchasesV2GetInstance) | **Get** /v2/inAppPurchases/{id} | 
[**InAppPurchasesV2IapPriceScheduleGetToOneRelated**](InAppPurchasesAPI.md#InAppPurchasesV2IapPriceScheduleGetToOneRelated) | **Get** /v2/inAppPurchases/{id}/iapPriceSchedule | 
[**InAppPurchasesV2ImagesGetToManyRelated**](InAppPurchasesAPI.md#InAppPurchasesV2ImagesGetToManyRelated) | **Get** /v2/inAppPurchases/{id}/images | 
[**InAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelated**](InAppPurchasesAPI.md#InAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelated) | **Get** /v2/inAppPurchases/{id}/inAppPurchaseAvailability | 
[**InAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelated**](InAppPurchasesAPI.md#InAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelated) | **Get** /v2/inAppPurchases/{id}/inAppPurchaseLocalizations | 
[**InAppPurchasesV2PricePointsGetToManyRelated**](InAppPurchasesAPI.md#InAppPurchasesV2PricePointsGetToManyRelated) | **Get** /v2/inAppPurchases/{id}/pricePoints | 
[**InAppPurchasesV2PromotedPurchaseGetToOneRelated**](InAppPurchasesAPI.md#InAppPurchasesV2PromotedPurchaseGetToOneRelated) | **Get** /v2/inAppPurchases/{id}/promotedPurchase | 
[**InAppPurchasesV2UpdateInstance**](InAppPurchasesAPI.md#InAppPurchasesV2UpdateInstance) | **Patch** /v2/inAppPurchases/{id} | 



## InAppPurchasesGetInstance

> InAppPurchaseResponse InAppPurchasesGetInstance(ctx, id).FieldsInAppPurchases(fieldsInAppPurchases).Include(include).LimitApps(limitApps).Execute()



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
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitApps := int32(56) // int32 | maximum number of related apps returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesGetInstance(context.Background(), id).FieldsInAppPurchases(fieldsInAppPurchases).Include(include).LimitApps(limitApps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesGetInstance`: InAppPurchaseResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitApps** | **int32** | maximum number of related apps returned (when they are included) | 

### Return type

[**InAppPurchaseResponse**](InAppPurchaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2AppStoreReviewScreenshotGetToOneRelated

> InAppPurchaseAppStoreReviewScreenshotResponse InAppPurchasesV2AppStoreReviewScreenshotGetToOneRelated(ctx, id).FieldsInAppPurchaseAppStoreReviewScreenshots(fieldsInAppPurchaseAppStoreReviewScreenshots).FieldsInAppPurchases(fieldsInAppPurchases).Include(include).Execute()



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
	fieldsInAppPurchaseAppStoreReviewScreenshots := []string{"FieldsInAppPurchaseAppStoreReviewScreenshots_example"} // []string | the fields to include for returned resources of type inAppPurchaseAppStoreReviewScreenshots (optional)
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2AppStoreReviewScreenshotGetToOneRelated(context.Background(), id).FieldsInAppPurchaseAppStoreReviewScreenshots(fieldsInAppPurchaseAppStoreReviewScreenshots).FieldsInAppPurchases(fieldsInAppPurchases).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2AppStoreReviewScreenshotGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2AppStoreReviewScreenshotGetToOneRelated`: InAppPurchaseAppStoreReviewScreenshotResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2AppStoreReviewScreenshotGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2AppStoreReviewScreenshotGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchaseAppStoreReviewScreenshots** | **[]string** | the fields to include for returned resources of type inAppPurchaseAppStoreReviewScreenshots | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchaseAppStoreReviewScreenshotResponse**](InAppPurchaseAppStoreReviewScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2ContentGetToOneRelated

> InAppPurchaseContentResponse InAppPurchasesV2ContentGetToOneRelated(ctx, id).FieldsInAppPurchaseContents(fieldsInAppPurchaseContents).FieldsInAppPurchases(fieldsInAppPurchases).Include(include).Execute()



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
	fieldsInAppPurchaseContents := []string{"FieldsInAppPurchaseContents_example"} // []string | the fields to include for returned resources of type inAppPurchaseContents (optional)
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2ContentGetToOneRelated(context.Background(), id).FieldsInAppPurchaseContents(fieldsInAppPurchaseContents).FieldsInAppPurchases(fieldsInAppPurchases).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2ContentGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2ContentGetToOneRelated`: InAppPurchaseContentResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2ContentGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2ContentGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchaseContents** | **[]string** | the fields to include for returned resources of type inAppPurchaseContents | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchaseContentResponse**](InAppPurchaseContentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2CreateInstance

> InAppPurchaseV2Response InAppPurchasesV2CreateInstance(ctx).InAppPurchaseV2CreateRequest(inAppPurchaseV2CreateRequest).Execute()



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
	inAppPurchaseV2CreateRequest := *openapiclient.NewInAppPurchaseV2CreateRequest(*openapiclient.NewInAppPurchaseV2CreateRequestData("Type_example", *openapiclient.NewInAppPurchaseV2CreateRequestDataAttributes("Name_example", "ProductId_example", openapiclient.InAppPurchaseType("CONSUMABLE")), *openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // InAppPurchaseV2CreateRequest | InAppPurchase representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2CreateInstance(context.Background()).InAppPurchaseV2CreateRequest(inAppPurchaseV2CreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2CreateInstance`: InAppPurchaseV2Response
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2CreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inAppPurchaseV2CreateRequest** | [**InAppPurchaseV2CreateRequest**](InAppPurchaseV2CreateRequest.md) | InAppPurchase representation | 

### Return type

[**InAppPurchaseV2Response**](InAppPurchaseV2Response.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2DeleteInstance

> InAppPurchasesV2DeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2DeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2DeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInAppPurchasesV2DeleteInstanceRequest struct via the builder pattern


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


## InAppPurchasesV2GetInstance

> InAppPurchaseV2Response InAppPurchasesV2GetInstance(ctx, id).FieldsInAppPurchases(fieldsInAppPurchases).FieldsInAppPurchaseLocalizations(fieldsInAppPurchaseLocalizations).FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints).FieldsInAppPurchaseContents(fieldsInAppPurchaseContents).FieldsInAppPurchaseAppStoreReviewScreenshots(fieldsInAppPurchaseAppStoreReviewScreenshots).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsInAppPurchasePriceSchedules(fieldsInAppPurchasePriceSchedules).FieldsInAppPurchaseAvailabilities(fieldsInAppPurchaseAvailabilities).FieldsInAppPurchaseImages(fieldsInAppPurchaseImages).Include(include).LimitImages(limitImages).LimitInAppPurchaseLocalizations(limitInAppPurchaseLocalizations).LimitPricePoints(limitPricePoints).Execute()



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
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	fieldsInAppPurchaseLocalizations := []string{"FieldsInAppPurchaseLocalizations_example"} // []string | the fields to include for returned resources of type inAppPurchaseLocalizations (optional)
	fieldsInAppPurchasePricePoints := []string{"FieldsInAppPurchasePricePoints_example"} // []string | the fields to include for returned resources of type inAppPurchasePricePoints (optional)
	fieldsInAppPurchaseContents := []string{"FieldsInAppPurchaseContents_example"} // []string | the fields to include for returned resources of type inAppPurchaseContents (optional)
	fieldsInAppPurchaseAppStoreReviewScreenshots := []string{"FieldsInAppPurchaseAppStoreReviewScreenshots_example"} // []string | the fields to include for returned resources of type inAppPurchaseAppStoreReviewScreenshots (optional)
	fieldsPromotedPurchases := []string{"FieldsPromotedPurchases_example"} // []string | the fields to include for returned resources of type promotedPurchases (optional)
	fieldsInAppPurchasePriceSchedules := []string{"FieldsInAppPurchasePriceSchedules_example"} // []string | the fields to include for returned resources of type inAppPurchasePriceSchedules (optional)
	fieldsInAppPurchaseAvailabilities := []string{"FieldsInAppPurchaseAvailabilities_example"} // []string | the fields to include for returned resources of type inAppPurchaseAvailabilities (optional)
	fieldsInAppPurchaseImages := []string{"FieldsInAppPurchaseImages_example"} // []string | the fields to include for returned resources of type inAppPurchaseImages (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitImages := int32(56) // int32 | maximum number of related images returned (when they are included) (optional)
	limitInAppPurchaseLocalizations := int32(56) // int32 | maximum number of related inAppPurchaseLocalizations returned (when they are included) (optional)
	limitPricePoints := int32(56) // int32 | maximum number of related pricePoints returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2GetInstance(context.Background(), id).FieldsInAppPurchases(fieldsInAppPurchases).FieldsInAppPurchaseLocalizations(fieldsInAppPurchaseLocalizations).FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints).FieldsInAppPurchaseContents(fieldsInAppPurchaseContents).FieldsInAppPurchaseAppStoreReviewScreenshots(fieldsInAppPurchaseAppStoreReviewScreenshots).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsInAppPurchasePriceSchedules(fieldsInAppPurchasePriceSchedules).FieldsInAppPurchaseAvailabilities(fieldsInAppPurchaseAvailabilities).FieldsInAppPurchaseImages(fieldsInAppPurchaseImages).Include(include).LimitImages(limitImages).LimitInAppPurchaseLocalizations(limitInAppPurchaseLocalizations).LimitPricePoints(limitPricePoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2GetInstance`: InAppPurchaseV2Response
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **fieldsInAppPurchaseLocalizations** | **[]string** | the fields to include for returned resources of type inAppPurchaseLocalizations | 
 **fieldsInAppPurchasePricePoints** | **[]string** | the fields to include for returned resources of type inAppPurchasePricePoints | 
 **fieldsInAppPurchaseContents** | **[]string** | the fields to include for returned resources of type inAppPurchaseContents | 
 **fieldsInAppPurchaseAppStoreReviewScreenshots** | **[]string** | the fields to include for returned resources of type inAppPurchaseAppStoreReviewScreenshots | 
 **fieldsPromotedPurchases** | **[]string** | the fields to include for returned resources of type promotedPurchases | 
 **fieldsInAppPurchasePriceSchedules** | **[]string** | the fields to include for returned resources of type inAppPurchasePriceSchedules | 
 **fieldsInAppPurchaseAvailabilities** | **[]string** | the fields to include for returned resources of type inAppPurchaseAvailabilities | 
 **fieldsInAppPurchaseImages** | **[]string** | the fields to include for returned resources of type inAppPurchaseImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitImages** | **int32** | maximum number of related images returned (when they are included) | 
 **limitInAppPurchaseLocalizations** | **int32** | maximum number of related inAppPurchaseLocalizations returned (when they are included) | 
 **limitPricePoints** | **int32** | maximum number of related pricePoints returned (when they are included) | 

### Return type

[**InAppPurchaseV2Response**](InAppPurchaseV2Response.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2IapPriceScheduleGetToOneRelated

> InAppPurchasePriceScheduleResponse InAppPurchasesV2IapPriceScheduleGetToOneRelated(ctx, id).FieldsInAppPurchasePriceSchedules(fieldsInAppPurchasePriceSchedules).FieldsInAppPurchases(fieldsInAppPurchases).FieldsTerritories(fieldsTerritories).FieldsInAppPurchasePrices(fieldsInAppPurchasePrices).Include(include).LimitManualPrices(limitManualPrices).LimitAutomaticPrices(limitAutomaticPrices).Execute()



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
	fieldsInAppPurchasePriceSchedules := []string{"FieldsInAppPurchasePriceSchedules_example"} // []string | the fields to include for returned resources of type inAppPurchasePriceSchedules (optional)
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	fieldsInAppPurchasePrices := []string{"FieldsInAppPurchasePrices_example"} // []string | the fields to include for returned resources of type inAppPurchasePrices (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitManualPrices := int32(56) // int32 | maximum number of related manualPrices returned (when they are included) (optional)
	limitAutomaticPrices := int32(56) // int32 | maximum number of related automaticPrices returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2IapPriceScheduleGetToOneRelated(context.Background(), id).FieldsInAppPurchasePriceSchedules(fieldsInAppPurchasePriceSchedules).FieldsInAppPurchases(fieldsInAppPurchases).FieldsTerritories(fieldsTerritories).FieldsInAppPurchasePrices(fieldsInAppPurchasePrices).Include(include).LimitManualPrices(limitManualPrices).LimitAutomaticPrices(limitAutomaticPrices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2IapPriceScheduleGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2IapPriceScheduleGetToOneRelated`: InAppPurchasePriceScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2IapPriceScheduleGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2IapPriceScheduleGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchasePriceSchedules** | **[]string** | the fields to include for returned resources of type inAppPurchasePriceSchedules | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **fieldsInAppPurchasePrices** | **[]string** | the fields to include for returned resources of type inAppPurchasePrices | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitManualPrices** | **int32** | maximum number of related manualPrices returned (when they are included) | 
 **limitAutomaticPrices** | **int32** | maximum number of related automaticPrices returned (when they are included) | 

### Return type

[**InAppPurchasePriceScheduleResponse**](InAppPurchasePriceScheduleResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2ImagesGetToManyRelated

> InAppPurchaseImagesResponse InAppPurchasesV2ImagesGetToManyRelated(ctx, id).FieldsInAppPurchaseImages(fieldsInAppPurchaseImages).FieldsInAppPurchases(fieldsInAppPurchases).Limit(limit).Include(include).Execute()



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
	fieldsInAppPurchaseImages := []string{"FieldsInAppPurchaseImages_example"} // []string | the fields to include for returned resources of type inAppPurchaseImages (optional)
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2ImagesGetToManyRelated(context.Background(), id).FieldsInAppPurchaseImages(fieldsInAppPurchaseImages).FieldsInAppPurchases(fieldsInAppPurchases).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2ImagesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2ImagesGetToManyRelated`: InAppPurchaseImagesResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2ImagesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2ImagesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchaseImages** | **[]string** | the fields to include for returned resources of type inAppPurchaseImages | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchaseImagesResponse**](InAppPurchaseImagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelated

> InAppPurchaseAvailabilityResponse InAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelated(ctx, id).FieldsInAppPurchaseAvailabilities(fieldsInAppPurchaseAvailabilities).FieldsTerritories(fieldsTerritories).Include(include).LimitAvailableTerritories(limitAvailableTerritories).Execute()



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
	fieldsInAppPurchaseAvailabilities := []string{"FieldsInAppPurchaseAvailabilities_example"} // []string | the fields to include for returned resources of type inAppPurchaseAvailabilities (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAvailableTerritories := int32(56) // int32 | maximum number of related availableTerritories returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelated(context.Background(), id).FieldsInAppPurchaseAvailabilities(fieldsInAppPurchaseAvailabilities).FieldsTerritories(fieldsTerritories).Include(include).LimitAvailableTerritories(limitAvailableTerritories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelated`: InAppPurchaseAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2InAppPurchaseAvailabilityGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchaseAvailabilities** | **[]string** | the fields to include for returned resources of type inAppPurchaseAvailabilities | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAvailableTerritories** | **int32** | maximum number of related availableTerritories returned (when they are included) | 

### Return type

[**InAppPurchaseAvailabilityResponse**](InAppPurchaseAvailabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelated

> InAppPurchaseLocalizationsResponse InAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelated(ctx, id).FieldsInAppPurchaseLocalizations(fieldsInAppPurchaseLocalizations).FieldsInAppPurchases(fieldsInAppPurchases).Limit(limit).Include(include).Execute()



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
	fieldsInAppPurchaseLocalizations := []string{"FieldsInAppPurchaseLocalizations_example"} // []string | the fields to include for returned resources of type inAppPurchaseLocalizations (optional)
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelated(context.Background(), id).FieldsInAppPurchaseLocalizations(fieldsInAppPurchaseLocalizations).FieldsInAppPurchases(fieldsInAppPurchases).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelated`: InAppPurchaseLocalizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2InAppPurchaseLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchaseLocalizations** | **[]string** | the fields to include for returned resources of type inAppPurchaseLocalizations | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchaseLocalizationsResponse**](InAppPurchaseLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2PricePointsGetToManyRelated

> InAppPurchasePricePointsResponse InAppPurchasesV2PricePointsGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()



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
	filterTerritory := []string{"Inner_example"} // []string | filter by id(s) of related 'territory' (optional)
	fieldsInAppPurchasePricePoints := []string{"FieldsInAppPurchasePricePoints_example"} // []string | the fields to include for returned resources of type inAppPurchasePricePoints (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2PricePointsGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2PricePointsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2PricePointsGetToManyRelated`: InAppPurchasePricePointsResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2PricePointsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2PricePointsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsInAppPurchasePricePoints** | **[]string** | the fields to include for returned resources of type inAppPurchasePricePoints | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchasePricePointsResponse**](InAppPurchasePricePointsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2PromotedPurchaseGetToOneRelated

> PromotedPurchaseResponse InAppPurchasesV2PromotedPurchaseGetToOneRelated(ctx, id).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsInAppPurchases(fieldsInAppPurchases).FieldsSubscriptions(fieldsSubscriptions).FieldsPromotedPurchaseImages(fieldsPromotedPurchaseImages).Include(include).LimitPromotionImages(limitPromotionImages).Execute()



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
	fieldsPromotedPurchases := []string{"FieldsPromotedPurchases_example"} // []string | the fields to include for returned resources of type promotedPurchases (optional)
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	fieldsPromotedPurchaseImages := []string{"FieldsPromotedPurchaseImages_example"} // []string | the fields to include for returned resources of type promotedPurchaseImages (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitPromotionImages := int32(56) // int32 | maximum number of related promotionImages returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2PromotedPurchaseGetToOneRelated(context.Background(), id).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsInAppPurchases(fieldsInAppPurchases).FieldsSubscriptions(fieldsSubscriptions).FieldsPromotedPurchaseImages(fieldsPromotedPurchaseImages).Include(include).LimitPromotionImages(limitPromotionImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2PromotedPurchaseGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2PromotedPurchaseGetToOneRelated`: PromotedPurchaseResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2PromotedPurchaseGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2PromotedPurchaseGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPromotedPurchases** | **[]string** | the fields to include for returned resources of type promotedPurchases | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **fieldsPromotedPurchaseImages** | **[]string** | the fields to include for returned resources of type promotedPurchaseImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitPromotionImages** | **int32** | maximum number of related promotionImages returned (when they are included) | 

### Return type

[**PromotedPurchaseResponse**](PromotedPurchaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasesV2UpdateInstance

> InAppPurchaseV2Response InAppPurchasesV2UpdateInstance(ctx, id).InAppPurchaseV2UpdateRequest(inAppPurchaseV2UpdateRequest).Execute()



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
	inAppPurchaseV2UpdateRequest := *openapiclient.NewInAppPurchaseV2UpdateRequest(*openapiclient.NewInAppPurchaseV2UpdateRequestData("Type_example", "Id_example")) // InAppPurchaseV2UpdateRequest | InAppPurchase representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasesAPI.InAppPurchasesV2UpdateInstance(context.Background(), id).InAppPurchaseV2UpdateRequest(inAppPurchaseV2UpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasesAPI.InAppPurchasesV2UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasesV2UpdateInstance`: InAppPurchaseV2Response
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasesAPI.InAppPurchasesV2UpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasesV2UpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inAppPurchaseV2UpdateRequest** | [**InAppPurchaseV2UpdateRequest**](InAppPurchaseV2UpdateRequest.md) | InAppPurchase representation | 

### Return type

[**InAppPurchaseV2Response**](InAppPurchaseV2Response.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

