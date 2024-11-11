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

// checks if the ScmProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmProviderResponse{}

// ScmProviderResponse struct for ScmProviderResponse
type ScmProviderResponse struct {
	Data ScmProvider `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _ScmProviderResponse ScmProviderResponse

// NewScmProviderResponse instantiates a new ScmProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmProviderResponse(data ScmProvider, links DocumentLinks) *ScmProviderResponse {
	this := ScmProviderResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewScmProviderResponseWithDefaults instantiates a new ScmProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmProviderResponseWithDefaults() *ScmProviderResponse {
	this := ScmProviderResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ScmProviderResponse) GetData() ScmProvider {
	if o == nil {
		var ret ScmProvider
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ScmProviderResponse) GetDataOk() (*ScmProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ScmProviderResponse) SetData(v ScmProvider) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ScmProviderResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ScmProviderResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ScmProviderResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o ScmProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *ScmProviderResponse) UnmarshalJSON(data []byte) (err error) {
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

	varScmProviderResponse := _ScmProviderResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScmProviderResponse)

	if err != nil {
		return err
	}

	*o = ScmProviderResponse(varScmProviderResponse)

	return err
}

type NullableScmProviderResponse struct {
	value *ScmProviderResponse
	isSet bool
}

func (v NullableScmProviderResponse) Get() *ScmProviderResponse {
	return v.value
}

func (v *NullableScmProviderResponse) Set(val *ScmProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScmProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScmProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmProviderResponse(val *ScmProviderResponse) *NullableScmProviderResponse {
	return &NullableScmProviderResponse{value: val, isSet: true}
}

func (v NullableScmProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


