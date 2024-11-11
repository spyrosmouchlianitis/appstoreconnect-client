# \AlternativeDistributionKeysAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlternativeDistributionKeysCreateInstance**](AlternativeDistributionKeysAPI.md#AlternativeDistributionKeysCreateInstance) | **Post** /v1/alternativeDistributionKeys | 
[**AlternativeDistributionKeysDeleteInstance**](AlternativeDistributionKeysAPI.md#AlternativeDistributionKeysDeleteInstance) | **Delete** /v1/alternativeDistributionKeys/{id} | 
[**AlternativeDistributionKeysGetCollection**](AlternativeDistributionKeysAPI.md#AlternativeDistributionKeysGetCollection) | **Get** /v1/alternativeDistributionKeys | 
[**AlternativeDistributionKeysGetInstance**](AlternativeDistributionKeysAPI.md#AlternativeDistributionKeysGetInstance) | **Get** /v1/alternativeDistributionKeys/{id} | 



## AlternativeDistributionKeysCreateInstance

> AlternativeDistributionKeyResponse AlternativeDistributionKeysCreateInstance(ctx).AlternativeDistributionKeyCreateRequest(alternativeDistributionKeyCreateRequest).Execute()



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
	alternativeDistributionKeyCreateRequest := *openapiclient.NewAlternativeDistributionKeyCreateRequest(*openapiclient.NewAlternativeDistributionKeyCreateRequestData("Type_example", *openapiclient.NewAlternativeDistributionKeyCreateRequestDataAttributes("PublicKey_example"))) // AlternativeDistributionKeyCreateRequest | AlternativeDistributionKey representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionKeysAPI.AlternativeDistributionKeysCreateInstance(context.Background()).AlternativeDistributionKeyCreateRequest(alternativeDistributionKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionKeysAPI.AlternativeDistributionKeysCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionKeysCreateInstance`: AlternativeDistributionKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionKeysAPI.AlternativeDistributionKeysCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionKeysCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alternativeDistributionKeyCreateRequest** | [**AlternativeDistributionKeyCreateRequest**](AlternativeDistributionKeyCreateRequest.md) | AlternativeDistributionKey representation | 

### Return type

[**AlternativeDistributionKeyResponse**](AlternativeDistributionKeyResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlternativeDistributionKeysDeleteInstance

> AlternativeDistributionKeysDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AlternativeDistributionKeysAPI.AlternativeDistributionKeysDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionKeysAPI.AlternativeDistributionKeysDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAlternativeDistributionKeysDeleteInstanceRequest struct via the builder pattern


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


## AlternativeDistributionKeysGetCollection

> AlternativeDistributionKeysResponse AlternativeDistributionKeysGetCollection(ctx).ExistsApp(existsApp).FieldsAlternativeDistributionKeys(fieldsAlternativeDistributionKeys).Limit(limit).Execute()



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
	existsApp := true // bool | filter by existence or non-existence of related 'app' (optional)
	fieldsAlternativeDistributionKeys := []string{"FieldsAlternativeDistributionKeys_example"} // []string | the fields to include for returned resources of type alternativeDistributionKeys (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionKeysAPI.AlternativeDistributionKeysGetCollection(context.Background()).ExistsApp(existsApp).FieldsAlternativeDistributionKeys(fieldsAlternativeDistributionKeys).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionKeysAPI.AlternativeDistributionKeysGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionKeysGetCollection`: AlternativeDistributionKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionKeysAPI.AlternativeDistributionKeysGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionKeysGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **existsApp** | **bool** | filter by existence or non-existence of related &#39;app&#39; | 
 **fieldsAlternativeDistributionKeys** | **[]string** | the fields to include for returned resources of type alternativeDistributionKeys | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AlternativeDistributionKeysResponse**](AlternativeDistributionKeysResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlternativeDistributionKeysGetInstance

> AlternativeDistributionKeyResponse AlternativeDistributionKeysGetInstance(ctx, id).FieldsAlternativeDistributionKeys(fieldsAlternativeDistributionKeys).Execute()



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
	fieldsAlternativeDistributionKeys := []string{"FieldsAlternativeDistributionKeys_example"} // []string | the fields to include for returned resources of type alternativeDistributionKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionKeysAPI.AlternativeDistributionKeysGetInstance(context.Background(), id).FieldsAlternativeDistributionKeys(fieldsAlternativeDistributionKeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionKeysAPI.AlternativeDistributionKeysGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionKeysGetInstance`: AlternativeDistributionKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionKeysAPI.AlternativeDistributionKeysGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionKeysGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionKeys** | **[]string** | the fields to include for returned resources of type alternativeDistributionKeys | 

### Return type

[**AlternativeDistributionKeyResponse**](AlternativeDistributionKeyResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

