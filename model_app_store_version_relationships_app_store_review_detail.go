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

// checks if the AppStoreVersionRelationshipsAppStoreReviewDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionRelationshipsAppStoreReviewDetail{}

// AppStoreVersionRelationshipsAppStoreReviewDetail struct for AppStoreVersionRelationshipsAppStoreReviewDetail
type AppStoreVersionRelationshipsAppStoreReviewDetail struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData `json:"data,omitempty"`
}

// NewAppStoreVersionRelationshipsAppStoreReviewDetail instantiates a new AppStoreVersionRelationshipsAppStoreReviewDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionRelationshipsAppStoreReviewDetail() *AppStoreVersionRelationshipsAppStoreReviewDetail {
	this := AppStoreVersionRelationshipsAppStoreReviewDetail{}
	return &this
}

// NewAppStoreVersionRelationshipsAppStoreReviewDetailWithDefaults instantiates a new AppStoreVersionRelationshipsAppStoreReviewDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionRelationshipsAppStoreReviewDetailWithDefaults() *AppStoreVersionRelationshipsAppStoreReviewDetail {
	this := AppStoreVersionRelationshipsAppStoreReviewDetail{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsAppStoreReviewDetail) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsAppStoreReviewDetail) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsAppStoreReviewDetail) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *AppStoreVersionRelationshipsAppStoreReviewDetail) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsAppStoreReviewDetail) GetData() AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData {
	if o == nil || IsNil(o.Data) {
		var ret AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsAppStoreReviewDetail) GetDataOk() (*AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsAppStoreReviewDetail) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData and assigns it to the Data field.
func (o *AppStoreVersionRelationshipsAppStoreReviewDetail) SetData(v AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData) {
	o.Data = &v
}

func (o AppStoreVersionRelationshipsAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionRelationshipsAppStoreReviewDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppStoreVersionRelationshipsAppStoreReviewDetail struct {
	value *AppStoreVersionRelationshipsAppStoreReviewDetail
	isSet bool
}

func (v NullableAppStoreVersionRelationshipsAppStoreReviewDetail) Get() *AppStoreVersionRelationshipsAppStoreReviewDetail {
	return v.value
}

func (v *NullableAppStoreVersionRelationshipsAppStoreReviewDetail) Set(val *AppStoreVersionRelationshipsAppStoreReviewDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionRelationshipsAppStoreReviewDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionRelationshipsAppStoreReviewDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionRelationshipsAppStoreReviewDetail(val *AppStoreVersionRelationshipsAppStoreReviewDetail) *NullableAppStoreVersionRelationshipsAppStoreReviewDetail {
	return &NullableAppStoreVersionRelationshipsAppStoreReviewDetail{value: val, isSet: true}
}

func (v NullableAppStoreVersionRelationshipsAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionRelationshipsAppStoreReviewDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


