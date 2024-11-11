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

// checks if the AppStoreVersionExperimentTreatmentCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentCreateRequestDataRelationships{}

// AppStoreVersionExperimentTreatmentCreateRequestDataRelationships struct for AppStoreVersionExperimentTreatmentCreateRequestDataRelationships
type AppStoreVersionExperimentTreatmentCreateRequestDataRelationships struct {
	AppStoreVersionExperiment *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment `json:"appStoreVersionExperiment,omitempty"`
	AppStoreVersionExperimentV2 *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment `json:"appStoreVersionExperimentV2,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentCreateRequestDataRelationships instantiates a new AppStoreVersionExperimentTreatmentCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentCreateRequestDataRelationships() *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships {
	this := AppStoreVersionExperimentTreatmentCreateRequestDataRelationships{}
	return &this
}

// NewAppStoreVersionExperimentTreatmentCreateRequestDataRelationshipsWithDefaults instantiates a new AppStoreVersionExperimentTreatmentCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentCreateRequestDataRelationshipsWithDefaults() *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships {
	this := AppStoreVersionExperimentTreatmentCreateRequestDataRelationships{}
	return &this
}

// GetAppStoreVersionExperiment returns the AppStoreVersionExperiment field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) GetAppStoreVersionExperiment() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	if o == nil || IsNil(o.AppStoreVersionExperiment) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
		return ret
	}
	return *o.AppStoreVersionExperiment
}

// GetAppStoreVersionExperimentOk returns a tuple with the AppStoreVersionExperiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) GetAppStoreVersionExperimentOk() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperiment) {
		return nil, false
	}
	return o.AppStoreVersionExperiment, true
}

// HasAppStoreVersionExperiment returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) HasAppStoreVersionExperiment() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperiment) {
		return true
	}

	return false
}

// SetAppStoreVersionExperiment gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment and assigns it to the AppStoreVersionExperiment field.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) SetAppStoreVersionExperiment(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) {
	o.AppStoreVersionExperiment = &v
}

// GetAppStoreVersionExperimentV2 returns the AppStoreVersionExperimentV2 field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) GetAppStoreVersionExperimentV2() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	if o == nil || IsNil(o.AppStoreVersionExperimentV2) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
		return ret
	}
	return *o.AppStoreVersionExperimentV2
}

// GetAppStoreVersionExperimentV2Ok returns a tuple with the AppStoreVersionExperimentV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) GetAppStoreVersionExperimentV2Ok() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperimentV2) {
		return nil, false
	}
	return o.AppStoreVersionExperimentV2, true
}

// HasAppStoreVersionExperimentV2 returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) HasAppStoreVersionExperimentV2() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperimentV2) {
		return true
	}

	return false
}

// SetAppStoreVersionExperimentV2 gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment and assigns it to the AppStoreVersionExperimentV2 field.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) SetAppStoreVersionExperimentV2(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) {
	o.AppStoreVersionExperimentV2 = &v
}

func (o AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppStoreVersionExperiment) {
		toSerialize["appStoreVersionExperiment"] = o.AppStoreVersionExperiment
	}
	if !IsNil(o.AppStoreVersionExperimentV2) {
		toSerialize["appStoreVersionExperimentV2"] = o.AppStoreVersionExperimentV2
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships struct {
	value *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships) Get() *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships) Set(val *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships(val *AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) *NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships {
	return &NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

