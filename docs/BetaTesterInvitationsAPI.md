# \BetaTesterInvitationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaTesterInvitationsCreateInstance**](BetaTesterInvitationsAPI.md#BetaTesterInvitationsCreateInstance) | **Post** /v1/betaTesterInvitations | 



## BetaTesterInvitationsCreateInstance

> BetaTesterInvitationResponse BetaTesterInvitationsCreateInstance(ctx).BetaTesterInvitationCreateRequest(betaTesterInvitationCreateRequest).Execute()



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
	betaTesterInvitationCreateRequest := *openapiclient.NewBetaTesterInvitationCreateRequest(*openapiclient.NewBetaTesterInvitationCreateRequestData("Type_example", *openapiclient.NewBetaTesterInvitationCreateRequestDataRelationships(*openapiclient.NewBetaTesterInvitationCreateRequestDataRelationshipsBetaTester(*openapiclient.NewBetaGroupRelationshipsBetaTestersDataInner("Type_example", "Id_example")), *openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // BetaTesterInvitationCreateRequest | BetaTesterInvitation representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaTesterInvitationsAPI.BetaTesterInvitationsCreateInstance(context.Background()).BetaTesterInvitationCreateRequest(betaTesterInvitationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTesterInvitationsAPI.BetaTesterInvitationsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTesterInvitationsCreateInstance`: BetaTesterInvitationResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTesterInvitationsAPI.BetaTesterInvitationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaTesterInvitationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaTesterInvitationCreateRequest** | [**BetaTesterInvitationCreateRequest**](BetaTesterInvitationCreateRequest.md) | BetaTesterInvitation representation | 

### Return type

[**BetaTesterInvitationResponse**](BetaTesterInvitationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

