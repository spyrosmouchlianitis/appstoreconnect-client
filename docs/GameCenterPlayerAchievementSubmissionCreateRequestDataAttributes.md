# GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | **string** |  | 
**ChallengeIds** | Pointer to **[]string** |  | [optional] 
**PercentageAchieved** | **int32** |  | 
**ScopedPlayerId** | **string** |  | 
**SubmittedDate** | Pointer to **time.Time** |  | [optional] 
**VendorIdentifier** | **string** |  | 

## Methods

### NewGameCenterPlayerAchievementSubmissionCreateRequestDataAttributes

`func NewGameCenterPlayerAchievementSubmissionCreateRequestDataAttributes(bundleId string, percentageAchieved int32, scopedPlayerId string, vendorIdentifier string, ) *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes`

NewGameCenterPlayerAchievementSubmissionCreateRequestDataAttributes instantiates a new GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterPlayerAchievementSubmissionCreateRequestDataAttributesWithDefaults

`func NewGameCenterPlayerAchievementSubmissionCreateRequestDataAttributesWithDefaults() *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes`

NewGameCenterPlayerAchievementSubmissionCreateRequestDataAttributesWithDefaults instantiates a new GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetChallengeIds

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetChallengeIds() []string`

GetChallengeIds returns the ChallengeIds field if non-nil, zero value otherwise.

### GetChallengeIdsOk

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetChallengeIdsOk() (*[]string, bool)`

GetChallengeIdsOk returns a tuple with the ChallengeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeIds

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) SetChallengeIds(v []string)`

SetChallengeIds sets ChallengeIds field to given value.

### HasChallengeIds

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) HasChallengeIds() bool`

HasChallengeIds returns a boolean if a field has been set.

### GetPercentageAchieved

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetPercentageAchieved() int32`

GetPercentageAchieved returns the PercentageAchieved field if non-nil, zero value otherwise.

### GetPercentageAchievedOk

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetPercentageAchievedOk() (*int32, bool)`

GetPercentageAchievedOk returns a tuple with the PercentageAchieved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageAchieved

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) SetPercentageAchieved(v int32)`

SetPercentageAchieved sets PercentageAchieved field to given value.


### GetScopedPlayerId

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetScopedPlayerId() string`

GetScopedPlayerId returns the ScopedPlayerId field if non-nil, zero value otherwise.

### GetScopedPlayerIdOk

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetScopedPlayerIdOk() (*string, bool)`

GetScopedPlayerIdOk returns a tuple with the ScopedPlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedPlayerId

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) SetScopedPlayerId(v string)`

SetScopedPlayerId sets ScopedPlayerId field to given value.


### GetSubmittedDate

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetSubmittedDate() time.Time`

GetSubmittedDate returns the SubmittedDate field if non-nil, zero value otherwise.

### GetSubmittedDateOk

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetSubmittedDateOk() (*time.Time, bool)`

GetSubmittedDateOk returns a tuple with the SubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDate

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) SetSubmittedDate(v time.Time)`

SetSubmittedDate sets SubmittedDate field to given value.

### HasSubmittedDate

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) HasSubmittedDate() bool`

HasSubmittedDate returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *GameCenterPlayerAchievementSubmissionCreateRequestDataAttributes) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


