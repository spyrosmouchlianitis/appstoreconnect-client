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

// checks if the AppCustomProductPageLocalizationInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationInlineCreate{}

// AppCustomProductPageLocalizationInlineCreate struct for AppCustomProductPageLocalizationInlineCreate
type AppCustomProductPageLocalizationInlineCreate struct {
	Type string `json:"type"`
	Id *string `json:"id,omitempty"`
	Attributes AppCustomProductPageLocalizationInlineCreateAttributes `json:"attributes"`
	Relationships *AppCustomProductPageLocalizationInlineCreateRelationships `json:"relationships,omitempty"`
}

type _AppCustomProductPageLocalizationInlineCreate AppCustomProductPageLocalizationInlineCreate

// NewAppCustomProductPageLocalizationInlineCreate instantiates a new AppCustomProductPageLocalizationInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationInlineCreate(type_ string, attributes AppCustomProductPageLocalizationInlineCreateAttributes) *AppCustomProductPageLocalizationInlineCreate {
	this := AppCustomProductPageLocalizationInlineCreate{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAppCustomProductPageLocalizationInlineCreateWithDefaults instantiates a new AppCustomProductPageLocalizationInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationInlineCreateWithDefaults() *AppCustomProductPageLocalizationInlineCreate {
	this := AppCustomProductPageLocalizationInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *AppCustomProductPageLocalizationInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppCustomProductPageLocalizationInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppCustomProductPageLocalizationInlineCreate) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value
func (o *AppCustomProductPageLocalizationInlineCreate) GetAttributes() AppCustomProductPageLocalizationInlineCreateAttributes {
	if o == nil {
		var ret AppCustomProductPageLocalizationInlineCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationInlineCreate) GetAttributesOk() (*AppCustomProductPageLocalizationInlineCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppCustomProductPageLocalizationInlineCreate) SetAttributes(v AppCustomProductPageLocalizationInlineCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationInlineCreate) GetRelationships() AppCustomProductPageLocalizationInlineCreateRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppCustomProductPageLocalizationInlineCreateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationInlineCreate) GetRelationshipsOk() (*AppCustomProductPageLocalizationInlineCreateRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationInlineCreate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppCustomProductPageLocalizationInlineCreateRelationships and assigns it to the Relationships field.
func (o *AppCustomProductPageLocalizationInlineCreate) SetRelationships(v AppCustomProductPageLocalizationInlineCreateRelationships) {
	o.Relationships = &v
}

func (o AppCustomProductPageLocalizationInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *AppCustomProductPageLocalizationInlineCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varAppCustomProductPageLocalizationInlineCreate := _AppCustomProductPageLocalizationInlineCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppCustomProductPageLocalizationInlineCreate)

	if err != nil {
		return err
	}

	*o = AppCustomProductPageLocalizationInlineCreate(varAppCustomProductPageLocalizationInlineCreate)

	return err
}

type NullableAppCustomProductPageLocalizationInlineCreate struct {
	value *AppCustomProductPageLocalizationInlineCreate
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationInlineCreate) Get() *AppCustomProductPageLocalizationInlineCreate {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationInlineCreate) Set(val *AppCustomProductPageLocalizationInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationInlineCreate(val *AppCustomProductPageLocalizationInlineCreate) *NullableAppCustomProductPageLocalizationInlineCreate {
	return &NullableAppCustomProductPageLocalizationInlineCreate{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


