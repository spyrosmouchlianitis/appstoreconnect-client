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

// checks if the BetaAppClipInvocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocation{}

// BetaAppClipInvocation struct for BetaAppClipInvocation
type BetaAppClipInvocation struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BetaAppClipInvocationAttributes `json:"attributes,omitempty"`
	Relationships *BetaAppClipInvocationRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _BetaAppClipInvocation BetaAppClipInvocation

// NewBetaAppClipInvocation instantiates a new BetaAppClipInvocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocation(type_ string, id string) *BetaAppClipInvocation {
	this := BetaAppClipInvocation{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBetaAppClipInvocationWithDefaults instantiates a new BetaAppClipInvocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationWithDefaults() *BetaAppClipInvocation {
	this := BetaAppClipInvocation{}
	return &this
}

// GetType returns the Type field value
func (o *BetaAppClipInvocation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaAppClipInvocation) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaAppClipInvocation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaAppClipInvocation) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BetaAppClipInvocation) GetAttributes() BetaAppClipInvocationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BetaAppClipInvocationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocation) GetAttributesOk() (*BetaAppClipInvocationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BetaAppClipInvocation) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaAppClipInvocationAttributes and assigns it to the Attributes field.
func (o *BetaAppClipInvocation) SetAttributes(v BetaAppClipInvocationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BetaAppClipInvocation) GetRelationships() BetaAppClipInvocationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BetaAppClipInvocationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocation) GetRelationshipsOk() (*BetaAppClipInvocationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BetaAppClipInvocation) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BetaAppClipInvocationRelationships and assigns it to the Relationships field.
func (o *BetaAppClipInvocation) SetRelationships(v BetaAppClipInvocationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BetaAppClipInvocation) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocation) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BetaAppClipInvocation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *BetaAppClipInvocation) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o BetaAppClipInvocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocation) ToMap() (map[string]interface{}, error) {
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

func (o *BetaAppClipInvocation) UnmarshalJSON(data []byte) (err error) {
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

	varBetaAppClipInvocation := _BetaAppClipInvocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaAppClipInvocation)

	if err != nil {
		return err
	}

	*o = BetaAppClipInvocation(varBetaAppClipInvocation)

	return err
}

type NullableBetaAppClipInvocation struct {
	value *BetaAppClipInvocation
	isSet bool
}

func (v NullableBetaAppClipInvocation) Get() *BetaAppClipInvocation {
	return v.value
}

func (v *NullableBetaAppClipInvocation) Set(val *BetaAppClipInvocation) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocation) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocation(val *BetaAppClipInvocation) *NullableBetaAppClipInvocation {
	return &NullableBetaAppClipInvocation{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


