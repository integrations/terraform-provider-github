package ghinstallation

import (
	gh "github.com/octokit/go-sdk/pkg/github/models"
)

// InstallationTokenOptions allow restricting a token's access to specific repositories.
// These options are taken from google/go-github's implementation.
// See more at https://github.com/google/go-github/blob/8c1232a5960307a6383998e1aa2dd71711343810/github/apps.go
type InstallationTokenOptions struct {
	// The IDs of the repositories that the installation token can access.
	// Providing repository IDs restricts the access of an installation token to specific repositories.
	RepositoryIDs []int64 `json:"repository_ids,omitempty"`

	// The names of the repositories that the installation token can access.
	// Providing repository names restricts the access of an installation token to specific repositories.
	Repositories []string `json:"repositories,omitempty"`

	// The permissions granted to the access token.
	// The permissions object includes the permission names and their access type.
	Permissions *gh.AppPermissions `json:"permissions,omitempty"`
}
