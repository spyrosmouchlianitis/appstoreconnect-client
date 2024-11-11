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

// checks if the AppEventLocalizationRelationshipsAppEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationRelationshipsAppEvent{}

// AppEventLocalizationRelationshipsAppEvent struct for AppEventLocalizationRelationshipsAppEvent
type AppEventLocalizationRelationshipsAppEvent struct {
	Data *AppEventLocalizationRelationshipsAppEventData `json:"data,omitempty"`
}

// NewAppEventLocalizationRelationshipsAppEvent instantiates a new AppEventLocalizationRelationshipsAppEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationRelationshipsAppEvent() *AppEventLocalizationRelationshipsAppEvent {
	this := AppEventLocalizationRelationshipsAppEvent{}
	return &this
}

// NewAppEventLocalizationRelationshipsAppEventWithDefaults instantiates a new AppEventLocalizationRelationshipsAppEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationRelationshipsAppEventWithDefaults() *AppEventLocalizationRelationshipsAppEvent {
	this := AppEventLocalizationRelationshipsAppEvent{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppEventLocalizationRelationshipsAppEvent) GetData() AppEventLocalizationRelationshipsAppEventData {
	if o == nil || IsNil(o.Data) {
		var ret AppEventLocalizationRelationshipsAppEventData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEvent) GetDataOk() (*AppEventLocalizationRelationshipsAppEventData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppEventLocalizationRelationshipsAppEvent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppEventLocalizationRelationshipsAppEventData and assigns it to the Data field.
func (o *AppEventLocalizationRelationshipsAppEvent) SetData(v AppEventLocalizationRelationshipsAppEventData) {
	o.Data = &v
}

func (o AppEventLocalizationRelationshipsAppEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationRelationshipsAppEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppEventLocalizationRelationshipsAppEvent struct {
	value *AppEventLocalizationRelationshipsAppEvent
	isSet bool
}

func (v NullableAppEventLocalizationRelationshipsAppEvent) Get() *AppEventLocalizationRelationshipsAppEvent {
	return v.value
}

func (v *NullableAppEventLocalizationRelationshipsAppEvent) Set(val *AppEventLocalizationRelationshipsAppEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationRelationshipsAppEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationRelationshipsAppEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationRelationshipsAppEvent(val *AppEventLocalizationRelationshipsAppEvent) *NullableAppEventLocalizationRelationshipsAppEvent {
	return &NullableAppEventLocalizationRelationshipsAppEvent{value: val, isSet: true}
}

func (v NullableAppEventLocalizationRelationshipsAppEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationRelationshipsAppEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


