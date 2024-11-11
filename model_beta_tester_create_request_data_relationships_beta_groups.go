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

// checks if the BetaTesterCreateRequestDataRelationshipsBetaGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterCreateRequestDataRelationshipsBetaGroups{}

// BetaTesterCreateRequestDataRelationshipsBetaGroups struct for BetaTesterCreateRequestDataRelationshipsBetaGroups
type BetaTesterCreateRequestDataRelationshipsBetaGroups struct {
	Data []AppRelationshipsBetaGroupsDataInner `json:"data,omitempty"`
}

// NewBetaTesterCreateRequestDataRelationshipsBetaGroups instantiates a new BetaTesterCreateRequestDataRelationshipsBetaGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterCreateRequestDataRelationshipsBetaGroups() *BetaTesterCreateRequestDataRelationshipsBetaGroups {
	this := BetaTesterCreateRequestDataRelationshipsBetaGroups{}
	return &this
}

// NewBetaTesterCreateRequestDataRelationshipsBetaGroupsWithDefaults instantiates a new BetaTesterCreateRequestDataRelationshipsBetaGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterCreateRequestDataRelationshipsBetaGroupsWithDefaults() *BetaTesterCreateRequestDataRelationshipsBetaGroups {
	this := BetaTesterCreateRequestDataRelationshipsBetaGroups{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BetaTesterCreateRequestDataRelationshipsBetaGroups) GetData() []AppRelationshipsBetaGroupsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppRelationshipsBetaGroupsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestDataRelationshipsBetaGroups) GetDataOk() ([]AppRelationshipsBetaGroupsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BetaTesterCreateRequestDataRelationshipsBetaGroups) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppRelationshipsBetaGroupsDataInner and assigns it to the Data field.
func (o *BetaTesterCreateRequestDataRelationshipsBetaGroups) SetData(v []AppRelationshipsBetaGroupsDataInner) {
	o.Data = v
}

func (o BetaTesterCreateRequestDataRelationshipsBetaGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterCreateRequestDataRelationshipsBetaGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBetaTesterCreateRequestDataRelationshipsBetaGroups struct {
	value *BetaTesterCreateRequestDataRelationshipsBetaGroups
	isSet bool
}

func (v NullableBetaTesterCreateRequestDataRelationshipsBetaGroups) Get() *BetaTesterCreateRequestDataRelationshipsBetaGroups {
	return v.value
}

func (v *NullableBetaTesterCreateRequestDataRelationshipsBetaGroups) Set(val *BetaTesterCreateRequestDataRelationshipsBetaGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterCreateRequestDataRelationshipsBetaGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterCreateRequestDataRelationshipsBetaGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterCreateRequestDataRelationshipsBetaGroups(val *BetaTesterCreateRequestDataRelationshipsBetaGroups) *NullableBetaTesterCreateRequestDataRelationshipsBetaGroups {
	return &NullableBetaTesterCreateRequestDataRelationshipsBetaGroups{value: val, isSet: true}
}

func (v NullableBetaTesterCreateRequestDataRelationshipsBetaGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterCreateRequestDataRelationshipsBetaGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


