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

// checks if the WinBackOfferPricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WinBackOfferPricesResponse{}

// WinBackOfferPricesResponse struct for WinBackOfferPricesResponse
type WinBackOfferPricesResponse struct {
	Data []WinBackOfferPrice `json:"data"`
	Included []SubscriptionOfferCodePricesResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _WinBackOfferPricesResponse WinBackOfferPricesResponse

// NewWinBackOfferPricesResponse instantiates a new WinBackOfferPricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWinBackOfferPricesResponse(data []WinBackOfferPrice, links PagedDocumentLinks) *WinBackOfferPricesResponse {
	this := WinBackOfferPricesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewWinBackOfferPricesResponseWithDefaults instantiates a new WinBackOfferPricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWinBackOfferPricesResponseWithDefaults() *WinBackOfferPricesResponse {
	this := WinBackOfferPricesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *WinBackOfferPricesResponse) GetData() []WinBackOfferPrice {
	if o == nil {
		var ret []WinBackOfferPrice
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferPricesResponse) GetDataOk() ([]WinBackOfferPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *WinBackOfferPricesResponse) SetData(v []WinBackOfferPrice) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WinBackOfferPricesResponse) GetIncluded() []SubscriptionOfferCodePricesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionOfferCodePricesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferPricesResponse) GetIncludedOk() ([]SubscriptionOfferCodePricesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WinBackOfferPricesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionOfferCodePricesResponseIncludedInner and assigns it to the Included field.
func (o *WinBackOfferPricesResponse) SetIncluded(v []SubscriptionOfferCodePricesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *WinBackOfferPricesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferPricesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *WinBackOfferPricesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WinBackOfferPricesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferPricesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WinBackOfferPricesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *WinBackOfferPricesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o WinBackOfferPricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WinBackOfferPricesResponse) ToMap() (map[string]interface{}, error) {
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

func (o *WinBackOfferPricesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varWinBackOfferPricesResponse := _WinBackOfferPricesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWinBackOfferPricesResponse)

	if err != nil {
		return err
	}

	*o = WinBackOfferPricesResponse(varWinBackOfferPricesResponse)

	return err
}

type NullableWinBackOfferPricesResponse struct {
	value *WinBackOfferPricesResponse
	isSet bool
}

func (v NullableWinBackOfferPricesResponse) Get() *WinBackOfferPricesResponse {
	return v.value
}

func (v *NullableWinBackOfferPricesResponse) Set(val *WinBackOfferPricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWinBackOfferPricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWinBackOfferPricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWinBackOfferPricesResponse(val *WinBackOfferPricesResponse) *NullableWinBackOfferPricesResponse {
	return &NullableWinBackOfferPricesResponse{value: val, isSet: true}
}

func (v NullableWinBackOfferPricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWinBackOfferPricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


