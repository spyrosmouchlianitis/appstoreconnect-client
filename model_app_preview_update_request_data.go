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

// checks if the AppPreviewUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewUpdateRequestData{}

// AppPreviewUpdateRequestData struct for AppPreviewUpdateRequestData
type AppPreviewUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppPreviewUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

type _AppPreviewUpdateRequestData AppPreviewUpdateRequestData

// NewAppPreviewUpdateRequestData instantiates a new AppPreviewUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewUpdateRequestData(type_ string, id string) *AppPreviewUpdateRequestData {
	this := AppPreviewUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppPreviewUpdateRequestDataWithDefaults instantiates a new AppPreviewUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewUpdateRequestDataWithDefaults() *AppPreviewUpdateRequestData {
	this := AppPreviewUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppPreviewUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPreviewUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPreviewUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPreviewUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPreviewUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPreviewUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppPreviewUpdateRequestData) GetAttributes() AppPreviewUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppPreviewUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewUpdateRequestData) GetAttributesOk() (*AppPreviewUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppPreviewUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppPreviewUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppPreviewUpdateRequestData) SetAttributes(v AppPreviewUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o AppPreviewUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *AppPreviewUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppPreviewUpdateRequestData := _AppPreviewUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppPreviewUpdateRequestData)

	if err != nil {
		return err
	}

	*o = AppPreviewUpdateRequestData(varAppPreviewUpdateRequestData)

	return err
}

type NullableAppPreviewUpdateRequestData struct {
	value *AppPreviewUpdateRequestData
	isSet bool
}

func (v NullableAppPreviewUpdateRequestData) Get() *AppPreviewUpdateRequestData {
	return v.value
}

func (v *NullableAppPreviewUpdateRequestData) Set(val *AppPreviewUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewUpdateRequestData(val *AppPreviewUpdateRequestData) *NullableAppPreviewUpdateRequestData {
	return &NullableAppPreviewUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppPreviewUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


