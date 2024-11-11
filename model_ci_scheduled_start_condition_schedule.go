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

// checks if the CiScheduledStartConditionSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiScheduledStartConditionSchedule{}

// CiScheduledStartConditionSchedule struct for CiScheduledStartConditionSchedule
type CiScheduledStartConditionSchedule struct {
	Frequency *string `json:"frequency,omitempty"`
	Days []string `json:"days,omitempty"`
	Hour *int32 `json:"hour,omitempty"`
	Minute *int32 `json:"minute,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// NewCiScheduledStartConditionSchedule instantiates a new CiScheduledStartConditionSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiScheduledStartConditionSchedule() *CiScheduledStartConditionSchedule {
	this := CiScheduledStartConditionSchedule{}
	return &this
}

// NewCiScheduledStartConditionScheduleWithDefaults instantiates a new CiScheduledStartConditionSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiScheduledStartConditionScheduleWithDefaults() *CiScheduledStartConditionSchedule {
	this := CiScheduledStartConditionSchedule{}
	return &this
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *CiScheduledStartConditionSchedule) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiScheduledStartConditionSchedule) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *CiScheduledStartConditionSchedule) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *CiScheduledStartConditionSchedule) SetFrequency(v string) {
	o.Frequency = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *CiScheduledStartConditionSchedule) GetDays() []string {
	if o == nil || IsNil(o.Days) {
		var ret []string
		return ret
	}
	return o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiScheduledStartConditionSchedule) GetDaysOk() ([]string, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *CiScheduledStartConditionSchedule) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given []string and assigns it to the Days field.
func (o *CiScheduledStartConditionSchedule) SetDays(v []string) {
	o.Days = v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *CiScheduledStartConditionSchedule) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiScheduledStartConditionSchedule) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *CiScheduledStartConditionSchedule) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *CiScheduledStartConditionSchedule) SetHour(v int32) {
	o.Hour = &v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *CiScheduledStartConditionSchedule) GetMinute() int32 {
	if o == nil || IsNil(o.Minute) {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiScheduledStartConditionSchedule) GetMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *CiScheduledStartConditionSchedule) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *CiScheduledStartConditionSchedule) SetMinute(v int32) {
	o.Minute = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *CiScheduledStartConditionSchedule) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiScheduledStartConditionSchedule) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *CiScheduledStartConditionSchedule) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *CiScheduledStartConditionSchedule) SetTimezone(v string) {
	o.Timezone = &v
}

func (o CiScheduledStartConditionSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiScheduledStartConditionSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.Minute) {
		toSerialize["minute"] = o.Minute
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

type NullableCiScheduledStartConditionSchedule struct {
	value *CiScheduledStartConditionSchedule
	isSet bool
}

func (v NullableCiScheduledStartConditionSchedule) Get() *CiScheduledStartConditionSchedule {
	return v.value
}

func (v *NullableCiScheduledStartConditionSchedule) Set(val *CiScheduledStartConditionSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableCiScheduledStartConditionSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableCiScheduledStartConditionSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiScheduledStartConditionSchedule(val *CiScheduledStartConditionSchedule) *NullableCiScheduledStartConditionSchedule {
	return &NullableCiScheduledStartConditionSchedule{value: val, isSet: true}
}

func (v NullableCiScheduledStartConditionSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiScheduledStartConditionSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


