/*
 * Moreu API
 *
 * This is a moreu server.
 *
 * API version: 1.0.0
 * Contact: saltbo@foxmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type HttputilJsonResponse struct {
	Code int32 `json:"code,omitempty"`
	Data *interface{} `json:"data,omitempty"`
	Msg string `json:"msg,omitempty"`
}
