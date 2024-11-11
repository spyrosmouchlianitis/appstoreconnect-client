/*
App Store Connect API

Testing InAppPurchasePriceSchedulesAPIService

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

func Test_appstoreconnect-client_InAppPurchasePriceSchedulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InAppPurchasePriceSchedulesAPIService InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchasePriceSchedulesAPIService InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchasePriceSchedulesAPIService InAppPurchasePriceSchedulesCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchasePriceSchedulesAPIService InAppPurchasePriceSchedulesGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchasePriceSchedulesAPIService InAppPurchasePriceSchedulesManualPricesGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesAPI.InAppPurchasePriceSchedulesManualPricesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
