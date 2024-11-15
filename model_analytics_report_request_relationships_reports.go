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

// checks if the AnalyticsReportRequestRelationshipsReports type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsReportRequestRelationshipsReports{}

// AnalyticsReportRequestRelationshipsReports struct for AnalyticsReportRequestRelationshipsReports
type AnalyticsReportRequestRelationshipsReports struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AnalyticsReportRequestRelationshipsReportsDataInner `json:"data,omitempty"`
}

// NewAnalyticsReportRequestRelationshipsReports instantiates a new AnalyticsReportRequestRelationshipsReports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsReportRequestRelationshipsReports() *AnalyticsReportRequestRelationshipsReports {
	this := AnalyticsReportRequestRelationshipsReports{}
	return &this
}

// NewAnalyticsReportRequestRelationshipsReportsWithDefaults instantiates a new AnalyticsReportRequestRelationshipsReports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsReportRequestRelationshipsReportsWithDefaults() *AnalyticsReportRequestRelationshipsReports {
	this := AnalyticsReportRequestRelationshipsReports{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AnalyticsReportRequestRelationshipsReports) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestRelationshipsReports) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AnalyticsReportRequestRelationshipsReports) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *AnalyticsReportRequestRelationshipsReports) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AnalyticsReportRequestRelationshipsReports) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestRelationshipsReports) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AnalyticsReportRequestRelationshipsReports) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AnalyticsReportRequestRelationshipsReports) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AnalyticsReportRequestRelationshipsReports) GetData() []AnalyticsReportRequestRelationshipsReportsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AnalyticsReportRequestRelationshipsReportsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestRelationshipsReports) GetDataOk() ([]AnalyticsReportRequestRelationshipsReportsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AnalyticsReportRequestRelationshipsReports) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AnalyticsReportRequestRelationshipsReportsDataInner and assigns it to the Data field.
func (o *AnalyticsReportRequestRelationshipsReports) SetData(v []AnalyticsReportRequestRelationshipsReportsDataInner) {
	o.Data = v
}

func (o AnalyticsReportRequestRelationshipsReports) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsReportRequestRelationshipsReports) ToMap() (map[string]interface{}, error) {
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

type NullableAnalyticsReportRequestRelationshipsReports struct {
	value *AnalyticsReportRequestRelationshipsReports
	isSet bool
}

func (v NullableAnalyticsReportRequestRelationshipsReports) Get() *AnalyticsReportRequestRelationshipsReports {
	return v.value
}

func (v *NullableAnalyticsReportRequestRelationshipsReports) Set(val *AnalyticsReportRequestRelationshipsReports) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsReportRequestRelationshipsReports) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsReportRequestRelationshipsReports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsReportRequestRelationshipsReports(val *AnalyticsReportRequestRelationshipsReports) *NullableAnalyticsReportRequestRelationshipsReports {
	return &NullableAnalyticsReportRequestRelationshipsReports{value: val, isSet: true}
}

func (v NullableAnalyticsReportRequestRelationshipsReports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsReportRequestRelationshipsReports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


