# \SubscriptionPricePointsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionPricePointsEqualizationsGetToManyRelated**](SubscriptionPricePointsAPI.md#SubscriptionPricePointsEqualizationsGetToManyRelated) | **Get** /v1/subscriptionPricePoints/{id}/equalizations | 
[**SubscriptionPricePointsGetInstance**](SubscriptionPricePointsAPI.md#SubscriptionPricePointsGetInstance) | **Get** /v1/subscriptionPricePoints/{id} | 



## SubscriptionPricePointsEqualizationsGetToManyRelated

> SubscriptionPricePointsResponse SubscriptionPricePointsEqualizationsGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FilterSubscription(filterSubscription).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()



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
	filterSubscription := []string{"Inner_example"} // []string | filter by id(s) of related 'subscription' (optional)
	fieldsSubscriptionPricePoints := []string{"FieldsSubscriptionPricePoints_example"} // []string | the fields to include for returned resources of type subscriptionPricePoints (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionPricePointsAPI.SubscriptionPricePointsEqualizationsGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FilterSubscription(filterSubscription).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPricePointsAPI.SubscriptionPricePointsEqualizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPricePointsEqualizationsGetToManyRelated`: SubscriptionPricePointsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPricePointsAPI.SubscriptionPricePointsEqualizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPricePointsEqualizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **filterSubscription** | **[]string** | filter by id(s) of related &#39;subscription&#39; | 
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


## SubscriptionPricePointsGetInstance

> SubscriptionPricePointResponse SubscriptionPricePointsGetInstance(ctx, id).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Include(include).Execute()



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
	fieldsSubscriptionPricePoints := []string{"FieldsSubscriptionPricePoints_example"} // []string | the fields to include for returned resources of type subscriptionPricePoints (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionPricePointsAPI.SubscriptionPricePointsGetInstance(context.Background(), id).FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPricePointsAPI.SubscriptionPricePointsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPricePointsGetInstance`: SubscriptionPricePointResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPricePointsAPI.SubscriptionPricePointsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPricePointsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionPricePoints** | **[]string** | the fields to include for returned resources of type subscriptionPricePoints | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionPricePointResponse**](SubscriptionPricePointResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

