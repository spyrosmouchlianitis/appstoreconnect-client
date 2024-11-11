/*
App Store Connect API

Testing InAppPurchaseAppStoreReviewScreenshotsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package appstoreconnect-client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_appstoreconnect-client_InAppPurchaseAppStoreReviewScreenshotsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InAppPurchaseAppStoreReviewScreenshotsAPIService InAppPurchaseAppStoreReviewScreenshotsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InAppPurchaseAppStoreReviewScreenshotsAPI.InAppPurchaseAppStoreReviewScreenshotsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchaseAppStoreReviewScreenshotsAPIService InAppPurchaseAppStoreReviewScreenshotsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.InAppPurchaseAppStoreReviewScreenshotsAPI.InAppPurchaseAppStoreReviewScreenshotsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchaseAppStoreReviewScreenshotsAPIService InAppPurchaseAppStoreReviewScreenshotsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchaseAppStoreReviewScreenshotsAPI.InAppPurchaseAppStoreReviewScreenshotsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchaseAppStoreReviewScreenshotsAPIService InAppPurchaseAppStoreReviewScreenshotsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchaseAppStoreReviewScreenshotsAPI.InAppPurchaseAppStoreReviewScreenshotsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}