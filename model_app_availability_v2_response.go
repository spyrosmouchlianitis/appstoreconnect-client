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

// checks if the AppAvailabilityV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityV2Response{}

// AppAvailabilityV2Response struct for AppAvailabilityV2Response
type AppAvailabilityV2Response struct {
	Data AppAvailabilityV2 `json:"data"`
	Included []TerritoryAvailability `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _AppAvailabilityV2Response AppAvailabilityV2Response

// NewAppAvailabilityV2Response instantiates a new AppAvailabilityV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityV2Response(data AppAvailabilityV2, links DocumentLinks) *AppAvailabilityV2Response {
	this := AppAvailabilityV2Response{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppAvailabilityV2ResponseWithDefaults instantiates a new AppAvailabilityV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityV2ResponseWithDefaults() *AppAvailabilityV2Response {
	this := AppAvailabilityV2Response{}
	return &this
}

// GetData returns the Data field value
func (o *AppAvailabilityV2Response) GetData() AppAvailabilityV2 {
	if o == nil {
		var ret AppAvailabilityV2
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2Response) GetDataOk() (*AppAvailabilityV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppAvailabilityV2Response) SetData(v AppAvailabilityV2) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppAvailabilityV2Response) GetIncluded() []TerritoryAvailability {
	if o == nil || IsNil(o.Included) {
		var ret []TerritoryAvailability
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2Response) GetIncludedOk() ([]TerritoryAvailability, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppAvailabilityV2Response) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []TerritoryAvailability and assigns it to the Included field.
func (o *AppAvailabilityV2Response) SetIncluded(v []TerritoryAvailability) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppAvailabilityV2Response) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2Response) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppAvailabilityV2Response) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppAvailabilityV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppAvailabilityV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varAppAvailabilityV2Response := _AppAvailabilityV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppAvailabilityV2Response)

	if err != nil {
		return err
	}

	*o = AppAvailabilityV2Response(varAppAvailabilityV2Response)

	return err
}

type NullableAppAvailabilityV2Response struct {
	value *AppAvailabilityV2Response
	isSet bool
}

func (v NullableAppAvailabilityV2Response) Get() *AppAvailabilityV2Response {
	return v.value
}

func (v *NullableAppAvailabilityV2Response) Set(val *AppAvailabilityV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityV2Response(val *AppAvailabilityV2Response) *NullableAppAvailabilityV2Response {
	return &NullableAppAvailabilityV2Response{value: val, isSet: true}
}

func (v NullableAppAvailabilityV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

