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

// checks if the InAppPurchaseV2RelationshipsImages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2RelationshipsImages{}

// InAppPurchaseV2RelationshipsImages struct for InAppPurchaseV2RelationshipsImages
type InAppPurchaseV2RelationshipsImages struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []InAppPurchaseV2RelationshipsImagesDataInner `json:"data,omitempty"`
}

// NewInAppPurchaseV2RelationshipsImages instantiates a new InAppPurchaseV2RelationshipsImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2RelationshipsImages() *InAppPurchaseV2RelationshipsImages {
	this := InAppPurchaseV2RelationshipsImages{}
	return &this
}

// NewInAppPurchaseV2RelationshipsImagesWithDefaults instantiates a new InAppPurchaseV2RelationshipsImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2RelationshipsImagesWithDefaults() *InAppPurchaseV2RelationshipsImages {
	this := InAppPurchaseV2RelationshipsImages{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsImages) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsImages) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsImages) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *InAppPurchaseV2RelationshipsImages) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsImages) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsImages) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsImages) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *InAppPurchaseV2RelationshipsImages) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsImages) GetData() []InAppPurchaseV2RelationshipsImagesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []InAppPurchaseV2RelationshipsImagesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsImages) GetDataOk() ([]InAppPurchaseV2RelationshipsImagesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsImages) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []InAppPurchaseV2RelationshipsImagesDataInner and assigns it to the Data field.
func (o *InAppPurchaseV2RelationshipsImages) SetData(v []InAppPurchaseV2RelationshipsImagesDataInner) {
	o.Data = v
}

func (o InAppPurchaseV2RelationshipsImages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2RelationshipsImages) ToMap() (map[string]interface{}, error) {
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

type NullableInAppPurchaseV2RelationshipsImages struct {
	value *InAppPurchaseV2RelationshipsImages
	isSet bool
}

func (v NullableInAppPurchaseV2RelationshipsImages) Get() *InAppPurchaseV2RelationshipsImages {
	return v.value
}

func (v *NullableInAppPurchaseV2RelationshipsImages) Set(val *InAppPurchaseV2RelationshipsImages) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2RelationshipsImages) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2RelationshipsImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2RelationshipsImages(val *InAppPurchaseV2RelationshipsImages) *NullableInAppPurchaseV2RelationshipsImages {
	return &NullableInAppPurchaseV2RelationshipsImages{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2RelationshipsImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2RelationshipsImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

