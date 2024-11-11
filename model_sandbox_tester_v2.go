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

// checks if the SandboxTesterV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTesterV2{}

// SandboxTesterV2 struct for SandboxTesterV2
type SandboxTesterV2 struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *SandboxTesterV2Attributes `json:"attributes,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _SandboxTesterV2 SandboxTesterV2

// NewSandboxTesterV2 instantiates a new SandboxTesterV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTesterV2(type_ string, id string) *SandboxTesterV2 {
	this := SandboxTesterV2{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSandboxTesterV2WithDefaults instantiates a new SandboxTesterV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTesterV2WithDefaults() *SandboxTesterV2 {
	this := SandboxTesterV2{}
	return &this
}

// GetType returns the Type field value
func (o *SandboxTesterV2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SandboxTesterV2) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SandboxTesterV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SandboxTesterV2) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SandboxTesterV2) GetAttributes() SandboxTesterV2Attributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SandboxTesterV2Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2) GetAttributesOk() (*SandboxTesterV2Attributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SandboxTesterV2) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SandboxTesterV2Attributes and assigns it to the Attributes field.
func (o *SandboxTesterV2) SetAttributes(v SandboxTesterV2Attributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SandboxTesterV2) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SandboxTesterV2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *SandboxTesterV2) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o SandboxTesterV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTesterV2) ToMap() (map[string]interface{}, error) {
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

func (o *SandboxTesterV2) UnmarshalJSON(data []byte) (err error) {
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

	varSandboxTesterV2 := _SandboxTesterV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSandboxTesterV2)

	if err != nil {
		return err
	}

	*o = SandboxTesterV2(varSandboxTesterV2)

	return err
}

type NullableSandboxTesterV2 struct {
	value *SandboxTesterV2
	isSet bool
}

func (v NullableSandboxTesterV2) Get() *SandboxTesterV2 {
	return v.value
}

func (v *NullableSandboxTesterV2) Set(val *SandboxTesterV2) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTesterV2) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTesterV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTesterV2(val *SandboxTesterV2) *NullableSandboxTesterV2 {
	return &NullableSandboxTesterV2{value: val, isSet: true}
}

func (v NullableSandboxTesterV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTesterV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


