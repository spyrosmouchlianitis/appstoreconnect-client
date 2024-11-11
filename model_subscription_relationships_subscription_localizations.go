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

// checks if the SubscriptionRelationshipsSubscriptionLocalizations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionRelationshipsSubscriptionLocalizations{}

// SubscriptionRelationshipsSubscriptionLocalizations struct for SubscriptionRelationshipsSubscriptionLocalizations
type SubscriptionRelationshipsSubscriptionLocalizations struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []SubscriptionRelationshipsSubscriptionLocalizationsDataInner `json:"data,omitempty"`
}

// NewSubscriptionRelationshipsSubscriptionLocalizations instantiates a new SubscriptionRelationshipsSubscriptionLocalizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionRelationshipsSubscriptionLocalizations() *SubscriptionRelationshipsSubscriptionLocalizations {
	this := SubscriptionRelationshipsSubscriptionLocalizations{}
	return &this
}

// NewSubscriptionRelationshipsSubscriptionLocalizationsWithDefaults instantiates a new SubscriptionRelationshipsSubscriptionLocalizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionRelationshipsSubscriptionLocalizationsWithDefaults() *SubscriptionRelationshipsSubscriptionLocalizations {
	this := SubscriptionRelationshipsSubscriptionLocalizations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) GetData() []SubscriptionRelationshipsSubscriptionLocalizationsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []SubscriptionRelationshipsSubscriptionLocalizationsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) GetDataOk() ([]SubscriptionRelationshipsSubscriptionLocalizationsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SubscriptionRelationshipsSubscriptionLocalizationsDataInner and assigns it to the Data field.
func (o *SubscriptionRelationshipsSubscriptionLocalizations) SetData(v []SubscriptionRelationshipsSubscriptionLocalizationsDataInner) {
	o.Data = v
}

func (o SubscriptionRelationshipsSubscriptionLocalizations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionRelationshipsSubscriptionLocalizations) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionRelationshipsSubscriptionLocalizations struct {
	value *SubscriptionRelationshipsSubscriptionLocalizations
	isSet bool
}

func (v NullableSubscriptionRelationshipsSubscriptionLocalizations) Get() *SubscriptionRelationshipsSubscriptionLocalizations {
	return v.value
}

func (v *NullableSubscriptionRelationshipsSubscriptionLocalizations) Set(val *SubscriptionRelationshipsSubscriptionLocalizations) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionRelationshipsSubscriptionLocalizations) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionRelationshipsSubscriptionLocalizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionRelationshipsSubscriptionLocalizations(val *SubscriptionRelationshipsSubscriptionLocalizations) *NullableSubscriptionRelationshipsSubscriptionLocalizations {
	return &NullableSubscriptionRelationshipsSubscriptionLocalizations{value: val, isSet: true}
}

func (v NullableSubscriptionRelationshipsSubscriptionLocalizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionRelationshipsSubscriptionLocalizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


