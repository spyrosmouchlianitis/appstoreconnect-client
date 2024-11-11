/*
App Store Connect API

Testing AppClipDefaultExperienceLocalizationsAPIService

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

func Test_appstoreconnect-client_AppClipDefaultExperienceLocalizationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppClipDefaultExperienceLocalizationsAPIService AppClipDefaultExperienceLocalizationsAppClipHeaderImageGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppClipDefaultExperienceLocalizationsAPI.AppClipDefaultExperienceLocalizationsAppClipHeaderImageGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppClipDefaultExperienceLocalizationsAPIService AppClipDefaultExperienceLocalizationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppClipDefaultExperienceLocalizationsAPI.AppClipDefaultExperienceLocalizationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppClipDefaultExperienceLocalizationsAPIService AppClipDefaultExperienceLocalizationsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AppClipDefaultExperienceLocalizationsAPI.AppClipDefaultExperienceLocalizationsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppClipDefaultExperienceLocalizationsAPIService AppClipDefaultExperienceLocalizationsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppClipDefaultExperienceLocalizationsAPI.AppClipDefaultExperienceLocalizationsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppClipDefaultExperienceLocalizationsAPIService AppClipDefaultExperienceLocalizationsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppClipDefaultExperienceLocalizationsAPI.AppClipDefaultExperienceLocalizationsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
