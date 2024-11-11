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

// checks if the GameCenterAchievementImageAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementImageAttributes{}

// GameCenterAchievementImageAttributes struct for GameCenterAchievementImageAttributes
type GameCenterAchievementImageAttributes struct {
	FileSize *int32 `json:"fileSize,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	ImageAsset *ImageAsset `json:"imageAsset,omitempty"`
	UploadOperations []UploadOperation `json:"uploadOperations,omitempty"`
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
}

// NewGameCenterAchievementImageAttributes instantiates a new GameCenterAchievementImageAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementImageAttributes() *GameCenterAchievementImageAttributes {
	this := GameCenterAchievementImageAttributes{}
	return &this
}

// NewGameCenterAchievementImageAttributesWithDefaults instantiates a new GameCenterAchievementImageAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementImageAttributesWithDefaults() *GameCenterAchievementImageAttributes {
	this := GameCenterAchievementImageAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *GameCenterAchievementImageAttributes) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *GameCenterAchievementImageAttributes) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *GameCenterAchievementImageAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *GameCenterAchievementImageAttributes) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *GameCenterAchievementImageAttributes) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *GameCenterAchievementImageAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetImageAsset returns the ImageAsset field value if set, zero value otherwise.
func (o *GameCenterAchievementImageAttributes) GetImageAsset() ImageAsset {
	if o == nil || IsNil(o.ImageAsset) {
		var ret ImageAsset
		return ret
	}
	return *o.ImageAsset
}

// GetImageAssetOk returns a tuple with the ImageAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageAttributes) GetImageAssetOk() (*ImageAsset, bool) {
	if o == nil || IsNil(o.ImageAsset) {
		return nil, false
	}
	return o.ImageAsset, true
}

// HasImageAsset returns a boolean if a field has been set.
func (o *GameCenterAchievementImageAttributes) HasImageAsset() bool {
	if o != nil && !IsNil(o.ImageAsset) {
		return true
	}

	return false
}

// SetImageAsset gets a reference to the given ImageAsset and assigns it to the ImageAsset field.
func (o *GameCenterAchievementImageAttributes) SetImageAsset(v ImageAsset) {
	o.ImageAsset = &v
}

// GetUploadOperations returns the UploadOperations field value if set, zero value otherwise.
func (o *GameCenterAchievementImageAttributes) GetUploadOperations() []UploadOperation {
	if o == nil || IsNil(o.UploadOperations) {
		var ret []UploadOperation
		return ret
	}
	return o.UploadOperations
}

// GetUploadOperationsOk returns a tuple with the UploadOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageAttributes) GetUploadOperationsOk() ([]UploadOperation, bool) {
	if o == nil || IsNil(o.UploadOperations) {
		return nil, false
	}
	return o.UploadOperations, true
}

// HasUploadOperations returns a boolean if a field has been set.
func (o *GameCenterAchievementImageAttributes) HasUploadOperations() bool {
	if o != nil && !IsNil(o.UploadOperations) {
		return true
	}

	return false
}

// SetUploadOperations gets a reference to the given []UploadOperation and assigns it to the UploadOperations field.
func (o *GameCenterAchievementImageAttributes) SetUploadOperations(v []UploadOperation) {
	o.UploadOperations = v
}

// GetAssetDeliveryState returns the AssetDeliveryState field value if set, zero value otherwise.
func (o *GameCenterAchievementImageAttributes) GetAssetDeliveryState() AppMediaAssetState {
	if o == nil || IsNil(o.AssetDeliveryState) {
		var ret AppMediaAssetState
		return ret
	}
	return *o.AssetDeliveryState
}

// GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool) {
	if o == nil || IsNil(o.AssetDeliveryState) {
		return nil, false
	}
	return o.AssetDeliveryState, true
}

// HasAssetDeliveryState returns a boolean if a field has been set.
func (o *GameCenterAchievementImageAttributes) HasAssetDeliveryState() bool {
	if o != nil && !IsNil(o.AssetDeliveryState) {
		return true
	}

	return false
}

// SetAssetDeliveryState gets a reference to the given AppMediaAssetState and assigns it to the AssetDeliveryState field.
func (o *GameCenterAchievementImageAttributes) SetAssetDeliveryState(v AppMediaAssetState) {
	o.AssetDeliveryState = &v
}

func (o GameCenterAchievementImageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementImageAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.ImageAsset) {
		toSerialize["imageAsset"] = o.ImageAsset
	}
	if !IsNil(o.UploadOperations) {
		toSerialize["uploadOperations"] = o.UploadOperations
	}
	if !IsNil(o.AssetDeliveryState) {
		toSerialize["assetDeliveryState"] = o.AssetDeliveryState
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementImageAttributes struct {
	value *GameCenterAchievementImageAttributes
	isSet bool
}

func (v NullableGameCenterAchievementImageAttributes) Get() *GameCenterAchievementImageAttributes {
	return v.value
}

func (v *NullableGameCenterAchievementImageAttributes) Set(val *GameCenterAchievementImageAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementImageAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementImageAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementImageAttributes(val *GameCenterAchievementImageAttributes) *NullableGameCenterAchievementImageAttributes {
	return &NullableGameCenterAchievementImageAttributes{value: val, isSet: true}
}

func (v NullableGameCenterAchievementImageAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementImageAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


