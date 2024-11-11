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

// checks if the AgeRatingDeclarationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgeRatingDeclarationResponse{}

// AgeRatingDeclarationResponse struct for AgeRatingDeclarationResponse
type AgeRatingDeclarationResponse struct {
	Data AgeRatingDeclaration `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AgeRatingDeclarationResponse AgeRatingDeclarationResponse

// NewAgeRatingDeclarationResponse instantiates a new AgeRatingDeclarationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgeRatingDeclarationResponse(data AgeRatingDeclaration, links DocumentLinks) *AgeRatingDeclarationResponse {
	this := AgeRatingDeclarationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAgeRatingDeclarationResponseWithDefaults instantiates a new AgeRatingDeclarationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeRatingDeclarationResponseWithDefaults() *AgeRatingDeclarationResponse {
	this := AgeRatingDeclarationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AgeRatingDeclarationResponse) GetData() AgeRatingDeclaration {
	if o == nil {
		var ret AgeRatingDeclaration
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationResponse) GetDataOk() (*AgeRatingDeclaration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgeRatingDeclarationResponse) SetData(v AgeRatingDeclaration) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AgeRatingDeclarationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AgeRatingDeclarationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AgeRatingDeclarationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgeRatingDeclarationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AgeRatingDeclarationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAgeRatingDeclarationResponse := _AgeRatingDeclarationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgeRatingDeclarationResponse)

	if err != nil {
		return err
	}

	*o = AgeRatingDeclarationResponse(varAgeRatingDeclarationResponse)

	return err
}

type NullableAgeRatingDeclarationResponse struct {
	value *AgeRatingDeclarationResponse
	isSet bool
}

func (v NullableAgeRatingDeclarationResponse) Get() *AgeRatingDeclarationResponse {
	return v.value
}

func (v *NullableAgeRatingDeclarationResponse) Set(val *AgeRatingDeclarationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgeRatingDeclarationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgeRatingDeclarationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgeRatingDeclarationResponse(val *AgeRatingDeclarationResponse) *NullableAgeRatingDeclarationResponse {
	return &NullableAgeRatingDeclarationResponse{value: val, isSet: true}
}

func (v NullableAgeRatingDeclarationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgeRatingDeclarationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


