package ghclient

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidOption is returned when a client or source option is given an invalid value.
	ErrInvalidOption = errors.New("invalid option")

	// ErrCacheInUse is returned when the cache database is locked by another provider configuration or process.
	ErrCacheInUse = errors.New("cache is already in use")

	// ErrInstallationNotFound is returned when the GitHub App has no installation for the requested owner.
	ErrInstallationNotFound = errors.New("app installation not found")
)

// errGHESBaseURLRequired is returned when GitHub Enterprise Server is enabled without a base URL.
var errGHESBaseURLRequired = fmt.Errorf("%w: base URL cannot be nil when ghes is enabled", ErrInvalidOption)
