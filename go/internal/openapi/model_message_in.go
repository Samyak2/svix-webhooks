/*
 * Svix
 *
 * The Svix server API documentation
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MessageIn struct for MessageIn
type MessageIn struct {
	EventType string `json:"eventType"`
	EventId *string `json:"eventId,omitempty"`
	Data map[string]interface{} `json:"data"`
}

// NewMessageIn instantiates a new MessageIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageIn(eventType string, data map[string]interface{}, ) *MessageIn {
	this := MessageIn{}
	this.EventType = eventType
	this.Data = data
	return &this
}

// NewMessageInWithDefaults instantiates a new MessageIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageInWithDefaults() *MessageIn {
	this := MessageIn{}
	return &this
}

// GetEventType returns the EventType field value
func (o *MessageIn) GetEventType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MessageIn) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MessageIn) SetEventType(v string) {
	o.EventType = v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *MessageIn) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageIn) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *MessageIn) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *MessageIn) SetEventId(v string) {
	o.EventId = &v
}

// GetData returns the Data field value
func (o *MessageIn) GetData() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MessageIn) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MessageIn) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o MessageIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMessageIn struct {
	value *MessageIn
	isSet bool
}

func (v NullableMessageIn) Get() *MessageIn {
	return v.value
}

func (v *NullableMessageIn) Set(val *MessageIn) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageIn) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageIn(val *MessageIn) *NullableMessageIn {
	return &NullableMessageIn{value: val, isSet: true}
}

func (v NullableMessageIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

