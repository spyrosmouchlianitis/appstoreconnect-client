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

// checks if the DevicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesResponse{}

// DevicesResponse struct for DevicesResponse
type DevicesResponse struct {
	Data []Device `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _DevicesResponse DevicesResponse

// NewDevicesResponse instantiates a new DevicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesResponse(data []Device, links PagedDocumentLinks) *DevicesResponse {
	this := DevicesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewDevicesResponseWithDefaults instantiates a new DevicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesResponseWithDefaults() *DevicesResponse {
	this := DevicesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DevicesResponse) GetData() []Device {
	if o == nil {
		var ret []Device
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DevicesResponse) GetDataOk() ([]Device, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DevicesResponse) SetData(v []Device) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *DevicesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DevicesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DevicesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DevicesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DevicesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *DevicesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o DevicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *DevicesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDevicesResponse := _DevicesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDevicesResponse)

	if err != nil {
		return err
	}

	*o = DevicesResponse(varDevicesResponse)

	return err
}

type NullableDevicesResponse struct {
	value *DevicesResponse
	isSet bool
}

func (v NullableDevicesResponse) Get() *DevicesResponse {
	return v.value
}

func (v *NullableDevicesResponse) Set(val *DevicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesResponse(val *DevicesResponse) *NullableDevicesResponse {
	return &NullableDevicesResponse{value: val, isSet: true}
}

func (v NullableDevicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


