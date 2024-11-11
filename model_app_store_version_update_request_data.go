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

// checks if the AppStoreVersionUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionUpdateRequestData{}

// AppStoreVersionUpdateRequestData struct for AppStoreVersionUpdateRequestData
type AppStoreVersionUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreVersionUpdateRequestDataAttributes `json:"attributes,omitempty"`
	Relationships *AppStoreVersionUpdateRequestDataRelationships `json:"relationships,omitempty"`
}

type _AppStoreVersionUpdateRequestData AppStoreVersionUpdateRequestData

// NewAppStoreVersionUpdateRequestData instantiates a new AppStoreVersionUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionUpdateRequestData(type_ string, id string) *AppStoreVersionUpdateRequestData {
	this := AppStoreVersionUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionUpdateRequestDataWithDefaults instantiates a new AppStoreVersionUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionUpdateRequestDataWithDefaults() *AppStoreVersionUpdateRequestData {
	this := AppStoreVersionUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestData) GetAttributes() AppStoreVersionUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppStoreVersionUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestData) GetAttributesOk() (*AppStoreVersionUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreVersionUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppStoreVersionUpdateRequestData) SetAttributes(v AppStoreVersionUpdateRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestData) GetRelationships() AppStoreVersionUpdateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppStoreVersionUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestData) GetRelationshipsOk() (*AppStoreVersionUpdateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppStoreVersionUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *AppStoreVersionUpdateRequestData) SetRelationships(v AppStoreVersionUpdateRequestDataRelationships) {
	o.Relationships = &v
}

func (o AppStoreVersionUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *AppStoreVersionUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionUpdateRequestData := _AppStoreVersionUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionUpdateRequestData)

	if err != nil {
		return err
	}

	*o = AppStoreVersionUpdateRequestData(varAppStoreVersionUpdateRequestData)

	return err
}

type NullableAppStoreVersionUpdateRequestData struct {
	value *AppStoreVersionUpdateRequestData
	isSet bool
}

func (v NullableAppStoreVersionUpdateRequestData) Get() *AppStoreVersionUpdateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionUpdateRequestData) Set(val *AppStoreVersionUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionUpdateRequestData(val *AppStoreVersionUpdateRequestData) *NullableAppStoreVersionUpdateRequestData {
	return &NullableAppStoreVersionUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


