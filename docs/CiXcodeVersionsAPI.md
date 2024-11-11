# \CiXcodeVersionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CiXcodeVersionsGetCollection**](CiXcodeVersionsAPI.md#CiXcodeVersionsGetCollection) | **Get** /v1/ciXcodeVersions | 
[**CiXcodeVersionsGetInstance**](CiXcodeVersionsAPI.md#CiXcodeVersionsGetInstance) | **Get** /v1/ciXcodeVersions/{id} | 
[**CiXcodeVersionsMacOsVersionsGetToManyRelated**](CiXcodeVersionsAPI.md#CiXcodeVersionsMacOsVersionsGetToManyRelated) | **Get** /v1/ciXcodeVersions/{id}/macOsVersions | 



## CiXcodeVersionsGetCollection

> CiXcodeVersionsResponse CiXcodeVersionsGetCollection(ctx).FieldsCiXcodeVersions(fieldsCiXcodeVersions).FieldsCiMacOsVersions(fieldsCiMacOsVersions).Limit(limit).Include(include).LimitMacOsVersions(limitMacOsVersions).Execute()



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
	fieldsCiXcodeVersions := []string{"FieldsCiXcodeVersions_example"} // []string | the fields to include for returned resources of type ciXcodeVersions (optional)
	fieldsCiMacOsVersions := []string{"FieldsCiMacOsVersions_example"} // []string | the fields to include for returned resources of type ciMacOsVersions (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitMacOsVersions := int32(56) // int32 | maximum number of related macOsVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiXcodeVersionsAPI.CiXcodeVersionsGetCollection(context.Background()).FieldsCiXcodeVersions(fieldsCiXcodeVersions).FieldsCiMacOsVersions(fieldsCiMacOsVersions).Limit(limit).Include(include).LimitMacOsVersions(limitMacOsVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiXcodeVersionsAPI.CiXcodeVersionsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiXcodeVersionsGetCollection`: CiXcodeVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiXcodeVersionsAPI.CiXcodeVersionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCiXcodeVersionsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldsCiXcodeVersions** | **[]string** | the fields to include for returned resources of type ciXcodeVersions | 
 **fieldsCiMacOsVersions** | **[]string** | the fields to include for returned resources of type ciMacOsVersions | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitMacOsVersions** | **int32** | maximum number of related macOsVersions returned (when they are included) | 

### Return type

[**CiXcodeVersionsResponse**](CiXcodeVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiXcodeVersionsGetInstance

> CiXcodeVersionResponse CiXcodeVersionsGetInstance(ctx, id).FieldsCiXcodeVersions(fieldsCiXcodeVersions).FieldsCiMacOsVersions(fieldsCiMacOsVersions).Include(include).LimitMacOsVersions(limitMacOsVersions).Execute()



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
	fieldsCiXcodeVersions := []string{"FieldsCiXcodeVersions_example"} // []string | the fields to include for returned resources of type ciXcodeVersions (optional)
	fieldsCiMacOsVersions := []string{"FieldsCiMacOsVersions_example"} // []string | the fields to include for returned resources of type ciMacOsVersions (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitMacOsVersions := int32(56) // int32 | maximum number of related macOsVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiXcodeVersionsAPI.CiXcodeVersionsGetInstance(context.Background(), id).FieldsCiXcodeVersions(fieldsCiXcodeVersions).FieldsCiMacOsVersions(fieldsCiMacOsVersions).Include(include).LimitMacOsVersions(limitMacOsVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiXcodeVersionsAPI.CiXcodeVersionsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiXcodeVersionsGetInstance`: CiXcodeVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `CiXcodeVersionsAPI.CiXcodeVersionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiXcodeVersionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiXcodeVersions** | **[]string** | the fields to include for returned resources of type ciXcodeVersions | 
 **fieldsCiMacOsVersions** | **[]string** | the fields to include for returned resources of type ciMacOsVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitMacOsVersions** | **int32** | maximum number of related macOsVersions returned (when they are included) | 

### Return type

[**CiXcodeVersionResponse**](CiXcodeVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiXcodeVersionsMacOsVersionsGetToManyRelated

> CiMacOsVersionsResponse CiXcodeVersionsMacOsVersionsGetToManyRelated(ctx, id).FieldsCiMacOsVersions(fieldsCiMacOsVersions).FieldsCiXcodeVersions(fieldsCiXcodeVersions).Limit(limit).Include(include).LimitXcodeVersions(limitXcodeVersions).Execute()



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
	fieldsCiMacOsVersions := []string{"FieldsCiMacOsVersions_example"} // []string | the fields to include for returned resources of type ciMacOsVersions (optional)
	fieldsCiXcodeVersions := []string{"FieldsCiXcodeVersions_example"} // []string | the fields to include for returned resources of type ciXcodeVersions (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitXcodeVersions := int32(56) // int32 | maximum number of related xcodeVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiXcodeVersionsAPI.CiXcodeVersionsMacOsVersionsGetToManyRelated(context.Background(), id).FieldsCiMacOsVersions(fieldsCiMacOsVersions).FieldsCiXcodeVersions(fieldsCiXcodeVersions).Limit(limit).Include(include).LimitXcodeVersions(limitXcodeVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiXcodeVersionsAPI.CiXcodeVersionsMacOsVersionsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiXcodeVersionsMacOsVersionsGetToManyRelated`: CiMacOsVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiXcodeVersionsAPI.CiXcodeVersionsMacOsVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiXcodeVersionsMacOsVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiMacOsVersions** | **[]string** | the fields to include for returned resources of type ciMacOsVersions | 
 **fieldsCiXcodeVersions** | **[]string** | the fields to include for returned resources of type ciXcodeVersions | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitXcodeVersions** | **int32** | maximum number of related xcodeVersions returned (when they are included) | 

### Return type

[**CiMacOsVersionsResponse**](CiMacOsVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

