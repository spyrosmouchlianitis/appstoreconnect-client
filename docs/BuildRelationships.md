# BuildRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreReleaseVersion** | Pointer to [**BuildRelationshipsPreReleaseVersion**](BuildRelationshipsPreReleaseVersion.md) |  | [optional] 
**IndividualTesters** | Pointer to [**BetaGroupRelationshipsBetaTesters**](BetaGroupRelationshipsBetaTesters.md) |  | [optional] 
**BetaGroups** | Pointer to [**AppRelationshipsBetaGroups**](AppRelationshipsBetaGroups.md) |  | [optional] 
**BetaBuildLocalizations** | Pointer to [**BuildRelationshipsBetaBuildLocalizations**](BuildRelationshipsBetaBuildLocalizations.md) |  | [optional] 
**AppEncryptionDeclaration** | Pointer to [**BuildRelationshipsAppEncryptionDeclaration**](BuildRelationshipsAppEncryptionDeclaration.md) |  | [optional] 
**BetaAppReviewSubmission** | Pointer to [**BuildRelationshipsBetaAppReviewSubmission**](BuildRelationshipsBetaAppReviewSubmission.md) |  | [optional] 
**App** | Pointer to [**BetaAppLocalizationRelationshipsApp**](BetaAppLocalizationRelationshipsApp.md) |  | [optional] 
**BuildBetaDetail** | Pointer to [**BuildRelationshipsBuildBetaDetail**](BuildRelationshipsBuildBetaDetail.md) |  | [optional] 
**AppStoreVersion** | Pointer to [**AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion**](AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion.md) |  | [optional] 
**Icons** | Pointer to [**BuildRelationshipsIcons**](BuildRelationshipsIcons.md) |  | [optional] 
**BuildBundles** | Pointer to [**BuildRelationshipsBuildBundles**](BuildRelationshipsBuildBundles.md) |  | [optional] 
**PerfPowerMetrics** | Pointer to [**AnalyticsReportInstanceRelationshipsSegments**](AnalyticsReportInstanceRelationshipsSegments.md) |  | [optional] 
**DiagnosticSignatures** | Pointer to [**AnalyticsReportInstanceRelationshipsSegments**](AnalyticsReportInstanceRelationshipsSegments.md) |  | [optional] 

## Methods

### NewBuildRelationships

`func NewBuildRelationships() *BuildRelationships`

NewBuildRelationships instantiates a new BuildRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildRelationshipsWithDefaults

`func NewBuildRelationshipsWithDefaults() *BuildRelationships`

NewBuildRelationshipsWithDefaults instantiates a new BuildRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreReleaseVersion

`func (o *BuildRelationships) GetPreReleaseVersion() BuildRelationshipsPreReleaseVersion`

GetPreReleaseVersion returns the PreReleaseVersion field if non-nil, zero value otherwise.

### GetPreReleaseVersionOk

`func (o *BuildRelationships) GetPreReleaseVersionOk() (*BuildRelationshipsPreReleaseVersion, bool)`

GetPreReleaseVersionOk returns a tuple with the PreReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreReleaseVersion

`func (o *BuildRelationships) SetPreReleaseVersion(v BuildRelationshipsPreReleaseVersion)`

SetPreReleaseVersion sets PreReleaseVersion field to given value.

### HasPreReleaseVersion

`func (o *BuildRelationships) HasPreReleaseVersion() bool`

HasPreReleaseVersion returns a boolean if a field has been set.

### GetIndividualTesters

`func (o *BuildRelationships) GetIndividualTesters() BetaGroupRelationshipsBetaTesters`

GetIndividualTesters returns the IndividualTesters field if non-nil, zero value otherwise.

### GetIndividualTestersOk

`func (o *BuildRelationships) GetIndividualTestersOk() (*BetaGroupRelationshipsBetaTesters, bool)`

GetIndividualTestersOk returns a tuple with the IndividualTesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualTesters

`func (o *BuildRelationships) SetIndividualTesters(v BetaGroupRelationshipsBetaTesters)`

SetIndividualTesters sets IndividualTesters field to given value.

### HasIndividualTesters

`func (o *BuildRelationships) HasIndividualTesters() bool`

HasIndividualTesters returns a boolean if a field has been set.

### GetBetaGroups

`func (o *BuildRelationships) GetBetaGroups() AppRelationshipsBetaGroups`

GetBetaGroups returns the BetaGroups field if non-nil, zero value otherwise.

### GetBetaGroupsOk

`func (o *BuildRelationships) GetBetaGroupsOk() (*AppRelationshipsBetaGroups, bool)`

GetBetaGroupsOk returns a tuple with the BetaGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaGroups

`func (o *BuildRelationships) SetBetaGroups(v AppRelationshipsBetaGroups)`

SetBetaGroups sets BetaGroups field to given value.

### HasBetaGroups

`func (o *BuildRelationships) HasBetaGroups() bool`

HasBetaGroups returns a boolean if a field has been set.

### GetBetaBuildLocalizations

`func (o *BuildRelationships) GetBetaBuildLocalizations() BuildRelationshipsBetaBuildLocalizations`

GetBetaBuildLocalizations returns the BetaBuildLocalizations field if non-nil, zero value otherwise.

### GetBetaBuildLocalizationsOk

`func (o *BuildRelationships) GetBetaBuildLocalizationsOk() (*BuildRelationshipsBetaBuildLocalizations, bool)`

GetBetaBuildLocalizationsOk returns a tuple with the BetaBuildLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaBuildLocalizations

`func (o *BuildRelationships) SetBetaBuildLocalizations(v BuildRelationshipsBetaBuildLocalizations)`

SetBetaBuildLocalizations sets BetaBuildLocalizations field to given value.

### HasBetaBuildLocalizations

`func (o *BuildRelationships) HasBetaBuildLocalizations() bool`

HasBetaBuildLocalizations returns a boolean if a field has been set.

### GetAppEncryptionDeclaration

`func (o *BuildRelationships) GetAppEncryptionDeclaration() BuildRelationshipsAppEncryptionDeclaration`

GetAppEncryptionDeclaration returns the AppEncryptionDeclaration field if non-nil, zero value otherwise.

### GetAppEncryptionDeclarationOk

`func (o *BuildRelationships) GetAppEncryptionDeclarationOk() (*BuildRelationshipsAppEncryptionDeclaration, bool)`

GetAppEncryptionDeclarationOk returns a tuple with the AppEncryptionDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEncryptionDeclaration

`func (o *BuildRelationships) SetAppEncryptionDeclaration(v BuildRelationshipsAppEncryptionDeclaration)`

SetAppEncryptionDeclaration sets AppEncryptionDeclaration field to given value.

### HasAppEncryptionDeclaration

`func (o *BuildRelationships) HasAppEncryptionDeclaration() bool`

HasAppEncryptionDeclaration returns a boolean if a field has been set.

### GetBetaAppReviewSubmission

`func (o *BuildRelationships) GetBetaAppReviewSubmission() BuildRelationshipsBetaAppReviewSubmission`

GetBetaAppReviewSubmission returns the BetaAppReviewSubmission field if non-nil, zero value otherwise.

### GetBetaAppReviewSubmissionOk

`func (o *BuildRelationships) GetBetaAppReviewSubmissionOk() (*BuildRelationshipsBetaAppReviewSubmission, bool)`

GetBetaAppReviewSubmissionOk returns a tuple with the BetaAppReviewSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaAppReviewSubmission

`func (o *BuildRelationships) SetBetaAppReviewSubmission(v BuildRelationshipsBetaAppReviewSubmission)`

SetBetaAppReviewSubmission sets BetaAppReviewSubmission field to given value.

### HasBetaAppReviewSubmission

`func (o *BuildRelationships) HasBetaAppReviewSubmission() bool`

HasBetaAppReviewSubmission returns a boolean if a field has been set.

### GetApp

`func (o *BuildRelationships) GetApp() BetaAppLocalizationRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *BuildRelationships) GetAppOk() (*BetaAppLocalizationRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *BuildRelationships) SetApp(v BetaAppLocalizationRelationshipsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *BuildRelationships) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetBuildBetaDetail

`func (o *BuildRelationships) GetBuildBetaDetail() BuildRelationshipsBuildBetaDetail`

GetBuildBetaDetail returns the BuildBetaDetail field if non-nil, zero value otherwise.

### GetBuildBetaDetailOk

`func (o *BuildRelationships) GetBuildBetaDetailOk() (*BuildRelationshipsBuildBetaDetail, bool)`

GetBuildBetaDetailOk returns a tuple with the BuildBetaDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildBetaDetail

`func (o *BuildRelationships) SetBuildBetaDetail(v BuildRelationshipsBuildBetaDetail)`

SetBuildBetaDetail sets BuildBetaDetail field to given value.

### HasBuildBetaDetail

`func (o *BuildRelationships) HasBuildBetaDetail() bool`

HasBuildBetaDetail returns a boolean if a field has been set.

### GetAppStoreVersion

`func (o *BuildRelationships) GetAppStoreVersion() AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion`

GetAppStoreVersion returns the AppStoreVersion field if non-nil, zero value otherwise.

### GetAppStoreVersionOk

`func (o *BuildRelationships) GetAppStoreVersionOk() (*AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion, bool)`

GetAppStoreVersionOk returns a tuple with the AppStoreVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreVersion

`func (o *BuildRelationships) SetAppStoreVersion(v AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion)`

SetAppStoreVersion sets AppStoreVersion field to given value.

### HasAppStoreVersion

`func (o *BuildRelationships) HasAppStoreVersion() bool`

HasAppStoreVersion returns a boolean if a field has been set.

### GetIcons

`func (o *BuildRelationships) GetIcons() BuildRelationshipsIcons`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *BuildRelationships) GetIconsOk() (*BuildRelationshipsIcons, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *BuildRelationships) SetIcons(v BuildRelationshipsIcons)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *BuildRelationships) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### GetBuildBundles

`func (o *BuildRelationships) GetBuildBundles() BuildRelationshipsBuildBundles`

GetBuildBundles returns the BuildBundles field if non-nil, zero value otherwise.

### GetBuildBundlesOk

`func (o *BuildRelationships) GetBuildBundlesOk() (*BuildRelationshipsBuildBundles, bool)`

GetBuildBundlesOk returns a tuple with the BuildBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildBundles

`func (o *BuildRelationships) SetBuildBundles(v BuildRelationshipsBuildBundles)`

SetBuildBundles sets BuildBundles field to given value.

### HasBuildBundles

`func (o *BuildRelationships) HasBuildBundles() bool`

HasBuildBundles returns a boolean if a field has been set.

### GetPerfPowerMetrics

`func (o *BuildRelationships) GetPerfPowerMetrics() AnalyticsReportInstanceRelationshipsSegments`

GetPerfPowerMetrics returns the PerfPowerMetrics field if non-nil, zero value otherwise.

### GetPerfPowerMetricsOk

`func (o *BuildRelationships) GetPerfPowerMetricsOk() (*AnalyticsReportInstanceRelationshipsSegments, bool)`

GetPerfPowerMetricsOk returns a tuple with the PerfPowerMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfPowerMetrics

`func (o *BuildRelationships) SetPerfPowerMetrics(v AnalyticsReportInstanceRelationshipsSegments)`

SetPerfPowerMetrics sets PerfPowerMetrics field to given value.

### HasPerfPowerMetrics

`func (o *BuildRelationships) HasPerfPowerMetrics() bool`

HasPerfPowerMetrics returns a boolean if a field has been set.

### GetDiagnosticSignatures

`func (o *BuildRelationships) GetDiagnosticSignatures() AnalyticsReportInstanceRelationshipsSegments`

GetDiagnosticSignatures returns the DiagnosticSignatures field if non-nil, zero value otherwise.

### GetDiagnosticSignaturesOk

`func (o *BuildRelationships) GetDiagnosticSignaturesOk() (*AnalyticsReportInstanceRelationshipsSegments, bool)`

GetDiagnosticSignaturesOk returns a tuple with the DiagnosticSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticSignatures

`func (o *BuildRelationships) SetDiagnosticSignatures(v AnalyticsReportInstanceRelationshipsSegments)`

SetDiagnosticSignatures sets DiagnosticSignatures field to given value.

### HasDiagnosticSignatures

`func (o *BuildRelationships) HasDiagnosticSignatures() bool`

HasDiagnosticSignatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


