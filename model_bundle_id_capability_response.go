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

// checks if the BundleIdCapabilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdCapabilityResponse{}

// BundleIdCapabilityResponse struct for BundleIdCapabilityResponse
type BundleIdCapabilityResponse struct {
	Data BundleIdCapability `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _BundleIdCapabilityResponse BundleIdCapabilityResponse

// NewBundleIdCapabilityResponse instantiates a new BundleIdCapabilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCapabilityResponse(data BundleIdCapability, links DocumentLinks) *BundleIdCapabilityResponse {
	this := BundleIdCapabilityResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBundleIdCapabilityResponseWithDefaults instantiates a new BundleIdCapabilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCapabilityResponseWithDefaults() *BundleIdCapabilityResponse {
	this := BundleIdCapabilityResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdCapabilityResponse) GetData() BundleIdCapability {
	if o == nil {
		var ret BundleIdCapability
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityResponse) GetDataOk() (*BundleIdCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleIdCapabilityResponse) SetData(v BundleIdCapability) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BundleIdCapabilityResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BundleIdCapabilityResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BundleIdCapabilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdCapabilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *BundleIdCapabilityResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBundleIdCapabilityResponse := _BundleIdCapabilityResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBundleIdCapabilityResponse)

	if err != nil {
		return err
	}

	*o = BundleIdCapabilityResponse(varBundleIdCapabilityResponse)

	return err
}

type NullableBundleIdCapabilityResponse struct {
	value *BundleIdCapabilityResponse
	isSet bool
}

func (v NullableBundleIdCapabilityResponse) Get() *BundleIdCapabilityResponse {
	return v.value
}

func (v *NullableBundleIdCapabilityResponse) Set(val *BundleIdCapabilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCapabilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCapabilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCapabilityResponse(val *BundleIdCapabilityResponse) *NullableBundleIdCapabilityResponse {
	return &NullableBundleIdCapabilityResponse{value: val, isSet: true}
}

func (v NullableBundleIdCapabilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCapabilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


