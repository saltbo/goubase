/*
 * Moreu API
 *
 * This is a moreu server.
 *
 * API version: 1.0
 * Contact: saltbo@foxmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type BindQueryUser struct {
	Limit int32 `json:"limit,omitempty"`
	Name string `json:"name,omitempty"`
	Offset int32 `json:"offset,omitempty"`
}
