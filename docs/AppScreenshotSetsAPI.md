# \AppScreenshotSetsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppScreenshotSetsAppScreenshotsGetToManyRelated**](AppScreenshotSetsAPI.md#AppScreenshotSetsAppScreenshotsGetToManyRelated) | **Get** /v1/appScreenshotSets/{id}/appScreenshots | 
[**AppScreenshotSetsAppScreenshotsGetToManyRelationship**](AppScreenshotSetsAPI.md#AppScreenshotSetsAppScreenshotsGetToManyRelationship) | **Get** /v1/appScreenshotSets/{id}/relationships/appScreenshots | 
[**AppScreenshotSetsAppScreenshotsReplaceToManyRelationship**](AppScreenshotSetsAPI.md#AppScreenshotSetsAppScreenshotsReplaceToManyRelationship) | **Patch** /v1/appScreenshotSets/{id}/relationships/appScreenshots | 
[**AppScreenshotSetsCreateInstance**](AppScreenshotSetsAPI.md#AppScreenshotSetsCreateInstance) | **Post** /v1/appScreenshotSets | 
[**AppScreenshotSetsDeleteInstance**](AppScreenshotSetsAPI.md#AppScreenshotSetsDeleteInstance) | **Delete** /v1/appScreenshotSets/{id} | 
[**AppScreenshotSetsGetInstance**](AppScreenshotSetsAPI.md#AppScreenshotSetsGetInstance) | **Get** /v1/appScreenshotSets/{id} | 



## AppScreenshotSetsAppScreenshotsGetToManyRelated

> AppScreenshotsResponse AppScreenshotSetsAppScreenshotsGetToManyRelated(ctx, id).FieldsAppScreenshots(fieldsAppScreenshots).FieldsAppScreenshotSets(fieldsAppScreenshotSets).Limit(limit).Include(include).Execute()



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
	fieldsAppScreenshots := []string{"FieldsAppScreenshots_example"} // []string | the fields to include for returned resources of type appScreenshots (optional)
	fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppScreenshotSetsAPI.AppScreenshotSetsAppScreenshotsGetToManyRelated(context.Background(), id).FieldsAppScreenshots(fieldsAppScreenshots).FieldsAppScreenshotSets(fieldsAppScreenshotSets).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsAPI.AppScreenshotSetsAppScreenshotsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppScreenshotSetsAppScreenshotsGetToManyRelated`: AppScreenshotsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppScreenshotSetsAPI.AppScreenshotSetsAppScreenshotsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsAppScreenshotsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppScreenshots** | **[]string** | the fields to include for returned resources of type appScreenshots | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppScreenshotsResponse**](AppScreenshotsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsAppScreenshotsGetToManyRelationship

> AppScreenshotSetAppScreenshotsLinkagesResponse AppScreenshotSetsAppScreenshotsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppScreenshotSetsAPI.AppScreenshotSetsAppScreenshotsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsAPI.AppScreenshotSetsAppScreenshotsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppScreenshotSetsAppScreenshotsGetToManyRelationship`: AppScreenshotSetAppScreenshotsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `AppScreenshotSetsAPI.AppScreenshotSetsAppScreenshotsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsAppScreenshotsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**AppScreenshotSetAppScreenshotsLinkagesResponse**](AppScreenshotSetAppScreenshotsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsAppScreenshotsReplaceToManyRelationship

> AppScreenshotSetsAppScreenshotsReplaceToManyRelationship(ctx, id).AppScreenshotSetAppScreenshotsLinkagesRequest(appScreenshotSetAppScreenshotsLinkagesRequest).Execute()



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
	appScreenshotSetAppScreenshotsLinkagesRequest := *openapiclient.NewAppScreenshotSetAppScreenshotsLinkagesRequest([]openapiclient.AppScreenshotSetRelationshipsAppScreenshotsDataInner{*openapiclient.NewAppScreenshotSetRelationshipsAppScreenshotsDataInner("Type_example", "Id_example")}) // AppScreenshotSetAppScreenshotsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppScreenshotSetsAPI.AppScreenshotSetsAppScreenshotsReplaceToManyRelationship(context.Background(), id).AppScreenshotSetAppScreenshotsLinkagesRequest(appScreenshotSetAppScreenshotsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsAPI.AppScreenshotSetsAppScreenshotsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppScreenshotSetsAppScreenshotsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appScreenshotSetAppScreenshotsLinkagesRequest** | [**AppScreenshotSetAppScreenshotsLinkagesRequest**](AppScreenshotSetAppScreenshotsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsCreateInstance

> AppScreenshotSetResponse AppScreenshotSetsCreateInstance(ctx).AppScreenshotSetCreateRequest(appScreenshotSetCreateRequest).Execute()



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
	appScreenshotSetCreateRequest := *openapiclient.NewAppScreenshotSetCreateRequest(*openapiclient.NewAppScreenshotSetCreateRequestData("Type_example", *openapiclient.NewAppScreenshotSetCreateRequestDataAttributes(openapiclient.ScreenshotDisplayType("APP_IPHONE_67")))) // AppScreenshotSetCreateRequest | AppScreenshotSet representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppScreenshotSetsAPI.AppScreenshotSetsCreateInstance(context.Background()).AppScreenshotSetCreateRequest(appScreenshotSetCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsAPI.AppScreenshotSetsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppScreenshotSetsCreateInstance`: AppScreenshotSetResponse
	fmt.Fprintf(os.Stdout, "Response from `AppScreenshotSetsAPI.AppScreenshotSetsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appScreenshotSetCreateRequest** | [**AppScreenshotSetCreateRequest**](AppScreenshotSetCreateRequest.md) | AppScreenshotSet representation | 

### Return type

[**AppScreenshotSetResponse**](AppScreenshotSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppScreenshotSetsDeleteInstance

> AppScreenshotSetsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppScreenshotSetsAPI.AppScreenshotSetsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsAPI.AppScreenshotSetsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppScreenshotSetsDeleteInstanceRequest struct via the builder pattern


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


## AppScreenshotSetsGetInstance

> AppScreenshotSetResponse AppScreenshotSetsGetInstance(ctx, id).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppScreenshots(fieldsAppScreenshots).Include(include).LimitAppScreenshots(limitAppScreenshots).Execute()



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
	fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
	fieldsAppScreenshots := []string{"FieldsAppScreenshots_example"} // []string | the fields to include for returned resources of type appScreenshots (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppScreenshots := int32(56) // int32 | maximum number of related appScreenshots returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppScreenshotSetsAPI.AppScreenshotSetsGetInstance(context.Background(), id).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppScreenshots(fieldsAppScreenshots).Include(include).LimitAppScreenshots(limitAppScreenshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppScreenshotSetsAPI.AppScreenshotSetsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppScreenshotSetsGetInstance`: AppScreenshotSetResponse
	fmt.Fprintf(os.Stdout, "Response from `AppScreenshotSetsAPI.AppScreenshotSetsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppScreenshotSetsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppScreenshots** | **[]string** | the fields to include for returned resources of type appScreenshots | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppScreenshots** | **int32** | maximum number of related appScreenshots returned (when they are included) | 

### Return type

[**AppScreenshotSetResponse**](AppScreenshotSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

