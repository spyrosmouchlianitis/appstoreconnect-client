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

// checks if the AppCategoryWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCategoryWithoutIncludesResponse{}

// AppCategoryWithoutIncludesResponse struct for AppCategoryWithoutIncludesResponse
type AppCategoryWithoutIncludesResponse struct {
	Data AppCategory `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AppCategoryWithoutIncludesResponse AppCategoryWithoutIncludesResponse

// NewAppCategoryWithoutIncludesResponse instantiates a new AppCategoryWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCategoryWithoutIncludesResponse(data AppCategory, links DocumentLinks) *AppCategoryWithoutIncludesResponse {
	this := AppCategoryWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppCategoryWithoutIncludesResponseWithDefaults instantiates a new AppCategoryWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCategoryWithoutIncludesResponseWithDefaults() *AppCategoryWithoutIncludesResponse {
	this := AppCategoryWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppCategoryWithoutIncludesResponse) GetData() AppCategory {
	if o == nil {
		var ret AppCategory
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCategoryWithoutIncludesResponse) GetDataOk() (*AppCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppCategoryWithoutIncludesResponse) SetData(v AppCategory) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppCategoryWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppCategoryWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppCategoryWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppCategoryWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCategoryWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppCategoryWithoutIncludesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppCategoryWithoutIncludesResponse := _AppCategoryWithoutIncludesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppCategoryWithoutIncludesResponse)

	if err != nil {
		return err
	}

	*o = AppCategoryWithoutIncludesResponse(varAppCategoryWithoutIncludesResponse)

	return err
}

type NullableAppCategoryWithoutIncludesResponse struct {
	value *AppCategoryWithoutIncludesResponse
	isSet bool
}

func (v NullableAppCategoryWithoutIncludesResponse) Get() *AppCategoryWithoutIncludesResponse {
	return v.value
}

func (v *NullableAppCategoryWithoutIncludesResponse) Set(val *AppCategoryWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCategoryWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCategoryWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCategoryWithoutIncludesResponse(val *AppCategoryWithoutIncludesResponse) *NullableAppCategoryWithoutIncludesResponse {
	return &NullableAppCategoryWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableAppCategoryWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCategoryWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


