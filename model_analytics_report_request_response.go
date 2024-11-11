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

// checks if the AnalyticsReportRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsReportRequestResponse{}

// AnalyticsReportRequestResponse struct for AnalyticsReportRequestResponse
type AnalyticsReportRequestResponse struct {
	Data AnalyticsReportRequest `json:"data"`
	Included []AnalyticsReport `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _AnalyticsReportRequestResponse AnalyticsReportRequestResponse

// NewAnalyticsReportRequestResponse instantiates a new AnalyticsReportRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsReportRequestResponse(data AnalyticsReportRequest, links DocumentLinks) *AnalyticsReportRequestResponse {
	this := AnalyticsReportRequestResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAnalyticsReportRequestResponseWithDefaults instantiates a new AnalyticsReportRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsReportRequestResponseWithDefaults() *AnalyticsReportRequestResponse {
	this := AnalyticsReportRequestResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AnalyticsReportRequestResponse) GetData() AnalyticsReportRequest {
	if o == nil {
		var ret AnalyticsReportRequest
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestResponse) GetDataOk() (*AnalyticsReportRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AnalyticsReportRequestResponse) SetData(v AnalyticsReportRequest) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AnalyticsReportRequestResponse) GetIncluded() []AnalyticsReport {
	if o == nil || IsNil(o.Included) {
		var ret []AnalyticsReport
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestResponse) GetIncludedOk() ([]AnalyticsReport, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AnalyticsReportRequestResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AnalyticsReport and assigns it to the Included field.
func (o *AnalyticsReportRequestResponse) SetIncluded(v []AnalyticsReport) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AnalyticsReportRequestResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AnalyticsReportRequestResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AnalyticsReportRequestResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AnalyticsReportRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsReportRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AnalyticsReportRequestResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAnalyticsReportRequestResponse := _AnalyticsReportRequestResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalyticsReportRequestResponse)

	if err != nil {
		return err
	}

	*o = AnalyticsReportRequestResponse(varAnalyticsReportRequestResponse)

	return err
}

type NullableAnalyticsReportRequestResponse struct {
	value *AnalyticsReportRequestResponse
	isSet bool
}

func (v NullableAnalyticsReportRequestResponse) Get() *AnalyticsReportRequestResponse {
	return v.value
}

func (v *NullableAnalyticsReportRequestResponse) Set(val *AnalyticsReportRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsReportRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsReportRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsReportRequestResponse(val *AnalyticsReportRequestResponse) *NullableAnalyticsReportRequestResponse {
	return &NullableAnalyticsReportRequestResponse{value: val, isSet: true}
}

func (v NullableAnalyticsReportRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsReportRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

