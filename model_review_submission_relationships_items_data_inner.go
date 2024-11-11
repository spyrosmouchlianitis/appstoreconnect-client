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

// checks if the ReviewSubmissionRelationshipsItemsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionRelationshipsItemsDataInner{}

// ReviewSubmissionRelationshipsItemsDataInner struct for ReviewSubmissionRelationshipsItemsDataInner
type ReviewSubmissionRelationshipsItemsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _ReviewSubmissionRelationshipsItemsDataInner ReviewSubmissionRelationshipsItemsDataInner

// NewReviewSubmissionRelationshipsItemsDataInner instantiates a new ReviewSubmissionRelationshipsItemsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionRelationshipsItemsDataInner(type_ string, id string) *ReviewSubmissionRelationshipsItemsDataInner {
	this := ReviewSubmissionRelationshipsItemsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewReviewSubmissionRelationshipsItemsDataInnerWithDefaults instantiates a new ReviewSubmissionRelationshipsItemsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionRelationshipsItemsDataInnerWithDefaults() *ReviewSubmissionRelationshipsItemsDataInner {
	this := ReviewSubmissionRelationshipsItemsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *ReviewSubmissionRelationshipsItemsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionRelationshipsItemsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReviewSubmissionRelationshipsItemsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ReviewSubmissionRelationshipsItemsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionRelationshipsItemsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewSubmissionRelationshipsItemsDataInner) SetId(v string) {
	o.Id = v
}

func (o ReviewSubmissionRelationshipsItemsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionRelationshipsItemsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ReviewSubmissionRelationshipsItemsDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varReviewSubmissionRelationshipsItemsDataInner := _ReviewSubmissionRelationshipsItemsDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReviewSubmissionRelationshipsItemsDataInner)

	if err != nil {
		return err
	}

	*o = ReviewSubmissionRelationshipsItemsDataInner(varReviewSubmissionRelationshipsItemsDataInner)

	return err
}

type NullableReviewSubmissionRelationshipsItemsDataInner struct {
	value *ReviewSubmissionRelationshipsItemsDataInner
	isSet bool
}

func (v NullableReviewSubmissionRelationshipsItemsDataInner) Get() *ReviewSubmissionRelationshipsItemsDataInner {
	return v.value
}

func (v *NullableReviewSubmissionRelationshipsItemsDataInner) Set(val *ReviewSubmissionRelationshipsItemsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionRelationshipsItemsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionRelationshipsItemsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionRelationshipsItemsDataInner(val *ReviewSubmissionRelationshipsItemsDataInner) *NullableReviewSubmissionRelationshipsItemsDataInner {
	return &NullableReviewSubmissionRelationshipsItemsDataInner{value: val, isSet: true}
}

func (v NullableReviewSubmissionRelationshipsItemsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionRelationshipsItemsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


