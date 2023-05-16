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
import (
	"time"
)
// AnalysisArchiveTransitionRule A rule for auto-archiving image analysis by time and/or tag-history
type AnalysisArchiveTransitionRule struct {
	Selector ImageSelector `json:"selector,omitempty"`
	// Number of images mapped to the tag that are newer
	TagVersionsNewer int32 `json:"tag_versions_newer,omitempty"`
	// Matches if the analysis is strictly older than this number of days
	AnalysisAgeDays int32 `json:"analysis_age_days,omitempty"`
	// The type of transition to make. If \"archive\", then archive an image from the working set and remove it from the working set. If \"delete\", then match against archived images and delete from the archive if match.
	Transition string `json:"transition"`
	// True if the rule applies to all accounts in the system. This is only available to admin users to update/modify, but all users with permission to list rules can see them
	SystemGlobal bool `json:"system_global,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	Exclude AnalysisArchiveTransitionRuleExclude `json:"exclude,omitempty"`
	// This is the maximum number of image analyses an account can have. Can only be set on system_global rules
	MaxImagesPerAccount int32 `json:"max_images_per_account,omitempty"`
}
