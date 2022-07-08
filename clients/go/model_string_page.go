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

// StringPage struct for StringPage
type StringPage struct {
	Items []string `json:"items,omitempty"`
	TotalItems *int32 `json:"totalItems,omitempty"`
	PageIndex NullableInt32 `json:"pageIndex,omitempty"`
	PageSize NullableInt32 `json:"pageSize,omitempty"`
	IsLastPage NullableBool `json:"isLastPage,omitempty"`
}

// NewStringPage instantiates a new StringPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringPage() *StringPage {
	this := StringPage{}
	return &this
}

// NewStringPageWithDefaults instantiates a new StringPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringPageWithDefaults() *StringPage {
	this := StringPage{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringPage) GetItems() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringPage) GetItemsOk() ([]string, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *StringPage) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *StringPage) SetItems(v []string) {
	o.Items = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *StringPage) GetTotalItems() int32 {
	if o == nil || o.TotalItems == nil {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringPage) GetTotalItemsOk() (*int32, bool) {
	if o == nil || o.TotalItems == nil {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *StringPage) HasTotalItems() bool {
	if o != nil && o.TotalItems != nil {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *StringPage) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringPage) GetPageIndex() int32 {
	if o == nil || o.PageIndex.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PageIndex.Get()
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringPage) GetPageIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageIndex.Get(), o.PageIndex.IsSet()
}

// HasPageIndex returns a boolean if a field has been set.
func (o *StringPage) HasPageIndex() bool {
	if o != nil && o.PageIndex.IsSet() {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given NullableInt32 and assigns it to the PageIndex field.
func (o *StringPage) SetPageIndex(v int32) {
	o.PageIndex.Set(&v)
}
// SetPageIndexNil sets the value for PageIndex to be an explicit nil
func (o *StringPage) SetPageIndexNil() {
	o.PageIndex.Set(nil)
}

// UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
func (o *StringPage) UnsetPageIndex() {
	o.PageIndex.Unset()
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringPage) GetPageSize() int32 {
	if o == nil || o.PageSize.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringPage) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *StringPage) HasPageSize() bool {
	if o != nil && o.PageSize.IsSet() {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given NullableInt32 and assigns it to the PageSize field.
func (o *StringPage) SetPageSize(v int32) {
	o.PageSize.Set(&v)
}
// SetPageSizeNil sets the value for PageSize to be an explicit nil
func (o *StringPage) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
func (o *StringPage) UnsetPageSize() {
	o.PageSize.Unset()
}

// GetIsLastPage returns the IsLastPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringPage) GetIsLastPage() bool {
	if o == nil || o.IsLastPage.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsLastPage.Get()
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringPage) GetIsLastPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsLastPage.Get(), o.IsLastPage.IsSet()
}

// HasIsLastPage returns a boolean if a field has been set.
func (o *StringPage) HasIsLastPage() bool {
	if o != nil && o.IsLastPage.IsSet() {
		return true
	}

	return false
}

// SetIsLastPage gets a reference to the given NullableBool and assigns it to the IsLastPage field.
func (o *StringPage) SetIsLastPage(v bool) {
	o.IsLastPage.Set(&v)
}
// SetIsLastPageNil sets the value for IsLastPage to be an explicit nil
func (o *StringPage) SetIsLastPageNil() {
	o.IsLastPage.Set(nil)
}

// UnsetIsLastPage ensures that no value is present for IsLastPage, not even an explicit nil
func (o *StringPage) UnsetIsLastPage() {
	o.IsLastPage.Unset()
}

func (o StringPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.TotalItems != nil {
		toSerialize["totalItems"] = o.TotalItems
	}
	if o.PageIndex.IsSet() {
		toSerialize["pageIndex"] = o.PageIndex.Get()
	}
	if o.PageSize.IsSet() {
		toSerialize["pageSize"] = o.PageSize.Get()
	}
	if o.IsLastPage.IsSet() {
		toSerialize["isLastPage"] = o.IsLastPage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableStringPage struct {
	value *StringPage
	isSet bool
}

func (v NullableStringPage) Get() *StringPage {
	return v.value
}

func (v *NullableStringPage) Set(val *StringPage) {
	v.value = val
	v.isSet = true
}

func (v NullableStringPage) IsSet() bool {
	return v.isSet
}

func (v *NullableStringPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringPage(val *StringPage) *NullableStringPage {
	return &NullableStringPage{value: val, isSet: true}
}

func (v NullableStringPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

