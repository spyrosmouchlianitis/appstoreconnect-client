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

// checks if the CiTestResultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiTestResultResponse{}

// CiTestResultResponse struct for CiTestResultResponse
type CiTestResultResponse struct {
	Data CiTestResult `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _CiTestResultResponse CiTestResultResponse

// NewCiTestResultResponse instantiates a new CiTestResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiTestResultResponse(data CiTestResult, links DocumentLinks) *CiTestResultResponse {
	this := CiTestResultResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiTestResultResponseWithDefaults instantiates a new CiTestResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiTestResultResponseWithDefaults() *CiTestResultResponse {
	this := CiTestResultResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiTestResultResponse) GetData() CiTestResult {
	if o == nil {
		var ret CiTestResult
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiTestResultResponse) GetDataOk() (*CiTestResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiTestResultResponse) SetData(v CiTestResult) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *CiTestResultResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiTestResultResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiTestResultResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o CiTestResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiTestResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CiTestResultResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCiTestResultResponse := _CiTestResultResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiTestResultResponse)

	if err != nil {
		return err
	}

	*o = CiTestResultResponse(varCiTestResultResponse)

	return err
}

type NullableCiTestResultResponse struct {
	value *CiTestResultResponse
	isSet bool
}

func (v NullableCiTestResultResponse) Get() *CiTestResultResponse {
	return v.value
}

func (v *NullableCiTestResultResponse) Set(val *CiTestResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiTestResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiTestResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiTestResultResponse(val *CiTestResultResponse) *NullableCiTestResultResponse {
	return &NullableCiTestResultResponse{value: val, isSet: true}
}

func (v NullableCiTestResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiTestResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


