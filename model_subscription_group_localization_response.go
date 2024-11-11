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

// checks if the SubscriptionGroupLocalizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalizationResponse{}

// SubscriptionGroupLocalizationResponse struct for SubscriptionGroupLocalizationResponse
type SubscriptionGroupLocalizationResponse struct {
	Data SubscriptionGroupLocalization `json:"data"`
	Included []SubscriptionGroup `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _SubscriptionGroupLocalizationResponse SubscriptionGroupLocalizationResponse

// NewSubscriptionGroupLocalizationResponse instantiates a new SubscriptionGroupLocalizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalizationResponse(data SubscriptionGroupLocalization, links DocumentLinks) *SubscriptionGroupLocalizationResponse {
	this := SubscriptionGroupLocalizationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionGroupLocalizationResponseWithDefaults instantiates a new SubscriptionGroupLocalizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationResponseWithDefaults() *SubscriptionGroupLocalizationResponse {
	this := SubscriptionGroupLocalizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGroupLocalizationResponse) GetData() SubscriptionGroupLocalization {
	if o == nil {
		var ret SubscriptionGroupLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationResponse) GetDataOk() (*SubscriptionGroupLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionGroupLocalizationResponse) SetData(v SubscriptionGroupLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionGroupLocalizationResponse) GetIncluded() []SubscriptionGroup {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionGroup
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationResponse) GetIncludedOk() ([]SubscriptionGroup, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionGroupLocalizationResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionGroup and assigns it to the Included field.
func (o *SubscriptionGroupLocalizationResponse) SetIncluded(v []SubscriptionGroup) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionGroupLocalizationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionGroupLocalizationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionGroupLocalizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *SubscriptionGroupLocalizationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionGroupLocalizationResponse := _SubscriptionGroupLocalizationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionGroupLocalizationResponse)

	if err != nil {
		return err
	}

	*o = SubscriptionGroupLocalizationResponse(varSubscriptionGroupLocalizationResponse)

	return err
}

type NullableSubscriptionGroupLocalizationResponse struct {
	value *SubscriptionGroupLocalizationResponse
	isSet bool
}

func (v NullableSubscriptionGroupLocalizationResponse) Get() *SubscriptionGroupLocalizationResponse {
	return v.value
}

func (v *NullableSubscriptionGroupLocalizationResponse) Set(val *SubscriptionGroupLocalizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalizationResponse(val *SubscriptionGroupLocalizationResponse) *NullableSubscriptionGroupLocalizationResponse {
	return &NullableSubscriptionGroupLocalizationResponse{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


