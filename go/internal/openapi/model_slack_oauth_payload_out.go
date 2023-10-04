/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SlackOauthPayloadOut struct for SlackOauthPayloadOut
type SlackOauthPayloadOut struct {
	Error NullableString `json:"error,omitempty"`
	IncomingWebhook *IncomingWebhook `json:"incoming_webhook,omitempty"`
	Ok bool `json:"ok"`
}

// NewSlackOauthPayloadOut instantiates a new SlackOauthPayloadOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackOauthPayloadOut(ok bool) *SlackOauthPayloadOut {
	this := SlackOauthPayloadOut{}
	this.Ok = ok
	return &this
}

// NewSlackOauthPayloadOutWithDefaults instantiates a new SlackOauthPayloadOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackOauthPayloadOutWithDefaults() *SlackOauthPayloadOut {
	this := SlackOauthPayloadOut{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SlackOauthPayloadOut) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SlackOauthPayloadOut) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *SlackOauthPayloadOut) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *SlackOauthPayloadOut) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *SlackOauthPayloadOut) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *SlackOauthPayloadOut) UnsetError() {
	o.Error.Unset()
}

// GetIncomingWebhook returns the IncomingWebhook field value if set, zero value otherwise.
func (o *SlackOauthPayloadOut) GetIncomingWebhook() IncomingWebhook {
	if o == nil || o.IncomingWebhook == nil {
		var ret IncomingWebhook
		return ret
	}
	return *o.IncomingWebhook
}

// GetIncomingWebhookOk returns a tuple with the IncomingWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackOauthPayloadOut) GetIncomingWebhookOk() (*IncomingWebhook, bool) {
	if o == nil || o.IncomingWebhook == nil {
		return nil, false
	}
	return o.IncomingWebhook, true
}

// HasIncomingWebhook returns a boolean if a field has been set.
func (o *SlackOauthPayloadOut) HasIncomingWebhook() bool {
	if o != nil && o.IncomingWebhook != nil {
		return true
	}

	return false
}

// SetIncomingWebhook gets a reference to the given IncomingWebhook and assigns it to the IncomingWebhook field.
func (o *SlackOauthPayloadOut) SetIncomingWebhook(v IncomingWebhook) {
	o.IncomingWebhook = &v
}

// GetOk returns the Ok field value
func (o *SlackOauthPayloadOut) GetOk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value
// and a boolean to check if the value has been set.
func (o *SlackOauthPayloadOut) GetOkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ok, true
}

// SetOk sets field value
func (o *SlackOauthPayloadOut) SetOk(v bool) {
	o.Ok = v
}

func (o SlackOauthPayloadOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.IncomingWebhook != nil {
		toSerialize["incoming_webhook"] = o.IncomingWebhook
	}
	if true {
		toSerialize["ok"] = o.Ok
	}
	return json.Marshal(toSerialize)
}

type NullableSlackOauthPayloadOut struct {
	value *SlackOauthPayloadOut
	isSet bool
}

func (v NullableSlackOauthPayloadOut) Get() *SlackOauthPayloadOut {
	return v.value
}

func (v *NullableSlackOauthPayloadOut) Set(val *SlackOauthPayloadOut) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackOauthPayloadOut) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackOauthPayloadOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackOauthPayloadOut(val *SlackOauthPayloadOut) *NullableSlackOauthPayloadOut {
	return &NullableSlackOauthPayloadOut{value: val, isSet: true}
}

func (v NullableSlackOauthPayloadOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackOauthPayloadOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

