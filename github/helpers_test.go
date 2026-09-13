package github

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/url"
	"strings"
	"testing"

	"golang.org/x/crypto/ssh"
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

// mustNewSshPublicKey returns a freshly generated public key in the authorized
// keys format accepted by the GitHub SSH key and SSH signing key endpoints.
func mustNewSshPublicKey(t *testing.T) string {
	t.Helper()

	pub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate Ed25519 key: %v", err)
	}

	sshPub, err := ssh.NewPublicKey(pub)
	if err != nil {
		t.Fatalf("failed to convert Ed25519 key to SSH public key: %v", err)
	}

	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(sshPub)))
}
