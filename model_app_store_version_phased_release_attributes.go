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

// checks if the AppStoreVersionPhasedReleaseAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPhasedReleaseAttributes{}

// AppStoreVersionPhasedReleaseAttributes struct for AppStoreVersionPhasedReleaseAttributes
type AppStoreVersionPhasedReleaseAttributes struct {
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	TotalPauseDuration *int32 `json:"totalPauseDuration,omitempty"`
	CurrentDayNumber *int32 `json:"currentDayNumber,omitempty"`
}

// NewAppStoreVersionPhasedReleaseAttributes instantiates a new AppStoreVersionPhasedReleaseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPhasedReleaseAttributes() *AppStoreVersionPhasedReleaseAttributes {
	this := AppStoreVersionPhasedReleaseAttributes{}
	return &this
}

// NewAppStoreVersionPhasedReleaseAttributesWithDefaults instantiates a new AppStoreVersionPhasedReleaseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPhasedReleaseAttributesWithDefaults() *AppStoreVersionPhasedReleaseAttributes {
	this := AppStoreVersionPhasedReleaseAttributes{}
	return &this
}

// GetPhasedReleaseState returns the PhasedReleaseState field value if set, zero value otherwise.
func (o *AppStoreVersionPhasedReleaseAttributes) GetPhasedReleaseState() PhasedReleaseState {
	if o == nil || IsNil(o.PhasedReleaseState) {
		var ret PhasedReleaseState
		return ret
	}
	return *o.PhasedReleaseState
}

// GetPhasedReleaseStateOk returns a tuple with the PhasedReleaseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseAttributes) GetPhasedReleaseStateOk() (*PhasedReleaseState, bool) {
	if o == nil || IsNil(o.PhasedReleaseState) {
		return nil, false
	}
	return o.PhasedReleaseState, true
}

// HasPhasedReleaseState returns a boolean if a field has been set.
func (o *AppStoreVersionPhasedReleaseAttributes) HasPhasedReleaseState() bool {
	if o != nil && !IsNil(o.PhasedReleaseState) {
		return true
	}

	return false
}

// SetPhasedReleaseState gets a reference to the given PhasedReleaseState and assigns it to the PhasedReleaseState field.
func (o *AppStoreVersionPhasedReleaseAttributes) SetPhasedReleaseState(v PhasedReleaseState) {
	o.PhasedReleaseState = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *AppStoreVersionPhasedReleaseAttributes) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseAttributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *AppStoreVersionPhasedReleaseAttributes) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *AppStoreVersionPhasedReleaseAttributes) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetTotalPauseDuration returns the TotalPauseDuration field value if set, zero value otherwise.
func (o *AppStoreVersionPhasedReleaseAttributes) GetTotalPauseDuration() int32 {
	if o == nil || IsNil(o.TotalPauseDuration) {
		var ret int32
		return ret
	}
	return *o.TotalPauseDuration
}

// GetTotalPauseDurationOk returns a tuple with the TotalPauseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseAttributes) GetTotalPauseDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPauseDuration) {
		return nil, false
	}
	return o.TotalPauseDuration, true
}

// HasTotalPauseDuration returns a boolean if a field has been set.
func (o *AppStoreVersionPhasedReleaseAttributes) HasTotalPauseDuration() bool {
	if o != nil && !IsNil(o.TotalPauseDuration) {
		return true
	}

	return false
}

// SetTotalPauseDuration gets a reference to the given int32 and assigns it to the TotalPauseDuration field.
func (o *AppStoreVersionPhasedReleaseAttributes) SetTotalPauseDuration(v int32) {
	o.TotalPauseDuration = &v
}

// GetCurrentDayNumber returns the CurrentDayNumber field value if set, zero value otherwise.
func (o *AppStoreVersionPhasedReleaseAttributes) GetCurrentDayNumber() int32 {
	if o == nil || IsNil(o.CurrentDayNumber) {
		var ret int32
		return ret
	}
	return *o.CurrentDayNumber
}

// GetCurrentDayNumberOk returns a tuple with the CurrentDayNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseAttributes) GetCurrentDayNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentDayNumber) {
		return nil, false
	}
	return o.CurrentDayNumber, true
}

// HasCurrentDayNumber returns a boolean if a field has been set.
func (o *AppStoreVersionPhasedReleaseAttributes) HasCurrentDayNumber() bool {
	if o != nil && !IsNil(o.CurrentDayNumber) {
		return true
	}

	return false
}

// SetCurrentDayNumber gets a reference to the given int32 and assigns it to the CurrentDayNumber field.
func (o *AppStoreVersionPhasedReleaseAttributes) SetCurrentDayNumber(v int32) {
	o.CurrentDayNumber = &v
}

func (o AppStoreVersionPhasedReleaseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPhasedReleaseAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhasedReleaseState) {
		toSerialize["phasedReleaseState"] = o.PhasedReleaseState
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.TotalPauseDuration) {
		toSerialize["totalPauseDuration"] = o.TotalPauseDuration
	}
	if !IsNil(o.CurrentDayNumber) {
		toSerialize["currentDayNumber"] = o.CurrentDayNumber
	}
	return toSerialize, nil
}

type NullableAppStoreVersionPhasedReleaseAttributes struct {
	value *AppStoreVersionPhasedReleaseAttributes
	isSet bool
}

func (v NullableAppStoreVersionPhasedReleaseAttributes) Get() *AppStoreVersionPhasedReleaseAttributes {
	return v.value
}

func (v *NullableAppStoreVersionPhasedReleaseAttributes) Set(val *AppStoreVersionPhasedReleaseAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPhasedReleaseAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPhasedReleaseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPhasedReleaseAttributes(val *AppStoreVersionPhasedReleaseAttributes) *NullableAppStoreVersionPhasedReleaseAttributes {
	return &NullableAppStoreVersionPhasedReleaseAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionPhasedReleaseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPhasedReleaseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


