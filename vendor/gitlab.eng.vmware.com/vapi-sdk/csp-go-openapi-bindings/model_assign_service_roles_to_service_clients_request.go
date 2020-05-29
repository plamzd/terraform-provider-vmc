/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// AssignServiceRolesToServiceClientsRequest struct for AssignServiceRolesToServiceClientsRequest
type AssignServiceRolesToServiceClientsRequest struct {
	// The service definition id (without \"external\" prefix
	ServiceDefinitionId string `json:"serviceDefinitionId,omitempty"`
	// The service roles
	ServiceRoles []ServiceRoleBindingDto `json:"serviceRoles,omitempty"`
	// Service Client Ids
	ClientIds []string `json:"clientIds"`
}