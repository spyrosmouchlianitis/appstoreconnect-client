# \CiBuildActionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CiBuildActionsArtifactsGetToManyRelated**](CiBuildActionsAPI.md#CiBuildActionsArtifactsGetToManyRelated) | **Get** /v1/ciBuildActions/{id}/artifacts | 
[**CiBuildActionsBuildRunGetToOneRelated**](CiBuildActionsAPI.md#CiBuildActionsBuildRunGetToOneRelated) | **Get** /v1/ciBuildActions/{id}/buildRun | 
[**CiBuildActionsGetInstance**](CiBuildActionsAPI.md#CiBuildActionsGetInstance) | **Get** /v1/ciBuildActions/{id} | 
[**CiBuildActionsIssuesGetToManyRelated**](CiBuildActionsAPI.md#CiBuildActionsIssuesGetToManyRelated) | **Get** /v1/ciBuildActions/{id}/issues | 
[**CiBuildActionsTestResultsGetToManyRelated**](CiBuildActionsAPI.md#CiBuildActionsTestResultsGetToManyRelated) | **Get** /v1/ciBuildActions/{id}/testResults | 



## CiBuildActionsArtifactsGetToManyRelated

> CiArtifactsResponse CiBuildActionsArtifactsGetToManyRelated(ctx, id).FieldsCiArtifacts(fieldsCiArtifacts).Limit(limit).Execute()



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
	fieldsCiArtifacts := []string{"FieldsCiArtifacts_example"} // []string | the fields to include for returned resources of type ciArtifacts (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildActionsAPI.CiBuildActionsArtifactsGetToManyRelated(context.Background(), id).FieldsCiArtifacts(fieldsCiArtifacts).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildActionsAPI.CiBuildActionsArtifactsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildActionsArtifactsGetToManyRelated`: CiArtifactsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildActionsAPI.CiBuildActionsArtifactsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildActionsArtifactsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiArtifacts** | **[]string** | the fields to include for returned resources of type ciArtifacts | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**CiArtifactsResponse**](CiArtifactsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiBuildActionsBuildRunGetToOneRelated

> CiBuildRunResponse CiBuildActionsBuildRunGetToOneRelated(ctx, id).FieldsCiBuildRuns(fieldsCiBuildRuns).FieldsBuilds(fieldsBuilds).FieldsCiWorkflows(fieldsCiWorkflows).FieldsCiProducts(fieldsCiProducts).FieldsScmGitReferences(fieldsScmGitReferences).FieldsScmPullRequests(fieldsScmPullRequests).Include(include).LimitBuilds(limitBuilds).Execute()



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
	fieldsCiBuildRuns := []string{"FieldsCiBuildRuns_example"} // []string | the fields to include for returned resources of type ciBuildRuns (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	fieldsCiWorkflows := []string{"FieldsCiWorkflows_example"} // []string | the fields to include for returned resources of type ciWorkflows (optional)
	fieldsCiProducts := []string{"FieldsCiProducts_example"} // []string | the fields to include for returned resources of type ciProducts (optional)
	fieldsScmGitReferences := []string{"FieldsScmGitReferences_example"} // []string | the fields to include for returned resources of type scmGitReferences (optional)
	fieldsScmPullRequests := []string{"FieldsScmPullRequests_example"} // []string | the fields to include for returned resources of type scmPullRequests (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildActionsAPI.CiBuildActionsBuildRunGetToOneRelated(context.Background(), id).FieldsCiBuildRuns(fieldsCiBuildRuns).FieldsBuilds(fieldsBuilds).FieldsCiWorkflows(fieldsCiWorkflows).FieldsCiProducts(fieldsCiProducts).FieldsScmGitReferences(fieldsScmGitReferences).FieldsScmPullRequests(fieldsScmPullRequests).Include(include).LimitBuilds(limitBuilds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildActionsAPI.CiBuildActionsBuildRunGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildActionsBuildRunGetToOneRelated`: CiBuildRunResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildActionsAPI.CiBuildActionsBuildRunGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildActionsBuildRunGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiBuildRuns** | **[]string** | the fields to include for returned resources of type ciBuildRuns | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsCiWorkflows** | **[]string** | the fields to include for returned resources of type ciWorkflows | 
 **fieldsCiProducts** | **[]string** | the fields to include for returned resources of type ciProducts | 
 **fieldsScmGitReferences** | **[]string** | the fields to include for returned resources of type scmGitReferences | 
 **fieldsScmPullRequests** | **[]string** | the fields to include for returned resources of type scmPullRequests | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**CiBuildRunResponse**](CiBuildRunResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiBuildActionsGetInstance

> CiBuildActionResponse CiBuildActionsGetInstance(ctx, id).FieldsCiBuildActions(fieldsCiBuildActions).FieldsCiBuildRuns(fieldsCiBuildRuns).Include(include).Execute()



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
	fieldsCiBuildActions := []string{"FieldsCiBuildActions_example"} // []string | the fields to include for returned resources of type ciBuildActions (optional)
	fieldsCiBuildRuns := []string{"FieldsCiBuildRuns_example"} // []string | the fields to include for returned resources of type ciBuildRuns (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildActionsAPI.CiBuildActionsGetInstance(context.Background(), id).FieldsCiBuildActions(fieldsCiBuildActions).FieldsCiBuildRuns(fieldsCiBuildRuns).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildActionsAPI.CiBuildActionsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildActionsGetInstance`: CiBuildActionResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildActionsAPI.CiBuildActionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildActionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiBuildActions** | **[]string** | the fields to include for returned resources of type ciBuildActions | 
 **fieldsCiBuildRuns** | **[]string** | the fields to include for returned resources of type ciBuildRuns | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**CiBuildActionResponse**](CiBuildActionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiBuildActionsIssuesGetToManyRelated

> CiIssuesResponse CiBuildActionsIssuesGetToManyRelated(ctx, id).FieldsCiIssues(fieldsCiIssues).Limit(limit).Execute()



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
	fieldsCiIssues := []string{"FieldsCiIssues_example"} // []string | the fields to include for returned resources of type ciIssues (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildActionsAPI.CiBuildActionsIssuesGetToManyRelated(context.Background(), id).FieldsCiIssues(fieldsCiIssues).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildActionsAPI.CiBuildActionsIssuesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildActionsIssuesGetToManyRelated`: CiIssuesResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildActionsAPI.CiBuildActionsIssuesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildActionsIssuesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiIssues** | **[]string** | the fields to include for returned resources of type ciIssues | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**CiIssuesResponse**](CiIssuesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiBuildActionsTestResultsGetToManyRelated

> CiTestResultsResponse CiBuildActionsTestResultsGetToManyRelated(ctx, id).FieldsCiTestResults(fieldsCiTestResults).Limit(limit).Execute()



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
	fieldsCiTestResults := []string{"FieldsCiTestResults_example"} // []string | the fields to include for returned resources of type ciTestResults (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildActionsAPI.CiBuildActionsTestResultsGetToManyRelated(context.Background(), id).FieldsCiTestResults(fieldsCiTestResults).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildActionsAPI.CiBuildActionsTestResultsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildActionsTestResultsGetToManyRelated`: CiTestResultsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildActionsAPI.CiBuildActionsTestResultsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildActionsTestResultsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiTestResults** | **[]string** | the fields to include for returned resources of type ciTestResults | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**CiTestResultsResponse**](CiTestResultsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

