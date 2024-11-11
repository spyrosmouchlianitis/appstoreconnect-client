# \AlternativeDistributionDomainsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlternativeDistributionDomainsCreateInstance**](AlternativeDistributionDomainsAPI.md#AlternativeDistributionDomainsCreateInstance) | **Post** /v1/alternativeDistributionDomains | 
[**AlternativeDistributionDomainsDeleteInstance**](AlternativeDistributionDomainsAPI.md#AlternativeDistributionDomainsDeleteInstance) | **Delete** /v1/alternativeDistributionDomains/{id} | 
[**AlternativeDistributionDomainsGetCollection**](AlternativeDistributionDomainsAPI.md#AlternativeDistributionDomainsGetCollection) | **Get** /v1/alternativeDistributionDomains | 
[**AlternativeDistributionDomainsGetInstance**](AlternativeDistributionDomainsAPI.md#AlternativeDistributionDomainsGetInstance) | **Get** /v1/alternativeDistributionDomains/{id} | 



## AlternativeDistributionDomainsCreateInstance

> AlternativeDistributionDomainResponse AlternativeDistributionDomainsCreateInstance(ctx).AlternativeDistributionDomainCreateRequest(alternativeDistributionDomainCreateRequest).Execute()



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
	alternativeDistributionDomainCreateRequest := *openapiclient.NewAlternativeDistributionDomainCreateRequest(*openapiclient.NewAlternativeDistributionDomainCreateRequestData("Type_example", *openapiclient.NewAlternativeDistributionDomainCreateRequestDataAttributes("Domain_example", "ReferenceName_example"))) // AlternativeDistributionDomainCreateRequest | AlternativeDistributionDomain representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsCreateInstance(context.Background()).AlternativeDistributionDomainCreateRequest(alternativeDistributionDomainCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionDomainsCreateInstance`: AlternativeDistributionDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionDomainsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alternativeDistributionDomainCreateRequest** | [**AlternativeDistributionDomainCreateRequest**](AlternativeDistributionDomainCreateRequest.md) | AlternativeDistributionDomain representation | 

### Return type

[**AlternativeDistributionDomainResponse**](AlternativeDistributionDomainResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlternativeDistributionDomainsDeleteInstance

> AlternativeDistributionDomainsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlternativeDistributionDomainsDeleteInstanceRequest struct via the builder pattern


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


## AlternativeDistributionDomainsGetCollection

> AlternativeDistributionDomainsResponse AlternativeDistributionDomainsGetCollection(ctx).FieldsAlternativeDistributionDomains(fieldsAlternativeDistributionDomains).Limit(limit).Execute()



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
	fieldsAlternativeDistributionDomains := []string{"FieldsAlternativeDistributionDomains_example"} // []string | the fields to include for returned resources of type alternativeDistributionDomains (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsGetCollection(context.Background()).FieldsAlternativeDistributionDomains(fieldsAlternativeDistributionDomains).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionDomainsGetCollection`: AlternativeDistributionDomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionDomainsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldsAlternativeDistributionDomains** | **[]string** | the fields to include for returned resources of type alternativeDistributionDomains | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AlternativeDistributionDomainsResponse**](AlternativeDistributionDomainsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlternativeDistributionDomainsGetInstance

> AlternativeDistributionDomainResponse AlternativeDistributionDomainsGetInstance(ctx, id).FieldsAlternativeDistributionDomains(fieldsAlternativeDistributionDomains).Execute()



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
	fieldsAlternativeDistributionDomains := []string{"FieldsAlternativeDistributionDomains_example"} // []string | the fields to include for returned resources of type alternativeDistributionDomains (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsGetInstance(context.Background(), id).FieldsAlternativeDistributionDomains(fieldsAlternativeDistributionDomains).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionDomainsGetInstance`: AlternativeDistributionDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionDomainsAPI.AlternativeDistributionDomainsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionDomainsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionDomains** | **[]string** | the fields to include for returned resources of type alternativeDistributionDomains | 

### Return type

[**AlternativeDistributionDomainResponse**](AlternativeDistributionDomainResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

