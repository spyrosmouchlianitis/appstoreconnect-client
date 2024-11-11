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

// checks if the AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData{}

// AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData struct for AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData
type AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData

// NewAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData instantiates a new AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData(type_ string, id string) *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData {
	this := AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentDataWithDefaults instantiates a new AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentDataWithDefaults() *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData {
	this := AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) SetId(v string) {
	o.Id = v
}

func (o AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData := _AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData)

	if err != nil {
		return err
	}

	*o = AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData(varAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData)

	return err
}

type NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData struct {
	value *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) Get() *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) Set(val *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData(val *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) *NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData {
	return &NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperimentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


