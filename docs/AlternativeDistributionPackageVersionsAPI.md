# \AlternativeDistributionPackageVersionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlternativeDistributionPackageVersionsDeltasGetToManyRelated**](AlternativeDistributionPackageVersionsAPI.md#AlternativeDistributionPackageVersionsDeltasGetToManyRelated) | **Get** /v1/alternativeDistributionPackageVersions/{id}/deltas | 
[**AlternativeDistributionPackageVersionsGetInstance**](AlternativeDistributionPackageVersionsAPI.md#AlternativeDistributionPackageVersionsGetInstance) | **Get** /v1/alternativeDistributionPackageVersions/{id} | 
[**AlternativeDistributionPackageVersionsVariantsGetToManyRelated**](AlternativeDistributionPackageVersionsAPI.md#AlternativeDistributionPackageVersionsVariantsGetToManyRelated) | **Get** /v1/alternativeDistributionPackageVersions/{id}/variants | 



## AlternativeDistributionPackageVersionsDeltasGetToManyRelated

> AlternativeDistributionPackageDeltasResponse AlternativeDistributionPackageVersionsDeltasGetToManyRelated(ctx, id).FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas).Limit(limit).Execute()



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
	fieldsAlternativeDistributionPackageDeltas := []string{"FieldsAlternativeDistributionPackageDeltas_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageDeltas (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsDeltasGetToManyRelated(context.Background(), id).FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsDeltasGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionPackageVersionsDeltasGetToManyRelated`: AlternativeDistributionPackageDeltasResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsDeltasGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionPackageDeltas** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageDeltas | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AlternativeDistributionPackageDeltasResponse**](AlternativeDistributionPackageDeltasResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlternativeDistributionPackageVersionsGetInstance

> AlternativeDistributionPackageVersionResponse AlternativeDistributionPackageVersionsGetInstance(ctx, id).FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions).FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants).FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas).Include(include).LimitDeltas(limitDeltas).LimitVariants(limitVariants).Execute()



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
	fieldsAlternativeDistributionPackageVersions := []string{"FieldsAlternativeDistributionPackageVersions_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageVersions (optional)
	fieldsAlternativeDistributionPackageVariants := []string{"FieldsAlternativeDistributionPackageVariants_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageVariants (optional)
	fieldsAlternativeDistributionPackageDeltas := []string{"FieldsAlternativeDistributionPackageDeltas_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageDeltas (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitDeltas := int32(56) // int32 | maximum number of related deltas returned (when they are included) (optional)
	limitVariants := int32(56) // int32 | maximum number of related variants returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsGetInstance(context.Background(), id).FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions).FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants).FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas).Include(include).LimitDeltas(limitDeltas).LimitVariants(limitVariants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionPackageVersionsGetInstance`: AlternativeDistributionPackageVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionPackageVersionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionPackageVersions** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageVersions | 
 **fieldsAlternativeDistributionPackageVariants** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageVariants | 
 **fieldsAlternativeDistributionPackageDeltas** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageDeltas | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitDeltas** | **int32** | maximum number of related deltas returned (when they are included) | 
 **limitVariants** | **int32** | maximum number of related variants returned (when they are included) | 

### Return type

[**AlternativeDistributionPackageVersionResponse**](AlternativeDistributionPackageVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlternativeDistributionPackageVersionsVariantsGetToManyRelated

> AlternativeDistributionPackageVariantsResponse AlternativeDistributionPackageVersionsVariantsGetToManyRelated(ctx, id).FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants).Limit(limit).Execute()



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
	fieldsAlternativeDistributionPackageVariants := []string{"FieldsAlternativeDistributionPackageVariants_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageVariants (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsVariantsGetToManyRelated(context.Background(), id).FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsVariantsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionPackageVersionsVariantsGetToManyRelated`: AlternativeDistributionPackageVariantsResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionPackageVersionsAPI.AlternativeDistributionPackageVersionsVariantsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionPackageVariants** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageVariants | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AlternativeDistributionPackageVariantsResponse**](AlternativeDistributionPackageVariantsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

