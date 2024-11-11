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

// checks if the AppStoreReviewDetailRelationshipsAppStoreReviewAttachments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewDetailRelationshipsAppStoreReviewAttachments{}

// AppStoreReviewDetailRelationshipsAppStoreReviewAttachments struct for AppStoreReviewDetailRelationshipsAppStoreReviewAttachments
type AppStoreReviewDetailRelationshipsAppStoreReviewAttachments struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsDataInner `json:"data,omitempty"`
}

// NewAppStoreReviewDetailRelationshipsAppStoreReviewAttachments instantiates a new AppStoreReviewDetailRelationshipsAppStoreReviewAttachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailRelationshipsAppStoreReviewAttachments() *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments {
	this := AppStoreReviewDetailRelationshipsAppStoreReviewAttachments{}
	return &this
}

// NewAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsWithDefaults instantiates a new AppStoreReviewDetailRelationshipsAppStoreReviewAttachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsWithDefaults() *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments {
	this := AppStoreReviewDetailRelationshipsAppStoreReviewAttachments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) GetData() []AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) GetDataOk() ([]AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsDataInner and assigns it to the Data field.
func (o *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) SetData(v []AppStoreReviewDetailRelationshipsAppStoreReviewAttachmentsDataInner) {
	o.Data = v
}

func (o AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) ToMap() (map[string]interface{}, error) {
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

type NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments struct {
	value *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments
	isSet bool
}

func (v NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments) Get() *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments {
	return v.value
}

func (v *NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments) Set(val *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments(val *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) *NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments {
	return &NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailRelationshipsAppStoreReviewAttachments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

