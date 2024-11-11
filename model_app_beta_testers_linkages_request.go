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

// checks if the AppBetaTestersLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppBetaTestersLinkagesRequest{}

// AppBetaTestersLinkagesRequest struct for AppBetaTestersLinkagesRequest
type AppBetaTestersLinkagesRequest struct {
	Data []BetaGroupRelationshipsBetaTestersDataInner `json:"data"`
}

type _AppBetaTestersLinkagesRequest AppBetaTestersLinkagesRequest

// NewAppBetaTestersLinkagesRequest instantiates a new AppBetaTestersLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppBetaTestersLinkagesRequest(data []BetaGroupRelationshipsBetaTestersDataInner) *AppBetaTestersLinkagesRequest {
	this := AppBetaTestersLinkagesRequest{}
	this.Data = data
	return &this
}

// NewAppBetaTestersLinkagesRequestWithDefaults instantiates a new AppBetaTestersLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppBetaTestersLinkagesRequestWithDefaults() *AppBetaTestersLinkagesRequest {
	this := AppBetaTestersLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppBetaTestersLinkagesRequest) GetData() []BetaGroupRelationshipsBetaTestersDataInner {
	if o == nil {
		var ret []BetaGroupRelationshipsBetaTestersDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppBetaTestersLinkagesRequest) GetDataOk() ([]BetaGroupRelationshipsBetaTestersDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppBetaTestersLinkagesRequest) SetData(v []BetaGroupRelationshipsBetaTestersDataInner) {
	o.Data = v
}

func (o AppBetaTestersLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppBetaTestersLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppBetaTestersLinkagesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppBetaTestersLinkagesRequest := _AppBetaTestersLinkagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppBetaTestersLinkagesRequest)

	if err != nil {
		return err
	}

	*o = AppBetaTestersLinkagesRequest(varAppBetaTestersLinkagesRequest)

	return err
}

type NullableAppBetaTestersLinkagesRequest struct {
	value *AppBetaTestersLinkagesRequest
	isSet bool
}

func (v NullableAppBetaTestersLinkagesRequest) Get() *AppBetaTestersLinkagesRequest {
	return v.value
}

func (v *NullableAppBetaTestersLinkagesRequest) Set(val *AppBetaTestersLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppBetaTestersLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppBetaTestersLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppBetaTestersLinkagesRequest(val *AppBetaTestersLinkagesRequest) *NullableAppBetaTestersLinkagesRequest {
	return &NullableAppBetaTestersLinkagesRequest{value: val, isSet: true}
}

func (v NullableAppBetaTestersLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppBetaTestersLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


