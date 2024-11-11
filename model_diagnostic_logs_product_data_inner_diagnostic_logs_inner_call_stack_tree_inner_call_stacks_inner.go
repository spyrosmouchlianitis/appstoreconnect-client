/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner{}

// DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner struct for DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner
type DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner struct {
	CallStackRootFrames []DiagnosticLogCallStackNode `json:"callStackRootFrames,omitempty"`
}

// NewDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner instantiates a new DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner() *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner {
	this := DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner{}
	return &this
}

// NewDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInnerWithDefaults instantiates a new DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInnerWithDefaults() *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner {
	this := DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner{}
	return &this
}

// GetCallStackRootFrames returns the CallStackRootFrames field value if set, zero value otherwise.
func (o *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) GetCallStackRootFrames() []DiagnosticLogCallStackNode {
	if o == nil || IsNil(o.CallStackRootFrames) {
		var ret []DiagnosticLogCallStackNode
		return ret
	}
	return o.CallStackRootFrames
}

// GetCallStackRootFramesOk returns a tuple with the CallStackRootFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) GetCallStackRootFramesOk() ([]DiagnosticLogCallStackNode, bool) {
	if o == nil || IsNil(o.CallStackRootFrames) {
		return nil, false
	}
	return o.CallStackRootFrames, true
}

// HasCallStackRootFrames returns a boolean if a field has been set.
func (o *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) HasCallStackRootFrames() bool {
	if o != nil && !IsNil(o.CallStackRootFrames) {
		return true
	}

	return false
}

// SetCallStackRootFrames gets a reference to the given []DiagnosticLogCallStackNode and assigns it to the CallStackRootFrames field.
func (o *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) SetCallStackRootFrames(v []DiagnosticLogCallStackNode) {
	o.CallStackRootFrames = v
}

func (o DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallStackRootFrames) {
		toSerialize["callStackRootFrames"] = o.CallStackRootFrames
	}
	return toSerialize, nil
}

type NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner struct {
	value *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner
	isSet bool
}

func (v NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) Get() *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner {
	return v.value
}

func (v *NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) Set(val *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner(val *DiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) *NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner {
	return &NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner{value: val, isSet: true}
}

func (v NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticLogsProductDataInnerDiagnosticLogsInnerCallStackTreeInnerCallStacksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

