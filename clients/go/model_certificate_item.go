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

// CertificateItem struct for CertificateItem
type CertificateItem struct {
	Id string `json:"id"`
	Name string `json:"name"`
	PublicKey string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Channels []ChannelItem `json:"channels"`
}

// NewCertificateItem instantiates a new CertificateItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateItem(id string, name string, publicKey string, privateKey string, channels []ChannelItem) *CertificateItem {
	this := CertificateItem{}
	this.Id = id
	this.Name = name
	this.PublicKey = publicKey
	this.PrivateKey = privateKey
	this.Channels = channels
	return &this
}

// NewCertificateItemWithDefaults instantiates a new CertificateItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateItemWithDefaults() *CertificateItem {
	this := CertificateItem{}
	return &this
}

// GetId returns the Id field value
func (o *CertificateItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CertificateItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CertificateItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CertificateItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CertificateItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CertificateItem) SetName(v string) {
	o.Name = v
}

// GetPublicKey returns the PublicKey field value
func (o *CertificateItem) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *CertificateItem) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *CertificateItem) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *CertificateItem) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *CertificateItem) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *CertificateItem) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetChannels returns the Channels field value
func (o *CertificateItem) GetChannels() []ChannelItem {
	if o == nil {
		var ret []ChannelItem
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *CertificateItem) GetChannelsOk() ([]ChannelItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *CertificateItem) SetChannels(v []ChannelItem) {
	o.Channels = v
}

func (o CertificateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["publicKey"] = o.PublicKey
	}
	if true {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if true {
		toSerialize["channels"] = o.Channels
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateItem struct {
	value *CertificateItem
	isSet bool
}

func (v NullableCertificateItem) Get() *CertificateItem {
	return v.value
}

func (v *NullableCertificateItem) Set(val *CertificateItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateItem(val *CertificateItem) *NullableCertificateItem {
	return &NullableCertificateItem{value: val, isSet: true}
}

func (v NullableCertificateItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


