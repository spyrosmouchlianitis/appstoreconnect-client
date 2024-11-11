# \SubscriptionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsAppStoreReviewScreenshotGetToOneRelated**](SubscriptionsAPI.md#SubscriptionsAppStoreReviewScreenshotGetToOneRelated) | **Get** /v1/subscriptions/{id}/appStoreReviewScreenshot | 
[**SubscriptionsCreateInstance**](SubscriptionsAPI.md#SubscriptionsCreateInstance) | **Post** /v1/subscriptions | 
[**SubscriptionsDeleteInstance**](SubscriptionsAPI.md#SubscriptionsDeleteInstance) | **Delete** /v1/subscriptions/{id} | 
[**SubscriptionsGetInstance**](SubscriptionsAPI.md#SubscriptionsGetInstance) | **Get** /v1/subscriptions/{id} | 
[**SubscriptionsImagesGetToManyRelated**](SubscriptionsAPI.md#SubscriptionsImagesGetToManyRelated) | **Get** /v1/subscriptions/{id}/images | 
[**SubscriptionsIntroductoryOffersDeleteToManyRelationship**](SubscriptionsAPI.md#SubscriptionsIntroductoryOffersDeleteToManyRelationship) | **Delete** /v1/subscriptions/{id}/relationships/introductoryOffers | 
[**SubscriptionsIntroductoryOffersGetToManyRelated**](SubscriptionsAPI.md#SubscriptionsIntroductoryOffersGetToManyRelated) | **Get** /v1/subscriptions/{id}/introductoryOffers | 
[**SubscriptionsIntroductoryOffersGetToManyRelationship**](SubscriptionsAPI.md#SubscriptionsIntroductoryOffersGetToManyRelationship) | **Get** /v1/subscriptions/{id}/relationships/introductoryOffers | 
[**SubscriptionsOfferCodesGetToManyRelated**](SubscriptionsAPI.md#SubscriptionsOfferCodesGetToManyRelated) | **Get** /v1/subscriptions/{id}/offerCodes | 
[**SubscriptionsPricePointsGetToManyRelated**](SubscriptionsAPI.md#SubscriptionsPricePointsGetToManyRelated) | **Get** /v1/subscriptions/{id}/pricePoints | 
[**SubscriptionsPricesDeleteToManyRelationship**](SubscriptionsAPI.md#SubscriptionsPricesDeleteToManyRelationship) | **Delete** /v1/subscriptions/{id}/relationships/prices | 
[**SubscriptionsPricesGetToManyRelated**](SubscriptionsAPI.md#SubscriptionsPricesGetToManyRelated) | **Get** /v1/subscriptions/{id}/prices | 
[**SubscriptionsPricesGetToManyRelationship**](SubscriptionsAPI.md#SubscriptionsPricesGetToManyRelationship) | **Get** /v1/subscriptions/{id}/relationships/prices | 
[**SubscriptionsPromotedPurchaseGetToOneRelated**](SubscriptionsAPI.md#SubscriptionsPromotedPurchaseGetToOneRelated) | **Get** /v1/subscriptions/{id}/promotedPurchase | 
[**SubscriptionsPromotionalOffersGetToManyRelated**](SubscriptionsAPI.md#SubscriptionsPromotionalOffersGetToManyRelated) | **Get** /v1/subscriptions/{id}/promotionalOffers | 
[**SubscriptionsSubscriptionAvailabilityGetToOneRelated**](SubscriptionsAPI.md#SubscriptionsSubscriptionAvailabilityGetToOneRelated) | **Get** /v1/subscriptions/{id}/subscriptionAvailability | 
[**SubscriptionsSubscriptionLocalizationsGetToManyRelated**](SubscriptionsAPI.md#SubscriptionsSubscriptionLocalizationsGetToManyRelated) | **Get** /v1/subscriptions/{id}/subscriptionLocalizations | 
[**SubscriptionsUpdateInstance**](SubscriptionsAPI.md#SubscriptionsUpdateInstance) | **Patch** /v1/subscriptions/{id} | 
[**SubscriptionsWinBackOffersGetToManyRelated**](SubscriptionsAPI.md#SubscriptionsWinBackOffersGetToManyRelated) | **Get** /v1/subscriptions/{id}/winBackOffers | 



## SubscriptionsAppStoreReviewScreenshotGetToOneRelated

> SubscriptionAppStoreReviewScreenshotResponse SubscriptionsAppStoreReviewScreenshotGetToOneRelated(ctx, id).FieldsSubscriptionAppStoreReviewScreenshots(fieldsSubscriptionAppStoreReviewScreenshots).FieldsSubscriptions(fieldsSubscriptions).Include(include).Execute()



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
	fieldsSubscriptionAppStoreReviewScreenshots := []string{"FieldsSubscriptionAppStoreReviewScreenshots_example"} // []string | the fields to include for returned resources of type subscriptionAppStoreReviewScreenshots (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsAppStoreReviewScreenshotGetToOneRelated(context.Background(), id).FieldsSubscriptionAppStoreReviewScreenshots(fieldsSubscriptionAppStoreReviewScreenshots).FieldsSubscriptions(fieldsSubscriptions).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsAppStoreReviewScreenshotGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsAppStoreReviewScreenshotGetToOneRelated`: SubscriptionAppStoreReviewScreenshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsAppStoreReviewScreenshotGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsAppStoreReviewScreenshotGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionAppStoreReviewScreenshots** | **[]string** | the fields to include for returned resources of type subscriptionAppStoreReviewScreenshots | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionAppStoreReviewScreenshotResponse**](SubscriptionAppStoreReviewScreenshotResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsCreateInstance

> SubscriptionResponse SubscriptionsCreateInstance(ctx).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()



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
	subscriptionCreateRequest := *openapiclient.NewSubscriptionCreateRequest(*openapiclient.NewSubscriptionCreateRequestData("Type_example", *openapiclient.NewSubscriptionCreateRequestDataAttributes("Name_example", "ProductId_example"), *openapiclient.NewSubscriptionCreateRequestDataRelationships(*openapiclient.NewSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup(*openapiclient.NewAppRelationshipsSubscriptionGroupsDataInner("Type_example", "Id_example"))))) // SubscriptionCreateRequest | Subscription representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsCreateInstance(context.Background()).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsCreateInstance`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionCreateRequest** | [**SubscriptionCreateRequest**](SubscriptionCreateRequest.md) | Subscription representation | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsDeleteInstance

> SubscriptionsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.SubscriptionsAPI.SubscriptionsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionsDeleteInstanceRequest struct via the builder pattern


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


## SubscriptionsGetInstance

> SubscriptionResponse SubscriptionsGetInstance(ctx, id).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionLocalizations(fieldsSubscriptionLocalizations).FieldsSubscriptionAppStoreReviewScreenshots(fieldsSubscriptionAppStoreReviewScreenshots).FieldsSubscriptionIntroductoryOffers(fieldsSubscriptionIntroductoryOffers).FieldsSubscriptionPromotionalOffers(fieldsSubscriptionPromotionalOffers).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).FieldsSubscriptionPrices(fieldsSubscriptionPrices).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities).FieldsWinBackOffers(fieldsWinBackOffers).FieldsSubscriptionImages(fieldsSubscriptionImages).Include(include).LimitImages(limitImages).LimitIntroductoryOffers(limitIntroductoryOffers).LimitOfferCodes(limitOfferCodes).LimitPrices(limitPrices).LimitPromotionalOffers(limitPromotionalOffers).LimitSubscriptionLocalizations(limitSubscriptionLocalizations).LimitWinBackOffers(limitWinBackOffers).Execute()



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
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	fieldsSubscriptionLocalizations := []string{"FieldsSubscriptionLocalizations_example"} // []string | the fields to include for returned resources of type subscriptionLocalizations (optional)
	fieldsSubscriptionAppStoreReviewScreenshots := []string{"FieldsSubscriptionAppStoreReviewScreenshots_example"} // []string | the fields to include for returned resources of type subscriptionAppStoreReviewScreenshots (optional)
	fieldsSubscriptionIntroductoryOffers := []string{"FieldsSubscriptionIntroductoryOffers_example"} // []string | the fields to include for returned resources of type subscriptionIntroductoryOffers (optional)
	fieldsSubscriptionPromotionalOffers := []string{"FieldsSubscriptionPromotionalOffers_example"} // []string | the fields to include for returned resources of type subscriptionPromotionalOffers (optional)
	fieldsSubscriptionOfferCodes := []string{"FieldsSubscriptionOfferCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodes (optional)
	fieldsSubscriptionPrices := []string{"FieldsSubscriptionPrices_example"} // []string | the fields to include for returned resources of type subscriptionPrices (optional)
	fieldsPromotedPurchases := []string{"FieldsPromotedPurchases_example"} // []string | the fields to include for returned resources of type promotedPurchases (optional)
	fieldsSubscriptionAvailabilities := []string{"FieldsSubscriptionAvailabilities_example"} // []string | the fields to include for returned resources of type subscriptionAvailabilities (optional)
	fieldsWinBackOffers := []string{"FieldsWinBackOffers_example"} // []string | the fields to include for returned resources of type winBackOffers (optional)
	fieldsSubscriptionImages := []string{"FieldsSubscriptionImages_example"} // []string | the fields to include for returned resources of type subscriptionImages (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitImages := int32(56) // int32 | maximum number of related images returned (when they are included) (optional)
	limitIntroductoryOffers := int32(56) // int32 | maximum number of related introductoryOffers returned (when they are included) (optional)
	limitOfferCodes := int32(56) // int32 | maximum number of related offerCodes returned (when they are included) (optional)
	limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)
	limitPromotionalOffers := int32(56) // int32 | maximum number of related promotionalOffers returned (when they are included) (optional)
	limitSubscriptionLocalizations := int32(56) // int32 | maximum number of related subscriptionLocalizations returned (when they are included) (optional)
	limitWinBackOffers := int32(56) // int32 | maximum number of related winBackOffers returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGetInstance(context.Background(), id).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionLocalizations(fieldsSubscriptionLocalizations).FieldsSubscriptionAppStoreReviewScreenshots(fieldsSubscriptionAppStoreReviewScreenshots).FieldsSubscriptionIntroductoryOffers(fieldsSubscriptionIntroductoryOffers).FieldsSubscriptionPromotionalOffers(fieldsSubscriptionPromotionalOffers).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).FieldsSubscriptionPrices(fieldsSubscriptionPrices).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities).FieldsWinBackOffers(fieldsWinBackOffers).FieldsSubscriptionImages(fieldsSubscriptionImages).Include(include).LimitImages(limitImages).LimitIntroductoryOffers(limitIntroductoryOffers).LimitOfferCodes(limitOfferCodes).LimitPrices(limitPrices).LimitPromotionalOffers(limitPromotionalOffers).LimitSubscriptionLocalizations(limitSubscriptionLocalizations).LimitWinBackOffers(limitWinBackOffers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGetInstance`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **fieldsSubscriptionLocalizations** | **[]string** | the fields to include for returned resources of type subscriptionLocalizations | 
 **fieldsSubscriptionAppStoreReviewScreenshots** | **[]string** | the fields to include for returned resources of type subscriptionAppStoreReviewScreenshots | 
 **fieldsSubscriptionIntroductoryOffers** | **[]string** | the fields to include for returned resources of type subscriptionIntroductoryOffers | 
 **fieldsSubscriptionPromotionalOffers** | **[]string** | the fields to include for returned resources of type subscriptionPromotionalOffers | 
 **fieldsSubscriptionOfferCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodes | 
 **fieldsSubscriptionPrices** | **[]string** | the fields to include for returned resources of type subscriptionPrices | 
 **fieldsPromotedPurchases** | **[]string** | the fields to include for returned resources of type promotedPurchases | 
 **fieldsSubscriptionAvailabilities** | **[]string** | the fields to include for returned resources of type subscriptionAvailabilities | 
 **fieldsWinBackOffers** | **[]string** | the fields to include for returned resources of type winBackOffers | 
 **fieldsSubscriptionImages** | **[]string** | the fields to include for returned resources of type subscriptionImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitImages** | **int32** | maximum number of related images returned (when they are included) | 
 **limitIntroductoryOffers** | **int32** | maximum number of related introductoryOffers returned (when they are included) | 
 **limitOfferCodes** | **int32** | maximum number of related offerCodes returned (when they are included) | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 
 **limitPromotionalOffers** | **int32** | maximum number of related promotionalOffers returned (when they are included) | 
 **limitSubscriptionLocalizations** | **int32** | maximum number of related subscriptionLocalizations returned (when they are included) | 
 **limitWinBackOffers** | **int32** | maximum number of related winBackOffers returned (when they are included) | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsImagesGetToManyRelated

> SubscriptionImagesResponse SubscriptionsImagesGetToManyRelated(ctx, id).FieldsSubscriptionImages(fieldsSubscriptionImages).FieldsSubscriptions(fieldsSubscriptions).Limit(limit).Include(include).Execute()



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
	fieldsSubscriptionImages := []string{"FieldsSubscriptionImages_example"} // []string | the fields to include for returned resources of type subscriptionImages (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsImagesGetToManyRelated(context.Background(), id).FieldsSubscriptionImages(fieldsSubscriptionImages).FieldsSubscriptions(fieldsSubscriptions).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsImagesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsImagesGetToManyRelated`: SubscriptionImagesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsImagesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsImagesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionImages** | **[]string** | the fields to include for returned resources of type subscriptionImages | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionImagesResponse**](SubscriptionImagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIntroductoryOffersDeleteToManyRelationship

> SubscriptionsIntroductoryOffersDeleteToManyRelationship(ctx, id).SubscriptionIntroductoryOffersLinkagesRequest(subscriptionIntroductoryOffersLinkagesRequest).Execute()



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
	subscriptionIntroductoryOffersLinkagesRequest := *openapiclient.NewSubscriptionIntroductoryOffersLinkagesRequest([]openapiclient.SubscriptionRelationshipsIntroductoryOffersDataInner{*openapiclient.NewSubscriptionRelationshipsIntroductoryOffersDataInner("Type_example", "Id_example")}) // SubscriptionIntroductoryOffersLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionsAPI.SubscriptionsIntroductoryOffersDeleteToManyRelationship(context.Background(), id).SubscriptionIntroductoryOffersLinkagesRequest(subscriptionIntroductoryOffersLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIntroductoryOffersDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionsIntroductoryOffersDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionIntroductoryOffersLinkagesRequest** | [**SubscriptionIntroductoryOffersLinkagesRequest**](SubscriptionIntroductoryOffersLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIntroductoryOffersGetToManyRelated

> SubscriptionIntroductoryOffersResponse SubscriptionsIntroductoryOffersGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsSubscriptionIntroductoryOffers(fieldsSubscriptionIntroductoryOffers).FieldsSubscriptions(fieldsSubscriptions).FieldsTerritories(fieldsTerritories).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Limit(limit).Include(include).Execute()



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
	fieldsSubscriptionIntroductoryOffers := []string{"FieldsSubscriptionIntroductoryOffers_example"} // []string | the fields to include for returned resources of type subscriptionIntroductoryOffers (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	fieldsSubscriptionPricePoints := []string{"FieldsSubscriptionPricePoints_example"} // []string | the fields to include for returned resources of type subscriptionPricePoints (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIntroductoryOffersGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsSubscriptionIntroductoryOffers(fieldsSubscriptionIntroductoryOffers).FieldsSubscriptions(fieldsSubscriptions).FieldsTerritories(fieldsTerritories).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIntroductoryOffersGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIntroductoryOffersGetToManyRelated`: SubscriptionIntroductoryOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIntroductoryOffersGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIntroductoryOffersGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsSubscriptionIntroductoryOffers** | **[]string** | the fields to include for returned resources of type subscriptionIntroductoryOffers | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **fieldsSubscriptionPricePoints** | **[]string** | the fields to include for returned resources of type subscriptionPricePoints | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionIntroductoryOffersResponse**](SubscriptionIntroductoryOffersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIntroductoryOffersGetToManyRelationship

> SubscriptionIntroductoryOffersLinkagesResponse SubscriptionsIntroductoryOffersGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIntroductoryOffersGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIntroductoryOffersGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIntroductoryOffersGetToManyRelationship`: SubscriptionIntroductoryOffersLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIntroductoryOffersGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIntroductoryOffersGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**SubscriptionIntroductoryOffersLinkagesResponse**](SubscriptionIntroductoryOffersLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsOfferCodesGetToManyRelated

> SubscriptionOfferCodesResponse SubscriptionsOfferCodesGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionOfferCodeOneTimeUseCodes(fieldsSubscriptionOfferCodeOneTimeUseCodes).FieldsSubscriptionOfferCodeCustomCodes(fieldsSubscriptionOfferCodeCustomCodes).FieldsSubscriptionOfferCodePrices(fieldsSubscriptionOfferCodePrices).Limit(limit).Include(include).LimitOneTimeUseCodes(limitOneTimeUseCodes).LimitCustomCodes(limitCustomCodes).LimitPrices(limitPrices).Execute()



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
	filterTerritory := []string{"Inner_example"} // []string | filter by territory (optional)
	fieldsSubscriptionOfferCodes := []string{"FieldsSubscriptionOfferCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodes (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	fieldsSubscriptionOfferCodeOneTimeUseCodes := []string{"FieldsSubscriptionOfferCodeOneTimeUseCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodeOneTimeUseCodes (optional)
	fieldsSubscriptionOfferCodeCustomCodes := []string{"FieldsSubscriptionOfferCodeCustomCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodeCustomCodes (optional)
	fieldsSubscriptionOfferCodePrices := []string{"FieldsSubscriptionOfferCodePrices_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodePrices (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitOneTimeUseCodes := int32(56) // int32 | maximum number of related oneTimeUseCodes returned (when they are included) (optional)
	limitCustomCodes := int32(56) // int32 | maximum number of related customCodes returned (when they are included) (optional)
	limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsOfferCodesGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionOfferCodeOneTimeUseCodes(fieldsSubscriptionOfferCodeOneTimeUseCodes).FieldsSubscriptionOfferCodeCustomCodes(fieldsSubscriptionOfferCodeCustomCodes).FieldsSubscriptionOfferCodePrices(fieldsSubscriptionOfferCodePrices).Limit(limit).Include(include).LimitOneTimeUseCodes(limitOneTimeUseCodes).LimitCustomCodes(limitCustomCodes).LimitPrices(limitPrices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsOfferCodesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsOfferCodesGetToManyRelated`: SubscriptionOfferCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsOfferCodesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsOfferCodesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by territory | 
 **fieldsSubscriptionOfferCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodes | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **fieldsSubscriptionOfferCodeOneTimeUseCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodeOneTimeUseCodes | 
 **fieldsSubscriptionOfferCodeCustomCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodeCustomCodes | 
 **fieldsSubscriptionOfferCodePrices** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodePrices | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitOneTimeUseCodes** | **int32** | maximum number of related oneTimeUseCodes returned (when they are included) | 
 **limitCustomCodes** | **int32** | maximum number of related customCodes returned (when they are included) | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 

### Return type

[**SubscriptionOfferCodesResponse**](SubscriptionOfferCodesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPricePointsGetToManyRelated

> SubscriptionPricePointsResponse SubscriptionsPricePointsGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()



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
	fieldsSubscriptionPricePoints := []string{"FieldsSubscriptionPricePoints_example"} // []string | the fields to include for returned resources of type subscriptionPricePoints (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsPricePointsGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPricePointsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsPricePointsGetToManyRelated`: SubscriptionPricePointsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsPricePointsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPricePointsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsSubscriptionPricePoints** | **[]string** | the fields to include for returned resources of type subscriptionPricePoints | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionPricePointsResponse**](SubscriptionPricePointsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPricesDeleteToManyRelationship

> SubscriptionsPricesDeleteToManyRelationship(ctx, id).SubscriptionPricesLinkagesRequest(subscriptionPricesLinkagesRequest).Execute()



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
	subscriptionPricesLinkagesRequest := *openapiclient.NewSubscriptionPricesLinkagesRequest([]openapiclient.SubscriptionRelationshipsPricesDataInner{*openapiclient.NewSubscriptionRelationshipsPricesDataInner("Type_example", "Id_example")}) // SubscriptionPricesLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionsAPI.SubscriptionsPricesDeleteToManyRelationship(context.Background(), id).SubscriptionPricesLinkagesRequest(subscriptionPricesLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPricesDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionsPricesDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionPricesLinkagesRequest** | [**SubscriptionPricesLinkagesRequest**](SubscriptionPricesLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPricesGetToManyRelated

> SubscriptionPricesResponse SubscriptionsPricesGetToManyRelated(ctx, id).FilterSubscriptionPricePoint(filterSubscriptionPricePoint).FilterTerritory(filterTerritory).FieldsSubscriptionPrices(fieldsSubscriptionPrices).FieldsTerritories(fieldsTerritories).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Limit(limit).Include(include).Execute()



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
	filterSubscriptionPricePoint := []string{"Inner_example"} // []string | filter by id(s) of related 'subscriptionPricePoint' (optional)
	filterTerritory := []string{"Inner_example"} // []string | filter by id(s) of related 'territory' (optional)
	fieldsSubscriptionPrices := []string{"FieldsSubscriptionPrices_example"} // []string | the fields to include for returned resources of type subscriptionPrices (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	fieldsSubscriptionPricePoints := []string{"FieldsSubscriptionPricePoints_example"} // []string | the fields to include for returned resources of type subscriptionPricePoints (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsPricesGetToManyRelated(context.Background(), id).FilterSubscriptionPricePoint(filterSubscriptionPricePoint).FilterTerritory(filterTerritory).FieldsSubscriptionPrices(fieldsSubscriptionPrices).FieldsTerritories(fieldsTerritories).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPricesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsPricesGetToManyRelated`: SubscriptionPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsPricesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPricesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterSubscriptionPricePoint** | **[]string** | filter by id(s) of related &#39;subscriptionPricePoint&#39; | 
 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsSubscriptionPrices** | **[]string** | the fields to include for returned resources of type subscriptionPrices | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **fieldsSubscriptionPricePoints** | **[]string** | the fields to include for returned resources of type subscriptionPricePoints | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionPricesResponse**](SubscriptionPricesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPricesGetToManyRelationship

> SubscriptionPricesLinkagesResponse SubscriptionsPricesGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsPricesGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPricesGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsPricesGetToManyRelationship`: SubscriptionPricesLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsPricesGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPricesGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**SubscriptionPricesLinkagesResponse**](SubscriptionPricesLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPromotedPurchaseGetToOneRelated

> PromotedPurchaseResponse SubscriptionsPromotedPurchaseGetToOneRelated(ctx, id).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsInAppPurchases(fieldsInAppPurchases).FieldsSubscriptions(fieldsSubscriptions).FieldsPromotedPurchaseImages(fieldsPromotedPurchaseImages).Include(include).LimitPromotionImages(limitPromotionImages).Execute()



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
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsPromotedPurchaseGetToOneRelated(context.Background(), id).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsInAppPurchases(fieldsInAppPurchases).FieldsSubscriptions(fieldsSubscriptions).FieldsPromotedPurchaseImages(fieldsPromotedPurchaseImages).Include(include).LimitPromotionImages(limitPromotionImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPromotedPurchaseGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsPromotedPurchaseGetToOneRelated`: PromotedPurchaseResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsPromotedPurchaseGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPromotedPurchaseGetToOneRelatedRequest struct via the builder pattern


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


## SubscriptionsPromotionalOffersGetToManyRelated

> SubscriptionPromotionalOffersResponse SubscriptionsPromotionalOffersGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsSubscriptionPromotionalOffers(fieldsSubscriptionPromotionalOffers).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionPromotionalOfferPrices(fieldsSubscriptionPromotionalOfferPrices).Limit(limit).Include(include).LimitPrices(limitPrices).Execute()



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
	filterTerritory := []string{"Inner_example"} // []string | filter by territory (optional)
	fieldsSubscriptionPromotionalOffers := []string{"FieldsSubscriptionPromotionalOffers_example"} // []string | the fields to include for returned resources of type subscriptionPromotionalOffers (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	fieldsSubscriptionPromotionalOfferPrices := []string{"FieldsSubscriptionPromotionalOfferPrices_example"} // []string | the fields to include for returned resources of type subscriptionPromotionalOfferPrices (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsPromotionalOffersGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsSubscriptionPromotionalOffers(fieldsSubscriptionPromotionalOffers).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionPromotionalOfferPrices(fieldsSubscriptionPromotionalOfferPrices).Limit(limit).Include(include).LimitPrices(limitPrices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPromotionalOffersGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsPromotionalOffersGetToManyRelated`: SubscriptionPromotionalOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsPromotionalOffersGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPromotionalOffersGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by territory | 
 **fieldsSubscriptionPromotionalOffers** | **[]string** | the fields to include for returned resources of type subscriptionPromotionalOffers | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **fieldsSubscriptionPromotionalOfferPrices** | **[]string** | the fields to include for returned resources of type subscriptionPromotionalOfferPrices | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 

### Return type

[**SubscriptionPromotionalOffersResponse**](SubscriptionPromotionalOffersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionAvailabilityGetToOneRelated

> SubscriptionAvailabilityResponse SubscriptionsSubscriptionAvailabilityGetToOneRelated(ctx, id).FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities).FieldsSubscriptions(fieldsSubscriptions).FieldsTerritories(fieldsTerritories).Include(include).LimitAvailableTerritories(limitAvailableTerritories).Execute()



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
	fieldsSubscriptionAvailabilities := []string{"FieldsSubscriptionAvailabilities_example"} // []string | the fields to include for returned resources of type subscriptionAvailabilities (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAvailableTerritories := int32(56) // int32 | maximum number of related availableTerritories returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsSubscriptionAvailabilityGetToOneRelated(context.Background(), id).FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities).FieldsSubscriptions(fieldsSubscriptions).FieldsTerritories(fieldsTerritories).Include(include).LimitAvailableTerritories(limitAvailableTerritories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsSubscriptionAvailabilityGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsSubscriptionAvailabilityGetToOneRelated`: SubscriptionAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsSubscriptionAvailabilityGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionAvailabilityGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionAvailabilities** | **[]string** | the fields to include for returned resources of type subscriptionAvailabilities | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
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


## SubscriptionsSubscriptionLocalizationsGetToManyRelated

> SubscriptionLocalizationsResponse SubscriptionsSubscriptionLocalizationsGetToManyRelated(ctx, id).FieldsSubscriptionLocalizations(fieldsSubscriptionLocalizations).FieldsSubscriptions(fieldsSubscriptions).Limit(limit).Include(include).Execute()



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
	fieldsSubscriptionLocalizations := []string{"FieldsSubscriptionLocalizations_example"} // []string | the fields to include for returned resources of type subscriptionLocalizations (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsSubscriptionLocalizationsGetToManyRelated(context.Background(), id).FieldsSubscriptionLocalizations(fieldsSubscriptionLocalizations).FieldsSubscriptions(fieldsSubscriptions).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsSubscriptionLocalizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsSubscriptionLocalizationsGetToManyRelated`: SubscriptionLocalizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsSubscriptionLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionLocalizations** | **[]string** | the fields to include for returned resources of type subscriptionLocalizations | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionLocalizationsResponse**](SubscriptionLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdateInstance

> SubscriptionResponse SubscriptionsUpdateInstance(ctx, id).SubscriptionUpdateRequest(subscriptionUpdateRequest).Execute()



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
	subscriptionUpdateRequest := *openapiclient.NewSubscriptionUpdateRequest(*openapiclient.NewSubscriptionUpdateRequestData("Type_example", "Id_example")) // SubscriptionUpdateRequest | Subscription representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsUpdateInstance(context.Background(), id).SubscriptionUpdateRequest(subscriptionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsUpdateInstance`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionUpdateRequest** | [**SubscriptionUpdateRequest**](SubscriptionUpdateRequest.md) | Subscription representation | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsWinBackOffersGetToManyRelated

> WinBackOffersResponse SubscriptionsWinBackOffersGetToManyRelated(ctx, id).FieldsWinBackOffers(fieldsWinBackOffers).FieldsWinBackOfferPrices(fieldsWinBackOfferPrices).Limit(limit).Include(include).LimitPrices(limitPrices).Execute()



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
	fieldsWinBackOffers := []string{"FieldsWinBackOffers_example"} // []string | the fields to include for returned resources of type winBackOffers (optional)
	fieldsWinBackOfferPrices := []string{"FieldsWinBackOfferPrices_example"} // []string | the fields to include for returned resources of type winBackOfferPrices (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsWinBackOffersGetToManyRelated(context.Background(), id).FieldsWinBackOffers(fieldsWinBackOffers).FieldsWinBackOfferPrices(fieldsWinBackOfferPrices).Limit(limit).Include(include).LimitPrices(limitPrices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsWinBackOffersGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsWinBackOffersGetToManyRelated`: WinBackOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsWinBackOffersGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsWinBackOffersGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsWinBackOffers** | **[]string** | the fields to include for returned resources of type winBackOffers | 
 **fieldsWinBackOfferPrices** | **[]string** | the fields to include for returned resources of type winBackOfferPrices | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 

### Return type

[**WinBackOffersResponse**](WinBackOffersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

