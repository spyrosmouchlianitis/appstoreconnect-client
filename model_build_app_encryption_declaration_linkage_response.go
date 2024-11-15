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

// checks if the BuildAppEncryptionDeclarationLinkageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildAppEncryptionDeclarationLinkageResponse{}

// BuildAppEncryptionDeclarationLinkageResponse struct for BuildAppEncryptionDeclarationLinkageResponse
type BuildAppEncryptionDeclarationLinkageResponse struct {
	Data AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _BuildAppEncryptionDeclarationLinkageResponse BuildAppEncryptionDeclarationLinkageResponse

// NewBuildAppEncryptionDeclarationLinkageResponse instantiates a new BuildAppEncryptionDeclarationLinkageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildAppEncryptionDeclarationLinkageResponse(data AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData, links DocumentLinks) *BuildAppEncryptionDeclarationLinkageResponse {
	this := BuildAppEncryptionDeclarationLinkageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildAppEncryptionDeclarationLinkageResponseWithDefaults instantiates a new BuildAppEncryptionDeclarationLinkageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildAppEncryptionDeclarationLinkageResponseWithDefaults() *BuildAppEncryptionDeclarationLinkageResponse {
	this := BuildAppEncryptionDeclarationLinkageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildAppEncryptionDeclarationLinkageResponse) GetData() AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData {
	if o == nil {
		var ret AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildAppEncryptionDeclarationLinkageResponse) GetDataOk() (*AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildAppEncryptionDeclarationLinkageResponse) SetData(v AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BuildAppEncryptionDeclarationLinkageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildAppEncryptionDeclarationLinkageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildAppEncryptionDeclarationLinkageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BuildAppEncryptionDeclarationLinkageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildAppEncryptionDeclarationLinkageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *BuildAppEncryptionDeclarationLinkageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varBuildAppEncryptionDeclarationLinkageResponse := _BuildAppEncryptionDeclarationLinkageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildAppEncryptionDeclarationLinkageResponse)

	if err != nil {
		return err
	}

	*o = BuildAppEncryptionDeclarationLinkageResponse(varBuildAppEncryptionDeclarationLinkageResponse)

	return err
}

type NullableBuildAppEncryptionDeclarationLinkageResponse struct {
	value *BuildAppEncryptionDeclarationLinkageResponse
	isSet bool
}

func (v NullableBuildAppEncryptionDeclarationLinkageResponse) Get() *BuildAppEncryptionDeclarationLinkageResponse {
	return v.value
}

func (v *NullableBuildAppEncryptionDeclarationLinkageResponse) Set(val *BuildAppEncryptionDeclarationLinkageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildAppEncryptionDeclarationLinkageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildAppEncryptionDeclarationLinkageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildAppEncryptionDeclarationLinkageResponse(val *BuildAppEncryptionDeclarationLinkageResponse) *NullableBuildAppEncryptionDeclarationLinkageResponse {
	return &NullableBuildAppEncryptionDeclarationLinkageResponse{value: val, isSet: true}
}

func (v NullableBuildAppEncryptionDeclarationLinkageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildAppEncryptionDeclarationLinkageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


