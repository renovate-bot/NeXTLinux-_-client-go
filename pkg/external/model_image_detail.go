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
// ImageDetail A metadata detail record for a specific image. Multiple detail records may map a single catalog image.
type ImageDetail struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	// Full docker-pullable tag string referencing the image
	Fulltag string `json:"fulltag,omitempty"`
	// Full docker-pullable digest string including the registry url and repository necessary get the image
	Fulldigest string `json:"fulldigest,omitempty"`
	UserId string `json:"userId,omitempty"`
	ImageId string `json:"imageId,omitempty"`
	Registry string `json:"registry,omitempty"`
	Repo string `json:"repo,omitempty"`
	Dockerfile *string `json:"dockerfile,omitempty"`
	// The parent Nextlinux Image record to which this detail maps
	ImageDigest string `json:"imageDigest,omitempty"`
}
