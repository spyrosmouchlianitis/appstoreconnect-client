# GameCenterMatchmakingQueueUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterMatchmakingQueueUpdateRequestDataAttributes**](GameCenterMatchmakingQueueUpdateRequestDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterMatchmakingQueueRelationships**](GameCenterMatchmakingQueueRelationships.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingQueueUpdateRequestData

`func NewGameCenterMatchmakingQueueUpdateRequestData(type_ string, id string, ) *GameCenterMatchmakingQueueUpdateRequestData`

NewGameCenterMatchmakingQueueUpdateRequestData instantiates a new GameCenterMatchmakingQueueUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingQueueUpdateRequestDataWithDefaults

`func NewGameCenterMatchmakingQueueUpdateRequestDataWithDefaults() *GameCenterMatchmakingQueueUpdateRequestData`

NewGameCenterMatchmakingQueueUpdateRequestDataWithDefaults instantiates a new GameCenterMatchmakingQueueUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterMatchmakingQueueUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterMatchmakingQueueUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterMatchmakingQueueUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterMatchmakingQueueUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterMatchmakingQueueUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterMatchmakingQueueUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterMatchmakingQueueUpdateRequestData) GetAttributes() GameCenterMatchmakingQueueUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterMatchmakingQueueUpdateRequestData) GetAttributesOk() (*GameCenterMatchmakingQueueUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterMatchmakingQueueUpdateRequestData) SetAttributes(v GameCenterMatchmakingQueueUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterMatchmakingQueueUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterMatchmakingQueueUpdateRequestData) GetRelationships() GameCenterMatchmakingQueueRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterMatchmakingQueueUpdateRequestData) GetRelationshipsOk() (*GameCenterMatchmakingQueueRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterMatchmakingQueueUpdateRequestData) SetRelationships(v GameCenterMatchmakingQueueRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterMatchmakingQueueUpdateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


