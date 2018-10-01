/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

type MapTest struct {
	MapMapOfString  map[string]map[string]string `json:"map_map_of_string,omitempty"`
	MapOfEnumString map[string]string            `json:"map_of_enum_string,omitempty"`
	DirectMap       map[string]bool              `json:"direct_map,omitempty"`
	IndirectMap     map[string]bool              `json:"indirect_map,omitempty"`
}
