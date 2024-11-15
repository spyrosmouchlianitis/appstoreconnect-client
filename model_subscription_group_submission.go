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

// checks if the SubscriptionGroupSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupSubmission{}

// SubscriptionGroupSubmission struct for SubscriptionGroupSubmission
type SubscriptionGroupSubmission struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _SubscriptionGroupSubmission SubscriptionGroupSubmission

// NewSubscriptionGroupSubmission instantiates a new SubscriptionGroupSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupSubmission(type_ string, id string) *SubscriptionGroupSubmission {
	this := SubscriptionGroupSubmission{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionGroupSubmissionWithDefaults instantiates a new SubscriptionGroupSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupSubmissionWithDefaults() *SubscriptionGroupSubmission {
	this := SubscriptionGroupSubmission{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionGroupSubmission) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupSubmission) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionGroupSubmission) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionGroupSubmission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupSubmission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionGroupSubmission) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionGroupSubmission) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupSubmission) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionGroupSubmission) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *SubscriptionGroupSubmission) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o SubscriptionGroupSubmission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupSubmission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *SubscriptionGroupSubmission) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionGroupSubmission := _SubscriptionGroupSubmission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionGroupSubmission)

	if err != nil {
		return err
	}

	*o = SubscriptionGroupSubmission(varSubscriptionGroupSubmission)

	return err
}

type NullableSubscriptionGroupSubmission struct {
	value *SubscriptionGroupSubmission
	isSet bool
}

func (v NullableSubscriptionGroupSubmission) Get() *SubscriptionGroupSubmission {
	return v.value
}

func (v *NullableSubscriptionGroupSubmission) Set(val *SubscriptionGroupSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupSubmission(val *SubscriptionGroupSubmission) *NullableSubscriptionGroupSubmission {
	return &NullableSubscriptionGroupSubmission{value: val, isSet: true}
}

func (v NullableSubscriptionGroupSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


