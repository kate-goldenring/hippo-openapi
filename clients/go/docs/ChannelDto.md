# ChannelDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**RevisionSelectionStrategy** | Pointer to [**ChannelRevisionSelectionStrategy**](ChannelRevisionSelectionStrategy.md) |  | [optional] 
**ActiveRevision** | Pointer to [**Revision**](Revision.md) |  | [optional] 
**RangeRule** | Pointer to **NullableString** |  | [optional] 
**EnvironmentVariables** | Pointer to [**[]EnvironmentVariableDto**](EnvironmentVariableDto.md) |  | [optional] 

## Methods

### NewChannelDto

`func NewChannelDto() *ChannelDto`

NewChannelDto instantiates a new ChannelDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelDtoWithDefaults

`func NewChannelDtoWithDefaults() *ChannelDto`

NewChannelDtoWithDefaults instantiates a new ChannelDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppId

`func (o *ChannelDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ChannelDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ChannelDto) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ChannelDto) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetName

`func (o *ChannelDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ChannelDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ChannelDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDomain

`func (o *ChannelDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ChannelDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ChannelDto) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ChannelDto) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *ChannelDto) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *ChannelDto) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetRevisionSelectionStrategy

`func (o *ChannelDto) GetRevisionSelectionStrategy() ChannelRevisionSelectionStrategy`

GetRevisionSelectionStrategy returns the RevisionSelectionStrategy field if non-nil, zero value otherwise.

### GetRevisionSelectionStrategyOk

`func (o *ChannelDto) GetRevisionSelectionStrategyOk() (*ChannelRevisionSelectionStrategy, bool)`

GetRevisionSelectionStrategyOk returns a tuple with the RevisionSelectionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionSelectionStrategy

`func (o *ChannelDto) SetRevisionSelectionStrategy(v ChannelRevisionSelectionStrategy)`

SetRevisionSelectionStrategy sets RevisionSelectionStrategy field to given value.

### HasRevisionSelectionStrategy

`func (o *ChannelDto) HasRevisionSelectionStrategy() bool`

HasRevisionSelectionStrategy returns a boolean if a field has been set.

### GetActiveRevision

`func (o *ChannelDto) GetActiveRevision() Revision`

GetActiveRevision returns the ActiveRevision field if non-nil, zero value otherwise.

### GetActiveRevisionOk

`func (o *ChannelDto) GetActiveRevisionOk() (*Revision, bool)`

GetActiveRevisionOk returns a tuple with the ActiveRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRevision

`func (o *ChannelDto) SetActiveRevision(v Revision)`

SetActiveRevision sets ActiveRevision field to given value.

### HasActiveRevision

`func (o *ChannelDto) HasActiveRevision() bool`

HasActiveRevision returns a boolean if a field has been set.

### GetRangeRule

`func (o *ChannelDto) GetRangeRule() string`

GetRangeRule returns the RangeRule field if non-nil, zero value otherwise.

### GetRangeRuleOk

`func (o *ChannelDto) GetRangeRuleOk() (*string, bool)`

GetRangeRuleOk returns a tuple with the RangeRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeRule

`func (o *ChannelDto) SetRangeRule(v string)`

SetRangeRule sets RangeRule field to given value.

### HasRangeRule

`func (o *ChannelDto) HasRangeRule() bool`

HasRangeRule returns a boolean if a field has been set.

### SetRangeRuleNil

`func (o *ChannelDto) SetRangeRuleNil(b bool)`

 SetRangeRuleNil sets the value for RangeRule to be an explicit nil

### UnsetRangeRule
`func (o *ChannelDto) UnsetRangeRule()`

UnsetRangeRule ensures that no value is present for RangeRule, not even an explicit nil
### GetEnvironmentVariables

`func (o *ChannelDto) GetEnvironmentVariables() []EnvironmentVariableDto`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ChannelDto) GetEnvironmentVariablesOk() (*[]EnvironmentVariableDto, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ChannelDto) SetEnvironmentVariables(v []EnvironmentVariableDto)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ChannelDto) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *ChannelDto) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *ChannelDto) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

