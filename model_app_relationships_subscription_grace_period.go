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

// checks if the AppRelationshipsSubscriptionGracePeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsSubscriptionGracePeriod{}

// AppRelationshipsSubscriptionGracePeriod struct for AppRelationshipsSubscriptionGracePeriod
type AppRelationshipsSubscriptionGracePeriod struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *AppRelationshipsSubscriptionGracePeriodData `json:"data,omitempty"`
}

// NewAppRelationshipsSubscriptionGracePeriod instantiates a new AppRelationshipsSubscriptionGracePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsSubscriptionGracePeriod() *AppRelationshipsSubscriptionGracePeriod {
	this := AppRelationshipsSubscriptionGracePeriod{}
	return &this
}

// NewAppRelationshipsSubscriptionGracePeriodWithDefaults instantiates a new AppRelationshipsSubscriptionGracePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsSubscriptionGracePeriodWithDefaults() *AppRelationshipsSubscriptionGracePeriod {
	this := AppRelationshipsSubscriptionGracePeriod{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsSubscriptionGracePeriod) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsSubscriptionGracePeriod) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsSubscriptionGracePeriod) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *AppRelationshipsSubscriptionGracePeriod) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsSubscriptionGracePeriod) GetData() AppRelationshipsSubscriptionGracePeriodData {
	if o == nil || IsNil(o.Data) {
		var ret AppRelationshipsSubscriptionGracePeriodData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsSubscriptionGracePeriod) GetDataOk() (*AppRelationshipsSubscriptionGracePeriodData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsSubscriptionGracePeriod) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppRelationshipsSubscriptionGracePeriodData and assigns it to the Data field.
func (o *AppRelationshipsSubscriptionGracePeriod) SetData(v AppRelationshipsSubscriptionGracePeriodData) {
	o.Data = &v
}

func (o AppRelationshipsSubscriptionGracePeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsSubscriptionGracePeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppRelationshipsSubscriptionGracePeriod struct {
	value *AppRelationshipsSubscriptionGracePeriod
	isSet bool
}

func (v NullableAppRelationshipsSubscriptionGracePeriod) Get() *AppRelationshipsSubscriptionGracePeriod {
	return v.value
}

func (v *NullableAppRelationshipsSubscriptionGracePeriod) Set(val *AppRelationshipsSubscriptionGracePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsSubscriptionGracePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsSubscriptionGracePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsSubscriptionGracePeriod(val *AppRelationshipsSubscriptionGracePeriod) *NullableAppRelationshipsSubscriptionGracePeriod {
	return &NullableAppRelationshipsSubscriptionGracePeriod{value: val, isSet: true}
}

func (v NullableAppRelationshipsSubscriptionGracePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsSubscriptionGracePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


