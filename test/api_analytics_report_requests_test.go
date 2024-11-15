/*
App Store Connect API

Testing AnalyticsReportRequestsAPIService

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

func Test_appstoreconnectclient_AnalyticsReportRequestsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AnalyticsReportRequestsAPIService AnalyticsReportRequestsCreateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AnalyticsReportRequestsAPI.AnalyticsReportRequestsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsReportRequestsAPIService AnalyticsReportRequestsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.AnalyticsReportRequestsAPI.AnalyticsReportRequestsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsReportRequestsAPIService AnalyticsReportRequestsGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AnalyticsReportRequestsAPI.AnalyticsReportRequestsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnalyticsReportRequestsAPIService AnalyticsReportRequestsReportsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AnalyticsReportRequestsAPI.AnalyticsReportRequestsReportsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
