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

// checks if the PagingInformationPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingInformationPaging{}

// PagingInformationPaging struct for PagingInformationPaging
type PagingInformationPaging struct {
	Total *int32 `json:"total,omitempty"`
	Limit int32 `json:"limit"`
}

type _PagingInformationPaging PagingInformationPaging

// NewPagingInformationPaging instantiates a new PagingInformationPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingInformationPaging(limit int32) *PagingInformationPaging {
	this := PagingInformationPaging{}
	this.Limit = limit
	return &this
}

// NewPagingInformationPagingWithDefaults instantiates a new PagingInformationPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingInformationPagingWithDefaults() *PagingInformationPaging {
	this := PagingInformationPaging{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PagingInformationPaging) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingInformationPaging) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PagingInformationPaging) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PagingInformationPaging) SetTotal(v int32) {
	o.Total = &v
}

// GetLimit returns the Limit field value
func (o *PagingInformationPaging) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PagingInformationPaging) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PagingInformationPaging) SetLimit(v int32) {
	o.Limit = v
}

func (o PagingInformationPaging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingInformationPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

func (o *PagingInformationPaging) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
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

	varPagingInformationPaging := _PagingInformationPaging{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPagingInformationPaging)

	if err != nil {
		return err
	}

	*o = PagingInformationPaging(varPagingInformationPaging)

	return err
}

type NullablePagingInformationPaging struct {
	value *PagingInformationPaging
	isSet bool
}

func (v NullablePagingInformationPaging) Get() *PagingInformationPaging {
	return v.value
}

func (v *NullablePagingInformationPaging) Set(val *PagingInformationPaging) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingInformationPaging) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingInformationPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingInformationPaging(val *PagingInformationPaging) *NullablePagingInformationPaging {
	return &NullablePagingInformationPaging{value: val, isSet: true}
}

func (v NullablePagingInformationPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingInformationPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


