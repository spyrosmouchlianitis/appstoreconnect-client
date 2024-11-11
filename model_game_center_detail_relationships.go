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

// checks if the GameCenterDetailRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailRelationships{}

// GameCenterDetailRelationships struct for GameCenterDetailRelationships
type GameCenterDetailRelationships struct {
	App *AlternativeDistributionKeyCreateRequestDataRelationshipsApp `json:"app,omitempty"`
	GameCenterAppVersions *GameCenterAppVersionRelationshipsCompatibilityVersions `json:"gameCenterAppVersions,omitempty"`
	GameCenterGroup *GameCenterDetailRelationshipsGameCenterGroup `json:"gameCenterGroup,omitempty"`
	GameCenterLeaderboards *GameCenterDetailRelationshipsGameCenterLeaderboards `json:"gameCenterLeaderboards,omitempty"`
	GameCenterLeaderboardSets *GameCenterDetailRelationshipsGameCenterLeaderboardSets `json:"gameCenterLeaderboardSets,omitempty"`
	GameCenterAchievements *GameCenterDetailRelationshipsGameCenterAchievements `json:"gameCenterAchievements,omitempty"`
	DefaultLeaderboard *GameCenterDetailRelationshipsDefaultLeaderboard `json:"defaultLeaderboard,omitempty"`
	DefaultGroupLeaderboard *GameCenterDetailRelationshipsDefaultLeaderboard `json:"defaultGroupLeaderboard,omitempty"`
	AchievementReleases *GameCenterAchievementRelationshipsReleases `json:"achievementReleases,omitempty"`
	LeaderboardReleases *GameCenterDetailRelationshipsLeaderboardReleases `json:"leaderboardReleases,omitempty"`
	LeaderboardSetReleases *GameCenterDetailRelationshipsLeaderboardSetReleases `json:"leaderboardSetReleases,omitempty"`
}

// NewGameCenterDetailRelationships instantiates a new GameCenterDetailRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailRelationships() *GameCenterDetailRelationships {
	this := GameCenterDetailRelationships{}
	return &this
}

// NewGameCenterDetailRelationshipsWithDefaults instantiates a new GameCenterDetailRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailRelationshipsWithDefaults() *GameCenterDetailRelationships {
	this := GameCenterDetailRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetApp() AlternativeDistributionKeyCreateRequestDataRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AlternativeDistributionKeyCreateRequestDataRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetAppOk() (*AlternativeDistributionKeyCreateRequestDataRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AlternativeDistributionKeyCreateRequestDataRelationshipsApp and assigns it to the App field.
func (o *GameCenterDetailRelationships) SetApp(v AlternativeDistributionKeyCreateRequestDataRelationshipsApp) {
	o.App = &v
}

// GetGameCenterAppVersions returns the GameCenterAppVersions field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetGameCenterAppVersions() GameCenterAppVersionRelationshipsCompatibilityVersions {
	if o == nil || IsNil(o.GameCenterAppVersions) {
		var ret GameCenterAppVersionRelationshipsCompatibilityVersions
		return ret
	}
	return *o.GameCenterAppVersions
}

// GetGameCenterAppVersionsOk returns a tuple with the GameCenterAppVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetGameCenterAppVersionsOk() (*GameCenterAppVersionRelationshipsCompatibilityVersions, bool) {
	if o == nil || IsNil(o.GameCenterAppVersions) {
		return nil, false
	}
	return o.GameCenterAppVersions, true
}

// HasGameCenterAppVersions returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasGameCenterAppVersions() bool {
	if o != nil && !IsNil(o.GameCenterAppVersions) {
		return true
	}

	return false
}

// SetGameCenterAppVersions gets a reference to the given GameCenterAppVersionRelationshipsCompatibilityVersions and assigns it to the GameCenterAppVersions field.
func (o *GameCenterDetailRelationships) SetGameCenterAppVersions(v GameCenterAppVersionRelationshipsCompatibilityVersions) {
	o.GameCenterAppVersions = &v
}

// GetGameCenterGroup returns the GameCenterGroup field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetGameCenterGroup() GameCenterDetailRelationshipsGameCenterGroup {
	if o == nil || IsNil(o.GameCenterGroup) {
		var ret GameCenterDetailRelationshipsGameCenterGroup
		return ret
	}
	return *o.GameCenterGroup
}

// GetGameCenterGroupOk returns a tuple with the GameCenterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetGameCenterGroupOk() (*GameCenterDetailRelationshipsGameCenterGroup, bool) {
	if o == nil || IsNil(o.GameCenterGroup) {
		return nil, false
	}
	return o.GameCenterGroup, true
}

// HasGameCenterGroup returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasGameCenterGroup() bool {
	if o != nil && !IsNil(o.GameCenterGroup) {
		return true
	}

	return false
}

// SetGameCenterGroup gets a reference to the given GameCenterDetailRelationshipsGameCenterGroup and assigns it to the GameCenterGroup field.
func (o *GameCenterDetailRelationships) SetGameCenterGroup(v GameCenterDetailRelationshipsGameCenterGroup) {
	o.GameCenterGroup = &v
}

// GetGameCenterLeaderboards returns the GameCenterLeaderboards field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetGameCenterLeaderboards() GameCenterDetailRelationshipsGameCenterLeaderboards {
	if o == nil || IsNil(o.GameCenterLeaderboards) {
		var ret GameCenterDetailRelationshipsGameCenterLeaderboards
		return ret
	}
	return *o.GameCenterLeaderboards
}

// GetGameCenterLeaderboardsOk returns a tuple with the GameCenterLeaderboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetGameCenterLeaderboardsOk() (*GameCenterDetailRelationshipsGameCenterLeaderboards, bool) {
	if o == nil || IsNil(o.GameCenterLeaderboards) {
		return nil, false
	}
	return o.GameCenterLeaderboards, true
}

// HasGameCenterLeaderboards returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasGameCenterLeaderboards() bool {
	if o != nil && !IsNil(o.GameCenterLeaderboards) {
		return true
	}

	return false
}

// SetGameCenterLeaderboards gets a reference to the given GameCenterDetailRelationshipsGameCenterLeaderboards and assigns it to the GameCenterLeaderboards field.
func (o *GameCenterDetailRelationships) SetGameCenterLeaderboards(v GameCenterDetailRelationshipsGameCenterLeaderboards) {
	o.GameCenterLeaderboards = &v
}

// GetGameCenterLeaderboardSets returns the GameCenterLeaderboardSets field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetGameCenterLeaderboardSets() GameCenterDetailRelationshipsGameCenterLeaderboardSets {
	if o == nil || IsNil(o.GameCenterLeaderboardSets) {
		var ret GameCenterDetailRelationshipsGameCenterLeaderboardSets
		return ret
	}
	return *o.GameCenterLeaderboardSets
}

// GetGameCenterLeaderboardSetsOk returns a tuple with the GameCenterLeaderboardSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetGameCenterLeaderboardSetsOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardSets, bool) {
	if o == nil || IsNil(o.GameCenterLeaderboardSets) {
		return nil, false
	}
	return o.GameCenterLeaderboardSets, true
}

// HasGameCenterLeaderboardSets returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasGameCenterLeaderboardSets() bool {
	if o != nil && !IsNil(o.GameCenterLeaderboardSets) {
		return true
	}

	return false
}

// SetGameCenterLeaderboardSets gets a reference to the given GameCenterDetailRelationshipsGameCenterLeaderboardSets and assigns it to the GameCenterLeaderboardSets field.
func (o *GameCenterDetailRelationships) SetGameCenterLeaderboardSets(v GameCenterDetailRelationshipsGameCenterLeaderboardSets) {
	o.GameCenterLeaderboardSets = &v
}

// GetGameCenterAchievements returns the GameCenterAchievements field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetGameCenterAchievements() GameCenterDetailRelationshipsGameCenterAchievements {
	if o == nil || IsNil(o.GameCenterAchievements) {
		var ret GameCenterDetailRelationshipsGameCenterAchievements
		return ret
	}
	return *o.GameCenterAchievements
}

// GetGameCenterAchievementsOk returns a tuple with the GameCenterAchievements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetGameCenterAchievementsOk() (*GameCenterDetailRelationshipsGameCenterAchievements, bool) {
	if o == nil || IsNil(o.GameCenterAchievements) {
		return nil, false
	}
	return o.GameCenterAchievements, true
}

// HasGameCenterAchievements returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasGameCenterAchievements() bool {
	if o != nil && !IsNil(o.GameCenterAchievements) {
		return true
	}

	return false
}

// SetGameCenterAchievements gets a reference to the given GameCenterDetailRelationshipsGameCenterAchievements and assigns it to the GameCenterAchievements field.
func (o *GameCenterDetailRelationships) SetGameCenterAchievements(v GameCenterDetailRelationshipsGameCenterAchievements) {
	o.GameCenterAchievements = &v
}

// GetDefaultLeaderboard returns the DefaultLeaderboard field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetDefaultLeaderboard() GameCenterDetailRelationshipsDefaultLeaderboard {
	if o == nil || IsNil(o.DefaultLeaderboard) {
		var ret GameCenterDetailRelationshipsDefaultLeaderboard
		return ret
	}
	return *o.DefaultLeaderboard
}

// GetDefaultLeaderboardOk returns a tuple with the DefaultLeaderboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetDefaultLeaderboardOk() (*GameCenterDetailRelationshipsDefaultLeaderboard, bool) {
	if o == nil || IsNil(o.DefaultLeaderboard) {
		return nil, false
	}
	return o.DefaultLeaderboard, true
}

// HasDefaultLeaderboard returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasDefaultLeaderboard() bool {
	if o != nil && !IsNil(o.DefaultLeaderboard) {
		return true
	}

	return false
}

// SetDefaultLeaderboard gets a reference to the given GameCenterDetailRelationshipsDefaultLeaderboard and assigns it to the DefaultLeaderboard field.
func (o *GameCenterDetailRelationships) SetDefaultLeaderboard(v GameCenterDetailRelationshipsDefaultLeaderboard) {
	o.DefaultLeaderboard = &v
}

// GetDefaultGroupLeaderboard returns the DefaultGroupLeaderboard field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetDefaultGroupLeaderboard() GameCenterDetailRelationshipsDefaultLeaderboard {
	if o == nil || IsNil(o.DefaultGroupLeaderboard) {
		var ret GameCenterDetailRelationshipsDefaultLeaderboard
		return ret
	}
	return *o.DefaultGroupLeaderboard
}

// GetDefaultGroupLeaderboardOk returns a tuple with the DefaultGroupLeaderboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetDefaultGroupLeaderboardOk() (*GameCenterDetailRelationshipsDefaultLeaderboard, bool) {
	if o == nil || IsNil(o.DefaultGroupLeaderboard) {
		return nil, false
	}
	return o.DefaultGroupLeaderboard, true
}

// HasDefaultGroupLeaderboard returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasDefaultGroupLeaderboard() bool {
	if o != nil && !IsNil(o.DefaultGroupLeaderboard) {
		return true
	}

	return false
}

// SetDefaultGroupLeaderboard gets a reference to the given GameCenterDetailRelationshipsDefaultLeaderboard and assigns it to the DefaultGroupLeaderboard field.
func (o *GameCenterDetailRelationships) SetDefaultGroupLeaderboard(v GameCenterDetailRelationshipsDefaultLeaderboard) {
	o.DefaultGroupLeaderboard = &v
}

// GetAchievementReleases returns the AchievementReleases field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetAchievementReleases() GameCenterAchievementRelationshipsReleases {
	if o == nil || IsNil(o.AchievementReleases) {
		var ret GameCenterAchievementRelationshipsReleases
		return ret
	}
	return *o.AchievementReleases
}

// GetAchievementReleasesOk returns a tuple with the AchievementReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetAchievementReleasesOk() (*GameCenterAchievementRelationshipsReleases, bool) {
	if o == nil || IsNil(o.AchievementReleases) {
		return nil, false
	}
	return o.AchievementReleases, true
}

// HasAchievementReleases returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasAchievementReleases() bool {
	if o != nil && !IsNil(o.AchievementReleases) {
		return true
	}

	return false
}

// SetAchievementReleases gets a reference to the given GameCenterAchievementRelationshipsReleases and assigns it to the AchievementReleases field.
func (o *GameCenterDetailRelationships) SetAchievementReleases(v GameCenterAchievementRelationshipsReleases) {
	o.AchievementReleases = &v
}

// GetLeaderboardReleases returns the LeaderboardReleases field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetLeaderboardReleases() GameCenterDetailRelationshipsLeaderboardReleases {
	if o == nil || IsNil(o.LeaderboardReleases) {
		var ret GameCenterDetailRelationshipsLeaderboardReleases
		return ret
	}
	return *o.LeaderboardReleases
}

// GetLeaderboardReleasesOk returns a tuple with the LeaderboardReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetLeaderboardReleasesOk() (*GameCenterDetailRelationshipsLeaderboardReleases, bool) {
	if o == nil || IsNil(o.LeaderboardReleases) {
		return nil, false
	}
	return o.LeaderboardReleases, true
}

// HasLeaderboardReleases returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasLeaderboardReleases() bool {
	if o != nil && !IsNil(o.LeaderboardReleases) {
		return true
	}

	return false
}

// SetLeaderboardReleases gets a reference to the given GameCenterDetailRelationshipsLeaderboardReleases and assigns it to the LeaderboardReleases field.
func (o *GameCenterDetailRelationships) SetLeaderboardReleases(v GameCenterDetailRelationshipsLeaderboardReleases) {
	o.LeaderboardReleases = &v
}

// GetLeaderboardSetReleases returns the LeaderboardSetReleases field value if set, zero value otherwise.
func (o *GameCenterDetailRelationships) GetLeaderboardSetReleases() GameCenterDetailRelationshipsLeaderboardSetReleases {
	if o == nil || IsNil(o.LeaderboardSetReleases) {
		var ret GameCenterDetailRelationshipsLeaderboardSetReleases
		return ret
	}
	return *o.LeaderboardSetReleases
}

// GetLeaderboardSetReleasesOk returns a tuple with the LeaderboardSetReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationships) GetLeaderboardSetReleasesOk() (*GameCenterDetailRelationshipsLeaderboardSetReleases, bool) {
	if o == nil || IsNil(o.LeaderboardSetReleases) {
		return nil, false
	}
	return o.LeaderboardSetReleases, true
}

// HasLeaderboardSetReleases returns a boolean if a field has been set.
func (o *GameCenterDetailRelationships) HasLeaderboardSetReleases() bool {
	if o != nil && !IsNil(o.LeaderboardSetReleases) {
		return true
	}

	return false
}

// SetLeaderboardSetReleases gets a reference to the given GameCenterDetailRelationshipsLeaderboardSetReleases and assigns it to the LeaderboardSetReleases field.
func (o *GameCenterDetailRelationships) SetLeaderboardSetReleases(v GameCenterDetailRelationshipsLeaderboardSetReleases) {
	o.LeaderboardSetReleases = &v
}

func (o GameCenterDetailRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.GameCenterAppVersions) {
		toSerialize["gameCenterAppVersions"] = o.GameCenterAppVersions
	}
	if !IsNil(o.GameCenterGroup) {
		toSerialize["gameCenterGroup"] = o.GameCenterGroup
	}
	if !IsNil(o.GameCenterLeaderboards) {
		toSerialize["gameCenterLeaderboards"] = o.GameCenterLeaderboards
	}
	if !IsNil(o.GameCenterLeaderboardSets) {
		toSerialize["gameCenterLeaderboardSets"] = o.GameCenterLeaderboardSets
	}
	if !IsNil(o.GameCenterAchievements) {
		toSerialize["gameCenterAchievements"] = o.GameCenterAchievements
	}
	if !IsNil(o.DefaultLeaderboard) {
		toSerialize["defaultLeaderboard"] = o.DefaultLeaderboard
	}
	if !IsNil(o.DefaultGroupLeaderboard) {
		toSerialize["defaultGroupLeaderboard"] = o.DefaultGroupLeaderboard
	}
	if !IsNil(o.AchievementReleases) {
		toSerialize["achievementReleases"] = o.AchievementReleases
	}
	if !IsNil(o.LeaderboardReleases) {
		toSerialize["leaderboardReleases"] = o.LeaderboardReleases
	}
	if !IsNil(o.LeaderboardSetReleases) {
		toSerialize["leaderboardSetReleases"] = o.LeaderboardSetReleases
	}
	return toSerialize, nil
}

type NullableGameCenterDetailRelationships struct {
	value *GameCenterDetailRelationships
	isSet bool
}

func (v NullableGameCenterDetailRelationships) Get() *GameCenterDetailRelationships {
	return v.value
}

func (v *NullableGameCenterDetailRelationships) Set(val *GameCenterDetailRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailRelationships(val *GameCenterDetailRelationships) *NullableGameCenterDetailRelationships {
	return &NullableGameCenterDetailRelationships{value: val, isSet: true}
}

func (v NullableGameCenterDetailRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


