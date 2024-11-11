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

// checks if the CiArtifactAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiArtifactAttributes{}

// CiArtifactAttributes struct for CiArtifactAttributes
type CiArtifactAttributes struct {
	FileType *string `json:"fileType,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	FileSize *int32 `json:"fileSize,omitempty"`
	DownloadUrl *string `json:"downloadUrl,omitempty"`
}

// NewCiArtifactAttributes instantiates a new CiArtifactAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiArtifactAttributes() *CiArtifactAttributes {
	this := CiArtifactAttributes{}
	return &this
}

// NewCiArtifactAttributesWithDefaults instantiates a new CiArtifactAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiArtifactAttributesWithDefaults() *CiArtifactAttributes {
	this := CiArtifactAttributes{}
	return &this
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *CiArtifactAttributes) GetFileType() string {
	if o == nil || IsNil(o.FileType) {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiArtifactAttributes) GetFileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileType) {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *CiArtifactAttributes) HasFileType() bool {
	if o != nil && !IsNil(o.FileType) {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *CiArtifactAttributes) SetFileType(v string) {
	o.FileType = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *CiArtifactAttributes) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiArtifactAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *CiArtifactAttributes) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *CiArtifactAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *CiArtifactAttributes) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiArtifactAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *CiArtifactAttributes) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *CiArtifactAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *CiArtifactAttributes) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiArtifactAttributes) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *CiArtifactAttributes) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *CiArtifactAttributes) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

func (o CiArtifactAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiArtifactAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileType) {
		toSerialize["fileType"] = o.FileType
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	return toSerialize, nil
}

type NullableCiArtifactAttributes struct {
	value *CiArtifactAttributes
	isSet bool
}

func (v NullableCiArtifactAttributes) Get() *CiArtifactAttributes {
	return v.value
}

func (v *NullableCiArtifactAttributes) Set(val *CiArtifactAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCiArtifactAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCiArtifactAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiArtifactAttributes(val *CiArtifactAttributes) *NullableCiArtifactAttributes {
	return &NullableCiArtifactAttributes{value: val, isSet: true}
}

func (v NullableCiArtifactAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiArtifactAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


