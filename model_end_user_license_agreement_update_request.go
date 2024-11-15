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

// checks if the EndUserLicenseAgreementUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndUserLicenseAgreementUpdateRequest{}

// EndUserLicenseAgreementUpdateRequest struct for EndUserLicenseAgreementUpdateRequest
type EndUserLicenseAgreementUpdateRequest struct {
	Data EndUserLicenseAgreementUpdateRequestData `json:"data"`
}

type _EndUserLicenseAgreementUpdateRequest EndUserLicenseAgreementUpdateRequest

// NewEndUserLicenseAgreementUpdateRequest instantiates a new EndUserLicenseAgreementUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementUpdateRequest(data EndUserLicenseAgreementUpdateRequestData) *EndUserLicenseAgreementUpdateRequest {
	this := EndUserLicenseAgreementUpdateRequest{}
	this.Data = data
	return &this
}

// NewEndUserLicenseAgreementUpdateRequestWithDefaults instantiates a new EndUserLicenseAgreementUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementUpdateRequestWithDefaults() *EndUserLicenseAgreementUpdateRequest {
	this := EndUserLicenseAgreementUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *EndUserLicenseAgreementUpdateRequest) GetData() EndUserLicenseAgreementUpdateRequestData {
	if o == nil {
		var ret EndUserLicenseAgreementUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementUpdateRequest) GetDataOk() (*EndUserLicenseAgreementUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EndUserLicenseAgreementUpdateRequest) SetData(v EndUserLicenseAgreementUpdateRequestData) {
	o.Data = v
}

func (o EndUserLicenseAgreementUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndUserLicenseAgreementUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EndUserLicenseAgreementUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varEndUserLicenseAgreementUpdateRequest := _EndUserLicenseAgreementUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndUserLicenseAgreementUpdateRequest)

	if err != nil {
		return err
	}

	*o = EndUserLicenseAgreementUpdateRequest(varEndUserLicenseAgreementUpdateRequest)

	return err
}

type NullableEndUserLicenseAgreementUpdateRequest struct {
	value *EndUserLicenseAgreementUpdateRequest
	isSet bool
}

func (v NullableEndUserLicenseAgreementUpdateRequest) Get() *EndUserLicenseAgreementUpdateRequest {
	return v.value
}

func (v *NullableEndUserLicenseAgreementUpdateRequest) Set(val *EndUserLicenseAgreementUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementUpdateRequest(val *EndUserLicenseAgreementUpdateRequest) *NullableEndUserLicenseAgreementUpdateRequest {
	return &NullableEndUserLicenseAgreementUpdateRequest{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


