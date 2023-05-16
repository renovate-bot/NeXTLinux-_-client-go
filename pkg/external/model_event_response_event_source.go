/*
 * Nextlinux Engine API Server
 *
 * This is the Nextlinux Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.16
 * Contact: nurmi@nextlinux.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
// EventResponseEventSource struct for EventResponseEventSource
type EventResponseEventSource struct {
	Servicename string `json:"servicename,omitempty"`
	Hostid string `json:"hostid,omitempty"`
	BaseUrl string `json:"base_url,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}
