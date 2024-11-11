/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the InAppPurchaseV2Relationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2Relationships{}

// InAppPurchaseV2Relationships struct for InAppPurchaseV2Relationships
type InAppPurchaseV2Relationships struct {
	InAppPurchaseLocalizations *InAppPurchaseV2RelationshipsInAppPurchaseLocalizations `json:"inAppPurchaseLocalizations,omitempty"`
	PricePoints *InAppPurchaseV2RelationshipsPricePoints `json:"pricePoints,omitempty"`
	Content *InAppPurchaseV2RelationshipsContent `json:"content,omitempty"`
	AppStoreReviewScreenshot *InAppPurchaseV2RelationshipsAppStoreReviewScreenshot `json:"appStoreReviewScreenshot,omitempty"`
	PromotedPurchase *InAppPurchaseV2RelationshipsPromotedPurchase `json:"promotedPurchase,omitempty"`
	IapPriceSchedule *InAppPurchaseV2RelationshipsIapPriceSchedule `json:"iapPriceSchedule,omitempty"`
	InAppPurchaseAvailability *InAppPurchaseV2RelationshipsInAppPurchaseAvailability `json:"inAppPurchaseAvailability,omitempty"`
	Images *InAppPurchaseV2RelationshipsImages `json:"images,omitempty"`
}

// NewInAppPurchaseV2Relationships instantiates a new InAppPurchaseV2Relationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2Relationships() *InAppPurchaseV2Relationships {
	this := InAppPurchaseV2Relationships{}
	return &this
}

// NewInAppPurchaseV2RelationshipsWithDefaults instantiates a new InAppPurchaseV2Relationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2RelationshipsWithDefaults() *InAppPurchaseV2Relationships {
	this := InAppPurchaseV2Relationships{}
	return &this
}

// GetInAppPurchaseLocalizations returns the InAppPurchaseLocalizations field value if set, zero value otherwise.
func (o *InAppPurchaseV2Relationships) GetInAppPurchaseLocalizations() InAppPurchaseV2RelationshipsInAppPurchaseLocalizations {
	if o == nil || IsNil(o.InAppPurchaseLocalizations) {
		var ret InAppPurchaseV2RelationshipsInAppPurchaseLocalizations
		return ret
	}
	return *o.InAppPurchaseLocalizations
}

// GetInAppPurchaseLocalizationsOk returns a tuple with the InAppPurchaseLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Relationships) GetInAppPurchaseLocalizationsOk() (*InAppPurchaseV2RelationshipsInAppPurchaseLocalizations, bool) {
	if o == nil || IsNil(o.InAppPurchaseLocalizations) {
		return nil, false
	}
	return o.InAppPurchaseLocalizations, true
}

// HasInAppPurchaseLocalizations returns a boolean if a field has been set.
func (o *InAppPurchaseV2Relationships) HasInAppPurchaseLocalizations() bool {
	if o != nil && !IsNil(o.InAppPurchaseLocalizations) {
		return true
	}

	return false
}

// SetInAppPurchaseLocalizations gets a reference to the given InAppPurchaseV2RelationshipsInAppPurchaseLocalizations and assigns it to the InAppPurchaseLocalizations field.
func (o *InAppPurchaseV2Relationships) SetInAppPurchaseLocalizations(v InAppPurchaseV2RelationshipsInAppPurchaseLocalizations) {
	o.InAppPurchaseLocalizations = &v
}

// GetPricePoints returns the PricePoints field value if set, zero value otherwise.
func (o *InAppPurchaseV2Relationships) GetPricePoints() InAppPurchaseV2RelationshipsPricePoints {
	if o == nil || IsNil(o.PricePoints) {
		var ret InAppPurchaseV2RelationshipsPricePoints
		return ret
	}
	return *o.PricePoints
}

// GetPricePointsOk returns a tuple with the PricePoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Relationships) GetPricePointsOk() (*InAppPurchaseV2RelationshipsPricePoints, bool) {
	if o == nil || IsNil(o.PricePoints) {
		return nil, false
	}
	return o.PricePoints, true
}

// HasPricePoints returns a boolean if a field has been set.
func (o *InAppPurchaseV2Relationships) HasPricePoints() bool {
	if o != nil && !IsNil(o.PricePoints) {
		return true
	}

	return false
}

// SetPricePoints gets a reference to the given InAppPurchaseV2RelationshipsPricePoints and assigns it to the PricePoints field.
func (o *InAppPurchaseV2Relationships) SetPricePoints(v InAppPurchaseV2RelationshipsPricePoints) {
	o.PricePoints = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *InAppPurchaseV2Relationships) GetContent() InAppPurchaseV2RelationshipsContent {
	if o == nil || IsNil(o.Content) {
		var ret InAppPurchaseV2RelationshipsContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Relationships) GetContentOk() (*InAppPurchaseV2RelationshipsContent, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *InAppPurchaseV2Relationships) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given InAppPurchaseV2RelationshipsContent and assigns it to the Content field.
func (o *InAppPurchaseV2Relationships) SetContent(v InAppPurchaseV2RelationshipsContent) {
	o.Content = &v
}

// GetAppStoreReviewScreenshot returns the AppStoreReviewScreenshot field value if set, zero value otherwise.
func (o *InAppPurchaseV2Relationships) GetAppStoreReviewScreenshot() InAppPurchaseV2RelationshipsAppStoreReviewScreenshot {
	if o == nil || IsNil(o.AppStoreReviewScreenshot) {
		var ret InAppPurchaseV2RelationshipsAppStoreReviewScreenshot
		return ret
	}
	return *o.AppStoreReviewScreenshot
}

// GetAppStoreReviewScreenshotOk returns a tuple with the AppStoreReviewScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Relationships) GetAppStoreReviewScreenshotOk() (*InAppPurchaseV2RelationshipsAppStoreReviewScreenshot, bool) {
	if o == nil || IsNil(o.AppStoreReviewScreenshot) {
		return nil, false
	}
	return o.AppStoreReviewScreenshot, true
}

// HasAppStoreReviewScreenshot returns a boolean if a field has been set.
func (o *InAppPurchaseV2Relationships) HasAppStoreReviewScreenshot() bool {
	if o != nil && !IsNil(o.AppStoreReviewScreenshot) {
		return true
	}

	return false
}

// SetAppStoreReviewScreenshot gets a reference to the given InAppPurchaseV2RelationshipsAppStoreReviewScreenshot and assigns it to the AppStoreReviewScreenshot field.
func (o *InAppPurchaseV2Relationships) SetAppStoreReviewScreenshot(v InAppPurchaseV2RelationshipsAppStoreReviewScreenshot) {
	o.AppStoreReviewScreenshot = &v
}

// GetPromotedPurchase returns the PromotedPurchase field value if set, zero value otherwise.
func (o *InAppPurchaseV2Relationships) GetPromotedPurchase() InAppPurchaseV2RelationshipsPromotedPurchase {
	if o == nil || IsNil(o.PromotedPurchase) {
		var ret InAppPurchaseV2RelationshipsPromotedPurchase
		return ret
	}
	return *o.PromotedPurchase
}

// GetPromotedPurchaseOk returns a tuple with the PromotedPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Relationships) GetPromotedPurchaseOk() (*InAppPurchaseV2RelationshipsPromotedPurchase, bool) {
	if o == nil || IsNil(o.PromotedPurchase) {
		return nil, false
	}
	return o.PromotedPurchase, true
}

// HasPromotedPurchase returns a boolean if a field has been set.
func (o *InAppPurchaseV2Relationships) HasPromotedPurchase() bool {
	if o != nil && !IsNil(o.PromotedPurchase) {
		return true
	}

	return false
}

// SetPromotedPurchase gets a reference to the given InAppPurchaseV2RelationshipsPromotedPurchase and assigns it to the PromotedPurchase field.
func (o *InAppPurchaseV2Relationships) SetPromotedPurchase(v InAppPurchaseV2RelationshipsPromotedPurchase) {
	o.PromotedPurchase = &v
}

// GetIapPriceSchedule returns the IapPriceSchedule field value if set, zero value otherwise.
func (o *InAppPurchaseV2Relationships) GetIapPriceSchedule() InAppPurchaseV2RelationshipsIapPriceSchedule {
	if o == nil || IsNil(o.IapPriceSchedule) {
		var ret InAppPurchaseV2RelationshipsIapPriceSchedule
		return ret
	}
	return *o.IapPriceSchedule
}

// GetIapPriceScheduleOk returns a tuple with the IapPriceSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Relationships) GetIapPriceScheduleOk() (*InAppPurchaseV2RelationshipsIapPriceSchedule, bool) {
	if o == nil || IsNil(o.IapPriceSchedule) {
		return nil, false
	}
	return o.IapPriceSchedule, true
}

// HasIapPriceSchedule returns a boolean if a field has been set.
func (o *InAppPurchaseV2Relationships) HasIapPriceSchedule() bool {
	if o != nil && !IsNil(o.IapPriceSchedule) {
		return true
	}

	return false
}

// SetIapPriceSchedule gets a reference to the given InAppPurchaseV2RelationshipsIapPriceSchedule and assigns it to the IapPriceSchedule field.
func (o *InAppPurchaseV2Relationships) SetIapPriceSchedule(v InAppPurchaseV2RelationshipsIapPriceSchedule) {
	o.IapPriceSchedule = &v
}

// GetInAppPurchaseAvailability returns the InAppPurchaseAvailability field value if set, zero value otherwise.
func (o *InAppPurchaseV2Relationships) GetInAppPurchaseAvailability() InAppPurchaseV2RelationshipsInAppPurchaseAvailability {
	if o == nil || IsNil(o.InAppPurchaseAvailability) {
		var ret InAppPurchaseV2RelationshipsInAppPurchaseAvailability
		return ret
	}
	return *o.InAppPurchaseAvailability
}

// GetInAppPurchaseAvailabilityOk returns a tuple with the InAppPurchaseAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Relationships) GetInAppPurchaseAvailabilityOk() (*InAppPurchaseV2RelationshipsInAppPurchaseAvailability, bool) {
	if o == nil || IsNil(o.InAppPurchaseAvailability) {
		return nil, false
	}
	return o.InAppPurchaseAvailability, true
}

// HasInAppPurchaseAvailability returns a boolean if a field has been set.
func (o *InAppPurchaseV2Relationships) HasInAppPurchaseAvailability() bool {
	if o != nil && !IsNil(o.InAppPurchaseAvailability) {
		return true
	}

	return false
}

// SetInAppPurchaseAvailability gets a reference to the given InAppPurchaseV2RelationshipsInAppPurchaseAvailability and assigns it to the InAppPurchaseAvailability field.
func (o *InAppPurchaseV2Relationships) SetInAppPurchaseAvailability(v InAppPurchaseV2RelationshipsInAppPurchaseAvailability) {
	o.InAppPurchaseAvailability = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *InAppPurchaseV2Relationships) GetImages() InAppPurchaseV2RelationshipsImages {
	if o == nil || IsNil(o.Images) {
		var ret InAppPurchaseV2RelationshipsImages
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2Relationships) GetImagesOk() (*InAppPurchaseV2RelationshipsImages, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *InAppPurchaseV2Relationships) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given InAppPurchaseV2RelationshipsImages and assigns it to the Images field.
func (o *InAppPurchaseV2Relationships) SetImages(v InAppPurchaseV2RelationshipsImages) {
	o.Images = &v
}

func (o InAppPurchaseV2Relationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2Relationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InAppPurchaseLocalizations) {
		toSerialize["inAppPurchaseLocalizations"] = o.InAppPurchaseLocalizations
	}
	if !IsNil(o.PricePoints) {
		toSerialize["pricePoints"] = o.PricePoints
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.AppStoreReviewScreenshot) {
		toSerialize["appStoreReviewScreenshot"] = o.AppStoreReviewScreenshot
	}
	if !IsNil(o.PromotedPurchase) {
		toSerialize["promotedPurchase"] = o.PromotedPurchase
	}
	if !IsNil(o.IapPriceSchedule) {
		toSerialize["iapPriceSchedule"] = o.IapPriceSchedule
	}
	if !IsNil(o.InAppPurchaseAvailability) {
		toSerialize["inAppPurchaseAvailability"] = o.InAppPurchaseAvailability
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	return toSerialize, nil
}

type NullableInAppPurchaseV2Relationships struct {
	value *InAppPurchaseV2Relationships
	isSet bool
}

func (v NullableInAppPurchaseV2Relationships) Get() *InAppPurchaseV2Relationships {
	return v.value
}

func (v *NullableInAppPurchaseV2Relationships) Set(val *InAppPurchaseV2Relationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2Relationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2Relationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2Relationships(val *InAppPurchaseV2Relationships) *NullableInAppPurchaseV2Relationships {
	return &NullableInAppPurchaseV2Relationships{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2Relationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2Relationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


