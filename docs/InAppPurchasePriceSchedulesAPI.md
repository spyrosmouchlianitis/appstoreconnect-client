# \InAppPurchasePriceSchedulesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated**](InAppPurchasePriceSchedulesAPI.md#InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated) | **Get** /v1/inAppPurchasePriceSchedules/{id}/automaticPrices | 
[**InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated**](InAppPurchasePriceSchedulesAPI.md#InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated) | **Get** /v1/inAppPurchasePriceSchedules/{id}/baseTerritory | 
[**InAppPurchasePriceSchedulesCreateInstance**](InAppPurchasePriceSchedulesAPI.md#InAppPurchasePriceSchedulesCreateInstance) | **Post** /v1/inAppPurchasePriceSchedules | 
[**InAppPurchasePriceSchedulesGetInstance**](InAppPurchasePriceSchedulesAPI.md#InAppPurchasePriceSchedulesGetInstance) | **Get** /v1/inAppPurchasePriceSchedules/{id} | 
[**InAppPurchasePriceSchedulesManualPricesGetToManyRelated**](InAppPurchasePriceSchedulesAPI.md#InAppPurchasePriceSchedulesManualPricesGetToManyRelated) | **Get** /v1/inAppPurchasePriceSchedules/{id}/manualPrices | 



## InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated

> InAppPurchasePricesResponse InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsInAppPurchasePrices(fieldsInAppPurchasePrices).FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()



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
	filterTerritory := []string{"Inner_example"} // []string | filter by id(s) of related 'territory' (optional)
	fieldsInAppPurchasePrices := []string{"FieldsInAppPurchasePrices_example"} // []string | the fields to include for returned resources of type inAppPurchasePrices (optional)
	fieldsInAppPurchasePricePoints := []string{"FieldsInAppPurchasePricePoints_example"} // []string | the fields to include for returned resources of type inAppPurchasePricePoints (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsInAppPurchasePrices(fieldsInAppPurchasePrices).FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated`: InAppPurchasePricesResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasePriceSchedulesAutomaticPricesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsInAppPurchasePrices** | **[]string** | the fields to include for returned resources of type inAppPurchasePrices | 
 **fieldsInAppPurchasePricePoints** | **[]string** | the fields to include for returned resources of type inAppPurchasePricePoints | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchasePricesResponse**](InAppPurchasePricesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated

> TerritoryResponse InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated(ctx, id).FieldsTerritories(fieldsTerritories).Execute()



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
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated(context.Background(), id).FieldsTerritories(fieldsTerritories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated`: TerritoryResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasePriceSchedulesBaseTerritoryGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 

### Return type

[**TerritoryResponse**](TerritoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasePriceSchedulesCreateInstance

> InAppPurchasePriceScheduleResponse InAppPurchasePriceSchedulesCreateInstance(ctx).InAppPurchasePriceScheduleCreateRequest(inAppPurchasePriceScheduleCreateRequest).Execute()



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
	inAppPurchasePriceScheduleCreateRequest := *openapiclient.NewInAppPurchasePriceScheduleCreateRequest(*openapiclient.NewInAppPurchasePriceScheduleCreateRequestData("Type_example", *openapiclient.NewInAppPurchasePriceScheduleCreateRequestDataRelationships(*openapiclient.NewInAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2(*openapiclient.NewAppRelationshipsInAppPurchasesDataInner("Type_example", "Id_example")), *openapiclient.NewAppPriceScheduleCreateRequestDataRelationshipsBaseTerritory(*openapiclient.NewAppAvailabilityRelationshipsAvailableTerritoriesDataInner("Type_example", "Id_example")), *openapiclient.NewInAppPurchasePriceScheduleCreateRequestDataRelationshipsManualPrices([]openapiclient.InAppPurchasePriceScheduleRelationshipsManualPricesDataInner{*openapiclient.NewInAppPurchasePriceScheduleRelationshipsManualPricesDataInner("Type_example", "Id_example")})))) // InAppPurchasePriceScheduleCreateRequest | InAppPurchasePriceSchedule representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesCreateInstance(context.Background()).InAppPurchasePriceScheduleCreateRequest(inAppPurchasePriceScheduleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasePriceSchedulesCreateInstance`: InAppPurchasePriceScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasePriceSchedulesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inAppPurchasePriceScheduleCreateRequest** | [**InAppPurchasePriceScheduleCreateRequest**](InAppPurchasePriceScheduleCreateRequest.md) | InAppPurchasePriceSchedule representation | 

### Return type

[**InAppPurchasePriceScheduleResponse**](InAppPurchasePriceScheduleResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InAppPurchasePriceSchedulesGetInstance

> InAppPurchasePriceScheduleResponse InAppPurchasePriceSchedulesGetInstance(ctx, id).FieldsInAppPurchasePriceSchedules(fieldsInAppPurchasePriceSchedules).FieldsTerritories(fieldsTerritories).FieldsInAppPurchasePrices(fieldsInAppPurchasePrices).Include(include).LimitAutomaticPrices(limitAutomaticPrices).LimitManualPrices(limitManualPrices).Execute()



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
	fieldsInAppPurchasePriceSchedules := []string{"FieldsInAppPurchasePriceSchedules_example"} // []string | the fields to include for returned resources of type inAppPurchasePriceSchedules (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	fieldsInAppPurchasePrices := []string{"FieldsInAppPurchasePrices_example"} // []string | the fields to include for returned resources of type inAppPurchasePrices (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAutomaticPrices := int32(56) // int32 | maximum number of related automaticPrices returned (when they are included) (optional)
	limitManualPrices := int32(56) // int32 | maximum number of related manualPrices returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesGetInstance(context.Background(), id).FieldsInAppPurchasePriceSchedules(fieldsInAppPurchasePriceSchedules).FieldsTerritories(fieldsTerritories).FieldsInAppPurchasePrices(fieldsInAppPurchasePrices).Include(include).LimitAutomaticPrices(limitAutomaticPrices).LimitManualPrices(limitManualPrices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasePriceSchedulesGetInstance`: InAppPurchasePriceScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasePriceSchedulesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsInAppPurchasePriceSchedules** | **[]string** | the fields to include for returned resources of type inAppPurchasePriceSchedules | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **fieldsInAppPurchasePrices** | **[]string** | the fields to include for returned resources of type inAppPurchasePrices | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAutomaticPrices** | **int32** | maximum number of related automaticPrices returned (when they are included) | 
 **limitManualPrices** | **int32** | maximum number of related manualPrices returned (when they are included) | 

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


## InAppPurchasePriceSchedulesManualPricesGetToManyRelated

> InAppPurchasePricesResponse InAppPurchasePriceSchedulesManualPricesGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsInAppPurchasePrices(fieldsInAppPurchasePrices).FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()



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
	filterTerritory := []string{"Inner_example"} // []string | filter by id(s) of related 'territory' (optional)
	fieldsInAppPurchasePrices := []string{"FieldsInAppPurchasePrices_example"} // []string | the fields to include for returned resources of type inAppPurchasePrices (optional)
	fieldsInAppPurchasePricePoints := []string{"FieldsInAppPurchasePricePoints_example"} // []string | the fields to include for returned resources of type inAppPurchasePricePoints (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesManualPricesGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsInAppPurchasePrices(fieldsInAppPurchasePrices).FieldsInAppPurchasePricePoints(fieldsInAppPurchasePricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesManualPricesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InAppPurchasePriceSchedulesManualPricesGetToManyRelated`: InAppPurchasePricesResponse
	fmt.Fprintf(os.Stdout, "Response from `InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesManualPricesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiInAppPurchasePriceSchedulesManualPricesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsInAppPurchasePrices** | **[]string** | the fields to include for returned resources of type inAppPurchasePrices | 
 **fieldsInAppPurchasePricePoints** | **[]string** | the fields to include for returned resources of type inAppPurchasePricePoints | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**InAppPurchasePricesResponse**](InAppPurchasePricesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

