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

// checks if the AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip{}

// AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip struct for AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip
type AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip struct {
	Data AppClipAdvancedExperienceRelationshipsAppClipData `json:"data"`
}

type _AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip

// NewAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip instantiates a new AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip(data AppClipAdvancedExperienceRelationshipsAppClipData) *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip {
	this := AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip{}
	this.Data = data
	return &this
}

// NewAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClipWithDefaults instantiates a new AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClipWithDefaults() *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip {
	this := AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) GetData() AppClipAdvancedExperienceRelationshipsAppClipData {
	if o == nil {
		var ret AppClipAdvancedExperienceRelationshipsAppClipData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) GetDataOk() (*AppClipAdvancedExperienceRelationshipsAppClipData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) SetData(v AppClipAdvancedExperienceRelationshipsAppClipData) {
	o.Data = v
}

func (o AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip := _AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip)

	if err != nil {
		return err
	}

	*o = AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip(varAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip)

	return err
}

type NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip struct {
	value *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip
	isSet bool
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) Get() *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) Set(val *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip(val *AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) *NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip {
	return &NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


