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
// V1ApplicationTemplateVersion struct for V1ApplicationTemplateVersion
type V1ApplicationTemplateVersion struct {
	CreateTime time.Time `json:"createTime"`
	CreateUser string `json:"createUser"`
	Description string `json:"description"`
	UpdateTime time.Time `json:"updateTime"`
	Version string `json:"version"`
}