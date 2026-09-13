package github

import (
	"fmt"
	"strconv"
)

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

// userMember represents a GitHub team member with its identity and role.
type userMember struct {
	userIdentity
	role string
}

type userMembers []userMember

// flatten converts the userMembers slice into a format suitable for Terraform schema.
func (um userMembers) flatten() any {
	items := make([]any, len(um))

	for i, u := range um {
		items[i] = map[string]any{
			"username": u.login,
			"role":     u.role,
		}
	}

	return items
}

func newUserMembers(in []any) (userMembers, error) {
	members := make(userMembers, len(in))

	for i, v := range in {
		m, ok := v.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("unexpected type for team member: %T", v)
		}

		usernameVal, ok := m["username"]
		if !ok {
			return nil, fmt.Errorf("missing username for team member: %v", m)
		}

		username, ok := usernameVal.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected type for username: %T", usernameVal)
		}

		roleVal, ok := m["role"]
		if !ok {
			return nil, fmt.Errorf("missing role for team member: %v", m)
		}

		role, ok := roleVal.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected type for role: %T", roleVal)
		}

		members[i] = userMember{
			userIdentity: userIdentity{login: username},
			role:         role,
		}
	}

	return members, nil
}
