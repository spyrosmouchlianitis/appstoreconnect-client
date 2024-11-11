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

// checks if the BetaTesterUsagesV1MetricResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterUsagesV1MetricResponse{}

// BetaTesterUsagesV1MetricResponse struct for BetaTesterUsagesV1MetricResponse
type BetaTesterUsagesV1MetricResponse struct {
	Data []BetaTesterUsagesV1MetricResponseDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _BetaTesterUsagesV1MetricResponse BetaTesterUsagesV1MetricResponse

// NewBetaTesterUsagesV1MetricResponse instantiates a new BetaTesterUsagesV1MetricResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterUsagesV1MetricResponse(data []BetaTesterUsagesV1MetricResponseDataInner, links PagedDocumentLinks) *BetaTesterUsagesV1MetricResponse {
	this := BetaTesterUsagesV1MetricResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaTesterUsagesV1MetricResponseWithDefaults instantiates a new BetaTesterUsagesV1MetricResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterUsagesV1MetricResponseWithDefaults() *BetaTesterUsagesV1MetricResponse {
	this := BetaTesterUsagesV1MetricResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTesterUsagesV1MetricResponse) GetData() []BetaTesterUsagesV1MetricResponseDataInner {
	if o == nil {
		var ret []BetaTesterUsagesV1MetricResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTesterUsagesV1MetricResponse) GetDataOk() ([]BetaTesterUsagesV1MetricResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaTesterUsagesV1MetricResponse) SetData(v []BetaTesterUsagesV1MetricResponseDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaTesterUsagesV1MetricResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaTesterUsagesV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaTesterUsagesV1MetricResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaTesterUsagesV1MetricResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterUsagesV1MetricResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaTesterUsagesV1MetricResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaTesterUsagesV1MetricResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaTesterUsagesV1MetricResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterUsagesV1MetricResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *BetaTesterUsagesV1MetricResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBetaTesterUsagesV1MetricResponse := _BetaTesterUsagesV1MetricResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaTesterUsagesV1MetricResponse)

	if err != nil {
		return err
	}

	*o = BetaTesterUsagesV1MetricResponse(varBetaTesterUsagesV1MetricResponse)

	return err
}

type NullableBetaTesterUsagesV1MetricResponse struct {
	value *BetaTesterUsagesV1MetricResponse
	isSet bool
}

func (v NullableBetaTesterUsagesV1MetricResponse) Get() *BetaTesterUsagesV1MetricResponse {
	return v.value
}

func (v *NullableBetaTesterUsagesV1MetricResponse) Set(val *BetaTesterUsagesV1MetricResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterUsagesV1MetricResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterUsagesV1MetricResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterUsagesV1MetricResponse(val *BetaTesterUsagesV1MetricResponse) *NullableBetaTesterUsagesV1MetricResponse {
	return &NullableBetaTesterUsagesV1MetricResponse{value: val, isSet: true}
}

func (v NullableBetaTesterUsagesV1MetricResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterUsagesV1MetricResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


