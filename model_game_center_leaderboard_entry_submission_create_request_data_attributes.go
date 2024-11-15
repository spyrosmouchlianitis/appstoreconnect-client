/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes{}

// GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes struct for GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes
type GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes struct {
	BundleId string `json:"bundleId"`
	ChallengeIds []string `json:"challengeIds,omitempty"`
	Context *float64 `json:"context,omitempty"`
	ScopedPlayerId string `json:"scopedPlayerId"`
	Score float64 `json:"score"`
	SubmittedDate *time.Time `json:"submittedDate,omitempty"`
	VendorIdentifier string `json:"vendorIdentifier"`
}

type _GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes

// NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes instantiates a new GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes(bundleId string, scopedPlayerId string, score float64, vendorIdentifier string) *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes {
	this := GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes{}
	this.BundleId = bundleId
	this.ScopedPlayerId = scopedPlayerId
	this.Score = score
	this.VendorIdentifier = vendorIdentifier
	return &this
}

// NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributesWithDefaults() *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes {
	this := GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes{}
	return &this
}

// GetBundleId returns the BundleId field value
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetBundleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetBundleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleId, true
}

// SetBundleId sets field value
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetBundleId(v string) {
	o.BundleId = v
}

// GetChallengeIds returns the ChallengeIds field value if set, zero value otherwise.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetChallengeIds() []string {
	if o == nil || IsNil(o.ChallengeIds) {
		var ret []string
		return ret
	}
	return o.ChallengeIds
}

// GetChallengeIdsOk returns a tuple with the ChallengeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetChallengeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ChallengeIds) {
		return nil, false
	}
	return o.ChallengeIds, true
}

// HasChallengeIds returns a boolean if a field has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) HasChallengeIds() bool {
	if o != nil && !IsNil(o.ChallengeIds) {
		return true
	}

	return false
}

// SetChallengeIds gets a reference to the given []string and assigns it to the ChallengeIds field.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetChallengeIds(v []string) {
	o.ChallengeIds = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetContext() float64 {
	if o == nil || IsNil(o.Context) {
		var ret float64
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetContextOk() (*float64, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given float64 and assigns it to the Context field.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetContext(v float64) {
	o.Context = &v
}

// GetScopedPlayerId returns the ScopedPlayerId field value
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetScopedPlayerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScopedPlayerId
}

// GetScopedPlayerIdOk returns a tuple with the ScopedPlayerId field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetScopedPlayerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopedPlayerId, true
}

// SetScopedPlayerId sets field value
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetScopedPlayerId(v string) {
	o.ScopedPlayerId = v
}

// GetScore returns the Score field value
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetScore(v float64) {
	o.Score = v
}

// GetSubmittedDate returns the SubmittedDate field value if set, zero value otherwise.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetSubmittedDate() time.Time {
	if o == nil || IsNil(o.SubmittedDate) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedDate
}

// GetSubmittedDateOk returns a tuple with the SubmittedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetSubmittedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedDate) {
		return nil, false
	}
	return o.SubmittedDate, true
}

// HasSubmittedDate returns a boolean if a field has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) HasSubmittedDate() bool {
	if o != nil && !IsNil(o.SubmittedDate) {
		return true
	}

	return false
}

// SetSubmittedDate gets a reference to the given time.Time and assigns it to the SubmittedDate field.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetSubmittedDate(v time.Time) {
	o.SubmittedDate = &v
}

// GetVendorIdentifier returns the VendorIdentifier field value
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetVendorIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorIdentifier
}

// GetVendorIdentifierOk returns a tuple with the VendorIdentifier field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetVendorIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorIdentifier, true
}

// SetVendorIdentifier sets field value
func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetVendorIdentifier(v string) {
	o.VendorIdentifier = v
}

func (o GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bundleId"] = o.BundleId
	if !IsNil(o.ChallengeIds) {
		toSerialize["challengeIds"] = o.ChallengeIds
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	toSerialize["scopedPlayerId"] = o.ScopedPlayerId
	toSerialize["score"] = o.Score
	if !IsNil(o.SubmittedDate) {
		toSerialize["submittedDate"] = o.SubmittedDate
	}
	toSerialize["vendorIdentifier"] = o.VendorIdentifier
	return toSerialize, nil
}

func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bundleId",
		"scopedPlayerId",
		"score",
		"vendorIdentifier",
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

	varGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes := _GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes(varGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes)

	return err
}

type NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes struct {
	value *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) Get() *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) Set(val *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes(val *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) *NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes {
	return &NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


