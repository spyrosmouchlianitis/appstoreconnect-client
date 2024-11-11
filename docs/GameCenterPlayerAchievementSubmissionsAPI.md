# \GameCenterPlayerAchievementSubmissionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterPlayerAchievementSubmissionsCreateInstance**](GameCenterPlayerAchievementSubmissionsAPI.md#GameCenterPlayerAchievementSubmissionsCreateInstance) | **Post** /v1/gameCenterPlayerAchievementSubmissions | 



## GameCenterPlayerAchievementSubmissionsCreateInstance

> GameCenterPlayerAchievementSubmissionResponse GameCenterPlayerAchievementSubmissionsCreateInstance(ctx).GameCenterPlayerAchievementSubmissionCreateRequest(gameCenterPlayerAchievementSubmissionCreateRequest).Execute()



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
	gameCenterPlayerAchievementSubmissionCreateRequest := *openapiclient.NewGameCenterPlayerAchievementSubmissionCreateRequest(*openapiclient.NewGameCenterPlayerAchievementSubmissionCreateRequestData("Type_example", *openapiclient.NewGameCenterPlayerAchievementSubmissionCreateRequestDataAttributes("BundleId_example", int32(123), "ScopedPlayerId_example", "VendorIdentifier_example"))) // GameCenterPlayerAchievementSubmissionCreateRequest | GameCenterPlayerAchievementSubmission representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterPlayerAchievementSubmissionsAPI.GameCenterPlayerAchievementSubmissionsCreateInstance(context.Background()).GameCenterPlayerAchievementSubmissionCreateRequest(gameCenterPlayerAchievementSubmissionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterPlayerAchievementSubmissionsAPI.GameCenterPlayerAchievementSubmissionsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterPlayerAchievementSubmissionsCreateInstance`: GameCenterPlayerAchievementSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterPlayerAchievementSubmissionsAPI.GameCenterPlayerAchievementSubmissionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterPlayerAchievementSubmissionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterPlayerAchievementSubmissionCreateRequest** | [**GameCenterPlayerAchievementSubmissionCreateRequest**](GameCenterPlayerAchievementSubmissionCreateRequest.md) | GameCenterPlayerAchievementSubmission representation | 

### Return type

[**GameCenterPlayerAchievementSubmissionResponse**](GameCenterPlayerAchievementSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

