/*
App Store Connect API

Testing PreReleaseVersionsAPIService

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

func Test_appstoreconnectclient_PreReleaseVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PreReleaseVersionsAPIService PreReleaseVersionsAppGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PreReleaseVersionsAPI.PreReleaseVersionsAppGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PreReleaseVersionsAPIService PreReleaseVersionsBuildsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PreReleaseVersionsAPI.PreReleaseVersionsBuildsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PreReleaseVersionsAPIService PreReleaseVersionsGetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PreReleaseVersionsAPI.PreReleaseVersionsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PreReleaseVersionsAPIService PreReleaseVersionsGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PreReleaseVersionsAPI.PreReleaseVersionsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
