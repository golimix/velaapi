/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1ListAddonResponse struct for V1ListAddonResponse
type V1ListAddonResponse struct {
	Addons []V1AddonInfo `json:"addons"`
	Message string `json:"message,omitempty"`
}