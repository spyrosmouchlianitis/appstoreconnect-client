/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues{}

// BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues struct for BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues
type BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues struct {
	CrashCount *int32 `json:"crashCount,omitempty"`
	InstallCount *int32 `json:"installCount,omitempty"`
	SessionCount *int32 `json:"sessionCount,omitempty"`
	FeedbackCount *int32 `json:"feedbackCount,omitempty"`
	InviteCount *int32 `json:"inviteCount,omitempty"`
}

// NewBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues instantiates a new BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues() *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues {
	this := BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// NewBetaBuildUsagesV1MetricResponseDataInnerDataPointsValuesWithDefaults instantiates a new BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildUsagesV1MetricResponseDataInnerDataPointsValuesWithDefaults() *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues {
	this := BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// GetCrashCount returns the CrashCount field value if set, zero value otherwise.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetCrashCount() int32 {
	if o == nil || IsNil(o.CrashCount) {
		var ret int32
		return ret
	}
	return *o.CrashCount
}

// GetCrashCountOk returns a tuple with the CrashCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetCrashCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CrashCount) {
		return nil, false
	}
	return o.CrashCount, true
}

// HasCrashCount returns a boolean if a field has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) HasCrashCount() bool {
	if o != nil && !IsNil(o.CrashCount) {
		return true
	}

	return false
}

// SetCrashCount gets a reference to the given int32 and assigns it to the CrashCount field.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) SetCrashCount(v int32) {
	o.CrashCount = &v
}

// GetInstallCount returns the InstallCount field value if set, zero value otherwise.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetInstallCount() int32 {
	if o == nil || IsNil(o.InstallCount) {
		var ret int32
		return ret
	}
	return *o.InstallCount
}

// GetInstallCountOk returns a tuple with the InstallCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetInstallCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InstallCount) {
		return nil, false
	}
	return o.InstallCount, true
}

// HasInstallCount returns a boolean if a field has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) HasInstallCount() bool {
	if o != nil && !IsNil(o.InstallCount) {
		return true
	}

	return false
}

// SetInstallCount gets a reference to the given int32 and assigns it to the InstallCount field.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) SetInstallCount(v int32) {
	o.InstallCount = &v
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetSessionCount() int32 {
	if o == nil || IsNil(o.SessionCount) {
		var ret int32
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetSessionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SessionCount) {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) HasSessionCount() bool {
	if o != nil && !IsNil(o.SessionCount) {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given int32 and assigns it to the SessionCount field.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) SetSessionCount(v int32) {
	o.SessionCount = &v
}

// GetFeedbackCount returns the FeedbackCount field value if set, zero value otherwise.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetFeedbackCount() int32 {
	if o == nil || IsNil(o.FeedbackCount) {
		var ret int32
		return ret
	}
	return *o.FeedbackCount
}

// GetFeedbackCountOk returns a tuple with the FeedbackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetFeedbackCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FeedbackCount) {
		return nil, false
	}
	return o.FeedbackCount, true
}

// HasFeedbackCount returns a boolean if a field has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) HasFeedbackCount() bool {
	if o != nil && !IsNil(o.FeedbackCount) {
		return true
	}

	return false
}

// SetFeedbackCount gets a reference to the given int32 and assigns it to the FeedbackCount field.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) SetFeedbackCount(v int32) {
	o.FeedbackCount = &v
}

// GetInviteCount returns the InviteCount field value if set, zero value otherwise.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetInviteCount() int32 {
	if o == nil || IsNil(o.InviteCount) {
		var ret int32
		return ret
	}
	return *o.InviteCount
}

// GetInviteCountOk returns a tuple with the InviteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) GetInviteCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InviteCount) {
		return nil, false
	}
	return o.InviteCount, true
}

// HasInviteCount returns a boolean if a field has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) HasInviteCount() bool {
	if o != nil && !IsNil(o.InviteCount) {
		return true
	}

	return false
}

// SetInviteCount gets a reference to the given int32 and assigns it to the InviteCount field.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) SetInviteCount(v int32) {
	o.InviteCount = &v
}

func (o BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CrashCount) {
		toSerialize["crashCount"] = o.CrashCount
	}
	if !IsNil(o.InstallCount) {
		toSerialize["installCount"] = o.InstallCount
	}
	if !IsNil(o.SessionCount) {
		toSerialize["sessionCount"] = o.SessionCount
	}
	if !IsNil(o.FeedbackCount) {
		toSerialize["feedbackCount"] = o.FeedbackCount
	}
	if !IsNil(o.InviteCount) {
		toSerialize["inviteCount"] = o.InviteCount
	}
	return toSerialize, nil
}

type NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues struct {
	value *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues
	isSet bool
}

func (v NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) Get() *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues {
	return v.value
}

func (v *NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) Set(val *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues(val *BetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) *NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues {
	return &NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues{value: val, isSet: true}
}

func (v NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildUsagesV1MetricResponseDataInnerDataPointsValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


