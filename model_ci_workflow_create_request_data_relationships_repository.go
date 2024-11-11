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

// checks if the CiWorkflowCreateRequestDataRelationshipsRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowCreateRequestDataRelationshipsRepository{}

// CiWorkflowCreateRequestDataRelationshipsRepository struct for CiWorkflowCreateRequestDataRelationshipsRepository
type CiWorkflowCreateRequestDataRelationshipsRepository struct {
	Data CiProductRelationshipsPrimaryRepositoriesDataInner `json:"data"`
}

type _CiWorkflowCreateRequestDataRelationshipsRepository CiWorkflowCreateRequestDataRelationshipsRepository

// NewCiWorkflowCreateRequestDataRelationshipsRepository instantiates a new CiWorkflowCreateRequestDataRelationshipsRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowCreateRequestDataRelationshipsRepository(data CiProductRelationshipsPrimaryRepositoriesDataInner) *CiWorkflowCreateRequestDataRelationshipsRepository {
	this := CiWorkflowCreateRequestDataRelationshipsRepository{}
	this.Data = data
	return &this
}

// NewCiWorkflowCreateRequestDataRelationshipsRepositoryWithDefaults instantiates a new CiWorkflowCreateRequestDataRelationshipsRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowCreateRequestDataRelationshipsRepositoryWithDefaults() *CiWorkflowCreateRequestDataRelationshipsRepository {
	this := CiWorkflowCreateRequestDataRelationshipsRepository{}
	return &this
}

// GetData returns the Data field value
func (o *CiWorkflowCreateRequestDataRelationshipsRepository) GetData() CiProductRelationshipsPrimaryRepositoriesDataInner {
	if o == nil {
		var ret CiProductRelationshipsPrimaryRepositoriesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataRelationshipsRepository) GetDataOk() (*CiProductRelationshipsPrimaryRepositoriesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiWorkflowCreateRequestDataRelationshipsRepository) SetData(v CiProductRelationshipsPrimaryRepositoriesDataInner) {
	o.Data = v
}

func (o CiWorkflowCreateRequestDataRelationshipsRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowCreateRequestDataRelationshipsRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CiWorkflowCreateRequestDataRelationshipsRepository) UnmarshalJSON(data []byte) (err error) {
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

	varCiWorkflowCreateRequestDataRelationshipsRepository := _CiWorkflowCreateRequestDataRelationshipsRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiWorkflowCreateRequestDataRelationshipsRepository)

	if err != nil {
		return err
	}

	*o = CiWorkflowCreateRequestDataRelationshipsRepository(varCiWorkflowCreateRequestDataRelationshipsRepository)

	return err
}

type NullableCiWorkflowCreateRequestDataRelationshipsRepository struct {
	value *CiWorkflowCreateRequestDataRelationshipsRepository
	isSet bool
}

func (v NullableCiWorkflowCreateRequestDataRelationshipsRepository) Get() *CiWorkflowCreateRequestDataRelationshipsRepository {
	return v.value
}

func (v *NullableCiWorkflowCreateRequestDataRelationshipsRepository) Set(val *CiWorkflowCreateRequestDataRelationshipsRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowCreateRequestDataRelationshipsRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowCreateRequestDataRelationshipsRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowCreateRequestDataRelationshipsRepository(val *CiWorkflowCreateRequestDataRelationshipsRepository) *NullableCiWorkflowCreateRequestDataRelationshipsRepository {
	return &NullableCiWorkflowCreateRequestDataRelationshipsRepository{value: val, isSet: true}
}

func (v NullableCiWorkflowCreateRequestDataRelationshipsRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowCreateRequestDataRelationshipsRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


