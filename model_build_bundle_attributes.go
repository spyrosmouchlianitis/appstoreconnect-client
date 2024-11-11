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

// checks if the BuildBundleAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBundleAttributes{}

// BuildBundleAttributes struct for BuildBundleAttributes
type BuildBundleAttributes struct {
	BundleId *string `json:"bundleId,omitempty"`
	BundleType *string `json:"bundleType,omitempty"`
	SdkBuild *string `json:"sdkBuild,omitempty"`
	PlatformBuild *string `json:"platformBuild,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	HasSirikit *bool `json:"hasSirikit,omitempty"`
	HasOnDemandResources *bool `json:"hasOnDemandResources,omitempty"`
	HasPrerenderedIcon *bool `json:"hasPrerenderedIcon,omitempty"`
	UsesLocationServices *bool `json:"usesLocationServices,omitempty"`
	IsIosBuildMacAppStoreCompatible *bool `json:"isIosBuildMacAppStoreCompatible,omitempty"`
	IncludesSymbols *bool `json:"includesSymbols,omitempty"`
	DSYMUrl *string `json:"dSYMUrl,omitempty"`
	SupportedArchitectures []string `json:"supportedArchitectures,omitempty"`
	RequiredCapabilities []string `json:"requiredCapabilities,omitempty"`
	DeviceProtocols []string `json:"deviceProtocols,omitempty"`
	Locales []string `json:"locales,omitempty"`
	Entitlements *map[string]map[string]string `json:"entitlements,omitempty"`
}

// NewBuildBundleAttributes instantiates a new BuildBundleAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBundleAttributes() *BuildBundleAttributes {
	this := BuildBundleAttributes{}
	return &this
}

// NewBuildBundleAttributesWithDefaults instantiates a new BuildBundleAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBundleAttributesWithDefaults() *BuildBundleAttributes {
	this := BuildBundleAttributes{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *BuildBundleAttributes) SetBundleId(v string) {
	o.BundleId = &v
}

// GetBundleType returns the BundleType field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetBundleType() string {
	if o == nil || IsNil(o.BundleType) {
		var ret string
		return ret
	}
	return *o.BundleType
}

// GetBundleTypeOk returns a tuple with the BundleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetBundleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BundleType) {
		return nil, false
	}
	return o.BundleType, true
}

// HasBundleType returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasBundleType() bool {
	if o != nil && !IsNil(o.BundleType) {
		return true
	}

	return false
}

// SetBundleType gets a reference to the given string and assigns it to the BundleType field.
func (o *BuildBundleAttributes) SetBundleType(v string) {
	o.BundleType = &v
}

// GetSdkBuild returns the SdkBuild field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetSdkBuild() string {
	if o == nil || IsNil(o.SdkBuild) {
		var ret string
		return ret
	}
	return *o.SdkBuild
}

// GetSdkBuildOk returns a tuple with the SdkBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetSdkBuildOk() (*string, bool) {
	if o == nil || IsNil(o.SdkBuild) {
		return nil, false
	}
	return o.SdkBuild, true
}

// HasSdkBuild returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasSdkBuild() bool {
	if o != nil && !IsNil(o.SdkBuild) {
		return true
	}

	return false
}

// SetSdkBuild gets a reference to the given string and assigns it to the SdkBuild field.
func (o *BuildBundleAttributes) SetSdkBuild(v string) {
	o.SdkBuild = &v
}

// GetPlatformBuild returns the PlatformBuild field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetPlatformBuild() string {
	if o == nil || IsNil(o.PlatformBuild) {
		var ret string
		return ret
	}
	return *o.PlatformBuild
}

// GetPlatformBuildOk returns a tuple with the PlatformBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetPlatformBuildOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformBuild) {
		return nil, false
	}
	return o.PlatformBuild, true
}

// HasPlatformBuild returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasPlatformBuild() bool {
	if o != nil && !IsNil(o.PlatformBuild) {
		return true
	}

	return false
}

// SetPlatformBuild gets a reference to the given string and assigns it to the PlatformBuild field.
func (o *BuildBundleAttributes) SetPlatformBuild(v string) {
	o.PlatformBuild = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *BuildBundleAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetHasSirikit returns the HasSirikit field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetHasSirikit() bool {
	if o == nil || IsNil(o.HasSirikit) {
		var ret bool
		return ret
	}
	return *o.HasSirikit
}

// GetHasSirikitOk returns a tuple with the HasSirikit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetHasSirikitOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSirikit) {
		return nil, false
	}
	return o.HasSirikit, true
}

// HasHasSirikit returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasHasSirikit() bool {
	if o != nil && !IsNil(o.HasSirikit) {
		return true
	}

	return false
}

// SetHasSirikit gets a reference to the given bool and assigns it to the HasSirikit field.
func (o *BuildBundleAttributes) SetHasSirikit(v bool) {
	o.HasSirikit = &v
}

// GetHasOnDemandResources returns the HasOnDemandResources field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetHasOnDemandResources() bool {
	if o == nil || IsNil(o.HasOnDemandResources) {
		var ret bool
		return ret
	}
	return *o.HasOnDemandResources
}

// GetHasOnDemandResourcesOk returns a tuple with the HasOnDemandResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetHasOnDemandResourcesOk() (*bool, bool) {
	if o == nil || IsNil(o.HasOnDemandResources) {
		return nil, false
	}
	return o.HasOnDemandResources, true
}

// HasHasOnDemandResources returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasHasOnDemandResources() bool {
	if o != nil && !IsNil(o.HasOnDemandResources) {
		return true
	}

	return false
}

// SetHasOnDemandResources gets a reference to the given bool and assigns it to the HasOnDemandResources field.
func (o *BuildBundleAttributes) SetHasOnDemandResources(v bool) {
	o.HasOnDemandResources = &v
}

// GetHasPrerenderedIcon returns the HasPrerenderedIcon field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetHasPrerenderedIcon() bool {
	if o == nil || IsNil(o.HasPrerenderedIcon) {
		var ret bool
		return ret
	}
	return *o.HasPrerenderedIcon
}

// GetHasPrerenderedIconOk returns a tuple with the HasPrerenderedIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetHasPrerenderedIconOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPrerenderedIcon) {
		return nil, false
	}
	return o.HasPrerenderedIcon, true
}

// HasHasPrerenderedIcon returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasHasPrerenderedIcon() bool {
	if o != nil && !IsNil(o.HasPrerenderedIcon) {
		return true
	}

	return false
}

// SetHasPrerenderedIcon gets a reference to the given bool and assigns it to the HasPrerenderedIcon field.
func (o *BuildBundleAttributes) SetHasPrerenderedIcon(v bool) {
	o.HasPrerenderedIcon = &v
}

// GetUsesLocationServices returns the UsesLocationServices field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetUsesLocationServices() bool {
	if o == nil || IsNil(o.UsesLocationServices) {
		var ret bool
		return ret
	}
	return *o.UsesLocationServices
}

// GetUsesLocationServicesOk returns a tuple with the UsesLocationServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetUsesLocationServicesOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesLocationServices) {
		return nil, false
	}
	return o.UsesLocationServices, true
}

// HasUsesLocationServices returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasUsesLocationServices() bool {
	if o != nil && !IsNil(o.UsesLocationServices) {
		return true
	}

	return false
}

// SetUsesLocationServices gets a reference to the given bool and assigns it to the UsesLocationServices field.
func (o *BuildBundleAttributes) SetUsesLocationServices(v bool) {
	o.UsesLocationServices = &v
}

// GetIsIosBuildMacAppStoreCompatible returns the IsIosBuildMacAppStoreCompatible field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetIsIosBuildMacAppStoreCompatible() bool {
	if o == nil || IsNil(o.IsIosBuildMacAppStoreCompatible) {
		var ret bool
		return ret
	}
	return *o.IsIosBuildMacAppStoreCompatible
}

// GetIsIosBuildMacAppStoreCompatibleOk returns a tuple with the IsIosBuildMacAppStoreCompatible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetIsIosBuildMacAppStoreCompatibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIosBuildMacAppStoreCompatible) {
		return nil, false
	}
	return o.IsIosBuildMacAppStoreCompatible, true
}

// HasIsIosBuildMacAppStoreCompatible returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasIsIosBuildMacAppStoreCompatible() bool {
	if o != nil && !IsNil(o.IsIosBuildMacAppStoreCompatible) {
		return true
	}

	return false
}

// SetIsIosBuildMacAppStoreCompatible gets a reference to the given bool and assigns it to the IsIosBuildMacAppStoreCompatible field.
func (o *BuildBundleAttributes) SetIsIosBuildMacAppStoreCompatible(v bool) {
	o.IsIosBuildMacAppStoreCompatible = &v
}

// GetIncludesSymbols returns the IncludesSymbols field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetIncludesSymbols() bool {
	if o == nil || IsNil(o.IncludesSymbols) {
		var ret bool
		return ret
	}
	return *o.IncludesSymbols
}

// GetIncludesSymbolsOk returns a tuple with the IncludesSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetIncludesSymbolsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludesSymbols) {
		return nil, false
	}
	return o.IncludesSymbols, true
}

// HasIncludesSymbols returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasIncludesSymbols() bool {
	if o != nil && !IsNil(o.IncludesSymbols) {
		return true
	}

	return false
}

// SetIncludesSymbols gets a reference to the given bool and assigns it to the IncludesSymbols field.
func (o *BuildBundleAttributes) SetIncludesSymbols(v bool) {
	o.IncludesSymbols = &v
}

// GetDSYMUrl returns the DSYMUrl field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetDSYMUrl() string {
	if o == nil || IsNil(o.DSYMUrl) {
		var ret string
		return ret
	}
	return *o.DSYMUrl
}

// GetDSYMUrlOk returns a tuple with the DSYMUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetDSYMUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DSYMUrl) {
		return nil, false
	}
	return o.DSYMUrl, true
}

// HasDSYMUrl returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasDSYMUrl() bool {
	if o != nil && !IsNil(o.DSYMUrl) {
		return true
	}

	return false
}

// SetDSYMUrl gets a reference to the given string and assigns it to the DSYMUrl field.
func (o *BuildBundleAttributes) SetDSYMUrl(v string) {
	o.DSYMUrl = &v
}

// GetSupportedArchitectures returns the SupportedArchitectures field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetSupportedArchitectures() []string {
	if o == nil || IsNil(o.SupportedArchitectures) {
		var ret []string
		return ret
	}
	return o.SupportedArchitectures
}

// GetSupportedArchitecturesOk returns a tuple with the SupportedArchitectures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetSupportedArchitecturesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedArchitectures) {
		return nil, false
	}
	return o.SupportedArchitectures, true
}

// HasSupportedArchitectures returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasSupportedArchitectures() bool {
	if o != nil && !IsNil(o.SupportedArchitectures) {
		return true
	}

	return false
}

// SetSupportedArchitectures gets a reference to the given []string and assigns it to the SupportedArchitectures field.
func (o *BuildBundleAttributes) SetSupportedArchitectures(v []string) {
	o.SupportedArchitectures = v
}

// GetRequiredCapabilities returns the RequiredCapabilities field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetRequiredCapabilities() []string {
	if o == nil || IsNil(o.RequiredCapabilities) {
		var ret []string
		return ret
	}
	return o.RequiredCapabilities
}

// GetRequiredCapabilitiesOk returns a tuple with the RequiredCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetRequiredCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.RequiredCapabilities) {
		return nil, false
	}
	return o.RequiredCapabilities, true
}

// HasRequiredCapabilities returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasRequiredCapabilities() bool {
	if o != nil && !IsNil(o.RequiredCapabilities) {
		return true
	}

	return false
}

// SetRequiredCapabilities gets a reference to the given []string and assigns it to the RequiredCapabilities field.
func (o *BuildBundleAttributes) SetRequiredCapabilities(v []string) {
	o.RequiredCapabilities = v
}

// GetDeviceProtocols returns the DeviceProtocols field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetDeviceProtocols() []string {
	if o == nil || IsNil(o.DeviceProtocols) {
		var ret []string
		return ret
	}
	return o.DeviceProtocols
}

// GetDeviceProtocolsOk returns a tuple with the DeviceProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetDeviceProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceProtocols) {
		return nil, false
	}
	return o.DeviceProtocols, true
}

// HasDeviceProtocols returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasDeviceProtocols() bool {
	if o != nil && !IsNil(o.DeviceProtocols) {
		return true
	}

	return false
}

// SetDeviceProtocols gets a reference to the given []string and assigns it to the DeviceProtocols field.
func (o *BuildBundleAttributes) SetDeviceProtocols(v []string) {
	o.DeviceProtocols = v
}

// GetLocales returns the Locales field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetLocales() []string {
	if o == nil || IsNil(o.Locales) {
		var ret []string
		return ret
	}
	return o.Locales
}

// GetLocalesOk returns a tuple with the Locales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetLocalesOk() ([]string, bool) {
	if o == nil || IsNil(o.Locales) {
		return nil, false
	}
	return o.Locales, true
}

// HasLocales returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasLocales() bool {
	if o != nil && !IsNil(o.Locales) {
		return true
	}

	return false
}

// SetLocales gets a reference to the given []string and assigns it to the Locales field.
func (o *BuildBundleAttributes) SetLocales(v []string) {
	o.Locales = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *BuildBundleAttributes) GetEntitlements() map[string]map[string]string {
	if o == nil || IsNil(o.Entitlements) {
		var ret map[string]map[string]string
		return ret
	}
	return *o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleAttributes) GetEntitlementsOk() (*map[string]map[string]string, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *BuildBundleAttributes) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given map[string]map[string]string and assigns it to the Entitlements field.
func (o *BuildBundleAttributes) SetEntitlements(v map[string]map[string]string) {
	o.Entitlements = &v
}

func (o BuildBundleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBundleAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundleId"] = o.BundleId
	}
	if !IsNil(o.BundleType) {
		toSerialize["bundleType"] = o.BundleType
	}
	if !IsNil(o.SdkBuild) {
		toSerialize["sdkBuild"] = o.SdkBuild
	}
	if !IsNil(o.PlatformBuild) {
		toSerialize["platformBuild"] = o.PlatformBuild
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.HasSirikit) {
		toSerialize["hasSirikit"] = o.HasSirikit
	}
	if !IsNil(o.HasOnDemandResources) {
		toSerialize["hasOnDemandResources"] = o.HasOnDemandResources
	}
	if !IsNil(o.HasPrerenderedIcon) {
		toSerialize["hasPrerenderedIcon"] = o.HasPrerenderedIcon
	}
	if !IsNil(o.UsesLocationServices) {
		toSerialize["usesLocationServices"] = o.UsesLocationServices
	}
	if !IsNil(o.IsIosBuildMacAppStoreCompatible) {
		toSerialize["isIosBuildMacAppStoreCompatible"] = o.IsIosBuildMacAppStoreCompatible
	}
	if !IsNil(o.IncludesSymbols) {
		toSerialize["includesSymbols"] = o.IncludesSymbols
	}
	if !IsNil(o.DSYMUrl) {
		toSerialize["dSYMUrl"] = o.DSYMUrl
	}
	if !IsNil(o.SupportedArchitectures) {
		toSerialize["supportedArchitectures"] = o.SupportedArchitectures
	}
	if !IsNil(o.RequiredCapabilities) {
		toSerialize["requiredCapabilities"] = o.RequiredCapabilities
	}
	if !IsNil(o.DeviceProtocols) {
		toSerialize["deviceProtocols"] = o.DeviceProtocols
	}
	if !IsNil(o.Locales) {
		toSerialize["locales"] = o.Locales
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	return toSerialize, nil
}

type NullableBuildBundleAttributes struct {
	value *BuildBundleAttributes
	isSet bool
}

func (v NullableBuildBundleAttributes) Get() *BuildBundleAttributes {
	return v.value
}

func (v *NullableBuildBundleAttributes) Set(val *BuildBundleAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBundleAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBundleAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBundleAttributes(val *BuildBundleAttributes) *NullableBuildBundleAttributes {
	return &NullableBuildBundleAttributes{value: val, isSet: true}
}

func (v NullableBuildBundleAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBundleAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


