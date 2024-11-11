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

// checks if the GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues{}

// GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues struct for GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues
type GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues struct {
	Count *int32 `json:"count,omitempty"`
	AverageNumberOfRequests *float32 `json:"averageNumberOfRequests,omitempty"`
	P50NumberOfRequests *float32 `json:"p50NumberOfRequests,omitempty"`
	P95NumberOfRequests *float32 `json:"p95NumberOfRequests,omitempty"`
}

// NewGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues instantiates a new GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues() *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues {
	this := GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// NewGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValuesWithDefaults instantiates a new GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValuesWithDefaults() *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues {
	this := GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) SetCount(v int32) {
	o.Count = &v
}

// GetAverageNumberOfRequests returns the AverageNumberOfRequests field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) GetAverageNumberOfRequests() float32 {
	if o == nil || IsNil(o.AverageNumberOfRequests) {
		var ret float32
		return ret
	}
	return *o.AverageNumberOfRequests
}

// GetAverageNumberOfRequestsOk returns a tuple with the AverageNumberOfRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) GetAverageNumberOfRequestsOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageNumberOfRequests) {
		return nil, false
	}
	return o.AverageNumberOfRequests, true
}

// HasAverageNumberOfRequests returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) HasAverageNumberOfRequests() bool {
	if o != nil && !IsNil(o.AverageNumberOfRequests) {
		return true
	}

	return false
}

// SetAverageNumberOfRequests gets a reference to the given float32 and assigns it to the AverageNumberOfRequests field.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) SetAverageNumberOfRequests(v float32) {
	o.AverageNumberOfRequests = &v
}

// GetP50NumberOfRequests returns the P50NumberOfRequests field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) GetP50NumberOfRequests() float32 {
	if o == nil || IsNil(o.P50NumberOfRequests) {
		var ret float32
		return ret
	}
	return *o.P50NumberOfRequests
}

// GetP50NumberOfRequestsOk returns a tuple with the P50NumberOfRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) GetP50NumberOfRequestsOk() (*float32, bool) {
	if o == nil || IsNil(o.P50NumberOfRequests) {
		return nil, false
	}
	return o.P50NumberOfRequests, true
}

// HasP50NumberOfRequests returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) HasP50NumberOfRequests() bool {
	if o != nil && !IsNil(o.P50NumberOfRequests) {
		return true
	}

	return false
}

// SetP50NumberOfRequests gets a reference to the given float32 and assigns it to the P50NumberOfRequests field.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) SetP50NumberOfRequests(v float32) {
	o.P50NumberOfRequests = &v
}

// GetP95NumberOfRequests returns the P95NumberOfRequests field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) GetP95NumberOfRequests() float32 {
	if o == nil || IsNil(o.P95NumberOfRequests) {
		var ret float32
		return ret
	}
	return *o.P95NumberOfRequests
}

// GetP95NumberOfRequestsOk returns a tuple with the P95NumberOfRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) GetP95NumberOfRequestsOk() (*float32, bool) {
	if o == nil || IsNil(o.P95NumberOfRequests) {
		return nil, false
	}
	return o.P95NumberOfRequests, true
}

// HasP95NumberOfRequests returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) HasP95NumberOfRequests() bool {
	if o != nil && !IsNil(o.P95NumberOfRequests) {
		return true
	}

	return false
}

// SetP95NumberOfRequests gets a reference to the given float32 and assigns it to the P95NumberOfRequests field.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) SetP95NumberOfRequests(v float32) {
	o.P95NumberOfRequests = &v
}

func (o GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.AverageNumberOfRequests) {
		toSerialize["averageNumberOfRequests"] = o.AverageNumberOfRequests
	}
	if !IsNil(o.P50NumberOfRequests) {
		toSerialize["p50NumberOfRequests"] = o.P50NumberOfRequests
	}
	if !IsNil(o.P95NumberOfRequests) {
		toSerialize["p95NumberOfRequests"] = o.P95NumberOfRequests
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues struct {
	value *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) Get() *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) Set(val *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues(val *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) *NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues {
	return &NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


