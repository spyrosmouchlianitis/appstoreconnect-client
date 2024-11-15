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

// checks if the SubscriptionImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionImage{}

// SubscriptionImage struct for SubscriptionImage
type SubscriptionImage struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *InAppPurchaseImageAttributes `json:"attributes,omitempty"`
	Relationships *SubscriptionAppStoreReviewScreenshotRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _SubscriptionImage SubscriptionImage

// NewSubscriptionImage instantiates a new SubscriptionImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionImage(type_ string, id string) *SubscriptionImage {
	this := SubscriptionImage{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionImageWithDefaults instantiates a new SubscriptionImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionImageWithDefaults() *SubscriptionImage {
	this := SubscriptionImage{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionImage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionImage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionImage) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionImage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionImage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionImage) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionImage) GetAttributes() InAppPurchaseImageAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret InAppPurchaseImageAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionImage) GetAttributesOk() (*InAppPurchaseImageAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionImage) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given InAppPurchaseImageAttributes and assigns it to the Attributes field.
func (o *SubscriptionImage) SetAttributes(v InAppPurchaseImageAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionImage) GetRelationships() SubscriptionAppStoreReviewScreenshotRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionAppStoreReviewScreenshotRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionImage) GetRelationshipsOk() (*SubscriptionAppStoreReviewScreenshotRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionImage) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionAppStoreReviewScreenshotRelationships and assigns it to the Relationships field.
func (o *SubscriptionImage) SetRelationships(v SubscriptionAppStoreReviewScreenshotRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionImage) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionImage) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionImage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *SubscriptionImage) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o SubscriptionImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *SubscriptionImage) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionImage := _SubscriptionImage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionImage)

	if err != nil {
		return err
	}

	*o = SubscriptionImage(varSubscriptionImage)

	return err
}

type NullableSubscriptionImage struct {
	value *SubscriptionImage
	isSet bool
}

func (v NullableSubscriptionImage) Get() *SubscriptionImage {
	return v.value
}

func (v *NullableSubscriptionImage) Set(val *SubscriptionImage) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionImage) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionImage(val *SubscriptionImage) *NullableSubscriptionImage {
	return &NullableSubscriptionImage{value: val, isSet: true}
}

func (v NullableSubscriptionImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


