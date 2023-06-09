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
// ModelApplicationTrait struct for ModelApplicationTrait
type ModelApplicationTrait struct {
	Alias string `json:"alias"`
	CreateTime time.Time `json:"createTime"`
	Description string `json:"description"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Type string `json:"type"`
	UpdateTime time.Time `json:"updateTime"`
}
