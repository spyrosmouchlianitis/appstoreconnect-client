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

// checks if the InAppPurchaseAppStoreReviewScreenshotCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAppStoreReviewScreenshotCreateRequestData{}

// InAppPurchaseAppStoreReviewScreenshotCreateRequestData struct for InAppPurchaseAppStoreReviewScreenshotCreateRequestData
type InAppPurchaseAppStoreReviewScreenshotCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes `json:"attributes"`
	Relationships InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships `json:"relationships"`
}

type _InAppPurchaseAppStoreReviewScreenshotCreateRequestData InAppPurchaseAppStoreReviewScreenshotCreateRequestData

// NewInAppPurchaseAppStoreReviewScreenshotCreateRequestData instantiates a new InAppPurchaseAppStoreReviewScreenshotCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAppStoreReviewScreenshotCreateRequestData(type_ string, attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes, relationships InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships) *InAppPurchaseAppStoreReviewScreenshotCreateRequestData {
	this := InAppPurchaseAppStoreReviewScreenshotCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewInAppPurchaseAppStoreReviewScreenshotCreateRequestDataWithDefaults instantiates a new InAppPurchaseAppStoreReviewScreenshotCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAppStoreReviewScreenshotCreateRequestDataWithDefaults() *InAppPurchaseAppStoreReviewScreenshotCreateRequestData {
	this := InAppPurchaseAppStoreReviewScreenshotCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) GetAttributes() AppClipAdvancedExperienceImageCreateRequestDataAttributes {
	if o == nil {
		var ret AppClipAdvancedExperienceImageCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) SetAttributes(v AppClipAdvancedExperienceImageCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) GetRelationships() InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships {
	if o == nil {
		var ret InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) GetRelationshipsOk() (*InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) SetRelationships(v InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o InAppPurchaseAppStoreReviewScreenshotCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAppStoreReviewScreenshotCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varInAppPurchaseAppStoreReviewScreenshotCreateRequestData := _InAppPurchaseAppStoreReviewScreenshotCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInAppPurchaseAppStoreReviewScreenshotCreateRequestData)

	if err != nil {
		return err
	}

	*o = InAppPurchaseAppStoreReviewScreenshotCreateRequestData(varInAppPurchaseAppStoreReviewScreenshotCreateRequestData)

	return err
}

type NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData struct {
	value *InAppPurchaseAppStoreReviewScreenshotCreateRequestData
	isSet bool
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData) Get() *InAppPurchaseAppStoreReviewScreenshotCreateRequestData {
	return v.value
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData) Set(val *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData(val *InAppPurchaseAppStoreReviewScreenshotCreateRequestData) *NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData {
	return &NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData{value: val, isSet: true}
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


