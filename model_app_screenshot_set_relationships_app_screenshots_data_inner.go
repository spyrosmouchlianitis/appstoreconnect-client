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

// checks if the AppScreenshotSetRelationshipsAppScreenshotsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotSetRelationshipsAppScreenshotsDataInner{}

// AppScreenshotSetRelationshipsAppScreenshotsDataInner struct for AppScreenshotSetRelationshipsAppScreenshotsDataInner
type AppScreenshotSetRelationshipsAppScreenshotsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _AppScreenshotSetRelationshipsAppScreenshotsDataInner AppScreenshotSetRelationshipsAppScreenshotsDataInner

// NewAppScreenshotSetRelationshipsAppScreenshotsDataInner instantiates a new AppScreenshotSetRelationshipsAppScreenshotsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSetRelationshipsAppScreenshotsDataInner(type_ string, id string) *AppScreenshotSetRelationshipsAppScreenshotsDataInner {
	this := AppScreenshotSetRelationshipsAppScreenshotsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppScreenshotSetRelationshipsAppScreenshotsDataInnerWithDefaults instantiates a new AppScreenshotSetRelationshipsAppScreenshotsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetRelationshipsAppScreenshotsDataInnerWithDefaults() *AppScreenshotSetRelationshipsAppScreenshotsDataInner {
	this := AppScreenshotSetRelationshipsAppScreenshotsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppScreenshotSetRelationshipsAppScreenshotsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetRelationshipsAppScreenshotsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppScreenshotSetRelationshipsAppScreenshotsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppScreenshotSetRelationshipsAppScreenshotsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetRelationshipsAppScreenshotsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppScreenshotSetRelationshipsAppScreenshotsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppScreenshotSetRelationshipsAppScreenshotsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotSetRelationshipsAppScreenshotsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AppScreenshotSetRelationshipsAppScreenshotsDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varAppScreenshotSetRelationshipsAppScreenshotsDataInner := _AppScreenshotSetRelationshipsAppScreenshotsDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppScreenshotSetRelationshipsAppScreenshotsDataInner)

	if err != nil {
		return err
	}

	*o = AppScreenshotSetRelationshipsAppScreenshotsDataInner(varAppScreenshotSetRelationshipsAppScreenshotsDataInner)

	return err
}

type NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner struct {
	value *AppScreenshotSetRelationshipsAppScreenshotsDataInner
	isSet bool
}

func (v NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner) Get() *AppScreenshotSetRelationshipsAppScreenshotsDataInner {
	return v.value
}

func (v *NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner) Set(val *AppScreenshotSetRelationshipsAppScreenshotsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetRelationshipsAppScreenshotsDataInner(val *AppScreenshotSetRelationshipsAppScreenshotsDataInner) *NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner {
	return &NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner{value: val, isSet: true}
}

func (v NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetRelationshipsAppScreenshotsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

