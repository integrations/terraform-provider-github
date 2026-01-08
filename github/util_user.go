package github

import "fmt"

// userWithLogin is an interface representing a GitHub user that has a login.
type userWithLogin interface {
	getLogin() string
}

// userIdentity represents a GitHub user by their login.
type userIdentity struct {
	login string
}

// getLogin returns the login of the user.
func (u userIdentity) getLogin() string {
	return u.login
}

// userCollaborator represents a GitHub user collaborator with its identity and permission level.
type userCollaborator struct {
	userIdentity
	permission   string
	invitationID *int64
}

// flatten converts the userCollaborator into a format suitable for Terraform schema.
func (u userCollaborator) flatten() any {
	m := map[string]any{
		"username":   u.login,
		"permission": u.permission,
	}

	if u.invitationID != nil {
		m["invitation_id"] = *u.invitationID
	}

	return m
}

// userCollaborators is a slice of userCollaborator.
type userCollaborators []userCollaborator

// flatten converts the userCollaborators slice into a format suitable for Terraform schema.
func (uc userCollaborators) flatten() any {
	items := make([]any, len(uc))

	for i, u := range uc {
		items[i] = u.flatten()
	}

	return items
}

// getUserCollaborators converts a slice of any type to a slice of userCollaborator.
func getUserCollaborators(col []any) (userCollaborators, error) {
	collaborators := make([]userCollaborator, len(col))

	for i, u := range col {
		m, ok := u.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("input invalid")
		}

		n, ok := m["username"]
		if !ok {
			return nil, fmt.Errorf("username missing")
		}

		username, ok := n.(string)
		if !ok || len(username) == 0 {
			return nil, fmt.Errorf("username invalid")
		}

		p, ok := m["permission"]
		if !ok {
			return nil, fmt.Errorf("permission missing")
		}

		permission, ok := p.(string)
		if !ok || len(permission) == 0 {
			return nil, fmt.Errorf("permission invalid")
		}

		uc := userCollaborator{
			userIdentity: userIdentity{
				login: username,
			},
			permission: permission,
		}

		collaborators[i] = uc
	}

	return collaborators, nil
}

// checkDuplicateUsers checks for duplicate usernames in a slice of userWithLogin.
func checkDuplicateUsers[T userWithLogin](users []T) bool {
	seen := make(map[string]any)

	for _, u := range users {
		login := u.getLogin()
		if _, ok := seen[login]; ok {
			return true
		}
		seen[login] = struct{}{}
	}

	return false
}
