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

// checks if the AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes{}

// AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes struct for AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes
type AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes struct {
	Locale string `json:"locale"`
}

type _AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes

// NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes instantiates a new AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes(locale string) *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes {
	this := AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	return &this
}

// NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributesWithDefaults() *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes {
	this := AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetLocale returns the Locale field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

func (o AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locale"] = o.Locale
	return toSerialize, nil
}

func (o *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locale",
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

	varAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes := _AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes(varAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes)

	return err
}

type NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes struct {
	value *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) Get() *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) Set(val *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes(val *AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes {
	return &NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

