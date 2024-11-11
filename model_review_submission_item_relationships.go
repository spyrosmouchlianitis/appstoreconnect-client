/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the ReviewSubmissionItemRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionItemRelationships{}

// ReviewSubmissionItemRelationships struct for ReviewSubmissionItemRelationships
type ReviewSubmissionItemRelationships struct {
	AppStoreVersion *AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion `json:"appStoreVersion,omitempty"`
	AppCustomProductPageVersion *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion `json:"appCustomProductPageVersion,omitempty"`
	AppStoreVersionExperiment *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment `json:"appStoreVersionExperiment,omitempty"`
	AppStoreVersionExperimentV2 *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment `json:"appStoreVersionExperimentV2,omitempty"`
	AppEvent *AppEventLocalizationRelationshipsAppEvent `json:"appEvent,omitempty"`
}

// NewReviewSubmissionItemRelationships instantiates a new ReviewSubmissionItemRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionItemRelationships() *ReviewSubmissionItemRelationships {
	this := ReviewSubmissionItemRelationships{}
	return &this
}

// NewReviewSubmissionItemRelationshipsWithDefaults instantiates a new ReviewSubmissionItemRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionItemRelationshipsWithDefaults() *ReviewSubmissionItemRelationships {
	this := ReviewSubmissionItemRelationships{}
	return &this
}

// GetAppStoreVersion returns the AppStoreVersion field value if set, zero value otherwise.
func (o *ReviewSubmissionItemRelationships) GetAppStoreVersion() AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.AppStoreVersion) {
		var ret AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemRelationships) GetAppStoreVersionOk() (*AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.AppStoreVersion) {
		return nil, false
	}
	return o.AppStoreVersion, true
}

// HasAppStoreVersion returns a boolean if a field has been set.
func (o *ReviewSubmissionItemRelationships) HasAppStoreVersion() bool {
	if o != nil && !IsNil(o.AppStoreVersion) {
		return true
	}

	return false
}

// SetAppStoreVersion gets a reference to the given AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion and assigns it to the AppStoreVersion field.
func (o *ReviewSubmissionItemRelationships) SetAppStoreVersion(v AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion) {
	o.AppStoreVersion = &v
}

// GetAppCustomProductPageVersion returns the AppCustomProductPageVersion field value if set, zero value otherwise.
func (o *ReviewSubmissionItemRelationships) GetAppCustomProductPageVersion() AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion {
	if o == nil || IsNil(o.AppCustomProductPageVersion) {
		var ret AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion
		return ret
	}
	return *o.AppCustomProductPageVersion
}

// GetAppCustomProductPageVersionOk returns a tuple with the AppCustomProductPageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemRelationships) GetAppCustomProductPageVersionOk() (*AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion, bool) {
	if o == nil || IsNil(o.AppCustomProductPageVersion) {
		return nil, false
	}
	return o.AppCustomProductPageVersion, true
}

// HasAppCustomProductPageVersion returns a boolean if a field has been set.
func (o *ReviewSubmissionItemRelationships) HasAppCustomProductPageVersion() bool {
	if o != nil && !IsNil(o.AppCustomProductPageVersion) {
		return true
	}

	return false
}

// SetAppCustomProductPageVersion gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion and assigns it to the AppCustomProductPageVersion field.
func (o *ReviewSubmissionItemRelationships) SetAppCustomProductPageVersion(v AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) {
	o.AppCustomProductPageVersion = &v
}

// GetAppStoreVersionExperiment returns the AppStoreVersionExperiment field value if set, zero value otherwise.
func (o *ReviewSubmissionItemRelationships) GetAppStoreVersionExperiment() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	if o == nil || IsNil(o.AppStoreVersionExperiment) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
		return ret
	}
	return *o.AppStoreVersionExperiment
}

// GetAppStoreVersionExperimentOk returns a tuple with the AppStoreVersionExperiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemRelationships) GetAppStoreVersionExperimentOk() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperiment) {
		return nil, false
	}
	return o.AppStoreVersionExperiment, true
}

// HasAppStoreVersionExperiment returns a boolean if a field has been set.
func (o *ReviewSubmissionItemRelationships) HasAppStoreVersionExperiment() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperiment) {
		return true
	}

	return false
}

// SetAppStoreVersionExperiment gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment and assigns it to the AppStoreVersionExperiment field.
func (o *ReviewSubmissionItemRelationships) SetAppStoreVersionExperiment(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) {
	o.AppStoreVersionExperiment = &v
}

// GetAppStoreVersionExperimentV2 returns the AppStoreVersionExperimentV2 field value if set, zero value otherwise.
func (o *ReviewSubmissionItemRelationships) GetAppStoreVersionExperimentV2() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	if o == nil || IsNil(o.AppStoreVersionExperimentV2) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
		return ret
	}
	return *o.AppStoreVersionExperimentV2
}

// GetAppStoreVersionExperimentV2Ok returns a tuple with the AppStoreVersionExperimentV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemRelationships) GetAppStoreVersionExperimentV2Ok() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperimentV2) {
		return nil, false
	}
	return o.AppStoreVersionExperimentV2, true
}

// HasAppStoreVersionExperimentV2 returns a boolean if a field has been set.
func (o *ReviewSubmissionItemRelationships) HasAppStoreVersionExperimentV2() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperimentV2) {
		return true
	}

	return false
}

// SetAppStoreVersionExperimentV2 gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment and assigns it to the AppStoreVersionExperimentV2 field.
func (o *ReviewSubmissionItemRelationships) SetAppStoreVersionExperimentV2(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) {
	o.AppStoreVersionExperimentV2 = &v
}

// GetAppEvent returns the AppEvent field value if set, zero value otherwise.
func (o *ReviewSubmissionItemRelationships) GetAppEvent() AppEventLocalizationRelationshipsAppEvent {
	if o == nil || IsNil(o.AppEvent) {
		var ret AppEventLocalizationRelationshipsAppEvent
		return ret
	}
	return *o.AppEvent
}

// GetAppEventOk returns a tuple with the AppEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemRelationships) GetAppEventOk() (*AppEventLocalizationRelationshipsAppEvent, bool) {
	if o == nil || IsNil(o.AppEvent) {
		return nil, false
	}
	return o.AppEvent, true
}

// HasAppEvent returns a boolean if a field has been set.
func (o *ReviewSubmissionItemRelationships) HasAppEvent() bool {
	if o != nil && !IsNil(o.AppEvent) {
		return true
	}

	return false
}

// SetAppEvent gets a reference to the given AppEventLocalizationRelationshipsAppEvent and assigns it to the AppEvent field.
func (o *ReviewSubmissionItemRelationships) SetAppEvent(v AppEventLocalizationRelationshipsAppEvent) {
	o.AppEvent = &v
}

func (o ReviewSubmissionItemRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionItemRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppStoreVersion) {
		toSerialize["appStoreVersion"] = o.AppStoreVersion
	}
	if !IsNil(o.AppCustomProductPageVersion) {
		toSerialize["appCustomProductPageVersion"] = o.AppCustomProductPageVersion
	}
	if !IsNil(o.AppStoreVersionExperiment) {
		toSerialize["appStoreVersionExperiment"] = o.AppStoreVersionExperiment
	}
	if !IsNil(o.AppStoreVersionExperimentV2) {
		toSerialize["appStoreVersionExperimentV2"] = o.AppStoreVersionExperimentV2
	}
	if !IsNil(o.AppEvent) {
		toSerialize["appEvent"] = o.AppEvent
	}
	return toSerialize, nil
}

type NullableReviewSubmissionItemRelationships struct {
	value *ReviewSubmissionItemRelationships
	isSet bool
}

func (v NullableReviewSubmissionItemRelationships) Get() *ReviewSubmissionItemRelationships {
	return v.value
}

func (v *NullableReviewSubmissionItemRelationships) Set(val *ReviewSubmissionItemRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionItemRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionItemRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionItemRelationships(val *ReviewSubmissionItemRelationships) *NullableReviewSubmissionItemRelationships {
	return &NullableReviewSubmissionItemRelationships{value: val, isSet: true}
}

func (v NullableReviewSubmissionItemRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionItemRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


