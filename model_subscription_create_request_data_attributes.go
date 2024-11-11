/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SubscriptionCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCreateRequestDataAttributes{}

// SubscriptionCreateRequestDataAttributes struct for SubscriptionCreateRequestDataAttributes
type SubscriptionCreateRequestDataAttributes struct {
	Name string `json:"name"`
	ProductId string `json:"productId"`
	FamilySharable *bool `json:"familySharable,omitempty"`
	SubscriptionPeriod *string `json:"subscriptionPeriod,omitempty"`
	ReviewNote *string `json:"reviewNote,omitempty"`
	GroupLevel *int32 `json:"groupLevel,omitempty"`
}

type _SubscriptionCreateRequestDataAttributes SubscriptionCreateRequestDataAttributes

// NewSubscriptionCreateRequestDataAttributes instantiates a new SubscriptionCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCreateRequestDataAttributes(name string, productId string) *SubscriptionCreateRequestDataAttributes {
	this := SubscriptionCreateRequestDataAttributes{}
	this.Name = name
	this.ProductId = productId
	return &this
}

// NewSubscriptionCreateRequestDataAttributesWithDefaults instantiates a new SubscriptionCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCreateRequestDataAttributesWithDefaults() *SubscriptionCreateRequestDataAttributes {
	this := SubscriptionCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *SubscriptionCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriptionCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetProductId returns the ProductId field value
func (o *SubscriptionCreateRequestDataAttributes) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequestDataAttributes) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *SubscriptionCreateRequestDataAttributes) SetProductId(v string) {
	o.ProductId = v
}

// GetFamilySharable returns the FamilySharable field value if set, zero value otherwise.
func (o *SubscriptionCreateRequestDataAttributes) GetFamilySharable() bool {
	if o == nil || IsNil(o.FamilySharable) {
		var ret bool
		return ret
	}
	return *o.FamilySharable
}

// GetFamilySharableOk returns a tuple with the FamilySharable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequestDataAttributes) GetFamilySharableOk() (*bool, bool) {
	if o == nil || IsNil(o.FamilySharable) {
		return nil, false
	}
	return o.FamilySharable, true
}

// HasFamilySharable returns a boolean if a field has been set.
func (o *SubscriptionCreateRequestDataAttributes) HasFamilySharable() bool {
	if o != nil && !IsNil(o.FamilySharable) {
		return true
	}

	return false
}

// SetFamilySharable gets a reference to the given bool and assigns it to the FamilySharable field.
func (o *SubscriptionCreateRequestDataAttributes) SetFamilySharable(v bool) {
	o.FamilySharable = &v
}

// GetSubscriptionPeriod returns the SubscriptionPeriod field value if set, zero value otherwise.
func (o *SubscriptionCreateRequestDataAttributes) GetSubscriptionPeriod() string {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		var ret string
		return ret
	}
	return *o.SubscriptionPeriod
}

// GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequestDataAttributes) GetSubscriptionPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		return nil, false
	}
	return o.SubscriptionPeriod, true
}

// HasSubscriptionPeriod returns a boolean if a field has been set.
func (o *SubscriptionCreateRequestDataAttributes) HasSubscriptionPeriod() bool {
	if o != nil && !IsNil(o.SubscriptionPeriod) {
		return true
	}

	return false
}

// SetSubscriptionPeriod gets a reference to the given string and assigns it to the SubscriptionPeriod field.
func (o *SubscriptionCreateRequestDataAttributes) SetSubscriptionPeriod(v string) {
	o.SubscriptionPeriod = &v
}

// GetReviewNote returns the ReviewNote field value if set, zero value otherwise.
func (o *SubscriptionCreateRequestDataAttributes) GetReviewNote() string {
	if o == nil || IsNil(o.ReviewNote) {
		var ret string
		return ret
	}
	return *o.ReviewNote
}

// GetReviewNoteOk returns a tuple with the ReviewNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequestDataAttributes) GetReviewNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewNote) {
		return nil, false
	}
	return o.ReviewNote, true
}

// HasReviewNote returns a boolean if a field has been set.
func (o *SubscriptionCreateRequestDataAttributes) HasReviewNote() bool {
	if o != nil && !IsNil(o.ReviewNote) {
		return true
	}

	return false
}

// SetReviewNote gets a reference to the given string and assigns it to the ReviewNote field.
func (o *SubscriptionCreateRequestDataAttributes) SetReviewNote(v string) {
	o.ReviewNote = &v
}

// GetGroupLevel returns the GroupLevel field value if set, zero value otherwise.
func (o *SubscriptionCreateRequestDataAttributes) GetGroupLevel() int32 {
	if o == nil || IsNil(o.GroupLevel) {
		var ret int32
		return ret
	}
	return *o.GroupLevel
}

// GetGroupLevelOk returns a tuple with the GroupLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateRequestDataAttributes) GetGroupLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupLevel) {
		return nil, false
	}
	return o.GroupLevel, true
}

// HasGroupLevel returns a boolean if a field has been set.
func (o *SubscriptionCreateRequestDataAttributes) HasGroupLevel() bool {
	if o != nil && !IsNil(o.GroupLevel) {
		return true
	}

	return false
}

// SetGroupLevel gets a reference to the given int32 and assigns it to the GroupLevel field.
func (o *SubscriptionCreateRequestDataAttributes) SetGroupLevel(v int32) {
	o.GroupLevel = &v
}

func (o SubscriptionCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["productId"] = o.ProductId
	if !IsNil(o.FamilySharable) {
		toSerialize["familySharable"] = o.FamilySharable
	}
	if !IsNil(o.SubscriptionPeriod) {
		toSerialize["subscriptionPeriod"] = o.SubscriptionPeriod
	}
	if !IsNil(o.ReviewNote) {
		toSerialize["reviewNote"] = o.ReviewNote
	}
	if !IsNil(o.GroupLevel) {
		toSerialize["groupLevel"] = o.GroupLevel
	}
	return toSerialize, nil
}

func (o *SubscriptionCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"productId",
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

	varSubscriptionCreateRequestDataAttributes := _SubscriptionCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = SubscriptionCreateRequestDataAttributes(varSubscriptionCreateRequestDataAttributes)

	return err
}

type NullableSubscriptionCreateRequestDataAttributes struct {
	value *SubscriptionCreateRequestDataAttributes
	isSet bool
}

func (v NullableSubscriptionCreateRequestDataAttributes) Get() *SubscriptionCreateRequestDataAttributes {
	return v.value
}

func (v *NullableSubscriptionCreateRequestDataAttributes) Set(val *SubscriptionCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCreateRequestDataAttributes(val *SubscriptionCreateRequestDataAttributes) *NullableSubscriptionCreateRequestDataAttributes {
	return &NullableSubscriptionCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


