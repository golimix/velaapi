/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
import (
	"time"
)
// V1TerraformProvider struct for V1TerraformProvider
type V1TerraformProvider struct {
	CreateTime time.Time `json:"createTime"`
	Name string `json:"name"`
	Provider string `json:"provider"`
	Region string `json:"region"`
}