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

// checks if the CiBranchPatterns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBranchPatterns{}

// CiBranchPatterns struct for CiBranchPatterns
type CiBranchPatterns struct {
	IsAllMatch *bool `json:"isAllMatch,omitempty"`
	Patterns []CiBranchPatternsPatternsInner `json:"patterns,omitempty"`
}

// NewCiBranchPatterns instantiates a new CiBranchPatterns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBranchPatterns() *CiBranchPatterns {
	this := CiBranchPatterns{}
	return &this
}

// NewCiBranchPatternsWithDefaults instantiates a new CiBranchPatterns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBranchPatternsWithDefaults() *CiBranchPatterns {
	this := CiBranchPatterns{}
	return &this
}

// GetIsAllMatch returns the IsAllMatch field value if set, zero value otherwise.
func (o *CiBranchPatterns) GetIsAllMatch() bool {
	if o == nil || IsNil(o.IsAllMatch) {
		var ret bool
		return ret
	}
	return *o.IsAllMatch
}

// GetIsAllMatchOk returns a tuple with the IsAllMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBranchPatterns) GetIsAllMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAllMatch) {
		return nil, false
	}
	return o.IsAllMatch, true
}

// HasIsAllMatch returns a boolean if a field has been set.
func (o *CiBranchPatterns) HasIsAllMatch() bool {
	if o != nil && !IsNil(o.IsAllMatch) {
		return true
	}

	return false
}

// SetIsAllMatch gets a reference to the given bool and assigns it to the IsAllMatch field.
func (o *CiBranchPatterns) SetIsAllMatch(v bool) {
	o.IsAllMatch = &v
}

// GetPatterns returns the Patterns field value if set, zero value otherwise.
func (o *CiBranchPatterns) GetPatterns() []CiBranchPatternsPatternsInner {
	if o == nil || IsNil(o.Patterns) {
		var ret []CiBranchPatternsPatternsInner
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBranchPatterns) GetPatternsOk() ([]CiBranchPatternsPatternsInner, bool) {
	if o == nil || IsNil(o.Patterns) {
		return nil, false
	}
	return o.Patterns, true
}

// HasPatterns returns a boolean if a field has been set.
func (o *CiBranchPatterns) HasPatterns() bool {
	if o != nil && !IsNil(o.Patterns) {
		return true
	}

	return false
}

// SetPatterns gets a reference to the given []CiBranchPatternsPatternsInner and assigns it to the Patterns field.
func (o *CiBranchPatterns) SetPatterns(v []CiBranchPatternsPatternsInner) {
	o.Patterns = v
}

func (o CiBranchPatterns) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBranchPatterns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAllMatch) {
		toSerialize["isAllMatch"] = o.IsAllMatch
	}
	if !IsNil(o.Patterns) {
		toSerialize["patterns"] = o.Patterns
	}
	return toSerialize, nil
}

type NullableCiBranchPatterns struct {
	value *CiBranchPatterns
	isSet bool
}

func (v NullableCiBranchPatterns) Get() *CiBranchPatterns {
	return v.value
}

func (v *NullableCiBranchPatterns) Set(val *CiBranchPatterns) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBranchPatterns) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBranchPatterns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBranchPatterns(val *CiBranchPatterns) *NullableCiBranchPatterns {
	return &NullableCiBranchPatterns{value: val, isSet: true}
}

func (v NullableCiBranchPatterns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBranchPatterns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


