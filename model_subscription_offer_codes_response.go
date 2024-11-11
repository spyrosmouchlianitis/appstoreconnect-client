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

// checks if the SubscriptionOfferCodesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodesResponse{}

// SubscriptionOfferCodesResponse struct for SubscriptionOfferCodesResponse
type SubscriptionOfferCodesResponse struct {
	Data []SubscriptionOfferCode `json:"data"`
	Included []SubscriptionOfferCodesResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _SubscriptionOfferCodesResponse SubscriptionOfferCodesResponse

// NewSubscriptionOfferCodesResponse instantiates a new SubscriptionOfferCodesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodesResponse(data []SubscriptionOfferCode, links PagedDocumentLinks) *SubscriptionOfferCodesResponse {
	this := SubscriptionOfferCodesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionOfferCodesResponseWithDefaults instantiates a new SubscriptionOfferCodesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodesResponseWithDefaults() *SubscriptionOfferCodesResponse {
	this := SubscriptionOfferCodesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionOfferCodesResponse) GetData() []SubscriptionOfferCode {
	if o == nil {
		var ret []SubscriptionOfferCode
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodesResponse) GetDataOk() ([]SubscriptionOfferCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionOfferCodesResponse) SetData(v []SubscriptionOfferCode) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionOfferCodesResponse) GetIncluded() []SubscriptionOfferCodesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionOfferCodesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodesResponse) GetIncludedOk() ([]SubscriptionOfferCodesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionOfferCodesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionOfferCodesResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionOfferCodesResponse) SetIncluded(v []SubscriptionOfferCodesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionOfferCodesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionOfferCodesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionOfferCodesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionOfferCodesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionOfferCodesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SubscriptionOfferCodesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *SubscriptionOfferCodesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varSubscriptionOfferCodesResponse := _SubscriptionOfferCodesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionOfferCodesResponse)

	if err != nil {
		return err
	}

	*o = SubscriptionOfferCodesResponse(varSubscriptionOfferCodesResponse)

	return err
}

type NullableSubscriptionOfferCodesResponse struct {
	value *SubscriptionOfferCodesResponse
	isSet bool
}

func (v NullableSubscriptionOfferCodesResponse) Get() *SubscriptionOfferCodesResponse {
	return v.value
}

func (v *NullableSubscriptionOfferCodesResponse) Set(val *SubscriptionOfferCodesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodesResponse(val *SubscriptionOfferCodesResponse) *NullableSubscriptionOfferCodesResponse {
	return &NullableSubscriptionOfferCodesResponse{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


