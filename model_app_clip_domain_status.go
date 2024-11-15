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

// checks if the AppClipDomainStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDomainStatus{}

// AppClipDomainStatus struct for AppClipDomainStatus
type AppClipDomainStatus struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipDomainStatusAttributes `json:"attributes,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _AppClipDomainStatus AppClipDomainStatus

// NewAppClipDomainStatus instantiates a new AppClipDomainStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDomainStatus(type_ string, id string) *AppClipDomainStatus {
	this := AppClipDomainStatus{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipDomainStatusWithDefaults instantiates a new AppClipDomainStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDomainStatusWithDefaults() *AppClipDomainStatus {
	this := AppClipDomainStatus{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipDomainStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipDomainStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipDomainStatus) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipDomainStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipDomainStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipDomainStatus) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipDomainStatus) GetAttributes() AppClipDomainStatusAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipDomainStatusAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDomainStatus) GetAttributesOk() (*AppClipDomainStatusAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipDomainStatus) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipDomainStatusAttributes and assigns it to the Attributes field.
func (o *AppClipDomainStatus) SetAttributes(v AppClipDomainStatusAttributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppClipDomainStatus) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDomainStatus) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppClipDomainStatus) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppClipDomainStatus) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppClipDomainStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDomainStatus) ToMap() (map[string]interface{}, error) {
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

func (o *AppClipDomainStatus) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipDomainStatus := _AppClipDomainStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipDomainStatus)

	if err != nil {
		return err
	}

	*o = AppClipDomainStatus(varAppClipDomainStatus)

	return err
}

type NullableAppClipDomainStatus struct {
	value *AppClipDomainStatus
	isSet bool
}

func (v NullableAppClipDomainStatus) Get() *AppClipDomainStatus {
	return v.value
}

func (v *NullableAppClipDomainStatus) Set(val *AppClipDomainStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDomainStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDomainStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDomainStatus(val *AppClipDomainStatus) *NullableAppClipDomainStatus {
	return &NullableAppClipDomainStatus{value: val, isSet: true}
}

func (v NullableAppClipDomainStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDomainStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


