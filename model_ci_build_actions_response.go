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

// checks if the CiBuildActionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildActionsResponse{}

// CiBuildActionsResponse struct for CiBuildActionsResponse
type CiBuildActionsResponse struct {
	Data []CiBuildAction `json:"data"`
	Included []CiBuildRun `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _CiBuildActionsResponse CiBuildActionsResponse

// NewCiBuildActionsResponse instantiates a new CiBuildActionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildActionsResponse(data []CiBuildAction, links PagedDocumentLinks) *CiBuildActionsResponse {
	this := CiBuildActionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiBuildActionsResponseWithDefaults instantiates a new CiBuildActionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildActionsResponseWithDefaults() *CiBuildActionsResponse {
	this := CiBuildActionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiBuildActionsResponse) GetData() []CiBuildAction {
	if o == nil {
		var ret []CiBuildAction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiBuildActionsResponse) GetDataOk() ([]CiBuildAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CiBuildActionsResponse) SetData(v []CiBuildAction) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CiBuildActionsResponse) GetIncluded() []CiBuildRun {
	if o == nil || IsNil(o.Included) {
		var ret []CiBuildRun
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionsResponse) GetIncludedOk() ([]CiBuildRun, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CiBuildActionsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CiBuildRun and assigns it to the Included field.
func (o *CiBuildActionsResponse) SetIncluded(v []CiBuildRun) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *CiBuildActionsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiBuildActionsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiBuildActionsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CiBuildActionsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CiBuildActionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *CiBuildActionsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o CiBuildActionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildActionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *CiBuildActionsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varCiBuildActionsResponse := _CiBuildActionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiBuildActionsResponse)

	if err != nil {
		return err
	}

	*o = CiBuildActionsResponse(varCiBuildActionsResponse)

	return err
}

type NullableCiBuildActionsResponse struct {
	value *CiBuildActionsResponse
	isSet bool
}

func (v NullableCiBuildActionsResponse) Get() *CiBuildActionsResponse {
	return v.value
}

func (v *NullableCiBuildActionsResponse) Set(val *CiBuildActionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildActionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildActionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildActionsResponse(val *CiBuildActionsResponse) *NullableCiBuildActionsResponse {
	return &NullableCiBuildActionsResponse{value: val, isSet: true}
}

func (v NullableCiBuildActionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildActionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


