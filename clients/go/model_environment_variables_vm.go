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

// EnvironmentVariablesVm struct for EnvironmentVariablesVm
type EnvironmentVariablesVm struct {
	EnvironmentVariables []EnvironmentVariableDto `json:"environmentVariables,omitempty"`
}

// NewEnvironmentVariablesVm instantiates a new EnvironmentVariablesVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariablesVm() *EnvironmentVariablesVm {
	this := EnvironmentVariablesVm{}
	return &this
}

// NewEnvironmentVariablesVmWithDefaults instantiates a new EnvironmentVariablesVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariablesVmWithDefaults() *EnvironmentVariablesVm {
	this := EnvironmentVariablesVm{}
	return &this
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentVariablesVm) GetEnvironmentVariables() []EnvironmentVariableDto {
	if o == nil  {
		var ret []EnvironmentVariableDto
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentVariablesVm) GetEnvironmentVariablesOk() (*[]EnvironmentVariableDto, bool) {
	if o == nil || o.EnvironmentVariables == nil {
		return nil, false
	}
	return &o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *EnvironmentVariablesVm) HasEnvironmentVariables() bool {
	if o != nil && o.EnvironmentVariables != nil {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []EnvironmentVariableDto and assigns it to the EnvironmentVariables field.
func (o *EnvironmentVariablesVm) SetEnvironmentVariables(v []EnvironmentVariableDto) {
	o.EnvironmentVariables = v
}

func (o EnvironmentVariablesVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnvironmentVariables != nil {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariablesVm struct {
	value *EnvironmentVariablesVm
	isSet bool
}

func (v NullableEnvironmentVariablesVm) Get() *EnvironmentVariablesVm {
	return v.value
}

func (v *NullableEnvironmentVariablesVm) Set(val *EnvironmentVariablesVm) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariablesVm) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariablesVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariablesVm(val *EnvironmentVariablesVm) *NullableEnvironmentVariablesVm {
	return &NullableEnvironmentVariablesVm{value: val, isSet: true}
}

func (v NullableEnvironmentVariablesVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariablesVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


