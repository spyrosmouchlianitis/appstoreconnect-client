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

// checks if the CiXcodeVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiXcodeVersionResponse{}

// CiXcodeVersionResponse struct for CiXcodeVersionResponse
type CiXcodeVersionResponse struct {
	Data CiXcodeVersion `json:"data"`
	Included []CiMacOsVersion `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _CiXcodeVersionResponse CiXcodeVersionResponse

// NewCiXcodeVersionResponse instantiates a new CiXcodeVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiXcodeVersionResponse(data CiXcodeVersion, links DocumentLinks) *CiXcodeVersionResponse {
	this := CiXcodeVersionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiXcodeVersionResponseWithDefaults instantiates a new CiXcodeVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiXcodeVersionResponseWithDefaults() *CiXcodeVersionResponse {
	this := CiXcodeVersionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiXcodeVersionResponse) GetData() CiXcodeVersion {
	if o == nil {
		var ret CiXcodeVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionResponse) GetDataOk() (*CiXcodeVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiXcodeVersionResponse) SetData(v CiXcodeVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CiXcodeVersionResponse) GetIncluded() []CiMacOsVersion {
	if o == nil || IsNil(o.Included) {
		var ret []CiMacOsVersion
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionResponse) GetIncludedOk() ([]CiMacOsVersion, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CiXcodeVersionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CiMacOsVersion and assigns it to the Included field.
func (o *CiXcodeVersionResponse) SetIncluded(v []CiMacOsVersion) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *CiXcodeVersionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiXcodeVersionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o CiXcodeVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiXcodeVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CiXcodeVersionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCiXcodeVersionResponse := _CiXcodeVersionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiXcodeVersionResponse)

	if err != nil {
		return err
	}

	*o = CiXcodeVersionResponse(varCiXcodeVersionResponse)

	return err
}

type NullableCiXcodeVersionResponse struct {
	value *CiXcodeVersionResponse
	isSet bool
}

func (v NullableCiXcodeVersionResponse) Get() *CiXcodeVersionResponse {
	return v.value
}

func (v *NullableCiXcodeVersionResponse) Set(val *CiXcodeVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiXcodeVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiXcodeVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiXcodeVersionResponse(val *CiXcodeVersionResponse) *NullableCiXcodeVersionResponse {
	return &NullableCiXcodeVersionResponse{value: val, isSet: true}
}

func (v NullableCiXcodeVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiXcodeVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


