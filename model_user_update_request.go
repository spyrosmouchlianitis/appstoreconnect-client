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

// checks if the UserUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUpdateRequest{}

// UserUpdateRequest struct for UserUpdateRequest
type UserUpdateRequest struct {
	Data UserUpdateRequestData `json:"data"`
}

type _UserUpdateRequest UserUpdateRequest

// NewUserUpdateRequest instantiates a new UserUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateRequest(data UserUpdateRequestData) *UserUpdateRequest {
	this := UserUpdateRequest{}
	this.Data = data
	return &this
}

// NewUserUpdateRequestWithDefaults instantiates a new UserUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateRequestWithDefaults() *UserUpdateRequest {
	this := UserUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *UserUpdateRequest) GetData() UserUpdateRequestData {
	if o == nil {
		var ret UserUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserUpdateRequest) GetDataOk() (*UserUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserUpdateRequest) SetData(v UserUpdateRequestData) {
	o.Data = v
}

func (o UserUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UserUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUserUpdateRequest := _UserUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserUpdateRequest)

	if err != nil {
		return err
	}

	*o = UserUpdateRequest(varUserUpdateRequest)

	return err
}

type NullableUserUpdateRequest struct {
	value *UserUpdateRequest
	isSet bool
}

func (v NullableUserUpdateRequest) Get() *UserUpdateRequest {
	return v.value
}

func (v *NullableUserUpdateRequest) Set(val *UserUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateRequest(val *UserUpdateRequest) *NullableUserUpdateRequest {
	return &NullableUserUpdateRequest{value: val, isSet: true}
}

func (v NullableUserUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


