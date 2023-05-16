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
// VulnUpdateNotificationPayload struct for VulnUpdateNotificationPayload
type VulnUpdateNotificationPayload struct {
	UserId string `json:"userId,omitempty"`
	SubscriptionKey string `json:"subscription_key,omitempty"`
	SubscriptionType string `json:"subscription_type,omitempty"`
	NotificationId string `json:"notificationId,omitempty"`
	DiffVulnerabilityResult VulnDiffResult `json:"diff_vulnerability_result,omitempty"`
	ImageDigest string `json:"imageDigest,omitempty"`
	// List of Corresponding Image Annotations
	Annotations *interface{} `json:"annotations,omitempty"`
}
