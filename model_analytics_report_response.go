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

// checks if the AnalyticsReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsReportResponse{}

// AnalyticsReportResponse struct for AnalyticsReportResponse
type AnalyticsReportResponse struct {
	Data AnalyticsReport `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _AnalyticsReportResponse AnalyticsReportResponse

// NewAnalyticsReportResponse instantiates a new AnalyticsReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsReportResponse(data AnalyticsReport, links DocumentLinks) *AnalyticsReportResponse {
	this := AnalyticsReportResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAnalyticsReportResponseWithDefaults instantiates a new AnalyticsReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsReportResponseWithDefaults() *AnalyticsReportResponse {
	this := AnalyticsReportResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AnalyticsReportResponse) GetData() AnalyticsReport {
	if o == nil {
		var ret AnalyticsReport
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportResponse) GetDataOk() (*AnalyticsReport, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AnalyticsReportResponse) SetData(v AnalyticsReport) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AnalyticsReportResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AnalyticsReportResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AnalyticsReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AnalyticsReportResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAnalyticsReportResponse := _AnalyticsReportResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalyticsReportResponse)

	if err != nil {
		return err
	}

	*o = AnalyticsReportResponse(varAnalyticsReportResponse)

	return err
}

type NullableAnalyticsReportResponse struct {
	value *AnalyticsReportResponse
	isSet bool
}

func (v NullableAnalyticsReportResponse) Get() *AnalyticsReportResponse {
	return v.value
}

func (v *NullableAnalyticsReportResponse) Set(val *AnalyticsReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsReportResponse(val *AnalyticsReportResponse) *NullableAnalyticsReportResponse {
	return &NullableAnalyticsReportResponse{value: val, isSet: true}
}

func (v NullableAnalyticsReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


