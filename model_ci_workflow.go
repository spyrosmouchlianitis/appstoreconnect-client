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

// checks if the CiWorkflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflow{}

// CiWorkflow struct for CiWorkflow
type CiWorkflow struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *CiWorkflowAttributes `json:"attributes,omitempty"`
	Relationships *CiWorkflowRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _CiWorkflow CiWorkflow

// NewCiWorkflow instantiates a new CiWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflow(type_ string, id string) *CiWorkflow {
	this := CiWorkflow{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiWorkflowWithDefaults instantiates a new CiWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowWithDefaults() *CiWorkflow {
	this := CiWorkflow{}
	return &this
}

// GetType returns the Type field value
func (o *CiWorkflow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiWorkflow) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiWorkflow) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiWorkflow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiWorkflow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiWorkflow) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CiWorkflow) GetAttributes() CiWorkflowAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CiWorkflowAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflow) GetAttributesOk() (*CiWorkflowAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CiWorkflow) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CiWorkflowAttributes and assigns it to the Attributes field.
func (o *CiWorkflow) SetAttributes(v CiWorkflowAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CiWorkflow) GetRelationships() CiWorkflowRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CiWorkflowRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflow) GetRelationshipsOk() (*CiWorkflowRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CiWorkflow) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CiWorkflowRelationships and assigns it to the Relationships field.
func (o *CiWorkflow) SetRelationships(v CiWorkflowRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CiWorkflow) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflow) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CiWorkflow) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *CiWorkflow) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o CiWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflow) ToMap() (map[string]interface{}, error) {
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

func (o *CiWorkflow) UnmarshalJSON(data []byte) (err error) {
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

	varCiWorkflow := _CiWorkflow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiWorkflow)

	if err != nil {
		return err
	}

	*o = CiWorkflow(varCiWorkflow)

	return err
}

type NullableCiWorkflow struct {
	value *CiWorkflow
	isSet bool
}

func (v NullableCiWorkflow) Get() *CiWorkflow {
	return v.value
}

func (v *NullableCiWorkflow) Set(val *CiWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflow(val *CiWorkflow) *NullableCiWorkflow {
	return &NullableCiWorkflow{value: val, isSet: true}
}

func (v NullableCiWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

