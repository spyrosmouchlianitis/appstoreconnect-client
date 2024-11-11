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

// checks if the AppStoreVersionExperimentV2CreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentV2CreateRequestData{}

// AppStoreVersionExperimentV2CreateRequestData struct for AppStoreVersionExperimentV2CreateRequestData
type AppStoreVersionExperimentV2CreateRequestData struct {
	Type string `json:"type"`
	Attributes AppStoreVersionExperimentV2CreateRequestDataAttributes `json:"attributes"`
	Relationships AnalyticsReportRequestCreateRequestDataRelationships `json:"relationships"`
}

type _AppStoreVersionExperimentV2CreateRequestData AppStoreVersionExperimentV2CreateRequestData

// NewAppStoreVersionExperimentV2CreateRequestData instantiates a new AppStoreVersionExperimentV2CreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentV2CreateRequestData(type_ string, attributes AppStoreVersionExperimentV2CreateRequestDataAttributes, relationships AnalyticsReportRequestCreateRequestDataRelationships) *AppStoreVersionExperimentV2CreateRequestData {
	this := AppStoreVersionExperimentV2CreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionExperimentV2CreateRequestDataWithDefaults instantiates a new AppStoreVersionExperimentV2CreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentV2CreateRequestDataWithDefaults() *AppStoreVersionExperimentV2CreateRequestData {
	this := AppStoreVersionExperimentV2CreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionExperimentV2CreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2CreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionExperimentV2CreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppStoreVersionExperimentV2CreateRequestData) GetAttributes() AppStoreVersionExperimentV2CreateRequestDataAttributes {
	if o == nil {
		var ret AppStoreVersionExperimentV2CreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2CreateRequestData) GetAttributesOk() (*AppStoreVersionExperimentV2CreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppStoreVersionExperimentV2CreateRequestData) SetAttributes(v AppStoreVersionExperimentV2CreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionExperimentV2CreateRequestData) GetRelationships() AnalyticsReportRequestCreateRequestDataRelationships {
	if o == nil {
		var ret AnalyticsReportRequestCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentV2CreateRequestData) GetRelationshipsOk() (*AnalyticsReportRequestCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionExperimentV2CreateRequestData) SetRelationships(v AnalyticsReportRequestCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionExperimentV2CreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentV2CreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *AppStoreVersionExperimentV2CreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionExperimentV2CreateRequestData := _AppStoreVersionExperimentV2CreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionExperimentV2CreateRequestData)

	if err != nil {
		return err
	}

	*o = AppStoreVersionExperimentV2CreateRequestData(varAppStoreVersionExperimentV2CreateRequestData)

	return err
}

type NullableAppStoreVersionExperimentV2CreateRequestData struct {
	value *AppStoreVersionExperimentV2CreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionExperimentV2CreateRequestData) Get() *AppStoreVersionExperimentV2CreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionExperimentV2CreateRequestData) Set(val *AppStoreVersionExperimentV2CreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentV2CreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentV2CreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentV2CreateRequestData(val *AppStoreVersionExperimentV2CreateRequestData) *NullableAppStoreVersionExperimentV2CreateRequestData {
	return &NullableAppStoreVersionExperimentV2CreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentV2CreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentV2CreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


