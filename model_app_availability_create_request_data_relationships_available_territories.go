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

// checks if the AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories{}

// AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories struct for AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories
type AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories struct {
	Data []AppAvailabilityRelationshipsAvailableTerritoriesDataInner `json:"data"`
}

type _AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories

// NewAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories instantiates a new AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories(data []AppAvailabilityRelationshipsAvailableTerritoriesDataInner) *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories {
	this := AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories{}
	this.Data = data
	return &this
}

// NewAppAvailabilityCreateRequestDataRelationshipsAvailableTerritoriesWithDefaults instantiates a new AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityCreateRequestDataRelationshipsAvailableTerritoriesWithDefaults() *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories {
	this := AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories{}
	return &this
}

// GetData returns the Data field value
func (o *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) GetData() []AppAvailabilityRelationshipsAvailableTerritoriesDataInner {
	if o == nil {
		var ret []AppAvailabilityRelationshipsAvailableTerritoriesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) GetDataOk() ([]AppAvailabilityRelationshipsAvailableTerritoriesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) SetData(v []AppAvailabilityRelationshipsAvailableTerritoriesDataInner) {
	o.Data = v
}

func (o AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories := _AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories)

	if err != nil {
		return err
	}

	*o = AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories(varAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories)

	return err
}

type NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories struct {
	value *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories
	isSet bool
}

func (v NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) Get() *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories {
	return v.value
}

func (v *NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) Set(val *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories(val *AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) *NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories {
	return &NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories{value: val, isSet: true}
}

func (v NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


