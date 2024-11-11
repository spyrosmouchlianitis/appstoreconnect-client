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

// checks if the AppClipHeaderImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipHeaderImageResponse{}

// AppClipHeaderImageResponse struct for AppClipHeaderImageResponse
type AppClipHeaderImageResponse struct {
	Data AppClipHeaderImage `json:"data"`
	Included []AppClipDefaultExperienceLocalization `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _AppClipHeaderImageResponse AppClipHeaderImageResponse

// NewAppClipHeaderImageResponse instantiates a new AppClipHeaderImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipHeaderImageResponse(data AppClipHeaderImage, links DocumentLinks) *AppClipHeaderImageResponse {
	this := AppClipHeaderImageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppClipHeaderImageResponseWithDefaults instantiates a new AppClipHeaderImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipHeaderImageResponseWithDefaults() *AppClipHeaderImageResponse {
	this := AppClipHeaderImageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipHeaderImageResponse) GetData() AppClipHeaderImage {
	if o == nil {
		var ret AppClipHeaderImage
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageResponse) GetDataOk() (*AppClipHeaderImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipHeaderImageResponse) SetData(v AppClipHeaderImage) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppClipHeaderImageResponse) GetIncluded() []AppClipDefaultExperienceLocalization {
	if o == nil || IsNil(o.Included) {
		var ret []AppClipDefaultExperienceLocalization
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageResponse) GetIncludedOk() ([]AppClipDefaultExperienceLocalization, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppClipHeaderImageResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppClipDefaultExperienceLocalization and assigns it to the Included field.
func (o *AppClipHeaderImageResponse) SetIncluded(v []AppClipDefaultExperienceLocalization) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppClipHeaderImageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipHeaderImageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppClipHeaderImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipHeaderImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppClipHeaderImageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipHeaderImageResponse := _AppClipHeaderImageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipHeaderImageResponse)

	if err != nil {
		return err
	}

	*o = AppClipHeaderImageResponse(varAppClipHeaderImageResponse)

	return err
}

type NullableAppClipHeaderImageResponse struct {
	value *AppClipHeaderImageResponse
	isSet bool
}

func (v NullableAppClipHeaderImageResponse) Get() *AppClipHeaderImageResponse {
	return v.value
}

func (v *NullableAppClipHeaderImageResponse) Set(val *AppClipHeaderImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipHeaderImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipHeaderImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipHeaderImageResponse(val *AppClipHeaderImageResponse) *NullableAppClipHeaderImageResponse {
	return &NullableAppClipHeaderImageResponse{value: val, isSet: true}
}

func (v NullableAppClipHeaderImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipHeaderImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


