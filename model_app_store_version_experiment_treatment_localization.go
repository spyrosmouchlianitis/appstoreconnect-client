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

// checks if the AppStoreVersionExperimentTreatmentLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentLocalization{}

// AppStoreVersionExperimentTreatmentLocalization struct for AppStoreVersionExperimentTreatmentLocalization
type AppStoreVersionExperimentTreatmentLocalization struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreVersionExperimentTreatmentLocalizationAttributes `json:"attributes,omitempty"`
	Relationships *AppStoreVersionExperimentTreatmentLocalizationRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _AppStoreVersionExperimentTreatmentLocalization AppStoreVersionExperimentTreatmentLocalization

// NewAppStoreVersionExperimentTreatmentLocalization instantiates a new AppStoreVersionExperimentTreatmentLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentLocalization(type_ string, id string) *AppStoreVersionExperimentTreatmentLocalization {
	this := AppStoreVersionExperimentTreatmentLocalization{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionExperimentTreatmentLocalizationWithDefaults instantiates a new AppStoreVersionExperimentTreatmentLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentLocalizationWithDefaults() *AppStoreVersionExperimentTreatmentLocalization {
	this := AppStoreVersionExperimentTreatmentLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionExperimentTreatmentLocalization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionExperimentTreatmentLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionExperimentTreatmentLocalization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionExperimentTreatmentLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentLocalization) GetAttributes() AppStoreVersionExperimentTreatmentLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppStoreVersionExperimentTreatmentLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalization) GetAttributesOk() (*AppStoreVersionExperimentTreatmentLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentLocalization) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreVersionExperimentTreatmentLocalizationAttributes and assigns it to the Attributes field.
func (o *AppStoreVersionExperimentTreatmentLocalization) SetAttributes(v AppStoreVersionExperimentTreatmentLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentLocalization) GetRelationships() AppStoreVersionExperimentTreatmentLocalizationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppStoreVersionExperimentTreatmentLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalization) GetRelationshipsOk() (*AppStoreVersionExperimentTreatmentLocalizationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentLocalization) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppStoreVersionExperimentTreatmentLocalizationRelationships and assigns it to the Relationships field.
func (o *AppStoreVersionExperimentTreatmentLocalization) SetRelationships(v AppStoreVersionExperimentTreatmentLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentLocalization) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentLocalization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppStoreVersionExperimentTreatmentLocalization) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppStoreVersionExperimentTreatmentLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentLocalization) ToMap() (map[string]interface{}, error) {
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

func (o *AppStoreVersionExperimentTreatmentLocalization) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionExperimentTreatmentLocalization := _AppStoreVersionExperimentTreatmentLocalization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionExperimentTreatmentLocalization)

	if err != nil {
		return err
	}

	*o = AppStoreVersionExperimentTreatmentLocalization(varAppStoreVersionExperimentTreatmentLocalization)

	return err
}

type NullableAppStoreVersionExperimentTreatmentLocalization struct {
	value *AppStoreVersionExperimentTreatmentLocalization
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentLocalization) Get() *AppStoreVersionExperimentTreatmentLocalization {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalization) Set(val *AppStoreVersionExperimentTreatmentLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentLocalization(val *AppStoreVersionExperimentTreatmentLocalization) *NullableAppStoreVersionExperimentTreatmentLocalization {
	return &NullableAppStoreVersionExperimentTreatmentLocalization{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


