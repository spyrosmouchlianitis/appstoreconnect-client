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

// checks if the AppEventCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventCreateRequest{}

// AppEventCreateRequest struct for AppEventCreateRequest
type AppEventCreateRequest struct {
	Data AppEventCreateRequestData `json:"data"`
}

type _AppEventCreateRequest AppEventCreateRequest

// NewAppEventCreateRequest instantiates a new AppEventCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventCreateRequest(data AppEventCreateRequestData) *AppEventCreateRequest {
	this := AppEventCreateRequest{}
	this.Data = data
	return &this
}

// NewAppEventCreateRequestWithDefaults instantiates a new AppEventCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventCreateRequestWithDefaults() *AppEventCreateRequest {
	this := AppEventCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppEventCreateRequest) GetData() AppEventCreateRequestData {
	if o == nil {
		var ret AppEventCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEventCreateRequest) GetDataOk() (*AppEventCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEventCreateRequest) SetData(v AppEventCreateRequestData) {
	o.Data = v
}

func (o AppEventCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppEventCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppEventCreateRequest := _AppEventCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEventCreateRequest)

	if err != nil {
		return err
	}

	*o = AppEventCreateRequest(varAppEventCreateRequest)

	return err
}

type NullableAppEventCreateRequest struct {
	value *AppEventCreateRequest
	isSet bool
}

func (v NullableAppEventCreateRequest) Get() *AppEventCreateRequest {
	return v.value
}

func (v *NullableAppEventCreateRequest) Set(val *AppEventCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventCreateRequest(val *AppEventCreateRequest) *NullableAppEventCreateRequest {
	return &NullableAppEventCreateRequest{value: val, isSet: true}
}

func (v NullableAppEventCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


