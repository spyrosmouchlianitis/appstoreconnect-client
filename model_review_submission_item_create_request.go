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

// checks if the ReviewSubmissionItemCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionItemCreateRequest{}

// ReviewSubmissionItemCreateRequest struct for ReviewSubmissionItemCreateRequest
type ReviewSubmissionItemCreateRequest struct {
	Data ReviewSubmissionItemCreateRequestData `json:"data"`
}

type _ReviewSubmissionItemCreateRequest ReviewSubmissionItemCreateRequest

// NewReviewSubmissionItemCreateRequest instantiates a new ReviewSubmissionItemCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionItemCreateRequest(data ReviewSubmissionItemCreateRequestData) *ReviewSubmissionItemCreateRequest {
	this := ReviewSubmissionItemCreateRequest{}
	this.Data = data
	return &this
}

// NewReviewSubmissionItemCreateRequestWithDefaults instantiates a new ReviewSubmissionItemCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionItemCreateRequestWithDefaults() *ReviewSubmissionItemCreateRequest {
	this := ReviewSubmissionItemCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *ReviewSubmissionItemCreateRequest) GetData() ReviewSubmissionItemCreateRequestData {
	if o == nil {
		var ret ReviewSubmissionItemCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemCreateRequest) GetDataOk() (*ReviewSubmissionItemCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReviewSubmissionItemCreateRequest) SetData(v ReviewSubmissionItemCreateRequestData) {
	o.Data = v
}

func (o ReviewSubmissionItemCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionItemCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ReviewSubmissionItemCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varReviewSubmissionItemCreateRequest := _ReviewSubmissionItemCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReviewSubmissionItemCreateRequest)

	if err != nil {
		return err
	}

	*o = ReviewSubmissionItemCreateRequest(varReviewSubmissionItemCreateRequest)

	return err
}

type NullableReviewSubmissionItemCreateRequest struct {
	value *ReviewSubmissionItemCreateRequest
	isSet bool
}

func (v NullableReviewSubmissionItemCreateRequest) Get() *ReviewSubmissionItemCreateRequest {
	return v.value
}

func (v *NullableReviewSubmissionItemCreateRequest) Set(val *ReviewSubmissionItemCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionItemCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionItemCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionItemCreateRequest(val *ReviewSubmissionItemCreateRequest) *NullableReviewSubmissionItemCreateRequest {
	return &NullableReviewSubmissionItemCreateRequest{value: val, isSet: true}
}

func (v NullableReviewSubmissionItemCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionItemCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


