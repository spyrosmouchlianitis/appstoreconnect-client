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

// checks if the BetaAppLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppLocalizationCreateRequestData{}

// BetaAppLocalizationCreateRequestData struct for BetaAppLocalizationCreateRequestData
type BetaAppLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes BetaAppLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships AnalyticsReportRequestCreateRequestDataRelationships `json:"relationships"`
}

type _BetaAppLocalizationCreateRequestData BetaAppLocalizationCreateRequestData

// NewBetaAppLocalizationCreateRequestData instantiates a new BetaAppLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppLocalizationCreateRequestData(type_ string, attributes BetaAppLocalizationCreateRequestDataAttributes, relationships AnalyticsReportRequestCreateRequestDataRelationships) *BetaAppLocalizationCreateRequestData {
	this := BetaAppLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewBetaAppLocalizationCreateRequestDataWithDefaults instantiates a new BetaAppLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppLocalizationCreateRequestDataWithDefaults() *BetaAppLocalizationCreateRequestData {
	this := BetaAppLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BetaAppLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaAppLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BetaAppLocalizationCreateRequestData) GetAttributes() BetaAppLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret BetaAppLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationCreateRequestData) GetAttributesOk() (*BetaAppLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BetaAppLocalizationCreateRequestData) SetAttributes(v BetaAppLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *BetaAppLocalizationCreateRequestData) GetRelationships() AnalyticsReportRequestCreateRequestDataRelationships {
	if o == nil {
		var ret AnalyticsReportRequestCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationCreateRequestData) GetRelationshipsOk() (*AnalyticsReportRequestCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *BetaAppLocalizationCreateRequestData) SetRelationships(v AnalyticsReportRequestCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o BetaAppLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *BetaAppLocalizationCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
		"relationships",
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

	varBetaAppLocalizationCreateRequestData := _BetaAppLocalizationCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaAppLocalizationCreateRequestData)

	if err != nil {
		return err
	}

	*o = BetaAppLocalizationCreateRequestData(varBetaAppLocalizationCreateRequestData)

	return err
}

type NullableBetaAppLocalizationCreateRequestData struct {
	value *BetaAppLocalizationCreateRequestData
	isSet bool
}

func (v NullableBetaAppLocalizationCreateRequestData) Get() *BetaAppLocalizationCreateRequestData {
	return v.value
}

func (v *NullableBetaAppLocalizationCreateRequestData) Set(val *BetaAppLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppLocalizationCreateRequestData(val *BetaAppLocalizationCreateRequestData) *NullableBetaAppLocalizationCreateRequestData {
	return &NullableBetaAppLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableBetaAppLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

