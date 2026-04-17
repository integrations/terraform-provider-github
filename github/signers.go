package github

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

// PEMSigner signs JWTs using a local RSA private key in PEM format.
type PEMSigner struct {
	signer jose.Signer
}

// NewPEMSigner creates a PEMSigner from a PKCS1 PEM-encoded RSA private key.
func NewPEMSigner(pemData []byte) (*PEMSigner, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("no decodeable PEM data found")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: privateKey},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		return nil, err
	}

	return &PEMSigner{signer: signer}, nil
}

func (s *PEMSigner) SignJWT(_ context.Context, claims jwt.Claims) (string, error) {
	return jwt.Signed(s.signer).Claims(claims).CompactSerialize()
}
