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

// checks if the BetaTesterBuildsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterBuildsLinkagesResponse{}

// BetaTesterBuildsLinkagesResponse struct for BetaTesterBuildsLinkagesResponse
type BetaTesterBuildsLinkagesResponse struct {
	Data []AppEncryptionDeclarationRelationshipsBuildsDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _BetaTesterBuildsLinkagesResponse BetaTesterBuildsLinkagesResponse

// NewBetaTesterBuildsLinkagesResponse instantiates a new BetaTesterBuildsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterBuildsLinkagesResponse(data []AppEncryptionDeclarationRelationshipsBuildsDataInner, links PagedDocumentLinks) *BetaTesterBuildsLinkagesResponse {
	this := BetaTesterBuildsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaTesterBuildsLinkagesResponseWithDefaults instantiates a new BetaTesterBuildsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterBuildsLinkagesResponseWithDefaults() *BetaTesterBuildsLinkagesResponse {
	this := BetaTesterBuildsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTesterBuildsLinkagesResponse) GetData() []AppEncryptionDeclarationRelationshipsBuildsDataInner {
	if o == nil {
		var ret []AppEncryptionDeclarationRelationshipsBuildsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTesterBuildsLinkagesResponse) GetDataOk() ([]AppEncryptionDeclarationRelationshipsBuildsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaTesterBuildsLinkagesResponse) SetData(v []AppEncryptionDeclarationRelationshipsBuildsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaTesterBuildsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaTesterBuildsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaTesterBuildsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaTesterBuildsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterBuildsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaTesterBuildsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaTesterBuildsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaTesterBuildsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterBuildsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *BetaTesterBuildsLinkagesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBetaTesterBuildsLinkagesResponse := _BetaTesterBuildsLinkagesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaTesterBuildsLinkagesResponse)

	if err != nil {
		return err
	}

	*o = BetaTesterBuildsLinkagesResponse(varBetaTesterBuildsLinkagesResponse)

	return err
}

type NullableBetaTesterBuildsLinkagesResponse struct {
	value *BetaTesterBuildsLinkagesResponse
	isSet bool
}

func (v NullableBetaTesterBuildsLinkagesResponse) Get() *BetaTesterBuildsLinkagesResponse {
	return v.value
}

func (v *NullableBetaTesterBuildsLinkagesResponse) Set(val *BetaTesterBuildsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterBuildsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterBuildsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterBuildsLinkagesResponse(val *BetaTesterBuildsLinkagesResponse) *NullableBetaTesterBuildsLinkagesResponse {
	return &NullableBetaTesterBuildsLinkagesResponse{value: val, isSet: true}
}

func (v NullableBetaTesterBuildsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterBuildsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


