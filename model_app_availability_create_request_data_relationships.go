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

// checks if the AppAvailabilityCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityCreateRequestDataRelationships{}

// AppAvailabilityCreateRequestDataRelationships struct for AppAvailabilityCreateRequestDataRelationships
type AppAvailabilityCreateRequestDataRelationships struct {
	App AnalyticsReportRequestCreateRequestDataRelationshipsApp `json:"app"`
	AvailableTerritories AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories `json:"availableTerritories"`
}

type _AppAvailabilityCreateRequestDataRelationships AppAvailabilityCreateRequestDataRelationships

// NewAppAvailabilityCreateRequestDataRelationships instantiates a new AppAvailabilityCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityCreateRequestDataRelationships(app AnalyticsReportRequestCreateRequestDataRelationshipsApp, availableTerritories AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) *AppAvailabilityCreateRequestDataRelationships {
	this := AppAvailabilityCreateRequestDataRelationships{}
	this.App = app
	this.AvailableTerritories = availableTerritories
	return &this
}

// NewAppAvailabilityCreateRequestDataRelationshipsWithDefaults instantiates a new AppAvailabilityCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityCreateRequestDataRelationshipsWithDefaults() *AppAvailabilityCreateRequestDataRelationships {
	this := AppAvailabilityCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *AppAvailabilityCreateRequestDataRelationships) GetApp() AnalyticsReportRequestCreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AnalyticsReportRequestCreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityCreateRequestDataRelationships) GetAppOk() (*AnalyticsReportRequestCreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *AppAvailabilityCreateRequestDataRelationships) SetApp(v AnalyticsReportRequestCreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetAvailableTerritories returns the AvailableTerritories field value
func (o *AppAvailabilityCreateRequestDataRelationships) GetAvailableTerritories() AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories {
	if o == nil {
		var ret AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories
		return ret
	}

	return o.AvailableTerritories
}

// GetAvailableTerritoriesOk returns a tuple with the AvailableTerritories field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityCreateRequestDataRelationships) GetAvailableTerritoriesOk() (*AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableTerritories, true
}

// SetAvailableTerritories sets field value
func (o *AppAvailabilityCreateRequestDataRelationships) SetAvailableTerritories(v AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) {
	o.AvailableTerritories = v
}

func (o AppAvailabilityCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	toSerialize["availableTerritories"] = o.AvailableTerritories
	return toSerialize, nil
}

func (o *AppAvailabilityCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app",
		"availableTerritories",
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

	varAppAvailabilityCreateRequestDataRelationships := _AppAvailabilityCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppAvailabilityCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = AppAvailabilityCreateRequestDataRelationships(varAppAvailabilityCreateRequestDataRelationships)

	return err
}

type NullableAppAvailabilityCreateRequestDataRelationships struct {
	value *AppAvailabilityCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppAvailabilityCreateRequestDataRelationships) Get() *AppAvailabilityCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppAvailabilityCreateRequestDataRelationships) Set(val *AppAvailabilityCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityCreateRequestDataRelationships(val *AppAvailabilityCreateRequestDataRelationships) *NullableAppAvailabilityCreateRequestDataRelationships {
	return &NullableAppAvailabilityCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppAvailabilityCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


