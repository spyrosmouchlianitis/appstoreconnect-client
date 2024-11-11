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

// checks if the MarketplaceDomainCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceDomainCreateRequest{}

// MarketplaceDomainCreateRequest struct for MarketplaceDomainCreateRequest
type MarketplaceDomainCreateRequest struct {
	Data MarketplaceDomainCreateRequestData `json:"data"`
}

type _MarketplaceDomainCreateRequest MarketplaceDomainCreateRequest

// NewMarketplaceDomainCreateRequest instantiates a new MarketplaceDomainCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceDomainCreateRequest(data MarketplaceDomainCreateRequestData) *MarketplaceDomainCreateRequest {
	this := MarketplaceDomainCreateRequest{}
	this.Data = data
	return &this
}

// NewMarketplaceDomainCreateRequestWithDefaults instantiates a new MarketplaceDomainCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceDomainCreateRequestWithDefaults() *MarketplaceDomainCreateRequest {
	this := MarketplaceDomainCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *MarketplaceDomainCreateRequest) GetData() MarketplaceDomainCreateRequestData {
	if o == nil {
		var ret MarketplaceDomainCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MarketplaceDomainCreateRequest) GetDataOk() (*MarketplaceDomainCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MarketplaceDomainCreateRequest) SetData(v MarketplaceDomainCreateRequestData) {
	o.Data = v
}

func (o MarketplaceDomainCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceDomainCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *MarketplaceDomainCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varMarketplaceDomainCreateRequest := _MarketplaceDomainCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketplaceDomainCreateRequest)

	if err != nil {
		return err
	}

	*o = MarketplaceDomainCreateRequest(varMarketplaceDomainCreateRequest)

	return err
}

type NullableMarketplaceDomainCreateRequest struct {
	value *MarketplaceDomainCreateRequest
	isSet bool
}

func (v NullableMarketplaceDomainCreateRequest) Get() *MarketplaceDomainCreateRequest {
	return v.value
}

func (v *NullableMarketplaceDomainCreateRequest) Set(val *MarketplaceDomainCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceDomainCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceDomainCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceDomainCreateRequest(val *MarketplaceDomainCreateRequest) *NullableMarketplaceDomainCreateRequest {
	return &NullableMarketplaceDomainCreateRequest{value: val, isSet: true}
}

func (v NullableMarketplaceDomainCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceDomainCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


