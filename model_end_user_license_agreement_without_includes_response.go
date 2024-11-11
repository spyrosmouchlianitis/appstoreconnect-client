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

// checks if the EndUserLicenseAgreementWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndUserLicenseAgreementWithoutIncludesResponse{}

// EndUserLicenseAgreementWithoutIncludesResponse struct for EndUserLicenseAgreementWithoutIncludesResponse
type EndUserLicenseAgreementWithoutIncludesResponse struct {
	Data EndUserLicenseAgreement `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _EndUserLicenseAgreementWithoutIncludesResponse EndUserLicenseAgreementWithoutIncludesResponse

// NewEndUserLicenseAgreementWithoutIncludesResponse instantiates a new EndUserLicenseAgreementWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementWithoutIncludesResponse(data EndUserLicenseAgreement, links DocumentLinks) *EndUserLicenseAgreementWithoutIncludesResponse {
	this := EndUserLicenseAgreementWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewEndUserLicenseAgreementWithoutIncludesResponseWithDefaults instantiates a new EndUserLicenseAgreementWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementWithoutIncludesResponseWithDefaults() *EndUserLicenseAgreementWithoutIncludesResponse {
	this := EndUserLicenseAgreementWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *EndUserLicenseAgreementWithoutIncludesResponse) GetData() EndUserLicenseAgreement {
	if o == nil {
		var ret EndUserLicenseAgreement
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementWithoutIncludesResponse) GetDataOk() (*EndUserLicenseAgreement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EndUserLicenseAgreementWithoutIncludesResponse) SetData(v EndUserLicenseAgreement) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *EndUserLicenseAgreementWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *EndUserLicenseAgreementWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o EndUserLicenseAgreementWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndUserLicenseAgreementWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *EndUserLicenseAgreementWithoutIncludesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varEndUserLicenseAgreementWithoutIncludesResponse := _EndUserLicenseAgreementWithoutIncludesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndUserLicenseAgreementWithoutIncludesResponse)

	if err != nil {
		return err
	}

	*o = EndUserLicenseAgreementWithoutIncludesResponse(varEndUserLicenseAgreementWithoutIncludesResponse)

	return err
}

type NullableEndUserLicenseAgreementWithoutIncludesResponse struct {
	value *EndUserLicenseAgreementWithoutIncludesResponse
	isSet bool
}

func (v NullableEndUserLicenseAgreementWithoutIncludesResponse) Get() *EndUserLicenseAgreementWithoutIncludesResponse {
	return v.value
}

func (v *NullableEndUserLicenseAgreementWithoutIncludesResponse) Set(val *EndUserLicenseAgreementWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementWithoutIncludesResponse(val *EndUserLicenseAgreementWithoutIncludesResponse) *NullableEndUserLicenseAgreementWithoutIncludesResponse {
	return &NullableEndUserLicenseAgreementWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


