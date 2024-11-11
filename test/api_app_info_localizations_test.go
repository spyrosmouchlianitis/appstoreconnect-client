/*
App Store Connect API

Testing AppInfoLocalizationsAPIService

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

func Test_appstoreconnect-client_AppInfoLocalizationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppInfoLocalizationsAPIService AppInfoLocalizationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppInfoLocalizationsAPI.AppInfoLocalizationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppInfoLocalizationsAPIService AppInfoLocalizationsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AppInfoLocalizationsAPI.AppInfoLocalizationsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppInfoLocalizationsAPIService AppInfoLocalizationsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppInfoLocalizationsAPI.AppInfoLocalizationsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppInfoLocalizationsAPIService AppInfoLocalizationsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppInfoLocalizationsAPI.AppInfoLocalizationsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
