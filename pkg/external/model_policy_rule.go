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
// PolicyRule A rule that defines and decision value if the match is found true for a given image.
type PolicyRule struct {
	Id string `json:"id,omitempty"`
	Gate string `json:"gate"`
	Trigger string `json:"trigger"`
	Action string `json:"action"`
	Params []PolicyRuleParams `json:"params,omitempty"`
}