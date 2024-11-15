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

// checks if the GameCenterMatchmakingRuleAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleAttributes{}

// GameCenterMatchmakingRuleAttributes struct for GameCenterMatchmakingRuleAttributes
type GameCenterMatchmakingRuleAttributes struct {
	ReferenceName *string `json:"referenceName,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Expression *string `json:"expression,omitempty"`
	Weight *float32 `json:"weight,omitempty"`
}

// NewGameCenterMatchmakingRuleAttributes instantiates a new GameCenterMatchmakingRuleAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleAttributes() *GameCenterMatchmakingRuleAttributes {
	this := GameCenterMatchmakingRuleAttributes{}
	return &this
}

// NewGameCenterMatchmakingRuleAttributesWithDefaults instantiates a new GameCenterMatchmakingRuleAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleAttributesWithDefaults() *GameCenterMatchmakingRuleAttributes {
	this := GameCenterMatchmakingRuleAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleAttributes) GetReferenceName() string {
	if o == nil || IsNil(o.ReferenceName) {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceName) {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleAttributes) HasReferenceName() bool {
	if o != nil && !IsNil(o.ReferenceName) {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *GameCenterMatchmakingRuleAttributes) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleAttributes) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GameCenterMatchmakingRuleAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleAttributes) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleAttributes) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleAttributes) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GameCenterMatchmakingRuleAttributes) SetType(v string) {
	o.Type = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleAttributes) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleAttributes) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleAttributes) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *GameCenterMatchmakingRuleAttributes) SetExpression(v string) {
	o.Expression = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleAttributes) GetWeight() float32 {
	if o == nil || IsNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleAttributes) GetWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleAttributes) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *GameCenterMatchmakingRuleAttributes) SetWeight(v float32) {
	o.Weight = &v
}

func (o GameCenterMatchmakingRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceName) {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingRuleAttributes struct {
	value *GameCenterMatchmakingRuleAttributes
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleAttributes) Get() *GameCenterMatchmakingRuleAttributes {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleAttributes) Set(val *GameCenterMatchmakingRuleAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleAttributes(val *GameCenterMatchmakingRuleAttributes) *NullableGameCenterMatchmakingRuleAttributes {
	return &NullableGameCenterMatchmakingRuleAttributes{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


