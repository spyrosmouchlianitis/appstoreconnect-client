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

// checks if the BetaTesterRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterRelationships{}

// BetaTesterRelationships struct for BetaTesterRelationships
type BetaTesterRelationships struct {
	Apps *BetaTesterRelationshipsApps `json:"apps,omitempty"`
	BetaGroups *AppRelationshipsBetaGroups `json:"betaGroups,omitempty"`
	Builds *AppRelationshipsBuilds `json:"builds,omitempty"`
}

// NewBetaTesterRelationships instantiates a new BetaTesterRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterRelationships() *BetaTesterRelationships {
	this := BetaTesterRelationships{}
	return &this
}

// NewBetaTesterRelationshipsWithDefaults instantiates a new BetaTesterRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterRelationshipsWithDefaults() *BetaTesterRelationships {
	this := BetaTesterRelationships{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *BetaTesterRelationships) GetApps() BetaTesterRelationshipsApps {
	if o == nil || IsNil(o.Apps) {
		var ret BetaTesterRelationshipsApps
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterRelationships) GetAppsOk() (*BetaTesterRelationshipsApps, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *BetaTesterRelationships) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given BetaTesterRelationshipsApps and assigns it to the Apps field.
func (o *BetaTesterRelationships) SetApps(v BetaTesterRelationshipsApps) {
	o.Apps = &v
}

// GetBetaGroups returns the BetaGroups field value if set, zero value otherwise.
func (o *BetaTesterRelationships) GetBetaGroups() AppRelationshipsBetaGroups {
	if o == nil || IsNil(o.BetaGroups) {
		var ret AppRelationshipsBetaGroups
		return ret
	}
	return *o.BetaGroups
}

// GetBetaGroupsOk returns a tuple with the BetaGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterRelationships) GetBetaGroupsOk() (*AppRelationshipsBetaGroups, bool) {
	if o == nil || IsNil(o.BetaGroups) {
		return nil, false
	}
	return o.BetaGroups, true
}

// HasBetaGroups returns a boolean if a field has been set.
func (o *BetaTesterRelationships) HasBetaGroups() bool {
	if o != nil && !IsNil(o.BetaGroups) {
		return true
	}

	return false
}

// SetBetaGroups gets a reference to the given AppRelationshipsBetaGroups and assigns it to the BetaGroups field.
func (o *BetaTesterRelationships) SetBetaGroups(v AppRelationshipsBetaGroups) {
	o.BetaGroups = &v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *BetaTesterRelationships) GetBuilds() AppRelationshipsBuilds {
	if o == nil || IsNil(o.Builds) {
		var ret AppRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool) {
	if o == nil || IsNil(o.Builds) {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *BetaTesterRelationships) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given AppRelationshipsBuilds and assigns it to the Builds field.
func (o *BetaTesterRelationships) SetBuilds(v AppRelationshipsBuilds) {
	o.Builds = &v
}

func (o BetaTesterRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.BetaGroups) {
		toSerialize["betaGroups"] = o.BetaGroups
	}
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	return toSerialize, nil
}

type NullableBetaTesterRelationships struct {
	value *BetaTesterRelationships
	isSet bool
}

func (v NullableBetaTesterRelationships) Get() *BetaTesterRelationships {
	return v.value
}

func (v *NullableBetaTesterRelationships) Set(val *BetaTesterRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterRelationships(val *BetaTesterRelationships) *NullableBetaTesterRelationships {
	return &NullableBetaTesterRelationships{value: val, isSet: true}
}

func (v NullableBetaTesterRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

