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

// ChannelsVm struct for ChannelsVm
type ChannelsVm struct {
	Channels []ChannelDto `json:"channels"`
}

// NewChannelsVm instantiates a new ChannelsVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelsVm(channels []ChannelDto) *ChannelsVm {
	this := ChannelsVm{}
	this.Channels = channels
	return &this
}

// NewChannelsVmWithDefaults instantiates a new ChannelsVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelsVmWithDefaults() *ChannelsVm {
	this := ChannelsVm{}
	return &this
}

// GetChannels returns the Channels field value
func (o *ChannelsVm) GetChannels() []ChannelDto {
	if o == nil {
		var ret []ChannelDto
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *ChannelsVm) GetChannelsOk() (*[]ChannelDto, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Channels, true
}

// SetChannels sets field value
func (o *ChannelsVm) SetChannels(v []ChannelDto) {
	o.Channels = v
}

func (o ChannelsVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channels"] = o.Channels
	}
	return json.Marshal(toSerialize)
}

type NullableChannelsVm struct {
	value *ChannelsVm
	isSet bool
}

func (v NullableChannelsVm) Get() *ChannelsVm {
	return v.value
}

func (v *NullableChannelsVm) Set(val *ChannelsVm) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelsVm) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelsVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelsVm(val *ChannelsVm) *NullableChannelsVm {
	return &NullableChannelsVm{value: val, isSet: true}
}

func (v NullableChannelsVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelsVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


