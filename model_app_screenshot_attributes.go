/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the AppScreenshotAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotAttributes{}

// AppScreenshotAttributes struct for AppScreenshotAttributes
type AppScreenshotAttributes struct {
	FileSize *int32 `json:"fileSize,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	ImageAsset *ImageAsset `json:"imageAsset,omitempty"`
	AssetToken *string `json:"assetToken,omitempty"`
	AssetType *string `json:"assetType,omitempty"`
	UploadOperations []UploadOperation `json:"uploadOperations,omitempty"`
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
}

// NewAppScreenshotAttributes instantiates a new AppScreenshotAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotAttributes() *AppScreenshotAttributes {
	this := AppScreenshotAttributes{}
	return &this
}

// NewAppScreenshotAttributesWithDefaults instantiates a new AppScreenshotAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotAttributesWithDefaults() *AppScreenshotAttributes {
	this := AppScreenshotAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *AppScreenshotAttributes) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *AppScreenshotAttributes) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *AppScreenshotAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AppScreenshotAttributes) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AppScreenshotAttributes) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AppScreenshotAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetSourceFileChecksum returns the SourceFileChecksum field value if set, zero value otherwise.
func (o *AppScreenshotAttributes) GetSourceFileChecksum() string {
	if o == nil || IsNil(o.SourceFileChecksum) {
		var ret string
		return ret
	}
	return *o.SourceFileChecksum
}

// GetSourceFileChecksumOk returns a tuple with the SourceFileChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotAttributes) GetSourceFileChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.SourceFileChecksum) {
		return nil, false
	}
	return o.SourceFileChecksum, true
}

// HasSourceFileChecksum returns a boolean if a field has been set.
func (o *AppScreenshotAttributes) HasSourceFileChecksum() bool {
	if o != nil && !IsNil(o.SourceFileChecksum) {
		return true
	}

	return false
}

// SetSourceFileChecksum gets a reference to the given string and assigns it to the SourceFileChecksum field.
func (o *AppScreenshotAttributes) SetSourceFileChecksum(v string) {
	o.SourceFileChecksum = &v
}

// GetImageAsset returns the ImageAsset field value if set, zero value otherwise.
func (o *AppScreenshotAttributes) GetImageAsset() ImageAsset {
	if o == nil || IsNil(o.ImageAsset) {
		var ret ImageAsset
		return ret
	}
	return *o.ImageAsset
}

// GetImageAssetOk returns a tuple with the ImageAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotAttributes) GetImageAssetOk() (*ImageAsset, bool) {
	if o == nil || IsNil(o.ImageAsset) {
		return nil, false
	}
	return o.ImageAsset, true
}

// HasImageAsset returns a boolean if a field has been set.
func (o *AppScreenshotAttributes) HasImageAsset() bool {
	if o != nil && !IsNil(o.ImageAsset) {
		return true
	}

	return false
}

// SetImageAsset gets a reference to the given ImageAsset and assigns it to the ImageAsset field.
func (o *AppScreenshotAttributes) SetImageAsset(v ImageAsset) {
	o.ImageAsset = &v
}

// GetAssetToken returns the AssetToken field value if set, zero value otherwise.
func (o *AppScreenshotAttributes) GetAssetToken() string {
	if o == nil || IsNil(o.AssetToken) {
		var ret string
		return ret
	}
	return *o.AssetToken
}

// GetAssetTokenOk returns a tuple with the AssetToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotAttributes) GetAssetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AssetToken) {
		return nil, false
	}
	return o.AssetToken, true
}

// HasAssetToken returns a boolean if a field has been set.
func (o *AppScreenshotAttributes) HasAssetToken() bool {
	if o != nil && !IsNil(o.AssetToken) {
		return true
	}

	return false
}

// SetAssetToken gets a reference to the given string and assigns it to the AssetToken field.
func (o *AppScreenshotAttributes) SetAssetToken(v string) {
	o.AssetToken = &v
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *AppScreenshotAttributes) GetAssetType() string {
	if o == nil || IsNil(o.AssetType) {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotAttributes) GetAssetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssetType) {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *AppScreenshotAttributes) HasAssetType() bool {
	if o != nil && !IsNil(o.AssetType) {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *AppScreenshotAttributes) SetAssetType(v string) {
	o.AssetType = &v
}

// GetUploadOperations returns the UploadOperations field value if set, zero value otherwise.
func (o *AppScreenshotAttributes) GetUploadOperations() []UploadOperation {
	if o == nil || IsNil(o.UploadOperations) {
		var ret []UploadOperation
		return ret
	}
	return o.UploadOperations
}

// GetUploadOperationsOk returns a tuple with the UploadOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotAttributes) GetUploadOperationsOk() ([]UploadOperation, bool) {
	if o == nil || IsNil(o.UploadOperations) {
		return nil, false
	}
	return o.UploadOperations, true
}

// HasUploadOperations returns a boolean if a field has been set.
func (o *AppScreenshotAttributes) HasUploadOperations() bool {
	if o != nil && !IsNil(o.UploadOperations) {
		return true
	}

	return false
}

// SetUploadOperations gets a reference to the given []UploadOperation and assigns it to the UploadOperations field.
func (o *AppScreenshotAttributes) SetUploadOperations(v []UploadOperation) {
	o.UploadOperations = v
}

// GetAssetDeliveryState returns the AssetDeliveryState field value if set, zero value otherwise.
func (o *AppScreenshotAttributes) GetAssetDeliveryState() AppMediaAssetState {
	if o == nil || IsNil(o.AssetDeliveryState) {
		var ret AppMediaAssetState
		return ret
	}
	return *o.AssetDeliveryState
}

// GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool) {
	if o == nil || IsNil(o.AssetDeliveryState) {
		return nil, false
	}
	return o.AssetDeliveryState, true
}

// HasAssetDeliveryState returns a boolean if a field has been set.
func (o *AppScreenshotAttributes) HasAssetDeliveryState() bool {
	if o != nil && !IsNil(o.AssetDeliveryState) {
		return true
	}

	return false
}

// SetAssetDeliveryState gets a reference to the given AppMediaAssetState and assigns it to the AssetDeliveryState field.
func (o *AppScreenshotAttributes) SetAssetDeliveryState(v AppMediaAssetState) {
	o.AssetDeliveryState = &v
}

func (o AppScreenshotAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.SourceFileChecksum) {
		toSerialize["sourceFileChecksum"] = o.SourceFileChecksum
	}
	if !IsNil(o.ImageAsset) {
		toSerialize["imageAsset"] = o.ImageAsset
	}
	if !IsNil(o.AssetToken) {
		toSerialize["assetToken"] = o.AssetToken
	}
	if !IsNil(o.AssetType) {
		toSerialize["assetType"] = o.AssetType
	}
	if !IsNil(o.UploadOperations) {
		toSerialize["uploadOperations"] = o.UploadOperations
	}
	if !IsNil(o.AssetDeliveryState) {
		toSerialize["assetDeliveryState"] = o.AssetDeliveryState
	}
	return toSerialize, nil
}

type NullableAppScreenshotAttributes struct {
	value *AppScreenshotAttributes
	isSet bool
}

func (v NullableAppScreenshotAttributes) Get() *AppScreenshotAttributes {
	return v.value
}

func (v *NullableAppScreenshotAttributes) Set(val *AppScreenshotAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotAttributes(val *AppScreenshotAttributes) *NullableAppScreenshotAttributes {
	return &NullableAppScreenshotAttributes{value: val, isSet: true}
}

func (v NullableAppScreenshotAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


