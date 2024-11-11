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

// checks if the ReviewSubmissionRelationshipsItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionRelationshipsItems{}

// ReviewSubmissionRelationshipsItems struct for ReviewSubmissionRelationshipsItems
type ReviewSubmissionRelationshipsItems struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []ReviewSubmissionRelationshipsItemsDataInner `json:"data,omitempty"`
}

// NewReviewSubmissionRelationshipsItems instantiates a new ReviewSubmissionRelationshipsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionRelationshipsItems() *ReviewSubmissionRelationshipsItems {
	this := ReviewSubmissionRelationshipsItems{}
	return &this
}

// NewReviewSubmissionRelationshipsItemsWithDefaults instantiates a new ReviewSubmissionRelationshipsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionRelationshipsItemsWithDefaults() *ReviewSubmissionRelationshipsItems {
	this := ReviewSubmissionRelationshipsItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReviewSubmissionRelationshipsItems) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionRelationshipsItems) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReviewSubmissionRelationshipsItems) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *ReviewSubmissionRelationshipsItems) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReviewSubmissionRelationshipsItems) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionRelationshipsItems) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReviewSubmissionRelationshipsItems) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *ReviewSubmissionRelationshipsItems) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReviewSubmissionRelationshipsItems) GetData() []ReviewSubmissionRelationshipsItemsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []ReviewSubmissionRelationshipsItemsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionRelationshipsItems) GetDataOk() ([]ReviewSubmissionRelationshipsItemsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReviewSubmissionRelationshipsItems) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ReviewSubmissionRelationshipsItemsDataInner and assigns it to the Data field.
func (o *ReviewSubmissionRelationshipsItems) SetData(v []ReviewSubmissionRelationshipsItemsDataInner) {
	o.Data = v
}

func (o ReviewSubmissionRelationshipsItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionRelationshipsItems) ToMap() (map[string]interface{}, error) {
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

type NullableReviewSubmissionRelationshipsItems struct {
	value *ReviewSubmissionRelationshipsItems
	isSet bool
}

func (v NullableReviewSubmissionRelationshipsItems) Get() *ReviewSubmissionRelationshipsItems {
	return v.value
}

func (v *NullableReviewSubmissionRelationshipsItems) Set(val *ReviewSubmissionRelationshipsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionRelationshipsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionRelationshipsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionRelationshipsItems(val *ReviewSubmissionRelationshipsItems) *NullableReviewSubmissionRelationshipsItems {
	return &NullableReviewSubmissionRelationshipsItems{value: val, isSet: true}
}

func (v NullableReviewSubmissionRelationshipsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionRelationshipsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

