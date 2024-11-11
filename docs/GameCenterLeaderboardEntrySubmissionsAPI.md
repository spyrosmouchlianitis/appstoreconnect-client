# \GameCenterLeaderboardEntrySubmissionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardEntrySubmissionsCreateInstance**](GameCenterLeaderboardEntrySubmissionsAPI.md#GameCenterLeaderboardEntrySubmissionsCreateInstance) | **Post** /v1/gameCenterLeaderboardEntrySubmissions | 



## GameCenterLeaderboardEntrySubmissionsCreateInstance

> GameCenterLeaderboardEntrySubmissionResponse GameCenterLeaderboardEntrySubmissionsCreateInstance(ctx).GameCenterLeaderboardEntrySubmissionCreateRequest(gameCenterLeaderboardEntrySubmissionCreateRequest).Execute()



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
	gameCenterLeaderboardEntrySubmissionCreateRequest := *openapiclient.NewGameCenterLeaderboardEntrySubmissionCreateRequest(*openapiclient.NewGameCenterLeaderboardEntrySubmissionCreateRequestData("Type_example", *openapiclient.NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes("BundleId_example", "ScopedPlayerId_example", float64(123), "VendorIdentifier_example"))) // GameCenterLeaderboardEntrySubmissionCreateRequest | GameCenterLeaderboardEntrySubmission representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterLeaderboardEntrySubmissionsAPI.GameCenterLeaderboardEntrySubmissionsCreateInstance(context.Background()).GameCenterLeaderboardEntrySubmissionCreateRequest(gameCenterLeaderboardEntrySubmissionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardEntrySubmissionsAPI.GameCenterLeaderboardEntrySubmissionsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterLeaderboardEntrySubmissionsCreateInstance`: GameCenterLeaderboardEntrySubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardEntrySubmissionsAPI.GameCenterLeaderboardEntrySubmissionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardEntrySubmissionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardEntrySubmissionCreateRequest** | [**GameCenterLeaderboardEntrySubmissionCreateRequest**](GameCenterLeaderboardEntrySubmissionCreateRequest.md) | GameCenterLeaderboardEntrySubmission representation | 

### Return type

[**GameCenterLeaderboardEntrySubmissionResponse**](GameCenterLeaderboardEntrySubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

