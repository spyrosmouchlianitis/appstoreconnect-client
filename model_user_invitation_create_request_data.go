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

// checks if the UserInvitationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitationCreateRequestData{}

// UserInvitationCreateRequestData struct for UserInvitationCreateRequestData
type UserInvitationCreateRequestData struct {
	Type string `json:"type"`
	Attributes UserInvitationCreateRequestDataAttributes `json:"attributes"`
	Relationships *UserInvitationCreateRequestDataRelationships `json:"relationships,omitempty"`
}

type _UserInvitationCreateRequestData UserInvitationCreateRequestData

// NewUserInvitationCreateRequestData instantiates a new UserInvitationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationCreateRequestData(type_ string, attributes UserInvitationCreateRequestDataAttributes) *UserInvitationCreateRequestData {
	this := UserInvitationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewUserInvitationCreateRequestDataWithDefaults instantiates a new UserInvitationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationCreateRequestDataWithDefaults() *UserInvitationCreateRequestData {
	this := UserInvitationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *UserInvitationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserInvitationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *UserInvitationCreateRequestData) GetAttributes() UserInvitationCreateRequestDataAttributes {
	if o == nil {
		var ret UserInvitationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestData) GetAttributesOk() (*UserInvitationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UserInvitationCreateRequestData) SetAttributes(v UserInvitationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *UserInvitationCreateRequestData) GetRelationships() UserInvitationCreateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret UserInvitationCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestData) GetRelationshipsOk() (*UserInvitationCreateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *UserInvitationCreateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given UserInvitationCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *UserInvitationCreateRequestData) SetRelationships(v UserInvitationCreateRequestDataRelationships) {
	o.Relationships = &v
}

func (o UserInvitationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *UserInvitationCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varUserInvitationCreateRequestData := _UserInvitationCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserInvitationCreateRequestData)

	if err != nil {
		return err
	}

	*o = UserInvitationCreateRequestData(varUserInvitationCreateRequestData)

	return err
}

type NullableUserInvitationCreateRequestData struct {
	value *UserInvitationCreateRequestData
	isSet bool
}

func (v NullableUserInvitationCreateRequestData) Get() *UserInvitationCreateRequestData {
	return v.value
}

func (v *NullableUserInvitationCreateRequestData) Set(val *UserInvitationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationCreateRequestData(val *UserInvitationCreateRequestData) *NullableUserInvitationCreateRequestData {
	return &NullableUserInvitationCreateRequestData{value: val, isSet: true}
}

func (v NullableUserInvitationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


