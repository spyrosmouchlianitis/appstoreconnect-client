/*
App Store Connect API

Testing GameCenterLeaderboardSetsAPIService

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

func Test_appstoreconnect-client_GameCenterLeaderboardSetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationship", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationship", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationship", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationship", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsLocalizationsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsLocalizationsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsReleasesGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsReleasesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterLeaderboardSetsAPIService GameCenterLeaderboardSetsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
