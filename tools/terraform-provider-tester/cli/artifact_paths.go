package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func confinedArtifactPath(providerRoot, destination, description string) (string, error) {
	if providerRoot == "" {
		return "", fmt.Errorf("%s destination cannot be checked without provider root", description)
	}
	root, err := filepath.Abs(providerRoot)
	if err != nil {
		return "", fmt.Errorf("resolving provider root: %w", err)
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		return "", fmt.Errorf("resolving provider root symlinks: %w", err)
	}
	dest, err := filepath.Abs(destination)
	if err != nil {
		return "", fmt.Errorf("resolving %s destination: %w", description, err)
	}
	realDest, err := evalPathThroughExistingAncestor(dest)
	if err != nil {
		return "", fmt.Errorf("resolving %s destination symlinks: %w", description, err)
	}
	if sameOrDescendant(root, realDest) {
		return "", fmt.Errorf("%s destination %s resolves inside provider root %s; set XDG_CONFIG_HOME or HOME outside the provider checkout", description, dest, root)
	}
	return filepath.Clean(destination), nil
}

func evalPathThroughExistingAncestor(path string) (string, error) {
	clean := filepath.Clean(path)
	existing := clean
	var suffix []string
	for {
		if _, err := os.Lstat(existing); err == nil {
			break
		} else if !os.IsNotExist(err) {
			return "", err
		}
		parent := filepath.Dir(existing)
		if parent == existing {
			break
		}
		suffix = append([]string{filepath.Base(existing)}, suffix...)
		existing = parent
	}
	realExisting, err := filepath.EvalSymlinks(existing)
	if err != nil {
		return "", err
	}
	parts := append([]string{realExisting}, suffix...)
	return filepath.Clean(filepath.Join(parts...)), nil
}

func sameOrDescendant(root, path string) bool {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return false
	}
	return rel == "." || (rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)))
}
