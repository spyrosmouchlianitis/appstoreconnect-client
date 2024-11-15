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

// checks if the GameCenterMatchmakingTestRequestInlineCreateAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTestRequestInlineCreateAttributes{}

// GameCenterMatchmakingTestRequestInlineCreateAttributes struct for GameCenterMatchmakingTestRequestInlineCreateAttributes
type GameCenterMatchmakingTestRequestInlineCreateAttributes struct {
	RequestName string `json:"requestName"`
	SecondsInQueue int32 `json:"secondsInQueue"`
	Locale *string `json:"locale,omitempty"`
	Location *Location `json:"location,omitempty"`
	MinPlayers *int32 `json:"minPlayers,omitempty"`
	MaxPlayers *int32 `json:"maxPlayers,omitempty"`
	PlayerCount *int32 `json:"playerCount,omitempty"`
	BundleId string `json:"bundleId"`
	Platform Platform `json:"platform"`
	AppVersion string `json:"appVersion"`
}

type _GameCenterMatchmakingTestRequestInlineCreateAttributes GameCenterMatchmakingTestRequestInlineCreateAttributes

// NewGameCenterMatchmakingTestRequestInlineCreateAttributes instantiates a new GameCenterMatchmakingTestRequestInlineCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTestRequestInlineCreateAttributes(requestName string, secondsInQueue int32, bundleId string, platform Platform, appVersion string) *GameCenterMatchmakingTestRequestInlineCreateAttributes {
	this := GameCenterMatchmakingTestRequestInlineCreateAttributes{}
	this.RequestName = requestName
	this.SecondsInQueue = secondsInQueue
	this.BundleId = bundleId
	this.Platform = platform
	this.AppVersion = appVersion
	return &this
}

// NewGameCenterMatchmakingTestRequestInlineCreateAttributesWithDefaults instantiates a new GameCenterMatchmakingTestRequestInlineCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTestRequestInlineCreateAttributesWithDefaults() *GameCenterMatchmakingTestRequestInlineCreateAttributes {
	this := GameCenterMatchmakingTestRequestInlineCreateAttributes{}
	return &this
}

// GetRequestName returns the RequestName field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetRequestName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestName
}

// GetRequestNameOk returns a tuple with the RequestName field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetRequestNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestName, true
}

// SetRequestName sets field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetRequestName(v string) {
	o.RequestName = v
}

// GetSecondsInQueue returns the SecondsInQueue field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetSecondsInQueue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecondsInQueue
}

// GetSecondsInQueueOk returns a tuple with the SecondsInQueue field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetSecondsInQueueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondsInQueue, true
}

// SetSecondsInQueue sets field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetSecondsInQueue(v int32) {
	o.SecondsInQueue = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetLocale(v string) {
	o.Locale = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetLocation(v Location) {
	o.Location = &v
}

// GetMinPlayers returns the MinPlayers field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetMinPlayers() int32 {
	if o == nil || IsNil(o.MinPlayers) {
		var ret int32
		return ret
	}
	return *o.MinPlayers
}

// GetMinPlayersOk returns a tuple with the MinPlayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetMinPlayersOk() (*int32, bool) {
	if o == nil || IsNil(o.MinPlayers) {
		return nil, false
	}
	return o.MinPlayers, true
}

// HasMinPlayers returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasMinPlayers() bool {
	if o != nil && !IsNil(o.MinPlayers) {
		return true
	}

	return false
}

// SetMinPlayers gets a reference to the given int32 and assigns it to the MinPlayers field.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetMinPlayers(v int32) {
	o.MinPlayers = &v
}

// GetMaxPlayers returns the MaxPlayers field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetMaxPlayers() int32 {
	if o == nil || IsNil(o.MaxPlayers) {
		var ret int32
		return ret
	}
	return *o.MaxPlayers
}

// GetMaxPlayersOk returns a tuple with the MaxPlayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetMaxPlayersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPlayers) {
		return nil, false
	}
	return o.MaxPlayers, true
}

// HasMaxPlayers returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasMaxPlayers() bool {
	if o != nil && !IsNil(o.MaxPlayers) {
		return true
	}

	return false
}

// SetMaxPlayers gets a reference to the given int32 and assigns it to the MaxPlayers field.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetMaxPlayers(v int32) {
	o.MaxPlayers = &v
}

// GetPlayerCount returns the PlayerCount field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetPlayerCount() int32 {
	if o == nil || IsNil(o.PlayerCount) {
		var ret int32
		return ret
	}
	return *o.PlayerCount
}

// GetPlayerCountOk returns a tuple with the PlayerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetPlayerCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PlayerCount) {
		return nil, false
	}
	return o.PlayerCount, true
}

// HasPlayerCount returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasPlayerCount() bool {
	if o != nil && !IsNil(o.PlayerCount) {
		return true
	}

	return false
}

// SetPlayerCount gets a reference to the given int32 and assigns it to the PlayerCount field.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetPlayerCount(v int32) {
	o.PlayerCount = &v
}

// GetBundleId returns the BundleId field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetBundleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetBundleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleId, true
}

// SetBundleId sets field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetBundleId(v string) {
	o.BundleId = v
}

// GetPlatform returns the Platform field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetPlatform() Platform {
	if o == nil {
		var ret Platform
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetPlatformOk() (*Platform, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetPlatform(v Platform) {
	o.Platform = v
}

// GetAppVersion returns the AppVersion field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetAppVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppVersion, true
}

// SetAppVersion sets field value
func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetAppVersion(v string) {
	o.AppVersion = v
}

func (o GameCenterMatchmakingTestRequestInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTestRequestInlineCreateAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestName"] = o.RequestName
	toSerialize["secondsInQueue"] = o.SecondsInQueue
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.MinPlayers) {
		toSerialize["minPlayers"] = o.MinPlayers
	}
	if !IsNil(o.MaxPlayers) {
		toSerialize["maxPlayers"] = o.MaxPlayers
	}
	if !IsNil(o.PlayerCount) {
		toSerialize["playerCount"] = o.PlayerCount
	}
	toSerialize["bundleId"] = o.BundleId
	toSerialize["platform"] = o.Platform
	toSerialize["appVersion"] = o.AppVersion
	return toSerialize, nil
}

func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requestName",
		"secondsInQueue",
		"bundleId",
		"platform",
		"appVersion",
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

	varGameCenterMatchmakingTestRequestInlineCreateAttributes := _GameCenterMatchmakingTestRequestInlineCreateAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingTestRequestInlineCreateAttributes)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingTestRequestInlineCreateAttributes(varGameCenterMatchmakingTestRequestInlineCreateAttributes)

	return err
}

type NullableGameCenterMatchmakingTestRequestInlineCreateAttributes struct {
	value *GameCenterMatchmakingTestRequestInlineCreateAttributes
	isSet bool
}

func (v NullableGameCenterMatchmakingTestRequestInlineCreateAttributes) Get() *GameCenterMatchmakingTestRequestInlineCreateAttributes {
	return v.value
}

func (v *NullableGameCenterMatchmakingTestRequestInlineCreateAttributes) Set(val *GameCenterMatchmakingTestRequestInlineCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTestRequestInlineCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTestRequestInlineCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTestRequestInlineCreateAttributes(val *GameCenterMatchmakingTestRequestInlineCreateAttributes) *NullableGameCenterMatchmakingTestRequestInlineCreateAttributes {
	return &NullableGameCenterMatchmakingTestRequestInlineCreateAttributes{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTestRequestInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTestRequestInlineCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


