# \AlternativeDistributionPackagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlternativeDistributionPackagesCreateInstance**](AlternativeDistributionPackagesAPI.md#AlternativeDistributionPackagesCreateInstance) | **Post** /v1/alternativeDistributionPackages | 
[**AlternativeDistributionPackagesGetInstance**](AlternativeDistributionPackagesAPI.md#AlternativeDistributionPackagesGetInstance) | **Get** /v1/alternativeDistributionPackages/{id} | 
[**AlternativeDistributionPackagesVersionsGetToManyRelated**](AlternativeDistributionPackagesAPI.md#AlternativeDistributionPackagesVersionsGetToManyRelated) | **Get** /v1/alternativeDistributionPackages/{id}/versions | 



## AlternativeDistributionPackagesCreateInstance

> AlternativeDistributionPackageResponse AlternativeDistributionPackagesCreateInstance(ctx).AlternativeDistributionPackageCreateRequest(alternativeDistributionPackageCreateRequest).Execute()



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
	alternativeDistributionPackageCreateRequest := *openapiclient.NewAlternativeDistributionPackageCreateRequest(*openapiclient.NewAlternativeDistributionPackageCreateRequestData("Type_example", *openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationships(*openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData("Type_example", "Id_example"))))) // AlternativeDistributionPackageCreateRequest | AlternativeDistributionPackage representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesCreateInstance(context.Background()).AlternativeDistributionPackageCreateRequest(alternativeDistributionPackageCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionPackagesCreateInstance`: AlternativeDistributionPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionPackagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alternativeDistributionPackageCreateRequest** | [**AlternativeDistributionPackageCreateRequest**](AlternativeDistributionPackageCreateRequest.md) | AlternativeDistributionPackage representation | 

### Return type

[**AlternativeDistributionPackageResponse**](AlternativeDistributionPackageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlternativeDistributionPackagesGetInstance

> AlternativeDistributionPackageResponse AlternativeDistributionPackagesGetInstance(ctx, id).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions).Include(include).LimitVersions(limitVersions).Execute()



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
	fieldsAlternativeDistributionPackages := []string{"FieldsAlternativeDistributionPackages_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackages (optional)
	fieldsAlternativeDistributionPackageVersions := []string{"FieldsAlternativeDistributionPackageVersions_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageVersions (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitVersions := int32(56) // int32 | maximum number of related versions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesGetInstance(context.Background(), id).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions).Include(include).LimitVersions(limitVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionPackagesGetInstance`: AlternativeDistributionPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionPackagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionPackages** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackages | 
 **fieldsAlternativeDistributionPackageVersions** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitVersions** | **int32** | maximum number of related versions returned (when they are included) | 

### Return type

[**AlternativeDistributionPackageResponse**](AlternativeDistributionPackageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlternativeDistributionPackagesVersionsGetToManyRelated

> AlternativeDistributionPackageVersionsResponse AlternativeDistributionPackagesVersionsGetToManyRelated(ctx, id).FilterState(filterState).FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions).FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants).FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).Limit(limit).Include(include).LimitVariants(limitVariants).LimitDeltas(limitDeltas).Execute()



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
	filterState := []string{"FilterState_example"} // []string | filter by attribute 'state' (optional)
	fieldsAlternativeDistributionPackageVersions := []string{"FieldsAlternativeDistributionPackageVersions_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageVersions (optional)
	fieldsAlternativeDistributionPackageVariants := []string{"FieldsAlternativeDistributionPackageVariants_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageVariants (optional)
	fieldsAlternativeDistributionPackageDeltas := []string{"FieldsAlternativeDistributionPackageDeltas_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageDeltas (optional)
	fieldsAlternativeDistributionPackages := []string{"FieldsAlternativeDistributionPackages_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackages (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitVariants := int32(56) // int32 | maximum number of related variants returned (when they are included) (optional)
	limitDeltas := int32(56) // int32 | maximum number of related deltas returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesVersionsGetToManyRelated(context.Background(), id).FilterState(filterState).FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions).FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants).FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).Limit(limit).Include(include).LimitVariants(limitVariants).LimitDeltas(limitDeltas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesVersionsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionPackagesVersionsGetToManyRelated`: AlternativeDistributionPackageVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionPackagesAPI.AlternativeDistributionPackagesVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionPackagesVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterState** | **[]string** | filter by attribute &#39;state&#39; | 
 **fieldsAlternativeDistributionPackageVersions** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageVersions | 
 **fieldsAlternativeDistributionPackageVariants** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageVariants | 
 **fieldsAlternativeDistributionPackageDeltas** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageDeltas | 
 **fieldsAlternativeDistributionPackages** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackages | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitVariants** | **int32** | maximum number of related variants returned (when they are included) | 
 **limitDeltas** | **int32** | maximum number of related deltas returned (when they are included) | 

### Return type

[**AlternativeDistributionPackageVersionsResponse**](AlternativeDistributionPackageVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

