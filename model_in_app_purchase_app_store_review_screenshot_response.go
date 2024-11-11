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

// checks if the InAppPurchaseAppStoreReviewScreenshotResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAppStoreReviewScreenshotResponse{}

// InAppPurchaseAppStoreReviewScreenshotResponse struct for InAppPurchaseAppStoreReviewScreenshotResponse
type InAppPurchaseAppStoreReviewScreenshotResponse struct {
	Data InAppPurchaseAppStoreReviewScreenshot `json:"data"`
	Included []InAppPurchaseV2 `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _InAppPurchaseAppStoreReviewScreenshotResponse InAppPurchaseAppStoreReviewScreenshotResponse

// NewInAppPurchaseAppStoreReviewScreenshotResponse instantiates a new InAppPurchaseAppStoreReviewScreenshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAppStoreReviewScreenshotResponse(data InAppPurchaseAppStoreReviewScreenshot, links DocumentLinks) *InAppPurchaseAppStoreReviewScreenshotResponse {
	this := InAppPurchaseAppStoreReviewScreenshotResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewInAppPurchaseAppStoreReviewScreenshotResponseWithDefaults instantiates a new InAppPurchaseAppStoreReviewScreenshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAppStoreReviewScreenshotResponseWithDefaults() *InAppPurchaseAppStoreReviewScreenshotResponse {
	this := InAppPurchaseAppStoreReviewScreenshotResponse{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) GetData() InAppPurchaseAppStoreReviewScreenshot {
	if o == nil {
		var ret InAppPurchaseAppStoreReviewScreenshot
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) GetDataOk() (*InAppPurchaseAppStoreReviewScreenshot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) SetData(v InAppPurchaseAppStoreReviewScreenshot) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) GetIncluded() []InAppPurchaseV2 {
	if o == nil || IsNil(o.Included) {
		var ret []InAppPurchaseV2
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) GetIncludedOk() ([]InAppPurchaseV2, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []InAppPurchaseV2 and assigns it to the Included field.
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) SetIncluded(v []InAppPurchaseV2) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InAppPurchaseAppStoreReviewScreenshotResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o InAppPurchaseAppStoreReviewScreenshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAppStoreReviewScreenshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *InAppPurchaseAppStoreReviewScreenshotResponse) UnmarshalJSON(data []byte) (err error) {
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

	varInAppPurchaseAppStoreReviewScreenshotResponse := _InAppPurchaseAppStoreReviewScreenshotResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInAppPurchaseAppStoreReviewScreenshotResponse)

	if err != nil {
		return err
	}

	*o = InAppPurchaseAppStoreReviewScreenshotResponse(varInAppPurchaseAppStoreReviewScreenshotResponse)

	return err
}

type NullableInAppPurchaseAppStoreReviewScreenshotResponse struct {
	value *InAppPurchaseAppStoreReviewScreenshotResponse
	isSet bool
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotResponse) Get() *InAppPurchaseAppStoreReviewScreenshotResponse {
	return v.value
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotResponse) Set(val *InAppPurchaseAppStoreReviewScreenshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAppStoreReviewScreenshotResponse(val *InAppPurchaseAppStoreReviewScreenshotResponse) *NullableInAppPurchaseAppStoreReviewScreenshotResponse {
	return &NullableInAppPurchaseAppStoreReviewScreenshotResponse{value: val, isSet: true}
}

func (v NullableInAppPurchaseAppStoreReviewScreenshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAppStoreReviewScreenshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


