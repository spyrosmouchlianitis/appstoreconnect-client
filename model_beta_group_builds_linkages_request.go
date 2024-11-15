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

// checks if the BetaGroupBuildsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupBuildsLinkagesRequest{}

// BetaGroupBuildsLinkagesRequest struct for BetaGroupBuildsLinkagesRequest
type BetaGroupBuildsLinkagesRequest struct {
	Data []AppEncryptionDeclarationRelationshipsBuildsDataInner `json:"data"`
}

type _BetaGroupBuildsLinkagesRequest BetaGroupBuildsLinkagesRequest

// NewBetaGroupBuildsLinkagesRequest instantiates a new BetaGroupBuildsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupBuildsLinkagesRequest(data []AppEncryptionDeclarationRelationshipsBuildsDataInner) *BetaGroupBuildsLinkagesRequest {
	this := BetaGroupBuildsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewBetaGroupBuildsLinkagesRequestWithDefaults instantiates a new BetaGroupBuildsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupBuildsLinkagesRequestWithDefaults() *BetaGroupBuildsLinkagesRequest {
	this := BetaGroupBuildsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupBuildsLinkagesRequest) GetData() []AppEncryptionDeclarationRelationshipsBuildsDataInner {
	if o == nil {
		var ret []AppEncryptionDeclarationRelationshipsBuildsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupBuildsLinkagesRequest) GetDataOk() ([]AppEncryptionDeclarationRelationshipsBuildsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaGroupBuildsLinkagesRequest) SetData(v []AppEncryptionDeclarationRelationshipsBuildsDataInner) {
	o.Data = v
}

func (o BetaGroupBuildsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupBuildsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BetaGroupBuildsLinkagesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varBetaGroupBuildsLinkagesRequest := _BetaGroupBuildsLinkagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaGroupBuildsLinkagesRequest)

	if err != nil {
		return err
	}

	*o = BetaGroupBuildsLinkagesRequest(varBetaGroupBuildsLinkagesRequest)

	return err
}

type NullableBetaGroupBuildsLinkagesRequest struct {
	value *BetaGroupBuildsLinkagesRequest
	isSet bool
}

func (v NullableBetaGroupBuildsLinkagesRequest) Get() *BetaGroupBuildsLinkagesRequest {
	return v.value
}

func (v *NullableBetaGroupBuildsLinkagesRequest) Set(val *BetaGroupBuildsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupBuildsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupBuildsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupBuildsLinkagesRequest(val *BetaGroupBuildsLinkagesRequest) *NullableBetaGroupBuildsLinkagesRequest {
	return &NullableBetaGroupBuildsLinkagesRequest{value: val, isSet: true}
}

func (v NullableBetaGroupBuildsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupBuildsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


