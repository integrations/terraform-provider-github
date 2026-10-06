package github

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/url"
	"testing"
)

func mustNewURL(t *testing.T, s string) *url.URL {
	t.Helper()

	url, err := url.Parse(s)
	if err != nil {
		t.Fatalf("failed to parse test base URL: %s", err.Error())
	}

	return url
}

func mustNewPEM(t *testing.T) []byte {
	t.Helper()

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
}
