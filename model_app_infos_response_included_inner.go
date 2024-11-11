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

// AppInfosResponseIncludedInner - struct for AppInfosResponseIncludedInner
type AppInfosResponseIncludedInner struct {
	AgeRatingDeclaration *AgeRatingDeclaration
	App *App
	AppCategory *AppCategory
	AppInfoLocalization *AppInfoLocalization
}

// AgeRatingDeclarationAsAppInfosResponseIncludedInner is a convenience function that returns AgeRatingDeclaration wrapped in AppInfosResponseIncludedInner
func AgeRatingDeclarationAsAppInfosResponseIncludedInner(v *AgeRatingDeclaration) AppInfosResponseIncludedInner {
	return AppInfosResponseIncludedInner{
		AgeRatingDeclaration: v,
	}
}

// AppAsAppInfosResponseIncludedInner is a convenience function that returns App wrapped in AppInfosResponseIncludedInner
func AppAsAppInfosResponseIncludedInner(v *App) AppInfosResponseIncludedInner {
	return AppInfosResponseIncludedInner{
		App: v,
	}
}

// AppCategoryAsAppInfosResponseIncludedInner is a convenience function that returns AppCategory wrapped in AppInfosResponseIncludedInner
func AppCategoryAsAppInfosResponseIncludedInner(v *AppCategory) AppInfosResponseIncludedInner {
	return AppInfosResponseIncludedInner{
		AppCategory: v,
	}
}

// AppInfoLocalizationAsAppInfosResponseIncludedInner is a convenience function that returns AppInfoLocalization wrapped in AppInfosResponseIncludedInner
func AppInfoLocalizationAsAppInfosResponseIncludedInner(v *AppInfoLocalization) AppInfosResponseIncludedInner {
	return AppInfosResponseIncludedInner{
		AppInfoLocalization: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppInfosResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AgeRatingDeclaration
	err = newStrictDecoder(data).Decode(&dst.AgeRatingDeclaration)
	if err == nil {
		jsonAgeRatingDeclaration, _ := json.Marshal(dst.AgeRatingDeclaration)
		if string(jsonAgeRatingDeclaration) == "{}" { // empty struct
			dst.AgeRatingDeclaration = nil
		} else {
			if err = validator.Validate(dst.AgeRatingDeclaration); err != nil {
				dst.AgeRatingDeclaration = nil
			} else {
				match++
			}
		}
	} else {
		dst.AgeRatingDeclaration = nil
	}

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

	// try to unmarshal data into AppCategory
	err = newStrictDecoder(data).Decode(&dst.AppCategory)
	if err == nil {
		jsonAppCategory, _ := json.Marshal(dst.AppCategory)
		if string(jsonAppCategory) == "{}" { // empty struct
			dst.AppCategory = nil
		} else {
			if err = validator.Validate(dst.AppCategory); err != nil {
				dst.AppCategory = nil
			} else {
				match++
			}
		}
	} else {
		dst.AppCategory = nil
	}

	// try to unmarshal data into AppInfoLocalization
	err = newStrictDecoder(data).Decode(&dst.AppInfoLocalization)
	if err == nil {
		jsonAppInfoLocalization, _ := json.Marshal(dst.AppInfoLocalization)
		if string(jsonAppInfoLocalization) == "{}" { // empty struct
			dst.AppInfoLocalization = nil
		} else {
			if err = validator.Validate(dst.AppInfoLocalization); err != nil {
				dst.AppInfoLocalization = nil
			} else {
				match++
			}
		}
	} else {
		dst.AppInfoLocalization = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AgeRatingDeclaration = nil
		dst.App = nil
		dst.AppCategory = nil
		dst.AppInfoLocalization = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppInfosResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppInfosResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppInfosResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AgeRatingDeclaration != nil {
		return json.Marshal(&src.AgeRatingDeclaration)
	}

	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.AppCategory != nil {
		return json.Marshal(&src.AppCategory)
	}

	if src.AppInfoLocalization != nil {
		return json.Marshal(&src.AppInfoLocalization)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppInfosResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AgeRatingDeclaration != nil {
		return obj.AgeRatingDeclaration
	}

	if obj.App != nil {
		return obj.App
	}

	if obj.AppCategory != nil {
		return obj.AppCategory
	}

	if obj.AppInfoLocalization != nil {
		return obj.AppInfoLocalization
	}

	// all schemas are nil
	return nil
}

type NullableAppInfosResponseIncludedInner struct {
	value *AppInfosResponseIncludedInner
	isSet bool
}

func (v NullableAppInfosResponseIncludedInner) Get() *AppInfosResponseIncludedInner {
	return v.value
}

func (v *NullableAppInfosResponseIncludedInner) Set(val *AppInfosResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfosResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfosResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfosResponseIncludedInner(val *AppInfosResponseIncludedInner) *NullableAppInfosResponseIncludedInner {
	return &NullableAppInfosResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppInfosResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfosResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


