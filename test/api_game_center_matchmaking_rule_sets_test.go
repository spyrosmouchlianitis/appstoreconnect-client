/*
App Store Connect API

Testing GameCenterMatchmakingRuleSetsAPIService

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

func Test_appstoreconnectclient_GameCenterMatchmakingRuleSetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GameCenterMatchmakingRuleSetsAPIService GameCenterMatchmakingRuleSetsCreateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRuleSetsAPIService GameCenterMatchmakingRuleSetsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRuleSetsAPIService GameCenterMatchmakingRuleSetsGetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRuleSetsAPIService GameCenterMatchmakingRuleSetsGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRuleSetsAPIService GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRuleSetsAPIService GameCenterMatchmakingRuleSetsRulesGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsRulesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRuleSetsAPIService GameCenterMatchmakingRuleSetsTeamsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsTeamsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRuleSetsAPIService GameCenterMatchmakingRuleSetsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
