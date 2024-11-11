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

// checks if the UserInvitationCreateRequestDataRelationshipsVisibleApps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitationCreateRequestDataRelationshipsVisibleApps{}

// UserInvitationCreateRequestDataRelationshipsVisibleApps struct for UserInvitationCreateRequestDataRelationshipsVisibleApps
type UserInvitationCreateRequestDataRelationshipsVisibleApps struct {
	Data []AlternativeDistributionKeyCreateRequestDataRelationshipsAppData `json:"data,omitempty"`
}

// NewUserInvitationCreateRequestDataRelationshipsVisibleApps instantiates a new UserInvitationCreateRequestDataRelationshipsVisibleApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationCreateRequestDataRelationshipsVisibleApps() *UserInvitationCreateRequestDataRelationshipsVisibleApps {
	this := UserInvitationCreateRequestDataRelationshipsVisibleApps{}
	return &this
}

// NewUserInvitationCreateRequestDataRelationshipsVisibleAppsWithDefaults instantiates a new UserInvitationCreateRequestDataRelationshipsVisibleApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationCreateRequestDataRelationshipsVisibleAppsWithDefaults() *UserInvitationCreateRequestDataRelationshipsVisibleApps {
	this := UserInvitationCreateRequestDataRelationshipsVisibleApps{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserInvitationCreateRequestDataRelationshipsVisibleApps) GetData() []AlternativeDistributionKeyCreateRequestDataRelationshipsAppData {
	if o == nil || IsNil(o.Data) {
		var ret []AlternativeDistributionKeyCreateRequestDataRelationshipsAppData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestDataRelationshipsVisibleApps) GetDataOk() ([]AlternativeDistributionKeyCreateRequestDataRelationshipsAppData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserInvitationCreateRequestDataRelationshipsVisibleApps) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AlternativeDistributionKeyCreateRequestDataRelationshipsAppData and assigns it to the Data field.
func (o *UserInvitationCreateRequestDataRelationshipsVisibleApps) SetData(v []AlternativeDistributionKeyCreateRequestDataRelationshipsAppData) {
	o.Data = v
}

func (o UserInvitationCreateRequestDataRelationshipsVisibleApps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitationCreateRequestDataRelationshipsVisibleApps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUserInvitationCreateRequestDataRelationshipsVisibleApps struct {
	value *UserInvitationCreateRequestDataRelationshipsVisibleApps
	isSet bool
}

func (v NullableUserInvitationCreateRequestDataRelationshipsVisibleApps) Get() *UserInvitationCreateRequestDataRelationshipsVisibleApps {
	return v.value
}

func (v *NullableUserInvitationCreateRequestDataRelationshipsVisibleApps) Set(val *UserInvitationCreateRequestDataRelationshipsVisibleApps) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationCreateRequestDataRelationshipsVisibleApps) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationCreateRequestDataRelationshipsVisibleApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationCreateRequestDataRelationshipsVisibleApps(val *UserInvitationCreateRequestDataRelationshipsVisibleApps) *NullableUserInvitationCreateRequestDataRelationshipsVisibleApps {
	return &NullableUserInvitationCreateRequestDataRelationshipsVisibleApps{value: val, isSet: true}
}

func (v NullableUserInvitationCreateRequestDataRelationshipsVisibleApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationCreateRequestDataRelationshipsVisibleApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


