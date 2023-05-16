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
// Service A service status record
type Service struct {
	// The unique id of the host on which the service is executing
	Hostid string `json:"hostid,omitempty"`
	// Registered service name
	Servicename string `json:"servicename,omitempty"`
	// The url to reach the service, including port as needed
	BaseUrl string `json:"base_url,omitempty"`
	// A state indicating the condition of the service. Normal operation is 'registered'
	StatusMessage string `json:"status_message,omitempty"`
	ServiceDetail StatusResponse `json:"service_detail,omitempty"`
	Status bool `json:"status,omitempty"`
	// The version of the service as reported by the service implementation on registration
	Version string `json:"version,omitempty"`
}
