/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AppEventLocalizationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationCreateRequestDataRelationships{}

// AppEventLocalizationCreateRequestDataRelationships struct for AppEventLocalizationCreateRequestDataRelationships
type AppEventLocalizationCreateRequestDataRelationships struct {
	AppEvent AppEventLocalizationCreateRequestDataRelationshipsAppEvent `json:"appEvent"`
}

type _AppEventLocalizationCreateRequestDataRelationships AppEventLocalizationCreateRequestDataRelationships

// NewAppEventLocalizationCreateRequestDataRelationships instantiates a new AppEventLocalizationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationCreateRequestDataRelationships(appEvent AppEventLocalizationCreateRequestDataRelationshipsAppEvent) *AppEventLocalizationCreateRequestDataRelationships {
	this := AppEventLocalizationCreateRequestDataRelationships{}
	this.AppEvent = appEvent
	return &this
}

// NewAppEventLocalizationCreateRequestDataRelationshipsWithDefaults instantiates a new AppEventLocalizationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationCreateRequestDataRelationshipsWithDefaults() *AppEventLocalizationCreateRequestDataRelationships {
	this := AppEventLocalizationCreateRequestDataRelationships{}
	return &this
}

// GetAppEvent returns the AppEvent field value
func (o *AppEventLocalizationCreateRequestDataRelationships) GetAppEvent() AppEventLocalizationCreateRequestDataRelationshipsAppEvent {
	if o == nil {
		var ret AppEventLocalizationCreateRequestDataRelationshipsAppEvent
		return ret
	}

	return o.AppEvent
}

// GetAppEventOk returns a tuple with the AppEvent field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationCreateRequestDataRelationships) GetAppEventOk() (*AppEventLocalizationCreateRequestDataRelationshipsAppEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppEvent, true
}

// SetAppEvent sets field value
func (o *AppEventLocalizationCreateRequestDataRelationships) SetAppEvent(v AppEventLocalizationCreateRequestDataRelationshipsAppEvent) {
	o.AppEvent = v
}

func (o AppEventLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appEvent"] = o.AppEvent
	return toSerialize, nil
}

func (o *AppEventLocalizationCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appEvent",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAppEventLocalizationCreateRequestDataRelationships := _AppEventLocalizationCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEventLocalizationCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = AppEventLocalizationCreateRequestDataRelationships(varAppEventLocalizationCreateRequestDataRelationships)

	return err
}

type NullableAppEventLocalizationCreateRequestDataRelationships struct {
	value *AppEventLocalizationCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppEventLocalizationCreateRequestDataRelationships) Get() *AppEventLocalizationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppEventLocalizationCreateRequestDataRelationships) Set(val *AppEventLocalizationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationCreateRequestDataRelationships(val *AppEventLocalizationCreateRequestDataRelationships) *NullableAppEventLocalizationCreateRequestDataRelationships {
	return &NullableAppEventLocalizationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppEventLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


