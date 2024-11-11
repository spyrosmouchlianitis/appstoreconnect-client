/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"time"
)

// checks if the ReviewSubmissionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionAttributes{}

// ReviewSubmissionAttributes struct for ReviewSubmissionAttributes
type ReviewSubmissionAttributes struct {
	Platform *Platform `json:"platform,omitempty"`
	SubmittedDate *time.Time `json:"submittedDate,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewReviewSubmissionAttributes instantiates a new ReviewSubmissionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionAttributes() *ReviewSubmissionAttributes {
	this := ReviewSubmissionAttributes{}
	return &this
}

// NewReviewSubmissionAttributesWithDefaults instantiates a new ReviewSubmissionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionAttributesWithDefaults() *ReviewSubmissionAttributes {
	this := ReviewSubmissionAttributes{}
	return &this
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ReviewSubmissionAttributes) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionAttributes) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ReviewSubmissionAttributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *ReviewSubmissionAttributes) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetSubmittedDate returns the SubmittedDate field value if set, zero value otherwise.
func (o *ReviewSubmissionAttributes) GetSubmittedDate() time.Time {
	if o == nil || IsNil(o.SubmittedDate) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedDate
}

// GetSubmittedDateOk returns a tuple with the SubmittedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionAttributes) GetSubmittedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedDate) {
		return nil, false
	}
	return o.SubmittedDate, true
}

// HasSubmittedDate returns a boolean if a field has been set.
func (o *ReviewSubmissionAttributes) HasSubmittedDate() bool {
	if o != nil && !IsNil(o.SubmittedDate) {
		return true
	}

	return false
}

// SetSubmittedDate gets a reference to the given time.Time and assigns it to the SubmittedDate field.
func (o *ReviewSubmissionAttributes) SetSubmittedDate(v time.Time) {
	o.SubmittedDate = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ReviewSubmissionAttributes) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionAttributes) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ReviewSubmissionAttributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ReviewSubmissionAttributes) SetState(v string) {
	o.State = &v
}

func (o ReviewSubmissionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.SubmittedDate) {
		toSerialize["submittedDate"] = o.SubmittedDate
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableReviewSubmissionAttributes struct {
	value *ReviewSubmissionAttributes
	isSet bool
}

func (v NullableReviewSubmissionAttributes) Get() *ReviewSubmissionAttributes {
	return v.value
}

func (v *NullableReviewSubmissionAttributes) Set(val *ReviewSubmissionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionAttributes(val *ReviewSubmissionAttributes) *NullableReviewSubmissionAttributes {
	return &NullableReviewSubmissionAttributes{value: val, isSet: true}
}

func (v NullableReviewSubmissionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


