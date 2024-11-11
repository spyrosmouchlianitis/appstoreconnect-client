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

// checks if the AppAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAttributes{}

// AppAttributes struct for AppAttributes
type AppAttributes struct {
	Name *string `json:"name,omitempty"`
	BundleId *string `json:"bundleId,omitempty"`
	Sku *string `json:"sku,omitempty"`
	PrimaryLocale *string `json:"primaryLocale,omitempty"`
	IsOrEverWasMadeForKids *bool `json:"isOrEverWasMadeForKids,omitempty"`
	SubscriptionStatusUrl *string `json:"subscriptionStatusUrl,omitempty"`
	SubscriptionStatusUrlVersion *SubscriptionStatusUrlVersion `json:"subscriptionStatusUrlVersion,omitempty"`
	SubscriptionStatusUrlForSandbox *string `json:"subscriptionStatusUrlForSandbox,omitempty"`
	SubscriptionStatusUrlVersionForSandbox *SubscriptionStatusUrlVersion `json:"subscriptionStatusUrlVersionForSandbox,omitempty"`
	ContentRightsDeclaration *string `json:"contentRightsDeclaration,omitempty"`
	StreamlinedPurchasingEnabled *bool `json:"streamlinedPurchasingEnabled,omitempty"`
}

// NewAppAttributes instantiates a new AppAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAttributes() *AppAttributes {
	this := AppAttributes{}
	return &this
}

// NewAppAttributesWithDefaults instantiates a new AppAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAttributesWithDefaults() *AppAttributes {
	this := AppAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppAttributes) SetName(v string) {
	o.Name = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AppAttributes) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AppAttributes) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AppAttributes) SetBundleId(v string) {
	o.BundleId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *AppAttributes) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *AppAttributes) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *AppAttributes) SetSku(v string) {
	o.Sku = &v
}

// GetPrimaryLocale returns the PrimaryLocale field value if set, zero value otherwise.
func (o *AppAttributes) GetPrimaryLocale() string {
	if o == nil || IsNil(o.PrimaryLocale) {
		var ret string
		return ret
	}
	return *o.PrimaryLocale
}

// GetPrimaryLocaleOk returns a tuple with the PrimaryLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetPrimaryLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryLocale) {
		return nil, false
	}
	return o.PrimaryLocale, true
}

// HasPrimaryLocale returns a boolean if a field has been set.
func (o *AppAttributes) HasPrimaryLocale() bool {
	if o != nil && !IsNil(o.PrimaryLocale) {
		return true
	}

	return false
}

// SetPrimaryLocale gets a reference to the given string and assigns it to the PrimaryLocale field.
func (o *AppAttributes) SetPrimaryLocale(v string) {
	o.PrimaryLocale = &v
}

// GetIsOrEverWasMadeForKids returns the IsOrEverWasMadeForKids field value if set, zero value otherwise.
func (o *AppAttributes) GetIsOrEverWasMadeForKids() bool {
	if o == nil || IsNil(o.IsOrEverWasMadeForKids) {
		var ret bool
		return ret
	}
	return *o.IsOrEverWasMadeForKids
}

// GetIsOrEverWasMadeForKidsOk returns a tuple with the IsOrEverWasMadeForKids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetIsOrEverWasMadeForKidsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrEverWasMadeForKids) {
		return nil, false
	}
	return o.IsOrEverWasMadeForKids, true
}

// HasIsOrEverWasMadeForKids returns a boolean if a field has been set.
func (o *AppAttributes) HasIsOrEverWasMadeForKids() bool {
	if o != nil && !IsNil(o.IsOrEverWasMadeForKids) {
		return true
	}

	return false
}

// SetIsOrEverWasMadeForKids gets a reference to the given bool and assigns it to the IsOrEverWasMadeForKids field.
func (o *AppAttributes) SetIsOrEverWasMadeForKids(v bool) {
	o.IsOrEverWasMadeForKids = &v
}

// GetSubscriptionStatusUrl returns the SubscriptionStatusUrl field value if set, zero value otherwise.
func (o *AppAttributes) GetSubscriptionStatusUrl() string {
	if o == nil || IsNil(o.SubscriptionStatusUrl) {
		var ret string
		return ret
	}
	return *o.SubscriptionStatusUrl
}

// GetSubscriptionStatusUrlOk returns a tuple with the SubscriptionStatusUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetSubscriptionStatusUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionStatusUrl) {
		return nil, false
	}
	return o.SubscriptionStatusUrl, true
}

// HasSubscriptionStatusUrl returns a boolean if a field has been set.
func (o *AppAttributes) HasSubscriptionStatusUrl() bool {
	if o != nil && !IsNil(o.SubscriptionStatusUrl) {
		return true
	}

	return false
}

// SetSubscriptionStatusUrl gets a reference to the given string and assigns it to the SubscriptionStatusUrl field.
func (o *AppAttributes) SetSubscriptionStatusUrl(v string) {
	o.SubscriptionStatusUrl = &v
}

// GetSubscriptionStatusUrlVersion returns the SubscriptionStatusUrlVersion field value if set, zero value otherwise.
func (o *AppAttributes) GetSubscriptionStatusUrlVersion() SubscriptionStatusUrlVersion {
	if o == nil || IsNil(o.SubscriptionStatusUrlVersion) {
		var ret SubscriptionStatusUrlVersion
		return ret
	}
	return *o.SubscriptionStatusUrlVersion
}

// GetSubscriptionStatusUrlVersionOk returns a tuple with the SubscriptionStatusUrlVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetSubscriptionStatusUrlVersionOk() (*SubscriptionStatusUrlVersion, bool) {
	if o == nil || IsNil(o.SubscriptionStatusUrlVersion) {
		return nil, false
	}
	return o.SubscriptionStatusUrlVersion, true
}

// HasSubscriptionStatusUrlVersion returns a boolean if a field has been set.
func (o *AppAttributes) HasSubscriptionStatusUrlVersion() bool {
	if o != nil && !IsNil(o.SubscriptionStatusUrlVersion) {
		return true
	}

	return false
}

// SetSubscriptionStatusUrlVersion gets a reference to the given SubscriptionStatusUrlVersion and assigns it to the SubscriptionStatusUrlVersion field.
func (o *AppAttributes) SetSubscriptionStatusUrlVersion(v SubscriptionStatusUrlVersion) {
	o.SubscriptionStatusUrlVersion = &v
}

// GetSubscriptionStatusUrlForSandbox returns the SubscriptionStatusUrlForSandbox field value if set, zero value otherwise.
func (o *AppAttributes) GetSubscriptionStatusUrlForSandbox() string {
	if o == nil || IsNil(o.SubscriptionStatusUrlForSandbox) {
		var ret string
		return ret
	}
	return *o.SubscriptionStatusUrlForSandbox
}

// GetSubscriptionStatusUrlForSandboxOk returns a tuple with the SubscriptionStatusUrlForSandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetSubscriptionStatusUrlForSandboxOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionStatusUrlForSandbox) {
		return nil, false
	}
	return o.SubscriptionStatusUrlForSandbox, true
}

// HasSubscriptionStatusUrlForSandbox returns a boolean if a field has been set.
func (o *AppAttributes) HasSubscriptionStatusUrlForSandbox() bool {
	if o != nil && !IsNil(o.SubscriptionStatusUrlForSandbox) {
		return true
	}

	return false
}

// SetSubscriptionStatusUrlForSandbox gets a reference to the given string and assigns it to the SubscriptionStatusUrlForSandbox field.
func (o *AppAttributes) SetSubscriptionStatusUrlForSandbox(v string) {
	o.SubscriptionStatusUrlForSandbox = &v
}

// GetSubscriptionStatusUrlVersionForSandbox returns the SubscriptionStatusUrlVersionForSandbox field value if set, zero value otherwise.
func (o *AppAttributes) GetSubscriptionStatusUrlVersionForSandbox() SubscriptionStatusUrlVersion {
	if o == nil || IsNil(o.SubscriptionStatusUrlVersionForSandbox) {
		var ret SubscriptionStatusUrlVersion
		return ret
	}
	return *o.SubscriptionStatusUrlVersionForSandbox
}

// GetSubscriptionStatusUrlVersionForSandboxOk returns a tuple with the SubscriptionStatusUrlVersionForSandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetSubscriptionStatusUrlVersionForSandboxOk() (*SubscriptionStatusUrlVersion, bool) {
	if o == nil || IsNil(o.SubscriptionStatusUrlVersionForSandbox) {
		return nil, false
	}
	return o.SubscriptionStatusUrlVersionForSandbox, true
}

// HasSubscriptionStatusUrlVersionForSandbox returns a boolean if a field has been set.
func (o *AppAttributes) HasSubscriptionStatusUrlVersionForSandbox() bool {
	if o != nil && !IsNil(o.SubscriptionStatusUrlVersionForSandbox) {
		return true
	}

	return false
}

// SetSubscriptionStatusUrlVersionForSandbox gets a reference to the given SubscriptionStatusUrlVersion and assigns it to the SubscriptionStatusUrlVersionForSandbox field.
func (o *AppAttributes) SetSubscriptionStatusUrlVersionForSandbox(v SubscriptionStatusUrlVersion) {
	o.SubscriptionStatusUrlVersionForSandbox = &v
}

// GetContentRightsDeclaration returns the ContentRightsDeclaration field value if set, zero value otherwise.
func (o *AppAttributes) GetContentRightsDeclaration() string {
	if o == nil || IsNil(o.ContentRightsDeclaration) {
		var ret string
		return ret
	}
	return *o.ContentRightsDeclaration
}

// GetContentRightsDeclarationOk returns a tuple with the ContentRightsDeclaration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetContentRightsDeclarationOk() (*string, bool) {
	if o == nil || IsNil(o.ContentRightsDeclaration) {
		return nil, false
	}
	return o.ContentRightsDeclaration, true
}

// HasContentRightsDeclaration returns a boolean if a field has been set.
func (o *AppAttributes) HasContentRightsDeclaration() bool {
	if o != nil && !IsNil(o.ContentRightsDeclaration) {
		return true
	}

	return false
}

// SetContentRightsDeclaration gets a reference to the given string and assigns it to the ContentRightsDeclaration field.
func (o *AppAttributes) SetContentRightsDeclaration(v string) {
	o.ContentRightsDeclaration = &v
}

// GetStreamlinedPurchasingEnabled returns the StreamlinedPurchasingEnabled field value if set, zero value otherwise.
func (o *AppAttributes) GetStreamlinedPurchasingEnabled() bool {
	if o == nil || IsNil(o.StreamlinedPurchasingEnabled) {
		var ret bool
		return ret
	}
	return *o.StreamlinedPurchasingEnabled
}

// GetStreamlinedPurchasingEnabledOk returns a tuple with the StreamlinedPurchasingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetStreamlinedPurchasingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StreamlinedPurchasingEnabled) {
		return nil, false
	}
	return o.StreamlinedPurchasingEnabled, true
}

// HasStreamlinedPurchasingEnabled returns a boolean if a field has been set.
func (o *AppAttributes) HasStreamlinedPurchasingEnabled() bool {
	if o != nil && !IsNil(o.StreamlinedPurchasingEnabled) {
		return true
	}

	return false
}

// SetStreamlinedPurchasingEnabled gets a reference to the given bool and assigns it to the StreamlinedPurchasingEnabled field.
func (o *AppAttributes) SetStreamlinedPurchasingEnabled(v bool) {
	o.StreamlinedPurchasingEnabled = &v
}

func (o AppAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundleId"] = o.BundleId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.PrimaryLocale) {
		toSerialize["primaryLocale"] = o.PrimaryLocale
	}
	if !IsNil(o.IsOrEverWasMadeForKids) {
		toSerialize["isOrEverWasMadeForKids"] = o.IsOrEverWasMadeForKids
	}
	if !IsNil(o.SubscriptionStatusUrl) {
		toSerialize["subscriptionStatusUrl"] = o.SubscriptionStatusUrl
	}
	if !IsNil(o.SubscriptionStatusUrlVersion) {
		toSerialize["subscriptionStatusUrlVersion"] = o.SubscriptionStatusUrlVersion
	}
	if !IsNil(o.SubscriptionStatusUrlForSandbox) {
		toSerialize["subscriptionStatusUrlForSandbox"] = o.SubscriptionStatusUrlForSandbox
	}
	if !IsNil(o.SubscriptionStatusUrlVersionForSandbox) {
		toSerialize["subscriptionStatusUrlVersionForSandbox"] = o.SubscriptionStatusUrlVersionForSandbox
	}
	if !IsNil(o.ContentRightsDeclaration) {
		toSerialize["contentRightsDeclaration"] = o.ContentRightsDeclaration
	}
	if !IsNil(o.StreamlinedPurchasingEnabled) {
		toSerialize["streamlinedPurchasingEnabled"] = o.StreamlinedPurchasingEnabled
	}
	return toSerialize, nil
}

type NullableAppAttributes struct {
	value *AppAttributes
	isSet bool
}

func (v NullableAppAttributes) Get() *AppAttributes {
	return v.value
}

func (v *NullableAppAttributes) Set(val *AppAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAttributes(val *AppAttributes) *NullableAppAttributes {
	return &NullableAppAttributes{value: val, isSet: true}
}

func (v NullableAppAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


