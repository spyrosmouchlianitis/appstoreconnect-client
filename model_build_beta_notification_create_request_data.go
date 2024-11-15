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

// checks if the BuildBetaNotificationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaNotificationCreateRequestData{}

// BuildBetaNotificationCreateRequestData struct for BuildBetaNotificationCreateRequestData
type BuildBetaNotificationCreateRequestData struct {
	Type string `json:"type"`
	Relationships BetaAppReviewSubmissionCreateRequestDataRelationships `json:"relationships"`
}

type _BuildBetaNotificationCreateRequestData BuildBetaNotificationCreateRequestData

// NewBuildBetaNotificationCreateRequestData instantiates a new BuildBetaNotificationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaNotificationCreateRequestData(type_ string, relationships BetaAppReviewSubmissionCreateRequestDataRelationships) *BuildBetaNotificationCreateRequestData {
	this := BuildBetaNotificationCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewBuildBetaNotificationCreateRequestDataWithDefaults instantiates a new BuildBetaNotificationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaNotificationCreateRequestDataWithDefaults() *BuildBetaNotificationCreateRequestData {
	this := BuildBetaNotificationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BuildBetaNotificationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BuildBetaNotificationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BuildBetaNotificationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *BuildBetaNotificationCreateRequestData) GetRelationships() BetaAppReviewSubmissionCreateRequestDataRelationships {
	if o == nil {
		var ret BetaAppReviewSubmissionCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *BuildBetaNotificationCreateRequestData) GetRelationshipsOk() (*BetaAppReviewSubmissionCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *BuildBetaNotificationCreateRequestData) SetRelationships(v BetaAppReviewSubmissionCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o BuildBetaNotificationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaNotificationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *BuildBetaNotificationCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varBuildBetaNotificationCreateRequestData := _BuildBetaNotificationCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildBetaNotificationCreateRequestData)

	if err != nil {
		return err
	}

	*o = BuildBetaNotificationCreateRequestData(varBuildBetaNotificationCreateRequestData)

	return err
}

type NullableBuildBetaNotificationCreateRequestData struct {
	value *BuildBetaNotificationCreateRequestData
	isSet bool
}

func (v NullableBuildBetaNotificationCreateRequestData) Get() *BuildBetaNotificationCreateRequestData {
	return v.value
}

func (v *NullableBuildBetaNotificationCreateRequestData) Set(val *BuildBetaNotificationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaNotificationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaNotificationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaNotificationCreateRequestData(val *BuildBetaNotificationCreateRequestData) *NullableBuildBetaNotificationCreateRequestData {
	return &NullableBuildBetaNotificationCreateRequestData{value: val, isSet: true}
}

func (v NullableBuildBetaNotificationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaNotificationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


