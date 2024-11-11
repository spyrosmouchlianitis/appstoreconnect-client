# GameCenterLeaderboardEntrySubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterLeaderboardEntrySubmissionAttributes**](GameCenterLeaderboardEntrySubmissionAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardEntrySubmission

`func NewGameCenterLeaderboardEntrySubmission(type_ string, id string, ) *GameCenterLeaderboardEntrySubmission`

NewGameCenterLeaderboardEntrySubmission instantiates a new GameCenterLeaderboardEntrySubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardEntrySubmissionWithDefaults

`func NewGameCenterLeaderboardEntrySubmissionWithDefaults() *GameCenterLeaderboardEntrySubmission`

NewGameCenterLeaderboardEntrySubmissionWithDefaults instantiates a new GameCenterLeaderboardEntrySubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterLeaderboardEntrySubmission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterLeaderboardEntrySubmission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterLeaderboardEntrySubmission) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterLeaderboardEntrySubmission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterLeaderboardEntrySubmission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterLeaderboardEntrySubmission) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterLeaderboardEntrySubmission) GetAttributes() GameCenterLeaderboardEntrySubmissionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterLeaderboardEntrySubmission) GetAttributesOk() (*GameCenterLeaderboardEntrySubmissionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterLeaderboardEntrySubmission) SetAttributes(v GameCenterLeaderboardEntrySubmissionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterLeaderboardEntrySubmission) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardEntrySubmission) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardEntrySubmission) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardEntrySubmission) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterLeaderboardEntrySubmission) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


