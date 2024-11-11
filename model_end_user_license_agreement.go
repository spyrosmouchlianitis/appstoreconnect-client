/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EndUserLicenseAgreement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndUserLicenseAgreement{}

// EndUserLicenseAgreement struct for EndUserLicenseAgreement
type EndUserLicenseAgreement struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BetaLicenseAgreementAttributes `json:"attributes,omitempty"`
	Relationships *EndUserLicenseAgreementRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _EndUserLicenseAgreement EndUserLicenseAgreement

// NewEndUserLicenseAgreement instantiates a new EndUserLicenseAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreement(type_ string, id string) *EndUserLicenseAgreement {
	this := EndUserLicenseAgreement{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewEndUserLicenseAgreementWithDefaults instantiates a new EndUserLicenseAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementWithDefaults() *EndUserLicenseAgreement {
	this := EndUserLicenseAgreement{}
	return &this
}

// GetType returns the Type field value
func (o *EndUserLicenseAgreement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EndUserLicenseAgreement) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *EndUserLicenseAgreement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EndUserLicenseAgreement) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EndUserLicenseAgreement) GetAttributes() BetaLicenseAgreementAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BetaLicenseAgreementAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreement) GetAttributesOk() (*BetaLicenseAgreementAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EndUserLicenseAgreement) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaLicenseAgreementAttributes and assigns it to the Attributes field.
func (o *EndUserLicenseAgreement) SetAttributes(v BetaLicenseAgreementAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *EndUserLicenseAgreement) GetRelationships() EndUserLicenseAgreementRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret EndUserLicenseAgreementRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreement) GetRelationshipsOk() (*EndUserLicenseAgreementRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *EndUserLicenseAgreement) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given EndUserLicenseAgreementRelationships and assigns it to the Relationships field.
func (o *EndUserLicenseAgreement) SetRelationships(v EndUserLicenseAgreementRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EndUserLicenseAgreement) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreement) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EndUserLicenseAgreement) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *EndUserLicenseAgreement) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o EndUserLicenseAgreement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndUserLicenseAgreement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *EndUserLicenseAgreement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varEndUserLicenseAgreement := _EndUserLicenseAgreement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndUserLicenseAgreement)

	if err != nil {
		return err
	}

	*o = EndUserLicenseAgreement(varEndUserLicenseAgreement)

	return err
}

type NullableEndUserLicenseAgreement struct {
	value *EndUserLicenseAgreement
	isSet bool
}

func (v NullableEndUserLicenseAgreement) Get() *EndUserLicenseAgreement {
	return v.value
}

func (v *NullableEndUserLicenseAgreement) Set(val *EndUserLicenseAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreement(val *EndUserLicenseAgreement) *NullableEndUserLicenseAgreement {
	return &NullableEndUserLicenseAgreement{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

