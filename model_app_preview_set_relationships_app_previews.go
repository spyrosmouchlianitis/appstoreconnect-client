/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the AppPreviewSetRelationshipsAppPreviews type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetRelationshipsAppPreviews{}

// AppPreviewSetRelationshipsAppPreviews struct for AppPreviewSetRelationshipsAppPreviews
type AppPreviewSetRelationshipsAppPreviews struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppPreviewSetRelationshipsAppPreviewsDataInner `json:"data,omitempty"`
}

// NewAppPreviewSetRelationshipsAppPreviews instantiates a new AppPreviewSetRelationshipsAppPreviews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetRelationshipsAppPreviews() *AppPreviewSetRelationshipsAppPreviews {
	this := AppPreviewSetRelationshipsAppPreviews{}
	return &this
}

// NewAppPreviewSetRelationshipsAppPreviewsWithDefaults instantiates a new AppPreviewSetRelationshipsAppPreviews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetRelationshipsAppPreviewsWithDefaults() *AppPreviewSetRelationshipsAppPreviews {
	this := AppPreviewSetRelationshipsAppPreviews{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppPreviewSetRelationshipsAppPreviews) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationshipsAppPreviews) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppPreviewSetRelationshipsAppPreviews) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *AppPreviewSetRelationshipsAppPreviews) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppPreviewSetRelationshipsAppPreviews) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationshipsAppPreviews) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppPreviewSetRelationshipsAppPreviews) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppPreviewSetRelationshipsAppPreviews) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPreviewSetRelationshipsAppPreviews) GetData() []AppPreviewSetRelationshipsAppPreviewsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppPreviewSetRelationshipsAppPreviewsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationshipsAppPreviews) GetDataOk() ([]AppPreviewSetRelationshipsAppPreviewsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPreviewSetRelationshipsAppPreviews) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppPreviewSetRelationshipsAppPreviewsDataInner and assigns it to the Data field.
func (o *AppPreviewSetRelationshipsAppPreviews) SetData(v []AppPreviewSetRelationshipsAppPreviewsDataInner) {
	o.Data = v
}

func (o AppPreviewSetRelationshipsAppPreviews) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetRelationshipsAppPreviews) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppPreviewSetRelationshipsAppPreviews struct {
	value *AppPreviewSetRelationshipsAppPreviews
	isSet bool
}

func (v NullableAppPreviewSetRelationshipsAppPreviews) Get() *AppPreviewSetRelationshipsAppPreviews {
	return v.value
}

func (v *NullableAppPreviewSetRelationshipsAppPreviews) Set(val *AppPreviewSetRelationshipsAppPreviews) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetRelationshipsAppPreviews) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetRelationshipsAppPreviews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetRelationshipsAppPreviews(val *AppPreviewSetRelationshipsAppPreviews) *NullableAppPreviewSetRelationshipsAppPreviews {
	return &NullableAppPreviewSetRelationshipsAppPreviews{value: val, isSet: true}
}

func (v NullableAppPreviewSetRelationshipsAppPreviews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetRelationshipsAppPreviews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

