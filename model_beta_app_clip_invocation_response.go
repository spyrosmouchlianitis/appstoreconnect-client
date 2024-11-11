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

// checks if the BetaAppClipInvocationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationResponse{}

// BetaAppClipInvocationResponse struct for BetaAppClipInvocationResponse
type BetaAppClipInvocationResponse struct {
	Data BetaAppClipInvocation `json:"data"`
	Included []BetaAppClipInvocationLocalization `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _BetaAppClipInvocationResponse BetaAppClipInvocationResponse

// NewBetaAppClipInvocationResponse instantiates a new BetaAppClipInvocationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationResponse(data BetaAppClipInvocation, links DocumentLinks) *BetaAppClipInvocationResponse {
	this := BetaAppClipInvocationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaAppClipInvocationResponseWithDefaults instantiates a new BetaAppClipInvocationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationResponseWithDefaults() *BetaAppClipInvocationResponse {
	this := BetaAppClipInvocationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppClipInvocationResponse) GetData() BetaAppClipInvocation {
	if o == nil {
		var ret BetaAppClipInvocation
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationResponse) GetDataOk() (*BetaAppClipInvocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppClipInvocationResponse) SetData(v BetaAppClipInvocation) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaAppClipInvocationResponse) GetIncluded() []BetaAppClipInvocationLocalization {
	if o == nil || IsNil(o.Included) {
		var ret []BetaAppClipInvocationLocalization
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationResponse) GetIncludedOk() ([]BetaAppClipInvocationLocalization, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaAppClipInvocationResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BetaAppClipInvocationLocalization and assigns it to the Included field.
func (o *BetaAppClipInvocationResponse) SetIncluded(v []BetaAppClipInvocationLocalization) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BetaAppClipInvocationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaAppClipInvocationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaAppClipInvocationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *BetaAppClipInvocationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBetaAppClipInvocationResponse := _BetaAppClipInvocationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaAppClipInvocationResponse)

	if err != nil {
		return err
	}

	*o = BetaAppClipInvocationResponse(varBetaAppClipInvocationResponse)

	return err
}

type NullableBetaAppClipInvocationResponse struct {
	value *BetaAppClipInvocationResponse
	isSet bool
}

func (v NullableBetaAppClipInvocationResponse) Get() *BetaAppClipInvocationResponse {
	return v.value
}

func (v *NullableBetaAppClipInvocationResponse) Set(val *BetaAppClipInvocationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationResponse(val *BetaAppClipInvocationResponse) *NullableBetaAppClipInvocationResponse {
	return &NullableBetaAppClipInvocationResponse{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


