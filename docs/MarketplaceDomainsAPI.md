# \MarketplaceDomainsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketplaceDomainsCreateInstance**](MarketplaceDomainsAPI.md#MarketplaceDomainsCreateInstance) | **Post** /v1/marketplaceDomains | 
[**MarketplaceDomainsDeleteInstance**](MarketplaceDomainsAPI.md#MarketplaceDomainsDeleteInstance) | **Delete** /v1/marketplaceDomains/{id} | 
[**MarketplaceDomainsGetCollection**](MarketplaceDomainsAPI.md#MarketplaceDomainsGetCollection) | **Get** /v1/marketplaceDomains | 
[**MarketplaceDomainsGetInstance**](MarketplaceDomainsAPI.md#MarketplaceDomainsGetInstance) | **Get** /v1/marketplaceDomains/{id} | 



## MarketplaceDomainsCreateInstance

> MarketplaceDomainResponse MarketplaceDomainsCreateInstance(ctx).MarketplaceDomainCreateRequest(marketplaceDomainCreateRequest).Execute()



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
	marketplaceDomainCreateRequest := *openapiclient.NewMarketplaceDomainCreateRequest(*openapiclient.NewMarketplaceDomainCreateRequestData("Type_example", *openapiclient.NewAlternativeDistributionDomainCreateRequestDataAttributes("Domain_example", "ReferenceName_example"))) // MarketplaceDomainCreateRequest | MarketplaceDomain representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceDomainsAPI.MarketplaceDomainsCreateInstance(context.Background()).MarketplaceDomainCreateRequest(marketplaceDomainCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceDomainsAPI.MarketplaceDomainsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceDomainsCreateInstance`: MarketplaceDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceDomainsAPI.MarketplaceDomainsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceDomainsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceDomainCreateRequest** | [**MarketplaceDomainCreateRequest**](MarketplaceDomainCreateRequest.md) | MarketplaceDomain representation | 

### Return type

[**MarketplaceDomainResponse**](MarketplaceDomainResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceDomainsDeleteInstance

> MarketplaceDomainsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.MarketplaceDomainsAPI.MarketplaceDomainsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceDomainsAPI.MarketplaceDomainsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMarketplaceDomainsDeleteInstanceRequest struct via the builder pattern


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


## MarketplaceDomainsGetCollection

> MarketplaceDomainsResponse MarketplaceDomainsGetCollection(ctx).FieldsMarketplaceDomains(fieldsMarketplaceDomains).Limit(limit).Execute()



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
	fieldsMarketplaceDomains := []string{"FieldsMarketplaceDomains_example"} // []string | the fields to include for returned resources of type marketplaceDomains (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceDomainsAPI.MarketplaceDomainsGetCollection(context.Background()).FieldsMarketplaceDomains(fieldsMarketplaceDomains).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceDomainsAPI.MarketplaceDomainsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceDomainsGetCollection`: MarketplaceDomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceDomainsAPI.MarketplaceDomainsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceDomainsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldsMarketplaceDomains** | **[]string** | the fields to include for returned resources of type marketplaceDomains | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**MarketplaceDomainsResponse**](MarketplaceDomainsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceDomainsGetInstance

> MarketplaceDomainResponse MarketplaceDomainsGetInstance(ctx, id).FieldsMarketplaceDomains(fieldsMarketplaceDomains).Execute()



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
	fieldsMarketplaceDomains := []string{"FieldsMarketplaceDomains_example"} // []string | the fields to include for returned resources of type marketplaceDomains (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceDomainsAPI.MarketplaceDomainsGetInstance(context.Background(), id).FieldsMarketplaceDomains(fieldsMarketplaceDomains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceDomainsAPI.MarketplaceDomainsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceDomainsGetInstance`: MarketplaceDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceDomainsAPI.MarketplaceDomainsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceDomainsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsMarketplaceDomains** | **[]string** | the fields to include for returned resources of type marketplaceDomains | 

### Return type

[**MarketplaceDomainResponse**](MarketplaceDomainResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

