/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1ApplicationStatusResponse struct for V1ApplicationStatusResponse
type V1ApplicationStatusResponse struct {
	EnvName string `json:"envName"`
	Status CommonAppStatus `json:"status"`
}
