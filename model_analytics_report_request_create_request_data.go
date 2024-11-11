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

// checks if the AnalyticsReportRequestCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsReportRequestCreateRequestData{}

// AnalyticsReportRequestCreateRequestData struct for AnalyticsReportRequestCreateRequestData
type AnalyticsReportRequestCreateRequestData struct {
	Type string `json:"type"`
	Attributes AnalyticsReportRequestCreateRequestDataAttributes `json:"attributes"`
	Relationships AnalyticsReportRequestCreateRequestDataRelationships `json:"relationships"`
}

type _AnalyticsReportRequestCreateRequestData AnalyticsReportRequestCreateRequestData

// NewAnalyticsReportRequestCreateRequestData instantiates a new AnalyticsReportRequestCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsReportRequestCreateRequestData(type_ string, attributes AnalyticsReportRequestCreateRequestDataAttributes, relationships AnalyticsReportRequestCreateRequestDataRelationships) *AnalyticsReportRequestCreateRequestData {
	this := AnalyticsReportRequestCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAnalyticsReportRequestCreateRequestDataWithDefaults instantiates a new AnalyticsReportRequestCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsReportRequestCreateRequestDataWithDefaults() *AnalyticsReportRequestCreateRequestData {
	this := AnalyticsReportRequestCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AnalyticsReportRequestCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AnalyticsReportRequestCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AnalyticsReportRequestCreateRequestData) GetAttributes() AnalyticsReportRequestCreateRequestDataAttributes {
	if o == nil {
		var ret AnalyticsReportRequestCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestCreateRequestData) GetAttributesOk() (*AnalyticsReportRequestCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AnalyticsReportRequestCreateRequestData) SetAttributes(v AnalyticsReportRequestCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AnalyticsReportRequestCreateRequestData) GetRelationships() AnalyticsReportRequestCreateRequestDataRelationships {
	if o == nil {
		var ret AnalyticsReportRequestCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestCreateRequestData) GetRelationshipsOk() (*AnalyticsReportRequestCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AnalyticsReportRequestCreateRequestData) SetRelationships(v AnalyticsReportRequestCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AnalyticsReportRequestCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsReportRequestCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *AnalyticsReportRequestCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAnalyticsReportRequestCreateRequestData := _AnalyticsReportRequestCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalyticsReportRequestCreateRequestData)

	if err != nil {
		return err
	}

	*o = AnalyticsReportRequestCreateRequestData(varAnalyticsReportRequestCreateRequestData)

	return err
}

type NullableAnalyticsReportRequestCreateRequestData struct {
	value *AnalyticsReportRequestCreateRequestData
	isSet bool
}

func (v NullableAnalyticsReportRequestCreateRequestData) Get() *AnalyticsReportRequestCreateRequestData {
	return v.value
}

func (v *NullableAnalyticsReportRequestCreateRequestData) Set(val *AnalyticsReportRequestCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsReportRequestCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsReportRequestCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsReportRequestCreateRequestData(val *AnalyticsReportRequestCreateRequestData) *NullableAnalyticsReportRequestCreateRequestData {
	return &NullableAnalyticsReportRequestCreateRequestData{value: val, isSet: true}
}

func (v NullableAnalyticsReportRequestCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsReportRequestCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


