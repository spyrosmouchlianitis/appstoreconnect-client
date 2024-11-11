# AppInfoAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppStoreState** | Pointer to [**AppStoreVersionState**](AppStoreVersionState.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**AppStoreAgeRating** | Pointer to [**AppStoreAgeRating**](AppStoreAgeRating.md) |  | [optional] 
**AustraliaAgeRating** | Pointer to **string** |  | [optional] 
**BrazilAgeRating** | Pointer to [**BrazilAgeRating**](BrazilAgeRating.md) |  | [optional] 
**BrazilAgeRatingV2** | Pointer to **string** |  | [optional] 
**KoreaAgeRating** | Pointer to **string** |  | [optional] 
**KidsAgeBand** | Pointer to [**KidsAgeBand**](KidsAgeBand.md) |  | [optional] 

## Methods

### NewAppInfoAttributes

`func NewAppInfoAttributes() *AppInfoAttributes`

NewAppInfoAttributes instantiates a new AppInfoAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoAttributesWithDefaults

`func NewAppInfoAttributesWithDefaults() *AppInfoAttributes`

NewAppInfoAttributesWithDefaults instantiates a new AppInfoAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppStoreState

`func (o *AppInfoAttributes) GetAppStoreState() AppStoreVersionState`

GetAppStoreState returns the AppStoreState field if non-nil, zero value otherwise.

### GetAppStoreStateOk

`func (o *AppInfoAttributes) GetAppStoreStateOk() (*AppStoreVersionState, bool)`

GetAppStoreStateOk returns a tuple with the AppStoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreState

`func (o *AppInfoAttributes) SetAppStoreState(v AppStoreVersionState)`

SetAppStoreState sets AppStoreState field to given value.

### HasAppStoreState

`func (o *AppInfoAttributes) HasAppStoreState() bool`

HasAppStoreState returns a boolean if a field has been set.

### GetState

`func (o *AppInfoAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppInfoAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppInfoAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AppInfoAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAppStoreAgeRating

`func (o *AppInfoAttributes) GetAppStoreAgeRating() AppStoreAgeRating`

GetAppStoreAgeRating returns the AppStoreAgeRating field if non-nil, zero value otherwise.

### GetAppStoreAgeRatingOk

`func (o *AppInfoAttributes) GetAppStoreAgeRatingOk() (*AppStoreAgeRating, bool)`

GetAppStoreAgeRatingOk returns a tuple with the AppStoreAgeRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreAgeRating

`func (o *AppInfoAttributes) SetAppStoreAgeRating(v AppStoreAgeRating)`

SetAppStoreAgeRating sets AppStoreAgeRating field to given value.

### HasAppStoreAgeRating

`func (o *AppInfoAttributes) HasAppStoreAgeRating() bool`

HasAppStoreAgeRating returns a boolean if a field has been set.

### GetAustraliaAgeRating

`func (o *AppInfoAttributes) GetAustraliaAgeRating() string`

GetAustraliaAgeRating returns the AustraliaAgeRating field if non-nil, zero value otherwise.

### GetAustraliaAgeRatingOk

`func (o *AppInfoAttributes) GetAustraliaAgeRatingOk() (*string, bool)`

GetAustraliaAgeRatingOk returns a tuple with the AustraliaAgeRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAustraliaAgeRating

`func (o *AppInfoAttributes) SetAustraliaAgeRating(v string)`

SetAustraliaAgeRating sets AustraliaAgeRating field to given value.

### HasAustraliaAgeRating

`func (o *AppInfoAttributes) HasAustraliaAgeRating() bool`

HasAustraliaAgeRating returns a boolean if a field has been set.

### GetBrazilAgeRating

`func (o *AppInfoAttributes) GetBrazilAgeRating() BrazilAgeRating`

GetBrazilAgeRating returns the BrazilAgeRating field if non-nil, zero value otherwise.

### GetBrazilAgeRatingOk

`func (o *AppInfoAttributes) GetBrazilAgeRatingOk() (*BrazilAgeRating, bool)`

GetBrazilAgeRatingOk returns a tuple with the BrazilAgeRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrazilAgeRating

`func (o *AppInfoAttributes) SetBrazilAgeRating(v BrazilAgeRating)`

SetBrazilAgeRating sets BrazilAgeRating field to given value.

### HasBrazilAgeRating

`func (o *AppInfoAttributes) HasBrazilAgeRating() bool`

HasBrazilAgeRating returns a boolean if a field has been set.

### GetBrazilAgeRatingV2

`func (o *AppInfoAttributes) GetBrazilAgeRatingV2() string`

GetBrazilAgeRatingV2 returns the BrazilAgeRatingV2 field if non-nil, zero value otherwise.

### GetBrazilAgeRatingV2Ok

`func (o *AppInfoAttributes) GetBrazilAgeRatingV2Ok() (*string, bool)`

GetBrazilAgeRatingV2Ok returns a tuple with the BrazilAgeRatingV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrazilAgeRatingV2

`func (o *AppInfoAttributes) SetBrazilAgeRatingV2(v string)`

SetBrazilAgeRatingV2 sets BrazilAgeRatingV2 field to given value.

### HasBrazilAgeRatingV2

`func (o *AppInfoAttributes) HasBrazilAgeRatingV2() bool`

HasBrazilAgeRatingV2 returns a boolean if a field has been set.

### GetKoreaAgeRating

`func (o *AppInfoAttributes) GetKoreaAgeRating() string`

GetKoreaAgeRating returns the KoreaAgeRating field if non-nil, zero value otherwise.

### GetKoreaAgeRatingOk

`func (o *AppInfoAttributes) GetKoreaAgeRatingOk() (*string, bool)`

GetKoreaAgeRatingOk returns a tuple with the KoreaAgeRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKoreaAgeRating

`func (o *AppInfoAttributes) SetKoreaAgeRating(v string)`

SetKoreaAgeRating sets KoreaAgeRating field to given value.

### HasKoreaAgeRating

`func (o *AppInfoAttributes) HasKoreaAgeRating() bool`

HasKoreaAgeRating returns a boolean if a field has been set.

### GetKidsAgeBand

`func (o *AppInfoAttributes) GetKidsAgeBand() KidsAgeBand`

GetKidsAgeBand returns the KidsAgeBand field if non-nil, zero value otherwise.

### GetKidsAgeBandOk

`func (o *AppInfoAttributes) GetKidsAgeBandOk() (*KidsAgeBand, bool)`

GetKidsAgeBandOk returns a tuple with the KidsAgeBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKidsAgeBand

`func (o *AppInfoAttributes) SetKidsAgeBand(v KidsAgeBand)`

SetKidsAgeBand sets KidsAgeBand field to given value.

### HasKidsAgeBand

`func (o *AppInfoAttributes) HasKidsAgeBand() bool`

HasKidsAgeBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


