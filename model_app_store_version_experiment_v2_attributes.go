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

// checks if the AppStoreVersionExperimentV2Attributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentV2Attributes{}

// AppStoreVersionExperimentV2Attributes struct for AppStoreVersionExperimentV2Attributes
type AppStoreVersionExperimentV2Attributes struct {
	Name *string `json:"name,omitempty"`
	Platform *Platform `json:"platform,omitempty"`
	TrafficProportion *int32 `json:"trafficProportion,omitempty"`
	State *string `json:"state,omitempty"`
	ReviewRequired *bool `json:"reviewRequired,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
}

// NewAppStoreVersionExperimentV2Attributes instantiates a new AppStoreVersionExperimentV2Attributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentV2Attributes() *AppStoreVersionExperimentV2Attributes {
	this := AppStoreVersionExperimentV2Attributes{}
	return &this
}

// NewAppStoreVersionExperimentV2AttributesWithDefaults instantiates a new AppStoreVersionExperimentV2Attributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentV2AttributesWithDefaults() *AppStoreVersionExperimentV2Attributes {
	this := AppStoreVersionExperimentV2Attributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2Attributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2Attributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2Attributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppStoreVersionExperimentV2Attributes) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2Attributes) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2Attributes) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2Attributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *AppStoreVersionExperimentV2Attributes) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetTrafficProportion returns the TrafficProportion field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2Attributes) GetTrafficProportion() int32 {
	if o == nil || IsNil(o.TrafficProportion) {
		var ret int32
		return ret
	}
	return *o.TrafficProportion
}

// GetTrafficProportionOk returns a tuple with the TrafficProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2Attributes) GetTrafficProportionOk() (*int32, bool) {
	if o == nil || IsNil(o.TrafficProportion) {
		return nil, false
	}
	return o.TrafficProportion, true
}

// HasTrafficProportion returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2Attributes) HasTrafficProportion() bool {
	if o != nil && !IsNil(o.TrafficProportion) {
		return true
	}

	return false
}

// SetTrafficProportion gets a reference to the given int32 and assigns it to the TrafficProportion field.
func (o *AppStoreVersionExperimentV2Attributes) SetTrafficProportion(v int32) {
	o.TrafficProportion = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2Attributes) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2Attributes) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2Attributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AppStoreVersionExperimentV2Attributes) SetState(v string) {
	o.State = &v
}

// GetReviewRequired returns the ReviewRequired field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2Attributes) GetReviewRequired() bool {
	if o == nil || IsNil(o.ReviewRequired) {
		var ret bool
		return ret
	}
	return *o.ReviewRequired
}

// GetReviewRequiredOk returns a tuple with the ReviewRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2Attributes) GetReviewRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ReviewRequired) {
		return nil, false
	}
	return o.ReviewRequired, true
}

// HasReviewRequired returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2Attributes) HasReviewRequired() bool {
	if o != nil && !IsNil(o.ReviewRequired) {
		return true
	}

	return false
}

// SetReviewRequired gets a reference to the given bool and assigns it to the ReviewRequired field.
func (o *AppStoreVersionExperimentV2Attributes) SetReviewRequired(v bool) {
	o.ReviewRequired = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2Attributes) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2Attributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2Attributes) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *AppStoreVersionExperimentV2Attributes) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentV2Attributes) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2Attributes) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentV2Attributes) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *AppStoreVersionExperimentV2Attributes) SetEndDate(v time.Time) {
	o.EndDate = &v
}

func (o AppStoreVersionExperimentV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentV2Attributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.TrafficProportion) {
		toSerialize["trafficProportion"] = o.TrafficProportion
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ReviewRequired) {
		toSerialize["reviewRequired"] = o.ReviewRequired
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentV2Attributes struct {
	value *AppStoreVersionExperimentV2Attributes
	isSet bool
}

func (v NullableAppStoreVersionExperimentV2Attributes) Get() *AppStoreVersionExperimentV2Attributes {
	return v.value
}

func (v *NullableAppStoreVersionExperimentV2Attributes) Set(val *AppStoreVersionExperimentV2Attributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentV2Attributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentV2Attributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentV2Attributes(val *AppStoreVersionExperimentV2Attributes) *NullableAppStoreVersionExperimentV2Attributes {
	return &NullableAppStoreVersionExperimentV2Attributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentV2Attributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentV2Attributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


