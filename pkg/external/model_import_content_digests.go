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
// ImportContentDigests Digest of content to use in the final import
type ImportContentDigests struct {
	// Digest to use for the packages content
	Packages string `json:"packages"`
	// Digest for reference content for image config
	ImageConfig string `json:"image_config"`
	// Digest to reference content for the image manifest
	Manifest string `json:"manifest"`
	// Digest for reference content for parent manifest
	ParentManifest string `json:"parent_manifest,omitempty"`
	// Digest for reference content for dockerfile
	Dockerfile string `json:"dockerfile,omitempty"`
}
