/*
 * vela-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vealapi
// V1ClusterResourceInfo struct for V1ClusterResourceInfo
type V1ClusterResourceInfo struct {
	CpuCapacity int64 `json:"cpuCapacity"`
	CpuUsed int64 `json:"cpuUsed"`
	GpuCapacity int64 `json:"gpuCapacity,omitempty"`
	GpuUsed int64 `json:"gpuUsed,omitempty"`
	MasterNumber int32 `json:"masterNumber"`
	MemoryCapacity int64 `json:"memoryCapacity"`
	MemoryUsed int64 `json:"memoryUsed"`
	PodCapacity int64 `json:"podCapacity"`
	PodUsed int64 `json:"podUsed"`
	StorageClassList []string `json:"storageClassList,omitempty"`
	WorkerNumber int32 `json:"workerNumber"`
}