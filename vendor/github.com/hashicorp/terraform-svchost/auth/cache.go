// Copyright (c) HashiCorp, Inc.

package auth

import (
	"sync"

	svchost "github.com/hashicorp/terraform-svchost"
)

// CachingCredentialsSource creates a new credentials source that wraps another
// and caches its results in memory, on a per-hostname basis.
//
// No means is provided for expiration of cached credentials, so a caching
// credentials source should have a limited lifetime (one Terraform operation,
// for example) to ensure that time-limited credentials don't expire before
// their cache entries do.
func CachingCredentialsSource(source CredentialsSource) CredentialsSource {
	return &cachingCredentialsSource{
		source: source,
		cache:  map[svchost.Hostname]HostCredentials{},
	}
}

type cachingCredentialsSource struct {
	source CredentialsSource
	cache  map[svchost.Hostname]HostCredentials
	mu     sync.Mutex
}

// ForHost passes the given hostname on to the wrapped credentials source and
// caches the result to return for future requests with the same hostname.
//
// Both credentials and non-credentials (nil) responses are cached.
//
// No cache entry is created if the wrapped source returns an error, to allow
// the caller to retry the failing operation.
func (s *cachingCredentialsSource) ForHost(host svchost.Hostname) (HostCredentials, error) {
	s.mu.Lock()
	if cache, cached := s.cache[host]; cached {
		s.mu.Unlock()
		return cache, nil
	}
	s.mu.Unlock()

	result, err := s.source.ForHost(host)
	if err != nil {
		return result, err
	}

	s.mu.Lock()
	s.cache[host] = result
	s.mu.Unlock()
	return result, nil
}

func (s *cachingCredentialsSource) StoreForHost(host svchost.Hostname, credentials HostCredentialsWritable) error {
	// We'll delete the cache entry even if the store fails, since that just
	// means that the next read will go to the real store and get a chance to
	// see which object (old or new) is actually present.
	s.mu.Lock()
	delete(s.cache, host)
	s.mu.Unlock()
	return s.source.StoreForHost(host, credentials)
}

func (s *cachingCredentialsSource) ForgetForHost(host svchost.Hostname) error {
	// We'll delete the cache entry even if the store fails, since that just
	// means that the next read will go to the real store and get a chance to
	// see if the object is still present.
	s.mu.Lock()
	delete(s.cache, host)
	s.mu.Unlock()
	return s.source.ForgetForHost(host)
}
