/*
App Store Connect API

Testing GameCenterLeaderboardImagesAPIService

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

func Test_appstoreconnect-client_GameCenterLeaderboardImagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GameCenterLeaderboardImagesAPIService GameCenterLeaderboardImagesCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardImagesAPIService GameCenterLeaderboardImagesDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardImagesAPIService GameCenterLeaderboardImagesGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardImagesAPIService GameCenterLeaderboardImagesUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
