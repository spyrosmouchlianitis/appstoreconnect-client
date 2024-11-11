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

// checks if the CiMacOsVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiMacOsVersionsResponse{}

// CiMacOsVersionsResponse struct for CiMacOsVersionsResponse
type CiMacOsVersionsResponse struct {
	Data []CiMacOsVersion `json:"data"`
	Included []CiXcodeVersion `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _CiMacOsVersionsResponse CiMacOsVersionsResponse

// NewCiMacOsVersionsResponse instantiates a new CiMacOsVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiMacOsVersionsResponse(data []CiMacOsVersion, links PagedDocumentLinks) *CiMacOsVersionsResponse {
	this := CiMacOsVersionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiMacOsVersionsResponseWithDefaults instantiates a new CiMacOsVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiMacOsVersionsResponseWithDefaults() *CiMacOsVersionsResponse {
	this := CiMacOsVersionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiMacOsVersionsResponse) GetData() []CiMacOsVersion {
	if o == nil {
		var ret []CiMacOsVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionsResponse) GetDataOk() ([]CiMacOsVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CiMacOsVersionsResponse) SetData(v []CiMacOsVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CiMacOsVersionsResponse) GetIncluded() []CiXcodeVersion {
	if o == nil || IsNil(o.Included) {
		var ret []CiXcodeVersion
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionsResponse) GetIncludedOk() ([]CiXcodeVersion, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CiMacOsVersionsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CiXcodeVersion and assigns it to the Included field.
func (o *CiMacOsVersionsResponse) SetIncluded(v []CiXcodeVersion) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *CiMacOsVersionsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiMacOsVersionsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CiMacOsVersionsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CiMacOsVersionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *CiMacOsVersionsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o CiMacOsVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiMacOsVersionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *CiMacOsVersionsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCiMacOsVersionsResponse := _CiMacOsVersionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiMacOsVersionsResponse)

	if err != nil {
		return err
	}

	*o = CiMacOsVersionsResponse(varCiMacOsVersionsResponse)

	return err
}

type NullableCiMacOsVersionsResponse struct {
	value *CiMacOsVersionsResponse
	isSet bool
}

func (v NullableCiMacOsVersionsResponse) Get() *CiMacOsVersionsResponse {
	return v.value
}

func (v *NullableCiMacOsVersionsResponse) Set(val *CiMacOsVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiMacOsVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiMacOsVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiMacOsVersionsResponse(val *CiMacOsVersionsResponse) *NullableCiMacOsVersionsResponse {
	return &NullableCiMacOsVersionsResponse{value: val, isSet: true}
}

func (v NullableCiMacOsVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiMacOsVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

