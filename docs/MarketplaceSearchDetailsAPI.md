# \MarketplaceSearchDetailsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketplaceSearchDetailsCreateInstance**](MarketplaceSearchDetailsAPI.md#MarketplaceSearchDetailsCreateInstance) | **Post** /v1/marketplaceSearchDetails | 
[**MarketplaceSearchDetailsDeleteInstance**](MarketplaceSearchDetailsAPI.md#MarketplaceSearchDetailsDeleteInstance) | **Delete** /v1/marketplaceSearchDetails/{id} | 
[**MarketplaceSearchDetailsUpdateInstance**](MarketplaceSearchDetailsAPI.md#MarketplaceSearchDetailsUpdateInstance) | **Patch** /v1/marketplaceSearchDetails/{id} | 



## MarketplaceSearchDetailsCreateInstance

> MarketplaceSearchDetailResponse MarketplaceSearchDetailsCreateInstance(ctx).MarketplaceSearchDetailCreateRequest(marketplaceSearchDetailCreateRequest).Execute()



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
	marketplaceSearchDetailCreateRequest := *openapiclient.NewMarketplaceSearchDetailCreateRequest(*openapiclient.NewMarketplaceSearchDetailCreateRequestData("Type_example", *openapiclient.NewMarketplaceSearchDetailCreateRequestDataAttributes("CatalogUrl_example"), *openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // MarketplaceSearchDetailCreateRequest | MarketplaceSearchDetail representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceSearchDetailsAPI.MarketplaceSearchDetailsCreateInstance(context.Background()).MarketplaceSearchDetailCreateRequest(marketplaceSearchDetailCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceSearchDetailsAPI.MarketplaceSearchDetailsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceSearchDetailsCreateInstance`: MarketplaceSearchDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceSearchDetailsAPI.MarketplaceSearchDetailsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceSearchDetailsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceSearchDetailCreateRequest** | [**MarketplaceSearchDetailCreateRequest**](MarketplaceSearchDetailCreateRequest.md) | MarketplaceSearchDetail representation | 

### Return type

[**MarketplaceSearchDetailResponse**](MarketplaceSearchDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceSearchDetailsDeleteInstance

> MarketplaceSearchDetailsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.MarketplaceSearchDetailsAPI.MarketplaceSearchDetailsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceSearchDetailsAPI.MarketplaceSearchDetailsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMarketplaceSearchDetailsDeleteInstanceRequest struct via the builder pattern


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


## MarketplaceSearchDetailsUpdateInstance

> MarketplaceSearchDetailResponse MarketplaceSearchDetailsUpdateInstance(ctx, id).MarketplaceSearchDetailUpdateRequest(marketplaceSearchDetailUpdateRequest).Execute()



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
	marketplaceSearchDetailUpdateRequest := *openapiclient.NewMarketplaceSearchDetailUpdateRequest(*openapiclient.NewMarketplaceSearchDetailUpdateRequestData("Type_example", "Id_example")) // MarketplaceSearchDetailUpdateRequest | MarketplaceSearchDetail representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceSearchDetailsAPI.MarketplaceSearchDetailsUpdateInstance(context.Background(), id).MarketplaceSearchDetailUpdateRequest(marketplaceSearchDetailUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceSearchDetailsAPI.MarketplaceSearchDetailsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceSearchDetailsUpdateInstance`: MarketplaceSearchDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceSearchDetailsAPI.MarketplaceSearchDetailsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceSearchDetailsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceSearchDetailUpdateRequest** | [**MarketplaceSearchDetailUpdateRequest**](MarketplaceSearchDetailUpdateRequest.md) | MarketplaceSearchDetail representation | 

### Return type

[**MarketplaceSearchDetailResponse**](MarketplaceSearchDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

