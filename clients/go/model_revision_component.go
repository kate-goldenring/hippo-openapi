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

// RevisionComponent struct for RevisionComponent
type RevisionComponent struct {
	Created *time.Time `json:"created,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	LastModified *time.Time `json:"lastModified,omitempty"`
	LastModifiedBy NullableString `json:"lastModifiedBy,omitempty"`
	Id *string `json:"id,omitempty"`
	Source NullableString `json:"source,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Route NullableString `json:"route,omitempty"`
	RevisionId *string `json:"revisionId,omitempty"`
	Revision *Revision `json:"revision,omitempty"`
	DomainEvents []DomainEvent `json:"domainEvents,omitempty"`
}

// NewRevisionComponent instantiates a new RevisionComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionComponent() *RevisionComponent {
	this := RevisionComponent{}
	return &this
}

// NewRevisionComponentWithDefaults instantiates a new RevisionComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionComponentWithDefaults() *RevisionComponent {
	this := RevisionComponent{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RevisionComponent) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionComponent) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RevisionComponent) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RevisionComponent) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionComponent) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionComponent) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RevisionComponent) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *RevisionComponent) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *RevisionComponent) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *RevisionComponent) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *RevisionComponent) GetLastModified() time.Time {
	if o == nil || o.LastModified == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionComponent) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || o.LastModified == nil {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *RevisionComponent) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *RevisionComponent) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionComponent) GetLastModifiedBy() string {
	if o == nil || o.LastModifiedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionComponent) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *RevisionComponent) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *RevisionComponent) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *RevisionComponent) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *RevisionComponent) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RevisionComponent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionComponent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RevisionComponent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RevisionComponent) SetId(v string) {
	o.Id = &v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionComponent) GetSource() string {
	if o == nil || o.Source.Get() == nil {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionComponent) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *RevisionComponent) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *RevisionComponent) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *RevisionComponent) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *RevisionComponent) UnsetSource() {
	o.Source.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionComponent) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RevisionComponent) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RevisionComponent) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RevisionComponent) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RevisionComponent) UnsetName() {
	o.Name.Unset()
}

// GetRoute returns the Route field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionComponent) GetRoute() string {
	if o == nil || o.Route.Get() == nil {
		var ret string
		return ret
	}
	return *o.Route.Get()
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionComponent) GetRouteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Route.Get(), o.Route.IsSet()
}

// HasRoute returns a boolean if a field has been set.
func (o *RevisionComponent) HasRoute() bool {
	if o != nil && o.Route.IsSet() {
		return true
	}

	return false
}

// SetRoute gets a reference to the given NullableString and assigns it to the Route field.
func (o *RevisionComponent) SetRoute(v string) {
	o.Route.Set(&v)
}
// SetRouteNil sets the value for Route to be an explicit nil
func (o *RevisionComponent) SetRouteNil() {
	o.Route.Set(nil)
}

// UnsetRoute ensures that no value is present for Route, not even an explicit nil
func (o *RevisionComponent) UnsetRoute() {
	o.Route.Unset()
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *RevisionComponent) GetRevisionId() string {
	if o == nil || o.RevisionId == nil {
		var ret string
		return ret
	}
	return *o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionComponent) GetRevisionIdOk() (*string, bool) {
	if o == nil || o.RevisionId == nil {
		return nil, false
	}
	return o.RevisionId, true
}

// HasRevisionId returns a boolean if a field has been set.
func (o *RevisionComponent) HasRevisionId() bool {
	if o != nil && o.RevisionId != nil {
		return true
	}

	return false
}

// SetRevisionId gets a reference to the given string and assigns it to the RevisionId field.
func (o *RevisionComponent) SetRevisionId(v string) {
	o.RevisionId = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *RevisionComponent) GetRevision() Revision {
	if o == nil || o.Revision == nil {
		var ret Revision
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionComponent) GetRevisionOk() (*Revision, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *RevisionComponent) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given Revision and assigns it to the Revision field.
func (o *RevisionComponent) SetRevision(v Revision) {
	o.Revision = &v
}

// GetDomainEvents returns the DomainEvents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RevisionComponent) GetDomainEvents() []DomainEvent {
	if o == nil {
		var ret []DomainEvent
		return ret
	}
	return o.DomainEvents
}

// GetDomainEventsOk returns a tuple with the DomainEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RevisionComponent) GetDomainEventsOk() ([]DomainEvent, bool) {
	if o == nil || o.DomainEvents == nil {
		return nil, false
	}
	return o.DomainEvents, true
}

// HasDomainEvents returns a boolean if a field has been set.
func (o *RevisionComponent) HasDomainEvents() bool {
	if o != nil && o.DomainEvents != nil {
		return true
	}

	return false
}

// SetDomainEvents gets a reference to the given []DomainEvent and assigns it to the DomainEvents field.
func (o *RevisionComponent) SetDomainEvents(v []DomainEvent) {
	o.DomainEvents = v
}

func (o RevisionComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.LastModified != nil {
		toSerialize["lastModified"] = o.LastModified
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Route.IsSet() {
		toSerialize["route"] = o.Route.Get()
	}
	if o.RevisionId != nil {
		toSerialize["revisionId"] = o.RevisionId
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.DomainEvents != nil {
		toSerialize["domainEvents"] = o.DomainEvents
	}
	return json.Marshal(toSerialize)
}

type NullableRevisionComponent struct {
	value *RevisionComponent
	isSet bool
}

func (v NullableRevisionComponent) Get() *RevisionComponent {
	return v.value
}

func (v *NullableRevisionComponent) Set(val *RevisionComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionComponent(val *RevisionComponent) *NullableRevisionComponent {
	return &NullableRevisionComponent{value: val, isSet: true}
}

func (v NullableRevisionComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

