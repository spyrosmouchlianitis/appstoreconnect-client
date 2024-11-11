# \AppCustomProductPagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated**](AppCustomProductPagesAPI.md#AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated) | **Get** /v1/appCustomProductPages/{id}/appCustomProductPageVersions | 
[**AppCustomProductPagesCreateInstance**](AppCustomProductPagesAPI.md#AppCustomProductPagesCreateInstance) | **Post** /v1/appCustomProductPages | 
[**AppCustomProductPagesDeleteInstance**](AppCustomProductPagesAPI.md#AppCustomProductPagesDeleteInstance) | **Delete** /v1/appCustomProductPages/{id} | 
[**AppCustomProductPagesGetInstance**](AppCustomProductPagesAPI.md#AppCustomProductPagesGetInstance) | **Get** /v1/appCustomProductPages/{id} | 
[**AppCustomProductPagesUpdateInstance**](AppCustomProductPagesAPI.md#AppCustomProductPagesUpdateInstance) | **Patch** /v1/appCustomProductPages/{id} | 



## AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated

> AppCustomProductPageVersionsResponse AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated(ctx, id).FilterState(filterState).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).FieldsAppCustomProductPages(fieldsAppCustomProductPages).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).Limit(limit).Include(include).LimitAppCustomProductPageLocalizations(limitAppCustomProductPageLocalizations).Execute()



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
	filterState := []string{"FilterState_example"} // []string | filter by attribute 'state' (optional)
	fieldsAppCustomProductPageVersions := []string{"FieldsAppCustomProductPageVersions_example"} // []string | the fields to include for returned resources of type appCustomProductPageVersions (optional)
	fieldsAppCustomProductPages := []string{"FieldsAppCustomProductPages_example"} // []string | the fields to include for returned resources of type appCustomProductPages (optional)
	fieldsAppCustomProductPageLocalizations := []string{"FieldsAppCustomProductPageLocalizations_example"} // []string | the fields to include for returned resources of type appCustomProductPageLocalizations (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppCustomProductPageLocalizations := int32(56) // int32 | maximum number of related appCustomProductPageLocalizations returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPagesAPI.AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated(context.Background(), id).FilterState(filterState).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).FieldsAppCustomProductPages(fieldsAppCustomProductPages).FieldsAppCustomProductPageLocalizations(fieldsAppCustomProductPageLocalizations).Limit(limit).Include(include).LimitAppCustomProductPageLocalizations(limitAppCustomProductPageLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPagesAPI.AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated`: AppCustomProductPageVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPagesAPI.AppCustomProductPagesAppCustomProductPageVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPagesAppCustomProductPageVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterState** | **[]string** | filter by attribute &#39;state&#39; | 
 **fieldsAppCustomProductPageVersions** | **[]string** | the fields to include for returned resources of type appCustomProductPageVersions | 
 **fieldsAppCustomProductPages** | **[]string** | the fields to include for returned resources of type appCustomProductPages | 
 **fieldsAppCustomProductPageLocalizations** | **[]string** | the fields to include for returned resources of type appCustomProductPageLocalizations | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppCustomProductPageLocalizations** | **int32** | maximum number of related appCustomProductPageLocalizations returned (when they are included) | 

### Return type

[**AppCustomProductPageVersionsResponse**](AppCustomProductPageVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomProductPagesCreateInstance

> AppCustomProductPageResponse AppCustomProductPagesCreateInstance(ctx).AppCustomProductPageCreateRequest(appCustomProductPageCreateRequest).Execute()



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
	appCustomProductPageCreateRequest := *openapiclient.NewAppCustomProductPageCreateRequest(*openapiclient.NewAppCustomProductPageCreateRequestData("Type_example", *openapiclient.NewAppCustomProductPageCreateRequestDataAttributes("Name_example"), *openapiclient.NewAppCustomProductPageCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // AppCustomProductPageCreateRequest | AppCustomProductPage representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPagesAPI.AppCustomProductPagesCreateInstance(context.Background()).AppCustomProductPageCreateRequest(appCustomProductPageCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPagesAPI.AppCustomProductPagesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPagesCreateInstance`: AppCustomProductPageResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPagesAPI.AppCustomProductPagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appCustomProductPageCreateRequest** | [**AppCustomProductPageCreateRequest**](AppCustomProductPageCreateRequest.md) | AppCustomProductPage representation | 

### Return type

[**AppCustomProductPageResponse**](AppCustomProductPageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomProductPagesDeleteInstance

> AppCustomProductPagesDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppCustomProductPagesAPI.AppCustomProductPagesDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPagesAPI.AppCustomProductPagesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppCustomProductPagesDeleteInstanceRequest struct via the builder pattern


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


## AppCustomProductPagesGetInstance

> AppCustomProductPageResponse AppCustomProductPagesGetInstance(ctx, id).FieldsAppCustomProductPages(fieldsAppCustomProductPages).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).Include(include).LimitAppCustomProductPageVersions(limitAppCustomProductPageVersions).Execute()



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
	fieldsAppCustomProductPages := []string{"FieldsAppCustomProductPages_example"} // []string | the fields to include for returned resources of type appCustomProductPages (optional)
	fieldsAppCustomProductPageVersions := []string{"FieldsAppCustomProductPageVersions_example"} // []string | the fields to include for returned resources of type appCustomProductPageVersions (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppCustomProductPageVersions := int32(56) // int32 | maximum number of related appCustomProductPageVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPagesAPI.AppCustomProductPagesGetInstance(context.Background(), id).FieldsAppCustomProductPages(fieldsAppCustomProductPages).FieldsAppCustomProductPageVersions(fieldsAppCustomProductPageVersions).Include(include).LimitAppCustomProductPageVersions(limitAppCustomProductPageVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPagesAPI.AppCustomProductPagesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPagesGetInstance`: AppCustomProductPageResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPagesAPI.AppCustomProductPagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppCustomProductPages** | **[]string** | the fields to include for returned resources of type appCustomProductPages | 
 **fieldsAppCustomProductPageVersions** | **[]string** | the fields to include for returned resources of type appCustomProductPageVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppCustomProductPageVersions** | **int32** | maximum number of related appCustomProductPageVersions returned (when they are included) | 

### Return type

[**AppCustomProductPageResponse**](AppCustomProductPageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCustomProductPagesUpdateInstance

> AppCustomProductPageResponse AppCustomProductPagesUpdateInstance(ctx, id).AppCustomProductPageUpdateRequest(appCustomProductPageUpdateRequest).Execute()



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
	appCustomProductPageUpdateRequest := *openapiclient.NewAppCustomProductPageUpdateRequest(*openapiclient.NewAppCustomProductPageUpdateRequestData("Type_example", "Id_example")) // AppCustomProductPageUpdateRequest | AppCustomProductPage representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppCustomProductPagesAPI.AppCustomProductPagesUpdateInstance(context.Background(), id).AppCustomProductPageUpdateRequest(appCustomProductPageUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppCustomProductPagesAPI.AppCustomProductPagesUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCustomProductPagesUpdateInstance`: AppCustomProductPageResponse
	fmt.Fprintf(os.Stdout, "Response from `AppCustomProductPagesAPI.AppCustomProductPagesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCustomProductPagesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appCustomProductPageUpdateRequest** | [**AppCustomProductPageUpdateRequest**](AppCustomProductPageUpdateRequest.md) | AppCustomProductPage representation | 

### Return type

[**AppCustomProductPageResponse**](AppCustomProductPageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

