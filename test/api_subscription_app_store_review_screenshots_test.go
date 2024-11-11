/*
App Store Connect API

Testing SubscriptionAppStoreReviewScreenshotsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package appstoreconnectclient_test

import (
	"context"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_appstoreconnectclient_SubscriptionAppStoreReviewScreenshotsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubscriptionAppStoreReviewScreenshotsAPIService SubscriptionAppStoreReviewScreenshotsCreateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAppStoreReviewScreenshotsAPIService SubscriptionAppStoreReviewScreenshotsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAppStoreReviewScreenshotsAPIService SubscriptionAppStoreReviewScreenshotsGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionAppStoreReviewScreenshotsAPIService SubscriptionAppStoreReviewScreenshotsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.SubscriptionAppStoreReviewScreenshotsAPI.SubscriptionAppStoreReviewScreenshotsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
