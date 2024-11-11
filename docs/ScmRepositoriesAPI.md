# \ScmRepositoriesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScmRepositoriesGetCollection**](ScmRepositoriesAPI.md#ScmRepositoriesGetCollection) | **Get** /v1/scmRepositories | 
[**ScmRepositoriesGetInstance**](ScmRepositoriesAPI.md#ScmRepositoriesGetInstance) | **Get** /v1/scmRepositories/{id} | 
[**ScmRepositoriesGitReferencesGetToManyRelated**](ScmRepositoriesAPI.md#ScmRepositoriesGitReferencesGetToManyRelated) | **Get** /v1/scmRepositories/{id}/gitReferences | 
[**ScmRepositoriesPullRequestsGetToManyRelated**](ScmRepositoriesAPI.md#ScmRepositoriesPullRequestsGetToManyRelated) | **Get** /v1/scmRepositories/{id}/pullRequests | 



## ScmRepositoriesGetCollection

> ScmRepositoriesResponse ScmRepositoriesGetCollection(ctx).FilterId(filterId).FieldsScmRepositories(fieldsScmRepositories).Limit(limit).Include(include).Execute()



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
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScmRepositoriesAPI.ScmRepositoriesGetCollection(context.Background()).FilterId(filterId).FieldsScmRepositories(fieldsScmRepositories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScmRepositoriesAPI.ScmRepositoriesGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScmRepositoriesGetCollection`: ScmRepositoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `ScmRepositoriesAPI.ScmRepositoriesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScmRepositoriesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterId** | **[]string** | filter by id(s) | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
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


## ScmRepositoriesGetInstance

> ScmRepositoryResponse ScmRepositoriesGetInstance(ctx, id).FieldsScmRepositories(fieldsScmRepositories).Include(include).Execute()



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
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScmRepositoriesAPI.ScmRepositoriesGetInstance(context.Background(), id).FieldsScmRepositories(fieldsScmRepositories).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScmRepositoriesAPI.ScmRepositoriesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScmRepositoriesGetInstance`: ScmRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `ScmRepositoriesAPI.ScmRepositoriesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiScmRepositoriesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**ScmRepositoryResponse**](ScmRepositoryResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScmRepositoriesGitReferencesGetToManyRelated

> ScmGitReferencesResponse ScmRepositoriesGitReferencesGetToManyRelated(ctx, id).FieldsScmGitReferences(fieldsScmGitReferences).FieldsScmRepositories(fieldsScmRepositories).Limit(limit).Include(include).Execute()



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
	fieldsScmGitReferences := []string{"FieldsScmGitReferences_example"} // []string | the fields to include for returned resources of type scmGitReferences (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScmRepositoriesAPI.ScmRepositoriesGitReferencesGetToManyRelated(context.Background(), id).FieldsScmGitReferences(fieldsScmGitReferences).FieldsScmRepositories(fieldsScmRepositories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScmRepositoriesAPI.ScmRepositoriesGitReferencesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScmRepositoriesGitReferencesGetToManyRelated`: ScmGitReferencesResponse
	fmt.Fprintf(os.Stdout, "Response from `ScmRepositoriesAPI.ScmRepositoriesGitReferencesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiScmRepositoriesGitReferencesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsScmGitReferences** | **[]string** | the fields to include for returned resources of type scmGitReferences | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**ScmGitReferencesResponse**](ScmGitReferencesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScmRepositoriesPullRequestsGetToManyRelated

> ScmPullRequestsResponse ScmRepositoriesPullRequestsGetToManyRelated(ctx, id).FieldsScmPullRequests(fieldsScmPullRequests).FieldsScmRepositories(fieldsScmRepositories).Limit(limit).Include(include).Execute()



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
	fieldsScmPullRequests := []string{"FieldsScmPullRequests_example"} // []string | the fields to include for returned resources of type scmPullRequests (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScmRepositoriesAPI.ScmRepositoriesPullRequestsGetToManyRelated(context.Background(), id).FieldsScmPullRequests(fieldsScmPullRequests).FieldsScmRepositories(fieldsScmRepositories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScmRepositoriesAPI.ScmRepositoriesPullRequestsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScmRepositoriesPullRequestsGetToManyRelated`: ScmPullRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `ScmRepositoriesAPI.ScmRepositoriesPullRequestsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiScmRepositoriesPullRequestsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsScmPullRequests** | **[]string** | the fields to include for returned resources of type scmPullRequests | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**ScmPullRequestsResponse**](ScmPullRequestsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

