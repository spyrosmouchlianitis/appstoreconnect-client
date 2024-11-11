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

// checks if the ScmRepositoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmRepositoryResponse{}

// ScmRepositoryResponse struct for ScmRepositoryResponse
type ScmRepositoryResponse struct {
	Data ScmRepository `json:"data"`
	Included []ScmRepositoriesResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _ScmRepositoryResponse ScmRepositoryResponse

// NewScmRepositoryResponse instantiates a new ScmRepositoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmRepositoryResponse(data ScmRepository, links DocumentLinks) *ScmRepositoryResponse {
	this := ScmRepositoryResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewScmRepositoryResponseWithDefaults instantiates a new ScmRepositoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmRepositoryResponseWithDefaults() *ScmRepositoryResponse {
	this := ScmRepositoryResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ScmRepositoryResponse) GetData() ScmRepository {
	if o == nil {
		var ret ScmRepository
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ScmRepositoryResponse) GetDataOk() (*ScmRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ScmRepositoryResponse) SetData(v ScmRepository) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ScmRepositoryResponse) GetIncluded() []ScmRepositoriesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []ScmRepositoriesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryResponse) GetIncludedOk() ([]ScmRepositoriesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ScmRepositoryResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []ScmRepositoriesResponseIncludedInner and assigns it to the Included field.
func (o *ScmRepositoryResponse) SetIncluded(v []ScmRepositoriesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *ScmRepositoryResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ScmRepositoryResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ScmRepositoryResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o ScmRepositoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmRepositoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *ScmRepositoryResponse) UnmarshalJSON(data []byte) (err error) {
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

	varScmRepositoryResponse := _ScmRepositoryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScmRepositoryResponse)

	if err != nil {
		return err
	}

	*o = ScmRepositoryResponse(varScmRepositoryResponse)

	return err
}

type NullableScmRepositoryResponse struct {
	value *ScmRepositoryResponse
	isSet bool
}

func (v NullableScmRepositoryResponse) Get() *ScmRepositoryResponse {
	return v.value
}

func (v *NullableScmRepositoryResponse) Set(val *ScmRepositoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScmRepositoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScmRepositoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmRepositoryResponse(val *ScmRepositoryResponse) *NullableScmRepositoryResponse {
	return &NullableScmRepositoryResponse{value: val, isSet: true}
}

func (v NullableScmRepositoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmRepositoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


