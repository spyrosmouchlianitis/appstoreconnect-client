/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues{}

// AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues struct for AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues
type AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues struct {
	CrashCount *int32 `json:"crashCount,omitempty"`
	SessionCount *int32 `json:"sessionCount,omitempty"`
	FeedbackCount *int32 `json:"feedbackCount,omitempty"`
}

// NewAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues instantiates a new AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues() *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues {
	this := AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// NewAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValuesWithDefaults instantiates a new AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValuesWithDefaults() *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues {
	this := AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// GetCrashCount returns the CrashCount field value if set, zero value otherwise.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) GetCrashCount() int32 {
	if o == nil || IsNil(o.CrashCount) {
		var ret int32
		return ret
	}
	return *o.CrashCount
}

// GetCrashCountOk returns a tuple with the CrashCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) GetCrashCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CrashCount) {
		return nil, false
	}
	return o.CrashCount, true
}

// HasCrashCount returns a boolean if a field has been set.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) HasCrashCount() bool {
	if o != nil && !IsNil(o.CrashCount) {
		return true
	}

	return false
}

// SetCrashCount gets a reference to the given int32 and assigns it to the CrashCount field.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) SetCrashCount(v int32) {
	o.CrashCount = &v
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) GetSessionCount() int32 {
	if o == nil || IsNil(o.SessionCount) {
		var ret int32
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) GetSessionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SessionCount) {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) HasSessionCount() bool {
	if o != nil && !IsNil(o.SessionCount) {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given int32 and assigns it to the SessionCount field.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) SetSessionCount(v int32) {
	o.SessionCount = &v
}

// GetFeedbackCount returns the FeedbackCount field value if set, zero value otherwise.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) GetFeedbackCount() int32 {
	if o == nil || IsNil(o.FeedbackCount) {
		var ret int32
		return ret
	}
	return *o.FeedbackCount
}

// GetFeedbackCountOk returns a tuple with the FeedbackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) GetFeedbackCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FeedbackCount) {
		return nil, false
	}
	return o.FeedbackCount, true
}

// HasFeedbackCount returns a boolean if a field has been set.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) HasFeedbackCount() bool {
	if o != nil && !IsNil(o.FeedbackCount) {
		return true
	}

	return false
}

// SetFeedbackCount gets a reference to the given int32 and assigns it to the FeedbackCount field.
func (o *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) SetFeedbackCount(v int32) {
	o.FeedbackCount = &v
}

func (o AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CrashCount) {
		toSerialize["crashCount"] = o.CrashCount
	}
	if !IsNil(o.SessionCount) {
		toSerialize["sessionCount"] = o.SessionCount
	}
	if !IsNil(o.FeedbackCount) {
		toSerialize["feedbackCount"] = o.FeedbackCount
	}
	return toSerialize, nil
}

type NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues struct {
	value *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues
	isSet bool
}

func (v NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) Get() *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues {
	return v.value
}

func (v *NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) Set(val *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues(val *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) *NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues {
	return &NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues{value: val, isSet: true}
}

func (v NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsBetaTesterUsagesV1MetricResponseDataInnerDataPointsValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


