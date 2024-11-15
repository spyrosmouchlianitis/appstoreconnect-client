/*
App Store Connect API

Testing CiXcodeVersionsAPIService

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

func Test_appstoreconnectclient_CiXcodeVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CiXcodeVersionsAPIService CiXcodeVersionsGetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CiXcodeVersionsAPI.CiXcodeVersionsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CiXcodeVersionsAPIService CiXcodeVersionsGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CiXcodeVersionsAPI.CiXcodeVersionsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CiXcodeVersionsAPIService CiXcodeVersionsMacOsVersionsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CiXcodeVersionsAPI.CiXcodeVersionsMacOsVersionsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
