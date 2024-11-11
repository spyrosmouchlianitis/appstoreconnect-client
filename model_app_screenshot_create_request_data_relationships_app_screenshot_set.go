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

// checks if the AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet{}

// AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet struct for AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet
type AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet struct {
	Data AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner `json:"data"`
}

type _AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet

// NewAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet instantiates a new AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet(data AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet {
	this := AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet{}
	this.Data = data
	return &this
}

// NewAppScreenshotCreateRequestDataRelationshipsAppScreenshotSetWithDefaults instantiates a new AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotCreateRequestDataRelationshipsAppScreenshotSetWithDefaults() *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet {
	this := AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet{}
	return &this
}

// GetData returns the Data field value
func (o *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) GetData() AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner {
	if o == nil {
		var ret AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) GetDataOk() (*AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) SetData(v AppCustomProductPageLocalizationRelationshipsAppScreenshotSetsDataInner) {
	o.Data = v
}

func (o AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) UnmarshalJSON(data []byte) (err error) {
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

	varAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet := _AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet)

	if err != nil {
		return err
	}

	*o = AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet(varAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet)

	return err
}

type NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet struct {
	value *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet
	isSet bool
}

func (v NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) Get() *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet {
	return v.value
}

func (v *NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) Set(val *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet(val *AppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) *NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet {
	return &NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet{value: val, isSet: true}
}

func (v NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotCreateRequestDataRelationshipsAppScreenshotSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


