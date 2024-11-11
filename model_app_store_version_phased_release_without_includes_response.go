/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AppStoreVersionPhasedReleaseWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPhasedReleaseWithoutIncludesResponse{}

// AppStoreVersionPhasedReleaseWithoutIncludesResponse struct for AppStoreVersionPhasedReleaseWithoutIncludesResponse
type AppStoreVersionPhasedReleaseWithoutIncludesResponse struct {
	Data AppStoreVersionPhasedRelease `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AppStoreVersionPhasedReleaseWithoutIncludesResponse AppStoreVersionPhasedReleaseWithoutIncludesResponse

// NewAppStoreVersionPhasedReleaseWithoutIncludesResponse instantiates a new AppStoreVersionPhasedReleaseWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPhasedReleaseWithoutIncludesResponse(data AppStoreVersionPhasedRelease, links DocumentLinks) *AppStoreVersionPhasedReleaseWithoutIncludesResponse {
	this := AppStoreVersionPhasedReleaseWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionPhasedReleaseWithoutIncludesResponseWithDefaults instantiates a new AppStoreVersionPhasedReleaseWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPhasedReleaseWithoutIncludesResponseWithDefaults() *AppStoreVersionPhasedReleaseWithoutIncludesResponse {
	this := AppStoreVersionPhasedReleaseWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionPhasedReleaseWithoutIncludesResponse) GetData() AppStoreVersionPhasedRelease {
	if o == nil {
		var ret AppStoreVersionPhasedRelease
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseWithoutIncludesResponse) GetDataOk() (*AppStoreVersionPhasedRelease, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionPhasedReleaseWithoutIncludesResponse) SetData(v AppStoreVersionPhasedRelease) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionPhasedReleaseWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionPhasedReleaseWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionPhasedReleaseWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPhasedReleaseWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppStoreVersionPhasedReleaseWithoutIncludesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varAppStoreVersionPhasedReleaseWithoutIncludesResponse := _AppStoreVersionPhasedReleaseWithoutIncludesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionPhasedReleaseWithoutIncludesResponse)

	if err != nil {
		return err
	}

	*o = AppStoreVersionPhasedReleaseWithoutIncludesResponse(varAppStoreVersionPhasedReleaseWithoutIncludesResponse)

	return err
}

type NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse struct {
	value *AppStoreVersionPhasedReleaseWithoutIncludesResponse
	isSet bool
}

func (v NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse) Get() *AppStoreVersionPhasedReleaseWithoutIncludesResponse {
	return v.value
}

func (v *NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse) Set(val *AppStoreVersionPhasedReleaseWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPhasedReleaseWithoutIncludesResponse(val *AppStoreVersionPhasedReleaseWithoutIncludesResponse) *NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse {
	return &NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPhasedReleaseWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


