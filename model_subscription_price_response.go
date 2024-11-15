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

// checks if the SubscriptionPriceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPriceResponse{}

// SubscriptionPriceResponse struct for SubscriptionPriceResponse
type SubscriptionPriceResponse struct {
	Data SubscriptionPrice `json:"data"`
	Included []SubscriptionOfferCodePricesResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _SubscriptionPriceResponse SubscriptionPriceResponse

// NewSubscriptionPriceResponse instantiates a new SubscriptionPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPriceResponse(data SubscriptionPrice, links DocumentLinks) *SubscriptionPriceResponse {
	this := SubscriptionPriceResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionPriceResponseWithDefaults instantiates a new SubscriptionPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPriceResponseWithDefaults() *SubscriptionPriceResponse {
	this := SubscriptionPriceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionPriceResponse) GetData() SubscriptionPrice {
	if o == nil {
		var ret SubscriptionPrice
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceResponse) GetDataOk() (*SubscriptionPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionPriceResponse) SetData(v SubscriptionPrice) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionPriceResponse) GetIncluded() []SubscriptionOfferCodePricesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionOfferCodePricesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceResponse) GetIncludedOk() ([]SubscriptionOfferCodePricesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionPriceResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionOfferCodePricesResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionPriceResponse) SetIncluded(v []SubscriptionOfferCodePricesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionPriceResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionPriceResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *SubscriptionPriceResponse) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionPriceResponse := _SubscriptionPriceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPriceResponse)

	if err != nil {
		return err
	}

	*o = SubscriptionPriceResponse(varSubscriptionPriceResponse)

	return err
}

type NullableSubscriptionPriceResponse struct {
	value *SubscriptionPriceResponse
	isSet bool
}

func (v NullableSubscriptionPriceResponse) Get() *SubscriptionPriceResponse {
	return v.value
}

func (v *NullableSubscriptionPriceResponse) Set(val *SubscriptionPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPriceResponse(val *SubscriptionPriceResponse) *NullableSubscriptionPriceResponse {
	return &NullableSubscriptionPriceResponse{value: val, isSet: true}
}

func (v NullableSubscriptionPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


