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

// checks if the AppEventLocalizationUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationUpdateRequestData{}

// AppEventLocalizationUpdateRequestData struct for AppEventLocalizationUpdateRequestData
type AppEventLocalizationUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppEventLocalizationUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

type _AppEventLocalizationUpdateRequestData AppEventLocalizationUpdateRequestData

// NewAppEventLocalizationUpdateRequestData instantiates a new AppEventLocalizationUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationUpdateRequestData(type_ string, id string) *AppEventLocalizationUpdateRequestData {
	this := AppEventLocalizationUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppEventLocalizationUpdateRequestDataWithDefaults instantiates a new AppEventLocalizationUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationUpdateRequestDataWithDefaults() *AppEventLocalizationUpdateRequestData {
	this := AppEventLocalizationUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventLocalizationUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventLocalizationUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEventLocalizationUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventLocalizationUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppEventLocalizationUpdateRequestData) GetAttributes() AppEventLocalizationUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppEventLocalizationUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationUpdateRequestData) GetAttributesOk() (*AppEventLocalizationUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppEventLocalizationUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppEventLocalizationUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppEventLocalizationUpdateRequestData) SetAttributes(v AppEventLocalizationUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o AppEventLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *AppEventLocalizationUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppEventLocalizationUpdateRequestData := _AppEventLocalizationUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEventLocalizationUpdateRequestData)

	if err != nil {
		return err
	}

	*o = AppEventLocalizationUpdateRequestData(varAppEventLocalizationUpdateRequestData)

	return err
}

type NullableAppEventLocalizationUpdateRequestData struct {
	value *AppEventLocalizationUpdateRequestData
	isSet bool
}

func (v NullableAppEventLocalizationUpdateRequestData) Get() *AppEventLocalizationUpdateRequestData {
	return v.value
}

func (v *NullableAppEventLocalizationUpdateRequestData) Set(val *AppEventLocalizationUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationUpdateRequestData(val *AppEventLocalizationUpdateRequestData) *NullableAppEventLocalizationUpdateRequestData {
	return &NullableAppEventLocalizationUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppEventLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


