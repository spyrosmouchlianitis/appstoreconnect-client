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

// checks if the AnalyticsReportInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsReportInstance{}

// AnalyticsReportInstance struct for AnalyticsReportInstance
type AnalyticsReportInstance struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AnalyticsReportInstanceAttributes `json:"attributes,omitempty"`
	Relationships *AnalyticsReportInstanceRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _AnalyticsReportInstance AnalyticsReportInstance

// NewAnalyticsReportInstance instantiates a new AnalyticsReportInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsReportInstance(type_ string, id string) *AnalyticsReportInstance {
	this := AnalyticsReportInstance{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAnalyticsReportInstanceWithDefaults instantiates a new AnalyticsReportInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsReportInstanceWithDefaults() *AnalyticsReportInstance {
	this := AnalyticsReportInstance{}
	return &this
}

// GetType returns the Type field value
func (o *AnalyticsReportInstance) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportInstance) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AnalyticsReportInstance) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AnalyticsReportInstance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportInstance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AnalyticsReportInstance) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AnalyticsReportInstance) GetAttributes() AnalyticsReportInstanceAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AnalyticsReportInstanceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportInstance) GetAttributesOk() (*AnalyticsReportInstanceAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AnalyticsReportInstance) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AnalyticsReportInstanceAttributes and assigns it to the Attributes field.
func (o *AnalyticsReportInstance) SetAttributes(v AnalyticsReportInstanceAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AnalyticsReportInstance) GetRelationships() AnalyticsReportInstanceRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AnalyticsReportInstanceRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportInstance) GetRelationshipsOk() (*AnalyticsReportInstanceRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AnalyticsReportInstance) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AnalyticsReportInstanceRelationships and assigns it to the Relationships field.
func (o *AnalyticsReportInstance) SetRelationships(v AnalyticsReportInstanceRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AnalyticsReportInstance) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportInstance) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AnalyticsReportInstance) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AnalyticsReportInstance) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AnalyticsReportInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsReportInstance) ToMap() (map[string]interface{}, error) {
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

func (o *AnalyticsReportInstance) UnmarshalJSON(data []byte) (err error) {
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

	varAnalyticsReportInstance := _AnalyticsReportInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalyticsReportInstance)

	if err != nil {
		return err
	}

	*o = AnalyticsReportInstance(varAnalyticsReportInstance)

	return err
}

type NullableAnalyticsReportInstance struct {
	value *AnalyticsReportInstance
	isSet bool
}

func (v NullableAnalyticsReportInstance) Get() *AnalyticsReportInstance {
	return v.value
}

func (v *NullableAnalyticsReportInstance) Set(val *AnalyticsReportInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsReportInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsReportInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsReportInstance(val *AnalyticsReportInstance) *NullableAnalyticsReportInstance {
	return &NullableAnalyticsReportInstance{value: val, isSet: true}
}

func (v NullableAnalyticsReportInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsReportInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

