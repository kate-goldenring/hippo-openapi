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
	Id string `json:"id"`
	AppId string `json:"appId"`
	Name string `json:"name"`
	Domain string `json:"domain"`
	RevisionSelectionStrategy ChannelRevisionSelectionStrategy `json:"revisionSelectionStrategy"`
	ActiveRevision *RevisionDto `json:"activeRevision,omitempty"`
	RangeRule NullableString `json:"rangeRule,omitempty"`
	Certificate *CertificateDto `json:"certificate,omitempty"`
	EnvironmentVariables []EnvironmentVariableDto `json:"environmentVariables"`
	AppSummary *AppSummaryDto `json:"appSummary,omitempty"`
}

// NewChannelDto instantiates a new ChannelDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelDto(id string, appId string, name string, domain string, revisionSelectionStrategy ChannelRevisionSelectionStrategy, environmentVariables []EnvironmentVariableDto) *ChannelDto {
	this := ChannelDto{}
	this.Id = id
	this.AppId = appId
	this.Name = name
	this.Domain = domain
	this.RevisionSelectionStrategy = revisionSelectionStrategy
	this.EnvironmentVariables = environmentVariables
	return &this
}

// NewChannelDtoWithDefaults instantiates a new ChannelDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelDtoWithDefaults() *ChannelDto {
	this := ChannelDto{}
	return &this
}

// GetId returns the Id field value
func (o *ChannelDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelDto) SetId(v string) {
	o.Id = v
}

// GetAppId returns the AppId field value
func (o *ChannelDto) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *ChannelDto) SetAppId(v string) {
	o.AppId = v
}

// GetName returns the Name field value
func (o *ChannelDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChannelDto) SetName(v string) {
	o.Name = v
}

// GetDomain returns the Domain field value
func (o *ChannelDto) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ChannelDto) SetDomain(v string) {
	o.Domain = v
}

// GetRevisionSelectionStrategy returns the RevisionSelectionStrategy field value
func (o *ChannelDto) GetRevisionSelectionStrategy() ChannelRevisionSelectionStrategy {
	if o == nil {
		var ret ChannelRevisionSelectionStrategy
		return ret
	}

	return o.RevisionSelectionStrategy
}

// GetRevisionSelectionStrategyOk returns a tuple with the RevisionSelectionStrategy field value
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetRevisionSelectionStrategyOk() (*ChannelRevisionSelectionStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionSelectionStrategy, true
}

// SetRevisionSelectionStrategy sets field value
func (o *ChannelDto) SetRevisionSelectionStrategy(v ChannelRevisionSelectionStrategy) {
	o.RevisionSelectionStrategy = v
}

// GetActiveRevision returns the ActiveRevision field value if set, zero value otherwise.
func (o *ChannelDto) GetActiveRevision() RevisionDto {
	if o == nil || o.ActiveRevision == nil {
		var ret RevisionDto
		return ret
	}
	return *o.ActiveRevision
}

// GetActiveRevisionOk returns a tuple with the ActiveRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetActiveRevisionOk() (*RevisionDto, bool) {
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

// SetActiveRevision gets a reference to the given RevisionDto and assigns it to the ActiveRevision field.
func (o *ChannelDto) SetActiveRevision(v RevisionDto) {
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
	if o == nil {
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

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *ChannelDto) GetCertificate() CertificateDto {
	if o == nil || o.Certificate == nil {
		var ret CertificateDto
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetCertificateOk() (*CertificateDto, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *ChannelDto) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given CertificateDto and assigns it to the Certificate field.
func (o *ChannelDto) SetCertificate(v CertificateDto) {
	o.Certificate = &v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value
func (o *ChannelDto) GetEnvironmentVariables() []EnvironmentVariableDto {
	if o == nil {
		var ret []EnvironmentVariableDto
		return ret
	}

	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetEnvironmentVariablesOk() ([]EnvironmentVariableDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// SetEnvironmentVariables sets field value
func (o *ChannelDto) SetEnvironmentVariables(v []EnvironmentVariableDto) {
	o.EnvironmentVariables = v
}

// GetAppSummary returns the AppSummary field value if set, zero value otherwise.
func (o *ChannelDto) GetAppSummary() AppSummaryDto {
	if o == nil || o.AppSummary == nil {
		var ret AppSummaryDto
		return ret
	}
	return *o.AppSummary
}

// GetAppSummaryOk returns a tuple with the AppSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDto) GetAppSummaryOk() (*AppSummaryDto, bool) {
	if o == nil || o.AppSummary == nil {
		return nil, false
	}
	return o.AppSummary, true
}

// HasAppSummary returns a boolean if a field has been set.
func (o *ChannelDto) HasAppSummary() bool {
	if o != nil && o.AppSummary != nil {
		return true
	}

	return false
}

// SetAppSummary gets a reference to the given AppSummaryDto and assigns it to the AppSummary field.
func (o *ChannelDto) SetAppSummary(v AppSummaryDto) {
	o.AppSummary = &v
}

func (o ChannelDto) MarshalJSON() ([]byte, error) {
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


