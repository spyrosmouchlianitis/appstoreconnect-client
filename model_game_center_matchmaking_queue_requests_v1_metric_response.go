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

// checks if the GameCenterMatchmakingQueueRequestsV1MetricResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueRequestsV1MetricResponse{}

// GameCenterMatchmakingQueueRequestsV1MetricResponse struct for GameCenterMatchmakingQueueRequestsV1MetricResponse
type GameCenterMatchmakingQueueRequestsV1MetricResponse struct {
	Data []GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _GameCenterMatchmakingQueueRequestsV1MetricResponse GameCenterMatchmakingQueueRequestsV1MetricResponse

// NewGameCenterMatchmakingQueueRequestsV1MetricResponse instantiates a new GameCenterMatchmakingQueueRequestsV1MetricResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueRequestsV1MetricResponse(data []GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner, links PagedDocumentLinks) *GameCenterMatchmakingQueueRequestsV1MetricResponse {
	this := GameCenterMatchmakingQueueRequestsV1MetricResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterMatchmakingQueueRequestsV1MetricResponseWithDefaults instantiates a new GameCenterMatchmakingQueueRequestsV1MetricResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueRequestsV1MetricResponseWithDefaults() *GameCenterMatchmakingQueueRequestsV1MetricResponse {
	this := GameCenterMatchmakingQueueRequestsV1MetricResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetData() []GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner {
	if o == nil {
		var ret []GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetDataOk() ([]GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) SetData(v []GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterMatchmakingQueueRequestsV1MetricResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueRequestsV1MetricResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterMatchmakingQueueRequestsV1MetricResponse := _GameCenterMatchmakingQueueRequestsV1MetricResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingQueueRequestsV1MetricResponse)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingQueueRequestsV1MetricResponse(varGameCenterMatchmakingQueueRequestsV1MetricResponse)

	return err
}

type NullableGameCenterMatchmakingQueueRequestsV1MetricResponse struct {
	value *GameCenterMatchmakingQueueRequestsV1MetricResponse
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueRequestsV1MetricResponse) Get() *GameCenterMatchmakingQueueRequestsV1MetricResponse {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueRequestsV1MetricResponse) Set(val *GameCenterMatchmakingQueueRequestsV1MetricResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueRequestsV1MetricResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueRequestsV1MetricResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueRequestsV1MetricResponse(val *GameCenterMatchmakingQueueRequestsV1MetricResponse) *NullableGameCenterMatchmakingQueueRequestsV1MetricResponse {
	return &NullableGameCenterMatchmakingQueueRequestsV1MetricResponse{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueRequestsV1MetricResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueRequestsV1MetricResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


