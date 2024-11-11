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

// checks if the AnalyticsReportSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsReportSegment{}

// AnalyticsReportSegment struct for AnalyticsReportSegment
type AnalyticsReportSegment struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AnalyticsReportSegmentAttributes `json:"attributes,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _AnalyticsReportSegment AnalyticsReportSegment

// NewAnalyticsReportSegment instantiates a new AnalyticsReportSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsReportSegment(type_ string, id string) *AnalyticsReportSegment {
	this := AnalyticsReportSegment{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAnalyticsReportSegmentWithDefaults instantiates a new AnalyticsReportSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsReportSegmentWithDefaults() *AnalyticsReportSegment {
	this := AnalyticsReportSegment{}
	return &this
}

// GetType returns the Type field value
func (o *AnalyticsReportSegment) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportSegment) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AnalyticsReportSegment) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AnalyticsReportSegment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportSegment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AnalyticsReportSegment) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AnalyticsReportSegment) GetAttributes() AnalyticsReportSegmentAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AnalyticsReportSegmentAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportSegment) GetAttributesOk() (*AnalyticsReportSegmentAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AnalyticsReportSegment) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AnalyticsReportSegmentAttributes and assigns it to the Attributes field.
func (o *AnalyticsReportSegment) SetAttributes(v AnalyticsReportSegmentAttributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AnalyticsReportSegment) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportSegment) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AnalyticsReportSegment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AnalyticsReportSegment) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AnalyticsReportSegment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsReportSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *AnalyticsReportSegment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varAnalyticsReportSegment := _AnalyticsReportSegment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalyticsReportSegment)

	if err != nil {
		return err
	}

	*o = AnalyticsReportSegment(varAnalyticsReportSegment)

	return err
}

type NullableAnalyticsReportSegment struct {
	value *AnalyticsReportSegment
	isSet bool
}

func (v NullableAnalyticsReportSegment) Get() *AnalyticsReportSegment {
	return v.value
}

func (v *NullableAnalyticsReportSegment) Set(val *AnalyticsReportSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsReportSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsReportSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsReportSegment(val *AnalyticsReportSegment) *NullableAnalyticsReportSegment {
	return &NullableAnalyticsReportSegment{value: val, isSet: true}
}

func (v NullableAnalyticsReportSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsReportSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


