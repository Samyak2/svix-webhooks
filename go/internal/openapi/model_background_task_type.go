/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BackgroundTaskType the model 'BackgroundTaskType'
type BackgroundTaskType string

// List of BackgroundTaskType
const (
	BACKGROUNDTASKTYPE_ENDPOINT_REPLAY BackgroundTaskType = "endpoint.replay"
	BACKGROUNDTASKTYPE_ENDPOINT_RECOVER BackgroundTaskType = "endpoint.recover"
	BACKGROUNDTASKTYPE_APPLICATION_STATS BackgroundTaskType = "application.stats"
)

var allowedBackgroundTaskTypeEnumValues = []BackgroundTaskType{
	"endpoint.replay",
	"endpoint.recover",
	"application.stats",
}

func (v *BackgroundTaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackgroundTaskType(value)
	for _, existing := range allowedBackgroundTaskTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackgroundTaskType", value)
}

// NewBackgroundTaskTypeFromValue returns a pointer to a valid BackgroundTaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackgroundTaskTypeFromValue(v string) (*BackgroundTaskType, error) {
	ev := BackgroundTaskType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackgroundTaskType: valid values are %v", v, allowedBackgroundTaskTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackgroundTaskType) IsValid() bool {
	for _, existing := range allowedBackgroundTaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackgroundTaskType value
func (v BackgroundTaskType) Ptr() *BackgroundTaskType {
	return &v
}

type NullableBackgroundTaskType struct {
	value *BackgroundTaskType
	isSet bool
}

func (v NullableBackgroundTaskType) Get() *BackgroundTaskType {
	return v.value
}

func (v *NullableBackgroundTaskType) Set(val *BackgroundTaskType) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundTaskType) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundTaskType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundTaskType(val *BackgroundTaskType) *NullableBackgroundTaskType {
	return &NullableBackgroundTaskType{value: val, isSet: true}
}

func (v NullableBackgroundTaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundTaskType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

