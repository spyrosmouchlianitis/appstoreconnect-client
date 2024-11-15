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

// checks if the AlternativeDistributionPackageDeltaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeDistributionPackageDeltaResponse{}

// AlternativeDistributionPackageDeltaResponse struct for AlternativeDistributionPackageDeltaResponse
type AlternativeDistributionPackageDeltaResponse struct {
	Data AlternativeDistributionPackageDelta `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AlternativeDistributionPackageDeltaResponse AlternativeDistributionPackageDeltaResponse

// NewAlternativeDistributionPackageDeltaResponse instantiates a new AlternativeDistributionPackageDeltaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeDistributionPackageDeltaResponse(data AlternativeDistributionPackageDelta, links DocumentLinks) *AlternativeDistributionPackageDeltaResponse {
	this := AlternativeDistributionPackageDeltaResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAlternativeDistributionPackageDeltaResponseWithDefaults instantiates a new AlternativeDistributionPackageDeltaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeDistributionPackageDeltaResponseWithDefaults() *AlternativeDistributionPackageDeltaResponse {
	this := AlternativeDistributionPackageDeltaResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AlternativeDistributionPackageDeltaResponse) GetData() AlternativeDistributionPackageDelta {
	if o == nil {
		var ret AlternativeDistributionPackageDelta
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionPackageDeltaResponse) GetDataOk() (*AlternativeDistributionPackageDelta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AlternativeDistributionPackageDeltaResponse) SetData(v AlternativeDistributionPackageDelta) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AlternativeDistributionPackageDeltaResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionPackageDeltaResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AlternativeDistributionPackageDeltaResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AlternativeDistributionPackageDeltaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeDistributionPackageDeltaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AlternativeDistributionPackageDeltaResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAlternativeDistributionPackageDeltaResponse := _AlternativeDistributionPackageDeltaResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlternativeDistributionPackageDeltaResponse)

	if err != nil {
		return err
	}

	*o = AlternativeDistributionPackageDeltaResponse(varAlternativeDistributionPackageDeltaResponse)

	return err
}

type NullableAlternativeDistributionPackageDeltaResponse struct {
	value *AlternativeDistributionPackageDeltaResponse
	isSet bool
}

func (v NullableAlternativeDistributionPackageDeltaResponse) Get() *AlternativeDistributionPackageDeltaResponse {
	return v.value
}

func (v *NullableAlternativeDistributionPackageDeltaResponse) Set(val *AlternativeDistributionPackageDeltaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeDistributionPackageDeltaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeDistributionPackageDeltaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeDistributionPackageDeltaResponse(val *AlternativeDistributionPackageDeltaResponse) *NullableAlternativeDistributionPackageDeltaResponse {
	return &NullableAlternativeDistributionPackageDeltaResponse{value: val, isSet: true}
}

func (v NullableAlternativeDistributionPackageDeltaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeDistributionPackageDeltaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


