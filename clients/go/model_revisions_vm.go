/*
Hippo.Web

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hippo-openapi

import (
	"encoding/json"
)

// RevisionsVm struct for RevisionsVm
type RevisionsVm struct {
	Revisions []RevisionDto `json:"revisions,omitempty"`
}

// NewRevisionsVm instantiates a new RevisionsVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionsVm() *RevisionsVm {
	this := RevisionsVm{}
	return &this
}

// NewRevisionsVmWithDefaults instantiates a new RevisionsVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionsVmWithDefaults() *RevisionsVm {
	this := RevisionsVm{}
	return &this
}

// GetRevisions returns the Revisions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionsVm) GetRevisions() []RevisionDto {
	if o == nil  {
		var ret []RevisionDto
		return ret
	}
	return o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionsVm) GetRevisionsOk() (*[]RevisionDto, bool) {
	if o == nil || o.Revisions == nil {
		return nil, false
	}
	return &o.Revisions, true
}

// HasRevisions returns a boolean if a field has been set.
func (o *RevisionsVm) HasRevisions() bool {
	if o != nil && o.Revisions != nil {
		return true
	}

	return false
}

// SetRevisions gets a reference to the given []RevisionDto and assigns it to the Revisions field.
func (o *RevisionsVm) SetRevisions(v []RevisionDto) {
	o.Revisions = v
}

func (o RevisionsVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Revisions != nil {
		toSerialize["revisions"] = o.Revisions
	}
	return json.Marshal(toSerialize)
}

type NullableRevisionsVm struct {
	value *RevisionsVm
	isSet bool
}

func (v NullableRevisionsVm) Get() *RevisionsVm {
	return v.value
}

func (v *NullableRevisionsVm) Set(val *RevisionsVm) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionsVm) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionsVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionsVm(val *RevisionsVm) *NullableRevisionsVm {
	return &NullableRevisionsVm{value: val, isSet: true}
}

func (v NullableRevisionsVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionsVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


