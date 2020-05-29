/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient
// OidcUserInfoDto OIDC User Info
type OidcUserInfoDto struct {
	// The context name (equals to CSP organization ID) in which the token was issued.
	ContextName string `json:"context_name,omitempty"`
	// The user's account identifier, the acct claim can be a combination of the user's username and domain in URLEncoded(username)@domain format or the user's email or the UPN of the user.
	Acct string `json:"acct,omitempty"`
	// The domain the user belongs to.
	Domain string `json:"domain,omitempty"`
	// The user's email.
	Email string `json:"email,omitempty"`
	// The user's family name.
	FamilyName string `json:"family_name,omitempty"`
	// Group names the user belongs to. Property will be returned only if the client registered with 'group_names' scope.
	GroupNames []string `json:"group_names,omitempty"`
	// The user's username.
	Username string `json:"username,omitempty"`
	// The user on behalf of which the token was issued.
	Sub string `json:"sub,omitempty"`
	// True if the user's e-mail address has been verified; Otherwise false.
	EmailVerified bool `json:"email_verified,omitempty"`
	// Group ids the user belongs to. Property will be returned only if the client registered with 'group_ids' scope.
	GroupIds []string `json:"group_ids,omitempty"`
	// The context identifier in which the token was issued.
	Context string `json:"context,omitempty"`
	// The user's given name.
	GivenName string `json:"given_name,omitempty"`
}