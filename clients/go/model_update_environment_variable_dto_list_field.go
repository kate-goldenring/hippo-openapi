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

// UpdateEnvironmentVariableDtoListField struct for UpdateEnvironmentVariableDtoListField
type UpdateEnvironmentVariableDtoListField struct {
	Value []UpdateEnvironmentVariableDto `json:"value,omitempty"`
}

// NewUpdateEnvironmentVariableDtoListField instantiates a new UpdateEnvironmentVariableDtoListField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEnvironmentVariableDtoListField() *UpdateEnvironmentVariableDtoListField {
	this := UpdateEnvironmentVariableDtoListField{}
	return &this
}

// NewUpdateEnvironmentVariableDtoListFieldWithDefaults instantiates a new UpdateEnvironmentVariableDtoListField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEnvironmentVariableDtoListFieldWithDefaults() *UpdateEnvironmentVariableDtoListField {
	this := UpdateEnvironmentVariableDtoListField{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateEnvironmentVariableDtoListField) GetValue() []UpdateEnvironmentVariableDto {
	if o == nil {
		var ret []UpdateEnvironmentVariableDto
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateEnvironmentVariableDtoListField) GetValueOk() ([]UpdateEnvironmentVariableDto, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UpdateEnvironmentVariableDtoListField) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []UpdateEnvironmentVariableDto and assigns it to the Value field.
func (o *UpdateEnvironmentVariableDtoListField) SetValue(v []UpdateEnvironmentVariableDto) {
	o.Value = v
}

func (o UpdateEnvironmentVariableDtoListField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateEnvironmentVariableDtoListField struct {
	value *UpdateEnvironmentVariableDtoListField
	isSet bool
}

func (v NullableUpdateEnvironmentVariableDtoListField) Get() *UpdateEnvironmentVariableDtoListField {
	return v.value
}

func (v *NullableUpdateEnvironmentVariableDtoListField) Set(val *UpdateEnvironmentVariableDtoListField) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEnvironmentVariableDtoListField) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEnvironmentVariableDtoListField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEnvironmentVariableDtoListField(val *UpdateEnvironmentVariableDtoListField) *NullableUpdateEnvironmentVariableDtoListField {
	return &NullableUpdateEnvironmentVariableDtoListField{value: val, isSet: true}
}

func (v NullableUpdateEnvironmentVariableDtoListField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEnvironmentVariableDtoListField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


