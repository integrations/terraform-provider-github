package github

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	kmstypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
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

// AWSKMSClient abstracts the AWS KMS Sign API for testability.
type AWSKMSClient interface {
	Sign(ctx context.Context, params *kms.SignInput, optFns ...func(*kms.Options)) (*kms.SignOutput, error)
}

// AWSKMSSigner signs JWTs by delegating to AWS KMS.
// The private key never leaves the KMS boundary.
type AWSKMSSigner struct {
	client AWSKMSClient
	keyID  string
}

// NewAWSKMSSigner creates an AWSKMSSigner using the default AWS credential chain.
func NewAWSKMSSigner(ctx context.Context, keyID string) (*AWSKMSSigner, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}
	return &AWSKMSSigner{
		client: kms.NewFromConfig(cfg),
		keyID:  keyID,
	}, nil
}

func (s *AWSKMSSigner) SignJWT(ctx context.Context, claims jwt.Claims) (string, error) {
	headerJSON, err := json.Marshal(map[string]string{"alg": "RS256", "typ": "JWT"})
	if err != nil {
		return "", err
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	signingInput := base64.RawURLEncoding.EncodeToString(headerJSON) + "." +
		base64.RawURLEncoding.EncodeToString(claimsJSON)

	digest := sha256.Sum256([]byte(signingInput))

	out, err := s.client.Sign(ctx, &kms.SignInput{
		KeyId:            aws.String(s.keyID),
		Message:          digest[:],
		MessageType:      kmstypes.MessageTypeDigest,
		SigningAlgorithm: kmstypes.SigningAlgorithmSpecRsassaPkcs1V15Sha256,
	})
	if err != nil {
		return "", fmt.Errorf("KMS signing failed: %w", err)
	}

	return signingInput + "." + base64.RawURLEncoding.EncodeToString(out.Signature), nil
}
