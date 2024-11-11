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

// checks if the ReviewSubmissionUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionUpdateRequest{}

// ReviewSubmissionUpdateRequest struct for ReviewSubmissionUpdateRequest
type ReviewSubmissionUpdateRequest struct {
	Data ReviewSubmissionUpdateRequestData `json:"data"`
}

type _ReviewSubmissionUpdateRequest ReviewSubmissionUpdateRequest

// NewReviewSubmissionUpdateRequest instantiates a new ReviewSubmissionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionUpdateRequest(data ReviewSubmissionUpdateRequestData) *ReviewSubmissionUpdateRequest {
	this := ReviewSubmissionUpdateRequest{}
	this.Data = data
	return &this
}

// NewReviewSubmissionUpdateRequestWithDefaults instantiates a new ReviewSubmissionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionUpdateRequestWithDefaults() *ReviewSubmissionUpdateRequest {
	this := ReviewSubmissionUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *ReviewSubmissionUpdateRequest) GetData() ReviewSubmissionUpdateRequestData {
	if o == nil {
		var ret ReviewSubmissionUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionUpdateRequest) GetDataOk() (*ReviewSubmissionUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReviewSubmissionUpdateRequest) SetData(v ReviewSubmissionUpdateRequestData) {
	o.Data = v
}

func (o ReviewSubmissionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ReviewSubmissionUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varReviewSubmissionUpdateRequest := _ReviewSubmissionUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReviewSubmissionUpdateRequest)

	if err != nil {
		return err
	}

	*o = ReviewSubmissionUpdateRequest(varReviewSubmissionUpdateRequest)

	return err
}

type NullableReviewSubmissionUpdateRequest struct {
	value *ReviewSubmissionUpdateRequest
	isSet bool
}

func (v NullableReviewSubmissionUpdateRequest) Get() *ReviewSubmissionUpdateRequest {
	return v.value
}

func (v *NullableReviewSubmissionUpdateRequest) Set(val *ReviewSubmissionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionUpdateRequest(val *ReviewSubmissionUpdateRequest) *NullableReviewSubmissionUpdateRequest {
	return &NullableReviewSubmissionUpdateRequest{value: val, isSet: true}
}

func (v NullableReviewSubmissionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

