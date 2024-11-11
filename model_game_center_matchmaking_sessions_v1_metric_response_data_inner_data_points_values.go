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

// checks if the GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues{}

// GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues struct for GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues
type GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues struct {
	Count *int32 `json:"count,omitempty"`
	AveragePlayerCount *float32 `json:"averagePlayerCount,omitempty"`
	P50PlayerCount *float32 `json:"p50PlayerCount,omitempty"`
	P95PlayerCount *float32 `json:"p95PlayerCount,omitempty"`
}

// NewGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues instantiates a new GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues() *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues {
	this := GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// NewGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValuesWithDefaults instantiates a new GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValuesWithDefaults() *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues {
	this := GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) SetCount(v int32) {
	o.Count = &v
}

// GetAveragePlayerCount returns the AveragePlayerCount field value if set, zero value otherwise.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) GetAveragePlayerCount() float32 {
	if o == nil || IsNil(o.AveragePlayerCount) {
		var ret float32
		return ret
	}
	return *o.AveragePlayerCount
}

// GetAveragePlayerCountOk returns a tuple with the AveragePlayerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) GetAveragePlayerCountOk() (*float32, bool) {
	if o == nil || IsNil(o.AveragePlayerCount) {
		return nil, false
	}
	return o.AveragePlayerCount, true
}

// HasAveragePlayerCount returns a boolean if a field has been set.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) HasAveragePlayerCount() bool {
	if o != nil && !IsNil(o.AveragePlayerCount) {
		return true
	}

	return false
}

// SetAveragePlayerCount gets a reference to the given float32 and assigns it to the AveragePlayerCount field.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) SetAveragePlayerCount(v float32) {
	o.AveragePlayerCount = &v
}

// GetP50PlayerCount returns the P50PlayerCount field value if set, zero value otherwise.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) GetP50PlayerCount() float32 {
	if o == nil || IsNil(o.P50PlayerCount) {
		var ret float32
		return ret
	}
	return *o.P50PlayerCount
}

// GetP50PlayerCountOk returns a tuple with the P50PlayerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) GetP50PlayerCountOk() (*float32, bool) {
	if o == nil || IsNil(o.P50PlayerCount) {
		return nil, false
	}
	return o.P50PlayerCount, true
}

// HasP50PlayerCount returns a boolean if a field has been set.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) HasP50PlayerCount() bool {
	if o != nil && !IsNil(o.P50PlayerCount) {
		return true
	}

	return false
}

// SetP50PlayerCount gets a reference to the given float32 and assigns it to the P50PlayerCount field.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) SetP50PlayerCount(v float32) {
	o.P50PlayerCount = &v
}

// GetP95PlayerCount returns the P95PlayerCount field value if set, zero value otherwise.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) GetP95PlayerCount() float32 {
	if o == nil || IsNil(o.P95PlayerCount) {
		var ret float32
		return ret
	}
	return *o.P95PlayerCount
}

// GetP95PlayerCountOk returns a tuple with the P95PlayerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) GetP95PlayerCountOk() (*float32, bool) {
	if o == nil || IsNil(o.P95PlayerCount) {
		return nil, false
	}
	return o.P95PlayerCount, true
}

// HasP95PlayerCount returns a boolean if a field has been set.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) HasP95PlayerCount() bool {
	if o != nil && !IsNil(o.P95PlayerCount) {
		return true
	}

	return false
}

// SetP95PlayerCount gets a reference to the given float32 and assigns it to the P95PlayerCount field.
func (o *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) SetP95PlayerCount(v float32) {
	o.P95PlayerCount = &v
}

func (o GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.AveragePlayerCount) {
		toSerialize["averagePlayerCount"] = o.AveragePlayerCount
	}
	if !IsNil(o.P50PlayerCount) {
		toSerialize["p50PlayerCount"] = o.P50PlayerCount
	}
	if !IsNil(o.P95PlayerCount) {
		toSerialize["p95PlayerCount"] = o.P95PlayerCount
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues struct {
	value *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues
	isSet bool
}

func (v NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) Get() *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues {
	return v.value
}

func (v *NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) Set(val *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues(val *GameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) *NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues {
	return &NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingSessionsV1MetricResponseDataInnerDataPointsValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


