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

// checks if the ActorAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActorAttributes{}

// ActorAttributes struct for ActorAttributes
type ActorAttributes struct {
	ActorType *string `json:"actorType,omitempty"`
	UserFirstName *string `json:"userFirstName,omitempty"`
	UserLastName *string `json:"userLastName,omitempty"`
	UserEmail *string `json:"userEmail,omitempty"`
	ApiKeyId *string `json:"apiKeyId,omitempty"`
}

// NewActorAttributes instantiates a new ActorAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActorAttributes() *ActorAttributes {
	this := ActorAttributes{}
	return &this
}

// NewActorAttributesWithDefaults instantiates a new ActorAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorAttributesWithDefaults() *ActorAttributes {
	this := ActorAttributes{}
	return &this
}

// GetActorType returns the ActorType field value if set, zero value otherwise.
func (o *ActorAttributes) GetActorType() string {
	if o == nil || IsNil(o.ActorType) {
		var ret string
		return ret
	}
	return *o.ActorType
}

// GetActorTypeOk returns a tuple with the ActorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorAttributes) GetActorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActorType) {
		return nil, false
	}
	return o.ActorType, true
}

// HasActorType returns a boolean if a field has been set.
func (o *ActorAttributes) HasActorType() bool {
	if o != nil && !IsNil(o.ActorType) {
		return true
	}

	return false
}

// SetActorType gets a reference to the given string and assigns it to the ActorType field.
func (o *ActorAttributes) SetActorType(v string) {
	o.ActorType = &v
}

// GetUserFirstName returns the UserFirstName field value if set, zero value otherwise.
func (o *ActorAttributes) GetUserFirstName() string {
	if o == nil || IsNil(o.UserFirstName) {
		var ret string
		return ret
	}
	return *o.UserFirstName
}

// GetUserFirstNameOk returns a tuple with the UserFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorAttributes) GetUserFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserFirstName) {
		return nil, false
	}
	return o.UserFirstName, true
}

// HasUserFirstName returns a boolean if a field has been set.
func (o *ActorAttributes) HasUserFirstName() bool {
	if o != nil && !IsNil(o.UserFirstName) {
		return true
	}

	return false
}

// SetUserFirstName gets a reference to the given string and assigns it to the UserFirstName field.
func (o *ActorAttributes) SetUserFirstName(v string) {
	o.UserFirstName = &v
}

// GetUserLastName returns the UserLastName field value if set, zero value otherwise.
func (o *ActorAttributes) GetUserLastName() string {
	if o == nil || IsNil(o.UserLastName) {
		var ret string
		return ret
	}
	return *o.UserLastName
}

// GetUserLastNameOk returns a tuple with the UserLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorAttributes) GetUserLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserLastName) {
		return nil, false
	}
	return o.UserLastName, true
}

// HasUserLastName returns a boolean if a field has been set.
func (o *ActorAttributes) HasUserLastName() bool {
	if o != nil && !IsNil(o.UserLastName) {
		return true
	}

	return false
}

// SetUserLastName gets a reference to the given string and assigns it to the UserLastName field.
func (o *ActorAttributes) SetUserLastName(v string) {
	o.UserLastName = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *ActorAttributes) GetUserEmail() string {
	if o == nil || IsNil(o.UserEmail) {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorAttributes) GetUserEmailOk() (*string, bool) {
	if o == nil || IsNil(o.UserEmail) {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *ActorAttributes) HasUserEmail() bool {
	if o != nil && !IsNil(o.UserEmail) {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *ActorAttributes) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *ActorAttributes) GetApiKeyId() string {
	if o == nil || IsNil(o.ApiKeyId) {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActorAttributes) GetApiKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKeyId) {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *ActorAttributes) HasApiKeyId() bool {
	if o != nil && !IsNil(o.ApiKeyId) {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *ActorAttributes) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

func (o ActorAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActorAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActorType) {
		toSerialize["actorType"] = o.ActorType
	}
	if !IsNil(o.UserFirstName) {
		toSerialize["userFirstName"] = o.UserFirstName
	}
	if !IsNil(o.UserLastName) {
		toSerialize["userLastName"] = o.UserLastName
	}
	if !IsNil(o.UserEmail) {
		toSerialize["userEmail"] = o.UserEmail
	}
	if !IsNil(o.ApiKeyId) {
		toSerialize["apiKeyId"] = o.ApiKeyId
	}
	return toSerialize, nil
}

type NullableActorAttributes struct {
	value *ActorAttributes
	isSet bool
}

func (v NullableActorAttributes) Get() *ActorAttributes {
	return v.value
}

func (v *NullableActorAttributes) Set(val *ActorAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableActorAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableActorAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActorAttributes(val *ActorAttributes) *NullableActorAttributes {
	return &NullableActorAttributes{value: val, isSet: true}
}

func (v NullableActorAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActorAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


