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
// RegistryTagSource An image reference using a tag in a registry, this is the most common source type.
type RegistryTagSource struct {
	// A docker pull string (e.g. docker.io/nginx:latest, or docker.io/nginx@sha256:abd) to retrieve the image
	Pullstring string `json:"pullstring"`
	// Base64 encoded content of the dockerfile used to build the image, if available.
	Dockerfile string `json:"dockerfile,omitempty"`
}