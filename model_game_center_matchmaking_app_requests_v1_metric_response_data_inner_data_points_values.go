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

// checks if the GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues{}

// GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues struct for GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues
type GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues struct {
	Count *int32 `json:"count,omitempty"`
	AverageSecondsInQueue *float32 `json:"averageSecondsInQueue,omitempty"`
	P50SecondsInQueue *float32 `json:"p50SecondsInQueue,omitempty"`
	P95SecondsInQueue *float32 `json:"p95SecondsInQueue,omitempty"`
}

// NewGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues instantiates a new GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues() *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues {
	this := GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// NewGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValuesWithDefaults instantiates a new GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValuesWithDefaults() *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues {
	this := GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) SetCount(v int32) {
	o.Count = &v
}

// GetAverageSecondsInQueue returns the AverageSecondsInQueue field value if set, zero value otherwise.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) GetAverageSecondsInQueue() float32 {
	if o == nil || IsNil(o.AverageSecondsInQueue) {
		var ret float32
		return ret
	}
	return *o.AverageSecondsInQueue
}

// GetAverageSecondsInQueueOk returns a tuple with the AverageSecondsInQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) GetAverageSecondsInQueueOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageSecondsInQueue) {
		return nil, false
	}
	return o.AverageSecondsInQueue, true
}

// HasAverageSecondsInQueue returns a boolean if a field has been set.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) HasAverageSecondsInQueue() bool {
	if o != nil && !IsNil(o.AverageSecondsInQueue) {
		return true
	}

	return false
}

// SetAverageSecondsInQueue gets a reference to the given float32 and assigns it to the AverageSecondsInQueue field.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) SetAverageSecondsInQueue(v float32) {
	o.AverageSecondsInQueue = &v
}

// GetP50SecondsInQueue returns the P50SecondsInQueue field value if set, zero value otherwise.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) GetP50SecondsInQueue() float32 {
	if o == nil || IsNil(o.P50SecondsInQueue) {
		var ret float32
		return ret
	}
	return *o.P50SecondsInQueue
}

// GetP50SecondsInQueueOk returns a tuple with the P50SecondsInQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) GetP50SecondsInQueueOk() (*float32, bool) {
	if o == nil || IsNil(o.P50SecondsInQueue) {
		return nil, false
	}
	return o.P50SecondsInQueue, true
}

// HasP50SecondsInQueue returns a boolean if a field has been set.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) HasP50SecondsInQueue() bool {
	if o != nil && !IsNil(o.P50SecondsInQueue) {
		return true
	}

	return false
}

// SetP50SecondsInQueue gets a reference to the given float32 and assigns it to the P50SecondsInQueue field.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) SetP50SecondsInQueue(v float32) {
	o.P50SecondsInQueue = &v
}

// GetP95SecondsInQueue returns the P95SecondsInQueue field value if set, zero value otherwise.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) GetP95SecondsInQueue() float32 {
	if o == nil || IsNil(o.P95SecondsInQueue) {
		var ret float32
		return ret
	}
	return *o.P95SecondsInQueue
}

// GetP95SecondsInQueueOk returns a tuple with the P95SecondsInQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) GetP95SecondsInQueueOk() (*float32, bool) {
	if o == nil || IsNil(o.P95SecondsInQueue) {
		return nil, false
	}
	return o.P95SecondsInQueue, true
}

// HasP95SecondsInQueue returns a boolean if a field has been set.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) HasP95SecondsInQueue() bool {
	if o != nil && !IsNil(o.P95SecondsInQueue) {
		return true
	}

	return false
}

// SetP95SecondsInQueue gets a reference to the given float32 and assigns it to the P95SecondsInQueue field.
func (o *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) SetP95SecondsInQueue(v float32) {
	o.P95SecondsInQueue = &v
}

func (o GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.AverageSecondsInQueue) {
		toSerialize["averageSecondsInQueue"] = o.AverageSecondsInQueue
	}
	if !IsNil(o.P50SecondsInQueue) {
		toSerialize["p50SecondsInQueue"] = o.P50SecondsInQueue
	}
	if !IsNil(o.P95SecondsInQueue) {
		toSerialize["p95SecondsInQueue"] = o.P95SecondsInQueue
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues struct {
	value *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues
	isSet bool
}

func (v NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) Get() *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues {
	return v.value
}

func (v *NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) Set(val *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues(val *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) *NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues {
	return &NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPointsValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


