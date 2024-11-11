# \CiMacOsVersionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CiMacOsVersionsGetCollection**](CiMacOsVersionsAPI.md#CiMacOsVersionsGetCollection) | **Get** /v1/ciMacOsVersions | 
[**CiMacOsVersionsGetInstance**](CiMacOsVersionsAPI.md#CiMacOsVersionsGetInstance) | **Get** /v1/ciMacOsVersions/{id} | 
[**CiMacOsVersionsXcodeVersionsGetToManyRelated**](CiMacOsVersionsAPI.md#CiMacOsVersionsXcodeVersionsGetToManyRelated) | **Get** /v1/ciMacOsVersions/{id}/xcodeVersions | 



## CiMacOsVersionsGetCollection

> CiMacOsVersionsResponse CiMacOsVersionsGetCollection(ctx).FieldsCiMacOsVersions(fieldsCiMacOsVersions).FieldsCiXcodeVersions(fieldsCiXcodeVersions).Limit(limit).Include(include).LimitXcodeVersions(limitXcodeVersions).Execute()



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
	fieldsCiMacOsVersions := []string{"FieldsCiMacOsVersions_example"} // []string | the fields to include for returned resources of type ciMacOsVersions (optional)
	fieldsCiXcodeVersions := []string{"FieldsCiXcodeVersions_example"} // []string | the fields to include for returned resources of type ciXcodeVersions (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitXcodeVersions := int32(56) // int32 | maximum number of related xcodeVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiMacOsVersionsAPI.CiMacOsVersionsGetCollection(context.Background()).FieldsCiMacOsVersions(fieldsCiMacOsVersions).FieldsCiXcodeVersions(fieldsCiXcodeVersions).Limit(limit).Include(include).LimitXcodeVersions(limitXcodeVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiMacOsVersionsAPI.CiMacOsVersionsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiMacOsVersionsGetCollection`: CiMacOsVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiMacOsVersionsAPI.CiMacOsVersionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCiMacOsVersionsGetCollectionRequest struct via the builder pattern


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


## CiMacOsVersionsGetInstance

> CiMacOsVersionResponse CiMacOsVersionsGetInstance(ctx, id).FieldsCiMacOsVersions(fieldsCiMacOsVersions).FieldsCiXcodeVersions(fieldsCiXcodeVersions).Include(include).LimitXcodeVersions(limitXcodeVersions).Execute()



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
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitXcodeVersions := int32(56) // int32 | maximum number of related xcodeVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiMacOsVersionsAPI.CiMacOsVersionsGetInstance(context.Background(), id).FieldsCiMacOsVersions(fieldsCiMacOsVersions).FieldsCiXcodeVersions(fieldsCiXcodeVersions).Include(include).LimitXcodeVersions(limitXcodeVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiMacOsVersionsAPI.CiMacOsVersionsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiMacOsVersionsGetInstance`: CiMacOsVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `CiMacOsVersionsAPI.CiMacOsVersionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiMacOsVersionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiMacOsVersions** | **[]string** | the fields to include for returned resources of type ciMacOsVersions | 
 **fieldsCiXcodeVersions** | **[]string** | the fields to include for returned resources of type ciXcodeVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitXcodeVersions** | **int32** | maximum number of related xcodeVersions returned (when they are included) | 

### Return type

[**CiMacOsVersionResponse**](CiMacOsVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiMacOsVersionsXcodeVersionsGetToManyRelated

> CiXcodeVersionsResponse CiMacOsVersionsXcodeVersionsGetToManyRelated(ctx, id).FieldsCiXcodeVersions(fieldsCiXcodeVersions).FieldsCiMacOsVersions(fieldsCiMacOsVersions).Limit(limit).Include(include).LimitMacOsVersions(limitMacOsVersions).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitMacOsVersions := int32(56) // int32 | maximum number of related macOsVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiMacOsVersionsAPI.CiMacOsVersionsXcodeVersionsGetToManyRelated(context.Background(), id).FieldsCiXcodeVersions(fieldsCiXcodeVersions).FieldsCiMacOsVersions(fieldsCiMacOsVersions).Limit(limit).Include(include).LimitMacOsVersions(limitMacOsVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiMacOsVersionsAPI.CiMacOsVersionsXcodeVersionsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiMacOsVersionsXcodeVersionsGetToManyRelated`: CiXcodeVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiMacOsVersionsAPI.CiMacOsVersionsXcodeVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiMacOsVersionsXcodeVersionsGetToManyRelatedRequest struct via the builder pattern


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

