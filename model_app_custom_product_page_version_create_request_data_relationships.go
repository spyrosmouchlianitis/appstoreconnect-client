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

// checks if the AppCustomProductPageVersionCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionCreateRequestDataRelationships{}

// AppCustomProductPageVersionCreateRequestDataRelationships struct for AppCustomProductPageVersionCreateRequestDataRelationships
type AppCustomProductPageVersionCreateRequestDataRelationships struct {
	AppCustomProductPage AppCustomProductPageVersionCreateRequestDataRelationshipsAppCustomProductPage `json:"appCustomProductPage"`
	AppCustomProductPageLocalizations *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations `json:"appCustomProductPageLocalizations,omitempty"`
}

type _AppCustomProductPageVersionCreateRequestDataRelationships AppCustomProductPageVersionCreateRequestDataRelationships

// NewAppCustomProductPageVersionCreateRequestDataRelationships instantiates a new AppCustomProductPageVersionCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionCreateRequestDataRelationships(appCustomProductPage AppCustomProductPageVersionCreateRequestDataRelationshipsAppCustomProductPage) *AppCustomProductPageVersionCreateRequestDataRelationships {
	this := AppCustomProductPageVersionCreateRequestDataRelationships{}
	this.AppCustomProductPage = appCustomProductPage
	return &this
}

// NewAppCustomProductPageVersionCreateRequestDataRelationshipsWithDefaults instantiates a new AppCustomProductPageVersionCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionCreateRequestDataRelationshipsWithDefaults() *AppCustomProductPageVersionCreateRequestDataRelationships {
	this := AppCustomProductPageVersionCreateRequestDataRelationships{}
	return &this
}

// GetAppCustomProductPage returns the AppCustomProductPage field value
func (o *AppCustomProductPageVersionCreateRequestDataRelationships) GetAppCustomProductPage() AppCustomProductPageVersionCreateRequestDataRelationshipsAppCustomProductPage {
	if o == nil {
		var ret AppCustomProductPageVersionCreateRequestDataRelationshipsAppCustomProductPage
		return ret
	}

	return o.AppCustomProductPage
}

// GetAppCustomProductPageOk returns a tuple with the AppCustomProductPage field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionCreateRequestDataRelationships) GetAppCustomProductPageOk() (*AppCustomProductPageVersionCreateRequestDataRelationshipsAppCustomProductPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppCustomProductPage, true
}

// SetAppCustomProductPage sets field value
func (o *AppCustomProductPageVersionCreateRequestDataRelationships) SetAppCustomProductPage(v AppCustomProductPageVersionCreateRequestDataRelationshipsAppCustomProductPage) {
	o.AppCustomProductPage = v
}

// GetAppCustomProductPageLocalizations returns the AppCustomProductPageLocalizations field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionCreateRequestDataRelationships) GetAppCustomProductPageLocalizations() AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations {
	if o == nil || IsNil(o.AppCustomProductPageLocalizations) {
		var ret AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations
		return ret
	}
	return *o.AppCustomProductPageLocalizations
}

// GetAppCustomProductPageLocalizationsOk returns a tuple with the AppCustomProductPageLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionCreateRequestDataRelationships) GetAppCustomProductPageLocalizationsOk() (*AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations, bool) {
	if o == nil || IsNil(o.AppCustomProductPageLocalizations) {
		return nil, false
	}
	return o.AppCustomProductPageLocalizations, true
}

// HasAppCustomProductPageLocalizations returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionCreateRequestDataRelationships) HasAppCustomProductPageLocalizations() bool {
	if o != nil && !IsNil(o.AppCustomProductPageLocalizations) {
		return true
	}

	return false
}

// SetAppCustomProductPageLocalizations gets a reference to the given AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations and assigns it to the AppCustomProductPageLocalizations field.
func (o *AppCustomProductPageVersionCreateRequestDataRelationships) SetAppCustomProductPageLocalizations(v AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) {
	o.AppCustomProductPageLocalizations = &v
}

func (o AppCustomProductPageVersionCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appCustomProductPage"] = o.AppCustomProductPage
	if !IsNil(o.AppCustomProductPageLocalizations) {
		toSerialize["appCustomProductPageLocalizations"] = o.AppCustomProductPageLocalizations
	}
	return toSerialize, nil
}

func (o *AppCustomProductPageVersionCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appCustomProductPage",
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

	varAppCustomProductPageVersionCreateRequestDataRelationships := _AppCustomProductPageVersionCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppCustomProductPageVersionCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = AppCustomProductPageVersionCreateRequestDataRelationships(varAppCustomProductPageVersionCreateRequestDataRelationships)

	return err
}

type NullableAppCustomProductPageVersionCreateRequestDataRelationships struct {
	value *AppCustomProductPageVersionCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppCustomProductPageVersionCreateRequestDataRelationships) Get() *AppCustomProductPageVersionCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppCustomProductPageVersionCreateRequestDataRelationships) Set(val *AppCustomProductPageVersionCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionCreateRequestDataRelationships(val *AppCustomProductPageVersionCreateRequestDataRelationships) *NullableAppCustomProductPageVersionCreateRequestDataRelationships {
	return &NullableAppCustomProductPageVersionCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


