/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the InAppPurchaseV2Attributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2Attributes{}

// InAppPurchaseV2Attributes struct for InAppPurchaseV2Attributes
type InAppPurchaseV2Attributes struct {
	Name *string `json:"name,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	InAppPurchaseType *InAppPurchaseType `json:"inAppPurchaseType,omitempty"`
	State *InAppPurchaseState `json:"state,omitempty"`
	ReviewNote *string `json:"reviewNote,omitempty"`
	FamilySharable *bool `json:"familySharable,omitempty"`
	ContentHosting *bool `json:"contentHosting,omitempty"`
}

// NewInAppPurchaseV2Attributes instantiates a new InAppPurchaseV2Attributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2Attributes() *InAppPurchaseV2Attributes {
	this := InAppPurchaseV2Attributes{}
	return &this
}

// NewInAppPurchaseV2AttributesWithDefaults instantiates a new InAppPurchaseV2Attributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2AttributesWithDefaults() *InAppPurchaseV2Attributes {
	this := InAppPurchaseV2Attributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InAppPurchaseV2Attributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Attributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InAppPurchaseV2Attributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InAppPurchaseV2Attributes) SetName(v string) {
	o.Name = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *InAppPurchaseV2Attributes) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Attributes) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *InAppPurchaseV2Attributes) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *InAppPurchaseV2Attributes) SetProductId(v string) {
	o.ProductId = &v
}

// GetInAppPurchaseType returns the InAppPurchaseType field value if set, zero value otherwise.
func (o *InAppPurchaseV2Attributes) GetInAppPurchaseType() InAppPurchaseType {
	if o == nil || IsNil(o.InAppPurchaseType) {
		var ret InAppPurchaseType
		return ret
	}
	return *o.InAppPurchaseType
}

// GetInAppPurchaseTypeOk returns a tuple with the InAppPurchaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Attributes) GetInAppPurchaseTypeOk() (*InAppPurchaseType, bool) {
	if o == nil || IsNil(o.InAppPurchaseType) {
		return nil, false
	}
	return o.InAppPurchaseType, true
}

// HasInAppPurchaseType returns a boolean if a field has been set.
func (o *InAppPurchaseV2Attributes) HasInAppPurchaseType() bool {
	if o != nil && !IsNil(o.InAppPurchaseType) {
		return true
	}

	return false
}

// SetInAppPurchaseType gets a reference to the given InAppPurchaseType and assigns it to the InAppPurchaseType field.
func (o *InAppPurchaseV2Attributes) SetInAppPurchaseType(v InAppPurchaseType) {
	o.InAppPurchaseType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *InAppPurchaseV2Attributes) GetState() InAppPurchaseState {
	if o == nil || IsNil(o.State) {
		var ret InAppPurchaseState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Attributes) GetStateOk() (*InAppPurchaseState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *InAppPurchaseV2Attributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given InAppPurchaseState and assigns it to the State field.
func (o *InAppPurchaseV2Attributes) SetState(v InAppPurchaseState) {
	o.State = &v
}

// GetReviewNote returns the ReviewNote field value if set, zero value otherwise.
func (o *InAppPurchaseV2Attributes) GetReviewNote() string {
	if o == nil || IsNil(o.ReviewNote) {
		var ret string
		return ret
	}
	return *o.ReviewNote
}

// GetReviewNoteOk returns a tuple with the ReviewNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Attributes) GetReviewNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewNote) {
		return nil, false
	}
	return o.ReviewNote, true
}

// HasReviewNote returns a boolean if a field has been set.
func (o *InAppPurchaseV2Attributes) HasReviewNote() bool {
	if o != nil && !IsNil(o.ReviewNote) {
		return true
	}

	return false
}

// SetReviewNote gets a reference to the given string and assigns it to the ReviewNote field.
func (o *InAppPurchaseV2Attributes) SetReviewNote(v string) {
	o.ReviewNote = &v
}

// GetFamilySharable returns the FamilySharable field value if set, zero value otherwise.
func (o *InAppPurchaseV2Attributes) GetFamilySharable() bool {
	if o == nil || IsNil(o.FamilySharable) {
		var ret bool
		return ret
	}
	return *o.FamilySharable
}

// GetFamilySharableOk returns a tuple with the FamilySharable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Attributes) GetFamilySharableOk() (*bool, bool) {
	if o == nil || IsNil(o.FamilySharable) {
		return nil, false
	}
	return o.FamilySharable, true
}

// HasFamilySharable returns a boolean if a field has been set.
func (o *InAppPurchaseV2Attributes) HasFamilySharable() bool {
	if o != nil && !IsNil(o.FamilySharable) {
		return true
	}

	return false
}

// SetFamilySharable gets a reference to the given bool and assigns it to the FamilySharable field.
func (o *InAppPurchaseV2Attributes) SetFamilySharable(v bool) {
	o.FamilySharable = &v
}

// GetContentHosting returns the ContentHosting field value if set, zero value otherwise.
func (o *InAppPurchaseV2Attributes) GetContentHosting() bool {
	if o == nil || IsNil(o.ContentHosting) {
		var ret bool
		return ret
	}
	return *o.ContentHosting
}

// GetContentHostingOk returns a tuple with the ContentHosting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Attributes) GetContentHostingOk() (*bool, bool) {
	if o == nil || IsNil(o.ContentHosting) {
		return nil, false
	}
	return o.ContentHosting, true
}

// HasContentHosting returns a boolean if a field has been set.
func (o *InAppPurchaseV2Attributes) HasContentHosting() bool {
	if o != nil && !IsNil(o.ContentHosting) {
		return true
	}

	return false
}

// SetContentHosting gets a reference to the given bool and assigns it to the ContentHosting field.
func (o *InAppPurchaseV2Attributes) SetContentHosting(v bool) {
	o.ContentHosting = &v
}

func (o InAppPurchaseV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2Attributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.InAppPurchaseType) {
		toSerialize["inAppPurchaseType"] = o.InAppPurchaseType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ReviewNote) {
		toSerialize["reviewNote"] = o.ReviewNote
	}
	if !IsNil(o.FamilySharable) {
		toSerialize["familySharable"] = o.FamilySharable
	}
	if !IsNil(o.ContentHosting) {
		toSerialize["contentHosting"] = o.ContentHosting
	}
	return toSerialize, nil
}

type NullableInAppPurchaseV2Attributes struct {
	value *InAppPurchaseV2Attributes
	isSet bool
}

func (v NullableInAppPurchaseV2Attributes) Get() *InAppPurchaseV2Attributes {
	return v.value
}

func (v *NullableInAppPurchaseV2Attributes) Set(val *InAppPurchaseV2Attributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2Attributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2Attributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2Attributes(val *InAppPurchaseV2Attributes) *NullableInAppPurchaseV2Attributes {
	return &NullableInAppPurchaseV2Attributes{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2Attributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2Attributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


