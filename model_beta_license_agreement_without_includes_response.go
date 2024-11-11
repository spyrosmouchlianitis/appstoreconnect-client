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

// checks if the BetaLicenseAgreementWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaLicenseAgreementWithoutIncludesResponse{}

// BetaLicenseAgreementWithoutIncludesResponse struct for BetaLicenseAgreementWithoutIncludesResponse
type BetaLicenseAgreementWithoutIncludesResponse struct {
	Data BetaLicenseAgreement `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _BetaLicenseAgreementWithoutIncludesResponse BetaLicenseAgreementWithoutIncludesResponse

// NewBetaLicenseAgreementWithoutIncludesResponse instantiates a new BetaLicenseAgreementWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaLicenseAgreementWithoutIncludesResponse(data BetaLicenseAgreement, links DocumentLinks) *BetaLicenseAgreementWithoutIncludesResponse {
	this := BetaLicenseAgreementWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaLicenseAgreementWithoutIncludesResponseWithDefaults instantiates a new BetaLicenseAgreementWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaLicenseAgreementWithoutIncludesResponseWithDefaults() *BetaLicenseAgreementWithoutIncludesResponse {
	this := BetaLicenseAgreementWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaLicenseAgreementWithoutIncludesResponse) GetData() BetaLicenseAgreement {
	if o == nil {
		var ret BetaLicenseAgreement
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementWithoutIncludesResponse) GetDataOk() (*BetaLicenseAgreement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaLicenseAgreementWithoutIncludesResponse) SetData(v BetaLicenseAgreement) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaLicenseAgreementWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaLicenseAgreementWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaLicenseAgreementWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaLicenseAgreementWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *BetaLicenseAgreementWithoutIncludesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBetaLicenseAgreementWithoutIncludesResponse := _BetaLicenseAgreementWithoutIncludesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaLicenseAgreementWithoutIncludesResponse)

	if err != nil {
		return err
	}

	*o = BetaLicenseAgreementWithoutIncludesResponse(varBetaLicenseAgreementWithoutIncludesResponse)

	return err
}

type NullableBetaLicenseAgreementWithoutIncludesResponse struct {
	value *BetaLicenseAgreementWithoutIncludesResponse
	isSet bool
}

func (v NullableBetaLicenseAgreementWithoutIncludesResponse) Get() *BetaLicenseAgreementWithoutIncludesResponse {
	return v.value
}

func (v *NullableBetaLicenseAgreementWithoutIncludesResponse) Set(val *BetaLicenseAgreementWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaLicenseAgreementWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaLicenseAgreementWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaLicenseAgreementWithoutIncludesResponse(val *BetaLicenseAgreementWithoutIncludesResponse) *NullableBetaLicenseAgreementWithoutIncludesResponse {
	return &NullableBetaLicenseAgreementWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableBetaLicenseAgreementWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaLicenseAgreementWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

