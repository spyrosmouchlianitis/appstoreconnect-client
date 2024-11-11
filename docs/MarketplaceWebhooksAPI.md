# \MarketplaceWebhooksAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketplaceWebhooksCreateInstance**](MarketplaceWebhooksAPI.md#MarketplaceWebhooksCreateInstance) | **Post** /v1/marketplaceWebhooks | 
[**MarketplaceWebhooksDeleteInstance**](MarketplaceWebhooksAPI.md#MarketplaceWebhooksDeleteInstance) | **Delete** /v1/marketplaceWebhooks/{id} | 
[**MarketplaceWebhooksGetCollection**](MarketplaceWebhooksAPI.md#MarketplaceWebhooksGetCollection) | **Get** /v1/marketplaceWebhooks | 
[**MarketplaceWebhooksUpdateInstance**](MarketplaceWebhooksAPI.md#MarketplaceWebhooksUpdateInstance) | **Patch** /v1/marketplaceWebhooks/{id} | 



## MarketplaceWebhooksCreateInstance

> MarketplaceWebhookResponse MarketplaceWebhooksCreateInstance(ctx).MarketplaceWebhookCreateRequest(marketplaceWebhookCreateRequest).Execute()



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
	marketplaceWebhookCreateRequest := *openapiclient.NewMarketplaceWebhookCreateRequest(*openapiclient.NewMarketplaceWebhookCreateRequestData("Type_example", *openapiclient.NewMarketplaceWebhookCreateRequestDataAttributes("EndpointUrl_example", "Secret_example"))) // MarketplaceWebhookCreateRequest | MarketplaceWebhook representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceWebhooksAPI.MarketplaceWebhooksCreateInstance(context.Background()).MarketplaceWebhookCreateRequest(marketplaceWebhookCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceWebhooksAPI.MarketplaceWebhooksCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceWebhooksCreateInstance`: MarketplaceWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceWebhooksAPI.MarketplaceWebhooksCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceWebhooksCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceWebhookCreateRequest** | [**MarketplaceWebhookCreateRequest**](MarketplaceWebhookCreateRequest.md) | MarketplaceWebhook representation | 

### Return type

[**MarketplaceWebhookResponse**](MarketplaceWebhookResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceWebhooksDeleteInstance

> MarketplaceWebhooksDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.MarketplaceWebhooksAPI.MarketplaceWebhooksDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceWebhooksAPI.MarketplaceWebhooksDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMarketplaceWebhooksDeleteInstanceRequest struct via the builder pattern


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


## MarketplaceWebhooksGetCollection

> MarketplaceWebhooksResponse MarketplaceWebhooksGetCollection(ctx).FieldsMarketplaceWebhooks(fieldsMarketplaceWebhooks).Limit(limit).Execute()



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
	fieldsMarketplaceWebhooks := []string{"FieldsMarketplaceWebhooks_example"} // []string | the fields to include for returned resources of type marketplaceWebhooks (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceWebhooksAPI.MarketplaceWebhooksGetCollection(context.Background()).FieldsMarketplaceWebhooks(fieldsMarketplaceWebhooks).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceWebhooksAPI.MarketplaceWebhooksGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceWebhooksGetCollection`: MarketplaceWebhooksResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceWebhooksAPI.MarketplaceWebhooksGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceWebhooksGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldsMarketplaceWebhooks** | **[]string** | the fields to include for returned resources of type marketplaceWebhooks | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**MarketplaceWebhooksResponse**](MarketplaceWebhooksResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceWebhooksUpdateInstance

> MarketplaceWebhookResponse MarketplaceWebhooksUpdateInstance(ctx, id).MarketplaceWebhookUpdateRequest(marketplaceWebhookUpdateRequest).Execute()



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
	marketplaceWebhookUpdateRequest := *openapiclient.NewMarketplaceWebhookUpdateRequest(*openapiclient.NewMarketplaceWebhookUpdateRequestData("Type_example", "Id_example")) // MarketplaceWebhookUpdateRequest | MarketplaceWebhook representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceWebhooksAPI.MarketplaceWebhooksUpdateInstance(context.Background(), id).MarketplaceWebhookUpdateRequest(marketplaceWebhookUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceWebhooksAPI.MarketplaceWebhooksUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceWebhooksUpdateInstance`: MarketplaceWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceWebhooksAPI.MarketplaceWebhooksUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceWebhooksUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceWebhookUpdateRequest** | [**MarketplaceWebhookUpdateRequest**](MarketplaceWebhookUpdateRequest.md) | MarketplaceWebhook representation | 

### Return type

[**MarketplaceWebhookResponse**](MarketplaceWebhookResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

