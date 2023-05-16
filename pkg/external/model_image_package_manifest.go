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
// ImagePackageManifest struct for ImagePackageManifest
type ImagePackageManifest struct {
	Artifacts []ImportPackage `json:"artifacts"`
	Source ImportSource `json:"source"`
	Distro ImportDistribution `json:"distro"`
	Descriptor ImportDescriptor `json:"descriptor,omitempty"`
	Schema ImportSchema `json:"schema,omitempty"`
	ArtifactRelationships []ImportPackageRelationship `json:"artifactRelationships,omitempty"`
}
