/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IsoCreate struct {
	Name string `json:"name"`
	Size float32 `json:"size,omitempty"`
	Source string `json:"source,omitempty"`
	Service string `json:"service,omitempty"`
	Cloud string `json:"cloud,omitempty"`
	Metadata DiskMetadata `json:"metadata,omitempty"`
	Tag map[string]interface{} `json:"tag,omitempty"`
}
