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

// checks if the AppCustomProductPageCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageCreateRequest{}

// AppCustomProductPageCreateRequest struct for AppCustomProductPageCreateRequest
type AppCustomProductPageCreateRequest struct {
	Data AppCustomProductPageCreateRequestData `json:"data"`
	Included []AppCustomProductPageCreateRequestIncludedInner `json:"included,omitempty"`
}

type _AppCustomProductPageCreateRequest AppCustomProductPageCreateRequest

// NewAppCustomProductPageCreateRequest instantiates a new AppCustomProductPageCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageCreateRequest(data AppCustomProductPageCreateRequestData) *AppCustomProductPageCreateRequest {
	this := AppCustomProductPageCreateRequest{}
	this.Data = data
	return &this
}

// NewAppCustomProductPageCreateRequestWithDefaults instantiates a new AppCustomProductPageCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageCreateRequestWithDefaults() *AppCustomProductPageCreateRequest {
	this := AppCustomProductPageCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppCustomProductPageCreateRequest) GetData() AppCustomProductPageCreateRequestData {
	if o == nil {
		var ret AppCustomProductPageCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequest) GetDataOk() (*AppCustomProductPageCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppCustomProductPageCreateRequest) SetData(v AppCustomProductPageCreateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppCustomProductPageCreateRequest) GetIncluded() []AppCustomProductPageCreateRequestIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppCustomProductPageCreateRequestIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequest) GetIncludedOk() ([]AppCustomProductPageCreateRequestIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppCustomProductPageCreateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppCustomProductPageCreateRequestIncludedInner and assigns it to the Included field.
func (o *AppCustomProductPageCreateRequest) SetIncluded(v []AppCustomProductPageCreateRequestIncludedInner) {
	o.Included = v
}

func (o AppCustomProductPageCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *AppCustomProductPageCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppCustomProductPageCreateRequest := _AppCustomProductPageCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppCustomProductPageCreateRequest)

	if err != nil {
		return err
	}

	*o = AppCustomProductPageCreateRequest(varAppCustomProductPageCreateRequest)

	return err
}

type NullableAppCustomProductPageCreateRequest struct {
	value *AppCustomProductPageCreateRequest
	isSet bool
}

func (v NullableAppCustomProductPageCreateRequest) Get() *AppCustomProductPageCreateRequest {
	return v.value
}

func (v *NullableAppCustomProductPageCreateRequest) Set(val *AppCustomProductPageCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageCreateRequest(val *AppCustomProductPageCreateRequest) *NullableAppCustomProductPageCreateRequest {
	return &NullableAppCustomProductPageCreateRequest{value: val, isSet: true}
}

func (v NullableAppCustomProductPageCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


