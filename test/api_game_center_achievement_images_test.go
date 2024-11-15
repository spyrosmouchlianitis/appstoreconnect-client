/*
App Store Connect API

Testing GameCenterAchievementImagesAPIService

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

func Test_appstoreconnectclient_GameCenterAchievementImagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GameCenterAchievementImagesAPIService GameCenterAchievementImagesCreateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GameCenterAchievementImagesAPI.GameCenterAchievementImagesCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterAchievementImagesAPIService GameCenterAchievementImagesDeleteInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterAchievementImagesAPI.GameCenterAchievementImagesDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterAchievementImagesAPIService GameCenterAchievementImagesGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementImagesAPI.GameCenterAchievementImagesGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterAchievementImagesAPIService GameCenterAchievementImagesUpdateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementImagesAPI.GameCenterAchievementImagesUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
