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

// checks if the SubscriptionImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionImageResponse{}

// SubscriptionImageResponse struct for SubscriptionImageResponse
type SubscriptionImageResponse struct {
	Data SubscriptionImage `json:"data"`
	Included []Subscription `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _SubscriptionImageResponse SubscriptionImageResponse

// NewSubscriptionImageResponse instantiates a new SubscriptionImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionImageResponse(data SubscriptionImage, links DocumentLinks) *SubscriptionImageResponse {
	this := SubscriptionImageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionImageResponseWithDefaults instantiates a new SubscriptionImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionImageResponseWithDefaults() *SubscriptionImageResponse {
	this := SubscriptionImageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionImageResponse) GetData() SubscriptionImage {
	if o == nil {
		var ret SubscriptionImage
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionImageResponse) GetDataOk() (*SubscriptionImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionImageResponse) SetData(v SubscriptionImage) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionImageResponse) GetIncluded() []Subscription {
	if o == nil || IsNil(o.Included) {
		var ret []Subscription
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionImageResponse) GetIncludedOk() ([]Subscription, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionImageResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []Subscription and assigns it to the Included field.
func (o *SubscriptionImageResponse) SetIncluded(v []Subscription) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionImageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionImageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionImageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *SubscriptionImageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionImageResponse := _SubscriptionImageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionImageResponse)

	if err != nil {
		return err
	}

	*o = SubscriptionImageResponse(varSubscriptionImageResponse)

	return err
}

type NullableSubscriptionImageResponse struct {
	value *SubscriptionImageResponse
	isSet bool
}

func (v NullableSubscriptionImageResponse) Get() *SubscriptionImageResponse {
	return v.value
}

func (v *NullableSubscriptionImageResponse) Set(val *SubscriptionImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionImageResponse(val *SubscriptionImageResponse) *NullableSubscriptionImageResponse {
	return &NullableSubscriptionImageResponse{value: val, isSet: true}
}

func (v NullableSubscriptionImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


