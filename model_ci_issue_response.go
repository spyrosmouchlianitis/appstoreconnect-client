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

// checks if the CiIssueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiIssueResponse{}

// CiIssueResponse struct for CiIssueResponse
type CiIssueResponse struct {
	Data CiIssue `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _CiIssueResponse CiIssueResponse

// NewCiIssueResponse instantiates a new CiIssueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiIssueResponse(data CiIssue, links DocumentLinks) *CiIssueResponse {
	this := CiIssueResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiIssueResponseWithDefaults instantiates a new CiIssueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiIssueResponseWithDefaults() *CiIssueResponse {
	this := CiIssueResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiIssueResponse) GetData() CiIssue {
	if o == nil {
		var ret CiIssue
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiIssueResponse) GetDataOk() (*CiIssue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiIssueResponse) SetData(v CiIssue) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *CiIssueResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiIssueResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiIssueResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o CiIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiIssueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CiIssueResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCiIssueResponse := _CiIssueResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiIssueResponse)

	if err != nil {
		return err
	}

	*o = CiIssueResponse(varCiIssueResponse)

	return err
}

type NullableCiIssueResponse struct {
	value *CiIssueResponse
	isSet bool
}

func (v NullableCiIssueResponse) Get() *CiIssueResponse {
	return v.value
}

func (v *NullableCiIssueResponse) Set(val *CiIssueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiIssueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiIssueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiIssueResponse(val *CiIssueResponse) *NullableCiIssueResponse {
	return &NullableCiIssueResponse{value: val, isSet: true}
}

func (v NullableCiIssueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiIssueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


