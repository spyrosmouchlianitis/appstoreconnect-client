# \AppEventsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppEventsCreateInstance**](AppEventsAPI.md#AppEventsCreateInstance) | **Post** /v1/appEvents | 
[**AppEventsDeleteInstance**](AppEventsAPI.md#AppEventsDeleteInstance) | **Delete** /v1/appEvents/{id} | 
[**AppEventsGetInstance**](AppEventsAPI.md#AppEventsGetInstance) | **Get** /v1/appEvents/{id} | 
[**AppEventsLocalizationsGetToManyRelated**](AppEventsAPI.md#AppEventsLocalizationsGetToManyRelated) | **Get** /v1/appEvents/{id}/localizations | 
[**AppEventsUpdateInstance**](AppEventsAPI.md#AppEventsUpdateInstance) | **Patch** /v1/appEvents/{id} | 



## AppEventsCreateInstance

> AppEventResponse AppEventsCreateInstance(ctx).AppEventCreateRequest(appEventCreateRequest).Execute()



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
	appEventCreateRequest := *openapiclient.NewAppEventCreateRequest(*openapiclient.NewAppEventCreateRequestData("Type_example", *openapiclient.NewAppEventCreateRequestDataAttributes("ReferenceName_example"), *openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // AppEventCreateRequest | AppEvent representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventsCreateInstance(context.Background()).AppEventCreateRequest(appEventCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventsCreateInstance`: AppEventResponse
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEventsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appEventCreateRequest** | [**AppEventCreateRequest**](AppEventCreateRequest.md) | AppEvent representation | 

### Return type

[**AppEventResponse**](AppEventResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventsDeleteInstance

> AppEventsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppEventsAPI.AppEventsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppEventsDeleteInstanceRequest struct via the builder pattern


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


## AppEventsGetInstance

> AppEventResponse AppEventsGetInstance(ctx, id).FieldsAppEvents(fieldsAppEvents).FieldsAppEventLocalizations(fieldsAppEventLocalizations).Include(include).LimitLocalizations(limitLocalizations).Execute()



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
	fieldsAppEvents := []string{"FieldsAppEvents_example"} // []string | the fields to include for returned resources of type appEvents (optional)
	fieldsAppEventLocalizations := []string{"FieldsAppEventLocalizations_example"} // []string | the fields to include for returned resources of type appEventLocalizations (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventsGetInstance(context.Background(), id).FieldsAppEvents(fieldsAppEvents).FieldsAppEventLocalizations(fieldsAppEventLocalizations).Include(include).LimitLocalizations(limitLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventsGetInstance`: AppEventResponse
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEvents** | **[]string** | the fields to include for returned resources of type appEvents | 
 **fieldsAppEventLocalizations** | **[]string** | the fields to include for returned resources of type appEventLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 

### Return type

[**AppEventResponse**](AppEventResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventsLocalizationsGetToManyRelated

> AppEventLocalizationsResponse AppEventsLocalizationsGetToManyRelated(ctx, id).FieldsAppEventLocalizations(fieldsAppEventLocalizations).FieldsAppEvents(fieldsAppEvents).FieldsAppEventScreenshots(fieldsAppEventScreenshots).FieldsAppEventVideoClips(fieldsAppEventVideoClips).Limit(limit).Include(include).LimitAppEventScreenshots(limitAppEventScreenshots).LimitAppEventVideoClips(limitAppEventVideoClips).Execute()



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
	fieldsAppEventLocalizations := []string{"FieldsAppEventLocalizations_example"} // []string | the fields to include for returned resources of type appEventLocalizations (optional)
	fieldsAppEvents := []string{"FieldsAppEvents_example"} // []string | the fields to include for returned resources of type appEvents (optional)
	fieldsAppEventScreenshots := []string{"FieldsAppEventScreenshots_example"} // []string | the fields to include for returned resources of type appEventScreenshots (optional)
	fieldsAppEventVideoClips := []string{"FieldsAppEventVideoClips_example"} // []string | the fields to include for returned resources of type appEventVideoClips (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppEventScreenshots := int32(56) // int32 | maximum number of related appEventScreenshots returned (when they are included) (optional)
	limitAppEventVideoClips := int32(56) // int32 | maximum number of related appEventVideoClips returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventsLocalizationsGetToManyRelated(context.Background(), id).FieldsAppEventLocalizations(fieldsAppEventLocalizations).FieldsAppEvents(fieldsAppEvents).FieldsAppEventScreenshots(fieldsAppEventScreenshots).FieldsAppEventVideoClips(fieldsAppEventVideoClips).Limit(limit).Include(include).LimitAppEventScreenshots(limitAppEventScreenshots).LimitAppEventVideoClips(limitAppEventVideoClips).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventsLocalizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventsLocalizationsGetToManyRelated`: AppEventLocalizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventsLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventsLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEventLocalizations** | **[]string** | the fields to include for returned resources of type appEventLocalizations | 
 **fieldsAppEvents** | **[]string** | the fields to include for returned resources of type appEvents | 
 **fieldsAppEventScreenshots** | **[]string** | the fields to include for returned resources of type appEventScreenshots | 
 **fieldsAppEventVideoClips** | **[]string** | the fields to include for returned resources of type appEventVideoClips | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppEventScreenshots** | **int32** | maximum number of related appEventScreenshots returned (when they are included) | 
 **limitAppEventVideoClips** | **int32** | maximum number of related appEventVideoClips returned (when they are included) | 

### Return type

[**AppEventLocalizationsResponse**](AppEventLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEventsUpdateInstance

> AppEventResponse AppEventsUpdateInstance(ctx, id).AppEventUpdateRequest(appEventUpdateRequest).Execute()



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
	appEventUpdateRequest := *openapiclient.NewAppEventUpdateRequest(*openapiclient.NewAppEventUpdateRequestData("Type_example", "Id_example")) // AppEventUpdateRequest | AppEvent representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppEventsAPI.AppEventsUpdateInstance(context.Background(), id).AppEventUpdateRequest(appEventUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppEventsAPI.AppEventsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppEventsUpdateInstance`: AppEventResponse
	fmt.Fprintf(os.Stdout, "Response from `AppEventsAPI.AppEventsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEventsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appEventUpdateRequest** | [**AppEventUpdateRequest**](AppEventUpdateRequest.md) | AppEvent representation | 

### Return type

[**AppEventResponse**](AppEventResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

