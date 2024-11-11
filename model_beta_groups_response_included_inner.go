/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// BetaGroupsResponseIncludedInner - struct for BetaGroupsResponseIncludedInner
type BetaGroupsResponseIncludedInner struct {
	App *App
	BetaTester *BetaTester
	Build *Build
}

// AppAsBetaGroupsResponseIncludedInner is a convenience function that returns App wrapped in BetaGroupsResponseIncludedInner
func AppAsBetaGroupsResponseIncludedInner(v *App) BetaGroupsResponseIncludedInner {
	return BetaGroupsResponseIncludedInner{
		App: v,
	}
}

// BetaTesterAsBetaGroupsResponseIncludedInner is a convenience function that returns BetaTester wrapped in BetaGroupsResponseIncludedInner
func BetaTesterAsBetaGroupsResponseIncludedInner(v *BetaTester) BetaGroupsResponseIncludedInner {
	return BetaGroupsResponseIncludedInner{
		BetaTester: v,
	}
}

// BuildAsBetaGroupsResponseIncludedInner is a convenience function that returns Build wrapped in BetaGroupsResponseIncludedInner
func BuildAsBetaGroupsResponseIncludedInner(v *Build) BetaGroupsResponseIncludedInner {
	return BetaGroupsResponseIncludedInner{
		Build: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BetaGroupsResponseIncludedInner) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into Build
	err = newStrictDecoder(data).Decode(&dst.Build)
	if err == nil {
		jsonBuild, _ := json.Marshal(dst.Build)
		if string(jsonBuild) == "{}" { // empty struct
			dst.Build = nil
		} else {
			if err = validator.Validate(dst.Build); err != nil {
				dst.Build = nil
			} else {
				match++
			}
		}
	} else {
		dst.Build = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.App = nil
		dst.BetaTester = nil
		dst.Build = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BetaGroupsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BetaGroupsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BetaGroupsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.BetaTester != nil {
		return json.Marshal(&src.BetaTester)
	}

	if src.Build != nil {
		return json.Marshal(&src.Build)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BetaGroupsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.App != nil {
		return obj.App
	}

	if obj.BetaTester != nil {
		return obj.BetaTester
	}

	if obj.Build != nil {
		return obj.Build
	}

	// all schemas are nil
	return nil
}

type NullableBetaGroupsResponseIncludedInner struct {
	value *BetaGroupsResponseIncludedInner
	isSet bool
}

func (v NullableBetaGroupsResponseIncludedInner) Get() *BetaGroupsResponseIncludedInner {
	return v.value
}

func (v *NullableBetaGroupsResponseIncludedInner) Set(val *BetaGroupsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupsResponseIncludedInner(val *BetaGroupsResponseIncludedInner) *NullableBetaGroupsResponseIncludedInner {
	return &NullableBetaGroupsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableBetaGroupsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


