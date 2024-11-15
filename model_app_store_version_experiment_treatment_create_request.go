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

// checks if the AppStoreVersionExperimentTreatmentCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentCreateRequest{}

// AppStoreVersionExperimentTreatmentCreateRequest struct for AppStoreVersionExperimentTreatmentCreateRequest
type AppStoreVersionExperimentTreatmentCreateRequest struct {
	Data AppStoreVersionExperimentTreatmentCreateRequestData `json:"data"`
}

type _AppStoreVersionExperimentTreatmentCreateRequest AppStoreVersionExperimentTreatmentCreateRequest

// NewAppStoreVersionExperimentTreatmentCreateRequest instantiates a new AppStoreVersionExperimentTreatmentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentCreateRequest(data AppStoreVersionExperimentTreatmentCreateRequestData) *AppStoreVersionExperimentTreatmentCreateRequest {
	this := AppStoreVersionExperimentTreatmentCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionExperimentTreatmentCreateRequestWithDefaults instantiates a new AppStoreVersionExperimentTreatmentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentCreateRequestWithDefaults() *AppStoreVersionExperimentTreatmentCreateRequest {
	this := AppStoreVersionExperimentTreatmentCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionExperimentTreatmentCreateRequest) GetData() AppStoreVersionExperimentTreatmentCreateRequestData {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequest) GetDataOk() (*AppStoreVersionExperimentTreatmentCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionExperimentTreatmentCreateRequest) SetData(v AppStoreVersionExperimentTreatmentCreateRequestData) {
	o.Data = v
}

func (o AppStoreVersionExperimentTreatmentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppStoreVersionExperimentTreatmentCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionExperimentTreatmentCreateRequest := _AppStoreVersionExperimentTreatmentCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionExperimentTreatmentCreateRequest)

	if err != nil {
		return err
	}

	*o = AppStoreVersionExperimentTreatmentCreateRequest(varAppStoreVersionExperimentTreatmentCreateRequest)

	return err
}

type NullableAppStoreVersionExperimentTreatmentCreateRequest struct {
	value *AppStoreVersionExperimentTreatmentCreateRequest
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequest) Get() *AppStoreVersionExperimentTreatmentCreateRequest {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequest) Set(val *AppStoreVersionExperimentTreatmentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentCreateRequest(val *AppStoreVersionExperimentTreatmentCreateRequest) *NullableAppStoreVersionExperimentTreatmentCreateRequest {
	return &NullableAppStoreVersionExperimentTreatmentCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


