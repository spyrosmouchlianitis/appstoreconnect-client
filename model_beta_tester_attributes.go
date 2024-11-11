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

// checks if the BetaTesterAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterAttributes{}

// BetaTesterAttributes struct for BetaTesterAttributes
type BetaTesterAttributes struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Email *string `json:"email,omitempty"`
	InviteType *BetaInviteType `json:"inviteType,omitempty"`
	State *BetaTesterState `json:"state,omitempty"`
}

// NewBetaTesterAttributes instantiates a new BetaTesterAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterAttributes() *BetaTesterAttributes {
	this := BetaTesterAttributes{}
	return &this
}

// NewBetaTesterAttributesWithDefaults instantiates a new BetaTesterAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterAttributesWithDefaults() *BetaTesterAttributes {
	this := BetaTesterAttributes{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BetaTesterAttributes) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BetaTesterAttributes) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BetaTesterAttributes) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BetaTesterAttributes) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterAttributes) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BetaTesterAttributes) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BetaTesterAttributes) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BetaTesterAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterAttributes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BetaTesterAttributes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BetaTesterAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetInviteType returns the InviteType field value if set, zero value otherwise.
func (o *BetaTesterAttributes) GetInviteType() BetaInviteType {
	if o == nil || IsNil(o.InviteType) {
		var ret BetaInviteType
		return ret
	}
	return *o.InviteType
}

// GetInviteTypeOk returns a tuple with the InviteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterAttributes) GetInviteTypeOk() (*BetaInviteType, bool) {
	if o == nil || IsNil(o.InviteType) {
		return nil, false
	}
	return o.InviteType, true
}

// HasInviteType returns a boolean if a field has been set.
func (o *BetaTesterAttributes) HasInviteType() bool {
	if o != nil && !IsNil(o.InviteType) {
		return true
	}

	return false
}

// SetInviteType gets a reference to the given BetaInviteType and assigns it to the InviteType field.
func (o *BetaTesterAttributes) SetInviteType(v BetaInviteType) {
	o.InviteType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BetaTesterAttributes) GetState() BetaTesterState {
	if o == nil || IsNil(o.State) {
		var ret BetaTesterState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterAttributes) GetStateOk() (*BetaTesterState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BetaTesterAttributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given BetaTesterState and assigns it to the State field.
func (o *BetaTesterAttributes) SetState(v BetaTesterState) {
	o.State = &v
}

func (o BetaTesterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.InviteType) {
		toSerialize["inviteType"] = o.InviteType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableBetaTesterAttributes struct {
	value *BetaTesterAttributes
	isSet bool
}

func (v NullableBetaTesterAttributes) Get() *BetaTesterAttributes {
	return v.value
}

func (v *NullableBetaTesterAttributes) Set(val *BetaTesterAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterAttributes(val *BetaTesterAttributes) *NullableBetaTesterAttributes {
	return &NullableBetaTesterAttributes{value: val, isSet: true}
}

func (v NullableBetaTesterAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


