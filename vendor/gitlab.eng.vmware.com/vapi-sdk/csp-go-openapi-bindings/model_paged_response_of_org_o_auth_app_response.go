/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// PagedResponseOfOrgOAuthAppResponse struct for PagedResponseOfOrgOAuthAppResponse
type PagedResponseOfOrgOAuthAppResponse struct {
	// Relative path to previous page if exists. Not returned for POST requests
	PrevLink string `json:"prevLink,omitempty"`
	// Relative path to next page if exists. Not returned for POST requests
	NextLink string `json:"nextLink,omitempty"`
	// Total amount of results if available
	TotalResults int32 `json:"totalResults,omitempty"`
	// List of results
	Results []OrgOAuthAppResponse `json:"results,omitempty"`
}