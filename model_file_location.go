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

// checks if the FileLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileLocation{}

// FileLocation struct for FileLocation
type FileLocation struct {
	Path *string `json:"path,omitempty"`
	LineNumber *int32 `json:"lineNumber,omitempty"`
}

// NewFileLocation instantiates a new FileLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLocation() *FileLocation {
	this := FileLocation{}
	return &this
}

// NewFileLocationWithDefaults instantiates a new FileLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLocationWithDefaults() *FileLocation {
	this := FileLocation{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FileLocation) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLocation) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FileLocation) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FileLocation) SetPath(v string) {
	o.Path = &v
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *FileLocation) GetLineNumber() int32 {
	if o == nil || IsNil(o.LineNumber) {
		var ret int32
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLocation) GetLineNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.LineNumber) {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *FileLocation) HasLineNumber() bool {
	if o != nil && !IsNil(o.LineNumber) {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given int32 and assigns it to the LineNumber field.
func (o *FileLocation) SetLineNumber(v int32) {
	o.LineNumber = &v
}

func (o FileLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.LineNumber) {
		toSerialize["lineNumber"] = o.LineNumber
	}
	return toSerialize, nil
}

type NullableFileLocation struct {
	value *FileLocation
	isSet bool
}

func (v NullableFileLocation) Get() *FileLocation {
	return v.value
}

func (v *NullableFileLocation) Set(val *FileLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLocation(val *FileLocation) *NullableFileLocation {
	return &NullableFileLocation{value: val, isSet: true}
}

func (v NullableFileLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


