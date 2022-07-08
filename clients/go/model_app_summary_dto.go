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

// AppSummaryDto struct for AppSummaryDto
type AppSummaryDto struct {
	Id string `json:"id"`
	Name string `json:"name"`
	StorageId string `json:"storageId"`
	Channels []AppChannelListItem `json:"channels"`
}

// NewAppSummaryDto instantiates a new AppSummaryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSummaryDto(id string, name string, storageId string, channels []AppChannelListItem) *AppSummaryDto {
	this := AppSummaryDto{}
	this.Id = id
	this.Name = name
	this.StorageId = storageId
	this.Channels = channels
	return &this
}

// NewAppSummaryDtoWithDefaults instantiates a new AppSummaryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSummaryDtoWithDefaults() *AppSummaryDto {
	this := AppSummaryDto{}
	return &this
}

// GetId returns the Id field value
func (o *AppSummaryDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppSummaryDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppSummaryDto) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AppSummaryDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppSummaryDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppSummaryDto) SetName(v string) {
	o.Name = v
}

// GetStorageId returns the StorageId field value
func (o *AppSummaryDto) GetStorageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *AppSummaryDto) GetStorageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *AppSummaryDto) SetStorageId(v string) {
	o.StorageId = v
}

// GetChannels returns the Channels field value
func (o *AppSummaryDto) GetChannels() []AppChannelListItem {
	if o == nil {
		var ret []AppChannelListItem
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *AppSummaryDto) GetChannelsOk() ([]AppChannelListItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *AppSummaryDto) SetChannels(v []AppChannelListItem) {
	o.Channels = v
}

func (o AppSummaryDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["storageId"] = o.StorageId
	}
	if true {
		toSerialize["channels"] = o.Channels
	}
	return json.Marshal(toSerialize)
}

type NullableAppSummaryDto struct {
	value *AppSummaryDto
	isSet bool
}

func (v NullableAppSummaryDto) Get() *AppSummaryDto {
	return v.value
}

func (v *NullableAppSummaryDto) Set(val *AppSummaryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSummaryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSummaryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSummaryDto(val *AppSummaryDto) *NullableAppSummaryDto {
	return &NullableAppSummaryDto{value: val, isSet: true}
}

func (v NullableAppSummaryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSummaryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


