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
// ModelImageResource struct for ModelImageResource
type ModelImageResource struct {
	CreateTime time.Time `json:"createTime,omitempty"`
	Digest string `json:"digest"`
	Tag string `json:"tag"`
	Url string `json:"url"`
}
