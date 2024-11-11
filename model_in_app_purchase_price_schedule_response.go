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

// checks if the InAppPurchasePriceScheduleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePriceScheduleResponse{}

// InAppPurchasePriceScheduleResponse struct for InAppPurchasePriceScheduleResponse
type InAppPurchasePriceScheduleResponse struct {
	Data InAppPurchasePriceSchedule `json:"data"`
	Included []InAppPurchasePriceScheduleResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _InAppPurchasePriceScheduleResponse InAppPurchasePriceScheduleResponse

// NewInAppPurchasePriceScheduleResponse instantiates a new InAppPurchasePriceScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePriceScheduleResponse(data InAppPurchasePriceSchedule, links DocumentLinks) *InAppPurchasePriceScheduleResponse {
	this := InAppPurchasePriceScheduleResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewInAppPurchasePriceScheduleResponseWithDefaults instantiates a new InAppPurchasePriceScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceScheduleResponseWithDefaults() *InAppPurchasePriceScheduleResponse {
	this := InAppPurchasePriceScheduleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchasePriceScheduleResponse) GetData() InAppPurchasePriceSchedule {
	if o == nil {
		var ret InAppPurchasePriceSchedule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceScheduleResponse) GetDataOk() (*InAppPurchasePriceSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchasePriceScheduleResponse) SetData(v InAppPurchasePriceSchedule) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *InAppPurchasePriceScheduleResponse) GetIncluded() []InAppPurchasePriceScheduleResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []InAppPurchasePriceScheduleResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceScheduleResponse) GetIncludedOk() ([]InAppPurchasePriceScheduleResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *InAppPurchasePriceScheduleResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []InAppPurchasePriceScheduleResponseIncludedInner and assigns it to the Included field.
func (o *InAppPurchasePriceScheduleResponse) SetIncluded(v []InAppPurchasePriceScheduleResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *InAppPurchasePriceScheduleResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceScheduleResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InAppPurchasePriceScheduleResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o InAppPurchasePriceScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePriceScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *InAppPurchasePriceScheduleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varInAppPurchasePriceScheduleResponse := _InAppPurchasePriceScheduleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInAppPurchasePriceScheduleResponse)

	if err != nil {
		return err
	}

	*o = InAppPurchasePriceScheduleResponse(varInAppPurchasePriceScheduleResponse)

	return err
}

type NullableInAppPurchasePriceScheduleResponse struct {
	value *InAppPurchasePriceScheduleResponse
	isSet bool
}

func (v NullableInAppPurchasePriceScheduleResponse) Get() *InAppPurchasePriceScheduleResponse {
	return v.value
}

func (v *NullableInAppPurchasePriceScheduleResponse) Set(val *InAppPurchasePriceScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceScheduleResponse(val *InAppPurchasePriceScheduleResponse) *NullableInAppPurchasePriceScheduleResponse {
	return &NullableInAppPurchasePriceScheduleResponse{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


