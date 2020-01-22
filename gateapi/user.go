/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type User struct {

	AccountNonExpired bool `json:"accountNonExpired,omitempty"`

	Roles []string `json:"roles,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Username string `json:"username,omitempty"`

	CredentialsNonExpired bool `json:"credentialsNonExpired,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	AllowedAccounts []string `json:"allowedAccounts,omitempty"`

	Authorities []GrantedAuthority `json:"authorities,omitempty"`

	AccountNonLocked bool `json:"accountNonLocked,omitempty"`

	Email string `json:"email,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
}
