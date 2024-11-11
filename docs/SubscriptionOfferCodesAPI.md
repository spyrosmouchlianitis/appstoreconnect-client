# \SubscriptionOfferCodesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionOfferCodesCreateInstance**](SubscriptionOfferCodesAPI.md#SubscriptionOfferCodesCreateInstance) | **Post** /v1/subscriptionOfferCodes | 
[**SubscriptionOfferCodesCustomCodesGetToManyRelated**](SubscriptionOfferCodesAPI.md#SubscriptionOfferCodesCustomCodesGetToManyRelated) | **Get** /v1/subscriptionOfferCodes/{id}/customCodes | 
[**SubscriptionOfferCodesGetInstance**](SubscriptionOfferCodesAPI.md#SubscriptionOfferCodesGetInstance) | **Get** /v1/subscriptionOfferCodes/{id} | 
[**SubscriptionOfferCodesOneTimeUseCodesGetToManyRelated**](SubscriptionOfferCodesAPI.md#SubscriptionOfferCodesOneTimeUseCodesGetToManyRelated) | **Get** /v1/subscriptionOfferCodes/{id}/oneTimeUseCodes | 
[**SubscriptionOfferCodesPricesGetToManyRelated**](SubscriptionOfferCodesAPI.md#SubscriptionOfferCodesPricesGetToManyRelated) | **Get** /v1/subscriptionOfferCodes/{id}/prices | 
[**SubscriptionOfferCodesUpdateInstance**](SubscriptionOfferCodesAPI.md#SubscriptionOfferCodesUpdateInstance) | **Patch** /v1/subscriptionOfferCodes/{id} | 



## SubscriptionOfferCodesCreateInstance

> SubscriptionOfferCodeResponse SubscriptionOfferCodesCreateInstance(ctx).SubscriptionOfferCodeCreateRequest(subscriptionOfferCodeCreateRequest).Execute()



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
	subscriptionOfferCodeCreateRequest := *openapiclient.NewSubscriptionOfferCodeCreateRequest(*openapiclient.NewSubscriptionOfferCodeCreateRequestData("Type_example", *openapiclient.NewSubscriptionOfferCodeCreateRequestDataAttributes("Name_example", []openapiclient.SubscriptionCustomerEligibility{openapiclient.SubscriptionCustomerEligibility("NEW")}, openapiclient.SubscriptionOfferEligibility("STACK_WITH_INTRO_OFFERS"), openapiclient.SubscriptionOfferDuration("THREE_DAYS"), openapiclient.SubscriptionOfferMode("PAY_AS_YOU_GO"), int32(123)), *openapiclient.NewSubscriptionOfferCodeCreateRequestDataRelationships(*openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(*openapiclient.NewPromotedPurchaseRelationshipsSubscriptionData("Type_example", "Id_example")), *openapiclient.NewSubscriptionOfferCodeCreateRequestDataRelationshipsPrices([]openapiclient.SubscriptionOfferCodeRelationshipsPricesDataInner{*openapiclient.NewSubscriptionOfferCodeRelationshipsPricesDataInner("Type_example", "Id_example")})))) // SubscriptionOfferCodeCreateRequest | SubscriptionOfferCode representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionOfferCodesAPI.SubscriptionOfferCodesCreateInstance(context.Background()).SubscriptionOfferCodeCreateRequest(subscriptionOfferCodeCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionOfferCodesAPI.SubscriptionOfferCodesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOfferCodesCreateInstance`: SubscriptionOfferCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionOfferCodesAPI.SubscriptionOfferCodesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOfferCodesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionOfferCodeCreateRequest** | [**SubscriptionOfferCodeCreateRequest**](SubscriptionOfferCodeCreateRequest.md) | SubscriptionOfferCode representation | 

### Return type

[**SubscriptionOfferCodeResponse**](SubscriptionOfferCodeResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionOfferCodesCustomCodesGetToManyRelated

> SubscriptionOfferCodeCustomCodesResponse SubscriptionOfferCodesCustomCodesGetToManyRelated(ctx, id).FieldsSubscriptionOfferCodeCustomCodes(fieldsSubscriptionOfferCodeCustomCodes).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).Limit(limit).Include(include).Execute()



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
	fieldsSubscriptionOfferCodeCustomCodes := []string{"FieldsSubscriptionOfferCodeCustomCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodeCustomCodes (optional)
	fieldsSubscriptionOfferCodes := []string{"FieldsSubscriptionOfferCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodes (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionOfferCodesAPI.SubscriptionOfferCodesCustomCodesGetToManyRelated(context.Background(), id).FieldsSubscriptionOfferCodeCustomCodes(fieldsSubscriptionOfferCodeCustomCodes).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionOfferCodesAPI.SubscriptionOfferCodesCustomCodesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOfferCodesCustomCodesGetToManyRelated`: SubscriptionOfferCodeCustomCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionOfferCodesAPI.SubscriptionOfferCodesCustomCodesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOfferCodesCustomCodesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionOfferCodeCustomCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodeCustomCodes | 
 **fieldsSubscriptionOfferCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodes | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionOfferCodeCustomCodesResponse**](SubscriptionOfferCodeCustomCodesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionOfferCodesGetInstance

> SubscriptionOfferCodeResponse SubscriptionOfferCodesGetInstance(ctx, id).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).FieldsSubscriptionOfferCodeOneTimeUseCodes(fieldsSubscriptionOfferCodeOneTimeUseCodes).FieldsSubscriptionOfferCodeCustomCodes(fieldsSubscriptionOfferCodeCustomCodes).FieldsSubscriptionOfferCodePrices(fieldsSubscriptionOfferCodePrices).Include(include).LimitCustomCodes(limitCustomCodes).LimitOneTimeUseCodes(limitOneTimeUseCodes).LimitPrices(limitPrices).Execute()



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
	fieldsSubscriptionOfferCodes := []string{"FieldsSubscriptionOfferCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodes (optional)
	fieldsSubscriptionOfferCodeOneTimeUseCodes := []string{"FieldsSubscriptionOfferCodeOneTimeUseCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodeOneTimeUseCodes (optional)
	fieldsSubscriptionOfferCodeCustomCodes := []string{"FieldsSubscriptionOfferCodeCustomCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodeCustomCodes (optional)
	fieldsSubscriptionOfferCodePrices := []string{"FieldsSubscriptionOfferCodePrices_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodePrices (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitCustomCodes := int32(56) // int32 | maximum number of related customCodes returned (when they are included) (optional)
	limitOneTimeUseCodes := int32(56) // int32 | maximum number of related oneTimeUseCodes returned (when they are included) (optional)
	limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionOfferCodesAPI.SubscriptionOfferCodesGetInstance(context.Background(), id).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).FieldsSubscriptionOfferCodeOneTimeUseCodes(fieldsSubscriptionOfferCodeOneTimeUseCodes).FieldsSubscriptionOfferCodeCustomCodes(fieldsSubscriptionOfferCodeCustomCodes).FieldsSubscriptionOfferCodePrices(fieldsSubscriptionOfferCodePrices).Include(include).LimitCustomCodes(limitCustomCodes).LimitOneTimeUseCodes(limitOneTimeUseCodes).LimitPrices(limitPrices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionOfferCodesAPI.SubscriptionOfferCodesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOfferCodesGetInstance`: SubscriptionOfferCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionOfferCodesAPI.SubscriptionOfferCodesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOfferCodesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionOfferCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodes | 
 **fieldsSubscriptionOfferCodeOneTimeUseCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodeOneTimeUseCodes | 
 **fieldsSubscriptionOfferCodeCustomCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodeCustomCodes | 
 **fieldsSubscriptionOfferCodePrices** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodePrices | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitCustomCodes** | **int32** | maximum number of related customCodes returned (when they are included) | 
 **limitOneTimeUseCodes** | **int32** | maximum number of related oneTimeUseCodes returned (when they are included) | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 

### Return type

[**SubscriptionOfferCodeResponse**](SubscriptionOfferCodeResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionOfferCodesOneTimeUseCodesGetToManyRelated

> SubscriptionOfferCodeOneTimeUseCodesResponse SubscriptionOfferCodesOneTimeUseCodesGetToManyRelated(ctx, id).FieldsSubscriptionOfferCodeOneTimeUseCodes(fieldsSubscriptionOfferCodeOneTimeUseCodes).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).Limit(limit).Include(include).Execute()



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
	fieldsSubscriptionOfferCodeOneTimeUseCodes := []string{"FieldsSubscriptionOfferCodeOneTimeUseCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodeOneTimeUseCodes (optional)
	fieldsSubscriptionOfferCodes := []string{"FieldsSubscriptionOfferCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodes (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionOfferCodesAPI.SubscriptionOfferCodesOneTimeUseCodesGetToManyRelated(context.Background(), id).FieldsSubscriptionOfferCodeOneTimeUseCodes(fieldsSubscriptionOfferCodeOneTimeUseCodes).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionOfferCodesAPI.SubscriptionOfferCodesOneTimeUseCodesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOfferCodesOneTimeUseCodesGetToManyRelated`: SubscriptionOfferCodeOneTimeUseCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionOfferCodesAPI.SubscriptionOfferCodesOneTimeUseCodesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOfferCodesOneTimeUseCodesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionOfferCodeOneTimeUseCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodeOneTimeUseCodes | 
 **fieldsSubscriptionOfferCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodes | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionOfferCodeOneTimeUseCodesResponse**](SubscriptionOfferCodeOneTimeUseCodesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionOfferCodesPricesGetToManyRelated

> SubscriptionOfferCodePricesResponse SubscriptionOfferCodesPricesGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsSubscriptionOfferCodePrices(fieldsSubscriptionOfferCodePrices).FieldsTerritories(fieldsTerritories).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Limit(limit).Include(include).Execute()



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
	fieldsSubscriptionOfferCodePrices := []string{"FieldsSubscriptionOfferCodePrices_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodePrices (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	fieldsSubscriptionPricePoints := []string{"FieldsSubscriptionPricePoints_example"} // []string | the fields to include for returned resources of type subscriptionPricePoints (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionOfferCodesAPI.SubscriptionOfferCodesPricesGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsSubscriptionOfferCodePrices(fieldsSubscriptionOfferCodePrices).FieldsTerritories(fieldsTerritories).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionOfferCodesAPI.SubscriptionOfferCodesPricesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOfferCodesPricesGetToManyRelated`: SubscriptionOfferCodePricesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionOfferCodesAPI.SubscriptionOfferCodesPricesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOfferCodesPricesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsSubscriptionOfferCodePrices** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodePrices | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **fieldsSubscriptionPricePoints** | **[]string** | the fields to include for returned resources of type subscriptionPricePoints | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionOfferCodePricesResponse**](SubscriptionOfferCodePricesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionOfferCodesUpdateInstance

> SubscriptionOfferCodeResponse SubscriptionOfferCodesUpdateInstance(ctx, id).SubscriptionOfferCodeUpdateRequest(subscriptionOfferCodeUpdateRequest).Execute()



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
	subscriptionOfferCodeUpdateRequest := *openapiclient.NewSubscriptionOfferCodeUpdateRequest(*openapiclient.NewSubscriptionOfferCodeUpdateRequestData("Type_example", "Id_example")) // SubscriptionOfferCodeUpdateRequest | SubscriptionOfferCode representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionOfferCodesAPI.SubscriptionOfferCodesUpdateInstance(context.Background(), id).SubscriptionOfferCodeUpdateRequest(subscriptionOfferCodeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionOfferCodesAPI.SubscriptionOfferCodesUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOfferCodesUpdateInstance`: SubscriptionOfferCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionOfferCodesAPI.SubscriptionOfferCodesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOfferCodesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionOfferCodeUpdateRequest** | [**SubscriptionOfferCodeUpdateRequest**](SubscriptionOfferCodeUpdateRequest.md) | SubscriptionOfferCode representation | 

### Return type

[**SubscriptionOfferCodeResponse**](SubscriptionOfferCodeResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

