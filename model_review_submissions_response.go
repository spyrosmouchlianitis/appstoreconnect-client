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

// checks if the ReviewSubmissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionsResponse{}

// ReviewSubmissionsResponse struct for ReviewSubmissionsResponse
type ReviewSubmissionsResponse struct {
	Data []ReviewSubmission `json:"data"`
	Included []ReviewSubmissionsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _ReviewSubmissionsResponse ReviewSubmissionsResponse

// NewReviewSubmissionsResponse instantiates a new ReviewSubmissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionsResponse(data []ReviewSubmission, links PagedDocumentLinks) *ReviewSubmissionsResponse {
	this := ReviewSubmissionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewReviewSubmissionsResponseWithDefaults instantiates a new ReviewSubmissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionsResponseWithDefaults() *ReviewSubmissionsResponse {
	this := ReviewSubmissionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ReviewSubmissionsResponse) GetData() []ReviewSubmission {
	if o == nil {
		var ret []ReviewSubmission
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionsResponse) GetDataOk() ([]ReviewSubmission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ReviewSubmissionsResponse) SetData(v []ReviewSubmission) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ReviewSubmissionsResponse) GetIncluded() []ReviewSubmissionsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []ReviewSubmissionsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionsResponse) GetIncludedOk() ([]ReviewSubmissionsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ReviewSubmissionsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []ReviewSubmissionsResponseIncludedInner and assigns it to the Included field.
func (o *ReviewSubmissionsResponse) SetIncluded(v []ReviewSubmissionsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *ReviewSubmissionsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ReviewSubmissionsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReviewSubmissionsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReviewSubmissionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *ReviewSubmissionsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o ReviewSubmissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ReviewSubmissionsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varReviewSubmissionsResponse := _ReviewSubmissionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReviewSubmissionsResponse)

	if err != nil {
		return err
	}

	*o = ReviewSubmissionsResponse(varReviewSubmissionsResponse)

	return err
}

type NullableReviewSubmissionsResponse struct {
	value *ReviewSubmissionsResponse
	isSet bool
}

func (v NullableReviewSubmissionsResponse) Get() *ReviewSubmissionsResponse {
	return v.value
}

func (v *NullableReviewSubmissionsResponse) Set(val *ReviewSubmissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionsResponse(val *ReviewSubmissionsResponse) *NullableReviewSubmissionsResponse {
	return &NullableReviewSubmissionsResponse{value: val, isSet: true}
}

func (v NullableReviewSubmissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


