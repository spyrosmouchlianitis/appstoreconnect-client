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

// checks if the InAppPurchaseLocalizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseLocalizationResponse{}

// InAppPurchaseLocalizationResponse struct for InAppPurchaseLocalizationResponse
type InAppPurchaseLocalizationResponse struct {
	Data InAppPurchaseLocalization `json:"data"`
	Included []InAppPurchaseV2 `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _InAppPurchaseLocalizationResponse InAppPurchaseLocalizationResponse

// NewInAppPurchaseLocalizationResponse instantiates a new InAppPurchaseLocalizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseLocalizationResponse(data InAppPurchaseLocalization, links DocumentLinks) *InAppPurchaseLocalizationResponse {
	this := InAppPurchaseLocalizationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewInAppPurchaseLocalizationResponseWithDefaults instantiates a new InAppPurchaseLocalizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseLocalizationResponseWithDefaults() *InAppPurchaseLocalizationResponse {
	this := InAppPurchaseLocalizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseLocalizationResponse) GetData() InAppPurchaseLocalization {
	if o == nil {
		var ret InAppPurchaseLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationResponse) GetDataOk() (*InAppPurchaseLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseLocalizationResponse) SetData(v InAppPurchaseLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *InAppPurchaseLocalizationResponse) GetIncluded() []InAppPurchaseV2 {
	if o == nil || IsNil(o.Included) {
		var ret []InAppPurchaseV2
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationResponse) GetIncludedOk() ([]InAppPurchaseV2, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *InAppPurchaseLocalizationResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []InAppPurchaseV2 and assigns it to the Included field.
func (o *InAppPurchaseLocalizationResponse) SetIncluded(v []InAppPurchaseV2) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *InAppPurchaseLocalizationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InAppPurchaseLocalizationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o InAppPurchaseLocalizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseLocalizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *InAppPurchaseLocalizationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varInAppPurchaseLocalizationResponse := _InAppPurchaseLocalizationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInAppPurchaseLocalizationResponse)

	if err != nil {
		return err
	}

	*o = InAppPurchaseLocalizationResponse(varInAppPurchaseLocalizationResponse)

	return err
}

type NullableInAppPurchaseLocalizationResponse struct {
	value *InAppPurchaseLocalizationResponse
	isSet bool
}

func (v NullableInAppPurchaseLocalizationResponse) Get() *InAppPurchaseLocalizationResponse {
	return v.value
}

func (v *NullableInAppPurchaseLocalizationResponse) Set(val *InAppPurchaseLocalizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseLocalizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseLocalizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseLocalizationResponse(val *InAppPurchaseLocalizationResponse) *NullableInAppPurchaseLocalizationResponse {
	return &NullableInAppPurchaseLocalizationResponse{value: val, isSet: true}
}

func (v NullableInAppPurchaseLocalizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseLocalizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


