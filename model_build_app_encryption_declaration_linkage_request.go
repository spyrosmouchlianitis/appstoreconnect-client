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

// checks if the BuildAppEncryptionDeclarationLinkageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildAppEncryptionDeclarationLinkageRequest{}

// BuildAppEncryptionDeclarationLinkageRequest struct for BuildAppEncryptionDeclarationLinkageRequest
type BuildAppEncryptionDeclarationLinkageRequest struct {
	Data AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData `json:"data"`
}

type _BuildAppEncryptionDeclarationLinkageRequest BuildAppEncryptionDeclarationLinkageRequest

// NewBuildAppEncryptionDeclarationLinkageRequest instantiates a new BuildAppEncryptionDeclarationLinkageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildAppEncryptionDeclarationLinkageRequest(data AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData) *BuildAppEncryptionDeclarationLinkageRequest {
	this := BuildAppEncryptionDeclarationLinkageRequest{}
	this.Data = data
	return &this
}

// NewBuildAppEncryptionDeclarationLinkageRequestWithDefaults instantiates a new BuildAppEncryptionDeclarationLinkageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildAppEncryptionDeclarationLinkageRequestWithDefaults() *BuildAppEncryptionDeclarationLinkageRequest {
	this := BuildAppEncryptionDeclarationLinkageRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BuildAppEncryptionDeclarationLinkageRequest) GetData() AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData {
	if o == nil {
		var ret AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildAppEncryptionDeclarationLinkageRequest) GetDataOk() (*AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildAppEncryptionDeclarationLinkageRequest) SetData(v AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData) {
	o.Data = v
}

func (o BuildAppEncryptionDeclarationLinkageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildAppEncryptionDeclarationLinkageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BuildAppEncryptionDeclarationLinkageRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBuildAppEncryptionDeclarationLinkageRequest := _BuildAppEncryptionDeclarationLinkageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildAppEncryptionDeclarationLinkageRequest)

	if err != nil {
		return err
	}

	*o = BuildAppEncryptionDeclarationLinkageRequest(varBuildAppEncryptionDeclarationLinkageRequest)

	return err
}

type NullableBuildAppEncryptionDeclarationLinkageRequest struct {
	value *BuildAppEncryptionDeclarationLinkageRequest
	isSet bool
}

func (v NullableBuildAppEncryptionDeclarationLinkageRequest) Get() *BuildAppEncryptionDeclarationLinkageRequest {
	return v.value
}

func (v *NullableBuildAppEncryptionDeclarationLinkageRequest) Set(val *BuildAppEncryptionDeclarationLinkageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildAppEncryptionDeclarationLinkageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildAppEncryptionDeclarationLinkageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildAppEncryptionDeclarationLinkageRequest(val *BuildAppEncryptionDeclarationLinkageRequest) *NullableBuildAppEncryptionDeclarationLinkageRequest {
	return &NullableBuildAppEncryptionDeclarationLinkageRequest{value: val, isSet: true}
}

func (v NullableBuildAppEncryptionDeclarationLinkageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildAppEncryptionDeclarationLinkageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

