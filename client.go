/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer( "%5B", "[", "%5D", "]" )
)

// APIClient manages communication with the App Store Connect API API v3.6.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	ActorsAPI *ActorsAPIService

	AgeRatingDeclarationsAPI *AgeRatingDeclarationsAPIService

	AlternativeDistributionDomainsAPI *AlternativeDistributionDomainsAPIService

	AlternativeDistributionKeysAPI *AlternativeDistributionKeysAPIService

	AlternativeDistributionPackageDeltasAPI *AlternativeDistributionPackageDeltasAPIService

	AlternativeDistributionPackageVariantsAPI *AlternativeDistributionPackageVariantsAPIService

	AlternativeDistributionPackageVersionsAPI *AlternativeDistributionPackageVersionsAPIService

	AlternativeDistributionPackagesAPI *AlternativeDistributionPackagesAPIService

	AnalyticsReportInstancesAPI *AnalyticsReportInstancesAPIService

	AnalyticsReportRequestsAPI *AnalyticsReportRequestsAPIService

	AnalyticsReportSegmentsAPI *AnalyticsReportSegmentsAPIService

	AnalyticsReportsAPI *AnalyticsReportsAPIService

	AppAvailabilitiesAPI *AppAvailabilitiesAPIService

	AppCategoriesAPI *AppCategoriesAPIService

	AppClipAdvancedExperienceImagesAPI *AppClipAdvancedExperienceImagesAPIService

	AppClipAdvancedExperiencesAPI *AppClipAdvancedExperiencesAPIService

	AppClipAppStoreReviewDetailsAPI *AppClipAppStoreReviewDetailsAPIService

	AppClipDefaultExperienceLocalizationsAPI *AppClipDefaultExperienceLocalizationsAPIService

	AppClipDefaultExperiencesAPI *AppClipDefaultExperiencesAPIService

	AppClipHeaderImagesAPI *AppClipHeaderImagesAPIService

	AppClipsAPI *AppClipsAPIService

	AppCustomProductPageLocalizationsAPI *AppCustomProductPageLocalizationsAPIService

	AppCustomProductPageVersionsAPI *AppCustomProductPageVersionsAPIService

	AppCustomProductPagesAPI *AppCustomProductPagesAPIService

	AppEncryptionDeclarationDocumentsAPI *AppEncryptionDeclarationDocumentsAPIService

	AppEncryptionDeclarationsAPI *AppEncryptionDeclarationsAPIService

	AppEventLocalizationsAPI *AppEventLocalizationsAPIService

	AppEventScreenshotsAPI *AppEventScreenshotsAPIService

	AppEventVideoClipsAPI *AppEventVideoClipsAPIService

	AppEventsAPI *AppEventsAPIService

	AppInfoLocalizationsAPI *AppInfoLocalizationsAPIService

	AppInfosAPI *AppInfosAPIService

	AppPreOrdersAPI *AppPreOrdersAPIService

	AppPreviewSetsAPI *AppPreviewSetsAPIService

	AppPreviewsAPI *AppPreviewsAPIService

	AppPricePointsAPI *AppPricePointsAPIService

	AppPriceSchedulesAPI *AppPriceSchedulesAPIService

	AppScreenshotSetsAPI *AppScreenshotSetsAPIService

	AppScreenshotsAPI *AppScreenshotsAPIService

	AppStoreReviewAttachmentsAPI *AppStoreReviewAttachmentsAPIService

	AppStoreReviewDetailsAPI *AppStoreReviewDetailsAPIService

	AppStoreVersionExperimentTreatmentLocalizationsAPI *AppStoreVersionExperimentTreatmentLocalizationsAPIService

	AppStoreVersionExperimentTreatmentsAPI *AppStoreVersionExperimentTreatmentsAPIService

	AppStoreVersionExperimentsAPI *AppStoreVersionExperimentsAPIService

	AppStoreVersionLocalizationsAPI *AppStoreVersionLocalizationsAPIService

	AppStoreVersionPhasedReleasesAPI *AppStoreVersionPhasedReleasesAPIService

	AppStoreVersionPromotionsAPI *AppStoreVersionPromotionsAPIService

	AppStoreVersionReleaseRequestsAPI *AppStoreVersionReleaseRequestsAPIService

	AppStoreVersionSubmissionsAPI *AppStoreVersionSubmissionsAPIService

	AppStoreVersionsAPI *AppStoreVersionsAPIService

	AppsAPI *AppsAPIService

	BetaAppClipInvocationLocalizationsAPI *BetaAppClipInvocationLocalizationsAPIService

	BetaAppClipInvocationsAPI *BetaAppClipInvocationsAPIService

	BetaAppLocalizationsAPI *BetaAppLocalizationsAPIService

	BetaAppReviewDetailsAPI *BetaAppReviewDetailsAPIService

	BetaAppReviewSubmissionsAPI *BetaAppReviewSubmissionsAPIService

	BetaBuildLocalizationsAPI *BetaBuildLocalizationsAPIService

	BetaGroupsAPI *BetaGroupsAPIService

	BetaLicenseAgreementsAPI *BetaLicenseAgreementsAPIService

	BetaTesterInvitationsAPI *BetaTesterInvitationsAPIService

	BetaTestersAPI *BetaTestersAPIService

	BuildBetaDetailsAPI *BuildBetaDetailsAPIService

	BuildBetaNotificationsAPI *BuildBetaNotificationsAPIService

	BuildBundlesAPI *BuildBundlesAPIService

	BuildsAPI *BuildsAPIService

	BundleIdCapabilitiesAPI *BundleIdCapabilitiesAPIService

	BundleIdsAPI *BundleIdsAPIService

	CertificatesAPI *CertificatesAPIService

	CiArtifactsAPI *CiArtifactsAPIService

	CiBuildActionsAPI *CiBuildActionsAPIService

	CiBuildRunsAPI *CiBuildRunsAPIService

	CiIssuesAPI *CiIssuesAPIService

	CiMacOsVersionsAPI *CiMacOsVersionsAPIService

	CiProductsAPI *CiProductsAPIService

	CiTestResultsAPI *CiTestResultsAPIService

	CiWorkflowsAPI *CiWorkflowsAPIService

	CiXcodeVersionsAPI *CiXcodeVersionsAPIService

	CustomerReviewResponsesAPI *CustomerReviewResponsesAPIService

	CustomerReviewsAPI *CustomerReviewsAPIService

	DevicesAPI *DevicesAPIService

	DiagnosticSignaturesAPI *DiagnosticSignaturesAPIService

	EndAppAvailabilityPreOrdersAPI *EndAppAvailabilityPreOrdersAPIService

	EndUserLicenseAgreementsAPI *EndUserLicenseAgreementsAPIService

	FinanceReportsAPI *FinanceReportsAPIService

	GameCenterAchievementImagesAPI *GameCenterAchievementImagesAPIService

	GameCenterAchievementLocalizationsAPI *GameCenterAchievementLocalizationsAPIService

	GameCenterAchievementReleasesAPI *GameCenterAchievementReleasesAPIService

	GameCenterAchievementsAPI *GameCenterAchievementsAPIService

	GameCenterAppVersionsAPI *GameCenterAppVersionsAPIService

	GameCenterDetailsAPI *GameCenterDetailsAPIService

	GameCenterEnabledVersionsAPI *GameCenterEnabledVersionsAPIService

	GameCenterGroupsAPI *GameCenterGroupsAPIService

	GameCenterLeaderboardEntrySubmissionsAPI *GameCenterLeaderboardEntrySubmissionsAPIService

	GameCenterLeaderboardImagesAPI *GameCenterLeaderboardImagesAPIService

	GameCenterLeaderboardLocalizationsAPI *GameCenterLeaderboardLocalizationsAPIService

	GameCenterLeaderboardReleasesAPI *GameCenterLeaderboardReleasesAPIService

	GameCenterLeaderboardSetImagesAPI *GameCenterLeaderboardSetImagesAPIService

	GameCenterLeaderboardSetLocalizationsAPI *GameCenterLeaderboardSetLocalizationsAPIService

	GameCenterLeaderboardSetMemberLocalizationsAPI *GameCenterLeaderboardSetMemberLocalizationsAPIService

	GameCenterLeaderboardSetReleasesAPI *GameCenterLeaderboardSetReleasesAPIService

	GameCenterLeaderboardSetsAPI *GameCenterLeaderboardSetsAPIService

	GameCenterLeaderboardsAPI *GameCenterLeaderboardsAPIService

	GameCenterMatchmakingQueuesAPI *GameCenterMatchmakingQueuesAPIService

	GameCenterMatchmakingRuleSetTestsAPI *GameCenterMatchmakingRuleSetTestsAPIService

	GameCenterMatchmakingRuleSetsAPI *GameCenterMatchmakingRuleSetsAPIService

	GameCenterMatchmakingRulesAPI *GameCenterMatchmakingRulesAPIService

	GameCenterMatchmakingTeamsAPI *GameCenterMatchmakingTeamsAPIService

	GameCenterPlayerAchievementSubmissionsAPI *GameCenterPlayerAchievementSubmissionsAPIService

	InAppPurchaseAppStoreReviewScreenshotsAPI *InAppPurchaseAppStoreReviewScreenshotsAPIService

	InAppPurchaseAvailabilitiesAPI *InAppPurchaseAvailabilitiesAPIService

	InAppPurchaseContentsAPI *InAppPurchaseContentsAPIService

	InAppPurchaseImagesAPI *InAppPurchaseImagesAPIService

	InAppPurchaseLocalizationsAPI *InAppPurchaseLocalizationsAPIService

	InAppPurchasePriceSchedulesAPI *InAppPurchasePriceSchedulesAPIService

	InAppPurchaseSubmissionsAPI *InAppPurchaseSubmissionsAPIService

	InAppPurchasesAPI *InAppPurchasesAPIService

	MarketplaceDomainsAPI *MarketplaceDomainsAPIService

	MarketplaceSearchDetailsAPI *MarketplaceSearchDetailsAPIService

	MarketplaceWebhooksAPI *MarketplaceWebhooksAPIService

	MetricsAPI *MetricsAPIService

	PreReleaseVersionsAPI *PreReleaseVersionsAPIService

	ProfilesAPI *ProfilesAPIService

	PromotedPurchaseImagesAPI *PromotedPurchaseImagesAPIService

	PromotedPurchasesAPI *PromotedPurchasesAPIService

	ReviewSubmissionItemsAPI *ReviewSubmissionItemsAPIService

	ReviewSubmissionsAPI *ReviewSubmissionsAPIService

	RoutingAppCoveragesAPI *RoutingAppCoveragesAPIService

	SalesReportsAPI *SalesReportsAPIService

	SandboxTestersAPI *SandboxTestersAPIService

	SandboxTestersClearPurchaseHistoryRequestAPI *SandboxTestersClearPurchaseHistoryRequestAPIService

	ScmGitReferencesAPI *ScmGitReferencesAPIService

	ScmProvidersAPI *ScmProvidersAPIService

	ScmPullRequestsAPI *ScmPullRequestsAPIService

	ScmRepositoriesAPI *ScmRepositoriesAPIService

	SubscriptionAppStoreReviewScreenshotsAPI *SubscriptionAppStoreReviewScreenshotsAPIService

	SubscriptionAvailabilitiesAPI *SubscriptionAvailabilitiesAPIService

	SubscriptionGracePeriodsAPI *SubscriptionGracePeriodsAPIService

	SubscriptionGroupLocalizationsAPI *SubscriptionGroupLocalizationsAPIService

	SubscriptionGroupSubmissionsAPI *SubscriptionGroupSubmissionsAPIService

	SubscriptionGroupsAPI *SubscriptionGroupsAPIService

	SubscriptionImagesAPI *SubscriptionImagesAPIService

	SubscriptionIntroductoryOffersAPI *SubscriptionIntroductoryOffersAPIService

	SubscriptionLocalizationsAPI *SubscriptionLocalizationsAPIService

	SubscriptionOfferCodeCustomCodesAPI *SubscriptionOfferCodeCustomCodesAPIService

	SubscriptionOfferCodeOneTimeUseCodesAPI *SubscriptionOfferCodeOneTimeUseCodesAPIService

	SubscriptionOfferCodesAPI *SubscriptionOfferCodesAPIService

	SubscriptionPricePointsAPI *SubscriptionPricePointsAPIService

	SubscriptionPricesAPI *SubscriptionPricesAPIService

	SubscriptionPromotionalOffersAPI *SubscriptionPromotionalOffersAPIService

	SubscriptionSubmissionsAPI *SubscriptionSubmissionsAPIService

	SubscriptionsAPI *SubscriptionsAPIService

	TerritoriesAPI *TerritoriesAPIService

	TerritoryAvailabilitiesAPI *TerritoryAvailabilitiesAPIService

	UserInvitationsAPI *UserInvitationsAPIService

	UsersAPI *UsersAPIService

	WinBackOffersAPI *WinBackOffersAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.ActorsAPI = (*ActorsAPIService)(&c.common)
	c.AgeRatingDeclarationsAPI = (*AgeRatingDeclarationsAPIService)(&c.common)
	c.AlternativeDistributionDomainsAPI = (*AlternativeDistributionDomainsAPIService)(&c.common)
	c.AlternativeDistributionKeysAPI = (*AlternativeDistributionKeysAPIService)(&c.common)
	c.AlternativeDistributionPackageDeltasAPI = (*AlternativeDistributionPackageDeltasAPIService)(&c.common)
	c.AlternativeDistributionPackageVariantsAPI = (*AlternativeDistributionPackageVariantsAPIService)(&c.common)
	c.AlternativeDistributionPackageVersionsAPI = (*AlternativeDistributionPackageVersionsAPIService)(&c.common)
	c.AlternativeDistributionPackagesAPI = (*AlternativeDistributionPackagesAPIService)(&c.common)
	c.AnalyticsReportInstancesAPI = (*AnalyticsReportInstancesAPIService)(&c.common)
	c.AnalyticsReportRequestsAPI = (*AnalyticsReportRequestsAPIService)(&c.common)
	c.AnalyticsReportSegmentsAPI = (*AnalyticsReportSegmentsAPIService)(&c.common)
	c.AnalyticsReportsAPI = (*AnalyticsReportsAPIService)(&c.common)
	c.AppAvailabilitiesAPI = (*AppAvailabilitiesAPIService)(&c.common)
	c.AppCategoriesAPI = (*AppCategoriesAPIService)(&c.common)
	c.AppClipAdvancedExperienceImagesAPI = (*AppClipAdvancedExperienceImagesAPIService)(&c.common)
	c.AppClipAdvancedExperiencesAPI = (*AppClipAdvancedExperiencesAPIService)(&c.common)
	c.AppClipAppStoreReviewDetailsAPI = (*AppClipAppStoreReviewDetailsAPIService)(&c.common)
	c.AppClipDefaultExperienceLocalizationsAPI = (*AppClipDefaultExperienceLocalizationsAPIService)(&c.common)
	c.AppClipDefaultExperiencesAPI = (*AppClipDefaultExperiencesAPIService)(&c.common)
	c.AppClipHeaderImagesAPI = (*AppClipHeaderImagesAPIService)(&c.common)
	c.AppClipsAPI = (*AppClipsAPIService)(&c.common)
	c.AppCustomProductPageLocalizationsAPI = (*AppCustomProductPageLocalizationsAPIService)(&c.common)
	c.AppCustomProductPageVersionsAPI = (*AppCustomProductPageVersionsAPIService)(&c.common)
	c.AppCustomProductPagesAPI = (*AppCustomProductPagesAPIService)(&c.common)
	c.AppEncryptionDeclarationDocumentsAPI = (*AppEncryptionDeclarationDocumentsAPIService)(&c.common)
	c.AppEncryptionDeclarationsAPI = (*AppEncryptionDeclarationsAPIService)(&c.common)
	c.AppEventLocalizationsAPI = (*AppEventLocalizationsAPIService)(&c.common)
	c.AppEventScreenshotsAPI = (*AppEventScreenshotsAPIService)(&c.common)
	c.AppEventVideoClipsAPI = (*AppEventVideoClipsAPIService)(&c.common)
	c.AppEventsAPI = (*AppEventsAPIService)(&c.common)
	c.AppInfoLocalizationsAPI = (*AppInfoLocalizationsAPIService)(&c.common)
	c.AppInfosAPI = (*AppInfosAPIService)(&c.common)
	c.AppPreOrdersAPI = (*AppPreOrdersAPIService)(&c.common)
	c.AppPreviewSetsAPI = (*AppPreviewSetsAPIService)(&c.common)
	c.AppPreviewsAPI = (*AppPreviewsAPIService)(&c.common)
	c.AppPricePointsAPI = (*AppPricePointsAPIService)(&c.common)
	c.AppPriceSchedulesAPI = (*AppPriceSchedulesAPIService)(&c.common)
	c.AppScreenshotSetsAPI = (*AppScreenshotSetsAPIService)(&c.common)
	c.AppScreenshotsAPI = (*AppScreenshotsAPIService)(&c.common)
	c.AppStoreReviewAttachmentsAPI = (*AppStoreReviewAttachmentsAPIService)(&c.common)
	c.AppStoreReviewDetailsAPI = (*AppStoreReviewDetailsAPIService)(&c.common)
	c.AppStoreVersionExperimentTreatmentLocalizationsAPI = (*AppStoreVersionExperimentTreatmentLocalizationsAPIService)(&c.common)
	c.AppStoreVersionExperimentTreatmentsAPI = (*AppStoreVersionExperimentTreatmentsAPIService)(&c.common)
	c.AppStoreVersionExperimentsAPI = (*AppStoreVersionExperimentsAPIService)(&c.common)
	c.AppStoreVersionLocalizationsAPI = (*AppStoreVersionLocalizationsAPIService)(&c.common)
	c.AppStoreVersionPhasedReleasesAPI = (*AppStoreVersionPhasedReleasesAPIService)(&c.common)
	c.AppStoreVersionPromotionsAPI = (*AppStoreVersionPromotionsAPIService)(&c.common)
	c.AppStoreVersionReleaseRequestsAPI = (*AppStoreVersionReleaseRequestsAPIService)(&c.common)
	c.AppStoreVersionSubmissionsAPI = (*AppStoreVersionSubmissionsAPIService)(&c.common)
	c.AppStoreVersionsAPI = (*AppStoreVersionsAPIService)(&c.common)
	c.AppsAPI = (*AppsAPIService)(&c.common)
	c.BetaAppClipInvocationLocalizationsAPI = (*BetaAppClipInvocationLocalizationsAPIService)(&c.common)
	c.BetaAppClipInvocationsAPI = (*BetaAppClipInvocationsAPIService)(&c.common)
	c.BetaAppLocalizationsAPI = (*BetaAppLocalizationsAPIService)(&c.common)
	c.BetaAppReviewDetailsAPI = (*BetaAppReviewDetailsAPIService)(&c.common)
	c.BetaAppReviewSubmissionsAPI = (*BetaAppReviewSubmissionsAPIService)(&c.common)
	c.BetaBuildLocalizationsAPI = (*BetaBuildLocalizationsAPIService)(&c.common)
	c.BetaGroupsAPI = (*BetaGroupsAPIService)(&c.common)
	c.BetaLicenseAgreementsAPI = (*BetaLicenseAgreementsAPIService)(&c.common)
	c.BetaTesterInvitationsAPI = (*BetaTesterInvitationsAPIService)(&c.common)
	c.BetaTestersAPI = (*BetaTestersAPIService)(&c.common)
	c.BuildBetaDetailsAPI = (*BuildBetaDetailsAPIService)(&c.common)
	c.BuildBetaNotificationsAPI = (*BuildBetaNotificationsAPIService)(&c.common)
	c.BuildBundlesAPI = (*BuildBundlesAPIService)(&c.common)
	c.BuildsAPI = (*BuildsAPIService)(&c.common)
	c.BundleIdCapabilitiesAPI = (*BundleIdCapabilitiesAPIService)(&c.common)
	c.BundleIdsAPI = (*BundleIdsAPIService)(&c.common)
	c.CertificatesAPI = (*CertificatesAPIService)(&c.common)
	c.CiArtifactsAPI = (*CiArtifactsAPIService)(&c.common)
	c.CiBuildActionsAPI = (*CiBuildActionsAPIService)(&c.common)
	c.CiBuildRunsAPI = (*CiBuildRunsAPIService)(&c.common)
	c.CiIssuesAPI = (*CiIssuesAPIService)(&c.common)
	c.CiMacOsVersionsAPI = (*CiMacOsVersionsAPIService)(&c.common)
	c.CiProductsAPI = (*CiProductsAPIService)(&c.common)
	c.CiTestResultsAPI = (*CiTestResultsAPIService)(&c.common)
	c.CiWorkflowsAPI = (*CiWorkflowsAPIService)(&c.common)
	c.CiXcodeVersionsAPI = (*CiXcodeVersionsAPIService)(&c.common)
	c.CustomerReviewResponsesAPI = (*CustomerReviewResponsesAPIService)(&c.common)
	c.CustomerReviewsAPI = (*CustomerReviewsAPIService)(&c.common)
	c.DevicesAPI = (*DevicesAPIService)(&c.common)
	c.DiagnosticSignaturesAPI = (*DiagnosticSignaturesAPIService)(&c.common)
	c.EndAppAvailabilityPreOrdersAPI = (*EndAppAvailabilityPreOrdersAPIService)(&c.common)
	c.EndUserLicenseAgreementsAPI = (*EndUserLicenseAgreementsAPIService)(&c.common)
	c.FinanceReportsAPI = (*FinanceReportsAPIService)(&c.common)
	c.GameCenterAchievementImagesAPI = (*GameCenterAchievementImagesAPIService)(&c.common)
	c.GameCenterAchievementLocalizationsAPI = (*GameCenterAchievementLocalizationsAPIService)(&c.common)
	c.GameCenterAchievementReleasesAPI = (*GameCenterAchievementReleasesAPIService)(&c.common)
	c.GameCenterAchievementsAPI = (*GameCenterAchievementsAPIService)(&c.common)
	c.GameCenterAppVersionsAPI = (*GameCenterAppVersionsAPIService)(&c.common)
	c.GameCenterDetailsAPI = (*GameCenterDetailsAPIService)(&c.common)
	c.GameCenterEnabledVersionsAPI = (*GameCenterEnabledVersionsAPIService)(&c.common)
	c.GameCenterGroupsAPI = (*GameCenterGroupsAPIService)(&c.common)
	c.GameCenterLeaderboardEntrySubmissionsAPI = (*GameCenterLeaderboardEntrySubmissionsAPIService)(&c.common)
	c.GameCenterLeaderboardImagesAPI = (*GameCenterLeaderboardImagesAPIService)(&c.common)
	c.GameCenterLeaderboardLocalizationsAPI = (*GameCenterLeaderboardLocalizationsAPIService)(&c.common)
	c.GameCenterLeaderboardReleasesAPI = (*GameCenterLeaderboardReleasesAPIService)(&c.common)
	c.GameCenterLeaderboardSetImagesAPI = (*GameCenterLeaderboardSetImagesAPIService)(&c.common)
	c.GameCenterLeaderboardSetLocalizationsAPI = (*GameCenterLeaderboardSetLocalizationsAPIService)(&c.common)
	c.GameCenterLeaderboardSetMemberLocalizationsAPI = (*GameCenterLeaderboardSetMemberLocalizationsAPIService)(&c.common)
	c.GameCenterLeaderboardSetReleasesAPI = (*GameCenterLeaderboardSetReleasesAPIService)(&c.common)
	c.GameCenterLeaderboardSetsAPI = (*GameCenterLeaderboardSetsAPIService)(&c.common)
	c.GameCenterLeaderboardsAPI = (*GameCenterLeaderboardsAPIService)(&c.common)
	c.GameCenterMatchmakingQueuesAPI = (*GameCenterMatchmakingQueuesAPIService)(&c.common)
	c.GameCenterMatchmakingRuleSetTestsAPI = (*GameCenterMatchmakingRuleSetTestsAPIService)(&c.common)
	c.GameCenterMatchmakingRuleSetsAPI = (*GameCenterMatchmakingRuleSetsAPIService)(&c.common)
	c.GameCenterMatchmakingRulesAPI = (*GameCenterMatchmakingRulesAPIService)(&c.common)
	c.GameCenterMatchmakingTeamsAPI = (*GameCenterMatchmakingTeamsAPIService)(&c.common)
	c.GameCenterPlayerAchievementSubmissionsAPI = (*GameCenterPlayerAchievementSubmissionsAPIService)(&c.common)
	c.InAppPurchaseAppStoreReviewScreenshotsAPI = (*InAppPurchaseAppStoreReviewScreenshotsAPIService)(&c.common)
	c.InAppPurchaseAvailabilitiesAPI = (*InAppPurchaseAvailabilitiesAPIService)(&c.common)
	c.InAppPurchaseContentsAPI = (*InAppPurchaseContentsAPIService)(&c.common)
	c.InAppPurchaseImagesAPI = (*InAppPurchaseImagesAPIService)(&c.common)
	c.InAppPurchaseLocalizationsAPI = (*InAppPurchaseLocalizationsAPIService)(&c.common)
	c.InAppPurchasePriceSchedulesAPI = (*InAppPurchasePriceSchedulesAPIService)(&c.common)
	c.InAppPurchaseSubmissionsAPI = (*InAppPurchaseSubmissionsAPIService)(&c.common)
	c.InAppPurchasesAPI = (*InAppPurchasesAPIService)(&c.common)
	c.MarketplaceDomainsAPI = (*MarketplaceDomainsAPIService)(&c.common)
	c.MarketplaceSearchDetailsAPI = (*MarketplaceSearchDetailsAPIService)(&c.common)
	c.MarketplaceWebhooksAPI = (*MarketplaceWebhooksAPIService)(&c.common)
	c.MetricsAPI = (*MetricsAPIService)(&c.common)
	c.PreReleaseVersionsAPI = (*PreReleaseVersionsAPIService)(&c.common)
	c.ProfilesAPI = (*ProfilesAPIService)(&c.common)
	c.PromotedPurchaseImagesAPI = (*PromotedPurchaseImagesAPIService)(&c.common)
	c.PromotedPurchasesAPI = (*PromotedPurchasesAPIService)(&c.common)
	c.ReviewSubmissionItemsAPI = (*ReviewSubmissionItemsAPIService)(&c.common)
	c.ReviewSubmissionsAPI = (*ReviewSubmissionsAPIService)(&c.common)
	c.RoutingAppCoveragesAPI = (*RoutingAppCoveragesAPIService)(&c.common)
	c.SalesReportsAPI = (*SalesReportsAPIService)(&c.common)
	c.SandboxTestersAPI = (*SandboxTestersAPIService)(&c.common)
	c.SandboxTestersClearPurchaseHistoryRequestAPI = (*SandboxTestersClearPurchaseHistoryRequestAPIService)(&c.common)
	c.ScmGitReferencesAPI = (*ScmGitReferencesAPIService)(&c.common)
	c.ScmProvidersAPI = (*ScmProvidersAPIService)(&c.common)
	c.ScmPullRequestsAPI = (*ScmPullRequestsAPIService)(&c.common)
	c.ScmRepositoriesAPI = (*ScmRepositoriesAPIService)(&c.common)
	c.SubscriptionAppStoreReviewScreenshotsAPI = (*SubscriptionAppStoreReviewScreenshotsAPIService)(&c.common)
	c.SubscriptionAvailabilitiesAPI = (*SubscriptionAvailabilitiesAPIService)(&c.common)
	c.SubscriptionGracePeriodsAPI = (*SubscriptionGracePeriodsAPIService)(&c.common)
	c.SubscriptionGroupLocalizationsAPI = (*SubscriptionGroupLocalizationsAPIService)(&c.common)
	c.SubscriptionGroupSubmissionsAPI = (*SubscriptionGroupSubmissionsAPIService)(&c.common)
	c.SubscriptionGroupsAPI = (*SubscriptionGroupsAPIService)(&c.common)
	c.SubscriptionImagesAPI = (*SubscriptionImagesAPIService)(&c.common)
	c.SubscriptionIntroductoryOffersAPI = (*SubscriptionIntroductoryOffersAPIService)(&c.common)
	c.SubscriptionLocalizationsAPI = (*SubscriptionLocalizationsAPIService)(&c.common)
	c.SubscriptionOfferCodeCustomCodesAPI = (*SubscriptionOfferCodeCustomCodesAPIService)(&c.common)
	c.SubscriptionOfferCodeOneTimeUseCodesAPI = (*SubscriptionOfferCodeOneTimeUseCodesAPIService)(&c.common)
	c.SubscriptionOfferCodesAPI = (*SubscriptionOfferCodesAPIService)(&c.common)
	c.SubscriptionPricePointsAPI = (*SubscriptionPricePointsAPIService)(&c.common)
	c.SubscriptionPricesAPI = (*SubscriptionPricesAPIService)(&c.common)
	c.SubscriptionPromotionalOffersAPI = (*SubscriptionPromotionalOffersAPIService)(&c.common)
	c.SubscriptionSubmissionsAPI = (*SubscriptionSubmissionsAPIService)(&c.common)
	c.SubscriptionsAPI = (*SubscriptionsAPIService)(&c.common)
	c.TerritoriesAPI = (*TerritoriesAPIService)(&c.common)
	c.TerritoryAvailabilitiesAPI = (*TerritoryAvailabilitiesAPIService)(&c.common)
	c.UserInvitationsAPI = (*UserInvitationsAPIService)(&c.common)
	c.UsersAPI = (*UsersAPIService)(&c.common)
	c.WinBackOffersAPI = (*WinBackOffersAPIService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString( obj interface{}, key string ) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", obj)
	}
	var param,ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap,err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
			case reflect.Invalid:
				value = "invalid"

			case reflect.Struct:
				if t,ok := obj.(MappedNullable); ok {
					dataMap,err := t.ToMap()
					if err != nil {
						return
					}
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
					return
				}
				if t, ok := obj.(time.Time); ok {
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
					return
				}
				value = v.Type().String() + " value"
			case reflect.Slice:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				var lenIndValue = indValue.Len()
				for i:=0;i<lenIndValue;i++ {
					var arrayValue = indValue.Index(i)
					var keyPrefixForCollectionType = keyPrefix
					if style == "deepObject" {
						keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
					}
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
				}
				return

			case reflect.Map:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				iter := indValue.MapRange()
				for iter.Next() {
					k,v := iter.Key(), iter.Value()
					parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
				}
				return

			case reflect.Interface:
				fallthrough
			case reflect.Ptr:
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
				return

			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(v.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				value = strconv.FormatUint(v.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
			case reflect.Bool:
				value = strconv.FormatBool(v.Bool())
			case reflect.String:
				value = v.String()
			default:
				value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
		case url.Values:
			if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
				valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix) + "," + value)
			} else {
				valuesMap.Add(keyPrefix, value)
			}
			break
		case map[string]string:
			valuesMap[keyPrefix] = value
			break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
		fileBytes []byte
		fileName string
		formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
						return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
						return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}