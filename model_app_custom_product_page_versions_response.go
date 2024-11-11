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

// checks if the AppCustomProductPageVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionsResponse{}

// AppCustomProductPageVersionsResponse struct for AppCustomProductPageVersionsResponse
type AppCustomProductPageVersionsResponse struct {
	Data []AppCustomProductPageVersion `json:"data"`
	Included []AppCustomProductPageVersionsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _AppCustomProductPageVersionsResponse AppCustomProductPageVersionsResponse

// NewAppCustomProductPageVersionsResponse instantiates a new AppCustomProductPageVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionsResponse(data []AppCustomProductPageVersion, links PagedDocumentLinks) *AppCustomProductPageVersionsResponse {
	this := AppCustomProductPageVersionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppCustomProductPageVersionsResponseWithDefaults instantiates a new AppCustomProductPageVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionsResponseWithDefaults() *AppCustomProductPageVersionsResponse {
	this := AppCustomProductPageVersionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppCustomProductPageVersionsResponse) GetData() []AppCustomProductPageVersion {
	if o == nil {
		var ret []AppCustomProductPageVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionsResponse) GetDataOk() ([]AppCustomProductPageVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppCustomProductPageVersionsResponse) SetData(v []AppCustomProductPageVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionsResponse) GetIncluded() []AppCustomProductPageVersionsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppCustomProductPageVersionsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionsResponse) GetIncludedOk() ([]AppCustomProductPageVersionsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppCustomProductPageVersionsResponseIncludedInner and assigns it to the Included field.
func (o *AppCustomProductPageVersionsResponse) SetIncluded(v []AppCustomProductPageVersionsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppCustomProductPageVersionsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppCustomProductPageVersionsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppCustomProductPageVersionsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppCustomProductPageVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AppCustomProductPageVersionsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppCustomProductPageVersionsResponse := _AppCustomProductPageVersionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppCustomProductPageVersionsResponse)

	if err != nil {
		return err
	}

	*o = AppCustomProductPageVersionsResponse(varAppCustomProductPageVersionsResponse)

	return err
}

type NullableAppCustomProductPageVersionsResponse struct {
	value *AppCustomProductPageVersionsResponse
	isSet bool
}

func (v NullableAppCustomProductPageVersionsResponse) Get() *AppCustomProductPageVersionsResponse {
	return v.value
}

func (v *NullableAppCustomProductPageVersionsResponse) Set(val *AppCustomProductPageVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionsResponse(val *AppCustomProductPageVersionsResponse) *NullableAppCustomProductPageVersionsResponse {
	return &NullableAppCustomProductPageVersionsResponse{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

