# \SubscriptionAvailabilitiesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated**](SubscriptionAvailabilitiesAPI.md#SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated) | **Get** /v1/subscriptionAvailabilities/{id}/availableTerritories | 
[**SubscriptionAvailabilitiesCreateInstance**](SubscriptionAvailabilitiesAPI.md#SubscriptionAvailabilitiesCreateInstance) | **Post** /v1/subscriptionAvailabilities | 
[**SubscriptionAvailabilitiesGetInstance**](SubscriptionAvailabilitiesAPI.md#SubscriptionAvailabilitiesGetInstance) | **Get** /v1/subscriptionAvailabilities/{id} | 



## SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated

> TerritoriesResponse SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated(ctx, id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated(context.Background(), id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated`: TerritoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**TerritoriesResponse**](TerritoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionAvailabilitiesCreateInstance

> SubscriptionAvailabilityResponse SubscriptionAvailabilitiesCreateInstance(ctx).SubscriptionAvailabilityCreateRequest(subscriptionAvailabilityCreateRequest).Execute()



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
	subscriptionAvailabilityCreateRequest := *openapiclient.NewSubscriptionAvailabilityCreateRequest(*openapiclient.NewSubscriptionAvailabilityCreateRequestData("Type_example", *openapiclient.NewAppAvailabilityV2CreateRequestDataAttributes(false), *openapiclient.NewSubscriptionAvailabilityCreateRequestDataRelationships(*openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(*openapiclient.NewPromotedPurchaseRelationshipsSubscriptionData("Type_example", "Id_example")), *openapiclient.NewAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories([]openapiclient.AppAvailabilityRelationshipsAvailableTerritoriesDataInner{*openapiclient.NewAppAvailabilityRelationshipsAvailableTerritoriesDataInner("Type_example", "Id_example")})))) // SubscriptionAvailabilityCreateRequest | SubscriptionAvailability representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesCreateInstance(context.Background()).SubscriptionAvailabilityCreateRequest(subscriptionAvailabilityCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAvailabilitiesCreateInstance`: SubscriptionAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAvailabilitiesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionAvailabilityCreateRequest** | [**SubscriptionAvailabilityCreateRequest**](SubscriptionAvailabilityCreateRequest.md) | SubscriptionAvailability representation | 

### Return type

[**SubscriptionAvailabilityResponse**](SubscriptionAvailabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionAvailabilitiesGetInstance

> SubscriptionAvailabilityResponse SubscriptionAvailabilitiesGetInstance(ctx, id).FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities).FieldsTerritories(fieldsTerritories).Include(include).LimitAvailableTerritories(limitAvailableTerritories).Execute()



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
	fieldsSubscriptionAvailabilities := []string{"FieldsSubscriptionAvailabilities_example"} // []string | the fields to include for returned resources of type subscriptionAvailabilities (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAvailableTerritories := int32(56) // int32 | maximum number of related availableTerritories returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesGetInstance(context.Background(), id).FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities).FieldsTerritories(fieldsTerritories).Include(include).LimitAvailableTerritories(limitAvailableTerritories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAvailabilitiesGetInstance`: SubscriptionAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAvailabilitiesAPI.SubscriptionAvailabilitiesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAvailabilitiesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionAvailabilities** | **[]string** | the fields to include for returned resources of type subscriptionAvailabilities | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAvailableTerritories** | **int32** | maximum number of related availableTerritories returned (when they are included) | 

### Return type

[**SubscriptionAvailabilityResponse**](SubscriptionAvailabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

