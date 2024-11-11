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

// checks if the AppStoreVersionAppClipDefaultExperienceLinkageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionAppClipDefaultExperienceLinkageResponse{}

// AppStoreVersionAppClipDefaultExperienceLinkageResponse struct for AppStoreVersionAppClipDefaultExperienceLinkageResponse
type AppStoreVersionAppClipDefaultExperienceLinkageResponse struct {
	Data AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AppStoreVersionAppClipDefaultExperienceLinkageResponse AppStoreVersionAppClipDefaultExperienceLinkageResponse

// NewAppStoreVersionAppClipDefaultExperienceLinkageResponse instantiates a new AppStoreVersionAppClipDefaultExperienceLinkageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionAppClipDefaultExperienceLinkageResponse(data AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData, links DocumentLinks) *AppStoreVersionAppClipDefaultExperienceLinkageResponse {
	this := AppStoreVersionAppClipDefaultExperienceLinkageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionAppClipDefaultExperienceLinkageResponseWithDefaults instantiates a new AppStoreVersionAppClipDefaultExperienceLinkageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionAppClipDefaultExperienceLinkageResponseWithDefaults() *AppStoreVersionAppClipDefaultExperienceLinkageResponse {
	this := AppStoreVersionAppClipDefaultExperienceLinkageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionAppClipDefaultExperienceLinkageResponse) GetData() AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData {
	if o == nil {
		var ret AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAppClipDefaultExperienceLinkageResponse) GetDataOk() (*AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionAppClipDefaultExperienceLinkageResponse) SetData(v AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionAppClipDefaultExperienceLinkageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAppClipDefaultExperienceLinkageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionAppClipDefaultExperienceLinkageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionAppClipDefaultExperienceLinkageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionAppClipDefaultExperienceLinkageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppStoreVersionAppClipDefaultExperienceLinkageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionAppClipDefaultExperienceLinkageResponse := _AppStoreVersionAppClipDefaultExperienceLinkageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionAppClipDefaultExperienceLinkageResponse)

	if err != nil {
		return err
	}

	*o = AppStoreVersionAppClipDefaultExperienceLinkageResponse(varAppStoreVersionAppClipDefaultExperienceLinkageResponse)

	return err
}

type NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse struct {
	value *AppStoreVersionAppClipDefaultExperienceLinkageResponse
	isSet bool
}

func (v NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse) Get() *AppStoreVersionAppClipDefaultExperienceLinkageResponse {
	return v.value
}

func (v *NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse) Set(val *AppStoreVersionAppClipDefaultExperienceLinkageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionAppClipDefaultExperienceLinkageResponse(val *AppStoreVersionAppClipDefaultExperienceLinkageResponse) *NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse {
	return &NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionAppClipDefaultExperienceLinkageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


