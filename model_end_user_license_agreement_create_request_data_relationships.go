/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EndUserLicenseAgreementCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndUserLicenseAgreementCreateRequestDataRelationships{}

// EndUserLicenseAgreementCreateRequestDataRelationships struct for EndUserLicenseAgreementCreateRequestDataRelationships
type EndUserLicenseAgreementCreateRequestDataRelationships struct {
	App AnalyticsReportRequestCreateRequestDataRelationshipsApp `json:"app"`
	Territories AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories `json:"territories"`
}

type _EndUserLicenseAgreementCreateRequestDataRelationships EndUserLicenseAgreementCreateRequestDataRelationships

// NewEndUserLicenseAgreementCreateRequestDataRelationships instantiates a new EndUserLicenseAgreementCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementCreateRequestDataRelationships(app AnalyticsReportRequestCreateRequestDataRelationshipsApp, territories AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) *EndUserLicenseAgreementCreateRequestDataRelationships {
	this := EndUserLicenseAgreementCreateRequestDataRelationships{}
	this.App = app
	this.Territories = territories
	return &this
}

// NewEndUserLicenseAgreementCreateRequestDataRelationshipsWithDefaults instantiates a new EndUserLicenseAgreementCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementCreateRequestDataRelationshipsWithDefaults() *EndUserLicenseAgreementCreateRequestDataRelationships {
	this := EndUserLicenseAgreementCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) GetApp() AnalyticsReportRequestCreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AnalyticsReportRequestCreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) GetAppOk() (*AnalyticsReportRequestCreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) SetApp(v AnalyticsReportRequestCreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetTerritories returns the Territories field value
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) GetTerritories() AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories {
	if o == nil {
		var ret AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories
		return ret
	}

	return o.Territories
}

// GetTerritoriesOk returns a tuple with the Territories field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) GetTerritoriesOk() (*AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Territories, true
}

// SetTerritories sets field value
func (o *EndUserLicenseAgreementCreateRequestDataRelationships) SetTerritories(v AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) {
	o.Territories = v
}

func (o EndUserLicenseAgreementCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndUserLicenseAgreementCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	toSerialize["territories"] = o.Territories
	return toSerialize, nil
}

func (o *EndUserLicenseAgreementCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app",
		"territories",
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

	varEndUserLicenseAgreementCreateRequestDataRelationships := _EndUserLicenseAgreementCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndUserLicenseAgreementCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = EndUserLicenseAgreementCreateRequestDataRelationships(varEndUserLicenseAgreementCreateRequestDataRelationships)

	return err
}

type NullableEndUserLicenseAgreementCreateRequestDataRelationships struct {
	value *EndUserLicenseAgreementCreateRequestDataRelationships
	isSet bool
}

func (v NullableEndUserLicenseAgreementCreateRequestDataRelationships) Get() *EndUserLicenseAgreementCreateRequestDataRelationships {
	return v.value
}

func (v *NullableEndUserLicenseAgreementCreateRequestDataRelationships) Set(val *EndUserLicenseAgreementCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementCreateRequestDataRelationships(val *EndUserLicenseAgreementCreateRequestDataRelationships) *NullableEndUserLicenseAgreementCreateRequestDataRelationships {
	return &NullableEndUserLicenseAgreementCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


