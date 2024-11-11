/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the AppClipAdvancedExperienceRelationshipsLocalizations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceRelationshipsLocalizations{}

// AppClipAdvancedExperienceRelationshipsLocalizations struct for AppClipAdvancedExperienceRelationshipsLocalizations
type AppClipAdvancedExperienceRelationshipsLocalizations struct {
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppClipAdvancedExperienceRelationshipsLocalizationsDataInner `json:"data,omitempty"`
}

// NewAppClipAdvancedExperienceRelationshipsLocalizations instantiates a new AppClipAdvancedExperienceRelationshipsLocalizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceRelationshipsLocalizations() *AppClipAdvancedExperienceRelationshipsLocalizations {
	this := AppClipAdvancedExperienceRelationshipsLocalizations{}
	return &this
}

// NewAppClipAdvancedExperienceRelationshipsLocalizationsWithDefaults instantiates a new AppClipAdvancedExperienceRelationshipsLocalizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceRelationshipsLocalizationsWithDefaults() *AppClipAdvancedExperienceRelationshipsLocalizations {
	this := AppClipAdvancedExperienceRelationshipsLocalizations{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceRelationshipsLocalizations) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceRelationshipsLocalizations) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceRelationshipsLocalizations) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppClipAdvancedExperienceRelationshipsLocalizations) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceRelationshipsLocalizations) GetData() []AppClipAdvancedExperienceRelationshipsLocalizationsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppClipAdvancedExperienceRelationshipsLocalizationsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceRelationshipsLocalizations) GetDataOk() ([]AppClipAdvancedExperienceRelationshipsLocalizationsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceRelationshipsLocalizations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppClipAdvancedExperienceRelationshipsLocalizationsDataInner and assigns it to the Data field.
func (o *AppClipAdvancedExperienceRelationshipsLocalizations) SetData(v []AppClipAdvancedExperienceRelationshipsLocalizationsDataInner) {
	o.Data = v
}

func (o AppClipAdvancedExperienceRelationshipsLocalizations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceRelationshipsLocalizations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceRelationshipsLocalizations struct {
	value *AppClipAdvancedExperienceRelationshipsLocalizations
	isSet bool
}

func (v NullableAppClipAdvancedExperienceRelationshipsLocalizations) Get() *AppClipAdvancedExperienceRelationshipsLocalizations {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceRelationshipsLocalizations) Set(val *AppClipAdvancedExperienceRelationshipsLocalizations) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceRelationshipsLocalizations) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceRelationshipsLocalizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceRelationshipsLocalizations(val *AppClipAdvancedExperienceRelationshipsLocalizations) *NullableAppClipAdvancedExperienceRelationshipsLocalizations {
	return &NullableAppClipAdvancedExperienceRelationshipsLocalizations{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceRelationshipsLocalizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceRelationshipsLocalizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


