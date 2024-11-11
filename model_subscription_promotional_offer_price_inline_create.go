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

// checks if the SubscriptionPromotionalOfferPriceInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferPriceInlineCreate{}

// SubscriptionPromotionalOfferPriceInlineCreate struct for SubscriptionPromotionalOfferPriceInlineCreate
type SubscriptionPromotionalOfferPriceInlineCreate struct {
	Type string `json:"type"`
	Id *string `json:"id,omitempty"`
	Relationships *SubscriptionOfferCodePriceRelationships `json:"relationships,omitempty"`
}

type _SubscriptionPromotionalOfferPriceInlineCreate SubscriptionPromotionalOfferPriceInlineCreate

// NewSubscriptionPromotionalOfferPriceInlineCreate instantiates a new SubscriptionPromotionalOfferPriceInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferPriceInlineCreate(type_ string) *SubscriptionPromotionalOfferPriceInlineCreate {
	this := SubscriptionPromotionalOfferPriceInlineCreate{}
	this.Type = type_
	return &this
}

// NewSubscriptionPromotionalOfferPriceInlineCreateWithDefaults instantiates a new SubscriptionPromotionalOfferPriceInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferPriceInlineCreateWithDefaults() *SubscriptionPromotionalOfferPriceInlineCreate {
	this := SubscriptionPromotionalOfferPriceInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionPromotionalOfferPriceInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionPromotionalOfferPriceInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) GetRelationships() SubscriptionOfferCodePriceRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionOfferCodePriceRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) GetRelationshipsOk() (*SubscriptionOfferCodePriceRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionOfferCodePriceRelationships and assigns it to the Relationships field.
func (o *SubscriptionPromotionalOfferPriceInlineCreate) SetRelationships(v SubscriptionOfferCodePriceRelationships) {
	o.Relationships = &v
}

func (o SubscriptionPromotionalOfferPriceInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferPriceInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *SubscriptionPromotionalOfferPriceInlineCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varSubscriptionPromotionalOfferPriceInlineCreate := _SubscriptionPromotionalOfferPriceInlineCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPromotionalOfferPriceInlineCreate)

	if err != nil {
		return err
	}

	*o = SubscriptionPromotionalOfferPriceInlineCreate(varSubscriptionPromotionalOfferPriceInlineCreate)

	return err
}

type NullableSubscriptionPromotionalOfferPriceInlineCreate struct {
	value *SubscriptionPromotionalOfferPriceInlineCreate
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferPriceInlineCreate) Get() *SubscriptionPromotionalOfferPriceInlineCreate {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferPriceInlineCreate) Set(val *SubscriptionPromotionalOfferPriceInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferPriceInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferPriceInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferPriceInlineCreate(val *SubscriptionPromotionalOfferPriceInlineCreate) *NullableSubscriptionPromotionalOfferPriceInlineCreate {
	return &NullableSubscriptionPromotionalOfferPriceInlineCreate{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferPriceInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferPriceInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


