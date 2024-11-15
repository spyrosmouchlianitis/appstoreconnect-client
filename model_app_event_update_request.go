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

// checks if the AppEventUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventUpdateRequest{}

// AppEventUpdateRequest struct for AppEventUpdateRequest
type AppEventUpdateRequest struct {
	Data AppEventUpdateRequestData `json:"data"`
}

type _AppEventUpdateRequest AppEventUpdateRequest

// NewAppEventUpdateRequest instantiates a new AppEventUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventUpdateRequest(data AppEventUpdateRequestData) *AppEventUpdateRequest {
	this := AppEventUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppEventUpdateRequestWithDefaults instantiates a new AppEventUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventUpdateRequestWithDefaults() *AppEventUpdateRequest {
	this := AppEventUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppEventUpdateRequest) GetData() AppEventUpdateRequestData {
	if o == nil {
		var ret AppEventUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEventUpdateRequest) GetDataOk() (*AppEventUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEventUpdateRequest) SetData(v AppEventUpdateRequestData) {
	o.Data = v
}

func (o AppEventUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppEventUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppEventUpdateRequest := _AppEventUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEventUpdateRequest)

	if err != nil {
		return err
	}

	*o = AppEventUpdateRequest(varAppEventUpdateRequest)

	return err
}

type NullableAppEventUpdateRequest struct {
	value *AppEventUpdateRequest
	isSet bool
}

func (v NullableAppEventUpdateRequest) Get() *AppEventUpdateRequest {
	return v.value
}

func (v *NullableAppEventUpdateRequest) Set(val *AppEventUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventUpdateRequest(val *AppEventUpdateRequest) *NullableAppEventUpdateRequest {
	return &NullableAppEventUpdateRequest{value: val, isSet: true}
}

func (v NullableAppEventUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


