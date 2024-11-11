/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BetaAppClipInvocationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationCreateRequestDataRelationships{}

// BetaAppClipInvocationCreateRequestDataRelationships struct for BetaAppClipInvocationCreateRequestDataRelationships
type BetaAppClipInvocationCreateRequestDataRelationships struct {
	BuildBundle BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle `json:"buildBundle"`
	BetaAppClipInvocationLocalizations BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations `json:"betaAppClipInvocationLocalizations"`
}

type _BetaAppClipInvocationCreateRequestDataRelationships BetaAppClipInvocationCreateRequestDataRelationships

// NewBetaAppClipInvocationCreateRequestDataRelationships instantiates a new BetaAppClipInvocationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationCreateRequestDataRelationships(buildBundle BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle, betaAppClipInvocationLocalizations BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) *BetaAppClipInvocationCreateRequestDataRelationships {
	this := BetaAppClipInvocationCreateRequestDataRelationships{}
	this.BuildBundle = buildBundle
	this.BetaAppClipInvocationLocalizations = betaAppClipInvocationLocalizations
	return &this
}

// NewBetaAppClipInvocationCreateRequestDataRelationshipsWithDefaults instantiates a new BetaAppClipInvocationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationCreateRequestDataRelationshipsWithDefaults() *BetaAppClipInvocationCreateRequestDataRelationships {
	this := BetaAppClipInvocationCreateRequestDataRelationships{}
	return &this
}

// GetBuildBundle returns the BuildBundle field value
func (o *BetaAppClipInvocationCreateRequestDataRelationships) GetBuildBundle() BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle {
	if o == nil {
		var ret BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle
		return ret
	}

	return o.BuildBundle
}

// GetBuildBundleOk returns a tuple with the BuildBundle field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationCreateRequestDataRelationships) GetBuildBundleOk() (*BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildBundle, true
}

// SetBuildBundle sets field value
func (o *BetaAppClipInvocationCreateRequestDataRelationships) SetBuildBundle(v BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) {
	o.BuildBundle = v
}

// GetBetaAppClipInvocationLocalizations returns the BetaAppClipInvocationLocalizations field value
func (o *BetaAppClipInvocationCreateRequestDataRelationships) GetBetaAppClipInvocationLocalizations() BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations {
	if o == nil {
		var ret BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations
		return ret
	}

	return o.BetaAppClipInvocationLocalizations
}

// GetBetaAppClipInvocationLocalizationsOk returns a tuple with the BetaAppClipInvocationLocalizations field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationCreateRequestDataRelationships) GetBetaAppClipInvocationLocalizationsOk() (*BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BetaAppClipInvocationLocalizations, true
}

// SetBetaAppClipInvocationLocalizations sets field value
func (o *BetaAppClipInvocationCreateRequestDataRelationships) SetBetaAppClipInvocationLocalizations(v BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) {
	o.BetaAppClipInvocationLocalizations = v
}

func (o BetaAppClipInvocationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buildBundle"] = o.BuildBundle
	toSerialize["betaAppClipInvocationLocalizations"] = o.BetaAppClipInvocationLocalizations
	return toSerialize, nil
}

func (o *BetaAppClipInvocationCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buildBundle",
		"betaAppClipInvocationLocalizations",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBetaAppClipInvocationCreateRequestDataRelationships := _BetaAppClipInvocationCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaAppClipInvocationCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = BetaAppClipInvocationCreateRequestDataRelationships(varBetaAppClipInvocationCreateRequestDataRelationships)

	return err
}

type NullableBetaAppClipInvocationCreateRequestDataRelationships struct {
	value *BetaAppClipInvocationCreateRequestDataRelationships
	isSet bool
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationships) Get() *BetaAppClipInvocationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationships) Set(val *BetaAppClipInvocationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationCreateRequestDataRelationships(val *BetaAppClipInvocationCreateRequestDataRelationships) *NullableBetaAppClipInvocationCreateRequestDataRelationships {
	return &NullableBetaAppClipInvocationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


