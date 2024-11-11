/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"time"
)

// checks if the AlternativeDistributionPackageVersionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeDistributionPackageVersionAttributes{}

// AlternativeDistributionPackageVersionAttributes struct for AlternativeDistributionPackageVersionAttributes
type AlternativeDistributionPackageVersionAttributes struct {
	Url *string `json:"url,omitempty"`
	UrlExpirationDate *time.Time `json:"urlExpirationDate,omitempty"`
	Version *string `json:"version,omitempty"`
	FileChecksum *string `json:"fileChecksum,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewAlternativeDistributionPackageVersionAttributes instantiates a new AlternativeDistributionPackageVersionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeDistributionPackageVersionAttributes() *AlternativeDistributionPackageVersionAttributes {
	this := AlternativeDistributionPackageVersionAttributes{}
	return &this
}

// NewAlternativeDistributionPackageVersionAttributesWithDefaults instantiates a new AlternativeDistributionPackageVersionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeDistributionPackageVersionAttributesWithDefaults() *AlternativeDistributionPackageVersionAttributes {
	this := AlternativeDistributionPackageVersionAttributes{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AlternativeDistributionPackageVersionAttributes) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionPackageVersionAttributes) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AlternativeDistributionPackageVersionAttributes) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AlternativeDistributionPackageVersionAttributes) SetUrl(v string) {
	o.Url = &v
}

// GetUrlExpirationDate returns the UrlExpirationDate field value if set, zero value otherwise.
func (o *AlternativeDistributionPackageVersionAttributes) GetUrlExpirationDate() time.Time {
	if o == nil || IsNil(o.UrlExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.UrlExpirationDate
}

// GetUrlExpirationDateOk returns a tuple with the UrlExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionPackageVersionAttributes) GetUrlExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UrlExpirationDate) {
		return nil, false
	}
	return o.UrlExpirationDate, true
}

// HasUrlExpirationDate returns a boolean if a field has been set.
func (o *AlternativeDistributionPackageVersionAttributes) HasUrlExpirationDate() bool {
	if o != nil && !IsNil(o.UrlExpirationDate) {
		return true
	}

	return false
}

// SetUrlExpirationDate gets a reference to the given time.Time and assigns it to the UrlExpirationDate field.
func (o *AlternativeDistributionPackageVersionAttributes) SetUrlExpirationDate(v time.Time) {
	o.UrlExpirationDate = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AlternativeDistributionPackageVersionAttributes) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionPackageVersionAttributes) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AlternativeDistributionPackageVersionAttributes) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AlternativeDistributionPackageVersionAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetFileChecksum returns the FileChecksum field value if set, zero value otherwise.
func (o *AlternativeDistributionPackageVersionAttributes) GetFileChecksum() string {
	if o == nil || IsNil(o.FileChecksum) {
		var ret string
		return ret
	}
	return *o.FileChecksum
}

// GetFileChecksumOk returns a tuple with the FileChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionPackageVersionAttributes) GetFileChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.FileChecksum) {
		return nil, false
	}
	return o.FileChecksum, true
}

// HasFileChecksum returns a boolean if a field has been set.
func (o *AlternativeDistributionPackageVersionAttributes) HasFileChecksum() bool {
	if o != nil && !IsNil(o.FileChecksum) {
		return true
	}

	return false
}

// SetFileChecksum gets a reference to the given string and assigns it to the FileChecksum field.
func (o *AlternativeDistributionPackageVersionAttributes) SetFileChecksum(v string) {
	o.FileChecksum = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AlternativeDistributionPackageVersionAttributes) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionPackageVersionAttributes) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AlternativeDistributionPackageVersionAttributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AlternativeDistributionPackageVersionAttributes) SetState(v string) {
	o.State = &v
}

func (o AlternativeDistributionPackageVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeDistributionPackageVersionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UrlExpirationDate) {
		toSerialize["urlExpirationDate"] = o.UrlExpirationDate
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.FileChecksum) {
		toSerialize["fileChecksum"] = o.FileChecksum
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableAlternativeDistributionPackageVersionAttributes struct {
	value *AlternativeDistributionPackageVersionAttributes
	isSet bool
}

func (v NullableAlternativeDistributionPackageVersionAttributes) Get() *AlternativeDistributionPackageVersionAttributes {
	return v.value
}

func (v *NullableAlternativeDistributionPackageVersionAttributes) Set(val *AlternativeDistributionPackageVersionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeDistributionPackageVersionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeDistributionPackageVersionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeDistributionPackageVersionAttributes(val *AlternativeDistributionPackageVersionAttributes) *NullableAlternativeDistributionPackageVersionAttributes {
	return &NullableAlternativeDistributionPackageVersionAttributes{value: val, isSet: true}
}

func (v NullableAlternativeDistributionPackageVersionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeDistributionPackageVersionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


