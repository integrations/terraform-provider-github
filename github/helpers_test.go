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

func mustNewSSHKey(t *testing.T) (string, []byte) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate ed25519 key: %v", err)
	}

	sshPub, err := ssh.NewPublicKey(pub)
	if err != nil {
		t.Fatalf("failed to generate ssh public key: %v", err)
	}

	block, err := ssh.MarshalPrivateKey(priv, "") // second arg is a comment
	if err != nil {
		t.Fatalf("failed to marshal ssh private key: %v", err)
	}

	publicKey := strings.TrimSpace(string(ssh.MarshalAuthorizedKey(sshPub)))
	privateKey := pem.EncodeToMemory(block)

	return publicKey, privateKey
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
