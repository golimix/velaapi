/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1Platform struct for V1Platform
type V1Platform struct {
	Architecture string `json:"architecture"`
	Features []string `json:"features,omitempty"`
	Os string `json:"os"`
	OsFeatures []string `json:"os.features,omitempty"`
	OsVersion string `json:"os.version,omitempty"`
	Variant string `json:"variant,omitempty"`
}
