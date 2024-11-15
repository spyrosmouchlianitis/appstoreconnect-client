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

// checks if the AppPreviewSetCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetCreateRequestDataAttributes{}

// AppPreviewSetCreateRequestDataAttributes struct for AppPreviewSetCreateRequestDataAttributes
type AppPreviewSetCreateRequestDataAttributes struct {
	PreviewType PreviewType `json:"previewType"`
}

type _AppPreviewSetCreateRequestDataAttributes AppPreviewSetCreateRequestDataAttributes

// NewAppPreviewSetCreateRequestDataAttributes instantiates a new AppPreviewSetCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetCreateRequestDataAttributes(previewType PreviewType) *AppPreviewSetCreateRequestDataAttributes {
	this := AppPreviewSetCreateRequestDataAttributes{}
	this.PreviewType = previewType
	return &this
}

// NewAppPreviewSetCreateRequestDataAttributesWithDefaults instantiates a new AppPreviewSetCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetCreateRequestDataAttributesWithDefaults() *AppPreviewSetCreateRequestDataAttributes {
	this := AppPreviewSetCreateRequestDataAttributes{}
	return &this
}

// GetPreviewType returns the PreviewType field value
func (o *AppPreviewSetCreateRequestDataAttributes) GetPreviewType() PreviewType {
	if o == nil {
		var ret PreviewType
		return ret
	}

	return o.PreviewType
}

// GetPreviewTypeOk returns a tuple with the PreviewType field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetCreateRequestDataAttributes) GetPreviewTypeOk() (*PreviewType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviewType, true
}

// SetPreviewType sets field value
func (o *AppPreviewSetCreateRequestDataAttributes) SetPreviewType(v PreviewType) {
	o.PreviewType = v
}

func (o AppPreviewSetCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["previewType"] = o.PreviewType
	return toSerialize, nil
}

func (o *AppPreviewSetCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"previewType",
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

	varAppPreviewSetCreateRequestDataAttributes := _AppPreviewSetCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppPreviewSetCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = AppPreviewSetCreateRequestDataAttributes(varAppPreviewSetCreateRequestDataAttributes)

	return err
}

type NullableAppPreviewSetCreateRequestDataAttributes struct {
	value *AppPreviewSetCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppPreviewSetCreateRequestDataAttributes) Get() *AppPreviewSetCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppPreviewSetCreateRequestDataAttributes) Set(val *AppPreviewSetCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetCreateRequestDataAttributes(val *AppPreviewSetCreateRequestDataAttributes) *NullableAppPreviewSetCreateRequestDataAttributes {
	return &NullableAppPreviewSetCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppPreviewSetCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


