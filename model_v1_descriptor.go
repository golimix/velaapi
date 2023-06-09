/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1Descriptor struct for V1Descriptor
type V1Descriptor struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	Data string `json:"data,omitempty"`
	Digest string `json:"digest"`
	MediaType string `json:"mediaType"`
	Platform V1Platform `json:"platform,omitempty"`
	Size int64 `json:"size"`
	Urls []string `json:"urls,omitempty"`
}
