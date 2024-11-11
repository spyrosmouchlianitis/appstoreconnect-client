/*
App Store Connect API

Testing AppEncryptionDeclarationDocumentsAPIService

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

func Test_appstoreconnectclient_AppEncryptionDeclarationDocumentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppEncryptionDeclarationDocumentsAPIService AppEncryptionDeclarationDocumentsCreateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppEncryptionDeclarationDocumentsAPIService AppEncryptionDeclarationDocumentsGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppEncryptionDeclarationDocumentsAPIService AppEncryptionDeclarationDocumentsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
