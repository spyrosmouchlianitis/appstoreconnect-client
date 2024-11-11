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

// checks if the AppStoreReviewDetailAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewDetailAttributes{}

// AppStoreReviewDetailAttributes struct for AppStoreReviewDetailAttributes
type AppStoreReviewDetailAttributes struct {
	ContactFirstName *string `json:"contactFirstName,omitempty"`
	ContactLastName *string `json:"contactLastName,omitempty"`
	ContactPhone *string `json:"contactPhone,omitempty"`
	ContactEmail *string `json:"contactEmail,omitempty"`
	DemoAccountName *string `json:"demoAccountName,omitempty"`
	DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
	DemoAccountRequired *bool `json:"demoAccountRequired,omitempty"`
	Notes *string `json:"notes,omitempty"`
}

// NewAppStoreReviewDetailAttributes instantiates a new AppStoreReviewDetailAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailAttributes() *AppStoreReviewDetailAttributes {
	this := AppStoreReviewDetailAttributes{}
	return &this
}

// NewAppStoreReviewDetailAttributesWithDefaults instantiates a new AppStoreReviewDetailAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailAttributesWithDefaults() *AppStoreReviewDetailAttributes {
	this := AppStoreReviewDetailAttributes{}
	return &this
}

// GetContactFirstName returns the ContactFirstName field value if set, zero value otherwise.
func (o *AppStoreReviewDetailAttributes) GetContactFirstName() string {
	if o == nil || IsNil(o.ContactFirstName) {
		var ret string
		return ret
	}
	return *o.ContactFirstName
}

// GetContactFirstNameOk returns a tuple with the ContactFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailAttributes) GetContactFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactFirstName) {
		return nil, false
	}
	return o.ContactFirstName, true
}

// HasContactFirstName returns a boolean if a field has been set.
func (o *AppStoreReviewDetailAttributes) HasContactFirstName() bool {
	if o != nil && !IsNil(o.ContactFirstName) {
		return true
	}

	return false
}

// SetContactFirstName gets a reference to the given string and assigns it to the ContactFirstName field.
func (o *AppStoreReviewDetailAttributes) SetContactFirstName(v string) {
	o.ContactFirstName = &v
}

// GetContactLastName returns the ContactLastName field value if set, zero value otherwise.
func (o *AppStoreReviewDetailAttributes) GetContactLastName() string {
	if o == nil || IsNil(o.ContactLastName) {
		var ret string
		return ret
	}
	return *o.ContactLastName
}

// GetContactLastNameOk returns a tuple with the ContactLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailAttributes) GetContactLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactLastName) {
		return nil, false
	}
	return o.ContactLastName, true
}

// HasContactLastName returns a boolean if a field has been set.
func (o *AppStoreReviewDetailAttributes) HasContactLastName() bool {
	if o != nil && !IsNil(o.ContactLastName) {
		return true
	}

	return false
}

// SetContactLastName gets a reference to the given string and assigns it to the ContactLastName field.
func (o *AppStoreReviewDetailAttributes) SetContactLastName(v string) {
	o.ContactLastName = &v
}

// GetContactPhone returns the ContactPhone field value if set, zero value otherwise.
func (o *AppStoreReviewDetailAttributes) GetContactPhone() string {
	if o == nil || IsNil(o.ContactPhone) {
		var ret string
		return ret
	}
	return *o.ContactPhone
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailAttributes) GetContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ContactPhone) {
		return nil, false
	}
	return o.ContactPhone, true
}

// HasContactPhone returns a boolean if a field has been set.
func (o *AppStoreReviewDetailAttributes) HasContactPhone() bool {
	if o != nil && !IsNil(o.ContactPhone) {
		return true
	}

	return false
}

// SetContactPhone gets a reference to the given string and assigns it to the ContactPhone field.
func (o *AppStoreReviewDetailAttributes) SetContactPhone(v string) {
	o.ContactPhone = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *AppStoreReviewDetailAttributes) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailAttributes) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *AppStoreReviewDetailAttributes) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *AppStoreReviewDetailAttributes) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetDemoAccountName returns the DemoAccountName field value if set, zero value otherwise.
func (o *AppStoreReviewDetailAttributes) GetDemoAccountName() string {
	if o == nil || IsNil(o.DemoAccountName) {
		var ret string
		return ret
	}
	return *o.DemoAccountName
}

// GetDemoAccountNameOk returns a tuple with the DemoAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailAttributes) GetDemoAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.DemoAccountName) {
		return nil, false
	}
	return o.DemoAccountName, true
}

// HasDemoAccountName returns a boolean if a field has been set.
func (o *AppStoreReviewDetailAttributes) HasDemoAccountName() bool {
	if o != nil && !IsNil(o.DemoAccountName) {
		return true
	}

	return false
}

// SetDemoAccountName gets a reference to the given string and assigns it to the DemoAccountName field.
func (o *AppStoreReviewDetailAttributes) SetDemoAccountName(v string) {
	o.DemoAccountName = &v
}

// GetDemoAccountPassword returns the DemoAccountPassword field value if set, zero value otherwise.
func (o *AppStoreReviewDetailAttributes) GetDemoAccountPassword() string {
	if o == nil || IsNil(o.DemoAccountPassword) {
		var ret string
		return ret
	}
	return *o.DemoAccountPassword
}

// GetDemoAccountPasswordOk returns a tuple with the DemoAccountPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailAttributes) GetDemoAccountPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DemoAccountPassword) {
		return nil, false
	}
	return o.DemoAccountPassword, true
}

// HasDemoAccountPassword returns a boolean if a field has been set.
func (o *AppStoreReviewDetailAttributes) HasDemoAccountPassword() bool {
	if o != nil && !IsNil(o.DemoAccountPassword) {
		return true
	}

	return false
}

// SetDemoAccountPassword gets a reference to the given string and assigns it to the DemoAccountPassword field.
func (o *AppStoreReviewDetailAttributes) SetDemoAccountPassword(v string) {
	o.DemoAccountPassword = &v
}

// GetDemoAccountRequired returns the DemoAccountRequired field value if set, zero value otherwise.
func (o *AppStoreReviewDetailAttributes) GetDemoAccountRequired() bool {
	if o == nil || IsNil(o.DemoAccountRequired) {
		var ret bool
		return ret
	}
	return *o.DemoAccountRequired
}

// GetDemoAccountRequiredOk returns a tuple with the DemoAccountRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailAttributes) GetDemoAccountRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.DemoAccountRequired) {
		return nil, false
	}
	return o.DemoAccountRequired, true
}

// HasDemoAccountRequired returns a boolean if a field has been set.
func (o *AppStoreReviewDetailAttributes) HasDemoAccountRequired() bool {
	if o != nil && !IsNil(o.DemoAccountRequired) {
		return true
	}

	return false
}

// SetDemoAccountRequired gets a reference to the given bool and assigns it to the DemoAccountRequired field.
func (o *AppStoreReviewDetailAttributes) SetDemoAccountRequired(v bool) {
	o.DemoAccountRequired = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AppStoreReviewDetailAttributes) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailAttributes) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AppStoreReviewDetailAttributes) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AppStoreReviewDetailAttributes) SetNotes(v string) {
	o.Notes = &v
}

func (o AppStoreReviewDetailAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewDetailAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactFirstName) {
		toSerialize["contactFirstName"] = o.ContactFirstName
	}
	if !IsNil(o.ContactLastName) {
		toSerialize["contactLastName"] = o.ContactLastName
	}
	if !IsNil(o.ContactPhone) {
		toSerialize["contactPhone"] = o.ContactPhone
	}
	if !IsNil(o.ContactEmail) {
		toSerialize["contactEmail"] = o.ContactEmail
	}
	if !IsNil(o.DemoAccountName) {
		toSerialize["demoAccountName"] = o.DemoAccountName
	}
	if !IsNil(o.DemoAccountPassword) {
		toSerialize["demoAccountPassword"] = o.DemoAccountPassword
	}
	if !IsNil(o.DemoAccountRequired) {
		toSerialize["demoAccountRequired"] = o.DemoAccountRequired
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableAppStoreReviewDetailAttributes struct {
	value *AppStoreReviewDetailAttributes
	isSet bool
}

func (v NullableAppStoreReviewDetailAttributes) Get() *AppStoreReviewDetailAttributes {
	return v.value
}

func (v *NullableAppStoreReviewDetailAttributes) Set(val *AppStoreReviewDetailAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailAttributes(val *AppStoreReviewDetailAttributes) *NullableAppStoreReviewDetailAttributes {
	return &NullableAppStoreReviewDetailAttributes{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


