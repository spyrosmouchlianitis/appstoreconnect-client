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

// checks if the PromotedPurchaseRelationshipsPromotionImages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseRelationshipsPromotionImages{}

// PromotedPurchaseRelationshipsPromotionImages struct for PromotedPurchaseRelationshipsPromotionImages
type PromotedPurchaseRelationshipsPromotionImages struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []PromotedPurchaseRelationshipsPromotionImagesDataInner `json:"data,omitempty"`
}

// NewPromotedPurchaseRelationshipsPromotionImages instantiates a new PromotedPurchaseRelationshipsPromotionImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseRelationshipsPromotionImages() *PromotedPurchaseRelationshipsPromotionImages {
	this := PromotedPurchaseRelationshipsPromotionImages{}
	return &this
}

// NewPromotedPurchaseRelationshipsPromotionImagesWithDefaults instantiates a new PromotedPurchaseRelationshipsPromotionImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseRelationshipsPromotionImagesWithDefaults() *PromotedPurchaseRelationshipsPromotionImages {
	this := PromotedPurchaseRelationshipsPromotionImages{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PromotedPurchaseRelationshipsPromotionImages) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseRelationshipsPromotionImages) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PromotedPurchaseRelationshipsPromotionImages) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *PromotedPurchaseRelationshipsPromotionImages) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PromotedPurchaseRelationshipsPromotionImages) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseRelationshipsPromotionImages) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PromotedPurchaseRelationshipsPromotionImages) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *PromotedPurchaseRelationshipsPromotionImages) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PromotedPurchaseRelationshipsPromotionImages) GetData() []PromotedPurchaseRelationshipsPromotionImagesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []PromotedPurchaseRelationshipsPromotionImagesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseRelationshipsPromotionImages) GetDataOk() ([]PromotedPurchaseRelationshipsPromotionImagesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PromotedPurchaseRelationshipsPromotionImages) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PromotedPurchaseRelationshipsPromotionImagesDataInner and assigns it to the Data field.
func (o *PromotedPurchaseRelationshipsPromotionImages) SetData(v []PromotedPurchaseRelationshipsPromotionImagesDataInner) {
	o.Data = v
}

func (o PromotedPurchaseRelationshipsPromotionImages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseRelationshipsPromotionImages) ToMap() (map[string]interface{}, error) {
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

type NullablePromotedPurchaseRelationshipsPromotionImages struct {
	value *PromotedPurchaseRelationshipsPromotionImages
	isSet bool
}

func (v NullablePromotedPurchaseRelationshipsPromotionImages) Get() *PromotedPurchaseRelationshipsPromotionImages {
	return v.value
}

func (v *NullablePromotedPurchaseRelationshipsPromotionImages) Set(val *PromotedPurchaseRelationshipsPromotionImages) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseRelationshipsPromotionImages) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseRelationshipsPromotionImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseRelationshipsPromotionImages(val *PromotedPurchaseRelationshipsPromotionImages) *NullablePromotedPurchaseRelationshipsPromotionImages {
	return &NullablePromotedPurchaseRelationshipsPromotionImages{value: val, isSet: true}
}

func (v NullablePromotedPurchaseRelationshipsPromotionImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseRelationshipsPromotionImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


