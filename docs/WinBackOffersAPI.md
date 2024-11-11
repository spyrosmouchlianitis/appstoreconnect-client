# \WinBackOffersAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WinBackOffersCreateInstance**](WinBackOffersAPI.md#WinBackOffersCreateInstance) | **Post** /v1/winBackOffers | 
[**WinBackOffersDeleteInstance**](WinBackOffersAPI.md#WinBackOffersDeleteInstance) | **Delete** /v1/winBackOffers/{id} | 
[**WinBackOffersGetInstance**](WinBackOffersAPI.md#WinBackOffersGetInstance) | **Get** /v1/winBackOffers/{id} | 
[**WinBackOffersPricesGetToManyRelated**](WinBackOffersAPI.md#WinBackOffersPricesGetToManyRelated) | **Get** /v1/winBackOffers/{id}/prices | 
[**WinBackOffersUpdateInstance**](WinBackOffersAPI.md#WinBackOffersUpdateInstance) | **Patch** /v1/winBackOffers/{id} | 



## WinBackOffersCreateInstance

> WinBackOfferResponse WinBackOffersCreateInstance(ctx).WinBackOfferCreateRequest(winBackOfferCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	winBackOfferCreateRequest := *openapiclient.NewWinBackOfferCreateRequest(*openapiclient.NewWinBackOfferCreateRequestData("Type_example", *openapiclient.NewWinBackOfferCreateRequestDataAttributes("ReferenceName_example", "OfferId_example", openapiclient.SubscriptionOfferDuration("THREE_DAYS"), openapiclient.SubscriptionOfferMode("PAY_AS_YOU_GO"), int32(123), int32(123), *openapiclient.NewIntegerRange(), time.Now(), "Priority_example"), *openapiclient.NewWinBackOfferCreateRequestDataRelationships(*openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(*openapiclient.NewPromotedPurchaseRelationshipsSubscriptionData("Type_example", "Id_example")), *openapiclient.NewWinBackOfferCreateRequestDataRelationshipsPrices([]openapiclient.WinBackOfferRelationshipsPricesDataInner{*openapiclient.NewWinBackOfferRelationshipsPricesDataInner("Type_example", "Id_example")})))) // WinBackOfferCreateRequest | WinBackOffer representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WinBackOffersAPI.WinBackOffersCreateInstance(context.Background()).WinBackOfferCreateRequest(winBackOfferCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WinBackOffersAPI.WinBackOffersCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WinBackOffersCreateInstance`: WinBackOfferResponse
	fmt.Fprintf(os.Stdout, "Response from `WinBackOffersAPI.WinBackOffersCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWinBackOffersCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **winBackOfferCreateRequest** | [**WinBackOfferCreateRequest**](WinBackOfferCreateRequest.md) | WinBackOffer representation | 

### Return type

[**WinBackOfferResponse**](WinBackOfferResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WinBackOffersDeleteInstance

> WinBackOffersDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.WinBackOffersAPI.WinBackOffersDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WinBackOffersAPI.WinBackOffersDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWinBackOffersDeleteInstanceRequest struct via the builder pattern


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


## WinBackOffersGetInstance

> WinBackOfferResponse WinBackOffersGetInstance(ctx, id).FieldsWinBackOffers(fieldsWinBackOffers).FieldsWinBackOfferPrices(fieldsWinBackOfferPrices).Include(include).LimitPrices(limitPrices).Execute()



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
	fieldsWinBackOffers := []string{"FieldsWinBackOffers_example"} // []string | the fields to include for returned resources of type winBackOffers (optional)
	fieldsWinBackOfferPrices := []string{"FieldsWinBackOfferPrices_example"} // []string | the fields to include for returned resources of type winBackOfferPrices (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WinBackOffersAPI.WinBackOffersGetInstance(context.Background(), id).FieldsWinBackOffers(fieldsWinBackOffers).FieldsWinBackOfferPrices(fieldsWinBackOfferPrices).Include(include).LimitPrices(limitPrices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WinBackOffersAPI.WinBackOffersGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WinBackOffersGetInstance`: WinBackOfferResponse
	fmt.Fprintf(os.Stdout, "Response from `WinBackOffersAPI.WinBackOffersGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiWinBackOffersGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsWinBackOffers** | **[]string** | the fields to include for returned resources of type winBackOffers | 
 **fieldsWinBackOfferPrices** | **[]string** | the fields to include for returned resources of type winBackOfferPrices | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 

### Return type

[**WinBackOfferResponse**](WinBackOfferResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WinBackOffersPricesGetToManyRelated

> WinBackOfferPricesResponse WinBackOffersPricesGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsWinBackOfferPrices(fieldsWinBackOfferPrices).FieldsTerritories(fieldsTerritories).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Limit(limit).Include(include).Execute()



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
	fieldsWinBackOfferPrices := []string{"FieldsWinBackOfferPrices_example"} // []string | the fields to include for returned resources of type winBackOfferPrices (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	fieldsSubscriptionPricePoints := []string{"FieldsSubscriptionPricePoints_example"} // []string | the fields to include for returned resources of type subscriptionPricePoints (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WinBackOffersAPI.WinBackOffersPricesGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsWinBackOfferPrices(fieldsWinBackOfferPrices).FieldsTerritories(fieldsTerritories).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WinBackOffersAPI.WinBackOffersPricesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WinBackOffersPricesGetToManyRelated`: WinBackOfferPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `WinBackOffersAPI.WinBackOffersPricesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiWinBackOffersPricesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsWinBackOfferPrices** | **[]string** | the fields to include for returned resources of type winBackOfferPrices | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **fieldsSubscriptionPricePoints** | **[]string** | the fields to include for returned resources of type subscriptionPricePoints | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**WinBackOfferPricesResponse**](WinBackOfferPricesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WinBackOffersUpdateInstance

> WinBackOfferResponse WinBackOffersUpdateInstance(ctx, id).WinBackOfferUpdateRequest(winBackOfferUpdateRequest).Execute()



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
	winBackOfferUpdateRequest := *openapiclient.NewWinBackOfferUpdateRequest(*openapiclient.NewWinBackOfferUpdateRequestData("Type_example", "Id_example")) // WinBackOfferUpdateRequest | WinBackOffer representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WinBackOffersAPI.WinBackOffersUpdateInstance(context.Background(), id).WinBackOfferUpdateRequest(winBackOfferUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WinBackOffersAPI.WinBackOffersUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WinBackOffersUpdateInstance`: WinBackOfferResponse
	fmt.Fprintf(os.Stdout, "Response from `WinBackOffersAPI.WinBackOffersUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiWinBackOffersUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **winBackOfferUpdateRequest** | [**WinBackOfferUpdateRequest**](WinBackOfferUpdateRequest.md) | WinBackOffer representation | 

### Return type

[**WinBackOfferResponse**](WinBackOfferResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

