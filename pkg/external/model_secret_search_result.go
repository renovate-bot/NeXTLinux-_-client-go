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
// SecretSearchResult The retrieved file entry including content (b64 encoded)
type SecretSearchResult struct {
	Path string `json:"path,omitempty"`
	Matches []RegexContentMatch `json:"matches,omitempty"`
}
