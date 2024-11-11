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

// checks if the AlternativeDistributionKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeDistributionKeyResponse{}

// AlternativeDistributionKeyResponse struct for AlternativeDistributionKeyResponse
type AlternativeDistributionKeyResponse struct {
	Data AlternativeDistributionKey `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AlternativeDistributionKeyResponse AlternativeDistributionKeyResponse

// NewAlternativeDistributionKeyResponse instantiates a new AlternativeDistributionKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeDistributionKeyResponse(data AlternativeDistributionKey, links DocumentLinks) *AlternativeDistributionKeyResponse {
	this := AlternativeDistributionKeyResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAlternativeDistributionKeyResponseWithDefaults instantiates a new AlternativeDistributionKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeDistributionKeyResponseWithDefaults() *AlternativeDistributionKeyResponse {
	this := AlternativeDistributionKeyResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AlternativeDistributionKeyResponse) GetData() AlternativeDistributionKey {
	if o == nil {
		var ret AlternativeDistributionKey
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionKeyResponse) GetDataOk() (*AlternativeDistributionKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AlternativeDistributionKeyResponse) SetData(v AlternativeDistributionKey) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AlternativeDistributionKeyResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionKeyResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AlternativeDistributionKeyResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AlternativeDistributionKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeDistributionKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AlternativeDistributionKeyResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAlternativeDistributionKeyResponse := _AlternativeDistributionKeyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlternativeDistributionKeyResponse)

	if err != nil {
		return err
	}

	*o = AlternativeDistributionKeyResponse(varAlternativeDistributionKeyResponse)

	return err
}

type NullableAlternativeDistributionKeyResponse struct {
	value *AlternativeDistributionKeyResponse
	isSet bool
}

func (v NullableAlternativeDistributionKeyResponse) Get() *AlternativeDistributionKeyResponse {
	return v.value
}

func (v *NullableAlternativeDistributionKeyResponse) Set(val *AlternativeDistributionKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeDistributionKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeDistributionKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeDistributionKeyResponse(val *AlternativeDistributionKeyResponse) *NullableAlternativeDistributionKeyResponse {
	return &NullableAlternativeDistributionKeyResponse{value: val, isSet: true}
}

func (v NullableAlternativeDistributionKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeDistributionKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


