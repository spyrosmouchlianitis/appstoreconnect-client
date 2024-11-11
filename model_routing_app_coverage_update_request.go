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

// checks if the RoutingAppCoverageUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingAppCoverageUpdateRequest{}

// RoutingAppCoverageUpdateRequest struct for RoutingAppCoverageUpdateRequest
type RoutingAppCoverageUpdateRequest struct {
	Data RoutingAppCoverageUpdateRequestData `json:"data"`
}

type _RoutingAppCoverageUpdateRequest RoutingAppCoverageUpdateRequest

// NewRoutingAppCoverageUpdateRequest instantiates a new RoutingAppCoverageUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingAppCoverageUpdateRequest(data RoutingAppCoverageUpdateRequestData) *RoutingAppCoverageUpdateRequest {
	this := RoutingAppCoverageUpdateRequest{}
	this.Data = data
	return &this
}

// NewRoutingAppCoverageUpdateRequestWithDefaults instantiates a new RoutingAppCoverageUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingAppCoverageUpdateRequestWithDefaults() *RoutingAppCoverageUpdateRequest {
	this := RoutingAppCoverageUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *RoutingAppCoverageUpdateRequest) GetData() RoutingAppCoverageUpdateRequestData {
	if o == nil {
		var ret RoutingAppCoverageUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverageUpdateRequest) GetDataOk() (*RoutingAppCoverageUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RoutingAppCoverageUpdateRequest) SetData(v RoutingAppCoverageUpdateRequestData) {
	o.Data = v
}

func (o RoutingAppCoverageUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingAppCoverageUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RoutingAppCoverageUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varRoutingAppCoverageUpdateRequest := _RoutingAppCoverageUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoutingAppCoverageUpdateRequest)

	if err != nil {
		return err
	}

	*o = RoutingAppCoverageUpdateRequest(varRoutingAppCoverageUpdateRequest)

	return err
}

type NullableRoutingAppCoverageUpdateRequest struct {
	value *RoutingAppCoverageUpdateRequest
	isSet bool
}

func (v NullableRoutingAppCoverageUpdateRequest) Get() *RoutingAppCoverageUpdateRequest {
	return v.value
}

func (v *NullableRoutingAppCoverageUpdateRequest) Set(val *RoutingAppCoverageUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingAppCoverageUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingAppCoverageUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingAppCoverageUpdateRequest(val *RoutingAppCoverageUpdateRequest) *NullableRoutingAppCoverageUpdateRequest {
	return &NullableRoutingAppCoverageUpdateRequest{value: val, isSet: true}
}

func (v NullableRoutingAppCoverageUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingAppCoverageUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


