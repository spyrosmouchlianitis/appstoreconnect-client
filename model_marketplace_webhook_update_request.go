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

// checks if the MarketplaceWebhookUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceWebhookUpdateRequest{}

// MarketplaceWebhookUpdateRequest struct for MarketplaceWebhookUpdateRequest
type MarketplaceWebhookUpdateRequest struct {
	Data MarketplaceWebhookUpdateRequestData `json:"data"`
}

type _MarketplaceWebhookUpdateRequest MarketplaceWebhookUpdateRequest

// NewMarketplaceWebhookUpdateRequest instantiates a new MarketplaceWebhookUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceWebhookUpdateRequest(data MarketplaceWebhookUpdateRequestData) *MarketplaceWebhookUpdateRequest {
	this := MarketplaceWebhookUpdateRequest{}
	this.Data = data
	return &this
}

// NewMarketplaceWebhookUpdateRequestWithDefaults instantiates a new MarketplaceWebhookUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceWebhookUpdateRequestWithDefaults() *MarketplaceWebhookUpdateRequest {
	this := MarketplaceWebhookUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *MarketplaceWebhookUpdateRequest) GetData() MarketplaceWebhookUpdateRequestData {
	if o == nil {
		var ret MarketplaceWebhookUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MarketplaceWebhookUpdateRequest) GetDataOk() (*MarketplaceWebhookUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MarketplaceWebhookUpdateRequest) SetData(v MarketplaceWebhookUpdateRequestData) {
	o.Data = v
}

func (o MarketplaceWebhookUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceWebhookUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *MarketplaceWebhookUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varMarketplaceWebhookUpdateRequest := _MarketplaceWebhookUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketplaceWebhookUpdateRequest)

	if err != nil {
		return err
	}

	*o = MarketplaceWebhookUpdateRequest(varMarketplaceWebhookUpdateRequest)

	return err
}

type NullableMarketplaceWebhookUpdateRequest struct {
	value *MarketplaceWebhookUpdateRequest
	isSet bool
}

func (v NullableMarketplaceWebhookUpdateRequest) Get() *MarketplaceWebhookUpdateRequest {
	return v.value
}

func (v *NullableMarketplaceWebhookUpdateRequest) Set(val *MarketplaceWebhookUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceWebhookUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceWebhookUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceWebhookUpdateRequest(val *MarketplaceWebhookUpdateRequest) *NullableMarketplaceWebhookUpdateRequest {
	return &NullableMarketplaceWebhookUpdateRequest{value: val, isSet: true}
}

func (v NullableMarketplaceWebhookUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceWebhookUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


