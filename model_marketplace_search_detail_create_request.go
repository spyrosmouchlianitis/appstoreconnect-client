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

// checks if the MarketplaceSearchDetailCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceSearchDetailCreateRequest{}

// MarketplaceSearchDetailCreateRequest struct for MarketplaceSearchDetailCreateRequest
type MarketplaceSearchDetailCreateRequest struct {
	Data MarketplaceSearchDetailCreateRequestData `json:"data"`
}

type _MarketplaceSearchDetailCreateRequest MarketplaceSearchDetailCreateRequest

// NewMarketplaceSearchDetailCreateRequest instantiates a new MarketplaceSearchDetailCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceSearchDetailCreateRequest(data MarketplaceSearchDetailCreateRequestData) *MarketplaceSearchDetailCreateRequest {
	this := MarketplaceSearchDetailCreateRequest{}
	this.Data = data
	return &this
}

// NewMarketplaceSearchDetailCreateRequestWithDefaults instantiates a new MarketplaceSearchDetailCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceSearchDetailCreateRequestWithDefaults() *MarketplaceSearchDetailCreateRequest {
	this := MarketplaceSearchDetailCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *MarketplaceSearchDetailCreateRequest) GetData() MarketplaceSearchDetailCreateRequestData {
	if o == nil {
		var ret MarketplaceSearchDetailCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MarketplaceSearchDetailCreateRequest) GetDataOk() (*MarketplaceSearchDetailCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MarketplaceSearchDetailCreateRequest) SetData(v MarketplaceSearchDetailCreateRequestData) {
	o.Data = v
}

func (o MarketplaceSearchDetailCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceSearchDetailCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *MarketplaceSearchDetailCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varMarketplaceSearchDetailCreateRequest := _MarketplaceSearchDetailCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketplaceSearchDetailCreateRequest)

	if err != nil {
		return err
	}

	*o = MarketplaceSearchDetailCreateRequest(varMarketplaceSearchDetailCreateRequest)

	return err
}

type NullableMarketplaceSearchDetailCreateRequest struct {
	value *MarketplaceSearchDetailCreateRequest
	isSet bool
}

func (v NullableMarketplaceSearchDetailCreateRequest) Get() *MarketplaceSearchDetailCreateRequest {
	return v.value
}

func (v *NullableMarketplaceSearchDetailCreateRequest) Set(val *MarketplaceSearchDetailCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceSearchDetailCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceSearchDetailCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceSearchDetailCreateRequest(val *MarketplaceSearchDetailCreateRequest) *NullableMarketplaceSearchDetailCreateRequest {
	return &NullableMarketplaceSearchDetailCreateRequest{value: val, isSet: true}
}

func (v NullableMarketplaceSearchDetailCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceSearchDetailCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


