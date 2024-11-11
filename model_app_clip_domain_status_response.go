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

// checks if the AppClipDomainStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDomainStatusResponse{}

// AppClipDomainStatusResponse struct for AppClipDomainStatusResponse
type AppClipDomainStatusResponse struct {
	Data AppClipDomainStatus `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AppClipDomainStatusResponse AppClipDomainStatusResponse

// NewAppClipDomainStatusResponse instantiates a new AppClipDomainStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDomainStatusResponse(data AppClipDomainStatus, links DocumentLinks) *AppClipDomainStatusResponse {
	this := AppClipDomainStatusResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppClipDomainStatusResponseWithDefaults instantiates a new AppClipDomainStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDomainStatusResponseWithDefaults() *AppClipDomainStatusResponse {
	this := AppClipDomainStatusResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipDomainStatusResponse) GetData() AppClipDomainStatus {
	if o == nil {
		var ret AppClipDomainStatus
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipDomainStatusResponse) GetDataOk() (*AppClipDomainStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipDomainStatusResponse) SetData(v AppClipDomainStatus) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppClipDomainStatusResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipDomainStatusResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipDomainStatusResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppClipDomainStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDomainStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppClipDomainStatusResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipDomainStatusResponse := _AppClipDomainStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipDomainStatusResponse)

	if err != nil {
		return err
	}

	*o = AppClipDomainStatusResponse(varAppClipDomainStatusResponse)

	return err
}

type NullableAppClipDomainStatusResponse struct {
	value *AppClipDomainStatusResponse
	isSet bool
}

func (v NullableAppClipDomainStatusResponse) Get() *AppClipDomainStatusResponse {
	return v.value
}

func (v *NullableAppClipDomainStatusResponse) Set(val *AppClipDomainStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDomainStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDomainStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDomainStatusResponse(val *AppClipDomainStatusResponse) *NullableAppClipDomainStatusResponse {
	return &NullableAppClipDomainStatusResponse{value: val, isSet: true}
}

func (v NullableAppClipDomainStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDomainStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


