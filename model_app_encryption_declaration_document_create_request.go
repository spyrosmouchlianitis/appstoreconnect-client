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

// checks if the AppEncryptionDeclarationDocumentCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationDocumentCreateRequest{}

// AppEncryptionDeclarationDocumentCreateRequest struct for AppEncryptionDeclarationDocumentCreateRequest
type AppEncryptionDeclarationDocumentCreateRequest struct {
	Data AppEncryptionDeclarationDocumentCreateRequestData `json:"data"`
}

type _AppEncryptionDeclarationDocumentCreateRequest AppEncryptionDeclarationDocumentCreateRequest

// NewAppEncryptionDeclarationDocumentCreateRequest instantiates a new AppEncryptionDeclarationDocumentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationDocumentCreateRequest(data AppEncryptionDeclarationDocumentCreateRequestData) *AppEncryptionDeclarationDocumentCreateRequest {
	this := AppEncryptionDeclarationDocumentCreateRequest{}
	this.Data = data
	return &this
}

// NewAppEncryptionDeclarationDocumentCreateRequestWithDefaults instantiates a new AppEncryptionDeclarationDocumentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationDocumentCreateRequestWithDefaults() *AppEncryptionDeclarationDocumentCreateRequest {
	this := AppEncryptionDeclarationDocumentCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppEncryptionDeclarationDocumentCreateRequest) GetData() AppEncryptionDeclarationDocumentCreateRequestData {
	if o == nil {
		var ret AppEncryptionDeclarationDocumentCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationDocumentCreateRequest) GetDataOk() (*AppEncryptionDeclarationDocumentCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEncryptionDeclarationDocumentCreateRequest) SetData(v AppEncryptionDeclarationDocumentCreateRequestData) {
	o.Data = v
}

func (o AppEncryptionDeclarationDocumentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationDocumentCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppEncryptionDeclarationDocumentCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppEncryptionDeclarationDocumentCreateRequest := _AppEncryptionDeclarationDocumentCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEncryptionDeclarationDocumentCreateRequest)

	if err != nil {
		return err
	}

	*o = AppEncryptionDeclarationDocumentCreateRequest(varAppEncryptionDeclarationDocumentCreateRequest)

	return err
}

type NullableAppEncryptionDeclarationDocumentCreateRequest struct {
	value *AppEncryptionDeclarationDocumentCreateRequest
	isSet bool
}

func (v NullableAppEncryptionDeclarationDocumentCreateRequest) Get() *AppEncryptionDeclarationDocumentCreateRequest {
	return v.value
}

func (v *NullableAppEncryptionDeclarationDocumentCreateRequest) Set(val *AppEncryptionDeclarationDocumentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationDocumentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationDocumentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationDocumentCreateRequest(val *AppEncryptionDeclarationDocumentCreateRequest) *NullableAppEncryptionDeclarationDocumentCreateRequest {
	return &NullableAppEncryptionDeclarationDocumentCreateRequest{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationDocumentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationDocumentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


