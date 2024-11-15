/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest{}

// AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest struct for AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest
type AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest struct {
	Data AlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData `json:"data"`
}

type _AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest

// NewAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest instantiates a new AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest(data AlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData) *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest {
	this := AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest{}
	this.Data = data
	return &this
}

// NewAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequestWithDefaults instantiates a new AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequestWithDefaults() *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest {
	this := AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) GetData() AlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData {
	if o == nil {
		var ret AlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) GetDataOk() (*AlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) SetData(v AlternativeDistributionPackageCreateRequestDataRelationshipsAppStoreVersionData) {
	o.Data = v
}

func (o AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest := _AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest)

	if err != nil {
		return err
	}

	*o = AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest(varAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest)

	return err
}

type NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest struct {
	value *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest
	isSet bool
}

func (v NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) Get() *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest {
	return v.value
}

func (v *NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) Set(val *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest(val *AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) *NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest {
	return &NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


