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

// ChannelDto struct for ChannelDto
type ChannelDto struct {
	Id *string `json:"id,omitempty"`
	AppId *string `json:"appId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Domain NullableString `json:"domain,omitempty"`
	RevisionSelectionStrategy *ChannelRevisionSelectionStrategy `json:"revisionSelectionStrategy,omitempty"`
	ActiveRevision *Revision `json:"activeRevision,omitempty"`
	RangeRule NullableString `json:"rangeRule,omitempty"`
	EnvironmentVariables []EnvironmentVariableDto `json:"environmentVariables,omitempty"`
}

// NewChannelDto instantiates a new ChannelDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelDto() *ChannelDto {
	this := ChannelDto{}
	return &this
}

// NewChannelDtoWithDefaults instantiates a new ChannelDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelDtoWithDefaults() *ChannelDto {
	this := ChannelDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelDto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelDto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelDto) SetId(v string) {
	o.Id = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ChannelDto) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ChannelDto) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ChannelDto) SetAppId(v string) {
	o.AppId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelDto) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ChannelDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ChannelDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ChannelDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ChannelDto) UnsetName() {
	o.Name.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelDto) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelDto) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *ChannelDto) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *ChannelDto) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *ChannelDto) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *ChannelDto) UnsetDomain() {
	o.Domain.Unset()
}

// GetRevisionSelectionStrategy returns the RevisionSelectionStrategy field value if set, zero value otherwise.
func (o *ChannelDto) GetRevisionSelectionStrategy() ChannelRevisionSelectionStrategy {
	if o == nil || o.RevisionSelectionStrategy == nil {
		var ret ChannelRevisionSelectionStrategy
		return ret
	}
	return *o.RevisionSelectionStrategy
}

// GetRevisionSelectionStrategyOk returns a tuple with the RevisionSelectionStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetRevisionSelectionStrategyOk() (*ChannelRevisionSelectionStrategy, bool) {
	if o == nil || o.RevisionSelectionStrategy == nil {
		return nil, false
	}
	return o.RevisionSelectionStrategy, true
}

// HasRevisionSelectionStrategy returns a boolean if a field has been set.
func (o *ChannelDto) HasRevisionSelectionStrategy() bool {
	if o != nil && o.RevisionSelectionStrategy != nil {
		return true
	}

	return false
}

// SetRevisionSelectionStrategy gets a reference to the given ChannelRevisionSelectionStrategy and assigns it to the RevisionSelectionStrategy field.
func (o *ChannelDto) SetRevisionSelectionStrategy(v ChannelRevisionSelectionStrategy) {
	o.RevisionSelectionStrategy = &v
}

// GetActiveRevision returns the ActiveRevision field value if set, zero value otherwise.
func (o *ChannelDto) GetActiveRevision() Revision {
	if o == nil || o.ActiveRevision == nil {
		var ret Revision
		return ret
	}
	return *o.ActiveRevision
}

// GetActiveRevisionOk returns a tuple with the ActiveRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetActiveRevisionOk() (*Revision, bool) {
	if o == nil || o.ActiveRevision == nil {
		return nil, false
	}
	return o.ActiveRevision, true
}

// HasActiveRevision returns a boolean if a field has been set.
func (o *ChannelDto) HasActiveRevision() bool {
	if o != nil && o.ActiveRevision != nil {
		return true
	}

	return false
}

// SetActiveRevision gets a reference to the given Revision and assigns it to the ActiveRevision field.
func (o *ChannelDto) SetActiveRevision(v Revision) {
	o.ActiveRevision = &v
}

// GetRangeRule returns the RangeRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelDto) GetRangeRule() string {
	if o == nil || o.RangeRule.Get() == nil {
		var ret string
		return ret
	}
	return *o.RangeRule.Get()
}

// GetRangeRuleOk returns a tuple with the RangeRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelDto) GetRangeRuleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RangeRule.Get(), o.RangeRule.IsSet()
}

// HasRangeRule returns a boolean if a field has been set.
func (o *ChannelDto) HasRangeRule() bool {
	if o != nil && o.RangeRule.IsSet() {
		return true
	}

	return false
}

// SetRangeRule gets a reference to the given NullableString and assigns it to the RangeRule field.
func (o *ChannelDto) SetRangeRule(v string) {
	o.RangeRule.Set(&v)
}
// SetRangeRuleNil sets the value for RangeRule to be an explicit nil
func (o *ChannelDto) SetRangeRuleNil() {
	o.RangeRule.Set(nil)
}

// UnsetRangeRule ensures that no value is present for RangeRule, not even an explicit nil
func (o *ChannelDto) UnsetRangeRule() {
	o.RangeRule.Unset()
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelDto) GetEnvironmentVariables() []EnvironmentVariableDto {
	if o == nil  {
		var ret []EnvironmentVariableDto
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelDto) GetEnvironmentVariablesOk() (*[]EnvironmentVariableDto, bool) {
	if o == nil || o.EnvironmentVariables == nil {
		return nil, false
	}
	return &o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *ChannelDto) HasEnvironmentVariables() bool {
	if o != nil && o.EnvironmentVariables != nil {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []EnvironmentVariableDto and assigns it to the EnvironmentVariables field.
func (o *ChannelDto) SetEnvironmentVariables(v []EnvironmentVariableDto) {
	o.EnvironmentVariables = v
}

func (o ChannelDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.RevisionSelectionStrategy != nil {
		toSerialize["revisionSelectionStrategy"] = o.RevisionSelectionStrategy
	}
	if o.ActiveRevision != nil {
		toSerialize["activeRevision"] = o.ActiveRevision
	}
	if o.RangeRule.IsSet() {
		toSerialize["rangeRule"] = o.RangeRule.Get()
	}
	if o.EnvironmentVariables != nil {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	return json.Marshal(toSerialize)
}

type NullableChannelDto struct {
	value *ChannelDto
	isSet bool
}

func (v NullableChannelDto) Get() *ChannelDto {
	return v.value
}

func (v *NullableChannelDto) Set(val *ChannelDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelDto(val *ChannelDto) *NullableChannelDto {
	return &NullableChannelDto{value: val, isSet: true}
}

func (v NullableChannelDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

