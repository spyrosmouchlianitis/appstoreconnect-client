# AppAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**BundleId** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**PrimaryLocale** | Pointer to **string** |  | [optional] 
**IsOrEverWasMadeForKids** | Pointer to **bool** |  | [optional] 
**SubscriptionStatusUrl** | Pointer to **string** |  | [optional] 
**SubscriptionStatusUrlVersion** | Pointer to [**SubscriptionStatusUrlVersion**](SubscriptionStatusUrlVersion.md) |  | [optional] 
**SubscriptionStatusUrlForSandbox** | Pointer to **string** |  | [optional] 
**SubscriptionStatusUrlVersionForSandbox** | Pointer to [**SubscriptionStatusUrlVersion**](SubscriptionStatusUrlVersion.md) |  | [optional] 
**ContentRightsDeclaration** | Pointer to **string** |  | [optional] 
**StreamlinedPurchasingEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAppAttributes

`func NewAppAttributes() *AppAttributes`

NewAppAttributes instantiates a new AppAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAttributesWithDefaults

`func NewAppAttributesWithDefaults() *AppAttributes`

NewAppAttributesWithDefaults instantiates a new AppAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBundleId

`func (o *AppAttributes) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *AppAttributes) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *AppAttributes) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *AppAttributes) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetSku

`func (o *AppAttributes) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AppAttributes) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AppAttributes) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *AppAttributes) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPrimaryLocale

`func (o *AppAttributes) GetPrimaryLocale() string`

GetPrimaryLocale returns the PrimaryLocale field if non-nil, zero value otherwise.

### GetPrimaryLocaleOk

`func (o *AppAttributes) GetPrimaryLocaleOk() (*string, bool)`

GetPrimaryLocaleOk returns a tuple with the PrimaryLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLocale

`func (o *AppAttributes) SetPrimaryLocale(v string)`

SetPrimaryLocale sets PrimaryLocale field to given value.

### HasPrimaryLocale

`func (o *AppAttributes) HasPrimaryLocale() bool`

HasPrimaryLocale returns a boolean if a field has been set.

### GetIsOrEverWasMadeForKids

`func (o *AppAttributes) GetIsOrEverWasMadeForKids() bool`

GetIsOrEverWasMadeForKids returns the IsOrEverWasMadeForKids field if non-nil, zero value otherwise.

### GetIsOrEverWasMadeForKidsOk

`func (o *AppAttributes) GetIsOrEverWasMadeForKidsOk() (*bool, bool)`

GetIsOrEverWasMadeForKidsOk returns a tuple with the IsOrEverWasMadeForKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrEverWasMadeForKids

`func (o *AppAttributes) SetIsOrEverWasMadeForKids(v bool)`

SetIsOrEverWasMadeForKids sets IsOrEverWasMadeForKids field to given value.

### HasIsOrEverWasMadeForKids

`func (o *AppAttributes) HasIsOrEverWasMadeForKids() bool`

HasIsOrEverWasMadeForKids returns a boolean if a field has been set.

### GetSubscriptionStatusUrl

`func (o *AppAttributes) GetSubscriptionStatusUrl() string`

GetSubscriptionStatusUrl returns the SubscriptionStatusUrl field if non-nil, zero value otherwise.

### GetSubscriptionStatusUrlOk

`func (o *AppAttributes) GetSubscriptionStatusUrlOk() (*string, bool)`

GetSubscriptionStatusUrlOk returns a tuple with the SubscriptionStatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusUrl

`func (o *AppAttributes) SetSubscriptionStatusUrl(v string)`

SetSubscriptionStatusUrl sets SubscriptionStatusUrl field to given value.

### HasSubscriptionStatusUrl

`func (o *AppAttributes) HasSubscriptionStatusUrl() bool`

HasSubscriptionStatusUrl returns a boolean if a field has been set.

### GetSubscriptionStatusUrlVersion

`func (o *AppAttributes) GetSubscriptionStatusUrlVersion() SubscriptionStatusUrlVersion`

GetSubscriptionStatusUrlVersion returns the SubscriptionStatusUrlVersion field if non-nil, zero value otherwise.

### GetSubscriptionStatusUrlVersionOk

`func (o *AppAttributes) GetSubscriptionStatusUrlVersionOk() (*SubscriptionStatusUrlVersion, bool)`

GetSubscriptionStatusUrlVersionOk returns a tuple with the SubscriptionStatusUrlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusUrlVersion

`func (o *AppAttributes) SetSubscriptionStatusUrlVersion(v SubscriptionStatusUrlVersion)`

SetSubscriptionStatusUrlVersion sets SubscriptionStatusUrlVersion field to given value.

### HasSubscriptionStatusUrlVersion

`func (o *AppAttributes) HasSubscriptionStatusUrlVersion() bool`

HasSubscriptionStatusUrlVersion returns a boolean if a field has been set.

### GetSubscriptionStatusUrlForSandbox

`func (o *AppAttributes) GetSubscriptionStatusUrlForSandbox() string`

GetSubscriptionStatusUrlForSandbox returns the SubscriptionStatusUrlForSandbox field if non-nil, zero value otherwise.

### GetSubscriptionStatusUrlForSandboxOk

`func (o *AppAttributes) GetSubscriptionStatusUrlForSandboxOk() (*string, bool)`

GetSubscriptionStatusUrlForSandboxOk returns a tuple with the SubscriptionStatusUrlForSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusUrlForSandbox

`func (o *AppAttributes) SetSubscriptionStatusUrlForSandbox(v string)`

SetSubscriptionStatusUrlForSandbox sets SubscriptionStatusUrlForSandbox field to given value.

### HasSubscriptionStatusUrlForSandbox

`func (o *AppAttributes) HasSubscriptionStatusUrlForSandbox() bool`

HasSubscriptionStatusUrlForSandbox returns a boolean if a field has been set.

### GetSubscriptionStatusUrlVersionForSandbox

`func (o *AppAttributes) GetSubscriptionStatusUrlVersionForSandbox() SubscriptionStatusUrlVersion`

GetSubscriptionStatusUrlVersionForSandbox returns the SubscriptionStatusUrlVersionForSandbox field if non-nil, zero value otherwise.

### GetSubscriptionStatusUrlVersionForSandboxOk

`func (o *AppAttributes) GetSubscriptionStatusUrlVersionForSandboxOk() (*SubscriptionStatusUrlVersion, bool)`

GetSubscriptionStatusUrlVersionForSandboxOk returns a tuple with the SubscriptionStatusUrlVersionForSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatusUrlVersionForSandbox

`func (o *AppAttributes) SetSubscriptionStatusUrlVersionForSandbox(v SubscriptionStatusUrlVersion)`

SetSubscriptionStatusUrlVersionForSandbox sets SubscriptionStatusUrlVersionForSandbox field to given value.

### HasSubscriptionStatusUrlVersionForSandbox

`func (o *AppAttributes) HasSubscriptionStatusUrlVersionForSandbox() bool`

HasSubscriptionStatusUrlVersionForSandbox returns a boolean if a field has been set.

### GetContentRightsDeclaration

`func (o *AppAttributes) GetContentRightsDeclaration() string`

GetContentRightsDeclaration returns the ContentRightsDeclaration field if non-nil, zero value otherwise.

### GetContentRightsDeclarationOk

`func (o *AppAttributes) GetContentRightsDeclarationOk() (*string, bool)`

GetContentRightsDeclarationOk returns a tuple with the ContentRightsDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRightsDeclaration

`func (o *AppAttributes) SetContentRightsDeclaration(v string)`

SetContentRightsDeclaration sets ContentRightsDeclaration field to given value.

### HasContentRightsDeclaration

`func (o *AppAttributes) HasContentRightsDeclaration() bool`

HasContentRightsDeclaration returns a boolean if a field has been set.

### GetStreamlinedPurchasingEnabled

`func (o *AppAttributes) GetStreamlinedPurchasingEnabled() bool`

GetStreamlinedPurchasingEnabled returns the StreamlinedPurchasingEnabled field if non-nil, zero value otherwise.

### GetStreamlinedPurchasingEnabledOk

`func (o *AppAttributes) GetStreamlinedPurchasingEnabledOk() (*bool, bool)`

GetStreamlinedPurchasingEnabledOk returns a tuple with the StreamlinedPurchasingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamlinedPurchasingEnabled

`func (o *AppAttributes) SetStreamlinedPurchasingEnabled(v bool)`

SetStreamlinedPurchasingEnabled sets StreamlinedPurchasingEnabled field to given value.

### HasStreamlinedPurchasingEnabled

`func (o *AppAttributes) HasStreamlinedPurchasingEnabled() bool`

HasStreamlinedPurchasingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


