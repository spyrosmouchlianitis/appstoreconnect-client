/*
App Store Connect API

Testing BetaAppLocalizationsAPIService

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

func Test_appstoreconnect-client_BetaAppLocalizationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BetaAppLocalizationsAPIService BetaAppLocalizationsAppGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsAppGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaAppLocalizationsAPIService BetaAppLocalizationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaAppLocalizationsAPIService BetaAppLocalizationsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaAppLocalizationsAPIService BetaAppLocalizationsGetCollection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaAppLocalizationsAPIService BetaAppLocalizationsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaAppLocalizationsAPIService BetaAppLocalizationsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
