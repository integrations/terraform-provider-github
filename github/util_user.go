package github

import "strconv"

// userIdentity represents a GitHub user by their login.
type userIdentity struct {
	login string
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

// flattenInvitations converts the userCollaborators slice into a format suitable for Terraform schema.
func (uc userCollaborators) flattenInvitations() any {
	m := make(map[string]any)

	for _, u := range uc {
		if u.invitationID != nil {
			m[u.login] = strconv.FormatInt(*u.invitationID, 10)
		}
	}

	return m
}
