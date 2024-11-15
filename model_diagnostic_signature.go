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

// checks if the DiagnosticSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosticSignature{}

// DiagnosticSignature struct for DiagnosticSignature
type DiagnosticSignature struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *DiagnosticSignatureAttributes `json:"attributes,omitempty"`
	Relationships *DiagnosticSignatureRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _DiagnosticSignature DiagnosticSignature

// NewDiagnosticSignature instantiates a new DiagnosticSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticSignature(type_ string, id string) *DiagnosticSignature {
	this := DiagnosticSignature{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewDiagnosticSignatureWithDefaults instantiates a new DiagnosticSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticSignatureWithDefaults() *DiagnosticSignature {
	this := DiagnosticSignature{}
	return &this
}

// GetType returns the Type field value
func (o *DiagnosticSignature) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSignature) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DiagnosticSignature) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DiagnosticSignature) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSignature) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiagnosticSignature) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DiagnosticSignature) GetAttributes() DiagnosticSignatureAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret DiagnosticSignatureAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSignature) GetAttributesOk() (*DiagnosticSignatureAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DiagnosticSignature) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given DiagnosticSignatureAttributes and assigns it to the Attributes field.
func (o *DiagnosticSignature) SetAttributes(v DiagnosticSignatureAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *DiagnosticSignature) GetRelationships() DiagnosticSignatureRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret DiagnosticSignatureRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSignature) GetRelationshipsOk() (*DiagnosticSignatureRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *DiagnosticSignature) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given DiagnosticSignatureRelationships and assigns it to the Relationships field.
func (o *DiagnosticSignature) SetRelationships(v DiagnosticSignatureRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DiagnosticSignature) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSignature) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiagnosticSignature) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *DiagnosticSignature) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o DiagnosticSignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosticSignature) ToMap() (map[string]interface{}, error) {
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

func (o *DiagnosticSignature) UnmarshalJSON(data []byte) (err error) {
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

	varDiagnosticSignature := _DiagnosticSignature{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiagnosticSignature)

	if err != nil {
		return err
	}

	*o = DiagnosticSignature(varDiagnosticSignature)

	return err
}

type NullableDiagnosticSignature struct {
	value *DiagnosticSignature
	isSet bool
}

func (v NullableDiagnosticSignature) Get() *DiagnosticSignature {
	return v.value
}

func (v *NullableDiagnosticSignature) Set(val *DiagnosticSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticSignature(val *DiagnosticSignature) *NullableDiagnosticSignature {
	return &NullableDiagnosticSignature{value: val, isSet: true}
}

func (v NullableDiagnosticSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


