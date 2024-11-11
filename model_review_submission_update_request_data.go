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

// checks if the ReviewSubmissionUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionUpdateRequestData{}

// ReviewSubmissionUpdateRequestData struct for ReviewSubmissionUpdateRequestData
type ReviewSubmissionUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *ReviewSubmissionUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

type _ReviewSubmissionUpdateRequestData ReviewSubmissionUpdateRequestData

// NewReviewSubmissionUpdateRequestData instantiates a new ReviewSubmissionUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionUpdateRequestData(type_ string, id string) *ReviewSubmissionUpdateRequestData {
	this := ReviewSubmissionUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewReviewSubmissionUpdateRequestDataWithDefaults instantiates a new ReviewSubmissionUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionUpdateRequestDataWithDefaults() *ReviewSubmissionUpdateRequestData {
	this := ReviewSubmissionUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *ReviewSubmissionUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReviewSubmissionUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ReviewSubmissionUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewSubmissionUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ReviewSubmissionUpdateRequestData) GetAttributes() ReviewSubmissionUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ReviewSubmissionUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionUpdateRequestData) GetAttributesOk() (*ReviewSubmissionUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ReviewSubmissionUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ReviewSubmissionUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *ReviewSubmissionUpdateRequestData) SetAttributes(v ReviewSubmissionUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o ReviewSubmissionUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *ReviewSubmissionUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varReviewSubmissionUpdateRequestData := _ReviewSubmissionUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReviewSubmissionUpdateRequestData)

	if err != nil {
		return err
	}

	*o = ReviewSubmissionUpdateRequestData(varReviewSubmissionUpdateRequestData)

	return err
}

type NullableReviewSubmissionUpdateRequestData struct {
	value *ReviewSubmissionUpdateRequestData
	isSet bool
}

func (v NullableReviewSubmissionUpdateRequestData) Get() *ReviewSubmissionUpdateRequestData {
	return v.value
}

func (v *NullableReviewSubmissionUpdateRequestData) Set(val *ReviewSubmissionUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionUpdateRequestData(val *ReviewSubmissionUpdateRequestData) *NullableReviewSubmissionUpdateRequestData {
	return &NullableReviewSubmissionUpdateRequestData{value: val, isSet: true}
}

func (v NullableReviewSubmissionUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


