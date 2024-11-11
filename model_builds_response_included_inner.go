/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// BuildsResponseIncludedInner - struct for BuildsResponseIncludedInner
type BuildsResponseIncludedInner struct {
	App *App
	AppEncryptionDeclaration *AppEncryptionDeclaration
	AppStoreVersion *AppStoreVersion
	BetaAppReviewSubmission *BetaAppReviewSubmission
	BetaBuildLocalization *BetaBuildLocalization
	BetaGroup *BetaGroup
	BetaTester *BetaTester
	BuildBetaDetail *BuildBetaDetail
	BuildBundle *BuildBundle
	BuildIcon *BuildIcon
	PrereleaseVersion *PrereleaseVersion
}

// AppAsBuildsResponseIncludedInner is a convenience function that returns App wrapped in BuildsResponseIncludedInner
func AppAsBuildsResponseIncludedInner(v *App) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		App: v,
	}
}

// AppEncryptionDeclarationAsBuildsResponseIncludedInner is a convenience function that returns AppEncryptionDeclaration wrapped in BuildsResponseIncludedInner
func AppEncryptionDeclarationAsBuildsResponseIncludedInner(v *AppEncryptionDeclaration) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		AppEncryptionDeclaration: v,
	}
}

// AppStoreVersionAsBuildsResponseIncludedInner is a convenience function that returns AppStoreVersion wrapped in BuildsResponseIncludedInner
func AppStoreVersionAsBuildsResponseIncludedInner(v *AppStoreVersion) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		AppStoreVersion: v,
	}
}

// BetaAppReviewSubmissionAsBuildsResponseIncludedInner is a convenience function that returns BetaAppReviewSubmission wrapped in BuildsResponseIncludedInner
func BetaAppReviewSubmissionAsBuildsResponseIncludedInner(v *BetaAppReviewSubmission) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		BetaAppReviewSubmission: v,
	}
}

// BetaBuildLocalizationAsBuildsResponseIncludedInner is a convenience function that returns BetaBuildLocalization wrapped in BuildsResponseIncludedInner
func BetaBuildLocalizationAsBuildsResponseIncludedInner(v *BetaBuildLocalization) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		BetaBuildLocalization: v,
	}
}

// BetaGroupAsBuildsResponseIncludedInner is a convenience function that returns BetaGroup wrapped in BuildsResponseIncludedInner
func BetaGroupAsBuildsResponseIncludedInner(v *BetaGroup) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		BetaGroup: v,
	}
}

// BetaTesterAsBuildsResponseIncludedInner is a convenience function that returns BetaTester wrapped in BuildsResponseIncludedInner
func BetaTesterAsBuildsResponseIncludedInner(v *BetaTester) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		BetaTester: v,
	}
}

// BuildBetaDetailAsBuildsResponseIncludedInner is a convenience function that returns BuildBetaDetail wrapped in BuildsResponseIncludedInner
func BuildBetaDetailAsBuildsResponseIncludedInner(v *BuildBetaDetail) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		BuildBetaDetail: v,
	}
}

// BuildBundleAsBuildsResponseIncludedInner is a convenience function that returns BuildBundle wrapped in BuildsResponseIncludedInner
func BuildBundleAsBuildsResponseIncludedInner(v *BuildBundle) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		BuildBundle: v,
	}
}

// BuildIconAsBuildsResponseIncludedInner is a convenience function that returns BuildIcon wrapped in BuildsResponseIncludedInner
func BuildIconAsBuildsResponseIncludedInner(v *BuildIcon) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		BuildIcon: v,
	}
}

// PrereleaseVersionAsBuildsResponseIncludedInner is a convenience function that returns PrereleaseVersion wrapped in BuildsResponseIncludedInner
func PrereleaseVersionAsBuildsResponseIncludedInner(v *PrereleaseVersion) BuildsResponseIncludedInner {
	return BuildsResponseIncludedInner{
		PrereleaseVersion: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BuildsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into App
	err = newStrictDecoder(data).Decode(&dst.App)
	if err == nil {
		jsonApp, _ := json.Marshal(dst.App)
		if string(jsonApp) == "{}" { // empty struct
			dst.App = nil
		} else {
			if err = validator.Validate(dst.App); err != nil {
				dst.App = nil
			} else {
				match++
			}
		}
	} else {
		dst.App = nil
	}

	// try to unmarshal data into AppEncryptionDeclaration
	err = newStrictDecoder(data).Decode(&dst.AppEncryptionDeclaration)
	if err == nil {
		jsonAppEncryptionDeclaration, _ := json.Marshal(dst.AppEncryptionDeclaration)
		if string(jsonAppEncryptionDeclaration) == "{}" { // empty struct
			dst.AppEncryptionDeclaration = nil
		} else {
			if err = validator.Validate(dst.AppEncryptionDeclaration); err != nil {
				dst.AppEncryptionDeclaration = nil
			} else {
				match++
			}
		}
	} else {
		dst.AppEncryptionDeclaration = nil
	}

	// try to unmarshal data into AppStoreVersion
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersion)
	if err == nil {
		jsonAppStoreVersion, _ := json.Marshal(dst.AppStoreVersion)
		if string(jsonAppStoreVersion) == "{}" { // empty struct
			dst.AppStoreVersion = nil
		} else {
			if err = validator.Validate(dst.AppStoreVersion); err != nil {
				dst.AppStoreVersion = nil
			} else {
				match++
			}
		}
	} else {
		dst.AppStoreVersion = nil
	}

	// try to unmarshal data into BetaAppReviewSubmission
	err = newStrictDecoder(data).Decode(&dst.BetaAppReviewSubmission)
	if err == nil {
		jsonBetaAppReviewSubmission, _ := json.Marshal(dst.BetaAppReviewSubmission)
		if string(jsonBetaAppReviewSubmission) == "{}" { // empty struct
			dst.BetaAppReviewSubmission = nil
		} else {
			if err = validator.Validate(dst.BetaAppReviewSubmission); err != nil {
				dst.BetaAppReviewSubmission = nil
			} else {
				match++
			}
		}
	} else {
		dst.BetaAppReviewSubmission = nil
	}

	// try to unmarshal data into BetaBuildLocalization
	err = newStrictDecoder(data).Decode(&dst.BetaBuildLocalization)
	if err == nil {
		jsonBetaBuildLocalization, _ := json.Marshal(dst.BetaBuildLocalization)
		if string(jsonBetaBuildLocalization) == "{}" { // empty struct
			dst.BetaBuildLocalization = nil
		} else {
			if err = validator.Validate(dst.BetaBuildLocalization); err != nil {
				dst.BetaBuildLocalization = nil
			} else {
				match++
			}
		}
	} else {
		dst.BetaBuildLocalization = nil
	}

	// try to unmarshal data into BetaGroup
	err = newStrictDecoder(data).Decode(&dst.BetaGroup)
	if err == nil {
		jsonBetaGroup, _ := json.Marshal(dst.BetaGroup)
		if string(jsonBetaGroup) == "{}" { // empty struct
			dst.BetaGroup = nil
		} else {
			if err = validator.Validate(dst.BetaGroup); err != nil {
				dst.BetaGroup = nil
			} else {
				match++
			}
		}
	} else {
		dst.BetaGroup = nil
	}

	// try to unmarshal data into BetaTester
	err = newStrictDecoder(data).Decode(&dst.BetaTester)
	if err == nil {
		jsonBetaTester, _ := json.Marshal(dst.BetaTester)
		if string(jsonBetaTester) == "{}" { // empty struct
			dst.BetaTester = nil
		} else {
			if err = validator.Validate(dst.BetaTester); err != nil {
				dst.BetaTester = nil
			} else {
				match++
			}
		}
	} else {
		dst.BetaTester = nil
	}

	// try to unmarshal data into BuildBetaDetail
	err = newStrictDecoder(data).Decode(&dst.BuildBetaDetail)
	if err == nil {
		jsonBuildBetaDetail, _ := json.Marshal(dst.BuildBetaDetail)
		if string(jsonBuildBetaDetail) == "{}" { // empty struct
			dst.BuildBetaDetail = nil
		} else {
			if err = validator.Validate(dst.BuildBetaDetail); err != nil {
				dst.BuildBetaDetail = nil
			} else {
				match++
			}
		}
	} else {
		dst.BuildBetaDetail = nil
	}

	// try to unmarshal data into BuildBundle
	err = newStrictDecoder(data).Decode(&dst.BuildBundle)
	if err == nil {
		jsonBuildBundle, _ := json.Marshal(dst.BuildBundle)
		if string(jsonBuildBundle) == "{}" { // empty struct
			dst.BuildBundle = nil
		} else {
			if err = validator.Validate(dst.BuildBundle); err != nil {
				dst.BuildBundle = nil
			} else {
				match++
			}
		}
	} else {
		dst.BuildBundle = nil
	}

	// try to unmarshal data into BuildIcon
	err = newStrictDecoder(data).Decode(&dst.BuildIcon)
	if err == nil {
		jsonBuildIcon, _ := json.Marshal(dst.BuildIcon)
		if string(jsonBuildIcon) == "{}" { // empty struct
			dst.BuildIcon = nil
		} else {
			if err = validator.Validate(dst.BuildIcon); err != nil {
				dst.BuildIcon = nil
			} else {
				match++
			}
		}
	} else {
		dst.BuildIcon = nil
	}

	// try to unmarshal data into PrereleaseVersion
	err = newStrictDecoder(data).Decode(&dst.PrereleaseVersion)
	if err == nil {
		jsonPrereleaseVersion, _ := json.Marshal(dst.PrereleaseVersion)
		if string(jsonPrereleaseVersion) == "{}" { // empty struct
			dst.PrereleaseVersion = nil
		} else {
			if err = validator.Validate(dst.PrereleaseVersion); err != nil {
				dst.PrereleaseVersion = nil
			} else {
				match++
			}
		}
	} else {
		dst.PrereleaseVersion = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.App = nil
		dst.AppEncryptionDeclaration = nil
		dst.AppStoreVersion = nil
		dst.BetaAppReviewSubmission = nil
		dst.BetaBuildLocalization = nil
		dst.BetaGroup = nil
		dst.BetaTester = nil
		dst.BuildBetaDetail = nil
		dst.BuildBundle = nil
		dst.BuildIcon = nil
		dst.PrereleaseVersion = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BuildsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BuildsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BuildsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.AppEncryptionDeclaration != nil {
		return json.Marshal(&src.AppEncryptionDeclaration)
	}

	if src.AppStoreVersion != nil {
		return json.Marshal(&src.AppStoreVersion)
	}

	if src.BetaAppReviewSubmission != nil {
		return json.Marshal(&src.BetaAppReviewSubmission)
	}

	if src.BetaBuildLocalization != nil {
		return json.Marshal(&src.BetaBuildLocalization)
	}

	if src.BetaGroup != nil {
		return json.Marshal(&src.BetaGroup)
	}

	if src.BetaTester != nil {
		return json.Marshal(&src.BetaTester)
	}

	if src.BuildBetaDetail != nil {
		return json.Marshal(&src.BuildBetaDetail)
	}

	if src.BuildBundle != nil {
		return json.Marshal(&src.BuildBundle)
	}

	if src.BuildIcon != nil {
		return json.Marshal(&src.BuildIcon)
	}

	if src.PrereleaseVersion != nil {
		return json.Marshal(&src.PrereleaseVersion)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BuildsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.App != nil {
		return obj.App
	}

	if obj.AppEncryptionDeclaration != nil {
		return obj.AppEncryptionDeclaration
	}

	if obj.AppStoreVersion != nil {
		return obj.AppStoreVersion
	}

	if obj.BetaAppReviewSubmission != nil {
		return obj.BetaAppReviewSubmission
	}

	if obj.BetaBuildLocalization != nil {
		return obj.BetaBuildLocalization
	}

	if obj.BetaGroup != nil {
		return obj.BetaGroup
	}

	if obj.BetaTester != nil {
		return obj.BetaTester
	}

	if obj.BuildBetaDetail != nil {
		return obj.BuildBetaDetail
	}

	if obj.BuildBundle != nil {
		return obj.BuildBundle
	}

	if obj.BuildIcon != nil {
		return obj.BuildIcon
	}

	if obj.PrereleaseVersion != nil {
		return obj.PrereleaseVersion
	}

	// all schemas are nil
	return nil
}

type NullableBuildsResponseIncludedInner struct {
	value *BuildsResponseIncludedInner
	isSet bool
}

func (v NullableBuildsResponseIncludedInner) Get() *BuildsResponseIncludedInner {
	return v.value
}

func (v *NullableBuildsResponseIncludedInner) Set(val *BuildsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildsResponseIncludedInner(val *BuildsResponseIncludedInner) *NullableBuildsResponseIncludedInner {
	return &NullableBuildsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableBuildsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


