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
// NvdDataObject struct for NvdDataObject
type NvdDataObject struct {
	// NVD Vulnerability ID
	Id string `json:"id,omitempty"`
	CvssV2 Cvssv2Scores `json:"cvss_v2,omitempty"`
	CvssV3 Cvssv3Scores `json:"cvss_v3,omitempty"`
}