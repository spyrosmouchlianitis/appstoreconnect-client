# \SubscriptionGroupsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionGroupsCreateInstance**](SubscriptionGroupsAPI.md#SubscriptionGroupsCreateInstance) | **Post** /v1/subscriptionGroups | 
[**SubscriptionGroupsDeleteInstance**](SubscriptionGroupsAPI.md#SubscriptionGroupsDeleteInstance) | **Delete** /v1/subscriptionGroups/{id} | 
[**SubscriptionGroupsGetInstance**](SubscriptionGroupsAPI.md#SubscriptionGroupsGetInstance) | **Get** /v1/subscriptionGroups/{id} | 
[**SubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelated**](SubscriptionGroupsAPI.md#SubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelated) | **Get** /v1/subscriptionGroups/{id}/subscriptionGroupLocalizations | 
[**SubscriptionGroupsSubscriptionsGetToManyRelated**](SubscriptionGroupsAPI.md#SubscriptionGroupsSubscriptionsGetToManyRelated) | **Get** /v1/subscriptionGroups/{id}/subscriptions | 
[**SubscriptionGroupsUpdateInstance**](SubscriptionGroupsAPI.md#SubscriptionGroupsUpdateInstance) | **Patch** /v1/subscriptionGroups/{id} | 



## SubscriptionGroupsCreateInstance

> SubscriptionGroupResponse SubscriptionGroupsCreateInstance(ctx).SubscriptionGroupCreateRequest(subscriptionGroupCreateRequest).Execute()



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
	subscriptionGroupCreateRequest := *openapiclient.NewSubscriptionGroupCreateRequest(*openapiclient.NewSubscriptionGroupCreateRequestData("Type_example", *openapiclient.NewSubscriptionGroupCreateRequestDataAttributes("ReferenceName_example"), *openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // SubscriptionGroupCreateRequest | SubscriptionGroup representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionGroupsAPI.SubscriptionGroupsCreateInstance(context.Background()).SubscriptionGroupCreateRequest(subscriptionGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsAPI.SubscriptionGroupsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionGroupsCreateInstance`: SubscriptionGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsAPI.SubscriptionGroupsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionGroupCreateRequest** | [**SubscriptionGroupCreateRequest**](SubscriptionGroupCreateRequest.md) | SubscriptionGroup representation | 

### Return type

[**SubscriptionGroupResponse**](SubscriptionGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionGroupsDeleteInstance

> SubscriptionGroupsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.SubscriptionGroupsAPI.SubscriptionGroupsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsAPI.SubscriptionGroupsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionGroupsDeleteInstanceRequest struct via the builder pattern


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


## SubscriptionGroupsGetInstance

> SubscriptionGroupResponse SubscriptionGroupsGetInstance(ctx, id).FieldsSubscriptionGroups(fieldsSubscriptionGroups).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionGroupLocalizations(fieldsSubscriptionGroupLocalizations).Include(include).LimitSubscriptionGroupLocalizations(limitSubscriptionGroupLocalizations).LimitSubscriptions(limitSubscriptions).Execute()



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
	fieldsSubscriptionGroups := []string{"FieldsSubscriptionGroups_example"} // []string | the fields to include for returned resources of type subscriptionGroups (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	fieldsSubscriptionGroupLocalizations := []string{"FieldsSubscriptionGroupLocalizations_example"} // []string | the fields to include for returned resources of type subscriptionGroupLocalizations (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitSubscriptionGroupLocalizations := int32(56) // int32 | maximum number of related subscriptionGroupLocalizations returned (when they are included) (optional)
	limitSubscriptions := int32(56) // int32 | maximum number of related subscriptions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionGroupsAPI.SubscriptionGroupsGetInstance(context.Background(), id).FieldsSubscriptionGroups(fieldsSubscriptionGroups).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionGroupLocalizations(fieldsSubscriptionGroupLocalizations).Include(include).LimitSubscriptionGroupLocalizations(limitSubscriptionGroupLocalizations).LimitSubscriptions(limitSubscriptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsAPI.SubscriptionGroupsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionGroupsGetInstance`: SubscriptionGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsAPI.SubscriptionGroupsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionGroups** | **[]string** | the fields to include for returned resources of type subscriptionGroups | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **fieldsSubscriptionGroupLocalizations** | **[]string** | the fields to include for returned resources of type subscriptionGroupLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitSubscriptionGroupLocalizations** | **int32** | maximum number of related subscriptionGroupLocalizations returned (when they are included) | 
 **limitSubscriptions** | **int32** | maximum number of related subscriptions returned (when they are included) | 

### Return type

[**SubscriptionGroupResponse**](SubscriptionGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelated

> SubscriptionGroupLocalizationsResponse SubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelated(ctx, id).FieldsSubscriptionGroupLocalizations(fieldsSubscriptionGroupLocalizations).FieldsSubscriptionGroups(fieldsSubscriptionGroups).Limit(limit).Include(include).Execute()



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
	fieldsSubscriptionGroupLocalizations := []string{"FieldsSubscriptionGroupLocalizations_example"} // []string | the fields to include for returned resources of type subscriptionGroupLocalizations (optional)
	fieldsSubscriptionGroups := []string{"FieldsSubscriptionGroups_example"} // []string | the fields to include for returned resources of type subscriptionGroups (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionGroupsAPI.SubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelated(context.Background(), id).FieldsSubscriptionGroupLocalizations(fieldsSubscriptionGroupLocalizations).FieldsSubscriptionGroups(fieldsSubscriptionGroups).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsAPI.SubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelated`: SubscriptionGroupLocalizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsAPI.SubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupsSubscriptionGroupLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionGroupLocalizations** | **[]string** | the fields to include for returned resources of type subscriptionGroupLocalizations | 
 **fieldsSubscriptionGroups** | **[]string** | the fields to include for returned resources of type subscriptionGroups | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionGroupLocalizationsResponse**](SubscriptionGroupLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionGroupsSubscriptionsGetToManyRelated

> SubscriptionsResponse SubscriptionGroupsSubscriptionsGetToManyRelated(ctx, id).FilterProductId(filterProductId).FilterName(filterName).FilterState(filterState).Sort(sort).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionLocalizations(fieldsSubscriptionLocalizations).FieldsSubscriptionAppStoreReviewScreenshots(fieldsSubscriptionAppStoreReviewScreenshots).FieldsSubscriptionGroups(fieldsSubscriptionGroups).FieldsSubscriptionIntroductoryOffers(fieldsSubscriptionIntroductoryOffers).FieldsSubscriptionPromotionalOffers(fieldsSubscriptionPromotionalOffers).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).FieldsSubscriptionPrices(fieldsSubscriptionPrices).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities).FieldsWinBackOffers(fieldsWinBackOffers).FieldsSubscriptionImages(fieldsSubscriptionImages).Limit(limit).Include(include).LimitSubscriptionLocalizations(limitSubscriptionLocalizations).LimitIntroductoryOffers(limitIntroductoryOffers).LimitPromotionalOffers(limitPromotionalOffers).LimitOfferCodes(limitOfferCodes).LimitPrices(limitPrices).LimitWinBackOffers(limitWinBackOffers).LimitImages(limitImages).Execute()



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
	filterProductId := []string{"Inner_example"} // []string | filter by attribute 'productId' (optional)
	filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
	filterState := []string{"FilterState_example"} // []string | filter by attribute 'state' (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
	fieldsSubscriptions := []string{"FieldsSubscriptions_example"} // []string | the fields to include for returned resources of type subscriptions (optional)
	fieldsSubscriptionLocalizations := []string{"FieldsSubscriptionLocalizations_example"} // []string | the fields to include for returned resources of type subscriptionLocalizations (optional)
	fieldsSubscriptionAppStoreReviewScreenshots := []string{"FieldsSubscriptionAppStoreReviewScreenshots_example"} // []string | the fields to include for returned resources of type subscriptionAppStoreReviewScreenshots (optional)
	fieldsSubscriptionGroups := []string{"FieldsSubscriptionGroups_example"} // []string | the fields to include for returned resources of type subscriptionGroups (optional)
	fieldsSubscriptionIntroductoryOffers := []string{"FieldsSubscriptionIntroductoryOffers_example"} // []string | the fields to include for returned resources of type subscriptionIntroductoryOffers (optional)
	fieldsSubscriptionPromotionalOffers := []string{"FieldsSubscriptionPromotionalOffers_example"} // []string | the fields to include for returned resources of type subscriptionPromotionalOffers (optional)
	fieldsSubscriptionOfferCodes := []string{"FieldsSubscriptionOfferCodes_example"} // []string | the fields to include for returned resources of type subscriptionOfferCodes (optional)
	fieldsSubscriptionPrices := []string{"FieldsSubscriptionPrices_example"} // []string | the fields to include for returned resources of type subscriptionPrices (optional)
	fieldsPromotedPurchases := []string{"FieldsPromotedPurchases_example"} // []string | the fields to include for returned resources of type promotedPurchases (optional)
	fieldsSubscriptionAvailabilities := []string{"FieldsSubscriptionAvailabilities_example"} // []string | the fields to include for returned resources of type subscriptionAvailabilities (optional)
	fieldsWinBackOffers := []string{"FieldsWinBackOffers_example"} // []string | the fields to include for returned resources of type winBackOffers (optional)
	fieldsSubscriptionImages := []string{"FieldsSubscriptionImages_example"} // []string | the fields to include for returned resources of type subscriptionImages (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitSubscriptionLocalizations := int32(56) // int32 | maximum number of related subscriptionLocalizations returned (when they are included) (optional)
	limitIntroductoryOffers := int32(56) // int32 | maximum number of related introductoryOffers returned (when they are included) (optional)
	limitPromotionalOffers := int32(56) // int32 | maximum number of related promotionalOffers returned (when they are included) (optional)
	limitOfferCodes := int32(56) // int32 | maximum number of related offerCodes returned (when they are included) (optional)
	limitPrices := int32(56) // int32 | maximum number of related prices returned (when they are included) (optional)
	limitWinBackOffers := int32(56) // int32 | maximum number of related winBackOffers returned (when they are included) (optional)
	limitImages := int32(56) // int32 | maximum number of related images returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionGroupsAPI.SubscriptionGroupsSubscriptionsGetToManyRelated(context.Background(), id).FilterProductId(filterProductId).FilterName(filterName).FilterState(filterState).Sort(sort).FieldsSubscriptions(fieldsSubscriptions).FieldsSubscriptionLocalizations(fieldsSubscriptionLocalizations).FieldsSubscriptionAppStoreReviewScreenshots(fieldsSubscriptionAppStoreReviewScreenshots).FieldsSubscriptionGroups(fieldsSubscriptionGroups).FieldsSubscriptionIntroductoryOffers(fieldsSubscriptionIntroductoryOffers).FieldsSubscriptionPromotionalOffers(fieldsSubscriptionPromotionalOffers).FieldsSubscriptionOfferCodes(fieldsSubscriptionOfferCodes).FieldsSubscriptionPrices(fieldsSubscriptionPrices).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities).FieldsWinBackOffers(fieldsWinBackOffers).FieldsSubscriptionImages(fieldsSubscriptionImages).Limit(limit).Include(include).LimitSubscriptionLocalizations(limitSubscriptionLocalizations).LimitIntroductoryOffers(limitIntroductoryOffers).LimitPromotionalOffers(limitPromotionalOffers).LimitOfferCodes(limitOfferCodes).LimitPrices(limitPrices).LimitWinBackOffers(limitWinBackOffers).LimitImages(limitImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsAPI.SubscriptionGroupsSubscriptionsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionGroupsSubscriptionsGetToManyRelated`: SubscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsAPI.SubscriptionGroupsSubscriptionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupsSubscriptionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterProductId** | **[]string** | filter by attribute &#39;productId&#39; | 
 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterState** | **[]string** | filter by attribute &#39;state&#39; | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsSubscriptions** | **[]string** | the fields to include for returned resources of type subscriptions | 
 **fieldsSubscriptionLocalizations** | **[]string** | the fields to include for returned resources of type subscriptionLocalizations | 
 **fieldsSubscriptionAppStoreReviewScreenshots** | **[]string** | the fields to include for returned resources of type subscriptionAppStoreReviewScreenshots | 
 **fieldsSubscriptionGroups** | **[]string** | the fields to include for returned resources of type subscriptionGroups | 
 **fieldsSubscriptionIntroductoryOffers** | **[]string** | the fields to include for returned resources of type subscriptionIntroductoryOffers | 
 **fieldsSubscriptionPromotionalOffers** | **[]string** | the fields to include for returned resources of type subscriptionPromotionalOffers | 
 **fieldsSubscriptionOfferCodes** | **[]string** | the fields to include for returned resources of type subscriptionOfferCodes | 
 **fieldsSubscriptionPrices** | **[]string** | the fields to include for returned resources of type subscriptionPrices | 
 **fieldsPromotedPurchases** | **[]string** | the fields to include for returned resources of type promotedPurchases | 
 **fieldsSubscriptionAvailabilities** | **[]string** | the fields to include for returned resources of type subscriptionAvailabilities | 
 **fieldsWinBackOffers** | **[]string** | the fields to include for returned resources of type winBackOffers | 
 **fieldsSubscriptionImages** | **[]string** | the fields to include for returned resources of type subscriptionImages | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitSubscriptionLocalizations** | **int32** | maximum number of related subscriptionLocalizations returned (when they are included) | 
 **limitIntroductoryOffers** | **int32** | maximum number of related introductoryOffers returned (when they are included) | 
 **limitPromotionalOffers** | **int32** | maximum number of related promotionalOffers returned (when they are included) | 
 **limitOfferCodes** | **int32** | maximum number of related offerCodes returned (when they are included) | 
 **limitPrices** | **int32** | maximum number of related prices returned (when they are included) | 
 **limitWinBackOffers** | **int32** | maximum number of related winBackOffers returned (when they are included) | 
 **limitImages** | **int32** | maximum number of related images returned (when they are included) | 

### Return type

[**SubscriptionsResponse**](SubscriptionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionGroupsUpdateInstance

> SubscriptionGroupResponse SubscriptionGroupsUpdateInstance(ctx, id).SubscriptionGroupUpdateRequest(subscriptionGroupUpdateRequest).Execute()



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
	subscriptionGroupUpdateRequest := *openapiclient.NewSubscriptionGroupUpdateRequest(*openapiclient.NewSubscriptionGroupUpdateRequestData("Type_example", "Id_example")) // SubscriptionGroupUpdateRequest | SubscriptionGroup representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionGroupsAPI.SubscriptionGroupsUpdateInstance(context.Background(), id).SubscriptionGroupUpdateRequest(subscriptionGroupUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsAPI.SubscriptionGroupsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionGroupsUpdateInstance`: SubscriptionGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsAPI.SubscriptionGroupsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionGroupUpdateRequest** | [**SubscriptionGroupUpdateRequest**](SubscriptionGroupUpdateRequest.md) | SubscriptionGroup representation | 

### Return type

[**SubscriptionGroupResponse**](SubscriptionGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

