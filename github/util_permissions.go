package github

import (
	"fmt"

	"github.com/google/go-github/v45/github"
)

const (
	pullPermission     string = "pull"
	triagePermission   string = "triage"
	pushPermission     string = "push"
	maintainPermission string = "maintain"
	adminPermission    string = "admin"
	writePermission    string = "write"
	readPermission     string = "read"
)

func getInvitationPermission(i *github.RepositoryInvitation) (string, error) {
	// Permissions for some GitHub API routes are expressed as "read",
	// "write", and "admin"; in other places, they are expressed as "pull",
	// "push", and "admin".
	permissions := i.GetPermissions()
	if permissions == readPermission {
		return pullPermission, nil
	} else if permissions == writePermission {
		return pushPermission, nil
	} else if permissions == adminPermission {
		return adminPermission, nil
	} else if *i.Permissions == maintainPermission {
		return maintainPermission, nil
	} else if *i.Permissions == triagePermission {
		return triagePermission, nil
	}

	return "", fmt.Errorf("unexpected permission value: %v", permissions)
}
