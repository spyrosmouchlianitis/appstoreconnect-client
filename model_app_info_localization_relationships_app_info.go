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

// checks if the AppInfoLocalizationRelationshipsAppInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoLocalizationRelationshipsAppInfo{}

// AppInfoLocalizationRelationshipsAppInfo struct for AppInfoLocalizationRelationshipsAppInfo
type AppInfoLocalizationRelationshipsAppInfo struct {
	Data *AppInfoLocalizationRelationshipsAppInfoData `json:"data,omitempty"`
}

// NewAppInfoLocalizationRelationshipsAppInfo instantiates a new AppInfoLocalizationRelationshipsAppInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationRelationshipsAppInfo() *AppInfoLocalizationRelationshipsAppInfo {
	this := AppInfoLocalizationRelationshipsAppInfo{}
	return &this
}

// NewAppInfoLocalizationRelationshipsAppInfoWithDefaults instantiates a new AppInfoLocalizationRelationshipsAppInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationRelationshipsAppInfoWithDefaults() *AppInfoLocalizationRelationshipsAppInfo {
	this := AppInfoLocalizationRelationshipsAppInfo{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppInfoLocalizationRelationshipsAppInfo) GetData() AppInfoLocalizationRelationshipsAppInfoData {
	if o == nil || IsNil(o.Data) {
		var ret AppInfoLocalizationRelationshipsAppInfoData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationRelationshipsAppInfo) GetDataOk() (*AppInfoLocalizationRelationshipsAppInfoData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppInfoLocalizationRelationshipsAppInfo) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppInfoLocalizationRelationshipsAppInfoData and assigns it to the Data field.
func (o *AppInfoLocalizationRelationshipsAppInfo) SetData(v AppInfoLocalizationRelationshipsAppInfoData) {
	o.Data = &v
}

func (o AppInfoLocalizationRelationshipsAppInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoLocalizationRelationshipsAppInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppInfoLocalizationRelationshipsAppInfo struct {
	value *AppInfoLocalizationRelationshipsAppInfo
	isSet bool
}

func (v NullableAppInfoLocalizationRelationshipsAppInfo) Get() *AppInfoLocalizationRelationshipsAppInfo {
	return v.value
}

func (v *NullableAppInfoLocalizationRelationshipsAppInfo) Set(val *AppInfoLocalizationRelationshipsAppInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationRelationshipsAppInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationRelationshipsAppInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationRelationshipsAppInfo(val *AppInfoLocalizationRelationshipsAppInfo) *NullableAppInfoLocalizationRelationshipsAppInfo {
	return &NullableAppInfoLocalizationRelationshipsAppInfo{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationRelationshipsAppInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationRelationshipsAppInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


