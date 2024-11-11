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

// checks if the AppStoreVersionLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalization{}

// AppStoreVersionLocalization struct for AppStoreVersionLocalization
type AppStoreVersionLocalization struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreVersionLocalizationAttributes `json:"attributes,omitempty"`
	Relationships *AppStoreVersionLocalizationRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _AppStoreVersionLocalization AppStoreVersionLocalization

// NewAppStoreVersionLocalization instantiates a new AppStoreVersionLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalization(type_ string, id string) *AppStoreVersionLocalization {
	this := AppStoreVersionLocalization{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionLocalizationWithDefaults instantiates a new AppStoreVersionLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationWithDefaults() *AppStoreVersionLocalization {
	this := AppStoreVersionLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionLocalization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionLocalization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreVersionLocalization) GetAttributes() AppStoreVersionLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppStoreVersionLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetAttributesOk() (*AppStoreVersionLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreVersionLocalization) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreVersionLocalizationAttributes and assigns it to the Attributes field.
func (o *AppStoreVersionLocalization) SetAttributes(v AppStoreVersionLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppStoreVersionLocalization) GetRelationships() AppStoreVersionLocalizationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppStoreVersionLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetRelationshipsOk() (*AppStoreVersionLocalizationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppStoreVersionLocalization) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppStoreVersionLocalizationRelationships and assigns it to the Relationships field.
func (o *AppStoreVersionLocalization) SetRelationships(v AppStoreVersionLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionLocalization) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionLocalization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppStoreVersionLocalization) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppStoreVersionLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *AppStoreVersionLocalization) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionLocalization := _AppStoreVersionLocalization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionLocalization)

	if err != nil {
		return err
	}

	*o = AppStoreVersionLocalization(varAppStoreVersionLocalization)

	return err
}

type NullableAppStoreVersionLocalization struct {
	value *AppStoreVersionLocalization
	isSet bool
}

func (v NullableAppStoreVersionLocalization) Get() *AppStoreVersionLocalization {
	return v.value
}

func (v *NullableAppStoreVersionLocalization) Set(val *AppStoreVersionLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalization(val *AppStoreVersionLocalization) *NullableAppStoreVersionLocalization {
	return &NullableAppStoreVersionLocalization{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


