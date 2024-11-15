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

// checks if the InAppPurchaseV2CreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2CreateRequestDataAttributes{}

// InAppPurchaseV2CreateRequestDataAttributes struct for InAppPurchaseV2CreateRequestDataAttributes
type InAppPurchaseV2CreateRequestDataAttributes struct {
	Name string `json:"name"`
	ProductId string `json:"productId"`
	InAppPurchaseType InAppPurchaseType `json:"inAppPurchaseType"`
	ReviewNote *string `json:"reviewNote,omitempty"`
	FamilySharable *bool `json:"familySharable,omitempty"`
}

type _InAppPurchaseV2CreateRequestDataAttributes InAppPurchaseV2CreateRequestDataAttributes

// NewInAppPurchaseV2CreateRequestDataAttributes instantiates a new InAppPurchaseV2CreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2CreateRequestDataAttributes(name string, productId string, inAppPurchaseType InAppPurchaseType) *InAppPurchaseV2CreateRequestDataAttributes {
	this := InAppPurchaseV2CreateRequestDataAttributes{}
	this.Name = name
	this.ProductId = productId
	this.InAppPurchaseType = inAppPurchaseType
	return &this
}

// NewInAppPurchaseV2CreateRequestDataAttributesWithDefaults instantiates a new InAppPurchaseV2CreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2CreateRequestDataAttributesWithDefaults() *InAppPurchaseV2CreateRequestDataAttributes {
	this := InAppPurchaseV2CreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InAppPurchaseV2CreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetProductId returns the ProductId field value
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *InAppPurchaseV2CreateRequestDataAttributes) SetProductId(v string) {
	o.ProductId = v
}

// GetInAppPurchaseType returns the InAppPurchaseType field value
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetInAppPurchaseType() InAppPurchaseType {
	if o == nil {
		var ret InAppPurchaseType
		return ret
	}

	return o.InAppPurchaseType
}

// GetInAppPurchaseTypeOk returns a tuple with the InAppPurchaseType field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetInAppPurchaseTypeOk() (*InAppPurchaseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InAppPurchaseType, true
}

// SetInAppPurchaseType sets field value
func (o *InAppPurchaseV2CreateRequestDataAttributes) SetInAppPurchaseType(v InAppPurchaseType) {
	o.InAppPurchaseType = v
}

// GetReviewNote returns the ReviewNote field value if set, zero value otherwise.
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetReviewNote() string {
	if o == nil || IsNil(o.ReviewNote) {
		var ret string
		return ret
	}
	return *o.ReviewNote
}

// GetReviewNoteOk returns a tuple with the ReviewNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetReviewNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewNote) {
		return nil, false
	}
	return o.ReviewNote, true
}

// HasReviewNote returns a boolean if a field has been set.
func (o *InAppPurchaseV2CreateRequestDataAttributes) HasReviewNote() bool {
	if o != nil && !IsNil(o.ReviewNote) {
		return true
	}

	return false
}

// SetReviewNote gets a reference to the given string and assigns it to the ReviewNote field.
func (o *InAppPurchaseV2CreateRequestDataAttributes) SetReviewNote(v string) {
	o.ReviewNote = &v
}

// GetFamilySharable returns the FamilySharable field value if set, zero value otherwise.
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetFamilySharable() bool {
	if o == nil || IsNil(o.FamilySharable) {
		var ret bool
		return ret
	}
	return *o.FamilySharable
}

// GetFamilySharableOk returns a tuple with the FamilySharable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2CreateRequestDataAttributes) GetFamilySharableOk() (*bool, bool) {
	if o == nil || IsNil(o.FamilySharable) {
		return nil, false
	}
	return o.FamilySharable, true
}

// HasFamilySharable returns a boolean if a field has been set.
func (o *InAppPurchaseV2CreateRequestDataAttributes) HasFamilySharable() bool {
	if o != nil && !IsNil(o.FamilySharable) {
		return true
	}

	return false
}

// SetFamilySharable gets a reference to the given bool and assigns it to the FamilySharable field.
func (o *InAppPurchaseV2CreateRequestDataAttributes) SetFamilySharable(v bool) {
	o.FamilySharable = &v
}

func (o InAppPurchaseV2CreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2CreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["productId"] = o.ProductId
	toSerialize["inAppPurchaseType"] = o.InAppPurchaseType
	if !IsNil(o.ReviewNote) {
		toSerialize["reviewNote"] = o.ReviewNote
	}
	if !IsNil(o.FamilySharable) {
		toSerialize["familySharable"] = o.FamilySharable
	}
	return toSerialize, nil
}

func (o *InAppPurchaseV2CreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"productId",
		"inAppPurchaseType",
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

	varInAppPurchaseV2CreateRequestDataAttributes := _InAppPurchaseV2CreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInAppPurchaseV2CreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = InAppPurchaseV2CreateRequestDataAttributes(varInAppPurchaseV2CreateRequestDataAttributes)

	return err
}

type NullableInAppPurchaseV2CreateRequestDataAttributes struct {
	value *InAppPurchaseV2CreateRequestDataAttributes
	isSet bool
}

func (v NullableInAppPurchaseV2CreateRequestDataAttributes) Get() *InAppPurchaseV2CreateRequestDataAttributes {
	return v.value
}

func (v *NullableInAppPurchaseV2CreateRequestDataAttributes) Set(val *InAppPurchaseV2CreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2CreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2CreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2CreateRequestDataAttributes(val *InAppPurchaseV2CreateRequestDataAttributes) *NullableInAppPurchaseV2CreateRequestDataAttributes {
	return &NullableInAppPurchaseV2CreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2CreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2CreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


