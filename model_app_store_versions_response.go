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

// checks if the AppStoreVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionsResponse{}

// AppStoreVersionsResponse struct for AppStoreVersionsResponse
type AppStoreVersionsResponse struct {
	Data []AppStoreVersion `json:"data"`
	Included []AppStoreVersionsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _AppStoreVersionsResponse AppStoreVersionsResponse

// NewAppStoreVersionsResponse instantiates a new AppStoreVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionsResponse(data []AppStoreVersion, links PagedDocumentLinks) *AppStoreVersionsResponse {
	this := AppStoreVersionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionsResponseWithDefaults instantiates a new AppStoreVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionsResponseWithDefaults() *AppStoreVersionsResponse {
	this := AppStoreVersionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionsResponse) GetData() []AppStoreVersion {
	if o == nil {
		var ret []AppStoreVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionsResponse) GetDataOk() ([]AppStoreVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionsResponse) SetData(v []AppStoreVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppStoreVersionsResponse) GetIncluded() []AppStoreVersionsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppStoreVersionsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionsResponse) GetIncludedOk() ([]AppStoreVersionsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppStoreVersionsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppStoreVersionsResponseIncludedInner and assigns it to the Included field.
func (o *AppStoreVersionsResponse) SetIncluded(v []AppStoreVersionsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppStoreVersionsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppStoreVersionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppStoreVersionsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppStoreVersionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AppStoreVersionsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionsResponse := _AppStoreVersionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionsResponse)

	if err != nil {
		return err
	}

	*o = AppStoreVersionsResponse(varAppStoreVersionsResponse)

	return err
}

type NullableAppStoreVersionsResponse struct {
	value *AppStoreVersionsResponse
	isSet bool
}

func (v NullableAppStoreVersionsResponse) Get() *AppStoreVersionsResponse {
	return v.value
}

func (v *NullableAppStoreVersionsResponse) Set(val *AppStoreVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionsResponse(val *AppStoreVersionsResponse) *NullableAppStoreVersionsResponse {
	return &NullableAppStoreVersionsResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


