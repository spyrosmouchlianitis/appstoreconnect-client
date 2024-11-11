/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the AgeRatingDeclarationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgeRatingDeclarationAttributes{}

// AgeRatingDeclarationAttributes struct for AgeRatingDeclarationAttributes
type AgeRatingDeclarationAttributes struct {
	AlcoholTobaccoOrDrugUseOrReferences *string `json:"alcoholTobaccoOrDrugUseOrReferences,omitempty"`
	Contests *string `json:"contests,omitempty"`
	// Deprecated
	GamblingAndContests *bool `json:"gamblingAndContests,omitempty"`
	Gambling *bool `json:"gambling,omitempty"`
	GamblingSimulated *string `json:"gamblingSimulated,omitempty"`
	KidsAgeBand *KidsAgeBand `json:"kidsAgeBand,omitempty"`
	LootBox *bool `json:"lootBox,omitempty"`
	MedicalOrTreatmentInformation *string `json:"medicalOrTreatmentInformation,omitempty"`
	ProfanityOrCrudeHumor *string `json:"profanityOrCrudeHumor,omitempty"`
	SexualContentGraphicAndNudity *string `json:"sexualContentGraphicAndNudity,omitempty"`
	SexualContentOrNudity *string `json:"sexualContentOrNudity,omitempty"`
	HorrorOrFearThemes *string `json:"horrorOrFearThemes,omitempty"`
	MatureOrSuggestiveThemes *string `json:"matureOrSuggestiveThemes,omitempty"`
	UnrestrictedWebAccess *bool `json:"unrestrictedWebAccess,omitempty"`
	ViolenceCartoonOrFantasy *string `json:"violenceCartoonOrFantasy,omitempty"`
	ViolenceRealisticProlongedGraphicOrSadistic *string `json:"violenceRealisticProlongedGraphicOrSadistic,omitempty"`
	ViolenceRealistic *string `json:"violenceRealistic,omitempty"`
	AgeRatingOverride *string `json:"ageRatingOverride,omitempty"`
	KoreaAgeRatingOverride *string `json:"koreaAgeRatingOverride,omitempty"`
	// Deprecated
	SeventeenPlus *bool `json:"seventeenPlus,omitempty"`
}

// NewAgeRatingDeclarationAttributes instantiates a new AgeRatingDeclarationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgeRatingDeclarationAttributes() *AgeRatingDeclarationAttributes {
	this := AgeRatingDeclarationAttributes{}
	return &this
}

// NewAgeRatingDeclarationAttributesWithDefaults instantiates a new AgeRatingDeclarationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeRatingDeclarationAttributesWithDefaults() *AgeRatingDeclarationAttributes {
	this := AgeRatingDeclarationAttributes{}
	return &this
}

// GetAlcoholTobaccoOrDrugUseOrReferences returns the AlcoholTobaccoOrDrugUseOrReferences field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetAlcoholTobaccoOrDrugUseOrReferences() string {
	if o == nil || IsNil(o.AlcoholTobaccoOrDrugUseOrReferences) {
		var ret string
		return ret
	}
	return *o.AlcoholTobaccoOrDrugUseOrReferences
}

// GetAlcoholTobaccoOrDrugUseOrReferencesOk returns a tuple with the AlcoholTobaccoOrDrugUseOrReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetAlcoholTobaccoOrDrugUseOrReferencesOk() (*string, bool) {
	if o == nil || IsNil(o.AlcoholTobaccoOrDrugUseOrReferences) {
		return nil, false
	}
	return o.AlcoholTobaccoOrDrugUseOrReferences, true
}

// HasAlcoholTobaccoOrDrugUseOrReferences returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasAlcoholTobaccoOrDrugUseOrReferences() bool {
	if o != nil && !IsNil(o.AlcoholTobaccoOrDrugUseOrReferences) {
		return true
	}

	return false
}

// SetAlcoholTobaccoOrDrugUseOrReferences gets a reference to the given string and assigns it to the AlcoholTobaccoOrDrugUseOrReferences field.
func (o *AgeRatingDeclarationAttributes) SetAlcoholTobaccoOrDrugUseOrReferences(v string) {
	o.AlcoholTobaccoOrDrugUseOrReferences = &v
}

// GetContests returns the Contests field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetContests() string {
	if o == nil || IsNil(o.Contests) {
		var ret string
		return ret
	}
	return *o.Contests
}

// GetContestsOk returns a tuple with the Contests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetContestsOk() (*string, bool) {
	if o == nil || IsNil(o.Contests) {
		return nil, false
	}
	return o.Contests, true
}

// HasContests returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasContests() bool {
	if o != nil && !IsNil(o.Contests) {
		return true
	}

	return false
}

// SetContests gets a reference to the given string and assigns it to the Contests field.
func (o *AgeRatingDeclarationAttributes) SetContests(v string) {
	o.Contests = &v
}

// GetGamblingAndContests returns the GamblingAndContests field value if set, zero value otherwise.
// Deprecated
func (o *AgeRatingDeclarationAttributes) GetGamblingAndContests() bool {
	if o == nil || IsNil(o.GamblingAndContests) {
		var ret bool
		return ret
	}
	return *o.GamblingAndContests
}

// GetGamblingAndContestsOk returns a tuple with the GamblingAndContests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AgeRatingDeclarationAttributes) GetGamblingAndContestsOk() (*bool, bool) {
	if o == nil || IsNil(o.GamblingAndContests) {
		return nil, false
	}
	return o.GamblingAndContests, true
}

// HasGamblingAndContests returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasGamblingAndContests() bool {
	if o != nil && !IsNil(o.GamblingAndContests) {
		return true
	}

	return false
}

// SetGamblingAndContests gets a reference to the given bool and assigns it to the GamblingAndContests field.
// Deprecated
func (o *AgeRatingDeclarationAttributes) SetGamblingAndContests(v bool) {
	o.GamblingAndContests = &v
}

// GetGambling returns the Gambling field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetGambling() bool {
	if o == nil || IsNil(o.Gambling) {
		var ret bool
		return ret
	}
	return *o.Gambling
}

// GetGamblingOk returns a tuple with the Gambling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetGamblingOk() (*bool, bool) {
	if o == nil || IsNil(o.Gambling) {
		return nil, false
	}
	return o.Gambling, true
}

// HasGambling returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasGambling() bool {
	if o != nil && !IsNil(o.Gambling) {
		return true
	}

	return false
}

// SetGambling gets a reference to the given bool and assigns it to the Gambling field.
func (o *AgeRatingDeclarationAttributes) SetGambling(v bool) {
	o.Gambling = &v
}

// GetGamblingSimulated returns the GamblingSimulated field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetGamblingSimulated() string {
	if o == nil || IsNil(o.GamblingSimulated) {
		var ret string
		return ret
	}
	return *o.GamblingSimulated
}

// GetGamblingSimulatedOk returns a tuple with the GamblingSimulated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetGamblingSimulatedOk() (*string, bool) {
	if o == nil || IsNil(o.GamblingSimulated) {
		return nil, false
	}
	return o.GamblingSimulated, true
}

// HasGamblingSimulated returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasGamblingSimulated() bool {
	if o != nil && !IsNil(o.GamblingSimulated) {
		return true
	}

	return false
}

// SetGamblingSimulated gets a reference to the given string and assigns it to the GamblingSimulated field.
func (o *AgeRatingDeclarationAttributes) SetGamblingSimulated(v string) {
	o.GamblingSimulated = &v
}

// GetKidsAgeBand returns the KidsAgeBand field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetKidsAgeBand() KidsAgeBand {
	if o == nil || IsNil(o.KidsAgeBand) {
		var ret KidsAgeBand
		return ret
	}
	return *o.KidsAgeBand
}

// GetKidsAgeBandOk returns a tuple with the KidsAgeBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetKidsAgeBandOk() (*KidsAgeBand, bool) {
	if o == nil || IsNil(o.KidsAgeBand) {
		return nil, false
	}
	return o.KidsAgeBand, true
}

// HasKidsAgeBand returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasKidsAgeBand() bool {
	if o != nil && !IsNil(o.KidsAgeBand) {
		return true
	}

	return false
}

// SetKidsAgeBand gets a reference to the given KidsAgeBand and assigns it to the KidsAgeBand field.
func (o *AgeRatingDeclarationAttributes) SetKidsAgeBand(v KidsAgeBand) {
	o.KidsAgeBand = &v
}

// GetLootBox returns the LootBox field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetLootBox() bool {
	if o == nil || IsNil(o.LootBox) {
		var ret bool
		return ret
	}
	return *o.LootBox
}

// GetLootBoxOk returns a tuple with the LootBox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetLootBoxOk() (*bool, bool) {
	if o == nil || IsNil(o.LootBox) {
		return nil, false
	}
	return o.LootBox, true
}

// HasLootBox returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasLootBox() bool {
	if o != nil && !IsNil(o.LootBox) {
		return true
	}

	return false
}

// SetLootBox gets a reference to the given bool and assigns it to the LootBox field.
func (o *AgeRatingDeclarationAttributes) SetLootBox(v bool) {
	o.LootBox = &v
}

// GetMedicalOrTreatmentInformation returns the MedicalOrTreatmentInformation field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetMedicalOrTreatmentInformation() string {
	if o == nil || IsNil(o.MedicalOrTreatmentInformation) {
		var ret string
		return ret
	}
	return *o.MedicalOrTreatmentInformation
}

// GetMedicalOrTreatmentInformationOk returns a tuple with the MedicalOrTreatmentInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetMedicalOrTreatmentInformationOk() (*string, bool) {
	if o == nil || IsNil(o.MedicalOrTreatmentInformation) {
		return nil, false
	}
	return o.MedicalOrTreatmentInformation, true
}

// HasMedicalOrTreatmentInformation returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasMedicalOrTreatmentInformation() bool {
	if o != nil && !IsNil(o.MedicalOrTreatmentInformation) {
		return true
	}

	return false
}

// SetMedicalOrTreatmentInformation gets a reference to the given string and assigns it to the MedicalOrTreatmentInformation field.
func (o *AgeRatingDeclarationAttributes) SetMedicalOrTreatmentInformation(v string) {
	o.MedicalOrTreatmentInformation = &v
}

// GetProfanityOrCrudeHumor returns the ProfanityOrCrudeHumor field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetProfanityOrCrudeHumor() string {
	if o == nil || IsNil(o.ProfanityOrCrudeHumor) {
		var ret string
		return ret
	}
	return *o.ProfanityOrCrudeHumor
}

// GetProfanityOrCrudeHumorOk returns a tuple with the ProfanityOrCrudeHumor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetProfanityOrCrudeHumorOk() (*string, bool) {
	if o == nil || IsNil(o.ProfanityOrCrudeHumor) {
		return nil, false
	}
	return o.ProfanityOrCrudeHumor, true
}

// HasProfanityOrCrudeHumor returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasProfanityOrCrudeHumor() bool {
	if o != nil && !IsNil(o.ProfanityOrCrudeHumor) {
		return true
	}

	return false
}

// SetProfanityOrCrudeHumor gets a reference to the given string and assigns it to the ProfanityOrCrudeHumor field.
func (o *AgeRatingDeclarationAttributes) SetProfanityOrCrudeHumor(v string) {
	o.ProfanityOrCrudeHumor = &v
}

// GetSexualContentGraphicAndNudity returns the SexualContentGraphicAndNudity field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetSexualContentGraphicAndNudity() string {
	if o == nil || IsNil(o.SexualContentGraphicAndNudity) {
		var ret string
		return ret
	}
	return *o.SexualContentGraphicAndNudity
}

// GetSexualContentGraphicAndNudityOk returns a tuple with the SexualContentGraphicAndNudity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetSexualContentGraphicAndNudityOk() (*string, bool) {
	if o == nil || IsNil(o.SexualContentGraphicAndNudity) {
		return nil, false
	}
	return o.SexualContentGraphicAndNudity, true
}

// HasSexualContentGraphicAndNudity returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasSexualContentGraphicAndNudity() bool {
	if o != nil && !IsNil(o.SexualContentGraphicAndNudity) {
		return true
	}

	return false
}

// SetSexualContentGraphicAndNudity gets a reference to the given string and assigns it to the SexualContentGraphicAndNudity field.
func (o *AgeRatingDeclarationAttributes) SetSexualContentGraphicAndNudity(v string) {
	o.SexualContentGraphicAndNudity = &v
}

// GetSexualContentOrNudity returns the SexualContentOrNudity field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetSexualContentOrNudity() string {
	if o == nil || IsNil(o.SexualContentOrNudity) {
		var ret string
		return ret
	}
	return *o.SexualContentOrNudity
}

// GetSexualContentOrNudityOk returns a tuple with the SexualContentOrNudity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetSexualContentOrNudityOk() (*string, bool) {
	if o == nil || IsNil(o.SexualContentOrNudity) {
		return nil, false
	}
	return o.SexualContentOrNudity, true
}

// HasSexualContentOrNudity returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasSexualContentOrNudity() bool {
	if o != nil && !IsNil(o.SexualContentOrNudity) {
		return true
	}

	return false
}

// SetSexualContentOrNudity gets a reference to the given string and assigns it to the SexualContentOrNudity field.
func (o *AgeRatingDeclarationAttributes) SetSexualContentOrNudity(v string) {
	o.SexualContentOrNudity = &v
}

// GetHorrorOrFearThemes returns the HorrorOrFearThemes field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetHorrorOrFearThemes() string {
	if o == nil || IsNil(o.HorrorOrFearThemes) {
		var ret string
		return ret
	}
	return *o.HorrorOrFearThemes
}

// GetHorrorOrFearThemesOk returns a tuple with the HorrorOrFearThemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetHorrorOrFearThemesOk() (*string, bool) {
	if o == nil || IsNil(o.HorrorOrFearThemes) {
		return nil, false
	}
	return o.HorrorOrFearThemes, true
}

// HasHorrorOrFearThemes returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasHorrorOrFearThemes() bool {
	if o != nil && !IsNil(o.HorrorOrFearThemes) {
		return true
	}

	return false
}

// SetHorrorOrFearThemes gets a reference to the given string and assigns it to the HorrorOrFearThemes field.
func (o *AgeRatingDeclarationAttributes) SetHorrorOrFearThemes(v string) {
	o.HorrorOrFearThemes = &v
}

// GetMatureOrSuggestiveThemes returns the MatureOrSuggestiveThemes field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetMatureOrSuggestiveThemes() string {
	if o == nil || IsNil(o.MatureOrSuggestiveThemes) {
		var ret string
		return ret
	}
	return *o.MatureOrSuggestiveThemes
}

// GetMatureOrSuggestiveThemesOk returns a tuple with the MatureOrSuggestiveThemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetMatureOrSuggestiveThemesOk() (*string, bool) {
	if o == nil || IsNil(o.MatureOrSuggestiveThemes) {
		return nil, false
	}
	return o.MatureOrSuggestiveThemes, true
}

// HasMatureOrSuggestiveThemes returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasMatureOrSuggestiveThemes() bool {
	if o != nil && !IsNil(o.MatureOrSuggestiveThemes) {
		return true
	}

	return false
}

// SetMatureOrSuggestiveThemes gets a reference to the given string and assigns it to the MatureOrSuggestiveThemes field.
func (o *AgeRatingDeclarationAttributes) SetMatureOrSuggestiveThemes(v string) {
	o.MatureOrSuggestiveThemes = &v
}

// GetUnrestrictedWebAccess returns the UnrestrictedWebAccess field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetUnrestrictedWebAccess() bool {
	if o == nil || IsNil(o.UnrestrictedWebAccess) {
		var ret bool
		return ret
	}
	return *o.UnrestrictedWebAccess
}

// GetUnrestrictedWebAccessOk returns a tuple with the UnrestrictedWebAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetUnrestrictedWebAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.UnrestrictedWebAccess) {
		return nil, false
	}
	return o.UnrestrictedWebAccess, true
}

// HasUnrestrictedWebAccess returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasUnrestrictedWebAccess() bool {
	if o != nil && !IsNil(o.UnrestrictedWebAccess) {
		return true
	}

	return false
}

// SetUnrestrictedWebAccess gets a reference to the given bool and assigns it to the UnrestrictedWebAccess field.
func (o *AgeRatingDeclarationAttributes) SetUnrestrictedWebAccess(v bool) {
	o.UnrestrictedWebAccess = &v
}

// GetViolenceCartoonOrFantasy returns the ViolenceCartoonOrFantasy field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetViolenceCartoonOrFantasy() string {
	if o == nil || IsNil(o.ViolenceCartoonOrFantasy) {
		var ret string
		return ret
	}
	return *o.ViolenceCartoonOrFantasy
}

// GetViolenceCartoonOrFantasyOk returns a tuple with the ViolenceCartoonOrFantasy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetViolenceCartoonOrFantasyOk() (*string, bool) {
	if o == nil || IsNil(o.ViolenceCartoonOrFantasy) {
		return nil, false
	}
	return o.ViolenceCartoonOrFantasy, true
}

// HasViolenceCartoonOrFantasy returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasViolenceCartoonOrFantasy() bool {
	if o != nil && !IsNil(o.ViolenceCartoonOrFantasy) {
		return true
	}

	return false
}

// SetViolenceCartoonOrFantasy gets a reference to the given string and assigns it to the ViolenceCartoonOrFantasy field.
func (o *AgeRatingDeclarationAttributes) SetViolenceCartoonOrFantasy(v string) {
	o.ViolenceCartoonOrFantasy = &v
}

// GetViolenceRealisticProlongedGraphicOrSadistic returns the ViolenceRealisticProlongedGraphicOrSadistic field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetViolenceRealisticProlongedGraphicOrSadistic() string {
	if o == nil || IsNil(o.ViolenceRealisticProlongedGraphicOrSadistic) {
		var ret string
		return ret
	}
	return *o.ViolenceRealisticProlongedGraphicOrSadistic
}

// GetViolenceRealisticProlongedGraphicOrSadisticOk returns a tuple with the ViolenceRealisticProlongedGraphicOrSadistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetViolenceRealisticProlongedGraphicOrSadisticOk() (*string, bool) {
	if o == nil || IsNil(o.ViolenceRealisticProlongedGraphicOrSadistic) {
		return nil, false
	}
	return o.ViolenceRealisticProlongedGraphicOrSadistic, true
}

// HasViolenceRealisticProlongedGraphicOrSadistic returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasViolenceRealisticProlongedGraphicOrSadistic() bool {
	if o != nil && !IsNil(o.ViolenceRealisticProlongedGraphicOrSadistic) {
		return true
	}

	return false
}

// SetViolenceRealisticProlongedGraphicOrSadistic gets a reference to the given string and assigns it to the ViolenceRealisticProlongedGraphicOrSadistic field.
func (o *AgeRatingDeclarationAttributes) SetViolenceRealisticProlongedGraphicOrSadistic(v string) {
	o.ViolenceRealisticProlongedGraphicOrSadistic = &v
}

// GetViolenceRealistic returns the ViolenceRealistic field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetViolenceRealistic() string {
	if o == nil || IsNil(o.ViolenceRealistic) {
		var ret string
		return ret
	}
	return *o.ViolenceRealistic
}

// GetViolenceRealisticOk returns a tuple with the ViolenceRealistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetViolenceRealisticOk() (*string, bool) {
	if o == nil || IsNil(o.ViolenceRealistic) {
		return nil, false
	}
	return o.ViolenceRealistic, true
}

// HasViolenceRealistic returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasViolenceRealistic() bool {
	if o != nil && !IsNil(o.ViolenceRealistic) {
		return true
	}

	return false
}

// SetViolenceRealistic gets a reference to the given string and assigns it to the ViolenceRealistic field.
func (o *AgeRatingDeclarationAttributes) SetViolenceRealistic(v string) {
	o.ViolenceRealistic = &v
}

// GetAgeRatingOverride returns the AgeRatingOverride field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetAgeRatingOverride() string {
	if o == nil || IsNil(o.AgeRatingOverride) {
		var ret string
		return ret
	}
	return *o.AgeRatingOverride
}

// GetAgeRatingOverrideOk returns a tuple with the AgeRatingOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetAgeRatingOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.AgeRatingOverride) {
		return nil, false
	}
	return o.AgeRatingOverride, true
}

// HasAgeRatingOverride returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasAgeRatingOverride() bool {
	if o != nil && !IsNil(o.AgeRatingOverride) {
		return true
	}

	return false
}

// SetAgeRatingOverride gets a reference to the given string and assigns it to the AgeRatingOverride field.
func (o *AgeRatingDeclarationAttributes) SetAgeRatingOverride(v string) {
	o.AgeRatingOverride = &v
}

// GetKoreaAgeRatingOverride returns the KoreaAgeRatingOverride field value if set, zero value otherwise.
func (o *AgeRatingDeclarationAttributes) GetKoreaAgeRatingOverride() string {
	if o == nil || IsNil(o.KoreaAgeRatingOverride) {
		var ret string
		return ret
	}
	return *o.KoreaAgeRatingOverride
}

// GetKoreaAgeRatingOverrideOk returns a tuple with the KoreaAgeRatingOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationAttributes) GetKoreaAgeRatingOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.KoreaAgeRatingOverride) {
		return nil, false
	}
	return o.KoreaAgeRatingOverride, true
}

// HasKoreaAgeRatingOverride returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasKoreaAgeRatingOverride() bool {
	if o != nil && !IsNil(o.KoreaAgeRatingOverride) {
		return true
	}

	return false
}

// SetKoreaAgeRatingOverride gets a reference to the given string and assigns it to the KoreaAgeRatingOverride field.
func (o *AgeRatingDeclarationAttributes) SetKoreaAgeRatingOverride(v string) {
	o.KoreaAgeRatingOverride = &v
}

// GetSeventeenPlus returns the SeventeenPlus field value if set, zero value otherwise.
// Deprecated
func (o *AgeRatingDeclarationAttributes) GetSeventeenPlus() bool {
	if o == nil || IsNil(o.SeventeenPlus) {
		var ret bool
		return ret
	}
	return *o.SeventeenPlus
}

// GetSeventeenPlusOk returns a tuple with the SeventeenPlus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AgeRatingDeclarationAttributes) GetSeventeenPlusOk() (*bool, bool) {
	if o == nil || IsNil(o.SeventeenPlus) {
		return nil, false
	}
	return o.SeventeenPlus, true
}

// HasSeventeenPlus returns a boolean if a field has been set.
func (o *AgeRatingDeclarationAttributes) HasSeventeenPlus() bool {
	if o != nil && !IsNil(o.SeventeenPlus) {
		return true
	}

	return false
}

// SetSeventeenPlus gets a reference to the given bool and assigns it to the SeventeenPlus field.
// Deprecated
func (o *AgeRatingDeclarationAttributes) SetSeventeenPlus(v bool) {
	o.SeventeenPlus = &v
}

func (o AgeRatingDeclarationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgeRatingDeclarationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlcoholTobaccoOrDrugUseOrReferences) {
		toSerialize["alcoholTobaccoOrDrugUseOrReferences"] = o.AlcoholTobaccoOrDrugUseOrReferences
	}
	if !IsNil(o.Contests) {
		toSerialize["contests"] = o.Contests
	}
	if !IsNil(o.GamblingAndContests) {
		toSerialize["gamblingAndContests"] = o.GamblingAndContests
	}
	if !IsNil(o.Gambling) {
		toSerialize["gambling"] = o.Gambling
	}
	if !IsNil(o.GamblingSimulated) {
		toSerialize["gamblingSimulated"] = o.GamblingSimulated
	}
	if !IsNil(o.KidsAgeBand) {
		toSerialize["kidsAgeBand"] = o.KidsAgeBand
	}
	if !IsNil(o.LootBox) {
		toSerialize["lootBox"] = o.LootBox
	}
	if !IsNil(o.MedicalOrTreatmentInformation) {
		toSerialize["medicalOrTreatmentInformation"] = o.MedicalOrTreatmentInformation
	}
	if !IsNil(o.ProfanityOrCrudeHumor) {
		toSerialize["profanityOrCrudeHumor"] = o.ProfanityOrCrudeHumor
	}
	if !IsNil(o.SexualContentGraphicAndNudity) {
		toSerialize["sexualContentGraphicAndNudity"] = o.SexualContentGraphicAndNudity
	}
	if !IsNil(o.SexualContentOrNudity) {
		toSerialize["sexualContentOrNudity"] = o.SexualContentOrNudity
	}
	if !IsNil(o.HorrorOrFearThemes) {
		toSerialize["horrorOrFearThemes"] = o.HorrorOrFearThemes
	}
	if !IsNil(o.MatureOrSuggestiveThemes) {
		toSerialize["matureOrSuggestiveThemes"] = o.MatureOrSuggestiveThemes
	}
	if !IsNil(o.UnrestrictedWebAccess) {
		toSerialize["unrestrictedWebAccess"] = o.UnrestrictedWebAccess
	}
	if !IsNil(o.ViolenceCartoonOrFantasy) {
		toSerialize["violenceCartoonOrFantasy"] = o.ViolenceCartoonOrFantasy
	}
	if !IsNil(o.ViolenceRealisticProlongedGraphicOrSadistic) {
		toSerialize["violenceRealisticProlongedGraphicOrSadistic"] = o.ViolenceRealisticProlongedGraphicOrSadistic
	}
	if !IsNil(o.ViolenceRealistic) {
		toSerialize["violenceRealistic"] = o.ViolenceRealistic
	}
	if !IsNil(o.AgeRatingOverride) {
		toSerialize["ageRatingOverride"] = o.AgeRatingOverride
	}
	if !IsNil(o.KoreaAgeRatingOverride) {
		toSerialize["koreaAgeRatingOverride"] = o.KoreaAgeRatingOverride
	}
	if !IsNil(o.SeventeenPlus) {
		toSerialize["seventeenPlus"] = o.SeventeenPlus
	}
	return toSerialize, nil
}

type NullableAgeRatingDeclarationAttributes struct {
	value *AgeRatingDeclarationAttributes
	isSet bool
}

func (v NullableAgeRatingDeclarationAttributes) Get() *AgeRatingDeclarationAttributes {
	return v.value
}

func (v *NullableAgeRatingDeclarationAttributes) Set(val *AgeRatingDeclarationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAgeRatingDeclarationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAgeRatingDeclarationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgeRatingDeclarationAttributes(val *AgeRatingDeclarationAttributes) *NullableAgeRatingDeclarationAttributes {
	return &NullableAgeRatingDeclarationAttributes{value: val, isSet: true}
}

func (v NullableAgeRatingDeclarationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgeRatingDeclarationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


