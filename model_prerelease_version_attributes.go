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

// checks if the PrereleaseVersionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrereleaseVersionAttributes{}

// PrereleaseVersionAttributes struct for PrereleaseVersionAttributes
type PrereleaseVersionAttributes struct {
	Version *string `json:"version,omitempty"`
	Platform *Platform `json:"platform,omitempty"`
}

// NewPrereleaseVersionAttributes instantiates a new PrereleaseVersionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrereleaseVersionAttributes() *PrereleaseVersionAttributes {
	this := PrereleaseVersionAttributes{}
	return &this
}

// NewPrereleaseVersionAttributesWithDefaults instantiates a new PrereleaseVersionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrereleaseVersionAttributesWithDefaults() *PrereleaseVersionAttributes {
	this := PrereleaseVersionAttributes{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PrereleaseVersionAttributes) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrereleaseVersionAttributes) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PrereleaseVersionAttributes) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PrereleaseVersionAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *PrereleaseVersionAttributes) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrereleaseVersionAttributes) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *PrereleaseVersionAttributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *PrereleaseVersionAttributes) SetPlatform(v Platform) {
	o.Platform = &v
}

func (o PrereleaseVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrereleaseVersionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	return toSerialize, nil
}

type NullablePrereleaseVersionAttributes struct {
	value *PrereleaseVersionAttributes
	isSet bool
}

func (v NullablePrereleaseVersionAttributes) Get() *PrereleaseVersionAttributes {
	return v.value
}

func (v *NullablePrereleaseVersionAttributes) Set(val *PrereleaseVersionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePrereleaseVersionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePrereleaseVersionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrereleaseVersionAttributes(val *PrereleaseVersionAttributes) *NullablePrereleaseVersionAttributes {
	return &NullablePrereleaseVersionAttributes{value: val, isSet: true}
}

func (v NullablePrereleaseVersionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrereleaseVersionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


