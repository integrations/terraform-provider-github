package github

const (
	pullPermission  string = "pull"
	pushPermission  string = "push"
	writePermission string = "write"
	readPermission  string = "read"
)

func getPermission(permission string) string {
	// Permissions for some GitHub API routes are expressed as "read",
	// "write", and "admin"; in other places, they are expressed as "pull",
	// "push", and "admin".
	switch permission {
	case readPermission:
		return pullPermission
	case writePermission:
		return pushPermission
	default:
		return permission
	}
}
