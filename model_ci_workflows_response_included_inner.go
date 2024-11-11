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

// CiWorkflowsResponseIncludedInner - struct for CiWorkflowsResponseIncludedInner
type CiWorkflowsResponseIncludedInner struct {
	CiMacOsVersion *CiMacOsVersion
	CiProduct *CiProduct
	CiXcodeVersion *CiXcodeVersion
	ScmRepository *ScmRepository
}

// CiMacOsVersionAsCiWorkflowsResponseIncludedInner is a convenience function that returns CiMacOsVersion wrapped in CiWorkflowsResponseIncludedInner
func CiMacOsVersionAsCiWorkflowsResponseIncludedInner(v *CiMacOsVersion) CiWorkflowsResponseIncludedInner {
	return CiWorkflowsResponseIncludedInner{
		CiMacOsVersion: v,
	}
}

// CiProductAsCiWorkflowsResponseIncludedInner is a convenience function that returns CiProduct wrapped in CiWorkflowsResponseIncludedInner
func CiProductAsCiWorkflowsResponseIncludedInner(v *CiProduct) CiWorkflowsResponseIncludedInner {
	return CiWorkflowsResponseIncludedInner{
		CiProduct: v,
	}
}

// CiXcodeVersionAsCiWorkflowsResponseIncludedInner is a convenience function that returns CiXcodeVersion wrapped in CiWorkflowsResponseIncludedInner
func CiXcodeVersionAsCiWorkflowsResponseIncludedInner(v *CiXcodeVersion) CiWorkflowsResponseIncludedInner {
	return CiWorkflowsResponseIncludedInner{
		CiXcodeVersion: v,
	}
}

// ScmRepositoryAsCiWorkflowsResponseIncludedInner is a convenience function that returns ScmRepository wrapped in CiWorkflowsResponseIncludedInner
func ScmRepositoryAsCiWorkflowsResponseIncludedInner(v *ScmRepository) CiWorkflowsResponseIncludedInner {
	return CiWorkflowsResponseIncludedInner{
		ScmRepository: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CiWorkflowsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CiMacOsVersion
	err = newStrictDecoder(data).Decode(&dst.CiMacOsVersion)
	if err == nil {
		jsonCiMacOsVersion, _ := json.Marshal(dst.CiMacOsVersion)
		if string(jsonCiMacOsVersion) == "{}" { // empty struct
			dst.CiMacOsVersion = nil
		} else {
			if err = validator.Validate(dst.CiMacOsVersion); err != nil {
				dst.CiMacOsVersion = nil
			} else {
				match++
			}
		}
	} else {
		dst.CiMacOsVersion = nil
	}

	// try to unmarshal data into CiProduct
	err = newStrictDecoder(data).Decode(&dst.CiProduct)
	if err == nil {
		jsonCiProduct, _ := json.Marshal(dst.CiProduct)
		if string(jsonCiProduct) == "{}" { // empty struct
			dst.CiProduct = nil
		} else {
			if err = validator.Validate(dst.CiProduct); err != nil {
				dst.CiProduct = nil
			} else {
				match++
			}
		}
	} else {
		dst.CiProduct = nil
	}

	// try to unmarshal data into CiXcodeVersion
	err = newStrictDecoder(data).Decode(&dst.CiXcodeVersion)
	if err == nil {
		jsonCiXcodeVersion, _ := json.Marshal(dst.CiXcodeVersion)
		if string(jsonCiXcodeVersion) == "{}" { // empty struct
			dst.CiXcodeVersion = nil
		} else {
			if err = validator.Validate(dst.CiXcodeVersion); err != nil {
				dst.CiXcodeVersion = nil
			} else {
				match++
			}
		}
	} else {
		dst.CiXcodeVersion = nil
	}

	// try to unmarshal data into ScmRepository
	err = newStrictDecoder(data).Decode(&dst.ScmRepository)
	if err == nil {
		jsonScmRepository, _ := json.Marshal(dst.ScmRepository)
		if string(jsonScmRepository) == "{}" { // empty struct
			dst.ScmRepository = nil
		} else {
			if err = validator.Validate(dst.ScmRepository); err != nil {
				dst.ScmRepository = nil
			} else {
				match++
			}
		}
	} else {
		dst.ScmRepository = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CiMacOsVersion = nil
		dst.CiProduct = nil
		dst.CiXcodeVersion = nil
		dst.ScmRepository = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CiWorkflowsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CiWorkflowsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CiWorkflowsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.CiMacOsVersion != nil {
		return json.Marshal(&src.CiMacOsVersion)
	}

	if src.CiProduct != nil {
		return json.Marshal(&src.CiProduct)
	}

	if src.CiXcodeVersion != nil {
		return json.Marshal(&src.CiXcodeVersion)
	}

	if src.ScmRepository != nil {
		return json.Marshal(&src.ScmRepository)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CiWorkflowsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CiMacOsVersion != nil {
		return obj.CiMacOsVersion
	}

	if obj.CiProduct != nil {
		return obj.CiProduct
	}

	if obj.CiXcodeVersion != nil {
		return obj.CiXcodeVersion
	}

	if obj.ScmRepository != nil {
		return obj.ScmRepository
	}

	// all schemas are nil
	return nil
}

type NullableCiWorkflowsResponseIncludedInner struct {
	value *CiWorkflowsResponseIncludedInner
	isSet bool
}

func (v NullableCiWorkflowsResponseIncludedInner) Get() *CiWorkflowsResponseIncludedInner {
	return v.value
}

func (v *NullableCiWorkflowsResponseIncludedInner) Set(val *CiWorkflowsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowsResponseIncludedInner(val *CiWorkflowsResponseIncludedInner) *NullableCiWorkflowsResponseIncludedInner {
	return &NullableCiWorkflowsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableCiWorkflowsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


