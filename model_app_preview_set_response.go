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

// checks if the AppPreviewSetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetResponse{}

// AppPreviewSetResponse struct for AppPreviewSetResponse
type AppPreviewSetResponse struct {
	Data AppPreviewSet `json:"data"`
	Included []AppPreviewSetsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _AppPreviewSetResponse AppPreviewSetResponse

// NewAppPreviewSetResponse instantiates a new AppPreviewSetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetResponse(data AppPreviewSet, links DocumentLinks) *AppPreviewSetResponse {
	this := AppPreviewSetResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppPreviewSetResponseWithDefaults instantiates a new AppPreviewSetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetResponseWithDefaults() *AppPreviewSetResponse {
	this := AppPreviewSetResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppPreviewSetResponse) GetData() AppPreviewSet {
	if o == nil {
		var ret AppPreviewSet
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetResponse) GetDataOk() (*AppPreviewSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPreviewSetResponse) SetData(v AppPreviewSet) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppPreviewSetResponse) GetIncluded() []AppPreviewSetsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppPreviewSetsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetResponse) GetIncludedOk() ([]AppPreviewSetsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppPreviewSetResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppPreviewSetsResponseIncludedInner and assigns it to the Included field.
func (o *AppPreviewSetResponse) SetIncluded(v []AppPreviewSetsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppPreviewSetResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPreviewSetResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppPreviewSetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AppPreviewSetResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAppPreviewSetResponse := _AppPreviewSetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppPreviewSetResponse)

	if err != nil {
		return err
	}

	*o = AppPreviewSetResponse(varAppPreviewSetResponse)

	return err
}

type NullableAppPreviewSetResponse struct {
	value *AppPreviewSetResponse
	isSet bool
}

func (v NullableAppPreviewSetResponse) Get() *AppPreviewSetResponse {
	return v.value
}

func (v *NullableAppPreviewSetResponse) Set(val *AppPreviewSetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetResponse(val *AppPreviewSetResponse) *NullableAppPreviewSetResponse {
	return &NullableAppPreviewSetResponse{value: val, isSet: true}
}

func (v NullableAppPreviewSetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


