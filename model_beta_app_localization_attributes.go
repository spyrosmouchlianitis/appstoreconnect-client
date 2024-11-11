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

// checks if the BetaAppLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppLocalizationAttributes{}

// BetaAppLocalizationAttributes struct for BetaAppLocalizationAttributes
type BetaAppLocalizationAttributes struct {
	FeedbackEmail *string `json:"feedbackEmail,omitempty"`
	MarketingUrl *string `json:"marketingUrl,omitempty"`
	PrivacyPolicyUrl *string `json:"privacyPolicyUrl,omitempty"`
	TvOsPrivacyPolicy *string `json:"tvOsPrivacyPolicy,omitempty"`
	Description *string `json:"description,omitempty"`
	Locale *string `json:"locale,omitempty"`
}

// NewBetaAppLocalizationAttributes instantiates a new BetaAppLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppLocalizationAttributes() *BetaAppLocalizationAttributes {
	this := BetaAppLocalizationAttributes{}
	return &this
}

// NewBetaAppLocalizationAttributesWithDefaults instantiates a new BetaAppLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppLocalizationAttributesWithDefaults() *BetaAppLocalizationAttributes {
	this := BetaAppLocalizationAttributes{}
	return &this
}

// GetFeedbackEmail returns the FeedbackEmail field value if set, zero value otherwise.
func (o *BetaAppLocalizationAttributes) GetFeedbackEmail() string {
	if o == nil || IsNil(o.FeedbackEmail) {
		var ret string
		return ret
	}
	return *o.FeedbackEmail
}

// GetFeedbackEmailOk returns a tuple with the FeedbackEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationAttributes) GetFeedbackEmailOk() (*string, bool) {
	if o == nil || IsNil(o.FeedbackEmail) {
		return nil, false
	}
	return o.FeedbackEmail, true
}

// HasFeedbackEmail returns a boolean if a field has been set.
func (o *BetaAppLocalizationAttributes) HasFeedbackEmail() bool {
	if o != nil && !IsNil(o.FeedbackEmail) {
		return true
	}

	return false
}

// SetFeedbackEmail gets a reference to the given string and assigns it to the FeedbackEmail field.
func (o *BetaAppLocalizationAttributes) SetFeedbackEmail(v string) {
	o.FeedbackEmail = &v
}

// GetMarketingUrl returns the MarketingUrl field value if set, zero value otherwise.
func (o *BetaAppLocalizationAttributes) GetMarketingUrl() string {
	if o == nil || IsNil(o.MarketingUrl) {
		var ret string
		return ret
	}
	return *o.MarketingUrl
}

// GetMarketingUrlOk returns a tuple with the MarketingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationAttributes) GetMarketingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MarketingUrl) {
		return nil, false
	}
	return o.MarketingUrl, true
}

// HasMarketingUrl returns a boolean if a field has been set.
func (o *BetaAppLocalizationAttributes) HasMarketingUrl() bool {
	if o != nil && !IsNil(o.MarketingUrl) {
		return true
	}

	return false
}

// SetMarketingUrl gets a reference to the given string and assigns it to the MarketingUrl field.
func (o *BetaAppLocalizationAttributes) SetMarketingUrl(v string) {
	o.MarketingUrl = &v
}

// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value if set, zero value otherwise.
func (o *BetaAppLocalizationAttributes) GetPrivacyPolicyUrl() string {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyUrl
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationAttributes) GetPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		return nil, false
	}
	return o.PrivacyPolicyUrl, true
}

// HasPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *BetaAppLocalizationAttributes) HasPrivacyPolicyUrl() bool {
	if o != nil && !IsNil(o.PrivacyPolicyUrl) {
		return true
	}

	return false
}

// SetPrivacyPolicyUrl gets a reference to the given string and assigns it to the PrivacyPolicyUrl field.
func (o *BetaAppLocalizationAttributes) SetPrivacyPolicyUrl(v string) {
	o.PrivacyPolicyUrl = &v
}

// GetTvOsPrivacyPolicy returns the TvOsPrivacyPolicy field value if set, zero value otherwise.
func (o *BetaAppLocalizationAttributes) GetTvOsPrivacyPolicy() string {
	if o == nil || IsNil(o.TvOsPrivacyPolicy) {
		var ret string
		return ret
	}
	return *o.TvOsPrivacyPolicy
}

// GetTvOsPrivacyPolicyOk returns a tuple with the TvOsPrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationAttributes) GetTvOsPrivacyPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.TvOsPrivacyPolicy) {
		return nil, false
	}
	return o.TvOsPrivacyPolicy, true
}

// HasTvOsPrivacyPolicy returns a boolean if a field has been set.
func (o *BetaAppLocalizationAttributes) HasTvOsPrivacyPolicy() bool {
	if o != nil && !IsNil(o.TvOsPrivacyPolicy) {
		return true
	}

	return false
}

// SetTvOsPrivacyPolicy gets a reference to the given string and assigns it to the TvOsPrivacyPolicy field.
func (o *BetaAppLocalizationAttributes) SetTvOsPrivacyPolicy(v string) {
	o.TvOsPrivacyPolicy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BetaAppLocalizationAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BetaAppLocalizationAttributes) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BetaAppLocalizationAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *BetaAppLocalizationAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *BetaAppLocalizationAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *BetaAppLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

func (o BetaAppLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeedbackEmail) {
		toSerialize["feedbackEmail"] = o.FeedbackEmail
	}
	if !IsNil(o.MarketingUrl) {
		toSerialize["marketingUrl"] = o.MarketingUrl
	}
	if !IsNil(o.PrivacyPolicyUrl) {
		toSerialize["privacyPolicyUrl"] = o.PrivacyPolicyUrl
	}
	if !IsNil(o.TvOsPrivacyPolicy) {
		toSerialize["tvOsPrivacyPolicy"] = o.TvOsPrivacyPolicy
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableBetaAppLocalizationAttributes struct {
	value *BetaAppLocalizationAttributes
	isSet bool
}

func (v NullableBetaAppLocalizationAttributes) Get() *BetaAppLocalizationAttributes {
	return v.value
}

func (v *NullableBetaAppLocalizationAttributes) Set(val *BetaAppLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppLocalizationAttributes(val *BetaAppLocalizationAttributes) *NullableBetaAppLocalizationAttributes {
	return &NullableBetaAppLocalizationAttributes{value: val, isSet: true}
}

func (v NullableBetaAppLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


