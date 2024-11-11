# \ScmProvidersAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScmProvidersGetCollection**](ScmProvidersAPI.md#ScmProvidersGetCollection) | **Get** /v1/scmProviders | 
[**ScmProvidersGetInstance**](ScmProvidersAPI.md#ScmProvidersGetInstance) | **Get** /v1/scmProviders/{id} | 
[**ScmProvidersRepositoriesGetToManyRelated**](ScmProvidersAPI.md#ScmProvidersRepositoriesGetToManyRelated) | **Get** /v1/scmProviders/{id}/repositories | 



## ScmProvidersGetCollection

> ScmProvidersResponse ScmProvidersGetCollection(ctx).FieldsScmProviders(fieldsScmProviders).Limit(limit).Execute()



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
	fieldsScmProviders := []string{"FieldsScmProviders_example"} // []string | the fields to include for returned resources of type scmProviders (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScmProvidersAPI.ScmProvidersGetCollection(context.Background()).FieldsScmProviders(fieldsScmProviders).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScmProvidersAPI.ScmProvidersGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScmProvidersGetCollection`: ScmProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `ScmProvidersAPI.ScmProvidersGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScmProvidersGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldsScmProviders** | **[]string** | the fields to include for returned resources of type scmProviders | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**ScmProvidersResponse**](ScmProvidersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScmProvidersGetInstance

> ScmProviderResponse ScmProvidersGetInstance(ctx, id).FieldsScmProviders(fieldsScmProviders).Execute()



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
	fieldsScmProviders := []string{"FieldsScmProviders_example"} // []string | the fields to include for returned resources of type scmProviders (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScmProvidersAPI.ScmProvidersGetInstance(context.Background(), id).FieldsScmProviders(fieldsScmProviders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScmProvidersAPI.ScmProvidersGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScmProvidersGetInstance`: ScmProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ScmProvidersAPI.ScmProvidersGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiScmProvidersGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsScmProviders** | **[]string** | the fields to include for returned resources of type scmProviders | 

### Return type

[**ScmProviderResponse**](ScmProviderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScmProvidersRepositoriesGetToManyRelated

> ScmRepositoriesResponse ScmProvidersRepositoriesGetToManyRelated(ctx, id).FilterId(filterId).FieldsScmRepositories(fieldsScmRepositories).FieldsScmProviders(fieldsScmProviders).FieldsScmGitReferences(fieldsScmGitReferences).Limit(limit).Include(include).Execute()



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
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	fieldsScmProviders := []string{"FieldsScmProviders_example"} // []string | the fields to include for returned resources of type scmProviders (optional)
	fieldsScmGitReferences := []string{"FieldsScmGitReferences_example"} // []string | the fields to include for returned resources of type scmGitReferences (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScmProvidersAPI.ScmProvidersRepositoriesGetToManyRelated(context.Background(), id).FilterId(filterId).FieldsScmRepositories(fieldsScmRepositories).FieldsScmProviders(fieldsScmProviders).FieldsScmGitReferences(fieldsScmGitReferences).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScmProvidersAPI.ScmProvidersRepositoriesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScmProvidersRepositoriesGetToManyRelated`: ScmRepositoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `ScmProvidersAPI.ScmProvidersRepositoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiScmProvidersRepositoriesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterId** | **[]string** | filter by id(s) | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **fieldsScmProviders** | **[]string** | the fields to include for returned resources of type scmProviders | 
 **fieldsScmGitReferences** | **[]string** | the fields to include for returned resources of type scmGitReferences | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**ScmRepositoriesResponse**](ScmRepositoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

