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

// checks if the ScmProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmProvider{}

// ScmProvider struct for ScmProvider
type ScmProvider struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *ScmProviderAttributes `json:"attributes,omitempty"`
	Relationships *ScmProviderRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _ScmProvider ScmProvider

// NewScmProvider instantiates a new ScmProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmProvider(type_ string, id string) *ScmProvider {
	this := ScmProvider{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewScmProviderWithDefaults instantiates a new ScmProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmProviderWithDefaults() *ScmProvider {
	this := ScmProvider{}
	return &this
}

// GetType returns the Type field value
func (o *ScmProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScmProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScmProvider) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ScmProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScmProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScmProvider) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ScmProvider) GetAttributes() ScmProviderAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ScmProviderAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmProvider) GetAttributesOk() (*ScmProviderAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ScmProvider) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ScmProviderAttributes and assigns it to the Attributes field.
func (o *ScmProvider) SetAttributes(v ScmProviderAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ScmProvider) GetRelationships() ScmProviderRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ScmProviderRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmProvider) GetRelationshipsOk() (*ScmProviderRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ScmProvider) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ScmProviderRelationships and assigns it to the Relationships field.
func (o *ScmProvider) SetRelationships(v ScmProviderRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ScmProvider) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmProvider) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ScmProvider) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *ScmProvider) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o ScmProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmProvider) ToMap() (map[string]interface{}, error) {
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

func (o *ScmProvider) UnmarshalJSON(data []byte) (err error) {
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

	varScmProvider := _ScmProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScmProvider)

	if err != nil {
		return err
	}

	*o = ScmProvider(varScmProvider)

	return err
}

type NullableScmProvider struct {
	value *ScmProvider
	isSet bool
}

func (v NullableScmProvider) Get() *ScmProvider {
	return v.value
}

func (v *NullableScmProvider) Set(val *ScmProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableScmProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableScmProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmProvider(val *ScmProvider) *NullableScmProvider {
	return &NullableScmProvider{value: val, isSet: true}
}

func (v NullableScmProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


