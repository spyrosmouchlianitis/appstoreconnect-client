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

// checks if the BetaAppReviewSubmissionWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppReviewSubmissionWithoutIncludesResponse{}

// BetaAppReviewSubmissionWithoutIncludesResponse struct for BetaAppReviewSubmissionWithoutIncludesResponse
type BetaAppReviewSubmissionWithoutIncludesResponse struct {
	Data BetaAppReviewSubmission `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _BetaAppReviewSubmissionWithoutIncludesResponse BetaAppReviewSubmissionWithoutIncludesResponse

// NewBetaAppReviewSubmissionWithoutIncludesResponse instantiates a new BetaAppReviewSubmissionWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppReviewSubmissionWithoutIncludesResponse(data BetaAppReviewSubmission, links DocumentLinks) *BetaAppReviewSubmissionWithoutIncludesResponse {
	this := BetaAppReviewSubmissionWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaAppReviewSubmissionWithoutIncludesResponseWithDefaults instantiates a new BetaAppReviewSubmissionWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppReviewSubmissionWithoutIncludesResponseWithDefaults() *BetaAppReviewSubmissionWithoutIncludesResponse {
	this := BetaAppReviewSubmissionWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppReviewSubmissionWithoutIncludesResponse) GetData() BetaAppReviewSubmission {
	if o == nil {
		var ret BetaAppReviewSubmission
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmissionWithoutIncludesResponse) GetDataOk() (*BetaAppReviewSubmission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppReviewSubmissionWithoutIncludesResponse) SetData(v BetaAppReviewSubmission) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaAppReviewSubmissionWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmissionWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaAppReviewSubmissionWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaAppReviewSubmissionWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppReviewSubmissionWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *BetaAppReviewSubmissionWithoutIncludesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varBetaAppReviewSubmissionWithoutIncludesResponse := _BetaAppReviewSubmissionWithoutIncludesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaAppReviewSubmissionWithoutIncludesResponse)

	if err != nil {
		return err
	}

	*o = BetaAppReviewSubmissionWithoutIncludesResponse(varBetaAppReviewSubmissionWithoutIncludesResponse)

	return err
}

type NullableBetaAppReviewSubmissionWithoutIncludesResponse struct {
	value *BetaAppReviewSubmissionWithoutIncludesResponse
	isSet bool
}

func (v NullableBetaAppReviewSubmissionWithoutIncludesResponse) Get() *BetaAppReviewSubmissionWithoutIncludesResponse {
	return v.value
}

func (v *NullableBetaAppReviewSubmissionWithoutIncludesResponse) Set(val *BetaAppReviewSubmissionWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppReviewSubmissionWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppReviewSubmissionWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppReviewSubmissionWithoutIncludesResponse(val *BetaAppReviewSubmissionWithoutIncludesResponse) *NullableBetaAppReviewSubmissionWithoutIncludesResponse {
	return &NullableBetaAppReviewSubmissionWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableBetaAppReviewSubmissionWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppReviewSubmissionWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


