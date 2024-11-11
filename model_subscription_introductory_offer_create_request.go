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

// checks if the SubscriptionIntroductoryOfferCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferCreateRequest{}

// SubscriptionIntroductoryOfferCreateRequest struct for SubscriptionIntroductoryOfferCreateRequest
type SubscriptionIntroductoryOfferCreateRequest struct {
	Data SubscriptionIntroductoryOfferCreateRequestData `json:"data"`
	Included []SubscriptionPricePointInlineCreate `json:"included,omitempty"`
}

type _SubscriptionIntroductoryOfferCreateRequest SubscriptionIntroductoryOfferCreateRequest

// NewSubscriptionIntroductoryOfferCreateRequest instantiates a new SubscriptionIntroductoryOfferCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferCreateRequest(data SubscriptionIntroductoryOfferCreateRequestData) *SubscriptionIntroductoryOfferCreateRequest {
	this := SubscriptionIntroductoryOfferCreateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionIntroductoryOfferCreateRequestWithDefaults instantiates a new SubscriptionIntroductoryOfferCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferCreateRequestWithDefaults() *SubscriptionIntroductoryOfferCreateRequest {
	this := SubscriptionIntroductoryOfferCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionIntroductoryOfferCreateRequest) GetData() SubscriptionIntroductoryOfferCreateRequestData {
	if o == nil {
		var ret SubscriptionIntroductoryOfferCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferCreateRequest) GetDataOk() (*SubscriptionIntroductoryOfferCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionIntroductoryOfferCreateRequest) SetData(v SubscriptionIntroductoryOfferCreateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferCreateRequest) GetIncluded() []SubscriptionPricePointInlineCreate {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionPricePointInlineCreate
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferCreateRequest) GetIncludedOk() ([]SubscriptionPricePointInlineCreate, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferCreateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionPricePointInlineCreate and assigns it to the Included field.
func (o *SubscriptionIntroductoryOfferCreateRequest) SetIncluded(v []SubscriptionPricePointInlineCreate) {
	o.Included = v
}

func (o SubscriptionIntroductoryOfferCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *SubscriptionIntroductoryOfferCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionIntroductoryOfferCreateRequest := _SubscriptionIntroductoryOfferCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionIntroductoryOfferCreateRequest)

	if err != nil {
		return err
	}

	*o = SubscriptionIntroductoryOfferCreateRequest(varSubscriptionIntroductoryOfferCreateRequest)

	return err
}

type NullableSubscriptionIntroductoryOfferCreateRequest struct {
	value *SubscriptionIntroductoryOfferCreateRequest
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferCreateRequest) Get() *SubscriptionIntroductoryOfferCreateRequest {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequest) Set(val *SubscriptionIntroductoryOfferCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferCreateRequest(val *SubscriptionIntroductoryOfferCreateRequest) *NullableSubscriptionIntroductoryOfferCreateRequest {
	return &NullableSubscriptionIntroductoryOfferCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


