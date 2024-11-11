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

// checks if the BuildBetaDetailUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaDetailUpdateRequest{}

// BuildBetaDetailUpdateRequest struct for BuildBetaDetailUpdateRequest
type BuildBetaDetailUpdateRequest struct {
	Data BuildBetaDetailUpdateRequestData `json:"data"`
}

type _BuildBetaDetailUpdateRequest BuildBetaDetailUpdateRequest

// NewBuildBetaDetailUpdateRequest instantiates a new BuildBetaDetailUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaDetailUpdateRequest(data BuildBetaDetailUpdateRequestData) *BuildBetaDetailUpdateRequest {
	this := BuildBetaDetailUpdateRequest{}
	this.Data = data
	return &this
}

// NewBuildBetaDetailUpdateRequestWithDefaults instantiates a new BuildBetaDetailUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaDetailUpdateRequestWithDefaults() *BuildBetaDetailUpdateRequest {
	this := BuildBetaDetailUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BuildBetaDetailUpdateRequest) GetData() BuildBetaDetailUpdateRequestData {
	if o == nil {
		var ret BuildBetaDetailUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailUpdateRequest) GetDataOk() (*BuildBetaDetailUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildBetaDetailUpdateRequest) SetData(v BuildBetaDetailUpdateRequestData) {
	o.Data = v
}

func (o BuildBetaDetailUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaDetailUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BuildBetaDetailUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBuildBetaDetailUpdateRequest := _BuildBetaDetailUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildBetaDetailUpdateRequest)

	if err != nil {
		return err
	}

	*o = BuildBetaDetailUpdateRequest(varBuildBetaDetailUpdateRequest)

	return err
}

type NullableBuildBetaDetailUpdateRequest struct {
	value *BuildBetaDetailUpdateRequest
	isSet bool
}

func (v NullableBuildBetaDetailUpdateRequest) Get() *BuildBetaDetailUpdateRequest {
	return v.value
}

func (v *NullableBuildBetaDetailUpdateRequest) Set(val *BuildBetaDetailUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaDetailUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaDetailUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaDetailUpdateRequest(val *BuildBetaDetailUpdateRequest) *NullableBuildBetaDetailUpdateRequest {
	return &NullableBuildBetaDetailUpdateRequest{value: val, isSet: true}
}

func (v NullableBuildBetaDetailUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaDetailUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

