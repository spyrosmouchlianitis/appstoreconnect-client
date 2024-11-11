/*
App Store Connect API

Testing AppEncryptionDeclarationsAPIService

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

func Test_appstoreconnectclient_AppEncryptionDeclarationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppEncryptionDeclarationsAPIService AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppEncryptionDeclarationsAPIService AppEncryptionDeclarationsAppGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsAppGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppEncryptionDeclarationsAPIService AppEncryptionDeclarationsBuildsCreateToManyRelationship", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsBuildsCreateToManyRelationship(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppEncryptionDeclarationsAPIService AppEncryptionDeclarationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppEncryptionDeclarationsAPIService AppEncryptionDeclarationsGetCollection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppEncryptionDeclarationsAPIService AppEncryptionDeclarationsGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
