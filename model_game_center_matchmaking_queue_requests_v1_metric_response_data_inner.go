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

// checks if the GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner{}

// GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner struct for GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner
type GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner struct {
	DataPoints *GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPoints `json:"dataPoints,omitempty"`
	Dimensions *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInnerDimensions `json:"dimensions,omitempty"`
	Granularity *string `json:"granularity,omitempty"`
}

// NewGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner instantiates a new GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner() *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner {
	this := GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner{}
	return &this
}

// NewGameCenterMatchmakingQueueRequestsV1MetricResponseDataInnerWithDefaults instantiates a new GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueRequestsV1MetricResponseDataInnerWithDefaults() *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner {
	this := GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner{}
	return &this
}

// GetDataPoints returns the DataPoints field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) GetDataPoints() GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPoints {
	if o == nil || IsNil(o.DataPoints) {
		var ret GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPoints
		return ret
	}
	return *o.DataPoints
}

// GetDataPointsOk returns a tuple with the DataPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) GetDataPointsOk() (*GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPoints, bool) {
	if o == nil || IsNil(o.DataPoints) {
		return nil, false
	}
	return o.DataPoints, true
}

// HasDataPoints returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) HasDataPoints() bool {
	if o != nil && !IsNil(o.DataPoints) {
		return true
	}

	return false
}

// SetDataPoints gets a reference to the given GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPoints and assigns it to the DataPoints field.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) SetDataPoints(v GameCenterMatchmakingAppRequestsV1MetricResponseDataInnerDataPoints) {
	o.DataPoints = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) GetDimensions() GameCenterMatchmakingQueueRequestsV1MetricResponseDataInnerDimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret GameCenterMatchmakingQueueRequestsV1MetricResponseDataInnerDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) GetDimensionsOk() (*GameCenterMatchmakingQueueRequestsV1MetricResponseDataInnerDimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given GameCenterMatchmakingQueueRequestsV1MetricResponseDataInnerDimensions and assigns it to the Dimensions field.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) SetDimensions(v GameCenterMatchmakingQueueRequestsV1MetricResponseDataInnerDimensions) {
	o.Dimensions = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) GetGranularity() string {
	if o == nil || IsNil(o.Granularity) {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) GetGranularityOk() (*string, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) SetGranularity(v string) {
	o.Granularity = &v
}

func (o GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataPoints) {
		toSerialize["dataPoints"] = o.DataPoints
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.Granularity) {
		toSerialize["granularity"] = o.Granularity
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner struct {
	value *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) Get() *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) Set(val *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner(val *GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) *NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner {
	return &NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


