# ReviewSubmissionItemCreateRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewSubmission** | [**ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission**](ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission.md) |  | 
**AppStoreVersion** | Pointer to [**AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion**](AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion.md) |  | [optional] 
**AppCustomProductPageVersion** | Pointer to [**AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion**](AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion.md) |  | [optional] 
**AppStoreVersionExperiment** | Pointer to [**AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment**](AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment.md) |  | [optional] 
**AppStoreVersionExperimentV2** | Pointer to [**AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment**](AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment.md) |  | [optional] 
**AppEvent** | Pointer to [**AppEventLocalizationRelationshipsAppEvent**](AppEventLocalizationRelationshipsAppEvent.md) |  | [optional] 

## Methods

### NewReviewSubmissionItemCreateRequestDataRelationships

`func NewReviewSubmissionItemCreateRequestDataRelationships(reviewSubmission ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission, ) *ReviewSubmissionItemCreateRequestDataRelationships`

NewReviewSubmissionItemCreateRequestDataRelationships instantiates a new ReviewSubmissionItemCreateRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewSubmissionItemCreateRequestDataRelationshipsWithDefaults

`func NewReviewSubmissionItemCreateRequestDataRelationshipsWithDefaults() *ReviewSubmissionItemCreateRequestDataRelationships`

NewReviewSubmissionItemCreateRequestDataRelationshipsWithDefaults instantiates a new ReviewSubmissionItemCreateRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewSubmission

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetReviewSubmission() ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission`

GetReviewSubmission returns the ReviewSubmission field if non-nil, zero value otherwise.

### GetReviewSubmissionOk

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetReviewSubmissionOk() (*ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission, bool)`

GetReviewSubmissionOk returns a tuple with the ReviewSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewSubmission

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetReviewSubmission(v ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission)`

SetReviewSubmission sets ReviewSubmission field to given value.


### GetAppStoreVersion

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersion() AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion`

GetAppStoreVersion returns the AppStoreVersion field if non-nil, zero value otherwise.

### GetAppStoreVersionOk

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionOk() (*AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion, bool)`

GetAppStoreVersionOk returns a tuple with the AppStoreVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreVersion

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppStoreVersion(v AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion)`

SetAppStoreVersion sets AppStoreVersion field to given value.

### HasAppStoreVersion

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppStoreVersion() bool`

HasAppStoreVersion returns a boolean if a field has been set.

### GetAppCustomProductPageVersion

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppCustomProductPageVersion() AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion`

GetAppCustomProductPageVersion returns the AppCustomProductPageVersion field if non-nil, zero value otherwise.

### GetAppCustomProductPageVersionOk

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppCustomProductPageVersionOk() (*AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion, bool)`

GetAppCustomProductPageVersionOk returns a tuple with the AppCustomProductPageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCustomProductPageVersion

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppCustomProductPageVersion(v AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion)`

SetAppCustomProductPageVersion sets AppCustomProductPageVersion field to given value.

### HasAppCustomProductPageVersion

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppCustomProductPageVersion() bool`

HasAppCustomProductPageVersion returns a boolean if a field has been set.

### GetAppStoreVersionExperiment

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionExperiment() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment`

GetAppStoreVersionExperiment returns the AppStoreVersionExperiment field if non-nil, zero value otherwise.

### GetAppStoreVersionExperimentOk

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionExperimentOk() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool)`

GetAppStoreVersionExperimentOk returns a tuple with the AppStoreVersionExperiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreVersionExperiment

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppStoreVersionExperiment(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment)`

SetAppStoreVersionExperiment sets AppStoreVersionExperiment field to given value.

### HasAppStoreVersionExperiment

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppStoreVersionExperiment() bool`

HasAppStoreVersionExperiment returns a boolean if a field has been set.

### GetAppStoreVersionExperimentV2

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionExperimentV2() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment`

GetAppStoreVersionExperimentV2 returns the AppStoreVersionExperimentV2 field if non-nil, zero value otherwise.

### GetAppStoreVersionExperimentV2Ok

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionExperimentV2Ok() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool)`

GetAppStoreVersionExperimentV2Ok returns a tuple with the AppStoreVersionExperimentV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreVersionExperimentV2

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppStoreVersionExperimentV2(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment)`

SetAppStoreVersionExperimentV2 sets AppStoreVersionExperimentV2 field to given value.

### HasAppStoreVersionExperimentV2

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppStoreVersionExperimentV2() bool`

HasAppStoreVersionExperimentV2 returns a boolean if a field has been set.

### GetAppEvent

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppEvent() AppEventLocalizationRelationshipsAppEvent`

GetAppEvent returns the AppEvent field if non-nil, zero value otherwise.

### GetAppEventOk

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppEventOk() (*AppEventLocalizationRelationshipsAppEvent, bool)`

GetAppEventOk returns a tuple with the AppEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEvent

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppEvent(v AppEventLocalizationRelationshipsAppEvent)`

SetAppEvent sets AppEvent field to given value.

### HasAppEvent

`func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppEvent() bool`

HasAppEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


