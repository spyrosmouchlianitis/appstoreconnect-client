# \AppStoreVersionPromotionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionPromotionsCreateInstance**](AppStoreVersionPromotionsAPI.md#AppStoreVersionPromotionsCreateInstance) | **Post** /v1/appStoreVersionPromotions | 



## AppStoreVersionPromotionsCreateInstance

> AppStoreVersionPromotionResponse AppStoreVersionPromotionsCreateInstance(ctx).AppStoreVersionPromotionCreateRequest(appStoreVersionPromotionCreateRequest).Execute()



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
	appStoreVersionPromotionCreateRequest := *openapiclient.NewAppStoreVersionPromotionCreateRequest(*openapiclient.NewAppStoreVersionPromotionCreateRequestData("Type_example", *openapiclient.NewAppStoreVersionPromotionCreateRequestDataRelationships(*openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData("Type_example", "Id_example")), *openapiclient.NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationshipsAppStoreVersionExperimentTreatment(*openapiclient.NewAppStoreVersionExperimentTreatmentLocalizationRelationshipsAppStoreVersionExperimentTreatmentData("Type_example", "Id_example"))))) // AppStoreVersionPromotionCreateRequest | AppStoreVersionPromotion representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionPromotionsAPI.AppStoreVersionPromotionsCreateInstance(context.Background()).AppStoreVersionPromotionCreateRequest(appStoreVersionPromotionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionPromotionsAPI.AppStoreVersionPromotionsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionPromotionsCreateInstance`: AppStoreVersionPromotionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionPromotionsAPI.AppStoreVersionPromotionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionPromotionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreVersionPromotionCreateRequest** | [**AppStoreVersionPromotionCreateRequest**](AppStoreVersionPromotionCreateRequest.md) | AppStoreVersionPromotion representation | 

### Return type

[**AppStoreVersionPromotionResponse**](AppStoreVersionPromotionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

