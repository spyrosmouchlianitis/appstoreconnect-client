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

// checks if the UploadOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadOperation{}

// UploadOperation struct for UploadOperation
type UploadOperation struct {
	Method *string `json:"method,omitempty"`
	Url *string `json:"url,omitempty"`
	Length *int32 `json:"length,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	RequestHeaders []HttpHeader `json:"requestHeaders,omitempty"`
}

// NewUploadOperation instantiates a new UploadOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadOperation() *UploadOperation {
	this := UploadOperation{}
	return &this
}

// NewUploadOperationWithDefaults instantiates a new UploadOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadOperationWithDefaults() *UploadOperation {
	this := UploadOperation{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *UploadOperation) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadOperation) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *UploadOperation) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *UploadOperation) SetMethod(v string) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UploadOperation) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadOperation) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UploadOperation) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UploadOperation) SetUrl(v string) {
	o.Url = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *UploadOperation) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadOperation) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *UploadOperation) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *UploadOperation) SetLength(v int32) {
	o.Length = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *UploadOperation) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadOperation) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *UploadOperation) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *UploadOperation) SetOffset(v int32) {
	o.Offset = &v
}

// GetRequestHeaders returns the RequestHeaders field value if set, zero value otherwise.
func (o *UploadOperation) GetRequestHeaders() []HttpHeader {
	if o == nil || IsNil(o.RequestHeaders) {
		var ret []HttpHeader
		return ret
	}
	return o.RequestHeaders
}

// GetRequestHeadersOk returns a tuple with the RequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadOperation) GetRequestHeadersOk() ([]HttpHeader, bool) {
	if o == nil || IsNil(o.RequestHeaders) {
		return nil, false
	}
	return o.RequestHeaders, true
}

// HasRequestHeaders returns a boolean if a field has been set.
func (o *UploadOperation) HasRequestHeaders() bool {
	if o != nil && !IsNil(o.RequestHeaders) {
		return true
	}

	return false
}

// SetRequestHeaders gets a reference to the given []HttpHeader and assigns it to the RequestHeaders field.
func (o *UploadOperation) SetRequestHeaders(v []HttpHeader) {
	o.RequestHeaders = v
}

func (o UploadOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.RequestHeaders) {
		toSerialize["requestHeaders"] = o.RequestHeaders
	}
	return toSerialize, nil
}

type NullableUploadOperation struct {
	value *UploadOperation
	isSet bool
}

func (v NullableUploadOperation) Get() *UploadOperation {
	return v.value
}

func (v *NullableUploadOperation) Set(val *UploadOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadOperation(val *UploadOperation) *NullableUploadOperation {
	return &NullableUploadOperation{value: val, isSet: true}
}

func (v NullableUploadOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


