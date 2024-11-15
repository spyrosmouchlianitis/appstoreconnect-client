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

// checks if the AppStoreVersionReleaseRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionReleaseRequestResponse{}

// AppStoreVersionReleaseRequestResponse struct for AppStoreVersionReleaseRequestResponse
type AppStoreVersionReleaseRequestResponse struct {
	Data AppStoreVersionReleaseRequest `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AppStoreVersionReleaseRequestResponse AppStoreVersionReleaseRequestResponse

// NewAppStoreVersionReleaseRequestResponse instantiates a new AppStoreVersionReleaseRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionReleaseRequestResponse(data AppStoreVersionReleaseRequest, links DocumentLinks) *AppStoreVersionReleaseRequestResponse {
	this := AppStoreVersionReleaseRequestResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionReleaseRequestResponseWithDefaults instantiates a new AppStoreVersionReleaseRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionReleaseRequestResponseWithDefaults() *AppStoreVersionReleaseRequestResponse {
	this := AppStoreVersionReleaseRequestResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionReleaseRequestResponse) GetData() AppStoreVersionReleaseRequest {
	if o == nil {
		var ret AppStoreVersionReleaseRequest
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionReleaseRequestResponse) GetDataOk() (*AppStoreVersionReleaseRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionReleaseRequestResponse) SetData(v AppStoreVersionReleaseRequest) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionReleaseRequestResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionReleaseRequestResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionReleaseRequestResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionReleaseRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionReleaseRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppStoreVersionReleaseRequestResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionReleaseRequestResponse := _AppStoreVersionReleaseRequestResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionReleaseRequestResponse)

	if err != nil {
		return err
	}

	*o = AppStoreVersionReleaseRequestResponse(varAppStoreVersionReleaseRequestResponse)

	return err
}

type NullableAppStoreVersionReleaseRequestResponse struct {
	value *AppStoreVersionReleaseRequestResponse
	isSet bool
}

func (v NullableAppStoreVersionReleaseRequestResponse) Get() *AppStoreVersionReleaseRequestResponse {
	return v.value
}

func (v *NullableAppStoreVersionReleaseRequestResponse) Set(val *AppStoreVersionReleaseRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionReleaseRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionReleaseRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionReleaseRequestResponse(val *AppStoreVersionReleaseRequestResponse) *NullableAppStoreVersionReleaseRequestResponse {
	return &NullableAppStoreVersionReleaseRequestResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionReleaseRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionReleaseRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


