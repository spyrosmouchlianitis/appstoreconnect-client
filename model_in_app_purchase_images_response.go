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

// checks if the InAppPurchaseImagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseImagesResponse{}

// InAppPurchaseImagesResponse struct for InAppPurchaseImagesResponse
type InAppPurchaseImagesResponse struct {
	Data []InAppPurchaseImage `json:"data"`
	Included []InAppPurchaseV2 `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _InAppPurchaseImagesResponse InAppPurchaseImagesResponse

// NewInAppPurchaseImagesResponse instantiates a new InAppPurchaseImagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseImagesResponse(data []InAppPurchaseImage, links PagedDocumentLinks) *InAppPurchaseImagesResponse {
	this := InAppPurchaseImagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewInAppPurchaseImagesResponseWithDefaults instantiates a new InAppPurchaseImagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseImagesResponseWithDefaults() *InAppPurchaseImagesResponse {
	this := InAppPurchaseImagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseImagesResponse) GetData() []InAppPurchaseImage {
	if o == nil {
		var ret []InAppPurchaseImage
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseImagesResponse) GetDataOk() ([]InAppPurchaseImage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseImagesResponse) SetData(v []InAppPurchaseImage) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *InAppPurchaseImagesResponse) GetIncluded() []InAppPurchaseV2 {
	if o == nil || IsNil(o.Included) {
		var ret []InAppPurchaseV2
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseImagesResponse) GetIncludedOk() ([]InAppPurchaseV2, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *InAppPurchaseImagesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []InAppPurchaseV2 and assigns it to the Included field.
func (o *InAppPurchaseImagesResponse) SetIncluded(v []InAppPurchaseV2) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *InAppPurchaseImagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseImagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InAppPurchaseImagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InAppPurchaseImagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseImagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InAppPurchaseImagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *InAppPurchaseImagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o InAppPurchaseImagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseImagesResponse) ToMap() (map[string]interface{}, error) {
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

func (o *InAppPurchaseImagesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varInAppPurchaseImagesResponse := _InAppPurchaseImagesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInAppPurchaseImagesResponse)

	if err != nil {
		return err
	}

	*o = InAppPurchaseImagesResponse(varInAppPurchaseImagesResponse)

	return err
}

type NullableInAppPurchaseImagesResponse struct {
	value *InAppPurchaseImagesResponse
	isSet bool
}

func (v NullableInAppPurchaseImagesResponse) Get() *InAppPurchaseImagesResponse {
	return v.value
}

func (v *NullableInAppPurchaseImagesResponse) Set(val *InAppPurchaseImagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseImagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseImagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseImagesResponse(val *InAppPurchaseImagesResponse) *NullableInAppPurchaseImagesResponse {
	return &NullableInAppPurchaseImagesResponse{value: val, isSet: true}
}

func (v NullableInAppPurchaseImagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseImagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


