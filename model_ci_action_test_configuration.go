/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the CiActionTestConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiActionTestConfiguration{}

// CiActionTestConfiguration struct for CiActionTestConfiguration
type CiActionTestConfiguration struct {
	Kind *string `json:"kind,omitempty"`
	TestPlanName *string `json:"testPlanName,omitempty"`
	TestDestinations []CiTestDestination `json:"testDestinations,omitempty"`
}

// NewCiActionTestConfiguration instantiates a new CiActionTestConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiActionTestConfiguration() *CiActionTestConfiguration {
	this := CiActionTestConfiguration{}
	return &this
}

// NewCiActionTestConfigurationWithDefaults instantiates a new CiActionTestConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiActionTestConfigurationWithDefaults() *CiActionTestConfiguration {
	this := CiActionTestConfiguration{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CiActionTestConfiguration) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiActionTestConfiguration) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CiActionTestConfiguration) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CiActionTestConfiguration) SetKind(v string) {
	o.Kind = &v
}

// GetTestPlanName returns the TestPlanName field value if set, zero value otherwise.
func (o *CiActionTestConfiguration) GetTestPlanName() string {
	if o == nil || IsNil(o.TestPlanName) {
		var ret string
		return ret
	}
	return *o.TestPlanName
}

// GetTestPlanNameOk returns a tuple with the TestPlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiActionTestConfiguration) GetTestPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.TestPlanName) {
		return nil, false
	}
	return o.TestPlanName, true
}

// HasTestPlanName returns a boolean if a field has been set.
func (o *CiActionTestConfiguration) HasTestPlanName() bool {
	if o != nil && !IsNil(o.TestPlanName) {
		return true
	}

	return false
}

// SetTestPlanName gets a reference to the given string and assigns it to the TestPlanName field.
func (o *CiActionTestConfiguration) SetTestPlanName(v string) {
	o.TestPlanName = &v
}

// GetTestDestinations returns the TestDestinations field value if set, zero value otherwise.
func (o *CiActionTestConfiguration) GetTestDestinations() []CiTestDestination {
	if o == nil || IsNil(o.TestDestinations) {
		var ret []CiTestDestination
		return ret
	}
	return o.TestDestinations
}

// GetTestDestinationsOk returns a tuple with the TestDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiActionTestConfiguration) GetTestDestinationsOk() ([]CiTestDestination, bool) {
	if o == nil || IsNil(o.TestDestinations) {
		return nil, false
	}
	return o.TestDestinations, true
}

// HasTestDestinations returns a boolean if a field has been set.
func (o *CiActionTestConfiguration) HasTestDestinations() bool {
	if o != nil && !IsNil(o.TestDestinations) {
		return true
	}

	return false
}

// SetTestDestinations gets a reference to the given []CiTestDestination and assigns it to the TestDestinations field.
func (o *CiActionTestConfiguration) SetTestDestinations(v []CiTestDestination) {
	o.TestDestinations = v
}

func (o CiActionTestConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiActionTestConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.TestPlanName) {
		toSerialize["testPlanName"] = o.TestPlanName
	}
	if !IsNil(o.TestDestinations) {
		toSerialize["testDestinations"] = o.TestDestinations
	}
	return toSerialize, nil
}

type NullableCiActionTestConfiguration struct {
	value *CiActionTestConfiguration
	isSet bool
}

func (v NullableCiActionTestConfiguration) Get() *CiActionTestConfiguration {
	return v.value
}

func (v *NullableCiActionTestConfiguration) Set(val *CiActionTestConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCiActionTestConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCiActionTestConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiActionTestConfiguration(val *CiActionTestConfiguration) *NullableCiActionTestConfiguration {
	return &NullableCiActionTestConfiguration{value: val, isSet: true}
}

func (v NullableCiActionTestConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiActionTestConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


