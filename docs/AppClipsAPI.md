# \AppClipsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppClipsAppClipAdvancedExperiencesGetToManyRelated**](AppClipsAPI.md#AppClipsAppClipAdvancedExperiencesGetToManyRelated) | **Get** /v1/appClips/{id}/appClipAdvancedExperiences | 
[**AppClipsAppClipDefaultExperiencesGetToManyRelated**](AppClipsAPI.md#AppClipsAppClipDefaultExperiencesGetToManyRelated) | **Get** /v1/appClips/{id}/appClipDefaultExperiences | 
[**AppClipsGetInstance**](AppClipsAPI.md#AppClipsGetInstance) | **Get** /v1/appClips/{id} | 



## AppClipsAppClipAdvancedExperiencesGetToManyRelated

> AppClipAdvancedExperiencesResponse AppClipsAppClipAdvancedExperiencesGetToManyRelated(ctx, id).FilterStatus(filterStatus).FilterPlaceStatus(filterPlaceStatus).FilterAction(filterAction).FieldsAppClipAdvancedExperiences(fieldsAppClipAdvancedExperiences).FieldsAppClips(fieldsAppClips).FieldsAppClipAdvancedExperienceImages(fieldsAppClipAdvancedExperienceImages).FieldsAppClipAdvancedExperienceLocalizations(fieldsAppClipAdvancedExperienceLocalizations).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).Execute()



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
	filterStatus := []string{"FilterStatus_example"} // []string | filter by attribute 'status' (optional)
	filterPlaceStatus := []string{"FilterPlaceStatus_example"} // []string | filter by attribute 'placeStatus' (optional)
	filterAction := []string{"FilterAction_example"} // []string | filter by attribute 'action' (optional)
	fieldsAppClipAdvancedExperiences := []string{"FieldsAppClipAdvancedExperiences_example"} // []string | the fields to include for returned resources of type appClipAdvancedExperiences (optional)
	fieldsAppClips := []string{"FieldsAppClips_example"} // []string | the fields to include for returned resources of type appClips (optional)
	fieldsAppClipAdvancedExperienceImages := []string{"FieldsAppClipAdvancedExperienceImages_example"} // []string | the fields to include for returned resources of type appClipAdvancedExperienceImages (optional)
	fieldsAppClipAdvancedExperienceLocalizations := []string{"FieldsAppClipAdvancedExperienceLocalizations_example"} // []string | the fields to include for returned resources of type appClipAdvancedExperienceLocalizations (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppClipsAPI.AppClipsAppClipAdvancedExperiencesGetToManyRelated(context.Background(), id).FilterStatus(filterStatus).FilterPlaceStatus(filterPlaceStatus).FilterAction(filterAction).FieldsAppClipAdvancedExperiences(fieldsAppClipAdvancedExperiences).FieldsAppClips(fieldsAppClips).FieldsAppClipAdvancedExperienceImages(fieldsAppClipAdvancedExperienceImages).FieldsAppClipAdvancedExperienceLocalizations(fieldsAppClipAdvancedExperienceLocalizations).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppClipsAPI.AppClipsAppClipAdvancedExperiencesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppClipsAppClipAdvancedExperiencesGetToManyRelated`: AppClipAdvancedExperiencesResponse
	fmt.Fprintf(os.Stdout, "Response from `AppClipsAPI.AppClipsAppClipAdvancedExperiencesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipsAppClipAdvancedExperiencesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterStatus** | **[]string** | filter by attribute &#39;status&#39; | 
 **filterPlaceStatus** | **[]string** | filter by attribute &#39;placeStatus&#39; | 
 **filterAction** | **[]string** | filter by attribute &#39;action&#39; | 
 **fieldsAppClipAdvancedExperiences** | **[]string** | the fields to include for returned resources of type appClipAdvancedExperiences | 
 **fieldsAppClips** | **[]string** | the fields to include for returned resources of type appClips | 
 **fieldsAppClipAdvancedExperienceImages** | **[]string** | the fields to include for returned resources of type appClipAdvancedExperienceImages | 
 **fieldsAppClipAdvancedExperienceLocalizations** | **[]string** | the fields to include for returned resources of type appClipAdvancedExperienceLocalizations | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 

### Return type

[**AppClipAdvancedExperiencesResponse**](AppClipAdvancedExperiencesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipsAppClipDefaultExperiencesGetToManyRelated

> AppClipDefaultExperiencesResponse AppClipsAppClipDefaultExperiencesGetToManyRelated(ctx, id).ExistsReleaseWithAppStoreVersion(existsReleaseWithAppStoreVersion).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppClips(fieldsAppClips).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations).FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails).Limit(limit).Include(include).LimitAppClipDefaultExperienceLocalizations(limitAppClipDefaultExperienceLocalizations).Execute()



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
	existsReleaseWithAppStoreVersion := true // bool | filter by existence or non-existence of related 'releaseWithAppStoreVersion' (optional)
	fieldsAppClipDefaultExperiences := []string{"FieldsAppClipDefaultExperiences_example"} // []string | the fields to include for returned resources of type appClipDefaultExperiences (optional)
	fieldsAppClips := []string{"FieldsAppClips_example"} // []string | the fields to include for returned resources of type appClips (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsAppClipDefaultExperienceLocalizations := []string{"FieldsAppClipDefaultExperienceLocalizations_example"} // []string | the fields to include for returned resources of type appClipDefaultExperienceLocalizations (optional)
	fieldsAppClipAppStoreReviewDetails := []string{"FieldsAppClipAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appClipAppStoreReviewDetails (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppClipDefaultExperienceLocalizations := int32(56) // int32 | maximum number of related appClipDefaultExperienceLocalizations returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppClipsAPI.AppClipsAppClipDefaultExperiencesGetToManyRelated(context.Background(), id).ExistsReleaseWithAppStoreVersion(existsReleaseWithAppStoreVersion).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppClips(fieldsAppClips).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations).FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails).Limit(limit).Include(include).LimitAppClipDefaultExperienceLocalizations(limitAppClipDefaultExperienceLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppClipsAPI.AppClipsAppClipDefaultExperiencesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppClipsAppClipDefaultExperiencesGetToManyRelated`: AppClipDefaultExperiencesResponse
	fmt.Fprintf(os.Stdout, "Response from `AppClipsAPI.AppClipsAppClipDefaultExperiencesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipsAppClipDefaultExperiencesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **existsReleaseWithAppStoreVersion** | **bool** | filter by existence or non-existence of related &#39;releaseWithAppStoreVersion&#39; | 
 **fieldsAppClipDefaultExperiences** | **[]string** | the fields to include for returned resources of type appClipDefaultExperiences | 
 **fieldsAppClips** | **[]string** | the fields to include for returned resources of type appClips | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppClipDefaultExperienceLocalizations** | **[]string** | the fields to include for returned resources of type appClipDefaultExperienceLocalizations | 
 **fieldsAppClipAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appClipAppStoreReviewDetails | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppClipDefaultExperienceLocalizations** | **int32** | maximum number of related appClipDefaultExperienceLocalizations returned (when they are included) | 

### Return type

[**AppClipDefaultExperiencesResponse**](AppClipDefaultExperiencesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipsGetInstance

> AppClipResponse AppClipsGetInstance(ctx, id).FieldsAppClips(fieldsAppClips).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).Include(include).LimitAppClipDefaultExperiences(limitAppClipDefaultExperiences).Execute()



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
	fieldsAppClips := []string{"FieldsAppClips_example"} // []string | the fields to include for returned resources of type appClips (optional)
	fieldsAppClipDefaultExperiences := []string{"FieldsAppClipDefaultExperiences_example"} // []string | the fields to include for returned resources of type appClipDefaultExperiences (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppClipDefaultExperiences := int32(56) // int32 | maximum number of related appClipDefaultExperiences returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppClipsAPI.AppClipsGetInstance(context.Background(), id).FieldsAppClips(fieldsAppClips).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).Include(include).LimitAppClipDefaultExperiences(limitAppClipDefaultExperiences).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppClipsAPI.AppClipsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppClipsGetInstance`: AppClipResponse
	fmt.Fprintf(os.Stdout, "Response from `AppClipsAPI.AppClipsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppClips** | **[]string** | the fields to include for returned resources of type appClips | 
 **fieldsAppClipDefaultExperiences** | **[]string** | the fields to include for returned resources of type appClipDefaultExperiences | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppClipDefaultExperiences** | **int32** | maximum number of related appClipDefaultExperiences returned (when they are included) | 

### Return type

[**AppClipResponse**](AppClipResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

