/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1CreateRoleRequest struct for V1CreateRoleRequest
type V1CreateRoleRequest struct {
	Alias string `json:"alias"`
	Name string `json:"name"`
	Permissions []string `json:"permissions"`
}
