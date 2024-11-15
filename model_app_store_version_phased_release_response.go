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

// checks if the AppStoreVersionPhasedReleaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPhasedReleaseResponse{}

// AppStoreVersionPhasedReleaseResponse struct for AppStoreVersionPhasedReleaseResponse
type AppStoreVersionPhasedReleaseResponse struct {
	Data AppStoreVersionPhasedRelease `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AppStoreVersionPhasedReleaseResponse AppStoreVersionPhasedReleaseResponse

// NewAppStoreVersionPhasedReleaseResponse instantiates a new AppStoreVersionPhasedReleaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPhasedReleaseResponse(data AppStoreVersionPhasedRelease, links DocumentLinks) *AppStoreVersionPhasedReleaseResponse {
	this := AppStoreVersionPhasedReleaseResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionPhasedReleaseResponseWithDefaults instantiates a new AppStoreVersionPhasedReleaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPhasedReleaseResponseWithDefaults() *AppStoreVersionPhasedReleaseResponse {
	this := AppStoreVersionPhasedReleaseResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionPhasedReleaseResponse) GetData() AppStoreVersionPhasedRelease {
	if o == nil {
		var ret AppStoreVersionPhasedRelease
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseResponse) GetDataOk() (*AppStoreVersionPhasedRelease, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionPhasedReleaseResponse) SetData(v AppStoreVersionPhasedRelease) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionPhasedReleaseResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionPhasedReleaseResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionPhasedReleaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPhasedReleaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppStoreVersionPhasedReleaseResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionPhasedReleaseResponse := _AppStoreVersionPhasedReleaseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionPhasedReleaseResponse)

	if err != nil {
		return err
	}

	*o = AppStoreVersionPhasedReleaseResponse(varAppStoreVersionPhasedReleaseResponse)

	return err
}

type NullableAppStoreVersionPhasedReleaseResponse struct {
	value *AppStoreVersionPhasedReleaseResponse
	isSet bool
}

func (v NullableAppStoreVersionPhasedReleaseResponse) Get() *AppStoreVersionPhasedReleaseResponse {
	return v.value
}

func (v *NullableAppStoreVersionPhasedReleaseResponse) Set(val *AppStoreVersionPhasedReleaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPhasedReleaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPhasedReleaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPhasedReleaseResponse(val *AppStoreVersionPhasedReleaseResponse) *NullableAppStoreVersionPhasedReleaseResponse {
	return &NullableAppStoreVersionPhasedReleaseResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionPhasedReleaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPhasedReleaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


