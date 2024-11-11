/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle{}

// BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle struct for BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle
type BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle struct {
	Data BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleData `json:"data"`
}

type _BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle

// NewBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle instantiates a new BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle(data BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleData) *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle {
	this := BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle{}
	this.Data = data
	return &this
}

// NewBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleWithDefaults instantiates a new BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleWithDefaults() *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle {
	this := BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) GetData() BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleData {
	if o == nil {
		var ret BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) GetDataOk() (*BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) SetData(v BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleData) {
	o.Data = v
}

func (o BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle := _BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle)

	if err != nil {
		return err
	}

	*o = BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle(varBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle)

	return err
}

type NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle struct {
	value *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle
	isSet bool
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) Get() *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle {
	return v.value
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) Set(val *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle(val *BetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) *NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle {
	return &NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


