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

// checks if the BuildBetaGroupsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaGroupsLinkagesRequest{}

// BuildBetaGroupsLinkagesRequest struct for BuildBetaGroupsLinkagesRequest
type BuildBetaGroupsLinkagesRequest struct {
	Data []AppRelationshipsBetaGroupsDataInner `json:"data"`
}

type _BuildBetaGroupsLinkagesRequest BuildBetaGroupsLinkagesRequest

// NewBuildBetaGroupsLinkagesRequest instantiates a new BuildBetaGroupsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaGroupsLinkagesRequest(data []AppRelationshipsBetaGroupsDataInner) *BuildBetaGroupsLinkagesRequest {
	this := BuildBetaGroupsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewBuildBetaGroupsLinkagesRequestWithDefaults instantiates a new BuildBetaGroupsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaGroupsLinkagesRequestWithDefaults() *BuildBetaGroupsLinkagesRequest {
	this := BuildBetaGroupsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BuildBetaGroupsLinkagesRequest) GetData() []AppRelationshipsBetaGroupsDataInner {
	if o == nil {
		var ret []AppRelationshipsBetaGroupsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildBetaGroupsLinkagesRequest) GetDataOk() ([]AppRelationshipsBetaGroupsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BuildBetaGroupsLinkagesRequest) SetData(v []AppRelationshipsBetaGroupsDataInner) {
	o.Data = v
}

func (o BuildBetaGroupsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaGroupsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BuildBetaGroupsLinkagesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBuildBetaGroupsLinkagesRequest := _BuildBetaGroupsLinkagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildBetaGroupsLinkagesRequest)

	if err != nil {
		return err
	}

	*o = BuildBetaGroupsLinkagesRequest(varBuildBetaGroupsLinkagesRequest)

	return err
}

type NullableBuildBetaGroupsLinkagesRequest struct {
	value *BuildBetaGroupsLinkagesRequest
	isSet bool
}

func (v NullableBuildBetaGroupsLinkagesRequest) Get() *BuildBetaGroupsLinkagesRequest {
	return v.value
}

func (v *NullableBuildBetaGroupsLinkagesRequest) Set(val *BuildBetaGroupsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaGroupsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaGroupsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaGroupsLinkagesRequest(val *BuildBetaGroupsLinkagesRequest) *NullableBuildBetaGroupsLinkagesRequest {
	return &NullableBuildBetaGroupsLinkagesRequest{value: val, isSet: true}
}

func (v NullableBuildBetaGroupsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaGroupsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


