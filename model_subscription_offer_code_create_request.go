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

// checks if the SubscriptionOfferCodeCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCreateRequest{}

// SubscriptionOfferCodeCreateRequest struct for SubscriptionOfferCodeCreateRequest
type SubscriptionOfferCodeCreateRequest struct {
	Data SubscriptionOfferCodeCreateRequestData `json:"data"`
	Included []SubscriptionOfferCodePriceInlineCreate `json:"included,omitempty"`
}

type _SubscriptionOfferCodeCreateRequest SubscriptionOfferCodeCreateRequest

// NewSubscriptionOfferCodeCreateRequest instantiates a new SubscriptionOfferCodeCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCreateRequest(data SubscriptionOfferCodeCreateRequestData) *SubscriptionOfferCodeCreateRequest {
	this := SubscriptionOfferCodeCreateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionOfferCodeCreateRequestWithDefaults instantiates a new SubscriptionOfferCodeCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCreateRequestWithDefaults() *SubscriptionOfferCodeCreateRequest {
	this := SubscriptionOfferCodeCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionOfferCodeCreateRequest) GetData() SubscriptionOfferCodeCreateRequestData {
	if o == nil {
		var ret SubscriptionOfferCodeCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequest) GetDataOk() (*SubscriptionOfferCodeCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionOfferCodeCreateRequest) SetData(v SubscriptionOfferCodeCreateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCreateRequest) GetIncluded() []SubscriptionOfferCodePriceInlineCreate {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionOfferCodePriceInlineCreate
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequest) GetIncludedOk() ([]SubscriptionOfferCodePriceInlineCreate, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCreateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionOfferCodePriceInlineCreate and assigns it to the Included field.
func (o *SubscriptionOfferCodeCreateRequest) SetIncluded(v []SubscriptionOfferCodePriceInlineCreate) {
	o.Included = v
}

func (o SubscriptionOfferCodeCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *SubscriptionOfferCodeCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionOfferCodeCreateRequest := _SubscriptionOfferCodeCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionOfferCodeCreateRequest)

	if err != nil {
		return err
	}

	*o = SubscriptionOfferCodeCreateRequest(varSubscriptionOfferCodeCreateRequest)

	return err
}

type NullableSubscriptionOfferCodeCreateRequest struct {
	value *SubscriptionOfferCodeCreateRequest
	isSet bool
}

func (v NullableSubscriptionOfferCodeCreateRequest) Get() *SubscriptionOfferCodeCreateRequest {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCreateRequest) Set(val *SubscriptionOfferCodeCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCreateRequest(val *SubscriptionOfferCodeCreateRequest) *NullableSubscriptionOfferCodeCreateRequest {
	return &NullableSubscriptionOfferCodeCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


