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

// checks if the AppEncryptionDeclarationDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationDocumentResponse{}

// AppEncryptionDeclarationDocumentResponse struct for AppEncryptionDeclarationDocumentResponse
type AppEncryptionDeclarationDocumentResponse struct {
	Data AppEncryptionDeclarationDocument `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AppEncryptionDeclarationDocumentResponse AppEncryptionDeclarationDocumentResponse

// NewAppEncryptionDeclarationDocumentResponse instantiates a new AppEncryptionDeclarationDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationDocumentResponse(data AppEncryptionDeclarationDocument, links DocumentLinks) *AppEncryptionDeclarationDocumentResponse {
	this := AppEncryptionDeclarationDocumentResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppEncryptionDeclarationDocumentResponseWithDefaults instantiates a new AppEncryptionDeclarationDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationDocumentResponseWithDefaults() *AppEncryptionDeclarationDocumentResponse {
	this := AppEncryptionDeclarationDocumentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppEncryptionDeclarationDocumentResponse) GetData() AppEncryptionDeclarationDocument {
	if o == nil {
		var ret AppEncryptionDeclarationDocument
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationDocumentResponse) GetDataOk() (*AppEncryptionDeclarationDocument, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEncryptionDeclarationDocumentResponse) SetData(v AppEncryptionDeclarationDocument) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppEncryptionDeclarationDocumentResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationDocumentResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppEncryptionDeclarationDocumentResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppEncryptionDeclarationDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppEncryptionDeclarationDocumentResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppEncryptionDeclarationDocumentResponse := _AppEncryptionDeclarationDocumentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEncryptionDeclarationDocumentResponse)

	if err != nil {
		return err
	}

	*o = AppEncryptionDeclarationDocumentResponse(varAppEncryptionDeclarationDocumentResponse)

	return err
}

type NullableAppEncryptionDeclarationDocumentResponse struct {
	value *AppEncryptionDeclarationDocumentResponse
	isSet bool
}

func (v NullableAppEncryptionDeclarationDocumentResponse) Get() *AppEncryptionDeclarationDocumentResponse {
	return v.value
}

func (v *NullableAppEncryptionDeclarationDocumentResponse) Set(val *AppEncryptionDeclarationDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationDocumentResponse(val *AppEncryptionDeclarationDocumentResponse) *NullableAppEncryptionDeclarationDocumentResponse {
	return &NullableAppEncryptionDeclarationDocumentResponse{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


