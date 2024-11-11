# GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | **string** |  | 
**ChallengeIds** | Pointer to **[]string** |  | [optional] 
**Context** | Pointer to **float64** |  | [optional] 
**ScopedPlayerId** | **string** |  | 
**Score** | **float64** |  | 
**SubmittedDate** | Pointer to **time.Time** |  | [optional] 
**VendorIdentifier** | **string** |  | 

## Methods

### NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes

`func NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes(bundleId string, scopedPlayerId string, score float64, vendorIdentifier string, ) *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes`

NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes instantiates a new GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributesWithDefaults

`func NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributesWithDefaults() *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes`

NewGameCenterLeaderboardEntrySubmissionCreateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetChallengeIds

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetChallengeIds() []string`

GetChallengeIds returns the ChallengeIds field if non-nil, zero value otherwise.

### GetChallengeIdsOk

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetChallengeIdsOk() (*[]string, bool)`

GetChallengeIdsOk returns a tuple with the ChallengeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeIds

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetChallengeIds(v []string)`

SetChallengeIds sets ChallengeIds field to given value.

### HasChallengeIds

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) HasChallengeIds() bool`

HasChallengeIds returns a boolean if a field has been set.

### GetContext

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetContext() float64`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetContextOk() (*float64, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetContext(v float64)`

SetContext sets Context field to given value.

### HasContext

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetScopedPlayerId

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetScopedPlayerId() string`

GetScopedPlayerId returns the ScopedPlayerId field if non-nil, zero value otherwise.

### GetScopedPlayerIdOk

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetScopedPlayerIdOk() (*string, bool)`

GetScopedPlayerIdOk returns a tuple with the ScopedPlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedPlayerId

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetScopedPlayerId(v string)`

SetScopedPlayerId sets ScopedPlayerId field to given value.


### GetScore

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetScore(v float64)`

SetScore sets Score field to given value.


### GetSubmittedDate

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetSubmittedDate() time.Time`

GetSubmittedDate returns the SubmittedDate field if non-nil, zero value otherwise.

### GetSubmittedDateOk

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetSubmittedDateOk() (*time.Time, bool)`

GetSubmittedDateOk returns a tuple with the SubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDate

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetSubmittedDate(v time.Time)`

SetSubmittedDate sets SubmittedDate field to given value.

### HasSubmittedDate

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) HasSubmittedDate() bool`

HasSubmittedDate returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *GameCenterLeaderboardEntrySubmissionCreateRequestDataAttributes) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


