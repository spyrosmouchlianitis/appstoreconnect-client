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

// AppStoreVersionExperimentsV2ResponseIncludedInner - struct for AppStoreVersionExperimentsV2ResponseIncludedInner
type AppStoreVersionExperimentsV2ResponseIncludedInner struct {
	App *App
	AppStoreVersion *AppStoreVersion
	AppStoreVersionExperimentTreatment *AppStoreVersionExperimentTreatment
}

// AppAsAppStoreVersionExperimentsV2ResponseIncludedInner is a convenience function that returns App wrapped in AppStoreVersionExperimentsV2ResponseIncludedInner
func AppAsAppStoreVersionExperimentsV2ResponseIncludedInner(v *App) AppStoreVersionExperimentsV2ResponseIncludedInner {
	return AppStoreVersionExperimentsV2ResponseIncludedInner{
		App: v,
	}
}

// AppStoreVersionAsAppStoreVersionExperimentsV2ResponseIncludedInner is a convenience function that returns AppStoreVersion wrapped in AppStoreVersionExperimentsV2ResponseIncludedInner
func AppStoreVersionAsAppStoreVersionExperimentsV2ResponseIncludedInner(v *AppStoreVersion) AppStoreVersionExperimentsV2ResponseIncludedInner {
	return AppStoreVersionExperimentsV2ResponseIncludedInner{
		AppStoreVersion: v,
	}
}

// AppStoreVersionExperimentTreatmentAsAppStoreVersionExperimentsV2ResponseIncludedInner is a convenience function that returns AppStoreVersionExperimentTreatment wrapped in AppStoreVersionExperimentsV2ResponseIncludedInner
func AppStoreVersionExperimentTreatmentAsAppStoreVersionExperimentsV2ResponseIncludedInner(v *AppStoreVersionExperimentTreatment) AppStoreVersionExperimentsV2ResponseIncludedInner {
	return AppStoreVersionExperimentsV2ResponseIncludedInner{
		AppStoreVersionExperimentTreatment: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppStoreVersionExperimentsV2ResponseIncludedInner) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into AppStoreVersionExperimentTreatment
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperimentTreatment)
	if err == nil {
		jsonAppStoreVersionExperimentTreatment, _ := json.Marshal(dst.AppStoreVersionExperimentTreatment)
		if string(jsonAppStoreVersionExperimentTreatment) == "{}" { // empty struct
			dst.AppStoreVersionExperimentTreatment = nil
		} else {
			if err = validator.Validate(dst.AppStoreVersionExperimentTreatment); err != nil {
				dst.AppStoreVersionExperimentTreatment = nil
			} else {
				match++
			}
		}
	} else {
		dst.AppStoreVersionExperimentTreatment = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.App = nil
		dst.AppStoreVersion = nil
		dst.AppStoreVersionExperimentTreatment = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppStoreVersionExperimentsV2ResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppStoreVersionExperimentsV2ResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppStoreVersionExperimentsV2ResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.AppStoreVersion != nil {
		return json.Marshal(&src.AppStoreVersion)
	}

	if src.AppStoreVersionExperimentTreatment != nil {
		return json.Marshal(&src.AppStoreVersionExperimentTreatment)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppStoreVersionExperimentsV2ResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.App != nil {
		return obj.App
	}

	if obj.AppStoreVersion != nil {
		return obj.AppStoreVersion
	}

	if obj.AppStoreVersionExperimentTreatment != nil {
		return obj.AppStoreVersionExperimentTreatment
	}

	// all schemas are nil
	return nil
}

type NullableAppStoreVersionExperimentsV2ResponseIncludedInner struct {
	value *AppStoreVersionExperimentsV2ResponseIncludedInner
	isSet bool
}

func (v NullableAppStoreVersionExperimentsV2ResponseIncludedInner) Get() *AppStoreVersionExperimentsV2ResponseIncludedInner {
	return v.value
}

func (v *NullableAppStoreVersionExperimentsV2ResponseIncludedInner) Set(val *AppStoreVersionExperimentsV2ResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentsV2ResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentsV2ResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentsV2ResponseIncludedInner(val *AppStoreVersionExperimentsV2ResponseIncludedInner) *NullableAppStoreVersionExperimentsV2ResponseIncludedInner {
	return &NullableAppStoreVersionExperimentsV2ResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentsV2ResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentsV2ResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


