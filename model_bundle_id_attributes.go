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

// checks if the BundleIdAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdAttributes{}

// BundleIdAttributes struct for BundleIdAttributes
type BundleIdAttributes struct {
	Name *string `json:"name,omitempty"`
	Platform *BundleIdPlatform `json:"platform,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	SeedId *string `json:"seedId,omitempty"`
}

// NewBundleIdAttributes instantiates a new BundleIdAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdAttributes() *BundleIdAttributes {
	this := BundleIdAttributes{}
	return &this
}

// NewBundleIdAttributesWithDefaults instantiates a new BundleIdAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdAttributesWithDefaults() *BundleIdAttributes {
	this := BundleIdAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BundleIdAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BundleIdAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BundleIdAttributes) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *BundleIdAttributes) GetPlatform() BundleIdPlatform {
	if o == nil || IsNil(o.Platform) {
		var ret BundleIdPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdAttributes) GetPlatformOk() (*BundleIdPlatform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *BundleIdAttributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given BundleIdPlatform and assigns it to the Platform field.
func (o *BundleIdAttributes) SetPlatform(v BundleIdPlatform) {
	o.Platform = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *BundleIdAttributes) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdAttributes) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *BundleIdAttributes) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *BundleIdAttributes) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetSeedId returns the SeedId field value if set, zero value otherwise.
func (o *BundleIdAttributes) GetSeedId() string {
	if o == nil || IsNil(o.SeedId) {
		var ret string
		return ret
	}
	return *o.SeedId
}

// GetSeedIdOk returns a tuple with the SeedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdAttributes) GetSeedIdOk() (*string, bool) {
	if o == nil || IsNil(o.SeedId) {
		return nil, false
	}
	return o.SeedId, true
}

// HasSeedId returns a boolean if a field has been set.
func (o *BundleIdAttributes) HasSeedId() bool {
	if o != nil && !IsNil(o.SeedId) {
		return true
	}

	return false
}

// SetSeedId gets a reference to the given string and assigns it to the SeedId field.
func (o *BundleIdAttributes) SetSeedId(v string) {
	o.SeedId = &v
}

func (o BundleIdAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.SeedId) {
		toSerialize["seedId"] = o.SeedId
	}
	return toSerialize, nil
}

type NullableBundleIdAttributes struct {
	value *BundleIdAttributes
	isSet bool
}

func (v NullableBundleIdAttributes) Get() *BundleIdAttributes {
	return v.value
}

func (v *NullableBundleIdAttributes) Set(val *BundleIdAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdAttributes(val *BundleIdAttributes) *NullableBundleIdAttributes {
	return &NullableBundleIdAttributes{value: val, isSet: true}
}

func (v NullableBundleIdAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


