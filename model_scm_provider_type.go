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

// checks if the ScmProviderType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmProviderType{}

// ScmProviderType struct for ScmProviderType
type ScmProviderType struct {
	Kind *string `json:"kind,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	IsOnPremise *bool `json:"isOnPremise,omitempty"`
}

// NewScmProviderType instantiates a new ScmProviderType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmProviderType() *ScmProviderType {
	this := ScmProviderType{}
	return &this
}

// NewScmProviderTypeWithDefaults instantiates a new ScmProviderType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmProviderTypeWithDefaults() *ScmProviderType {
	this := ScmProviderType{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ScmProviderType) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmProviderType) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ScmProviderType) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ScmProviderType) SetKind(v string) {
	o.Kind = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ScmProviderType) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmProviderType) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ScmProviderType) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ScmProviderType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIsOnPremise returns the IsOnPremise field value if set, zero value otherwise.
func (o *ScmProviderType) GetIsOnPremise() bool {
	if o == nil || IsNil(o.IsOnPremise) {
		var ret bool
		return ret
	}
	return *o.IsOnPremise
}

// GetIsOnPremiseOk returns a tuple with the IsOnPremise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmProviderType) GetIsOnPremiseOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOnPremise) {
		return nil, false
	}
	return o.IsOnPremise, true
}

// HasIsOnPremise returns a boolean if a field has been set.
func (o *ScmProviderType) HasIsOnPremise() bool {
	if o != nil && !IsNil(o.IsOnPremise) {
		return true
	}

	return false
}

// SetIsOnPremise gets a reference to the given bool and assigns it to the IsOnPremise field.
func (o *ScmProviderType) SetIsOnPremise(v bool) {
	o.IsOnPremise = &v
}

func (o ScmProviderType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmProviderType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.IsOnPremise) {
		toSerialize["isOnPremise"] = o.IsOnPremise
	}
	return toSerialize, nil
}

type NullableScmProviderType struct {
	value *ScmProviderType
	isSet bool
}

func (v NullableScmProviderType) Get() *ScmProviderType {
	return v.value
}

func (v *NullableScmProviderType) Set(val *ScmProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableScmProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableScmProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmProviderType(val *ScmProviderType) *NullableScmProviderType {
	return &NullableScmProviderType{value: val, isSet: true}
}

func (v NullableScmProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


