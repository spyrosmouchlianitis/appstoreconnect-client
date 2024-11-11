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

// checks if the SubscriptionLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionLocalizationCreateRequestData{}

// SubscriptionLocalizationCreateRequestData struct for SubscriptionLocalizationCreateRequestData
type SubscriptionLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes InAppPurchaseLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships `json:"relationships"`
}

type _SubscriptionLocalizationCreateRequestData SubscriptionLocalizationCreateRequestData

// NewSubscriptionLocalizationCreateRequestData instantiates a new SubscriptionLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionLocalizationCreateRequestData(type_ string, attributes InAppPurchaseLocalizationCreateRequestDataAttributes, relationships SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships) *SubscriptionLocalizationCreateRequestData {
	this := SubscriptionLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewSubscriptionLocalizationCreateRequestDataWithDefaults instantiates a new SubscriptionLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionLocalizationCreateRequestDataWithDefaults() *SubscriptionLocalizationCreateRequestData {
	this := SubscriptionLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionLocalizationCreateRequestData) GetAttributes() InAppPurchaseLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret InAppPurchaseLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionLocalizationCreateRequestData) GetAttributesOk() (*InAppPurchaseLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionLocalizationCreateRequestData) SetAttributes(v InAppPurchaseLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *SubscriptionLocalizationCreateRequestData) GetRelationships() SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships {
	if o == nil {
		var ret SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SubscriptionLocalizationCreateRequestData) GetRelationshipsOk() (*SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SubscriptionLocalizationCreateRequestData) SetRelationships(v SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SubscriptionLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *SubscriptionLocalizationCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
		"relationships",
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

	varSubscriptionLocalizationCreateRequestData := _SubscriptionLocalizationCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionLocalizationCreateRequestData)

	if err != nil {
		return err
	}

	*o = SubscriptionLocalizationCreateRequestData(varSubscriptionLocalizationCreateRequestData)

	return err
}

type NullableSubscriptionLocalizationCreateRequestData struct {
	value *SubscriptionLocalizationCreateRequestData
	isSet bool
}

func (v NullableSubscriptionLocalizationCreateRequestData) Get() *SubscriptionLocalizationCreateRequestData {
	return v.value
}

func (v *NullableSubscriptionLocalizationCreateRequestData) Set(val *SubscriptionLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionLocalizationCreateRequestData(val *SubscriptionLocalizationCreateRequestData) *NullableSubscriptionLocalizationCreateRequestData {
	return &NullableSubscriptionLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


