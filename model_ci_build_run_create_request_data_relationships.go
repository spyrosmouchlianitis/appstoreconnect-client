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

// checks if the CiBuildRunCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunCreateRequestDataRelationships{}

// CiBuildRunCreateRequestDataRelationships struct for CiBuildRunCreateRequestDataRelationships
type CiBuildRunCreateRequestDataRelationships struct {
	BuildRun *CiBuildRunCreateRequestDataRelationshipsBuildRun `json:"buildRun,omitempty"`
	Workflow *CiBuildRunRelationshipsWorkflow `json:"workflow,omitempty"`
	SourceBranchOrTag *CiBuildRunRelationshipsSourceBranchOrTag `json:"sourceBranchOrTag,omitempty"`
	PullRequest *CiBuildRunRelationshipsPullRequest `json:"pullRequest,omitempty"`
}

// NewCiBuildRunCreateRequestDataRelationships instantiates a new CiBuildRunCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunCreateRequestDataRelationships() *CiBuildRunCreateRequestDataRelationships {
	this := CiBuildRunCreateRequestDataRelationships{}
	return &this
}

// NewCiBuildRunCreateRequestDataRelationshipsWithDefaults instantiates a new CiBuildRunCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunCreateRequestDataRelationshipsWithDefaults() *CiBuildRunCreateRequestDataRelationships {
	this := CiBuildRunCreateRequestDataRelationships{}
	return &this
}

// GetBuildRun returns the BuildRun field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestDataRelationships) GetBuildRun() CiBuildRunCreateRequestDataRelationshipsBuildRun {
	if o == nil || IsNil(o.BuildRun) {
		var ret CiBuildRunCreateRequestDataRelationshipsBuildRun
		return ret
	}
	return *o.BuildRun
}

// GetBuildRunOk returns a tuple with the BuildRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestDataRelationships) GetBuildRunOk() (*CiBuildRunCreateRequestDataRelationshipsBuildRun, bool) {
	if o == nil || IsNil(o.BuildRun) {
		return nil, false
	}
	return o.BuildRun, true
}

// HasBuildRun returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestDataRelationships) HasBuildRun() bool {
	if o != nil && !IsNil(o.BuildRun) {
		return true
	}

	return false
}

// SetBuildRun gets a reference to the given CiBuildRunCreateRequestDataRelationshipsBuildRun and assigns it to the BuildRun field.
func (o *CiBuildRunCreateRequestDataRelationships) SetBuildRun(v CiBuildRunCreateRequestDataRelationshipsBuildRun) {
	o.BuildRun = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestDataRelationships) GetWorkflow() CiBuildRunRelationshipsWorkflow {
	if o == nil || IsNil(o.Workflow) {
		var ret CiBuildRunRelationshipsWorkflow
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestDataRelationships) GetWorkflowOk() (*CiBuildRunRelationshipsWorkflow, bool) {
	if o == nil || IsNil(o.Workflow) {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestDataRelationships) HasWorkflow() bool {
	if o != nil && !IsNil(o.Workflow) {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given CiBuildRunRelationshipsWorkflow and assigns it to the Workflow field.
func (o *CiBuildRunCreateRequestDataRelationships) SetWorkflow(v CiBuildRunRelationshipsWorkflow) {
	o.Workflow = &v
}

// GetSourceBranchOrTag returns the SourceBranchOrTag field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestDataRelationships) GetSourceBranchOrTag() CiBuildRunRelationshipsSourceBranchOrTag {
	if o == nil || IsNil(o.SourceBranchOrTag) {
		var ret CiBuildRunRelationshipsSourceBranchOrTag
		return ret
	}
	return *o.SourceBranchOrTag
}

// GetSourceBranchOrTagOk returns a tuple with the SourceBranchOrTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestDataRelationships) GetSourceBranchOrTagOk() (*CiBuildRunRelationshipsSourceBranchOrTag, bool) {
	if o == nil || IsNil(o.SourceBranchOrTag) {
		return nil, false
	}
	return o.SourceBranchOrTag, true
}

// HasSourceBranchOrTag returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestDataRelationships) HasSourceBranchOrTag() bool {
	if o != nil && !IsNil(o.SourceBranchOrTag) {
		return true
	}

	return false
}

// SetSourceBranchOrTag gets a reference to the given CiBuildRunRelationshipsSourceBranchOrTag and assigns it to the SourceBranchOrTag field.
func (o *CiBuildRunCreateRequestDataRelationships) SetSourceBranchOrTag(v CiBuildRunRelationshipsSourceBranchOrTag) {
	o.SourceBranchOrTag = &v
}

// GetPullRequest returns the PullRequest field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestDataRelationships) GetPullRequest() CiBuildRunRelationshipsPullRequest {
	if o == nil || IsNil(o.PullRequest) {
		var ret CiBuildRunRelationshipsPullRequest
		return ret
	}
	return *o.PullRequest
}

// GetPullRequestOk returns a tuple with the PullRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestDataRelationships) GetPullRequestOk() (*CiBuildRunRelationshipsPullRequest, bool) {
	if o == nil || IsNil(o.PullRequest) {
		return nil, false
	}
	return o.PullRequest, true
}

// HasPullRequest returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestDataRelationships) HasPullRequest() bool {
	if o != nil && !IsNil(o.PullRequest) {
		return true
	}

	return false
}

// SetPullRequest gets a reference to the given CiBuildRunRelationshipsPullRequest and assigns it to the PullRequest field.
func (o *CiBuildRunCreateRequestDataRelationships) SetPullRequest(v CiBuildRunRelationshipsPullRequest) {
	o.PullRequest = &v
}

func (o CiBuildRunCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildRun) {
		toSerialize["buildRun"] = o.BuildRun
	}
	if !IsNil(o.Workflow) {
		toSerialize["workflow"] = o.Workflow
	}
	if !IsNil(o.SourceBranchOrTag) {
		toSerialize["sourceBranchOrTag"] = o.SourceBranchOrTag
	}
	if !IsNil(o.PullRequest) {
		toSerialize["pullRequest"] = o.PullRequest
	}
	return toSerialize, nil
}

type NullableCiBuildRunCreateRequestDataRelationships struct {
	value *CiBuildRunCreateRequestDataRelationships
	isSet bool
}

func (v NullableCiBuildRunCreateRequestDataRelationships) Get() *CiBuildRunCreateRequestDataRelationships {
	return v.value
}

func (v *NullableCiBuildRunCreateRequestDataRelationships) Set(val *CiBuildRunCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunCreateRequestDataRelationships(val *CiBuildRunCreateRequestDataRelationships) *NullableCiBuildRunCreateRequestDataRelationships {
	return &NullableCiBuildRunCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableCiBuildRunCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

