/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ReviewSubmissionItemCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionItemCreateRequestDataRelationships{}

// ReviewSubmissionItemCreateRequestDataRelationships struct for ReviewSubmissionItemCreateRequestDataRelationships
type ReviewSubmissionItemCreateRequestDataRelationships struct {
	ReviewSubmission ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission `json:"reviewSubmission"`
	AppStoreVersion *AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion `json:"appStoreVersion,omitempty"`
	AppCustomProductPageVersion *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion `json:"appCustomProductPageVersion,omitempty"`
	AppStoreVersionExperiment *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment `json:"appStoreVersionExperiment,omitempty"`
	AppStoreVersionExperimentV2 *AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment `json:"appStoreVersionExperimentV2,omitempty"`
	AppEvent *AppEventLocalizationRelationshipsAppEvent `json:"appEvent,omitempty"`
}

type _ReviewSubmissionItemCreateRequestDataRelationships ReviewSubmissionItemCreateRequestDataRelationships

// NewReviewSubmissionItemCreateRequestDataRelationships instantiates a new ReviewSubmissionItemCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionItemCreateRequestDataRelationships(reviewSubmission ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission) *ReviewSubmissionItemCreateRequestDataRelationships {
	this := ReviewSubmissionItemCreateRequestDataRelationships{}
	this.ReviewSubmission = reviewSubmission
	return &this
}

// NewReviewSubmissionItemCreateRequestDataRelationshipsWithDefaults instantiates a new ReviewSubmissionItemCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionItemCreateRequestDataRelationshipsWithDefaults() *ReviewSubmissionItemCreateRequestDataRelationships {
	this := ReviewSubmissionItemCreateRequestDataRelationships{}
	return &this
}

// GetReviewSubmission returns the ReviewSubmission field value
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetReviewSubmission() ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission {
	if o == nil {
		var ret ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission
		return ret
	}

	return o.ReviewSubmission
}

// GetReviewSubmissionOk returns a tuple with the ReviewSubmission field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetReviewSubmissionOk() (*ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewSubmission, true
}

// SetReviewSubmission sets field value
func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetReviewSubmission(v ReviewSubmissionItemCreateRequestDataRelationshipsReviewSubmission) {
	o.ReviewSubmission = v
}

// GetAppStoreVersion returns the AppStoreVersion field value if set, zero value otherwise.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersion() AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.AppStoreVersion) {
		var ret AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionOk() (*AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.AppStoreVersion) {
		return nil, false
	}
	return o.AppStoreVersion, true
}

// HasAppStoreVersion returns a boolean if a field has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppStoreVersion() bool {
	if o != nil && !IsNil(o.AppStoreVersion) {
		return true
	}

	return false
}

// SetAppStoreVersion gets a reference to the given AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion and assigns it to the AppStoreVersion field.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppStoreVersion(v AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion) {
	o.AppStoreVersion = &v
}

// GetAppCustomProductPageVersion returns the AppCustomProductPageVersion field value if set, zero value otherwise.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppCustomProductPageVersion() AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion {
	if o == nil || IsNil(o.AppCustomProductPageVersion) {
		var ret AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion
		return ret
	}
	return *o.AppCustomProductPageVersion
}

// GetAppCustomProductPageVersionOk returns a tuple with the AppCustomProductPageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppCustomProductPageVersionOk() (*AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion, bool) {
	if o == nil || IsNil(o.AppCustomProductPageVersion) {
		return nil, false
	}
	return o.AppCustomProductPageVersion, true
}

// HasAppCustomProductPageVersion returns a boolean if a field has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppCustomProductPageVersion() bool {
	if o != nil && !IsNil(o.AppCustomProductPageVersion) {
		return true
	}

	return false
}

// SetAppCustomProductPageVersion gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion and assigns it to the AppCustomProductPageVersion field.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppCustomProductPageVersion(v AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) {
	o.AppCustomProductPageVersion = &v
}

// GetAppStoreVersionExperiment returns the AppStoreVersionExperiment field value if set, zero value otherwise.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionExperiment() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	if o == nil || IsNil(o.AppStoreVersionExperiment) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
		return ret
	}
	return *o.AppStoreVersionExperiment
}

// GetAppStoreVersionExperimentOk returns a tuple with the AppStoreVersionExperiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionExperimentOk() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperiment) {
		return nil, false
	}
	return o.AppStoreVersionExperiment, true
}

// HasAppStoreVersionExperiment returns a boolean if a field has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppStoreVersionExperiment() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperiment) {
		return true
	}

	return false
}

// SetAppStoreVersionExperiment gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment and assigns it to the AppStoreVersionExperiment field.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppStoreVersionExperiment(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) {
	o.AppStoreVersionExperiment = &v
}

// GetAppStoreVersionExperimentV2 returns the AppStoreVersionExperimentV2 field value if set, zero value otherwise.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionExperimentV2() AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment {
	if o == nil || IsNil(o.AppStoreVersionExperimentV2) {
		var ret AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment
		return ret
	}
	return *o.AppStoreVersionExperimentV2
}

// GetAppStoreVersionExperimentV2Ok returns a tuple with the AppStoreVersionExperimentV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppStoreVersionExperimentV2Ok() (*AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperimentV2) {
		return nil, false
	}
	return o.AppStoreVersionExperimentV2, true
}

// HasAppStoreVersionExperimentV2 returns a boolean if a field has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppStoreVersionExperimentV2() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperimentV2) {
		return true
	}

	return false
}

// SetAppStoreVersionExperimentV2 gets a reference to the given AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment and assigns it to the AppStoreVersionExperimentV2 field.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppStoreVersionExperimentV2(v AppStoreVersionExperimentTreatmentRelationshipsAppStoreVersionExperiment) {
	o.AppStoreVersionExperimentV2 = &v
}

// GetAppEvent returns the AppEvent field value if set, zero value otherwise.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppEvent() AppEventLocalizationRelationshipsAppEvent {
	if o == nil || IsNil(o.AppEvent) {
		var ret AppEventLocalizationRelationshipsAppEvent
		return ret
	}
	return *o.AppEvent
}

// GetAppEventOk returns a tuple with the AppEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) GetAppEventOk() (*AppEventLocalizationRelationshipsAppEvent, bool) {
	if o == nil || IsNil(o.AppEvent) {
		return nil, false
	}
	return o.AppEvent, true
}

// HasAppEvent returns a boolean if a field has been set.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) HasAppEvent() bool {
	if o != nil && !IsNil(o.AppEvent) {
		return true
	}

	return false
}

// SetAppEvent gets a reference to the given AppEventLocalizationRelationshipsAppEvent and assigns it to the AppEvent field.
func (o *ReviewSubmissionItemCreateRequestDataRelationships) SetAppEvent(v AppEventLocalizationRelationshipsAppEvent) {
	o.AppEvent = &v
}

func (o ReviewSubmissionItemCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionItemCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reviewSubmission"] = o.ReviewSubmission
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

func (o *ReviewSubmissionItemCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reviewSubmission",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReviewSubmissionItemCreateRequestDataRelationships := _ReviewSubmissionItemCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReviewSubmissionItemCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = ReviewSubmissionItemCreateRequestDataRelationships(varReviewSubmissionItemCreateRequestDataRelationships)

	return err
}

type NullableReviewSubmissionItemCreateRequestDataRelationships struct {
	value *ReviewSubmissionItemCreateRequestDataRelationships
	isSet bool
}

func (v NullableReviewSubmissionItemCreateRequestDataRelationships) Get() *ReviewSubmissionItemCreateRequestDataRelationships {
	return v.value
}

func (v *NullableReviewSubmissionItemCreateRequestDataRelationships) Set(val *ReviewSubmissionItemCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionItemCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionItemCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionItemCreateRequestDataRelationships(val *ReviewSubmissionItemCreateRequestDataRelationships) *NullableReviewSubmissionItemCreateRequestDataRelationships {
	return &NullableReviewSubmissionItemCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableReviewSubmissionItemCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionItemCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

