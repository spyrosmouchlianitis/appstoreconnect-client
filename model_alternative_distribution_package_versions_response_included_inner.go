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

// AlternativeDistributionPackageVersionsResponseIncludedInner - struct for AlternativeDistributionPackageVersionsResponseIncludedInner
type AlternativeDistributionPackageVersionsResponseIncludedInner struct {
	AlternativeDistributionPackage *AlternativeDistributionPackage
	AlternativeDistributionPackageDelta *AlternativeDistributionPackageDelta
	AlternativeDistributionPackageVariant *AlternativeDistributionPackageVariant
}

// AlternativeDistributionPackageAsAlternativeDistributionPackageVersionsResponseIncludedInner is a convenience function that returns AlternativeDistributionPackage wrapped in AlternativeDistributionPackageVersionsResponseIncludedInner
func AlternativeDistributionPackageAsAlternativeDistributionPackageVersionsResponseIncludedInner(v *AlternativeDistributionPackage) AlternativeDistributionPackageVersionsResponseIncludedInner {
	return AlternativeDistributionPackageVersionsResponseIncludedInner{
		AlternativeDistributionPackage: v,
	}
}

// AlternativeDistributionPackageDeltaAsAlternativeDistributionPackageVersionsResponseIncludedInner is a convenience function that returns AlternativeDistributionPackageDelta wrapped in AlternativeDistributionPackageVersionsResponseIncludedInner
func AlternativeDistributionPackageDeltaAsAlternativeDistributionPackageVersionsResponseIncludedInner(v *AlternativeDistributionPackageDelta) AlternativeDistributionPackageVersionsResponseIncludedInner {
	return AlternativeDistributionPackageVersionsResponseIncludedInner{
		AlternativeDistributionPackageDelta: v,
	}
}

// AlternativeDistributionPackageVariantAsAlternativeDistributionPackageVersionsResponseIncludedInner is a convenience function that returns AlternativeDistributionPackageVariant wrapped in AlternativeDistributionPackageVersionsResponseIncludedInner
func AlternativeDistributionPackageVariantAsAlternativeDistributionPackageVersionsResponseIncludedInner(v *AlternativeDistributionPackageVariant) AlternativeDistributionPackageVersionsResponseIncludedInner {
	return AlternativeDistributionPackageVersionsResponseIncludedInner{
		AlternativeDistributionPackageVariant: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AlternativeDistributionPackageVersionsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AlternativeDistributionPackage
	err = newStrictDecoder(data).Decode(&dst.AlternativeDistributionPackage)
	if err == nil {
		jsonAlternativeDistributionPackage, _ := json.Marshal(dst.AlternativeDistributionPackage)
		if string(jsonAlternativeDistributionPackage) == "{}" { // empty struct
			dst.AlternativeDistributionPackage = nil
		} else {
			if err = validator.Validate(dst.AlternativeDistributionPackage); err != nil {
				dst.AlternativeDistributionPackage = nil
			} else {
				match++
			}
		}
	} else {
		dst.AlternativeDistributionPackage = nil
	}

	// try to unmarshal data into AlternativeDistributionPackageDelta
	err = newStrictDecoder(data).Decode(&dst.AlternativeDistributionPackageDelta)
	if err == nil {
		jsonAlternativeDistributionPackageDelta, _ := json.Marshal(dst.AlternativeDistributionPackageDelta)
		if string(jsonAlternativeDistributionPackageDelta) == "{}" { // empty struct
			dst.AlternativeDistributionPackageDelta = nil
		} else {
			if err = validator.Validate(dst.AlternativeDistributionPackageDelta); err != nil {
				dst.AlternativeDistributionPackageDelta = nil
			} else {
				match++
			}
		}
	} else {
		dst.AlternativeDistributionPackageDelta = nil
	}

	// try to unmarshal data into AlternativeDistributionPackageVariant
	err = newStrictDecoder(data).Decode(&dst.AlternativeDistributionPackageVariant)
	if err == nil {
		jsonAlternativeDistributionPackageVariant, _ := json.Marshal(dst.AlternativeDistributionPackageVariant)
		if string(jsonAlternativeDistributionPackageVariant) == "{}" { // empty struct
			dst.AlternativeDistributionPackageVariant = nil
		} else {
			if err = validator.Validate(dst.AlternativeDistributionPackageVariant); err != nil {
				dst.AlternativeDistributionPackageVariant = nil
			} else {
				match++
			}
		}
	} else {
		dst.AlternativeDistributionPackageVariant = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AlternativeDistributionPackage = nil
		dst.AlternativeDistributionPackageDelta = nil
		dst.AlternativeDistributionPackageVariant = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AlternativeDistributionPackageVersionsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AlternativeDistributionPackageVersionsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AlternativeDistributionPackageVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AlternativeDistributionPackage != nil {
		return json.Marshal(&src.AlternativeDistributionPackage)
	}

	if src.AlternativeDistributionPackageDelta != nil {
		return json.Marshal(&src.AlternativeDistributionPackageDelta)
	}

	if src.AlternativeDistributionPackageVariant != nil {
		return json.Marshal(&src.AlternativeDistributionPackageVariant)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AlternativeDistributionPackageVersionsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AlternativeDistributionPackage != nil {
		return obj.AlternativeDistributionPackage
	}

	if obj.AlternativeDistributionPackageDelta != nil {
		return obj.AlternativeDistributionPackageDelta
	}

	if obj.AlternativeDistributionPackageVariant != nil {
		return obj.AlternativeDistributionPackageVariant
	}

	// all schemas are nil
	return nil
}

type NullableAlternativeDistributionPackageVersionsResponseIncludedInner struct {
	value *AlternativeDistributionPackageVersionsResponseIncludedInner
	isSet bool
}

func (v NullableAlternativeDistributionPackageVersionsResponseIncludedInner) Get() *AlternativeDistributionPackageVersionsResponseIncludedInner {
	return v.value
}

func (v *NullableAlternativeDistributionPackageVersionsResponseIncludedInner) Set(val *AlternativeDistributionPackageVersionsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeDistributionPackageVersionsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeDistributionPackageVersionsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeDistributionPackageVersionsResponseIncludedInner(val *AlternativeDistributionPackageVersionsResponseIncludedInner) *NullableAlternativeDistributionPackageVersionsResponseIncludedInner {
	return &NullableAlternativeDistributionPackageVersionsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAlternativeDistributionPackageVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeDistributionPackageVersionsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


