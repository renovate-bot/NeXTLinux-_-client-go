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
// ContentFilesResponseContent struct for ContentFilesResponseContent
type ContentFilesResponseContent struct {
	Filename string `json:"filename,omitempty"`
	Gid int32 `json:"gid,omitempty"`
	Linkdest *string `json:"linkdest,omitempty"`
	Mode string `json:"mode,omitempty"`
	Sha256 *string `json:"sha256,omitempty"`
	Size int32 `json:"size,omitempty"`
	Type string `json:"type,omitempty"`
	Uid int32 `json:"uid,omitempty"`
}
