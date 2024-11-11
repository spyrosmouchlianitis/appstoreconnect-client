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

// checks if the BetaTesterInvitationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterInvitationCreateRequestDataRelationships{}

// BetaTesterInvitationCreateRequestDataRelationships struct for BetaTesterInvitationCreateRequestDataRelationships
type BetaTesterInvitationCreateRequestDataRelationships struct {
	BetaTester BetaTesterInvitationCreateRequestDataRelationshipsBetaTester `json:"betaTester"`
	App AnalyticsReportRequestCreateRequestDataRelationshipsApp `json:"app"`
}

type _BetaTesterInvitationCreateRequestDataRelationships BetaTesterInvitationCreateRequestDataRelationships

// NewBetaTesterInvitationCreateRequestDataRelationships instantiates a new BetaTesterInvitationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterInvitationCreateRequestDataRelationships(betaTester BetaTesterInvitationCreateRequestDataRelationshipsBetaTester, app AnalyticsReportRequestCreateRequestDataRelationshipsApp) *BetaTesterInvitationCreateRequestDataRelationships {
	this := BetaTesterInvitationCreateRequestDataRelationships{}
	this.BetaTester = betaTester
	this.App = app
	return &this
}

// NewBetaTesterInvitationCreateRequestDataRelationshipsWithDefaults instantiates a new BetaTesterInvitationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterInvitationCreateRequestDataRelationshipsWithDefaults() *BetaTesterInvitationCreateRequestDataRelationships {
	this := BetaTesterInvitationCreateRequestDataRelationships{}
	return &this
}

// GetBetaTester returns the BetaTester field value
func (o *BetaTesterInvitationCreateRequestDataRelationships) GetBetaTester() BetaTesterInvitationCreateRequestDataRelationshipsBetaTester {
	if o == nil {
		var ret BetaTesterInvitationCreateRequestDataRelationshipsBetaTester
		return ret
	}

	return o.BetaTester
}

// GetBetaTesterOk returns a tuple with the BetaTester field value
// and a boolean to check if the value has been set.
func (o *BetaTesterInvitationCreateRequestDataRelationships) GetBetaTesterOk() (*BetaTesterInvitationCreateRequestDataRelationshipsBetaTester, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BetaTester, true
}

// SetBetaTester sets field value
func (o *BetaTesterInvitationCreateRequestDataRelationships) SetBetaTester(v BetaTesterInvitationCreateRequestDataRelationshipsBetaTester) {
	o.BetaTester = v
}

// GetApp returns the App field value
func (o *BetaTesterInvitationCreateRequestDataRelationships) GetApp() AnalyticsReportRequestCreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AnalyticsReportRequestCreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *BetaTesterInvitationCreateRequestDataRelationships) GetAppOk() (*AnalyticsReportRequestCreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *BetaTesterInvitationCreateRequestDataRelationships) SetApp(v AnalyticsReportRequestCreateRequestDataRelationshipsApp) {
	o.App = v
}

func (o BetaTesterInvitationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterInvitationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["betaTester"] = o.BetaTester
	toSerialize["app"] = o.App
	return toSerialize, nil
}

func (o *BetaTesterInvitationCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"betaTester",
		"app",
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

	varBetaTesterInvitationCreateRequestDataRelationships := _BetaTesterInvitationCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaTesterInvitationCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = BetaTesterInvitationCreateRequestDataRelationships(varBetaTesterInvitationCreateRequestDataRelationships)

	return err
}

type NullableBetaTesterInvitationCreateRequestDataRelationships struct {
	value *BetaTesterInvitationCreateRequestDataRelationships
	isSet bool
}

func (v NullableBetaTesterInvitationCreateRequestDataRelationships) Get() *BetaTesterInvitationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBetaTesterInvitationCreateRequestDataRelationships) Set(val *BetaTesterInvitationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterInvitationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterInvitationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterInvitationCreateRequestDataRelationships(val *BetaTesterInvitationCreateRequestDataRelationships) *NullableBetaTesterInvitationCreateRequestDataRelationships {
	return &NullableBetaTesterInvitationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBetaTesterInvitationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterInvitationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


