/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"time"
)

// checks if the BetaAppReviewSubmissionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppReviewSubmissionAttributes{}

// BetaAppReviewSubmissionAttributes struct for BetaAppReviewSubmissionAttributes
type BetaAppReviewSubmissionAttributes struct {
	BetaReviewState *BetaReviewState `json:"betaReviewState,omitempty"`
	SubmittedDate *time.Time `json:"submittedDate,omitempty"`
}

// NewBetaAppReviewSubmissionAttributes instantiates a new BetaAppReviewSubmissionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppReviewSubmissionAttributes() *BetaAppReviewSubmissionAttributes {
	this := BetaAppReviewSubmissionAttributes{}
	return &this
}

// NewBetaAppReviewSubmissionAttributesWithDefaults instantiates a new BetaAppReviewSubmissionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppReviewSubmissionAttributesWithDefaults() *BetaAppReviewSubmissionAttributes {
	this := BetaAppReviewSubmissionAttributes{}
	return &this
}

// GetBetaReviewState returns the BetaReviewState field value if set, zero value otherwise.
func (o *BetaAppReviewSubmissionAttributes) GetBetaReviewState() BetaReviewState {
	if o == nil || IsNil(o.BetaReviewState) {
		var ret BetaReviewState
		return ret
	}
	return *o.BetaReviewState
}

// GetBetaReviewStateOk returns a tuple with the BetaReviewState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmissionAttributes) GetBetaReviewStateOk() (*BetaReviewState, bool) {
	if o == nil || IsNil(o.BetaReviewState) {
		return nil, false
	}
	return o.BetaReviewState, true
}

// HasBetaReviewState returns a boolean if a field has been set.
func (o *BetaAppReviewSubmissionAttributes) HasBetaReviewState() bool {
	if o != nil && !IsNil(o.BetaReviewState) {
		return true
	}

	return false
}

// SetBetaReviewState gets a reference to the given BetaReviewState and assigns it to the BetaReviewState field.
func (o *BetaAppReviewSubmissionAttributes) SetBetaReviewState(v BetaReviewState) {
	o.BetaReviewState = &v
}

// GetSubmittedDate returns the SubmittedDate field value if set, zero value otherwise.
func (o *BetaAppReviewSubmissionAttributes) GetSubmittedDate() time.Time {
	if o == nil || IsNil(o.SubmittedDate) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedDate
}

// GetSubmittedDateOk returns a tuple with the SubmittedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmissionAttributes) GetSubmittedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedDate) {
		return nil, false
	}
	return o.SubmittedDate, true
}

// HasSubmittedDate returns a boolean if a field has been set.
func (o *BetaAppReviewSubmissionAttributes) HasSubmittedDate() bool {
	if o != nil && !IsNil(o.SubmittedDate) {
		return true
	}

	return false
}

// SetSubmittedDate gets a reference to the given time.Time and assigns it to the SubmittedDate field.
func (o *BetaAppReviewSubmissionAttributes) SetSubmittedDate(v time.Time) {
	o.SubmittedDate = &v
}

func (o BetaAppReviewSubmissionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppReviewSubmissionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BetaReviewState) {
		toSerialize["betaReviewState"] = o.BetaReviewState
	}
	if !IsNil(o.SubmittedDate) {
		toSerialize["submittedDate"] = o.SubmittedDate
	}
	return toSerialize, nil
}

type NullableBetaAppReviewSubmissionAttributes struct {
	value *BetaAppReviewSubmissionAttributes
	isSet bool
}

func (v NullableBetaAppReviewSubmissionAttributes) Get() *BetaAppReviewSubmissionAttributes {
	return v.value
}

func (v *NullableBetaAppReviewSubmissionAttributes) Set(val *BetaAppReviewSubmissionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppReviewSubmissionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppReviewSubmissionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppReviewSubmissionAttributes(val *BetaAppReviewSubmissionAttributes) *NullableBetaAppReviewSubmissionAttributes {
	return &NullableBetaAppReviewSubmissionAttributes{value: val, isSet: true}
}

func (v NullableBetaAppReviewSubmissionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppReviewSubmissionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


