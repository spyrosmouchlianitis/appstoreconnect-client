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

// checks if the AgeRatingDeclarationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgeRatingDeclarationUpdateRequest{}

// AgeRatingDeclarationUpdateRequest struct for AgeRatingDeclarationUpdateRequest
type AgeRatingDeclarationUpdateRequest struct {
	Data AgeRatingDeclarationUpdateRequestData `json:"data"`
}

type _AgeRatingDeclarationUpdateRequest AgeRatingDeclarationUpdateRequest

// NewAgeRatingDeclarationUpdateRequest instantiates a new AgeRatingDeclarationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgeRatingDeclarationUpdateRequest(data AgeRatingDeclarationUpdateRequestData) *AgeRatingDeclarationUpdateRequest {
	this := AgeRatingDeclarationUpdateRequest{}
	this.Data = data
	return &this
}

// NewAgeRatingDeclarationUpdateRequestWithDefaults instantiates a new AgeRatingDeclarationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeRatingDeclarationUpdateRequestWithDefaults() *AgeRatingDeclarationUpdateRequest {
	this := AgeRatingDeclarationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AgeRatingDeclarationUpdateRequest) GetData() AgeRatingDeclarationUpdateRequestData {
	if o == nil {
		var ret AgeRatingDeclarationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationUpdateRequest) GetDataOk() (*AgeRatingDeclarationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgeRatingDeclarationUpdateRequest) SetData(v AgeRatingDeclarationUpdateRequestData) {
	o.Data = v
}

func (o AgeRatingDeclarationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgeRatingDeclarationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AgeRatingDeclarationUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAgeRatingDeclarationUpdateRequest := _AgeRatingDeclarationUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgeRatingDeclarationUpdateRequest)

	if err != nil {
		return err
	}

	*o = AgeRatingDeclarationUpdateRequest(varAgeRatingDeclarationUpdateRequest)

	return err
}

type NullableAgeRatingDeclarationUpdateRequest struct {
	value *AgeRatingDeclarationUpdateRequest
	isSet bool
}

func (v NullableAgeRatingDeclarationUpdateRequest) Get() *AgeRatingDeclarationUpdateRequest {
	return v.value
}

func (v *NullableAgeRatingDeclarationUpdateRequest) Set(val *AgeRatingDeclarationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgeRatingDeclarationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgeRatingDeclarationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgeRatingDeclarationUpdateRequest(val *AgeRatingDeclarationUpdateRequest) *NullableAgeRatingDeclarationUpdateRequest {
	return &NullableAgeRatingDeclarationUpdateRequest{value: val, isSet: true}
}

func (v NullableAgeRatingDeclarationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgeRatingDeclarationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


