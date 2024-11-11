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

// checks if the AppStoreVersionBuildLinkageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionBuildLinkageResponse{}

// AppStoreVersionBuildLinkageResponse struct for AppStoreVersionBuildLinkageResponse
type AppStoreVersionBuildLinkageResponse struct {
	Data AppEncryptionDeclarationRelationshipsBuildsDataInner `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AppStoreVersionBuildLinkageResponse AppStoreVersionBuildLinkageResponse

// NewAppStoreVersionBuildLinkageResponse instantiates a new AppStoreVersionBuildLinkageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionBuildLinkageResponse(data AppEncryptionDeclarationRelationshipsBuildsDataInner, links DocumentLinks) *AppStoreVersionBuildLinkageResponse {
	this := AppStoreVersionBuildLinkageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionBuildLinkageResponseWithDefaults instantiates a new AppStoreVersionBuildLinkageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionBuildLinkageResponseWithDefaults() *AppStoreVersionBuildLinkageResponse {
	this := AppStoreVersionBuildLinkageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionBuildLinkageResponse) GetData() AppEncryptionDeclarationRelationshipsBuildsDataInner {
	if o == nil {
		var ret AppEncryptionDeclarationRelationshipsBuildsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionBuildLinkageResponse) GetDataOk() (*AppEncryptionDeclarationRelationshipsBuildsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionBuildLinkageResponse) SetData(v AppEncryptionDeclarationRelationshipsBuildsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionBuildLinkageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionBuildLinkageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionBuildLinkageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppStoreVersionBuildLinkageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionBuildLinkageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppStoreVersionBuildLinkageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionBuildLinkageResponse := _AppStoreVersionBuildLinkageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionBuildLinkageResponse)

	if err != nil {
		return err
	}

	*o = AppStoreVersionBuildLinkageResponse(varAppStoreVersionBuildLinkageResponse)

	return err
}

type NullableAppStoreVersionBuildLinkageResponse struct {
	value *AppStoreVersionBuildLinkageResponse
	isSet bool
}

func (v NullableAppStoreVersionBuildLinkageResponse) Get() *AppStoreVersionBuildLinkageResponse {
	return v.value
}

func (v *NullableAppStoreVersionBuildLinkageResponse) Set(val *AppStoreVersionBuildLinkageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionBuildLinkageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionBuildLinkageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionBuildLinkageResponse(val *AppStoreVersionBuildLinkageResponse) *NullableAppStoreVersionBuildLinkageResponse {
	return &NullableAppStoreVersionBuildLinkageResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionBuildLinkageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionBuildLinkageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


