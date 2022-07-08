/*
Hippo.Web

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hippo-openapi

import (
	"encoding/json"
	"time"
)

// ChannelItem struct for ChannelItem
type ChannelItem struct {
	Id string `json:"id"`
	AppId string `json:"appId"`
	Name string `json:"name"`
	Domain string `json:"domain"`
	RevisionSelectionStrategy ChannelRevisionSelectionStrategy `json:"revisionSelectionStrategy"`
	ActiveRevision *RevisionItem `json:"activeRevision,omitempty"`
	LastPublishAt NullableTime `json:"lastPublishAt,omitempty"`
	RangeRule NullableString `json:"rangeRule,omitempty"`
	Certificate *CertificateItem `json:"certificate,omitempty"`
	EnvironmentVariables []EnvironmentVariableItem `json:"environmentVariables"`
	AppSummary *AppSummaryDto `json:"appSummary,omitempty"`
}

// NewChannelItem instantiates a new ChannelItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelItem(id string, appId string, name string, domain string, revisionSelectionStrategy ChannelRevisionSelectionStrategy, environmentVariables []EnvironmentVariableItem) *ChannelItem {
	this := ChannelItem{}
	this.Id = id
	this.AppId = appId
	this.Name = name
	this.Domain = domain
	this.RevisionSelectionStrategy = revisionSelectionStrategy
	this.EnvironmentVariables = environmentVariables
	return &this
}

// NewChannelItemWithDefaults instantiates a new ChannelItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelItemWithDefaults() *ChannelItem {
	this := ChannelItem{}
	return &this
}

// GetId returns the Id field value
func (o *ChannelItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelItem) SetId(v string) {
	o.Id = v
}

// GetAppId returns the AppId field value
func (o *ChannelItem) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *ChannelItem) SetAppId(v string) {
	o.AppId = v
}

// GetName returns the Name field value
func (o *ChannelItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChannelItem) SetName(v string) {
	o.Name = v
}

// GetDomain returns the Domain field value
func (o *ChannelItem) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ChannelItem) SetDomain(v string) {
	o.Domain = v
}

// GetRevisionSelectionStrategy returns the RevisionSelectionStrategy field value
func (o *ChannelItem) GetRevisionSelectionStrategy() ChannelRevisionSelectionStrategy {
	if o == nil {
		var ret ChannelRevisionSelectionStrategy
		return ret
	}

	return o.RevisionSelectionStrategy
}

// GetRevisionSelectionStrategyOk returns a tuple with the RevisionSelectionStrategy field value
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetRevisionSelectionStrategyOk() (*ChannelRevisionSelectionStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionSelectionStrategy, true
}

// SetRevisionSelectionStrategy sets field value
func (o *ChannelItem) SetRevisionSelectionStrategy(v ChannelRevisionSelectionStrategy) {
	o.RevisionSelectionStrategy = v
}

// GetActiveRevision returns the ActiveRevision field value if set, zero value otherwise.
func (o *ChannelItem) GetActiveRevision() RevisionItem {
	if o == nil || o.ActiveRevision == nil {
		var ret RevisionItem
		return ret
	}
	return *o.ActiveRevision
}

// GetActiveRevisionOk returns a tuple with the ActiveRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetActiveRevisionOk() (*RevisionItem, bool) {
	if o == nil || o.ActiveRevision == nil {
		return nil, false
	}
	return o.ActiveRevision, true
}

// HasActiveRevision returns a boolean if a field has been set.
func (o *ChannelItem) HasActiveRevision() bool {
	if o != nil && o.ActiveRevision != nil {
		return true
	}

	return false
}

// SetActiveRevision gets a reference to the given RevisionItem and assigns it to the ActiveRevision field.
func (o *ChannelItem) SetActiveRevision(v RevisionItem) {
	o.ActiveRevision = &v
}

// GetLastPublishAt returns the LastPublishAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelItem) GetLastPublishAt() time.Time {
	if o == nil || o.LastPublishAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastPublishAt.Get()
}

// GetLastPublishAtOk returns a tuple with the LastPublishAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelItem) GetLastPublishAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPublishAt.Get(), o.LastPublishAt.IsSet()
}

// HasLastPublishAt returns a boolean if a field has been set.
func (o *ChannelItem) HasLastPublishAt() bool {
	if o != nil && o.LastPublishAt.IsSet() {
		return true
	}

	return false
}

// SetLastPublishAt gets a reference to the given NullableTime and assigns it to the LastPublishAt field.
func (o *ChannelItem) SetLastPublishAt(v time.Time) {
	o.LastPublishAt.Set(&v)
}
// SetLastPublishAtNil sets the value for LastPublishAt to be an explicit nil
func (o *ChannelItem) SetLastPublishAtNil() {
	o.LastPublishAt.Set(nil)
}

// UnsetLastPublishAt ensures that no value is present for LastPublishAt, not even an explicit nil
func (o *ChannelItem) UnsetLastPublishAt() {
	o.LastPublishAt.Unset()
}

// GetRangeRule returns the RangeRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelItem) GetRangeRule() string {
	if o == nil || o.RangeRule.Get() == nil {
		var ret string
		return ret
	}
	return *o.RangeRule.Get()
}

// GetRangeRuleOk returns a tuple with the RangeRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelItem) GetRangeRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RangeRule.Get(), o.RangeRule.IsSet()
}

// HasRangeRule returns a boolean if a field has been set.
func (o *ChannelItem) HasRangeRule() bool {
	if o != nil && o.RangeRule.IsSet() {
		return true
	}

	return false
}

// SetRangeRule gets a reference to the given NullableString and assigns it to the RangeRule field.
func (o *ChannelItem) SetRangeRule(v string) {
	o.RangeRule.Set(&v)
}
// SetRangeRuleNil sets the value for RangeRule to be an explicit nil
func (o *ChannelItem) SetRangeRuleNil() {
	o.RangeRule.Set(nil)
}

// UnsetRangeRule ensures that no value is present for RangeRule, not even an explicit nil
func (o *ChannelItem) UnsetRangeRule() {
	o.RangeRule.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *ChannelItem) GetCertificate() CertificateItem {
	if o == nil || o.Certificate == nil {
		var ret CertificateItem
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetCertificateOk() (*CertificateItem, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *ChannelItem) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given CertificateItem and assigns it to the Certificate field.
func (o *ChannelItem) SetCertificate(v CertificateItem) {
	o.Certificate = &v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value
func (o *ChannelItem) GetEnvironmentVariables() []EnvironmentVariableItem {
	if o == nil {
		var ret []EnvironmentVariableItem
		return ret
	}

	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetEnvironmentVariablesOk() ([]EnvironmentVariableItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// SetEnvironmentVariables sets field value
func (o *ChannelItem) SetEnvironmentVariables(v []EnvironmentVariableItem) {
	o.EnvironmentVariables = v
}

// GetAppSummary returns the AppSummary field value if set, zero value otherwise.
func (o *ChannelItem) GetAppSummary() AppSummaryDto {
	if o == nil || o.AppSummary == nil {
		var ret AppSummaryDto
		return ret
	}
	return *o.AppSummary
}

// GetAppSummaryOk returns a tuple with the AppSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelItem) GetAppSummaryOk() (*AppSummaryDto, bool) {
	if o == nil || o.AppSummary == nil {
		return nil, false
	}
	return o.AppSummary, true
}

// HasAppSummary returns a boolean if a field has been set.
func (o *ChannelItem) HasAppSummary() bool {
	if o != nil && o.AppSummary != nil {
		return true
	}

	return false
}

// SetAppSummary gets a reference to the given AppSummaryDto and assigns it to the AppSummary field.
func (o *ChannelItem) SetAppSummary(v AppSummaryDto) {
	o.AppSummary = &v
}

func (o ChannelItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["revisionSelectionStrategy"] = o.RevisionSelectionStrategy
	}
	if o.ActiveRevision != nil {
		toSerialize["activeRevision"] = o.ActiveRevision
	}
	if o.LastPublishAt.IsSet() {
		toSerialize["lastPublishAt"] = o.LastPublishAt.Get()
	}
	if o.RangeRule.IsSet() {
		toSerialize["rangeRule"] = o.RangeRule.Get()
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if true {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if o.AppSummary != nil {
		toSerialize["appSummary"] = o.AppSummary
	}
	return json.Marshal(toSerialize)
}

type NullableChannelItem struct {
	value *ChannelItem
	isSet bool
}

func (v NullableChannelItem) Get() *ChannelItem {
	return v.value
}

func (v *NullableChannelItem) Set(val *ChannelItem) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelItem) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelItem(val *ChannelItem) *NullableChannelItem {
	return &NullableChannelItem{value: val, isSet: true}
}

func (v NullableChannelItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


