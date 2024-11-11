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

// checks if the AppPreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewResponse{}

// AppPreviewResponse struct for AppPreviewResponse
type AppPreviewResponse struct {
	Data AppPreview `json:"data"`
	Included []AppPreviewSet `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _AppPreviewResponse AppPreviewResponse

// NewAppPreviewResponse instantiates a new AppPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewResponse(data AppPreview, links DocumentLinks) *AppPreviewResponse {
	this := AppPreviewResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppPreviewResponseWithDefaults instantiates a new AppPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewResponseWithDefaults() *AppPreviewResponse {
	this := AppPreviewResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppPreviewResponse) GetData() AppPreview {
	if o == nil {
		var ret AppPreview
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPreviewResponse) GetDataOk() (*AppPreview, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPreviewResponse) SetData(v AppPreview) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppPreviewResponse) GetIncluded() []AppPreviewSet {
	if o == nil || IsNil(o.Included) {
		var ret []AppPreviewSet
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewResponse) GetIncludedOk() ([]AppPreviewSet, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppPreviewResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppPreviewSet and assigns it to the Included field.
func (o *AppPreviewResponse) SetIncluded(v []AppPreviewSet) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppPreviewResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPreviewResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPreviewResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppPreviewResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppPreviewResponse := _AppPreviewResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppPreviewResponse)

	if err != nil {
		return err
	}

	*o = AppPreviewResponse(varAppPreviewResponse)

	return err
}

type NullableAppPreviewResponse struct {
	value *AppPreviewResponse
	isSet bool
}

func (v NullableAppPreviewResponse) Get() *AppPreviewResponse {
	return v.value
}

func (v *NullableAppPreviewResponse) Set(val *AppPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewResponse(val *AppPreviewResponse) *NullableAppPreviewResponse {
	return &NullableAppPreviewResponse{value: val, isSet: true}
}

func (v NullableAppPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


