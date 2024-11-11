# \AppInfoLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppInfoLocalizationsCreateInstance**](AppInfoLocalizationsAPI.md#AppInfoLocalizationsCreateInstance) | **Post** /v1/appInfoLocalizations | 
[**AppInfoLocalizationsDeleteInstance**](AppInfoLocalizationsAPI.md#AppInfoLocalizationsDeleteInstance) | **Delete** /v1/appInfoLocalizations/{id} | 
[**AppInfoLocalizationsGetInstance**](AppInfoLocalizationsAPI.md#AppInfoLocalizationsGetInstance) | **Get** /v1/appInfoLocalizations/{id} | 
[**AppInfoLocalizationsUpdateInstance**](AppInfoLocalizationsAPI.md#AppInfoLocalizationsUpdateInstance) | **Patch** /v1/appInfoLocalizations/{id} | 



## AppInfoLocalizationsCreateInstance

> AppInfoLocalizationResponse AppInfoLocalizationsCreateInstance(ctx).AppInfoLocalizationCreateRequest(appInfoLocalizationCreateRequest).Execute()



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
	appInfoLocalizationCreateRequest := *openapiclient.NewAppInfoLocalizationCreateRequest(*openapiclient.NewAppInfoLocalizationCreateRequestData("Type_example", *openapiclient.NewAppInfoLocalizationCreateRequestDataAttributes("Locale_example"), *openapiclient.NewAppInfoLocalizationCreateRequestDataRelationships(*openapiclient.NewAppInfoLocalizationCreateRequestDataRelationshipsAppInfo(*openapiclient.NewAppInfoLocalizationRelationshipsAppInfoData("Type_example", "Id_example"))))) // AppInfoLocalizationCreateRequest | AppInfoLocalization representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInfoLocalizationsAPI.AppInfoLocalizationsCreateInstance(context.Background()).AppInfoLocalizationCreateRequest(appInfoLocalizationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInfoLocalizationsAPI.AppInfoLocalizationsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInfoLocalizationsCreateInstance`: AppInfoLocalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AppInfoLocalizationsAPI.AppInfoLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppInfoLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appInfoLocalizationCreateRequest** | [**AppInfoLocalizationCreateRequest**](AppInfoLocalizationCreateRequest.md) | AppInfoLocalization representation | 

### Return type

[**AppInfoLocalizationResponse**](AppInfoLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfoLocalizationsDeleteInstance

> AppInfoLocalizationsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppInfoLocalizationsAPI.AppInfoLocalizationsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInfoLocalizationsAPI.AppInfoLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppInfoLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## AppInfoLocalizationsGetInstance

> AppInfoLocalizationResponse AppInfoLocalizationsGetInstance(ctx, id).FieldsAppInfoLocalizations(fieldsAppInfoLocalizations).Include(include).Execute()



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
	fieldsAppInfoLocalizations := []string{"FieldsAppInfoLocalizations_example"} // []string | the fields to include for returned resources of type appInfoLocalizations (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInfoLocalizationsAPI.AppInfoLocalizationsGetInstance(context.Background(), id).FieldsAppInfoLocalizations(fieldsAppInfoLocalizations).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInfoLocalizationsAPI.AppInfoLocalizationsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInfoLocalizationsGetInstance`: AppInfoLocalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AppInfoLocalizationsAPI.AppInfoLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppInfoLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppInfoLocalizations** | **[]string** | the fields to include for returned resources of type appInfoLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppInfoLocalizationResponse**](AppInfoLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppInfoLocalizationsUpdateInstance

> AppInfoLocalizationResponse AppInfoLocalizationsUpdateInstance(ctx, id).AppInfoLocalizationUpdateRequest(appInfoLocalizationUpdateRequest).Execute()



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
	appInfoLocalizationUpdateRequest := *openapiclient.NewAppInfoLocalizationUpdateRequest(*openapiclient.NewAppInfoLocalizationUpdateRequestData("Type_example", "Id_example")) // AppInfoLocalizationUpdateRequest | AppInfoLocalization representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppInfoLocalizationsAPI.AppInfoLocalizationsUpdateInstance(context.Background(), id).AppInfoLocalizationUpdateRequest(appInfoLocalizationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppInfoLocalizationsAPI.AppInfoLocalizationsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppInfoLocalizationsUpdateInstance`: AppInfoLocalizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AppInfoLocalizationsAPI.AppInfoLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppInfoLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appInfoLocalizationUpdateRequest** | [**AppInfoLocalizationUpdateRequest**](AppInfoLocalizationUpdateRequest.md) | AppInfoLocalization representation | 

### Return type

[**AppInfoLocalizationResponse**](AppInfoLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

