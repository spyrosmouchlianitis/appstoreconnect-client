# \AppStoreVersionReleaseRequestsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionReleaseRequestsCreateInstance**](AppStoreVersionReleaseRequestsAPI.md#AppStoreVersionReleaseRequestsCreateInstance) | **Post** /v1/appStoreVersionReleaseRequests | 



## AppStoreVersionReleaseRequestsCreateInstance

> AppStoreVersionReleaseRequestResponse AppStoreVersionReleaseRequestsCreateInstance(ctx).AppStoreVersionReleaseRequestCreateRequest(appStoreVersionReleaseRequestCreateRequest).Execute()



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
	appStoreVersionReleaseRequestCreateRequest := *openapiclient.NewAppStoreVersionReleaseRequestCreateRequest(*openapiclient.NewAppStoreVersionReleaseRequestCreateRequestData("Type_example", *openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationships(*openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData("Type_example", "Id_example"))))) // AppStoreVersionReleaseRequestCreateRequest | AppStoreVersionReleaseRequest representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionReleaseRequestsAPI.AppStoreVersionReleaseRequestsCreateInstance(context.Background()).AppStoreVersionReleaseRequestCreateRequest(appStoreVersionReleaseRequestCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionReleaseRequestsAPI.AppStoreVersionReleaseRequestsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionReleaseRequestsCreateInstance`: AppStoreVersionReleaseRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionReleaseRequestsAPI.AppStoreVersionReleaseRequestsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionReleaseRequestsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreVersionReleaseRequestCreateRequest** | [**AppStoreVersionReleaseRequestCreateRequest**](AppStoreVersionReleaseRequestCreateRequest.md) | AppStoreVersionReleaseRequest representation | 

### Return type

[**AppStoreVersionReleaseRequestResponse**](AppStoreVersionReleaseRequestResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

